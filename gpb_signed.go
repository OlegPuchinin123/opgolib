package opgolib

import (
	"bytes"
	"encoding/binary"
)

func (b *GPB) R8_signed() (int8, error) {
	var (
		buf *bytes.Buffer
		e   error
		i   int8
	)
	buf = bytes.NewBuffer(b.buf[b.pos : b.pos+1])
	b.pos++
	e = binary.Read(buf, binary.LittleEndian, &i)
	return i, e
}

func (b *GPB) R16_signed() (int16, error) {
	var (
		buf *bytes.Buffer
		e   error
		i   int16
	)
	buf = bytes.NewBuffer(b.buf[b.pos : b.pos+2])
	b.pos += 2
	e = binary.Read(buf, binary.LittleEndian, &i)
	return i, e
}

func (b *GPB) R32_signed() (int32, error) {
	var (
		buf *bytes.Buffer
		e   error
		i   int32
	)
	buf = bytes.NewBuffer(b.buf[b.pos : b.pos+4])
	b.pos += 4
	e = binary.Read(buf, binary.LittleEndian, &i)
	return i, e
}

func (b *GPB) R64_signed() (int64, error) {
	var (
		buf *bytes.Buffer
		e   error
		i   int64
	)
	buf = bytes.NewBuffer(b.buf[b.pos : b.pos+8])
	b.pos += 8
	e = binary.Read(buf, binary.LittleEndian, &i)
	return i, e
}

func (b *GPB) W8_signed(i int8) error {
	var (
		buf *bytes.Buffer
		e   error
	)
	if b.counter {
		b.size++
		return nil
	}
	buf = bytes.NewBuffer(nil)
	e = binary.Write(buf, binary.LittleEndian, i)
	if e != nil {
		return e
	}
	e = b.WD(buf.Bytes())
	return e
}

func (b *GPB) W16_signed(i int16) error {
	var (
		buf *bytes.Buffer
		e   error
	)
	if b.counter {
		b.size += 2
		return nil
	}
	buf = bytes.NewBuffer(nil)
	e = binary.Write(buf, binary.LittleEndian, i)
	if e != nil {
		return e
	}
	return b.WD(buf.Bytes())
}

func (b *GPB) W32_signed(i int32) error {
	var (
		buf *bytes.Buffer
		e   error
	)
	if b.counter {
		b.size += 4
		return nil
	}
	buf = bytes.NewBuffer(nil)
	e = binary.Write(buf, binary.LittleEndian, i)
	if e != nil {
		return e
	}
	return b.WD(buf.Bytes())
}

func (b *GPB) W64_signed(i int64) error {
	var (
		buf *bytes.Buffer
		e   error
	)
	if b.counter {
		b.size += 8
		return nil
	}
	buf = bytes.NewBuffer(nil)
	e = binary.Write(buf, binary.LittleEndian, i)
	if e != nil {
		return e
	}
	return b.WD(buf.Bytes())
}
