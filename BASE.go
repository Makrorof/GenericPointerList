package PointerList

import "sync"

type BASE interface {
	start() // if BaseList != nil => start()
	end()   // if BaseList != nil => defer end()
}

type lockerBase struct {
	BASE
	locker sync.Mutex
}

func (l *lockerBase) start() {
	l.locker.Lock()
}

func (l *lockerBase) end() {
	l.locker.Unlock()
}
