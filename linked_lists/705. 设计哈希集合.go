package linked_lists

import "container/list"

const base = 769

func hash(key int) int {
	return key % base
}

type MyHashSet struct {
	data []list.List
}

func NewMyHashSet() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}
func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	h := hash(key)
	s.data[h].PushBack(key)
}
func (s *MyHashSet) Remove(key int) {
	h := hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value == key {
			s.data[h].Remove(e)
			return
		}
	}
}
func (s *MyHashSet) Contains(key int) bool {
	h := hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
