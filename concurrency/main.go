package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	concurrency "github.com/stackpath/backend-developer-tests/concurrency/pkg"
)

func Waiting() {
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
}

func AdvWaiting(ctx context.Context) {
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
}

func main() {
	//fmt.Println("=============Testing Simple Pool===============")
	//sp := concurrency.NewSimplePool(4)

	// Create a bunch of tasks to do asynchronously
	//for i := 0; i < 10; i++ {
	//sp.Submit(Waiting)
	//}

	fmt.Println("=============Testing Advanced Pool=============")

	ctx := context.Background()
	ctx, cancleCtx := context.WithCancel(ctx)

	ap, err := concurrency.NewAdvancedPool(12, 4)
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i <= 40; i++ {
		err := ap.Submit(ctx, AdvWaiting)
		if err != nil {
			fmt.Println(err)
		}
		if i%30 == 0 {
			fmt.Println("CLOSE")
			ap.Close(cancleCtx)
		}
	}
}
