// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SkipGramOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsSkipGramOptions(buf []byte, offset flatbuffers.UOffsetT) *SkipGramOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SkipGramOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSkipGramOptions(buf []byte, offset flatbuffers.UOffsetT) *SkipGramOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SkipGramOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SkipGramOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SkipGramOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SkipGramOptions) NgramSize() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SkipGramOptions) MutateNgramSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *SkipGramOptions) MaxSkipSize() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SkipGramOptions) MutateMaxSkipSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *SkipGramOptions) IncludeAllNgrams() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SkipGramOptions) MutateIncludeAllNgrams(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func SkipGramOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func SkipGramOptionsAddNgramSize(builder *flatbuffers.Builder, ngramSize int32) {
	builder.PrependInt32Slot(0, ngramSize, 0)
}
func SkipGramOptionsAddMaxSkipSize(builder *flatbuffers.Builder, maxSkipSize int32) {
	builder.PrependInt32Slot(1, maxSkipSize, 0)
}
func SkipGramOptionsAddIncludeAllNgrams(builder *flatbuffers.Builder, includeAllNgrams bool) {
	builder.PrependBoolSlot(2, includeAllNgrams, false)
}
func SkipGramOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
