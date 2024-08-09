package main

import "math/rand"

func NewFid() string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	p1 := make([]rune, 10)
	for i := range p1 {
		p1[i] = chars[rand.Intn(len(chars))]
	}

	p2 := make([]rune, 11)
	for i := range p2 {
		p2[i] = chars[rand.Intn(len(chars))]
	}
	return string(p1) + "-" + string(p2)
}
