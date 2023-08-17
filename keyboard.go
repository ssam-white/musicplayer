package main


import (
    "log"
	"github.com/eiannone/keyboard"
)

func listenForEnter(enterPressed chan bool) {
	for {
		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if key == keyboard.KeyEnter {
			enterPressed <- true
			break
		}
	}
}
