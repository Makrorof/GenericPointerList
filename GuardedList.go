package PointerList

/////////////////////////////////
//          Guarded List       //
/////////////////////////////////

//List protected by mutex and containing only pointer variables
type GuardedPointerList[T any] interface {
	PointerList[T]
}

func NewGuardedPointerList[T any]() GuardedPointerList[T] {
	baseList := &lockerBase{}
	return &pointerList[T]{
		list: make([]*T, 0),
		BASE: baseList,
	}
}
