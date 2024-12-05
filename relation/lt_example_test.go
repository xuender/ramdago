package relation_test

import (
	"fmt"

	"github.com/xuender/rg/relation"
)

func ExampleLt() {
	fmt.Println(relation.Lt(2)(1))
	fmt.Println(relation.Lt('a')('z'))

	// Output:
	// false
	// true
}
