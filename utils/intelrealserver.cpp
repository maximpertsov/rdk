// g++ -std=c++17 intelreal.cpp -lrealsense2 -lhttpserver

#include <iostream>
#include <thread>
#include <vector>
#include <chrono>

#include <httpserver.hpp>

#include <librealsense2/rs.hpp>

class CameraOutput {
public:
    int width;
    int height;
    std::string ppmdata;
    std::vector<uint64_t> depth;
};

std::shared_ptr<CameraOutput> CameraOutputInstance;
long long numRequests = 0; // so we can throttle and not kill cpu

std::string my_write_ppm(const char *pixels, int x, int y, int bytes_per_pixel) {
    std::stringbuf buffer;
    std::ostream os (&buffer);

    os << "P6\n" << x << " " << y << "\n255\n";
    buffer.sputn((const char*)pixels, x * y * bytes_per_pixel);

    return buffer.str();
}

#define STB_IMAGE_WRITE_IMPLEMENTATION
#include "stb_image_write.h"

void cameraThread() {
    // Create a Pipeline - this serves as a top-level API for streaming and processing frames
    rs2::pipeline p;
    
    rs2::config cfg;
    cfg.enable_stream(RS2_STREAM_DEPTH);
    cfg.enable_stream(RS2_STREAM_COLOR);
    auto profile = p.start(cfg);

    auto sensor = profile.get_device().first<rs2::depth_sensor>();

    rs2::align align_to_color(RS2_STREAM_COLOR);
    
    long long prevNumRequests = 0; 
    
    while (true) {

        auto start = std::chrono::high_resolution_clock::now();

        std::shared_ptr<CameraOutput> output(new CameraOutput());
        // get next set of frames
        rs2::frameset frames = p.wait_for_frames();


        
        // make sure depth frame is aligned to color frame
        frames = align_to_color.process(frames);

        // do color frame
        auto vf = frames.get_color_frame();

        assert( vf.get_bytes_per_pixel() == 3);
        assert( vf.get_stride_in_bytes() == ( vf.get_width() * vf.get_bytes_per_pixel() ) );

        output->width = vf.get_width();
        output->height = vf.get_height();
        output->ppmdata = my_write_ppm((const char *)vf.get_data(),
                                       vf.get_width(),
                                       vf.get_height(),
                                       vf.get_bytes_per_pixel());
        

        // create depth map
        auto depth = frames.get_depth_frame();
        output->depth.push_back(depth.get_width());
        output->depth.push_back(depth.get_height());
        for ( auto x = 0; x < depth.get_width(); x++ ) {
            for ( auto y = 0; y < depth.get_height(); y++ ) {
                uint64_t d = 1000 * depth.get_distance(x, y);
                output->depth.push_back(d);
            }
        }
        
        // replace output data with this
        CameraOutputInstance = output;

        auto finish = std::chrono::high_resolution_clock::now();
        std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(finish-start).count() << "ms\n";
        
        if (numRequests == prevNumRequests) {
            sleep(1);
        } else {
            numRequests = prevNumRequests;
        }
    }
}

using namespace httpserver;

class hello_world_resource : public http_resource {
public:
    const std::shared_ptr<http_response> render(const http_request&) {
        return std::shared_ptr<http_response>(new string_response("Hello, World!\n"));
    }
};

class picture_resource : public http_resource {
public:
    const std::shared_ptr<http_response> render(const http_request&) {
        numRequests++;
        return std::shared_ptr<http_response>(new string_response(CameraOutputInstance->ppmdata, 200, "image/ppm"));
    }
};

class picture_resource_png : public http_resource {
public:
    const std::shared_ptr<http_response> render(const http_request&) {
        std::shared_ptr<CameraOutput> mine(CameraOutputInstance);
        const char * raw_data = mine->ppmdata.c_str();
        int len = mine->width * mine->height * 3;
        raw_data = raw_data + (mine->ppmdata.size() - len);

        int pngLen;
        auto out = stbi_write_png_to_mem((const unsigned char*)raw_data, mine->width * 3, mine->width, mine->height, 3, &pngLen);
        std::string s((char*)out, pngLen);
        STBIW_FREE(out);

        return std::shared_ptr<http_response>(new string_response(s, 200, "image/png"));
    }
};

class depth_resource : public http_resource {
public:
    const std::shared_ptr<http_response> render(const http_request&) {
        numRequests++;
        std::string s((char*)CameraOutputInstance->depth.data(), CameraOutputInstance->depth.size() * 8);
        return std::shared_ptr<http_response>(new string_response(s, 200, "binary"));
    }
};


int main(int argc, char** argv) {

    std::thread t(cameraThread);
    
    webserver ws = create_webserver(8181);

    hello_world_resource hwr;
    ws.register_resource("/", &hwr);

    picture_resource pr;
    ws.register_resource("/pic.ppm", &pr);

    picture_resource_png pr2;
    ws.register_resource("/pic.png", &pr2);

    depth_resource dr;
    ws.register_resource("/depth.dat", &dr);

    std::cout << "Starting to listen" << std::endl;
    
    ws.start(true);
    
    return 0;
}
