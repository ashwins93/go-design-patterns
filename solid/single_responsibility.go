package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount int = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(entry string) int {
	entryCount++
	newEntry := fmt.Sprintf("%d %s", entryCount, entry)
	j.entries = append(j.entries, newEntry)
	return entryCount
}

/*
	func (j *Jounral) RemoveEntry(entryNo int) bool {
		...
	}
*/

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// so far Journal is still adhering to SRP but the next part is a bad idea

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(
		filename, []byte(j.String()), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(
		filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)),
		0644,
	)
}

func main() {
	j := Journal{}
	j.AddEntry("Entry 1")
	j.AddEntry("Entry 2")
	j.Save("bad.txt")

	p := Persistence{"\n"}

	p.SaveToFile(&j, "better.txt")
}
