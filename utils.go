package main

import (
	"github.com/gofrs/uuid"
)

func UUID() string {
	u2, _ := uuid.NewV4()
	return u2.String()
}
