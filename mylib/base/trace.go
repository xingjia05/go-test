// Package base ...
package base

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func MyTrace1() {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}

}

// MyTrace ...
func MyTrace() {
	f, err := os.Create("/root/go/src/myapp/trace.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("1111")
}
