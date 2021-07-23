package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/7/21 4:49 下午
 * @Desc: Single Producer and Single Consumer - scenario
 */

var messages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

const consumerCount = 3
const producerCount = 4

func producer(jobs chan<- string, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range messages[idx] {
		fmt.Printf("Producer %v sending message \"%v\"\n", idx, msg)
		jobs <- msg
	}
}

func consumer(jobs <-chan string, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range jobs {
		fmt.Printf("Consumer %v consumed message \"%v\"\n", idx, msg)
	}
}

func main() {
	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	wp.Add(producerCount)
	wc.Add(consumerCount)
	jobs := make(chan string)

	for i := 0; i < producerCount; i++ {

		go producer(jobs, i, wp)
	}

	for i := 0; i < consumerCount; i++ {
		go consumer(jobs, i, wc)
	}

	wp.Wait()
	close(jobs)
	wc.Wait()

}
