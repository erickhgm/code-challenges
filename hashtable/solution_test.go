package template

import (
	"fmt"
	"testing"
)

func TestMyFunc(t *testing.T) {
	hash := NewHashTable(10)

	hash.Put(1, "one")
	hash.Put(2, "two")
	hash.Put(3, "three")
	hash.Put(4, "four")
	hash.Put(5, "five")
	hash.Put(6, "six")
	hash.Put(7, "seven")
	hash.Put(8, "eight")
	hash.Put(9, "nine")

	one := hash.Get(1)
	fmt.Println(one)

	ten := hash.Get(10)
	fmt.Println(ten)

	hash.Put(1, "one-one")
	one = hash.Get(1)
	fmt.Println(one)

	hash.Put(11, "eleven")
	eleven := hash.Get(11)
	fmt.Println(eleven)

	k111 := hash.Get(111)
	fmt.Println(k111)

	hash.Put(111, "111")
	k111 = hash.Get(111)
	fmt.Println(k111)

	hash.Remove(111)

	one = hash.Get(1)
	fmt.Println(one)

	eleven = hash.Get(11)
	fmt.Println(eleven)

	k111 = hash.Get(111)
	fmt.Println(k111)

	size := hash.Size()
	fmt.Println(size)
}
