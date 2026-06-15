package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dbData = [5]string{"data1", "data2", "data3", "data4", "data5"}

// WaitGroup is used to wait until all goroutines are done before moving on
var wg sync.WaitGroup = sync.WaitGroup{}

// mutex is used to lock shared resources so goroutines have to wait
var m sync.Mutex

func main() {
	// GOROUTINE TUTORIAL

	// 1. calling without goroutine
	fmt.Println("SCENE 1: Querying DB without goroutine")
	var t1 time.Time = time.Now()
	var resultList []string

	for i := range dbData {
		dbCall(i, &resultList, true, false, false)
	}

	fmt.Println("resultList:", resultList)
	fmt.Println("Total execution time:", time.Since(t1))

	// 2. calling with goroutine
	fmt.Println("\nSCENE 2: Querying DB with goroutine")
	var t2 time.Time = time.Now()
	var resultList2 []string

	for i := range dbData {
		wg.Add(1) // add 1 task to wait for
		go dbCall(i, &resultList2, true, true, false)
	}
	wg.Wait() // wait until all tasks in the WaitGroup is done

	fmt.Println("resultList2:", resultList2)
	fmt.Println("Total execution time:", time.Since(t2))

	// 3. calling with goroutine, without concurrency handling
	fmt.Println("\nSCENE 3: Querying DB with goroutine without concurrency handling")
	var t3 time.Time = time.Now()
	const count int = 200_000
	var resultList3 []string

	for i := 0; i < count; i++ {
		wg.Add(1) // add 1 task to wait for
		go dbCall(i%5, &resultList3, false, true, false)
	}
	wg.Wait() // wait until all tasks in the WaitGroup is done

	fmt.Println("expected count:", count)
	fmt.Println("length of resultList3:", len(resultList3))
	fmt.Println(count-len(resultList3), "MISSING DATA DUE TO CONCURRENCY")
	fmt.Println("Total execution time:", time.Since(t3))

	// 4. calling with goroutine, with concurrency handling
	fmt.Println("\nSCENE 4: Querying DB with goroutine with concurrency handling")
	var t4 time.Time = time.Now()
	const count2 int = 200_000
	var resultList4 []string

	for i := 0; i < count2; i++ {
		wg.Add(1) // add 1 task to wait for
		go dbCall(i%5, &resultList4, false, true, true)
	}
	wg.Wait() // wait until all tasks in the WaitGroup is done

	fmt.Println("expected count:", count2)
	fmt.Println("length of resultList4:", len(resultList4))
	fmt.Println(count2-len(resultList4), "MISSING DATA DUE TO CONCURRENCY")
	fmt.Println("Total execution time:", time.Since(t4))
}

// fake DB call to showcase delay
func dbCall(i int, resultList *[]string, isRandomDelay bool, isAsync bool, useLock bool) {
	// var t0 time.Time = time.Now()

	if isRandomDelay {
		var delay float32 = rand.Float32() * 1000
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	currData := dbData[i]
	// fmt.Printf("fetched %v after %v\n", currData, time.Since(t0))

	if useLock {
		// use mutex lock to avoid concurrency conflict
		m.Lock()
		*resultList = append(*resultList, currData)
		m.Unlock()
	} else {
		*resultList = append(*resultList, currData)
	}

	if isAsync {
		wg.Done() // remove 1 task to wait for
	}
}
