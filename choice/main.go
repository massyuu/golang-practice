package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 乱数のSeedを初期化
	rand.Seed(time.Now().UnixNano())

	// コマンドライン引数の数をカウント
	// 引数がなければエラーとする
	var c int
	c = len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s choice1 choise2...", os.Args[0])
		os.Exit(1)
	}

	// 乱数で選択したものを表示
	seed := rand.Intn(c)
	fmt.Println(seed)
	fmt.Println(os.Args[seed + 1])
}
