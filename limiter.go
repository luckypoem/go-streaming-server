package main

import "log"

type ConnectionLimiter struct {
	concurrentConnection int
	bucket               chan int
}

func NewConnectionLimiter(cc int) *ConnectionLimiter {
	return &ConnectionLimiter{
		concurrentConnection: cc,
		bucket:               make(chan int, cc),
	}
}

func (limiter *ConnectionLimiter) GetConnection() bool {
	if len(limiter.bucket) >= limiter.concurrentConnection {
		log.Println("Reached the rate limitation.")
		return false
	}

	limiter.bucket <- 1
	return true
}

func (limiter *ConnectionLimiter) ReleaseConnection() {
	c := <-limiter.bucket
	log.Printf("New connection coming %v", c)
}
