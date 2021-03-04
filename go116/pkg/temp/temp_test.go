package main

import (
	"testing"
	"time"
)

func Test_infiniteLoop(t *testing.T) {
	type args struct {
		done   chan bool
		ticker *time.Ticker
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "infiniteLoop 3ループ目完了後にシグナルが送信されたとき ループが終了すること",
			args: args{
				done:   make(chan bool),
				ticker: time.NewTicker(time.Second),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				time.Sleep(time.Second)
				tt.args.done <- true
				tt.args.ticker.Stop()
			}()
			infiniteLoop(tt.args.done, tt.args.ticker)
		})
	}
}

func Test_channelHandler(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channelHandler()
		})
	}
}
