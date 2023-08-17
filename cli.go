package main

import (
    "fmt"

	"github.com/manifoldco/promptui"
)

func ChooseSong(songs []Song) (int, error) {
	prompt := promptui.Select{
		Label: "Select a song",
		Items: songs,
		Templates: &promptui.SelectTemplates {
			Label:    "╭─────────────────────────────────────╮\n│ {{ . }}",
			Active:   "> {{ .Title | cyan }} ({{ .Artist }})",
			Inactive: "  {{ .Title | cyan }} ({{ .Artist }})",
			//Selected: "{{ .Title | green }}",
		},
		Size: len(songs),
	}

	index, _, err := prompt.Run()
	if err != nil {
		return -1, err
	}

	fmt.Printf("You selected: %s by %s\n", songs[index].Title, songs[index].Artist)
	return index, nil
}
