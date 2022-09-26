package PointerList

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	wg := &sync.WaitGroup{}
	forCount := 100

	log.Println("---Normal Pointer List---")
	log.Println("---Loading Pointer List---")
	normalList := NewPointerList[int]()

	for i := 0; i < forCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			log.Println("Thread ID:", index)
			normalList.Add(&index)
		}(i)
	}
	wg.Wait()
	log.Println("---Print Pointer List---")

	for i := 0; i < normalList.Count(); i++ {
		log.Println(i, ". => ", *normalList.Get(i))
	}
	/////////////////////////
	log.Println("---Guarded Pointer List---")
	log.Println("---Loading Guarded Pointer List---")

	guardedList := NewGuardedPointerList[int]()

	for i := 0; i < forCount; i++ {
		wg.Add(1)
		go func(index int, curWg *sync.WaitGroup) {
			defer curWg.Done()
			log.Println("Thread ID:", index)
			guardedList.Add(&index)
		}(i, wg)
	}
	wg.Wait()
	log.Println("---Print Pointer List---")

	for i := 0; i < guardedList.Count(); i++ {
		log.Println(i, ". => ", *guardedList.Get(i))
	}

	if forCount != guardedList.Count() {
		t.Errorf("forCount(%d) != guardedList(%d)", forCount, guardedList.Count())
	}

	if forCount != normalList.Count() {
		t.Logf("forCount(%d) != normalList(%d)", forCount, normalList.Count())
	}
}

func TestTagList(t *testing.T) {

	tagList := NewTagList[int]()

	var testVal int = 31

	tagList.Add("Test", &testVal)
}

func TestForeachTagList(t *testing.T) {
	list := NewGuardedTagList[int]()

	for i := 0; i < 25; i++ {
		newI := i
		list.Add(fmt.Sprint(i), &newI)
	}

	list.Foreach(func(key string, index int, current *int, removeCurItem func()) bool {
		log.Println("[", key, "] =>", index, ". =>", *current)

		if index%2 == 0 {
			removeCurItem()
		}

		return true
	})

	log.Println("Listele")

	list.Foreach(func(key string, index int, current *int, removeCurItem func()) bool {
		log.Println("[", key, "] =>", index, ". =>", *current)
		return true
	})
}
