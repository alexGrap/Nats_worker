package db

import (
	_ "github.com/lib/pq"
)

type Cash struct {
	memory map[string][]byte
}

var Memory Cash
