package Handler

import (
	"../Global"
	"../NetFrame"
	"../proto/dto"
)

type HandlerData struct {
	command    int32
	messages   []byte
	bytesStart int32
	bytesEnd   int32
	client     *Global.ClientState
}
type HandlerFunc interface {
	ReceiveMessage()
}

func NewHandlerData(decode *NetFrame.Decode, msg []byte, _client *Global.ClientState) *HandlerData {
	handlerData := &HandlerData{
		command:    decode.Command,
		bytesStart: decode.ReadPos,
		bytesEnd:   decode.Len + 4,
		messages:   msg,
		client:     _client,
	}
	return handlerData
}

func Handler(msg []byte, client *Global.ClientState) {
	var decode NetFrame.Decode
	decode.Read(msg)
	handlerData := NewHandlerData(&decode, msg, client)
	var p HandlerFunc
	//HANDLER CENTER
	switch decode.Thetype {
	case int32(DTO.MsgTypes_TYPE_LOGIN):
		{
			p = &Login{handlerData}
			p.ReceiveMessage()
		}
		break

	case int32(DTO.MsgTypes_TYPE_USER):
		//user
		break
	case int32(DTO.MsgTypes_TYPE_MATCH):
		{
			p = &Match{handlerData}
			p.ReceiveMessage()
		}
		break
	case int32(DTO.MsgTypes_TYPE_FIGHT):
		{
			p = &Fight{handlerData}
			p.ReceiveMessage()
		}
		break
	default:
		break
	}
}
