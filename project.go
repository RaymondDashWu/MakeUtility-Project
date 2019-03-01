// https://github.com/dyne/file-extension-list

package main

import (
	"log"
	"os"
)

// Pseudo brainstorm
// As you plug in an SD card, automatically moves to predesignated spot
// How to detect plug in of an sd card?
// Command line interface to determine where to move to
// TODO: account for file extension types and move them according to category in linked github

// notes
// get list of what's been plugged in /devs/tty

func main() {
	oldLocation := "OLD LOCATION THAT ACCEPTS MULTIPLE FILE TYPES. CERTAIN PRECODED EXTENSIONS"
	newLocation := "Specify new location through command line"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
