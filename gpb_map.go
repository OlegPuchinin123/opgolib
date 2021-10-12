
package opgolib


func (b *GPB) R_map() (map[string]string, error) {
	var (
		m map[string]string
		len uint8
		e error		
		i uint8
		key,value string

	)
	m = make(map[string]string)
	len, e = b.R8()
	if e != nil {
		return nil, nil
	}
	for i = 0; i < len; i++ {
		key,e = b.RSZ()
		if e != nil {
			return nil, e
		}
		value,e = b.RSZ()
		if e != nil {
			return nil, e
		}			
		m[key] = value
	}
	return m, e
}

func (b *GPB) W_map(m map[string]string) {
	b.W8(uint8(len(m)))
	for key, value := range m {
		b.WSZ(key)		
		b.WSZ(value)
	}
}
