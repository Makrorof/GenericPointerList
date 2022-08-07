package PointerList

import "sync"

type GuardedTagList[T any] interface {
	ToMap() map[string]GuardedPointerList[T]

	Get(key string) GuardedPointerList[T]

	//Adds a tag with the specified key and value to the list.
	Add(key string, value *T)

	//Removes all elements from the GuardedTagList.
	Clear()

	//Removes target list from the GuardedTagList.
	ClearList(key string)

	//Determines whether a tag is in the GuardedTagList.
	Contains(key string, value *T) bool

	//Inserts an element into the GuardedTagList at the specified index.
	Insert(index int, key string, value *T)

	//Removes the first occurrence of a specific object from the GuardedTagList.
	Remove(key string, value *T) bool

	//Removes the element at the specified index of the GuardedTagList.
	RemoveAt(key string, index int) bool

	//Returns the number of elements in a sequence.
	Count() int

	//Returns the number of elements in a sequence using the specified CountSelectTagListFunc[T]
	CountSelect(f CountSelectTagListFunc[T]) int
}

type guardedTagList[T any] struct {
	mapList map[string]GuardedPointerList[T]
	locker  sync.Mutex
}

func NewGuardedTagList[T any]() GuardedTagList[T] {
	return &guardedTagList[T]{
		mapList: make(map[string]GuardedPointerList[T]),
	}
}

func (l *guardedTagList[T]) Count() int {
	return len(l.mapList)
}

func (l *guardedTagList[T]) CountSelect(f CountSelectTagListFunc[T]) int {
	count := 0

	for key, list := range l.mapList {
		for index, current := range list.ToArray() {
			if f(key, index, current) {
				count++
			}
		}
	}

	return count
}

func (l *guardedTagList[T]) ToMap() map[string]GuardedPointerList[T] {
	l.locker.Lock()
	defer l.locker.Unlock()

	return l.mapList
}

func (l *guardedTagList[T]) Get(key string) GuardedPointerList[T] {
	l.locker.Lock()
	defer l.locker.Unlock()

	return l.mapList[key]
}

//Adds a tag with the specified key and value to the list.
func (l *guardedTagList[T]) Add(key string, value *T) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.mapList[key] == nil {
		l.mapList[key] = NewGuardedPointerList[T]()
	}

	l.mapList[key].Add(value)
}

//Removes all elements from the GuardedTagList.
func (l *guardedTagList[T]) Clear() {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.mapList = make(map[string]GuardedPointerList[T])
}

//Removes target list from the GuardedTagList.
func (l *guardedTagList[T]) ClearList(key string) {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.mapList[key] = nil
}

//Determines whether a tag is in the GuardedTagList.
func (l *guardedTagList[T]) Contains(key string, value *T) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.mapList[key] == nil {
		return false
	}

	return l.mapList[key].Contains(value)
}

//Inserts an element into the GuardedTagList at the specified index.
func (l *guardedTagList[T]) Insert(index int, key string, value *T) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.mapList[key] == nil {
		l.mapList[key] = NewGuardedPointerList[T]()
	}

	l.mapList[key].Insert(value, index)
}

//Removes the first occurrence of a specific object from the GuardedTagList.
func (l *guardedTagList[T]) Remove(key string, value *T) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.mapList[key] == nil {
		return false
	}

	return l.mapList[key].Remove(value)
}

//Removes the element at the specified index of the GuardedTagList.
func (l *guardedTagList[T]) RemoveAt(key string, index int) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.mapList[key] == nil {
		return false
	}

	return l.mapList[key].RemoveAt(index)
}
