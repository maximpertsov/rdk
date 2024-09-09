package resource

import (
	"cmp"
	"slices"
	"testing"
	"time"

	"go.viam.com/test"
)

// NewUnimplementedResource returns a resource that has all methods
// unimplemented.
func NewUnimplementedResource(name Name) Resource {
	return &unimplResource{Named: name.AsNamed()}
}

type unimplResource struct {
	Named
	AlwaysRebuild
	TriviallyCloseable
}

// NewResourceNameSet returns a flattened set of name strings from
// a collection of Name objects for the purposes of comparison
// in automated tests.
func NewResourceNameSet(resourceNames ...Name) map[Name]struct{} {
	set := make(map[Name]struct{}, len(resourceNames))
	for _, val := range resourceNames {
		set[val] = struct{}{}
	}
	return set
}

// newSortedResourceNames returns a new slice of resources names sorted by each
// resource's fully-qualified names for the purposes of comparison in automated tests.
func newSortedResourceNames(resourceNames []Name) []Name {
	sorted := make([]Name, len(resourceNames))
	copy(sorted, resourceNames)
	slices.SortStableFunc(sorted, func(r1, r2 Name) int {
		return cmp.Compare(r1.String(), r2.String())
	})
	return sorted
}

// VerifySameResourceNames asserts that two slices of Names contain the same
// resources.Names without considering order.
func VerifySameResourceNames(tb testing.TB, actual, expected []Name) {
	tb.Helper()

	test.That(tb, newSortedResourceNames(actual), test.ShouldResemble, newSortedResourceNames(expected))
}

// VerifySameResourceStatuses asserts that two slices of [Status] contain the
// same elements without considering order. Does not consider
// [Status.LastUpdated] timestamps when comparing.
func VerifySameResourceStatuses(tb testing.TB, actual, expected []Status) {
	tb.Helper()

	sortedActual := newSortedResourceStatuses(actual)
	sortedExpected := newSortedResourceStatuses(expected)

	for i := range sortedActual {
		sortedActual[i].LastUpdated = time.Time{}
	}
	for i := range sortedExpected {
		sortedExpected[i].LastUpdated = time.Time{}
	}

	test.That(tb, sortedActual, test.ShouldResemble, sortedExpected)
}

func newSortedResourceStatuses(resourceStatuses []Status) []Status {
	sorted := make([]Status, len(resourceStatuses))
	copy(sorted, resourceStatuses)
	slices.SortStableFunc(sorted, func(r1, r2 Status) int {
		return cmp.Compare(r1.Name.String(), r2.Name.String())
	})
	return sorted
}

// ExtractNames takes a slice of Name objects
// and returns a slice of name strings for the purposes of comparison
// in automated tests.
func ExtractNames(values ...Name) []string {
	var names []string
	for _, n := range values {
		names = append(names, n.ShortName())
	}
	return names
}

// SubtractNames removes values from the first slice of resource names.
func SubtractNames(from []Name, values ...Name) []Name {
	difference := make([]Name, 0, len(from))
	for _, n := range from {
		var found bool
		for _, v := range values {
			if n == v {
				found = true
				break
			}
		}
		if found {
			continue
		}
		difference = append(difference, n)
	}
	return difference
}

// VerifyTopologicallySortedLevels verifies each topological layer of a sort against the given levels from
// most dependencies to least dependencies.
func VerifyTopologicallySortedLevels(t *testing.T, g *Graph, levels [][]Name, exclusions ...Name) {
	sorted := g.TopologicalSortInLevels()
	sorted = SubtractNamesFromLevels(sorted, exclusions...)
	test.That(t, sorted, test.ShouldHaveLength, len(levels))

	for idx, level := range levels {
		t.Log("checking level", idx)
		test.That(t, sorted[idx], test.ShouldHaveLength, len(level))
		test.That(t, NewResourceNameSet(sorted[idx]...), test.ShouldResemble, NewResourceNameSet(level...))
	}
}

// SubtractNamesFromLevels removes values from each slice of resource names.
func SubtractNamesFromLevels(from [][]Name, values ...Name) [][]Name {
	differences := make([][]Name, 0, len(from))
	for _, names := range from {
		differences = append(differences, SubtractNames(names, values...))
	}
	return differences
}

// ConcatResourceNames takes a slice of slices of Name objects
// and returns a concatenated slice of Name for the purposes of comparison
// in automated tests.
func ConcatResourceNames(values ...[]Name) []Name {
	var rNames []Name
	for _, v := range values {
		rNames = append(rNames, v...)
	}
	return rNames
}

// ConcatResourceStatuses takes a slice of slices of Status objects and returns
// a concatenated slice of Status for the purposes of comparison in automated
// tests.
func ConcatResourceStatuses(values ...[]Status) []Status {
	var rs []Status
	for _, v := range values {
		rs = append(rs, v...)
	}
	return rs
}

// AddSuffixes takes a slice of Name objects and for each suffix,
// adds the suffix to every object, then returns the entire list.
func AddSuffixes(values []Name, suffixes ...string) []Name {
	var rNames []Name

	for _, s := range suffixes {
		for _, v := range values {
			newName := NewName(v.API, v.Name+s)
			rNames = append(rNames, newName)
		}
	}
	return rNames
}

// AddRemotes takes a slice of Name objects and for each remote,
// adds the remote to every object, then returns the entire list.
func AddRemotes(values []Name, remotes ...string) []Name {
	var rNames []Name

	for _, s := range remotes {
		for _, v := range values {
			v = v.PrependRemote(s)
			rNames = append(rNames, v)
		}
	}
	return rNames
}
