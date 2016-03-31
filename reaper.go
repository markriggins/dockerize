package main

// Adapted from https://github.com/miekg/dinit/sigchld.go which was
// Adapted from https://github.com/ramr/go-reaper/blob/master/reaper.go
// No license published there...

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type OnDeathFunc func(childPid int) error

func catchAllChildSignals() {
	var sigs = make(chan os.Signal, 3)
	var sig os.Signal
	signal.Notify(sigs, syscall.SIGCHLD)

	for {
		sig = <-sigs
		log.Printf("received signal %s\n", sig)
	}
}

// waiting for dying children cleans up zombies
// No good parent would want their child to wind up as a zombie -- right?
func waitForChildrenToDie() {
	var wstatus syscall.WaitStatus

	for {
		pid, err := syscall.Wait4(-1, &wstatus, 0, nil)
		if err == syscall.EINTR {
			log.Printf("wait4 pid %d interrupted: %+v", pid, wstatus)
		} else if err == syscall.ECHILD {
			log.Printf("wait4: No more children")
			time.Sleep(2000)
		} else {
			log.Printf("pid %d, finished, wstatus: %+v", pid, wstatus)
		}
	}
}

func reapChildren(onDeathFunc OnDeathFunc) {

	if onDeathFunc == nil {
		onDeathFunc = func(childPid int) error { return nil }
	}

	go catchAllChildSignals()
	go waitForChildrenToDie()
}
