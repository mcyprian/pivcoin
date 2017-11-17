package uuid

import (
	"crypto/rand"
	"io"
)

func GenerateID() []byte {
	id := make([]byte, 16)
	r := rand.Reader
	if _, err := io.ReadFull(r, id); err != nil {
		panic(err) // This shouldn't happen
	}
	return id
}
