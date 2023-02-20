package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	num := []int{15, 10, 9, 11, 19, 32, 8, 20, 21, 22, 38, 7, 5}
	totalcount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	allnum := make([]int, 0, totalcount)

	for {
		n := rand.Intn(len(num))
		exists := false
		for _, v := range allnum {
			if n == v {
				exists = true
			}
		}
		if !exists {
			allnum = append(allnum, n)
		}
		if len(allnum) == totalcount {
			break
		}
	}

	for _, i := range allnum {
		fmt.Printf("%d ", num[i])
	}
	fmt.Println("")
}
