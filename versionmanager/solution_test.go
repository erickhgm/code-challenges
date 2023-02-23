package template

import (
	"fmt"
	"testing"
)

func TestMyFunc(t *testing.T) {
	vm := NewVersionManager()

	vm.addVersion(1, false)
	vm.addVersion(2, true)
	vm.addVersion(3, true)
	vm.addVersion(4, false)
	vm.addVersion(5, true)
	vm.addVersion(6, true)

	out := vm.isCompatible(5, 6)
	fmt.Println(out)

	out = vm.isCompatible(4, 6)
	fmt.Println(out)

	out = vm.isCompatible(3, 6)
	fmt.Println(out)
}
