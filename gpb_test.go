/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

import (
	"fmt"
	"os"
	"testing"
)

func packet_new(size int) (*GPB, int) {
	var (
		gpb *GPB
	)
	if size == 0 {
		gpb = NewGPB(0, true)
	} else {
		gpb = NewGPB(size, false)
	}
	gpb.W8(0x1)
	gpb.W16(0xff00)
	gpb.W16(0xee00)
	gpb.W32(0xee00ee00)
	gpb.WS("OLEG")
	gpb.W8(0x0)
	gpb.WS("Puchinin")
	gpb.WS("Puchinin 2")
	gpb.W8(0x0)
	//HexDump(gpb.buf, os.Stdout)

	return gpb, gpb.size
}

func packet2_new(size int) (*GPB, int) {
	var (
		gpb *GPB
	)
	if size == 0 {
		gpb = NewGPB(0, true)
	} else {
		gpb = NewGPB(size, false)
	}
	gpb.W8_signed(-2)
	gpb.W16_signed(-3)
	gpb.W32_signed(-1)
	gpb.W64_signed(-1000)
	return gpb, gpb.size
}

func read_packet(gpb *GPB, t *testing.T) {
	var (
		b       byte
		w       uint16
		w2      uint16
		d       uint32
		name    string
		family  string
		family2 string
	)
	b, _ = gpb.R8()
	w, _ = gpb.R16()
	w2, _ = gpb.R16()
	d, _ = gpb.R32()
	name, _ = gpb.RSZ()
	family, _ = gpb.RS(8)
	family2, _ = gpb.RS(10)

	if b != 0x1 {
		t.Error("R8 loose")
	}
	if w != 0xff00 {
		t.Error("R16 loose")
	}
	if w2 != 0xee00 {
		t.Error("R16 2 loose")
	}
	if d != 0xee00ee00 {
		t.Error("R32 loose")
	}
	if name != "OLEG" {
		t.Error("RSZ loose")
	}
	if family != "Puchinin" {
		t.Error("RS loose")
	}
	if family2 != "Puchinin 2" {
		t.Error("RS 2 loose")
	}
}

func TestGPB_1(t *testing.T) {
	var (
		gpb  *GPB
		gpb2 *GPB
		size int
	)
	gpb = NewGPB(0, true)
	gpb.WSZ("Ono odnako")
	size = gpb.size
	gpb2 = NewGPB(size, false)
	gpb2.WSZ("Ono odnako")
	//HexDump(gpb2.buf, os.Stdout)
}

func TestGPB_2(t *testing.T) {
	var (
		gpb  *GPB
		size int
		i    int8
		i2   int16
		i3   int32
		i4   int64
	)

	gpb, size = packet_new(0)
	gpb, size = packet_new(size)
	//HexDump(gpb.buf, os.Stdout)
	gpb.pos = 0
	read_packet(gpb, t)
	gpb, size = packet2_new(0)
	gpb, size = packet2_new(size)
	gpb.Set_pos(0)
	i, _ = gpb.R8_signed()
	i2, _ = gpb.R16_signed()
	i3, _ = gpb.R32_signed()
	i4, _ = gpb.R64_signed()
	fmt.Printf("%d %d %d %d\n", i, i2, i3, i4)
	HexDump(gpb.buf, os.Stdout)
}

func packet_map_new(size int) (*GPB, int) {
	var (
		gpb *GPB
		m   map[string]string
	)
	if size == 0 {
		gpb = NewGPB(0, true)
	} else {
		gpb = NewGPB(size, false)
	}
	m = make(map[string]string)
	m["One"] = "Oleg"
	m["Two"] = "Vasya"
	gpb.W_map(m)
	return gpb, gpb.size
}

func TestGPB_map(t *testing.T) {
	var (
		size       int
		gpb        *GPB
		m          map[string]string
		key, value string
	)
	gpb, size = packet_map_new(0)
	gpb, size = packet_map_new(size)
	gpb.Set_pos(0)
	m, _ = gpb.R_map()
	for key, value = range m {
		fmt.Printf("%s %s\n", key, value)
	}

}
