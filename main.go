package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/the0val/midi-tab/melody"
)

func main() {
	tabText, err := ioutil.ReadFile("melody/testdata/harmonica-tabtemlate")
	if err != nil {
		log.Fatal(err)
		return
	}

	tab, err := melody.ParseTabTemplate(tabText)
	if err != nil {
		log.Fatal(err)
		return
	}

	mel, err := melody.ReadFile("melody/testdata/testmelody.mid")
	if err != nil {
		log.Fatal(err)
		return
	}

	mel.Transpose(-3 + 12)

	fmt.Println(tab.Tabulate(mel))
}
