package db

import (
	"serverGo/src/scribble"
)

var (
	DB, err = scribble.New("data", nil)
)


