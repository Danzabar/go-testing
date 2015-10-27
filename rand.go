package main

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(n int) string {
	part := make([]byte, n)
	for i := range part {
		part[i] = letters[rand.Intn(len(letters))]
	}
	return string(part)
}
