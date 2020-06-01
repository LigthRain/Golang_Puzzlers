package test1

import (
	"fmt"
	_ "unsafe"
)

//go:linkname inner src/test/ticker/linkName/outer.Outer
func inner() {
	fmt.Println("内部方法，你居然访问到了")
}
