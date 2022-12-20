package combinequeue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueueString(t *testing.T) {
	q1 := NewQueue[int]()
	q1.Add(1)
	q1.Add(2)
	q1.Add(3)
	q1.Add(4)

	q2 := NewQueue[string]()
	q2.Add("one")
	q2.Add("two")
	q2.Add("three")

	q3 := combine(q1, q2)

	out := q3.Remove()
	require.Equal(t, 1, *out)

	out = q3.Remove()
	require.Equal(t, "one", *out)

	out = q3.Remove()
	require.Equal(t, 2, *out)

	out = q3.Remove()
	require.Equal(t, "two", *out)

	out = q3.Remove()
	require.Equal(t, 3, *out)

	out = q3.Remove()
	require.Equal(t, "three", *out)

	out = q3.Remove()
	require.Equal(t, 4, *out)

	out = q3.Remove()
	require.Nil(t, out)
}
