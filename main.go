package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/the0val/midi-tab/melody"
)

func main() {
	midiPath := flag.String("file", "melody.mid", "Path to midi file of the melody.")
	templatePath := flag.String("template", "harmonica.tabtemplate", "Path to template file")
	transpose := flag.Int("transpose", 0, "number of steps to transpose melody before convertion. Can be negative.")
	flag.Parse()

	m, err := melody.ReadFile(*midiPath)
	if err != nil {
		log.Fatal(err)
	}
	m.Transpose(*transpose)

	templateText, err := ioutil.ReadFile(*templatePath)
	if err != nil {
		log.Fatal(err)
	}
	t, err := melody.ParseTabTemplate(templateText)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Tabulate(m))
}
