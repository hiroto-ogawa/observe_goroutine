package stacktrace

import (
	"fmt"
	"runtime"
	"time"
	"bytes"
	"regexp"
)

func ObserveGoroutine() error {
	// スタックトレース
	buf := make([]byte, 1024)

	for {
		// すべてのゴルーチンのスタックトレースを取得
		// falseにすると現在動いているgoroutineしかとらないようにすることができる
		n := runtime.Stack(buf, true)
		stackTrace := string(bytes.TrimRight(buf[:n], "\x00"))
		goroutine_ids := extractGoroutineID(stackTrace)

		fmt.Printf("< Stack trace >\n")

		for _, goroutine_id := range goroutine_ids {
			fmt.Println("goroutine " + goroutine_id[1]+ " is living")
		}
		// bufの中からスタックトレースが書き込まれた部分を表示される
		fmt.Printf("< Stack trace >\n%s\n", stackTrace)

		time.Sleep(3 * time.Second)
	}
}

// 文字列から数字を抽出する関数
func extractGoroutineID(stackTrace string) [][]string {
	pattern := regexp.MustCompile(`goroutine (\d+)`)

	// 正規表現パターンにマッチする部分を抽出
	matches := pattern.FindAllStringSubmatch(stackTrace, -1)

	// マッチした数字を出力
	// for _, match := range matches {
	// 	fmt.Println(match[1])
	// }

	return matches
}