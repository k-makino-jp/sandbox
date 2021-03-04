package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGKILL, // Signal Kill
		syscall.SIGTERM, // Signal Exit (kill command default)
		syscall.SIGINT,  // Signal Interrupt
		os.Interrupt)    // for Windows

	s := <-sig

	fmt.Println("signal: ", s)
}

func ticker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

func waitForInputSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		syscall.SIGKILL, // Signal Kill
		syscall.SIGTERM, // Signal Exit (kill command default)
		syscall.SIGINT,  // Signal Interrupt
		os.Interrupt)    // for Windows
	s := <-sig
	fmt.Println("SIGNAL:", s)
}

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
			// panic("PANIC")
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

func infiniteLoop(done <-chan bool, ticker *time.Ticker) {
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			fmt.Println("Hello World")
			// panic("PANIC")
		}
	}
}

func channelHandler() {
	done := make(chan bool)
	go func() {
		waitForInputSignal()
		done <- true
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	infiniteLoop(done, ticker)
}

func main() {
	// residentProcess()
	channelHandler()
}

// type timeTicker interface {
// }
// type timeTicker struct{}
