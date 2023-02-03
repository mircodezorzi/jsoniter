package jsoniter

import "sync"

type interner struct {
	sync.Pool
}

func newInterner() *interner {
	return &interner{
		Pool: sync.Pool{
			New: func() interface{} {
				return make(map[string]string)
			},
		},
	}
}

// String returns s, interned.
func (pool *interner) String(s string) string {
	m := pool.Get().(map[string]string)
	c, ok := m[s]
	if ok {
		pool.Put(m)
		return c
	}
	m[s] = s
	pool.Put(m)
	return s
}

// Bytes returns b converted to a string, interned.
func (pool *interner) Bytes(b []byte) string {
	m := pool.Get().(map[string]string)
	c, ok := m[string(b)]
	if ok {
		pool.Put(m)
		return c
	}
	s := string(b)
	m[s] = s
	pool.Put(m)
	return s
}
