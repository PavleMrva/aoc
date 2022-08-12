package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	horizontalPosition := 0
	depth := 0
	aim := 0
	for _, command := range strings.Split(string(content), "\n") {
		cmd := strings.Split(command, " ")

		if len(cmd) < 2 {
			continue
		}

		direction := cmd[0]
		units, _ := strconv.Atoi(cmd[1])

		if direction == "down" {
			aim += units
		} else if direction == "up" {
			aim -= units
		} else {
			horizontalPosition += units
			depth += aim * units
		}
	}

	fmt.Println(horizontalPosition * depth)
}
