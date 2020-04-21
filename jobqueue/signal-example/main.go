package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("Got:", sig)
		signal.Stop(sigs)
		time.Sleep(1 * time.Second)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
