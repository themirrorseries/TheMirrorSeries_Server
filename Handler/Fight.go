package Handler

type Fight struct{
	command int32
	messages []byte
	bytesStart int32
	bytesEnd int32
}

func NewFight(c, start, end int32, msg []byte) *Match{
	match := &Match{
		command:	c,
		bytesStart:	start,
		bytesEnd:	end,
		messages:	msg,
	}
	return match
}