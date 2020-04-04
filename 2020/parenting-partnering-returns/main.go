package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Task struct {
	ID int
	Start int
	End   int
}

type Person struct {
	Tasks []Task
}

type ByStart []Task

func (t ByStart) Len() int {
	return len(t)
}

func (t ByStart) Less(i, j int) bool {
	return t[i].Start < t[j].Start
}

func (t ByStart) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (p *Person) isFree(newTask Task) bool {
	for _, t := range p.Tasks {
		if overlap(t, newTask) {
			return false
		}
	}

	return true
}

func (p *Person) assign(task Task) {
	p.Tasks = append(p.Tasks, task)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numTests, _ := strconv.Atoi(scanner.Text())

	for testCase := 0; testCase < numTests; testCase++ {
		scanner.Scan()
		numTasks, _ := strconv.Atoi(scanner.Text())

		tasks := make([]Task, 0)

		for numTask := 0; numTask < numTasks; numTask++ {
			scanner.Scan()
			tasks = append(tasks, schedule(scanner.Text(), numTask))
		}

		fmt.Printf("Case #%d: %s\n", testCase+1, check(tasks))
	}
}

func schedule(s string, id int) Task {
	arr := strings.Split(s, " ")
	start, _ := strconv.Atoi(arr[0])
	end, _ := strconv.Atoi(arr[1])

	return Task{id, start, end}
}

func overlap(task1, task2 Task) bool {
	if task1.Start <= task2.Start {
		return task1.End > task2.Start
	}

	return task2.End > task1.Start
}

func check(tasks []Task) string {

	sort.Sort(ByStart(tasks))

	cameron := &Person{Tasks: make([]Task, 0)}
	jamie := &Person{Tasks: make([]Task, 0)}

	for _, t := range tasks {
		if cameron.isFree(t) {
			cameron.assign(t)
			continue
		}

		if jamie.isFree(t) {
			jamie.assign(t)
			continue
		}

		return "IMPOSSIBLE"
	}

	output := make([]string, len(tasks))

	for _, t := range cameron.Tasks {
		output[t.ID] = "C"
	}

	for _, t := range jamie.Tasks {
		output[t.ID] = "J"
	}


	return strings.Join(output, "")
}