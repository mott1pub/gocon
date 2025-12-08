package main

import (
	"hash/fnv"

	"github.com/google/uuid"
)

// hash the UUID into int
func GetRandomId() int {
	myuuid := uuid.New()
	h := fnv.New32a()
	h.Write(myuuid[:])
	hashedInit := int(h.Sum32())
	return hashedInit
}
