package main

import (
	"github.com/satori/go.uuid"
)

func UUID() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}
