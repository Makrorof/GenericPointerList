package PointerList

/////////////////////////////////
//        Pointer List         //
/////////////////////////////////

type PointerList[T any] interface {
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

	//Capacity() //TODO: ...
}

type pointerList[T any] struct {
	list []*T
}

func NewPointerList[T any]() PointerList[T] {
	return &pointerList[T]{
		list: make([]*T, 0),
	}
}

func (l *pointerList[T]) Get(index int) *T {
	if index >= len(l.list) || index < 0 {
		return nil
	}

	return l.list[index]
}

func (l *pointerList[T]) Count() int {
	return len(l.list)
}

func (l *pointerList[T]) Add(item *T) {
	l.list = append(l.list, item)
}

func (l *pointerList[T]) AddRange(item []*T) {
	l.list = append(l.list, item...)
}

func (l *pointerList[T]) Remove(targetItem *T) {
	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return
		}
	}
}

func (l *pointerList[T]) RemoveAt(index int) {
	if index >= len(l.list) || index < 0 {
		return
	}

	l.list = append(l.list[:index], l.list[index+1:]...)
}

func (l *pointerList[T]) ToArray() []*T {
	return l.list
}

func (l *pointerList[T]) Clear() {
	l.list = make([]*T, 0)
}

func (l *pointerList[T]) Contains(targetItem *T) bool {
	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			return true
		}
	}

	return false
}

func (l *pointerList[T]) Insert(targetItem *T, targetIndex int) error {
	if targetIndex > len(l.list) || targetIndex < 0 {
		return GetErrorf(IndexOutOfRange, len(l.list), len(l.list))
	}

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

func (l *pointerList[T]) InsertRange(targetItems []*T, targetIndex int) error {
	if targetIndex > len(l.list) || targetIndex < 0 {
		return GetErrorf(IndexOutOfRange, len(l.list), len(l.list))
	} else if len(targetItems) == 0 {
		return nil
	}

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

func (l *pointerList[T]) Reverse() {
	if len(l.list) == 0 {
		return
	}

	for i, j := 0, len(l.list)-1; i < j; i, j = i+1, j-1 {
		l.list[i], l.list[j] = l.list[j], l.list[i]
	}
}

func (l *pointerList[T]) Sort(f SortPointerFunc[T]) {
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
