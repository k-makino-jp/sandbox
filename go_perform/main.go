package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	fmt.Println("start")
	exec()
	fmt.Println("end")
}

func exec() {
	var someMap = make(map[string]string)

	doNothing("foo", 100000)
	addMap(someMap, "bar", 10000)
	addMap(someMap, "buzz", 1000000)
}

func addMap(s map[string]string, prefix string, count int) {
	for i := 0; i < count; i++ {
		key := prefix + "key" + strconv.Itoa(i)
		val := "value" + strconv.Itoa(i)
		s[key] = val
	}
}

func doNothing(prefix string, count int) {
	for i := 0; i < count; i++ {
		key := prefix + "key" + strconv.Itoa(i)
		val := "value" + strconv.Itoa(i)
		key = key + val
	}
}
