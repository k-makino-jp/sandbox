// mock作成コマンド
// mockgen -self_package=daemon -source=signal.go -destination=mock_daemon/mock_signal.go -package=daemon
package daemon

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Signal シグナル処理向けインターフェース
type Signal interface {
	Notify(c chan<- os.Signal, sig ...os.Signal)
}

type signalImpl struct{}

func (s signalImpl) Notify(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, sig...)
}

type Daemon interface {
	WaitForInputSignal()
}

type daemonImpl struct {
	signal Signal
}

func (d daemonImpl) WaitForInputSignal() {
	sig := make(chan os.Signal, 1)
	defer close(sig)
	d.signal.Notify(
		sig,
		// syscall.SIGKILL, // Signal Kill
		// syscall.SIGTERM, // Signal Exit (kill command default)
		// syscall.SIGINT,  // Signal Interrupt
		os.Interrupt, // SIGINT
		os.Kill)      // SIGKILL
	s := <-sig
	fmt.Println("SIGNAL:", s)
}

func waitForInputSignal() {
	sig := make(chan os.Signal, 1)
	defer close(sig)
	signal.Notify(
		sig,
		syscall.SIGKILL, // Signal Kill
		syscall.SIGTERM, // Signal Exit (kill command default)
		syscall.SIGINT,  // Signal Interrupt
		os.Interrupt,    // SIGINT
		os.Kill)         // SIGKILL
	s := <-sig
	fmt.Println("SIGNAL:", s)
}
