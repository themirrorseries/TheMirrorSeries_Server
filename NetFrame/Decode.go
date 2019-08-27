package NetFrame

import (
	"bytes"
	"encoding/binary"
)

//decode解析出前面一段字段
type Decode struct {
	Len     int32
	Thetype int32
	Command int32
	ReadPos int32
}

func (dec *Decode) Read(head []byte) {
	dec.ReadPos = 0
	binary.Read(bytes.NewBuffer(head[dec.ReadPos:]), binary.LittleEndian, &dec.Len)
	dec.ReadPos += int32(binary.Size(dec.Len))
	binary.Read(bytes.NewBuffer(head[dec.ReadPos:]), binary.LittleEndian, &dec.Thetype)
	dec.ReadPos += int32(binary.Size(dec.Thetype))
	binary.Read(bytes.NewBuffer(head[dec.ReadPos:]), binary.LittleEndian, &dec.Command)
	dec.ReadPos += int32(binary.Size(dec.Command))
}
