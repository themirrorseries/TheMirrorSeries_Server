package NetFrame

import (
	"bytes"
	"encoding/binary"
	)

//decode解析出前面一段字段                      
type Decode struct{
	len	int32
	thetype int32
	command int32
	readPos int32
}

func (dec *Decode)Read(head []byte){
	dec.readPos = 0;
	binary.Read(bytes.NewBuffer(head[dec.readPos:]), binary.BigEndian, &dec.len)
	dec.readPos += 4
	binary.Read(bytes.NewBuffer(head[dec.readPos:]), binary.BigEndian, &dec.thetype)
	dec.readPos += 4
	binary.Read(bytes.NewBuffer(head[dec.readPos:]), binary.BigEndian, &dec.command)
	dec.readPos += 4
}
