package main

import (
	"math/rand"
	"time"

	concurrency "github.com/stackpath/backend-developer-tests/concurrency/pkg"
)

func Waiting() {
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
}

func main() {
	sp := concurrency.NewSimplePool(4)

	// Create a bunch of tasks to do asynchronously
	for i := 0; i < 20; i++ {
		sp.Submit(Waiting)
	}

}
