package crc24q

//go:generate go run mktable.go

type Bytes interface {
	string | []byte
}

func Checksum[B Bytes](data B) uint32 {
	crc := uint32(0)
	for i := 0; i < len(data); i++ {
		crc = (crc << 8) ^ table[data[i]^byte(crc>>16)]
	}
	return crc & 0xFFFFFF
}

func Extract[B Bytes](data B, i int) uint32 {
	return uint32(data[i])<<16 | uint32(data[i+1])<<8 | uint32(data[i+2])
}
