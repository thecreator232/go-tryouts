package main

import (
	"context"
	"fmt"
	"time"

	"github.com/thecreator232/go-tryouts/source/goroutines"
)

func main() {

	f, err := goroutines.NewWorkerPool(context.Background(), 5, 10*time.Second)

	if err != nil {
		fmt.Println("error in creating pool")
		return
	}

	for i := 1; i <= 100; i++ {
		f.ConsumeEvent(context.Background(), &goroutines.Event{
			msg: fmt.Sprintf("msg {%s}", i),
		})
	}
}
