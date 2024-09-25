package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d *deck) generate() {
	faces := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	units := []string{"King", "Queen", "Jack"}
	for _, face := range faces {
		for _, unit := range units {
			card := fmt.Sprint(face, "Of", unit)
			*d = append(*d, card)
		}
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}
func (d deck) saveToFile(fileName string) error {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		return fmt.Errorf("error occured: %v", err)
	}
	return nil
}
func (d *deck) openFile(fileName string) {
	bs, err := os.OpenFile(fileName, os.O_CREATE, 0666)
	if err != nil {
		*d = deck{}
	}
	data, err := io.ReadAll(bs)
	temp := strings.Split(string(data), ",")

	*d = append(*d, temp...)

}
func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d {
		newI := r.Intn(len(d))
		d[i], d[newI] = d[newI], d[i]
	}
}
