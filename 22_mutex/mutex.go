package main

import (
	"fmt"
	"sync"
)

type Post struct{
	views int
	mu sync.Mutex
}

func(P *Post) increment(wg *sync.WaitGroup){
	defer func(){
		P.mu.Unlock()
		wg.Done()
	}()

	P.mu.Lock()
	P.views +=1
}


func main() {

	var wg sync.WaitGroup

	post := Post{
		views: 0,
	}

	for i:=0; i<100;i++{
		wg.Add(1)
		go post.increment(&wg)
	}

	wg.Wait()

	// post.increment()
	// post.increment()

	fmt.Println(post.views)
}
