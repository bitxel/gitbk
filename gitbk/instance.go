package gitbk

import (
	"github.com/bitxel/gitbk/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	done = make(chan bool)
	osSignal = make(chan os.Signal, 1)
)

func init() {
	signal.Notify(osSignal, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		for {
			s := <- osSignal
			log.Printf("signal (%d) %s received", s, s)
			done <- true
		}
	}()
}


func Start() {
	ticker := time.NewTicker(config.C.Global.BackupInterval)
	for {
		select {
		case <-done:
			log.Println("stop signal received, exit")
			return
		case <-ticker.C:
			log.Println("start to backup")
			SyncAll()
		}
	}
}

