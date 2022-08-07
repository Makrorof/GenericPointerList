package PointerList

type TagList[T any] interface {
	ToMap() map[string]PointerList[T]

	Get(key string) PointerList[T]

	//Adds a tag with the specified key and value to the list.
	Add(key string, value *T)

	//Removes all elements from the TagList.
	Clear()

	//Removes target list from the TagList.
	ClearList(key string)

	//Determines whether a tag is in the TagList.
	Contains(key string, value *T) bool

	//Inserts an element into the TagList at the specified index.
	Insert(index int, key string, value *T)

	//Removes the first occurrence of a specific object from the TagList.
	Remove(key string, value *T) bool

	//Removes the element at the specified index of the TagList.
	RemoveAt(key string, index int) bool

	//Returns the number of elements in a sequence.
	Count() int

	//Total count
	TotalCount() int

	//Returns the number of elements in a sequence.
	MapCount() map[string]int

	//Returns the number of elements in a sequence using the specified CountSelectTagListFunc[T]
	MapCountSelect(f CountSelectTagListFunc[T]) map[string]int

	//Returns the number of elements in a sequence using the specified CountSelectTagListFunc[T]
	CountSelect(f CountSelectTagListFunc[T]) int
}

type tagList[T any] struct {
	mapList map[string]PointerList[T]
}

func NewTagList[T any]() TagList[T] {
	return &tagList[T]{
		mapList: make(map[string]PointerList[T]),
	}
}

func (l *tagList[T]) ToMap() map[string]PointerList[T] {
	return l.mapList
}

func (l *tagList[T]) MapCount() map[string]int {
	count := make(map[string]int)

	for key, value := range l.mapList {
		count[key] = value.Count()
	}

	return count
}

func (l *tagList[T]) TotalCount() int {
	count := 0

	for _, value := range l.mapList {
		count += value.Count()
	}

	return count
}

func (l *tagList[T]) Count() int {
	return len(l.mapList)
}

func (l *tagList[T]) CountSelect(f CountSelectTagListFunc[T]) int {
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

func (l *tagList[T]) MapCountSelect(f CountSelectTagListFunc[T]) map[string]int {
	count := make(map[string]int)

	for key, list := range l.mapList {
		curCount := 0
		for index, current := range list.ToArray() {
			if f(key, index, current) {
				curCount++
			}
		}

		count[key] = curCount
	}

	return count
}

func (l *tagList[T]) Get(key string) PointerList[T] {
	return l.mapList[key]
}

//Adds a tag with the specified key and value to the list.
func (l *tagList[T]) Add(key string, value *T) {
	if l.mapList[key] == nil {
		l.mapList[key] = NewPointerList[T]()
	}

	l.mapList[key].Add(value)
}

//Removes all elements from the TagList.
func (l *tagList[T]) Clear() {
	l.mapList = make(map[string]PointerList[T])
}

//Removes target list from the TagList.
func (l *tagList[T]) ClearList(key string) {
	l.mapList[key] = nil
}

//Determines whether a tag is in the TagList.
func (l *tagList[T]) Contains(key string, value *T) bool {
	if l.mapList[key] == nil {
		return false
	}

	return l.mapList[key].Contains(value)
}

//Inserts an element into the TagList at the specified index.
func (l *tagList[T]) Insert(index int, key string, value *T) {
	if l.mapList[key] == nil {
		l.mapList[key] = NewPointerList[T]()
	}

	l.mapList[key].Insert(value, index)
}

//Removes the first occurrence of a specific object from the TagList.
func (l *tagList[T]) Remove(key string, value *T) bool {
	if l.mapList[key] == nil {
		return false
	}

	return l.mapList[key].Remove(value)
}

//Removes the element at the specified index of the TagList.
func (l *tagList[T]) RemoveAt(key string, index int) bool {
	if l.mapList[key] == nil {
		return false
	}

	return l.mapList[key].RemoveAt(index)
}
