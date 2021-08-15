package main

import (
	"log"
	"nitrohsu.com/futu/conf"
	"nitrohsu.com/futu/quote"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	quoteContext := quote.Quote{
		Config: conf.Config{
			Host: "127.0.0.1",
			Port: "11111",
		},
	}
	//
	err := quoteContext.Start()
	if err != nil {
		panic(err)
	}
	defer quoteContext.Close()

	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Printf("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
