package PointerList

/////////////////////////////////
//        Pointer List         //
/////////////////////////////////

//List containing only pointer variables
type PointerList[T any] interface {
	BASE

	ToArray() []*T
	//Returns the number of elements in a sequence.
	Count() int
	//Gets the element at the specified index.
	Get(index int) *T
	//Gets the element at next
	GetNext() *T
	//Gets the element at next
	GetNextBefore(f BeforeListFunc[T]) *T

	//Adds an object to the end of the PointerList[T].
	Add(item *T)
	//Adds the elements of the specified collection to the end of the PointerList[T].
	AddRange(item []*T)
	//Removes the first occurrence of a specific object from the PointerList[T].
	Remove(targetItem *T) bool
	//Removes the element at the specified index of the PointerList[T].
	RemoveAt(index int) bool
	//Removes the first occurrence of a specific object from the PointerList[T]. NoSafe
	RemoveNoSafe(targetItem *T) bool
	//Removes the element at the specified index of the PointerList[T]. NoSafe
	RemoveAtNoSafe(index int) bool
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
	//Searches for an element that matches the conditions defined by the specified predicate, and returns the first occurrence within the entire PointerList[T].
	Find(f FindPointerFunc[T]) *T
	//Retrieves all the elements that match the conditions defined by the specified predicate.
	FindAll(f FindPointerFunc[T]) []*T
	//Determines whether every element in the PointerList[T] matches the conditions defined by the specified predicate.
	TrueForAll(f TrueForAllPointerFunc[T]) bool
	//Find and remove
	FindAndRemove(f FindPointerFunc[T]) *T
	//Loop
	Foreach(f ForeachListFunc[T])
	//Capacity() //TODO: ...
}

type pointerList[T any] struct {
	BASE
	list      []*T
	lastIndex int
}

func NewPointerList[T any]() PointerList[T] {
	return &pointerList[T]{
		list: make([]*T, 0),
	}
}

//Gets the element at the specified index.
func (l *pointerList[T]) Get(index int) *T {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	if index >= len(l.list) || index < 0 {
		return nil
	}

	return l.list[index]
}

func (l *pointerList[T]) GetNext() *T {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	return l.getNext()
}

//Returns the number of elements in a sequence.
func (l *pointerList[T]) Count() int {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	return len(l.list)
}

//Adds an object to the end of the PointerList[T].
func (l *pointerList[T]) Add(item *T) {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	l.list = append(l.list, item)
}

//Adds the elements of the specified collection to the end of the PointerList[T].
func (l *pointerList[T]) AddRange(item []*T) {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	l.list = append(l.list, item...)
}

//Removes the first occurrence of a specific object from the PointerList[T].
func (l *pointerList[T]) Remove(targetItem *T) bool {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	return l.remove(targetItem)
}

//Removes the element at the specified index of the PointerList[T].
func (l *pointerList[T]) RemoveAt(index int) bool {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	return l.removeAt(index)
}

//Removes the first occurrence of a specific object from the PointerList[T]. NoSafe
func (l *pointerList[T]) RemoveNoSafe(targetItem *T) bool {
	return l.remove(targetItem)
}

//Removes the element at the specified index of the PointerList[T]. NoSafe
func (l *pointerList[T]) RemoveAtNoSafe(index int) bool {
	return l.removeAt(index)
}

//Removes all the elements that match the conditions defined by the specified predicate.
func (l *pointerList[T]) RemoveAll(f RemovePointerFunc[T]) {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	for i, i2 := 0, 0; i < len(l.list); i, i2 = i+1, i2+1 {
		if f(l.list[i], i2) {
			l.list = append(l.list[:i], l.list[i+1:]...)
			i--
		}
	}
}

func (l *pointerList[T]) ToArray() []*T {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	return l.list
}

//Removes all elements from the PointerList[T].
func (l *pointerList[T]) Clear() {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	l.list = make([]*T, 0)
}

//Determines whether an element is in the PointerList[T].
func (l *pointerList[T]) Contains(targetItem *T) bool {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			return true
		}
	}

	return false
}

//Inserts an element into the PointerList[T] at the specified index.
func (l *pointerList[T]) Insert(targetItem *T, targetIndex int) error {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

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
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

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
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	if len(l.list) == 0 {
		return
	}

	for i, j := 0, len(l.list)-1; i < j; i, j = i+1, j-1 {
		l.list[i], l.list[j] = l.list[j], l.list[i]
	}
}

//Sorts the elements in the entire PointerList[T] using the specified SortPointerFunc[T].
func (l *pointerList[T]) Sort(f SortPointerFunc[T]) {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

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

//Searches for an element that matches the conditions defined by the specified predicate, and returns the first occurrence within the entire PointerList[T].
func (l *pointerList[T]) Find(f FindPointerFunc[T]) *T {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	for i := 0; i < len(l.list); i++ {
		if f(i, l.list[i]) {
			return l.list[i]
		}
	}

	return nil
}

//Retrieves all the elements that match the conditions defined by the specified predicate.
func (l *pointerList[T]) FindAll(f FindPointerFunc[T]) []*T {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	retList := make([]*T, 0)

	for i := 0; i < len(l.list); i++ {
		if f(i, l.list[i]) {
			retList = append(retList, l.list[i])
		}
	}

	return retList
}

//Determines whether every element in the PointerList[T] matches the conditions defined by the specified predicate.
func (l *pointerList[T]) TrueForAll(f TrueForAllPointerFunc[T]) bool {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	for i := 0; i < len(l.list); i++ {
		if !f(l.list[i]) {
			return false
		}
	}

	return true
}

//Find and remove
func (l *pointerList[T]) FindAndRemove(f FindPointerFunc[T]) *T {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	for i := 0; i < len(l.list); i++ {
		if f(i, l.list[i]) {
			target := l.list[i]
			l.removeAt(i)
			return target
		}
	}

	return nil
}

func (l *pointerList[T]) GetNextBefore(f BeforeListFunc[T]) *T {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	for i := 0; i < len(l.list); i++ { //l.mapList[key].Count() => MaxCount
		currentItem := l.getNext()

		if currentItem == nil {
			break
		}

		if f(currentItem) {
			return currentItem
		}
	}

	return nil
}

func (l *pointerList[T]) Foreach(f ForeachListFunc[T]) {
	if l.BASE != nil {
		l.start()
		defer l.end()
	}

	for i := 0; i < len(l.list); i++ {
		if !f(i, l.list[i]) {
			break
		}
	}
}

/////////////////////////////////
//            PRIVATE          //
/////////////////////////////////

//Removes the element at the specified index of the PointerList[T].
func (l *pointerList[T]) removeAt(index int) bool {
	if index >= len(l.list) || index < 0 {
		return false
	}

	l.list = append(l.list[:index], l.list[index+1:]...)
	return true
}

//Removes the first occurrence of a specific object from the PointerList[T].
func (l *pointerList[T]) remove(targetItem *T) bool {
	for i := 0; i < len(l.list); i++ {
		if l.list[i] == targetItem {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return true
		}
	}

	return false
}

func (l *pointerList[T]) getNext() *T {
	if len(l.list) == 0 {
		return nil
	}

	if l.lastIndex >= len(l.list) {
		l.lastIndex = 0
	}

	selectedItem := l.list[l.lastIndex]
	l.lastIndex++

	return selectedItem
}
