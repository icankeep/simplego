package setx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Int64(t *testing.T) {
	a := assert.New(t)

	set := NewSet[int64](1, 2, 3, 8888888888, 9999999999)
	set.Add(1, 4)

	a.True(set.Contains(4))
	a.False(set.Contains(5))
	a.True(set.Len() == 6)

	a.True(set.Len() == len(set.ToSlice()))
	a.True(set.Len() == set.Clone().Len())
	a.Equal(set, set.Clone())
	a.ElementsMatch(set.ToSlice(), set.Clone().ToSlice())

	a.True(set.ContainsAll(set.Clone()))

	newSet := NewSet[int64](999988, 5599, 33999, 1930, 485033, 1, 2, 3)
	a.False(newSet.Remove(111))
	a.True(newSet.Len() == 8)
	a.False(set.ContainsAll(newSet))

	a.True(newSet.Contains(999988))
	a.True(newSet.Remove(999988, 5599))
	a.False(newSet.Contains(999988, 5599))

	a.True(newSet.Remove(5599, 33999, 1930, 485033))
	a.False(newSet.Contains(5599, 33999, 1930, 485033))
	a.True(set.ContainsAll(newSet))

	a.True(set.Clear())
	a.True(newSet.Clear())
	a.True(newSet.IsEmpty())
	a.True(set.IsEmpty())

	set = NewSet[int64](999988, 5599, 33999, 1930, 485033, 1, 2, 3)
	a.True(set.Contains(5599))
	set.RemoveAll(NewSet[int64](5599, 33999, 1111))
	a.True(set.Len() == 6)
	a.False(set.Contains(5599))
	a.False(set.Contains(33999))
}

func Test_UnionSet(t *testing.T) {
	a := assert.New(t)

	s1 := NewSet[string]("hahah", "testing")
	s2 := NewSet[string]("haha", "testing2", "testing")
	set := UnionSet(s1, s2)
	a.False(set.IsEmpty())
	a.True(set.Len() == 4)
	a.True(set.Contains("haha"))
	a.True(set.Contains("hahah"))
	a.False(set.Contains("hahazz"))

	slice1 := []string{"hahazz", "testing2zz", "testingzz"}
	slice2 := []string{"haha", "testingtt", "testing2zz"}
	set = NewSetFromSlice[string](slice2, slice1)
	a.False(set.IsEmpty())
	a.True(set.Len() == 5)
	a.True(set.Contains("haha"))
	a.True(set.Contains("hahazz"))
	a.False(set.Contains("zzzz"))
}

func Test_IntersectionSet(t *testing.T) {
	a := assert.New(t)

	s1 := NewSet[string]("hahah", "testing", "zzzz")
	s2 := NewSet[string]("haha", "testing2", "testing")

	set := IntersectionSet(s1, s2)
	a.False(set.IsEmpty())
	a.True(set.Len() == 1)
	a.False(set.Contains("haha"))
	a.False(set.Contains("hahah"))
	a.True(set.Contains("testing"))

	s1 = NewSet[string]("hahah", "testing", "zzzz")
	s2 = NewSet[string]("haha", "testing2")
	set = IntersectionSet(s1, s2)
	a.True(set.IsEmpty())
}

func Test_DiffSet(t *testing.T) {
	a := assert.New(t)

	s1 := NewSet[string]("hahah", "testing", "zzzz", "hhh")
	s2 := NewSet[string]("haha", "testing2", "testing")

	set := DiffSet(s1, s2)
	a.False(set.IsEmpty())
	a.True(set.Len() == 3)
	a.True(set.Contains("hahah", "zzzz"))
	a.False(set.Contains("testing"))
	a.False(set.Contains("testing", "xxxx"))

	s1 = NewSet[string]()
	s2 = NewSet[string]("haha", "testing2", "testing")
	set = DiffSet(s1, s2)
	a.True(set.IsEmpty())

	s1 = NewSet[string]("hha")
	s2 = NewSet[string]("haha", "testing2", "testing")
	set = DiffSet(s1, s2)
	a.True(set.Len() == 1)
}
