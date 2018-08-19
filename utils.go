package main

import (
	"github.com/satori/go.uuid"
)

func UUID() string {
	u2,_ := uuid.NewV4()
	return u2.String()
}
