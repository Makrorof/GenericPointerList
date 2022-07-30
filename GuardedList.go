package PointerList

import "sync"

/////////////////////////////////
//          Guarded List       //
/////////////////////////////////

type GuardedPointerList[T any] interface {
	ToArray() []*T
	Count() int
	Get(index int) *T
	Add(item *T)
	AddRange(item []*T)
	Remove(targetItem *T)
	Clear()
	Contains(targetItem *T) bool
	Insert(targetItem *T, targetIndex int) error
	InsertRange(targetItems []*T, targetIndex int) error
	Reverse()
	Sort(f SortPointerFunc[T])
}

type guardedPointerList[T any] struct {
	locker sync.Mutex
	list   []*T
}

func NewGuardedPointerList[T any]() GuardedPointerList[T] {
	return &guardedPointerList[T]{
		list: make([]*T, 0),
	}
}

func (l *guardedPointerList[T]) Get(index int) *T {
	if index >= len(l.list) || index < 0 {
		return nil
	}

	l.locker.Lock()
	defer l.locker.Unlock()

	return l.list[index]
}

func (l *guardedPointerList[T]) Count() int {
	return len(l.list)
}

func (l *guardedPointerList[T]) Add(item *T) {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.list = append(l.list, item)
}

func (l *guardedPointerList[T]) AddRange(item []*T) {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.list = append(l.list, item...)
}

func (l *guardedPointerList[T]) Remove(targetItem *T) {
	l.locker.Lock()
	defer l.locker.Unlock()

	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return
		}
	}
}

func (l *guardedPointerList[T]) RemoveAt(index int) {
	if index >= len(l.list) || index < 0 {
		return
	}

	l.locker.Lock()
	defer l.locker.Unlock()

	l.list = append(l.list[:index], l.list[index+1:]...)
}

func (l *guardedPointerList[T]) ToArray() []*T {
	return l.list
}

func (l *guardedPointerList[T]) Clear() {
	l.locker.Lock()
	defer l.locker.Unlock()

	l.list = make([]*T, 0)
}

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
