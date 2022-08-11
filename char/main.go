// Package main
package main

import (
	"fmt"
	"unicode/utf8"
)

// chapter9/sources/go-character-set-encoding/rune_encode_and_decode.go// rune -> []byte
func encodeRune(r rune) []byte {
	var buf [4]byte
	n := utf8.EncodeRune(buf[:], r)
	return buf[:n]
}

func decodeRune(b []byte) (rune, int) {
	return utf8.DecodeRune(b)
}

func main() {
	// r := 'a' // 'a' is a rune ascii code 97
	var r rune = 0x4E2D // '中' is a rune unicode code 20013
	b := encodeRune(r)
	r2, n := decodeRune(b)
	fmt.Printf("0x%X\n", b)  // "61"
	fmt.Printf("0x%X\n", r2) // "a"
	fmt.Printf("%d\n", n)    // 1
}
