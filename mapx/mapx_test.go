package mapx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Map(t *testing.T) {
	a := assert.New(t)

	m := NewMap[string, string]()
	v := m.Put("foo", "bar")
	a.Nil(v)
	a.Len(m, 1)

	v = m.Put("foo", "baz")
	a.NotNil(v)
	a.Len(m, 1)
	m.ContainsKey("foo")
	v = m.Get("foo")
	a.Equal(*v, "baz")

	v = m.Put("foo1", "baz")
	a.Nil(v)
	a.Len(m, 2)
	a.True(m.ContainsKey("foo1"))
	a.True(m.ContainsKey("foo"))
	a.Len(m.Keys(), 2)
	a.Len(m.Values(), 2)
	m.PutIfAbsent("foo", "baz")
	a.True(m.ContainsKey("foo"))
	a.Equal(*m.Get("foo"), "baz")
	a.Equal(m.GetOrDefault("foo", "zzz"), "baz")
	a.Equal(m.GetOrDefault("foo22", "zzz"), "zzz")

	m2 := FromMap[string, string](m)
	a.True(m2.Len() == m.Len())
	a.True(*m2.Get("foo") == "baz")

	a.False(m.IsEmpty())
	a.Nil(m.Remove("zzz"))
	a.Equal(*m.Remove("foo"), "baz")
	a.False(m.ContainsKey("foo"))
	m.PutIfAbsent("foo", "bazv2")
	a.True(*m.Get("foo") == "bazv2")
	a.Nil(v)
	m.Clear()
	a.Len(m, 0)
	a.True(m.IsEmpty())
}
