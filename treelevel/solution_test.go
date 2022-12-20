package treelevel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFuncTwoLevels(t *testing.T) {
	a := NewTree("A")
	b := NewTree("B")
	c := NewTree("C")

	a.Add(b)
	a.Add(c)

	assert.Equal(t, "A", a.value)
	assert.Equal(t, "B", b.value)
	assert.Equal(t, "C", c.value)

	assert.Len(t, a.children, 2)
	assert.Len(t, b.children, 0)
	assert.Len(t, c.children, 0)

	out := a.GetLevelWidth()
	assert.Equal(t, out, []int{1, 2})
}

func TestMyFuncThreeLevels(t *testing.T) {
	a := NewTree("A")
	b := NewTree("B")
	c := NewTree("C")
	d := NewTree("D")
	e := NewTree("E")
	f := NewTree("F")

	d.Add(f)
	b.Add(e)
	a.Add(d)
	a.Add(c)
	a.Add(b)

	assert.Equal(t, "A", a.value)
	assert.Equal(t, "B", b.value)
	assert.Equal(t, "C", c.value)
	assert.Equal(t, "D", d.value)
	assert.Equal(t, "E", e.value)
	assert.Equal(t, "F", f.value)

	assert.Len(t, a.children, 3)
	assert.Len(t, b.children, 1)
	assert.Len(t, c.children, 0)
	assert.Len(t, d.children, 1)
	assert.Len(t, e.children, 0)
	assert.Len(t, f.children, 0)

	out := a.GetLevelWidth()
	assert.Equal(t, out, []int{1, 3, 2})
}

func TestMyFuncFourLevels(t *testing.T) {
	a := NewTree("A")
	b := NewTree("B")
	c := NewTree("C")
	d := NewTree("D")
	e := NewTree("E")

	c.Add(e)
	b.Add(c)
	b.Add(d)
	a.Add(b)

	assert.Equal(t, "A", a.value)
	assert.Equal(t, "B", b.value)
	assert.Equal(t, "C", c.value)
	assert.Equal(t, "D", d.value)
	assert.Equal(t, "E", e.value)

	assert.Len(t, a.children, 1)
	assert.Len(t, b.children, 2)
	assert.Len(t, c.children, 1)
	assert.Len(t, d.children, 0)
	assert.Len(t, e.children, 0)

	out := a.GetLevelWidth()
	assert.Equal(t, out, []int{1, 1, 2, 1})
}
