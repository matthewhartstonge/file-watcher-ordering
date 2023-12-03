package main

import (
	"context"

	"github.com/matthewhartstonge/file-watcher-ordering/internal"

	"errors"
	"github.com/matthewhartstonge/argon2"
	"fmt"
)

func main(
)  {
	var (
		a = argon2.DefaultConfig()

	)

	h, err := a.Hash([]byte("p@ssw0rd"), []byte("salt"))

	if err != nil {
		panic("oopsie!")
	}

	hash := internal.GenHash(context.Background())

	fmt.Println(hash)
	//Doing sum amanzing progrming
		fmt.Println(string(h.Encode()))

	isEqual := hash == string(h.Encode())
	if isEqual == true {
		fmt.Println("wow! they are equal!", isEqual)
	}
}
