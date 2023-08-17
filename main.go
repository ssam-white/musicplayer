package main

import (
    "log"
)

func main() {
	songs, err := getSongs("songs.json")
	if err != nil {
		log.Fatal(err)
	}

	index, promptErr := ChooseSong(songs)
	if promptErr != nil {
		log.Fatal(promptErr)
	}

	var filePath string = songs[index].Path
	done := make(chan bool)
	go playMp3(filePath, done)

    <- done
}
