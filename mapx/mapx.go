package mapx

type Map[T comparable, V any] map[T]V

type IMap interface {
	Put(key, value any) any

	PutAll(ms ...Map[any, any])

	PutIfAbsent(key, value any)

	Get(key any) any

	GetOrDefault(key any, defaultValue any) any

	ContainsKey(key any) bool

	ContainsValue(value any) bool

	Remove(key any) any

	Clear()

	Len() int

	IsEmpty() bool

	Keys() []any

	Values() []any
}

func NewMap[T comparable, V any]() Map[T, V] {
	return make(Map[T, V])
}

func FromMap[T comparable, V any](ms ...Map[T, V]) Map[T, V] {
	newMap := NewMap[T, V]()
	newMap.PutAll(ms...)
	return newMap
}

// Put Associates the specified value with the specified key in this map (optional operation).
// If the map previously contained a mapping for the key, the old value is replaced by the specified value.
func (m Map[T, V]) Put(key T, value V) *V {
	oldValue, ok := m[key]
	m[key] = value
	if ok {
		return &oldValue
	}
	return nil
}

func (m Map[T, V]) PutAll(ms ...Map[T, V]) {
	for i := range ms {
		for k, v := range ms[i] {
			m.Put(k, v)
		}
	}
}

// PutIfAbsent If the specified key is not already associated with a value (or is mapped to null)
// associates it with the given value and returns null, else returns the current value.
func (m Map[T, V]) PutIfAbsent(key T, value V) {
	_, ok := m[key]
	if !ok {
		m[key] = value
	}
}

func (m Map[T, V]) Get(key T) *V {
	oldValue, ok := m[key]
	if ok {
		return &oldValue
	}
	return nil
}

func (m Map[T, V]) GetOrDefault(key T, defaultValue V) V {
	if value := m.Get(key); value != nil {
		return *value
	}
	return defaultValue
}

func (m Map[T, V]) ContainsKey(key T) bool {
	v := m.Get(key)
	if v == nil {
		return false
	}
	return true
}

//func (m Map[T, V]) ContainsValue(value V) bool {
//	for k := range m {
//		if m[k] == value {
//			return true
//		}
//	}
//	return false
//}

func (m Map[T, V]) Remove(key T) *V {
	oldValue := m.Get(key)
	delete(m, key)
	return oldValue
}

func (m Map[T, V]) Clear() {
	for k := range m {
		delete(m, k)
	}
}

func (m Map[T, V]) Len() int {
	return len(m)
}

func (m Map[T, V]) IsEmpty() bool {
	return m.Len() == 0
}

func (m Map[T, V]) Keys() []T {
	s := make([]T, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func (m Map[T, V]) Values() []V {
	s := make([]V, 0, len(m))
	for k := range m {
		s = append(s, m[k])
	}
	return s
}
