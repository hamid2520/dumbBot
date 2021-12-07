package main

import (
	"fmt"

	"github.com/hamid2520/dumbBot/bot"
	"github.com/hamid2520/dumbBot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return
}
