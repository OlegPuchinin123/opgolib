/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package opgolib

func (b *GPB) R8_all(count int) ([]byte, error) {
	return b.RD(count)
}

func (b *GPB) R16_all(count int) ([]uint16, error) {
	var (
		e     error
		w_arr []uint16
		w     uint16
		i     int
	)
	w_arr = make([]uint16, count)
	for i = 0; i < count; i++ {
		w, e = b.R16()
		if e != nil {
			return nil, e
		}
		w_arr[i] = w
	}
	return w_arr, nil
}

func (b *GPB) R32_all(count int) ([]uint32, error) {
	var (
		e     error
		d_arr []uint32
		d     uint32
		i     int
	)
	d_arr = make([]uint32, count)
	for i = 0; i < count; i++ {
		d, e = b.R32()
		if e != nil {
			return nil, e
		}
		d_arr[i] = d
	}
	return d_arr, nil
}

func (b *GPB) RSZ_all(count int) ([]string, error) {
	var (
		arr []string
		idx int
		e   error
	)
	arr = make([]string, count)
	for idx = 0; idx < count; idx++ {
		arr[idx], e = b.RSZ()
		if e != nil {
			return nil, e
		}
	}
	return arr, nil
}

func (b *GPB) W8_all(bs ...byte) error {
	return b.WD(bs)
}

func (b *GPB) W16_all(w ...uint16) error {
	var (
		e   error
		one uint16
	)
	for _, one = range w {
		e = b.W16(one)
		if e != nil {
			return e
		}
	}
	return nil
}

func (b *GPB) W32_all(d ...uint32) error {
	var (
		e   error
		one uint32
	)
	for _, one = range d {
		e = b.W32(one)
		if e != nil {
			return e
		}
	}
	return nil
}

func (b *GPB) WSZ_all(arr ...string) error {
	var (
		e error
		s string
	)
	for _, s = range arr {
		e = b.WSZ(s)
		if e != nil {
			return e
		}
	}
	return nil
}
