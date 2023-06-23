package stacktrace

import (
	"fmt"
	"runtime"
	"time"
	"bytes"
	"regexp"
)

func ObserveGoroutine() error {
	buf := make([]byte, 1024)
	var previous_goroutine_ids [][]string

	for {
		fmt.Print("\033[2J\033[H")
		// 全てのgoroutineのスタックトレースを取得する
		n := runtime.Stack(buf, true)
		stackTrace := string(bytes.TrimRight(buf[:n], "\x00"))
		goroutine_ids := extractGoroutineID(stackTrace)

		fmt.Printf("< goroutine status >\n")

		for _, goroutine_id := range goroutine_ids {
			fmt.Println("goroutine " + goroutine_id[1]+ " is living")

			elements := make(map[string]bool)

			for _, goroutine_id := range goroutine_ids {
				elements[goroutine_id[1]] = true
			}

			for _, previous_goroutine_id := range previous_goroutine_ids {
				if !elements[previous_goroutine_id[1]] {
					fmt.Println("goroutine " + previous_goroutine_id[1] + " is dead")
				}
			}
		}

		previous_goroutine_ids = goroutine_ids // 前回のループのgoroutine_idsを更新

		time.Sleep(10 * time.Second)
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