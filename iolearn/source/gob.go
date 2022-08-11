package source

import (
	"encoding/gob"
	"os"
)

type Player struct {
	Name  string
	Score int
}

//f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

func writeByGob(plays []*Player) error {
	f, err := os.Create("data.gob")
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(f)
	if err != nil {
		return err
	}
	for _, p := range plays {
		err := enc.Encode(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func readByGob(plays []*Player) error {
	f, err := os.Open("data.gob")
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(f)
	if err != nil {
		return err
	}
	for _, p := range plays {
		err := dec.Decode(p)
		if err != nil {
			return err
		}
	}

	return nil
}
