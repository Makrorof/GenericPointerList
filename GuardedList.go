package PointerList

import "sync"

/////////////////////////////////
//          Guarded List       //
/////////////////////////////////

//List protected by mutex and containing only pointer variables
type GuardedPointerList[T any] interface {
	ToArray() []*T
	//Returns the number of elements in a sequence.
	Count() int
	//Gets the element at the specified index.
	Get(index int) *T
	//Adds an object to the end of the GuardedPointerList[T].
	Add(item *T)
	//Adds the elements of the specified collection to the end of the GuardedPointerList[T].
	AddRange(item []*T)
	//Removes the first occurrence of a specific object from the GuardedPointerList[T].
	Remove(targetItem *T) bool
	//Removes the element at the specified index of the GuardedPointerList[T].
	RemoveAt(index int) bool
	//Removes all the elements that match the conditions defined by the specified predicate.
	RemoveAll(f RemovePointerFunc[T])
	//Removes all elements from the GuardedPointerList[T].
	Clear()
	//Determines whether an element is in the GuardedPointerList[T].
	Contains(targetItem *T) bool
	//Inserts an element into the GuardedPointerList[T] at the specified index.
	Insert(targetItem *T, targetIndex int) error
	//Inserts the elements of a collection into the GuardedPointerList[T] at the specified index.
	InsertRange(targetItems []*T, targetIndex int) error
	//Reverses the order of the elements in the entire GuardedPointerList[T].
	Reverse()
	//Sorts the elements in the entire GuardedPointerList[T] using the specified SortPointerFunc[T].
	Sort(f SortPointerFunc[T])
	//Searches for an element that matches the conditions defined by the specified predicate, and returns the first occurrence within the entire GuardedPointerList[T].
	Find(f FindPointerFunc[T]) *T
	//Retrieves all the elements that match the conditions defined by the specified predicate.
	FindAll(f FindPointerFunc[T]) []*T
	//Determines whether every element in the GuardedPointerList[T] matches the conditions defined by the specified predicate.
	TrueForAll(f TrueForAllPointerFunc[T]) bool
	//Find and remove
	FindAndRemove(f FindPointerFunc[T]) *T

	//Capacity() //TODO: ...
}

type guardedPointerList[T any] struct {
	list   []*T
	locker sync.Mutex
}

func NewGuardedPointerList[T any]() GuardedPointerList[T] {
	return &guardedPointerList[T]{
		list: make([]*T, 0),
	}
}

//Gets the element at the specified index.
func (l *guardedPointerList[T]) Get(index int) *T {
	if index >= len(l.list) || index < 0 {
		return nil
	}

	l.locker.Lock()
	defer l.locker.Unlock()

	return l.list[index]
}

//Returns the number of elements in a sequence.
func (l *guardedPointerList[T]) Count() int {
	return len(l.list)
}

//Adds an object to the end of the GuardedPointerList[T].
func (l *guardedPointerList[T]) Add(item *T) {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.list = append(l.list, item)
}

//Adds the elements of the specified collection to the end of the GuardedPointerList[T].
func (l *guardedPointerList[T]) AddRange(item []*T) {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.list = append(l.list, item...)
}

//Removes the first occurrence of a specific object from the GuardedPointerList[T].
func (l *guardedPointerList[T]) Remove(targetItem *T) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	return l.remove(targetItem)
}

//Removes the element at the specified index of the GuardedPointerList[T].
func (l *guardedPointerList[T]) RemoveAt(index int) bool {
	if index >= len(l.list) || index < 0 {
		return false
	}

	l.locker.Lock()
	defer l.locker.Unlock()

	return l.removeAt(index)
}

//Removes all the elements that match the conditions defined by the specified predicate.
func (l *guardedPointerList[T]) RemoveAll(f RemovePointerFunc[T]) {
	l.locker.Lock()
	defer l.locker.Unlock()

	for i, i2 := 0, 0; i < len(l.list); i, i2 = i+1, i2+1 {
		if f(l.list[i], i2) {
			l.list = append(l.list[:i], l.list[i+1:]...)
			i--
		}
	}
}

func (l *guardedPointerList[T]) ToArray() []*T {
	return l.list
}

//Removes all elements from the GuardedPointerList[T].
func (l *guardedPointerList[T]) Clear() {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.list = make([]*T, 0)
}

//Determines whether an element is in the GuardedPointerList[T].
func (l *guardedPointerList[T]) Contains(targetItem *T) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			return true
		}
	}

	return false
}

//Inserts an element into the GuardedPointerList[T] at the specified index.
func (l *guardedPointerList[T]) Insert(targetItem *T, targetIndex int) error {
	if targetIndex > len(l.list) || targetIndex < 0 {
		return GetErrorf(IndexOutOfRange, len(l.list), len(l.list))
	}

	l.locker.Lock()
	defer l.locker.Unlock()

	newArray := make([]*T, len(l.list)+1)

	for i, index := 0, 0; i < len(newArray); i, index = i+1, index+1 {
		if i == targetIndex {
			newArray[i] = targetItem
			index--
		} else {
			newArray[i] = l.list[index]
		}
	}

	l.list = newArray

	return nil
}

//Inserts the elements of a collection into the GuardedPointerList[T] at the specified index.
func (l *guardedPointerList[T]) InsertRange(targetItems []*T, targetIndex int) error {
	if targetIndex > len(l.list) || targetIndex < 0 {
		return GetErrorf(IndexOutOfRange, len(l.list), len(l.list))
	} else if len(targetItems) == 0 {
		return nil
	}

	l.locker.Lock()
	defer l.locker.Unlock()

	newArray := make([]*T, len(l.list)+len(targetItems))

	for i, index := 0, 0; i < len(newArray); i, index = i+1, index+1 {
		if i == targetIndex {
			for i2 := 0; i2 < len(targetItems); i2++ {
				newArray[i+i2] = targetItems[i2]
			}
			i += len(targetItems) - 1

			index--
		} else {
			newArray[i] = l.list[index]
		}
	}

	l.list = newArray

	return nil
}

//Reverses the order of the elements in the entire GuardedPointerList[T].
func (l *guardedPointerList[T]) Reverse() {
	if len(l.list) == 0 {
		return
	}

	l.locker.Lock()
	defer l.locker.Unlock()

	for i, j := 0, len(l.list)-1; i < j; i, j = i+1, j-1 {
		l.list[i], l.list[j] = l.list[j], l.list[i]
	}
}

//Sorts the elements in the entire GuardedPointerList[T] using the specified SortPointerFunc[T].
func (l *guardedPointerList[T]) Sort(f SortPointerFunc[T]) {
	l.locker.Lock()
	defer l.locker.Unlock()

	for i2 := 0; i2 < len(l.list); i2++ {
		for i := 0; i < len(l.list); i++ {
			if f(l.list[i], l.list[i2]) {
				l.list[i], l.list[i2] = l.list[i2], l.list[i]
			}
			/*if l.list[i] > l.list[i2] {
				l.list[i], l.list[i2] = l.list[i2], l.list[i]
			}*/
		}
	}
}

//Searches for an element that matches the conditions defined by the specified predicate, and returns the first occurrence within the entire GuardedPointerList[T].
func (l *guardedPointerList[T]) Find(f FindPointerFunc[T]) *T {
	l.locker.Lock()
	defer l.locker.Unlock()

	for i := 0; i < len(l.list); i++ {
		if f(l.list[i]) {
			return l.list[i]
		}
	}

	return nil
}

//Retrieves all the elements that match the conditions defined by the specified predicate.
func (l *guardedPointerList[T]) FindAll(f FindPointerFunc[T]) []*T {
	l.locker.Lock()
	defer l.locker.Unlock()

	retList := make([]*T, 0)

	for i := 0; i < len(l.list); i++ {
		if f(l.list[i]) {
			retList = append(retList, l.list[i])
		}
	}

	return retList
}

//Determines whether every element in the GuardedPointerList[T] matches the conditions defined by the specified predicate.
func (l *guardedPointerList[T]) TrueForAll(f TrueForAllPointerFunc[T]) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	for i := 0; i < len(l.list); i++ {
		if !f(l.list[i]) {
			return false
		}
	}

	return true
}

//Find and remove
func (l *guardedPointerList[T]) FindAndRemove(f FindPointerFunc[T]) *T {
	l.locker.Lock()
	defer l.locker.Unlock()

	for i := 0; i < len(l.list); i++ {
		if f(l.list[i]) {
			target := l.list[i]
			l.removeAt(i)
			return target
		}
	}

	return nil
}

/////////////////////////////////
//            PRIVATE          //
/////////////////////////////////

//Removes the element at the specified index of the GuardedPointerList[T].
func (l *guardedPointerList[T]) removeAt(index int) bool {
	if index >= len(l.list) || index < 0 {
		return false
	}

	l.list = append(l.list[:index], l.list[index+1:]...)
	return true
}

//Removes the first occurrence of a specific object from the GuardedPointerList[T].
func (l *guardedPointerList[T]) remove(targetItem *T) bool {
	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return true
		}
	}

	return false
}
