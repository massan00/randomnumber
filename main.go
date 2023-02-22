package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("引数に5,6,7いずれかの数が必要です")
	}

	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("引数の数値が正しくありません")
	}
	if !(arg == 5 || arg == 6 || arg == 7) {
		log.Fatal("引数は5,6,7のいずれかを入力してください")
	}

	// 選択する数がすべて格納された
	allCount := rotoSlice(arg)

	result := choiceNum(allCount, arg)

	fmt.Println(result)
}

func rotoSlice(i int) []int {
	var sliceLength int
	if i == 5 {
		sliceLength = 31
	} else if i == 6 {
		sliceLength = 43
	} else if i == 7 {
		sliceLength = 37
	}

	a := make([]int, sliceLength)
	for i := range a {
		a[i] = i + 1
	}

	return a
}

// 引数にすべての数が入ったスライスとos.Argsのカウントを受け取って最終的に選ぶ数を返す
func choiceNum(allCount []int, args int) []int {
	result := make([]int, args)
	allNumber := allCount
	for i := 0; i < args; i++ {
		num := rand.Intn(len(allNumber))
		//選択した数値をresultに格納
		result[i] = allNumber[num]

		//スライスから選択した要素を削除
		allNumber = append(allNumber[:num], allNumber[num+1:]...)

	}
	return result
}
