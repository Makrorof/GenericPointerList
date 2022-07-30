package PointerList

//Example: left.id > right.id | [5,4,3] => 5>4(true), 5>3(true)...4>5(false) => [3,4,5]
type SortPointerFunc[T any] func(left *T, right *T) bool

//Example: current.IsNull()-> true => deleted, current.IsNull()-> false => skip
type RemovePointerFunc[T any] func(current *T, index int) bool

//Example: return true; => finded, return false; => skip
type FindPointerFunc[T any] func(current *T) bool

//Example: return true; => skip, return false; => direct return false;
type TrueForAllPointerFunc[T any] func(current *T) bool
