package pakang

import (
	"testing"
	"github.com/taikedz/gocheck"
)

func TestSplitStringMultichar(t *testing.T) {
	res := SplitStringMultichar("12.04,1:5,1", ".,:")
	exp := []string{"12", "04", "1", "5", "1"}
	gocheck.EqualArr(t, exp, res)
}

func TestExtractInts(t *testing.T) {
	res, err := ExtractInts("12.04,1:5,1")
	if err != nil {
		t.Errorf("Error extracting ints: %s\n", err)
		return
	}
	exp := []int{12, 4, 1, 5, 1}
	gocheck.EqualArr(t, exp, res)
}

func TestArrIntsGte(t *testing.T) {
	gocheck.Equal(t, true,  ArrIntsGte([]int{1}, []int{1,0,0}))
	gocheck.Equal(t, true,  ArrIntsGte([]int{1,0}, []int{1,1}))
	gocheck.Equal(t, true,  ArrIntsGte([]int{1,0}, []int{1,0,1}))
	gocheck.Equal(t, true,  ArrIntsGte([]int{1,0,1}, []int{1,2}))
	gocheck.Equal(t, false, ArrIntsGte([]int{2,5}, []int{1,8}))
}

func TestArrIntsLte(t *testing.T) {
	gocheck.Equal(t, true,  ArrIntsLte([]int{1}, []int{1,0,0}))
	gocheck.Equal(t, true,  ArrIntsLte([]int{1,1}, []int{1,0}))
	gocheck.Equal(t, true,  ArrIntsLte([]int{1,0,1}, []int{1,0}))
	gocheck.Equal(t, true,  ArrIntsLte([]int{1,2}, []int{1,0,1}))
	gocheck.Equal(t, false, ArrIntsLte([]int{1,8}, []int{2,5}))
}