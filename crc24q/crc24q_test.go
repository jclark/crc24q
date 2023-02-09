package crc24q

import "testing"

// Example from RTCM 10403.2, section 4.2
var rtcmExample = []byte{
	0xD3, 0x00, 0x13, 0x3E, 0xD7, 0xD3, 0x02, 0x02, 0x98, 0x0E, 0xDE, 0xEF, 0x34, 0xB4, 0xBD, 0x62,
	0xAC, 0x09, 0x41, 0x98, 0x6F, 0x33, 0x36, 0x0B, 0x98,
}

func TestCRC24Q(t *testing.T) {
	testCRC24Q(t, "slice", rtcmExample)
	testCRC24Q(t, "string", string(rtcmExample))
}

func testCRC24Q[B Bytes](t *testing.T, name string, data B) {
	n := len(data) - 3
	actual := Checksum(data[0:n])
	expected := Extract(data, n)
	if actual != expected {
		t.Fatalf("%s: incorrect checksum: %06X", name, actual)
	}
}
