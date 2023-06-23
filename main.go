package main
import (
	"fmt"
	"bufio"
	"os"

	"knowledge_work_task/stacktrace"
	)
	
func main() {
	// goroutine 動作確認用コード
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panicが発生", r)
			}
		}()

		scanner := bufio.NewScanner(os.Stdin)
		// ユーザーからの入力を1行ずつ読み込んで出力する
		for scanner.Scan() {
			// 特定の文字列が入力されたらゴルーチンを終了させる
			if scanner.Text() == "test" {
				fmt.Println("入力された文字列は" + scanner.Text())
				fmt.Println("goroutine終了")
				break
			} else if scanner.Text() == "panic" {
				panic("")
				fmt.Println("goroutine終了")
				break
			} else {
				fmt.Println("入力された文字列は" + scanner.Text())
			}
		}
	}()


	stacktrace.ObserveGoroutine()
}