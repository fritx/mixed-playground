package linked_lists

import "container/list"

type entry struct {
	Key, Value int
}
type MyHashMap struct {
	data []list.List
}

func NewMyHashMap() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}
func (m *MyHashMap) Put(key int, value int) {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		// ** bug 1 - missing type assertion
		// if e.Value.Key == key {
		if et := e.Value.(entry); et.Key == key {
			// ** bug 2 - should operate e.Value directly
			// et.Value = value
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}
func (m *MyHashMap) Get(key int) int {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		// if e.Value.Key == key {
		if et := e.Value.(entry); et.Key == key {
			return et.Value
		}
	}
	return -1
}
func (m *MyHashMap) Remove(key int) {
	h := hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		// if e.Value.Key == key {
		if et := e.Value.(entry); et.Key == key {
			m.data[h].Remove(e)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
