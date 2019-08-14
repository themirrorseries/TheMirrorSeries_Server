package NetFrame

import (
	"bytes"
	"encoding/binary"
)

//encode 输出消息体前面一段bytes
//字段1：type		int32
//字段2：command	int32
//字段3：messages	[]byte		序列化后的消息体
//长度字段：	字段1之前增加int32字段表示要后面要发送的字段长度

type Encode struct {
	len      int32
	thetype  int32
	command  int32
	head     []byte
	writePos int32
}

func NewEncode(inputLen, inputType, inputCommand int32) *Encode {
	e := &Encode{
		len:     inputLen,
		thetype: inputType,
		command: inputCommand,
	}
	return e
}
func (enc *Encode) WriteInt32(num int32) {

	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.LittleEndian, num)
	copy(enc.head[enc.writePos:], buff.Bytes())
	enc.writePos += 4
}

func (enc *Encode) Write() {
	enc.head = make([]byte, 12)
	enc.writePos = 0
	enc.WriteInt32(enc.len)
	enc.WriteInt32(enc.thetype)
	enc.WriteInt32(enc.command)
}
func (enc *Encode) GetBytes() (head []byte) {
	return enc.head
}
