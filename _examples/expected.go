package main

import (
	"context"
	"fmt"

	"github.com/matthewhartstonge/argon2"

	"github.com/matthewhartstonge/file-watcher-ordering/internal"
)

func main() {
	a := argon2.DefaultConfig()

	h, err := a.Hash([]byte("p@ssw0rd"), []byte("salt"))
	if err != nil {
		panic("oopsie!")
	}

	hash := internal.GenHash(context.Background())

	fmt.Println(hash)
	// Doing sum amanzing progrming
	fmt.Println(string(h.Encode()))

	isEqual := hash == string(h.Encode())
	if isEqual == true {
		fmt.Println("wow! they are equal!", isEqual)
	}
}
