/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"errors"
)

type GPB struct {
	buf        []byte
	pos        int
	counter    bool
	size       int
	OutOfBound error
}

func NewGPB(size int, counter bool) *GPB {
	b := new(GPB)
	if size > 0 {
		b.buf = make([]byte, size)
		b.size = size
	} else {
		b.size = 0
	}
	b.pos = 0
	b.counter = counter
	b.OutOfBound = errors.New("Out of bound")
	return b
}

func NewGPBBuf(buf []byte) *GPB {
	b := new(GPB)
	b.size = len(buf)
	b.buf = buf
	b.pos = 0
	b.counter = false
	b.OutOfBound = errors.New("Out of bound")
	return b
}

func (b *GPB) Set_buf(buf []byte) {
	b.buf = buf
	b.size = len(buf)
	b.pos = 0
	b.counter = false
}

func (b *GPB) Set_pos(pos int) {
	b.pos = pos
}

func (b *GPB) Get_pos() int {
	return b.pos
}

func (b *GPB) Get_buf() []byte {
	return b.buf
}

func (b *GPB) Get_size() int {
	return b.size
}

func (b *GPB) W8(c uint8) error {
	if b.counter {
		b.size++
		return nil
	}
	if b.pos >= len(b.buf) {
		return b.OutOfBound
	}
	b.buf[b.pos] = c
	b.pos++
	return nil
}

func (b *GPB) W16(w uint16) error {
	var (
		e      error
		c1, c2 byte
	)
	if b.counter {
		b.size += 2
		return nil
	}
	c1 = byte(w & 0xFF)
	e = b.W8(c1)
	if e != nil {
		return e
	}
	c2 = byte((w >> 8) & 0xFf)
	e = b.W8(c2)
	if e != nil {
		return e
	}
	return nil
}

func (b *GPB) W32(dw uint32) error {
	var (
		e      error
		w1, w2 uint16
	)
	if b.counter {
		b.size += 4
		return nil
	}
	w1 = uint16(dw & 0xFFFF)
	e = b.W16(w1)
	if e != nil {
		return e
	}
	w2 = uint16((dw >> 16) & 0xFFFF)
	e = b.W16(w2)
	if e != nil {
		return e
	}
	return nil
}

func (b *GPB) W64(qw uint64) error {
	var (
		e        error
		dw1, dw2 uint32
	)
	if b.counter {
		b.size += 8
		return nil
	}
	dw1 = uint32(qw & 0xFFFFFFFF)
	e = b.W32(dw1)
	if e != nil {
		return e
	}
	dw2 = uint32((qw >> 32) & 0xFFFFFFFF)
	e = b.W32(dw2)
	if e != nil {
		return e
	}
	return nil
}

func (b *GPB) WS(str string) error {
	var (
		str_len int
	)
	str_len = len(str)
	if b.counter {
		b.size += str_len
		return nil
	}

	if (b.pos + str_len) >= len(b.buf) {
		return errors.New("Out of bound")
	}

	copy_ret := copy(b.buf[b.pos:], []byte(str))
	if copy_ret != str_len {
		return b.OutOfBound
	}
	b.pos += str_len
	return nil
}

func (b *GPB) WD(buf []byte) error {
	var (
		buf_len int
	)
	buf_len = len(buf)
	if b.counter {
		b.size += buf_len
		return nil
	}

	if (b.pos + buf_len) > len(b.buf) {
		return errors.New("Out of bound")
	}

	copy_ret := copy(b.buf[b.pos:], buf)
	if copy_ret != buf_len {
		return b.OutOfBound
	}
	b.pos += buf_len
	return nil
}

func (b *GPB) WSZ(str string) error {
	var (
		str_len int
	)
	str_len = len(str)
	if b.counter {
		b.size += str_len + 1
		return nil
	}
	if b.pos >= (b.size + str_len + 1) {
		return b.OutOfBound
	}
	copy_ret := copy(b.buf[b.pos:], []byte(str))
	if copy_ret != str_len {
		return b.OutOfBound
	}
	b.pos += str_len
	return b.W8(0)
}

func (b *GPB) R8() (byte, error) {
	var (
		c byte
	)

	if b.pos >= b.size {
		return 0, b.OutOfBound
	}

	c = b.buf[b.pos]
	b.pos++
	return c, nil
}

func (b *GPB) R16() (uint16, error) {
	var (
		w      uint16
		c1, c2 byte
		e      error
	)
	c1, e = b.R8()
	if e != nil {
		return 0, e
	}
	c2, e = b.R8()
	if e != nil {
		return 0, e
	}
	w = uint16(c1) | (uint16(c2) << 8)
	return w, nil
}

func (b *GPB) R32() (uint32, error) {
	var (
		dw     uint32
		w1, w2 uint16
		e      error
	)

	w1, e = b.R16()
	if e != nil {
		return 0, e
	}

	w2, e = b.R16()
	if e != nil {
		return 0, e
	}

	dw = uint32(w1) | (uint32(w2) << 16)
	return dw, nil
}

func (b *GPB) R64() (uint64, error) {
	var (
		qw       uint64
		dw1, dw2 uint32
		e        error
	)
	dw1, e = b.R32()
	if e != nil {
		return 0, e
	}
	dw2, e = b.R32()
	if e != nil {
		return 0, e
	}
	qw = uint64(dw1) | (uint64(dw2) << 32)
	return qw, nil
}

func (b *GPB) RS(asize int) (string, error) {
	var (
		high int
		s    string
	)
	if (b.pos + asize) > len(b.buf) {
		return "", b.OutOfBound
	}
	high = b.pos + asize
	s = string(b.buf[b.pos:high])
	b.pos += asize
	return s, nil
}

func (b *GPB) RD(asize int) ([]byte, error) {
	var (
		high int
		buf  []byte
	)
	if (b.pos + asize) > len(b.buf) {
		return nil, b.OutOfBound
	}
	high = b.pos + asize
	buf = b.buf[b.pos:high]
	b.pos += asize
	return buf, nil
}

func (b *GPB) RSZ() (string, error) {
	var (
		buf_len int
		high    int
		zero    bool
		s       string
	)
	zero = false
	buf_len = len(b.buf)
	for high = b.pos; high < buf_len; high++ {
		if b.buf[high] == 0 {
			zero = true
			break
		}
	}
	if zero {
		if high == b.pos {
			b.pos++
			return "", nil
		} else {
			s = string(b.buf[b.pos:high])
			b.pos = high + 1
			return s, nil
		}
	} else {
		return "", b.OutOfBound
	}
	return s, nil
}
