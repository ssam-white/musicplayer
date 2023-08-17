package main

import (
    "fmt"
    "log"
    "os"
    "encoding/json"
    "time"
    "io/ioutil"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Song struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Path   string `json:"path"`
}

func playMp3(filePath string, done chan bool) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	fmt.Println("Playing audio... Press Enter to stop.")

	<-done
}

func getSongs(songPath string) ([]Song, error) {
	file, err := ioutil.ReadFile(songPath)
	if err != nil {
		log.Fatal(err)
	}

	var songs []Song
	parseErr := json.Unmarshal(file, &songs)
	if parseErr != nil {
		return nil, parseErr
	}
	return songs, nil
}
