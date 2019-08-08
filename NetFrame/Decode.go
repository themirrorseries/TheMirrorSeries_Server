package NetFrame

import (
	"bytes"
	"encoding/binary"
	)

//decode解析出前面一段字段                      
type Decode struct{
	Len	int32
	Thetype int32
	Command int32
	ReadPos int32
}

func (dec *Decode)Read(head []byte){
	dec.ReadPos = 0;
	binary.Read(bytes.NewBuffer(head[dec.ReadPos:]), binary.BigEndian, &dec.Len)
	dec.ReadPos += 4
	binary.Read(bytes.NewBuffer(head[dec.ReadPos:]), binary.BigEndian, &dec.Thetype)
	dec.ReadPos += 4
	binary.Read(bytes.NewBuffer(head[dec.ReadPos:]), binary.BigEndian, &dec.Command)
	dec.ReadPos += 4
}
