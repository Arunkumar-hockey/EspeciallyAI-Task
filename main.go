package main

import (
	"encoding/binary"
	"fmt"
)

type DecodedStruct struct {
	Value1        uint16
	Value2        string
	Value3        uint8
	Value4        string
	Value5        uint16
	Value6        string
	Value7        uint32
}

func decodePacket(decoded DecodedStruct,packet []byte) DecodedStruct {
	

	decoded.Value1 = binary.BigEndian.Uint16(packet[:2])
	decoded.Value2 = string(packet[2:14])
	decoded.Value3 = packet[14]
	decoded.Value4 = string(packet[15:23])
	decoded.Value5 = binary.BigEndian.Uint16(packet[23:25])
	decoded.Value6 = string(packet[25:40])
	decoded.Value7 = binary.BigEndian.Uint32(packet[40:44])

	return decoded
}

func main() {
	packet := []byte{0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67, 0x38, 0x64, 0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03, 0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74, 0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73, 0x07, 0x5B, 0xCD, 0x15}
	var decoded DecodedStruct
	decodedStruct := decodePacket(decoded,packet)
	fmt.Printf("%+v\n", decodedStruct)
	//Print Output = {Value1:1234 Value2:keepdecoding Value3:56 Value4:dontstop Value5:789 Value6:congratulations Value7:123456789}
}
// {1234, "keepdecoding", 56, "dontstop", 789, "congratulations", , 123456789}