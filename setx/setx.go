package setx

type Set[T comparable] struct {
	m map[T]struct{}
}

type ISet interface {
	Add(item any)

	AddAll(set Set[any])

	Contains(item any) bool

	ContainsAll(set Set[any]) bool

	ToSlice() []any

	Remove(item any) bool

	RemoveAll(set Set[any]) bool

	Clear() bool

	Len() int

	IsEmpty() bool

	Clone()
}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{m: make(map[T]struct{})}
	s.Add(items...)
	return s
}

func NewSetFromSlice[T comparable](slices ...[]T) Set[T] {
	set := NewSet[T]()
	for i := range slices {
		set.Add(slices[i]...)
	}
	return set
}

func (s Set[T]) Add(items ...T) {
	for i := range items {
		s.m[items[i]] = struct{}{}
	}
}

func (s Set[T]) AddAll(set Set[T]) {
	for k := range set.m {
		s.Add(k)
	}
}

func (s Set[T]) Contains(items ...T) bool {
	for i := range items {
		_, ok := s.m[items[i]]
		if !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) ContainsAll(set Set[T]) bool {
	for k := range set.m {
		if !s.Contains(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) ToSlice() []T {
	items := make([]T, 0, s.Len())
	for k := range s.m {
		items = append(items, k)
	}
	return items
}

func (s Set[T]) Remove(items ...T) bool {
	modified := false
	for i := range items {
		if _, ok := s.m[items[i]]; ok {
			delete(s.m, items[i])
			modified = true
		}
	}
	return modified
}

func (s Set[T]) RemoveAll(set Set[T]) bool {
	modified := false
	for k := range set.m {
		delete(s.m, k)
		modified = true
	}
	return modified
}

func (s Set[T]) Clear() bool {
	modified := false
	for k := range s.m {
		delete(s.m, k)
		modified = true
	}
	return modified
}

func (s Set[T]) Len() int {
	return len(s.m)
}

func (s Set[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s Set[T]) Clone() Set[T] {
	newSet := NewSet[T]()
	newSet.Add(s.ToSlice()...)
	return newSet
}

// UnionSet return union of all sets
func UnionSet[T comparable](sets ...Set[T]) Set[T] {
	set := NewSet[T]()
	for i := range sets {
		set.AddAll(sets[i])
	}
	return set
}

// IntersectionSet return s1 & s2
func IntersectionSet[T comparable](s1, s2 Set[T]) Set[T] {
	set := NewSet[T]()
	for k := range s1.m {
		if s2.Contains(k) {
			set.Add(k)
		}
	}
	return set
}

// DiffSet return s1 - s2
func DiffSet[T comparable](s1, s2 Set[T]) Set[T] {
	set := NewSet[T]()
	for k := range s1.m {
		if s2.Contains(k) {
			continue
		}
		set.Add(k)
	}
	return set
}
