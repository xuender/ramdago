package slice_test

import (
	"fmt"

	"github.com/xuender/helper/slice"
)

func ExampleNone() {
	equals3 := func(num int) bool { return num == 3 }

	fmt.Println(slice.None([]int{1, 2, 4, 5}, equals3))
	fmt.Println(slice.None([]int{3, 3, 1, 3}, equals3))

	// Output:
	// true
	// false
}
