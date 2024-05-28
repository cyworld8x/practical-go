package workerpool

//worker pool design pattern is a creational design pattern that uses a fixed-size pool of workers to perform concurrent tasks.
//The worker pool design pattern is used to create a pool of workers that can perform concurrent tasks.
//The worker pool design pattern is useful when you have a large number of tasks that need to be performed concurrently.

//Context sample: 3 workers are used to perform 50 concurrent tasks in a fixed-size pool of workers.

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func Run() {
	const numJobs = 50
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
