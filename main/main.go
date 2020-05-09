// https://www.hackerrank.com/challenges/queue-using-two-stacks/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntQueue struct {
	elements []int
}

func NewIntQueue() *IntQueue {
	var queue = &IntQueue{
		elements: make([]int, 0),
	}

	return queue
}

func (queue *IntQueue) Enqueue(elem int) {
	queue.elements = append(queue.elements, elem)
}

func (queue *IntQueue) Dequeue() int {
	var value = queue.elements[0]
	queue.elements = queue.elements[1:]

	return value
}

func (queue *IntQueue) Peek() int {
	return queue.elements[0]
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var numOfCasesStr, _ = reader.ReadString('\n')
	var numOfCases, _ = strconv.Atoi(strings.TrimSpace(numOfCasesStr))

	var queue = NewIntQueue()

	for i := 0; i < numOfCases; i++ {
		var commandStr, _ = reader.ReadString('\n')
		var commands = strings.Split(strings.TrimSpace(commandStr), " ")
		var command, _ = strconv.Atoi(commands[0])

		switch command {
		case 1:
			// Enqueue. Read the new element
			var elem, _ = strconv.Atoi(commands[1])
			queue.Enqueue(elem)

		case 2:
			// Dequeue
			queue.Dequeue()

		case 3:
			// Print
			fmt.Println(queue.Peek())
		}
	}
}
