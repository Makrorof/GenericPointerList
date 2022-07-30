package PointerList

/////////////////////////////////
//        Pointer List         //
/////////////////////////////////

//List containing only pointer variables
type PointerList[T any] interface {
	ToArray() []*T
	//Returns the number of elements in a sequence.
	Count() int
	//Gets the element at the specified index.
	Get(index int) *T
	//Adds an object to the end of the PointerList[T].
	Add(item *T)
	//Adds the elements of the specified collection to the end of the PointerList[T].
	AddRange(item []*T)
	//Removes the first occurrence of a specific object from the PointerList[T].
	Remove(targetItem *T) bool
	//Removes the element at the specified index of the PointerList[T].
	RemoveAt(index int) bool
	//Removes all the elements that match the conditions defined by the specified predicate.
	RemoveAll(f RemovePointerFunc[T])
	//Removes all elements from the PointerList[T].
	Clear()
	//Determines whether an element is in the PointerList[T].
	Contains(targetItem *T) bool
	//Inserts an element into the PointerList[T] at the specified index.
	Insert(targetItem *T, targetIndex int) error
	//Inserts the elements of a collection into the PointerList[T] at the specified index.
	InsertRange(targetItems []*T, targetIndex int) error
	//Reverses the order of the elements in the entire PointerList[T].
	Reverse()
	//Sorts the elements in the entire PointerList[T] using the specified SortPointerFunc[T].
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

//Gets the element at the specified index.
func (l *pointerList[T]) Get(index int) *T {
	if index >= len(l.list) || index < 0 {
		return nil
	}

	return l.list[index]
}

//Returns the number of elements in a sequence.
func (l *pointerList[T]) Count() int {
	return len(l.list)
}

//Adds an object to the end of the PointerList[T].
func (l *pointerList[T]) Add(item *T) {
	l.list = append(l.list, item)
}

//Adds the elements of the specified collection to the end of the PointerList[T].
func (l *pointerList[T]) AddRange(item []*T) {
	l.list = append(l.list, item...)
}

//Removes the first occurrence of a specific object from the PointerList[T].
func (l *pointerList[T]) Remove(targetItem *T) bool {
	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return true
		}
	}

	return false
}

//Removes the element at the specified index of the PointerList[T].
func (l *pointerList[T]) RemoveAt(index int) bool {
	if index >= len(l.list) || index < 0 {
		return false
	}

	l.list = append(l.list[:index], l.list[index+1:]...)
	return true
}

//Removes all the elements that match the conditions defined by the specified predicate.
func (l *pointerList[T]) RemoveAll(f RemovePointerFunc[T]) {
	for i, i2 := 0, 0; i < len(l.list); i, i2 = i+1, i2+1 {
		if f(l.list[i], i2) {
			l.list = append(l.list[:i], l.list[i+1:]...)
			i--
		}
	}
}

func (l *pointerList[T]) ToArray() []*T {
	return l.list
}

//Removes all elements from the PointerList[T].
func (l *pointerList[T]) Clear() {
	l.list = make([]*T, 0)
}

//Determines whether an element is in the PointerList[T].
func (l *pointerList[T]) Contains(targetItem *T) bool {
	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			return true
		}
	}

	return false
}

//Inserts an element into the PointerList[T] at the specified index.
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

//Inserts the elements of a collection into the PointerList[T] at the specified index.
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

//Reverses the order of the elements in the entire PointerList[T].
func (l *pointerList[T]) Reverse() {
	if len(l.list) == 0 {
		return
	}

	for i, j := 0, len(l.list)-1; i < j; i, j = i+1, j-1 {
		l.list[i], l.list[j] = l.list[j], l.list[i]
	}
}

//Sorts the elements in the entire PointerList[T] using the specified SortPointerFunc[T].
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
