package march17

import (
	"encoding/hex"
	"math"
	"sort"
	"testing"
	"testing/quick"

	. "github.com/onsi/gomega"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sort[T constraints.Ordered](a []T) {
	for i := 0; i < len(a)-1; i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// Standard unit approach -----------------------------------------------------

func TestSortStandard(t *testing.T) {
	RegisterTestingT(t)

	t.Run("it does not panic if arg is empty", func(t *testing.T) {
		Sort([]int{})
	})
	t.Run("it sorts 2 reversed elements", func(t *testing.T) {
		fixture := []int{2, 1}
		Sort(fixture)
		Expect(fixture).To(Equal([]int{1, 2}))
	})
}

// Invariant ------------------------------------------------------------------

func TestSortWithInvariant(t *testing.T) {
	RegisterTestingT(t)

	itSorts := func(inputFromFW []int) bool {
		cpy := slices.Clone(inputFromFW)
		Sort(cpy)

		return slices.IsSorted(cpy)
	}

	Expect(quick.Check(itSorts, nil)).To(Succeed())
}

// Model  ---------------------------------------------------------------------

func MaxModel(xs []int) int {
	cpy := slices.Clone(xs)

	if len(cpy) == 0 {
		return 0
	}

	// Inefficient model
	slices.Sort(cpy)
	return cpy[len(cpy)-1]
}

func Max(xs []int) int {
	if len(xs) == 0 {
		return 0
	}

	max := math.MinInt
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}

func TestMax(t *testing.T) {
	RegisterTestingT(t)

	Expect(quick.CheckEqual(MaxModel, Max, &quick.Config{MaxCount: 1000})).To(Succeed())
}

// Oracle Model ---------------------------------------------------------------

func SortedOracleModel(xs []int) []int {
	cpy := slices.Clone(xs)
	sort.Ints(cpy)
	return cpy
}

func Sorted(xs []int) []int {
	cpy := slices.Clone(xs)
	Sort(cpy)
	return cpy
}

func TestSortedWithOracleModel(t *testing.T) {
	RegisterTestingT(t)

	Expect(quick.CheckEqual(SortedOracleModel, Sorted, nil)).To(Succeed())
}

// Symmetric tests ------------------------------------------------------------

func TestHexCodec(t *testing.T) {
	RegisterTestingT(t)

	itEncodesAndDecodes := func(xs []byte) bool {
		result, err := hex.DecodeString(hex.EncodeToString(xs))
		if err != nil {
			t.Logf("%v", err)
			return false
		}

		return slices.Equal(xs, result)
	}

	Expect(quick.Check(itEncodesAndDecodes, nil)).To(Succeed())
}

// Symmetric tests ------------------------------------------------------------

func Reversed[T constraints.Ordered](xs []T) []T {
	buffer := slices.Clone(xs)
	i := 0
	j := len(buffer) - 1
	for i < j {
		buffer[i], buffer[j] = buffer[j], buffer[i]
		i++
		j--
	}
	return buffer
}

func TestReversed(t *testing.T) {
	RegisterTestingT(t)

	itReverses := func(xs []int) bool {
		result := Reversed(Reversed(xs))

		return slices.Equal(xs, result)
	}

	Expect(quick.Check(itReverses, nil)).To(Succeed())
}

func TestReversed2(t *testing.T) {
	RegisterTestingT(t)

	itReversesReally := func(xs []int) bool {
		reversed := Reversed(xs)

		if !SlicesAreReversed(xs, reversed) {
			t.Log("Intermediate was not reversed")
			return false
		}

		result := Reversed(reversed)

		return slices.Equal(xs, result)
	}

	Expect(quick.Check(itReversesReally, nil)).To(Succeed())
}

func SlicesAreReversed[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	i := 0
	j := len(a) - 1
	for i <= j {
		if a[i] != b[j] {
			return false
		}
		i++
		j--
	}

	return true
}
