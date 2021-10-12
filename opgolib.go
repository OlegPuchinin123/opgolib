/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const VERSION = "2"

func SayHello() {
	fmt.Printf("opgolib is loaded !\n")
}

func Version() string {
	return VERSION
}

func HexDump(buf []byte, w io.Writer) {
	var (
		s    string
		n    int
		wbuf *bufio.Writer
	)
	n = 2
	wbuf = bufio.NewWriter(w)
	for _, b := range buf {
		s = fmt.Sprintf("%x", b)
		if len(s) == 2 {
			s = fmt.Sprintf("%x ", b)
		} else {
			s = fmt.Sprintf("0%x ", b)
		}
		if n == 16 {
			n = 1
			s = s + "\n"
		}
		wbuf.WriteString(s)
		n++
	}
	wbuf.WriteByte('\n')
	wbuf.Flush()
}

func RandomString(size int) string {
	var (
		b       []byte
		letters []byte
	)
	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b = make([]byte, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Seed() {
	rand.Seed(time.Now().UnixNano())
}
