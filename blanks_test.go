package goproptest

import (
	"testing"
	"testing/quick"

	. "github.com/onsi/gomega"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sort[T constraints.Ordered](a []T) {
	// for i := 0; i < len(a)-1; i++ {
	// 	for j := len(a) - 1; j > i; j-- {
	// 		if a[j] < a[j-1] {
	// 			a[j], a[j-1] = a[j-1], a[j]
	// 		}
	// 	}
	// }
}

// Standard unit approach -----------------------------------------------------

func TestSortStandard(t *testing.T) {
	RegisterTestingT(t)

	t.Run("it does not panic if arg is empty", func(t *testing.T) {
	})
	t.Run("it sorts 3 reversed elements", func(t *testing.T) {
	})
}

// Invariant ------------------------------------------------------------------

func TestSortWithInvariant(t *testing.T) {
	RegisterTestingT(t)
}

// Model  ---------------------------------------------------------------------

func BiggestModel(xs []int) int {
	// Inefficient model
	return 0
}

func Biggest(xs []int) int {
	return 0
}

func TestBiggest(t *testing.T) {
	RegisterTestingT(t)
}

// Oracle Model ---------------------------------------------------------------

func SortedOracleModel(xs []int) []int {
	return nil
}

func Sorted(xs []int) []int {
	return nil
}

func TestSortedWithOracleModel(t *testing.T) {
	RegisterTestingT(t)
}

// Symmetric tests ------------------------------------------------------------

func TestHexCodec(t *testing.T) {
	RegisterTestingT(t)

	// itEncodesAndDecodes :=

	// Expect(quick.Check(itEncodesAndDecodes, nil)).To(Succeed())
}

// Symmetric tests ------------------------------------------------------------

func Reversed[T constraints.Ordered](xs []T) []T {
	buffer := slices.Clone(xs)
	// i := 0
	// j := len(buffer) - 1
	// for i < j {
	// 	buffer[i], buffer[j] = buffer[j], buffer[i]
	// 	i++
	// 	j--
	// }
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
