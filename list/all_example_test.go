package list_test

import (
	"fmt"

	"github.com/xuender/rg/list"
	"github.com/xuender/rg/relation"
)

func ExampleAll() {
	equals3 := relation.Equals(3)
	fmt.Println(list.All(equals3)([]int{3, 3, 3, 3}))
	fmt.Println(list.All(equals3)([]int{3, 3, 1, 3}))

	// Output:
	// true
	// false
}
