package daemon

import (
	"fmt"
	"time"
)

// 1秒おきにHello Worldを出力する関数
// SIGNALを受け取った場合、Doneを出力し正常終了する
func residentProcess() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		waitForInputSignal()
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			fmt.Println("Hello World")
		}
	}
}

func practice() {
	message := make(chan string)
	go func() { message <- "hello" }()
	receivedMessage := <-message
	fmt.Println(receivedMessage)
	// close(message)
	go func() { message <- "hello" }()
	fmt.Println(<-message)
}

func main() {
	residentProcess()
}
