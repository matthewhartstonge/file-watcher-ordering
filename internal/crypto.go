package internal

import (
	"context"
	"github.com/matthewhartstonge/argon2"
)

func GenHash(_ context.Context) string {
	a := argon2.DefaultConfig()
	h, _ := a.Hash([]byte("p@ssw0rd"), []byte("salt"))

	return string(h.Encode())
}
