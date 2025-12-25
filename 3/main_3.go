package main

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() map[string]int {
	m := make(map[string]int)
	for k, v := range s.data {
		m[k] = v
	}
	return m
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.data[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.data[key]
	return v, ok
}
