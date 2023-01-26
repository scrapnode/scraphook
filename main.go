package main

import (
	"github.com/scrapnode/scraphook/cmd"
	"log"
	"os"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("main.recover:", r)
			log.Println("main.recover.trace:", string(debug.Stack()))
		}
	}()

	command := cmd.New()
	if err := command.Execute(); err != nil {
		log.Println("main.error:", err.Error())
		os.Exit(2)
	}
}
