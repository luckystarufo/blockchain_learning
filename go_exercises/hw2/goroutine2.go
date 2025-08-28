package main

import (
	"fmt"
	"time"
)

type Task struct {
	Name string
	Fn   func()
}

func RunTasks(tasks []Task) {
	resultCh := make(chan string, len(tasks))

	for _, task := range tasks {
		go func(t Task) {
			start := time.Now()
			t.Fn()
			duration := time.Since(start)
			resultCh <- fmt.Sprintf("%s executed in %v", t.Name, duration)
		}(task)
	}

	// 收集所有结果
	for i := 0; i < len(tasks); i++ {
		fmt.Println(<-resultCh)
	}
}

func main() {
	// test
	tasks := []Task{
		{"Task1", func() { time.Sleep(500 * time.Millisecond); fmt.Println("Task1 done") }},
		{"Task2", func() { time.Sleep(300 * time.Millisecond); fmt.Println("Task2 done") }},
		{"Task3", func() { time.Sleep(700 * time.Millisecond); fmt.Println("Task3 done") }},
	}
	RunTasks(tasks)
}
