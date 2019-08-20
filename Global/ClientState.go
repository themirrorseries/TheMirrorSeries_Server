package Global

import "net"

type ClientState struct {
	Client   net.Conn
	IsMatch  bool
	IsFight  bool
	RoomID   int32
	RoomSeat int32
	PlayerID int32
}

func NewClientState(client net.Conn) *ClientState {
	c := &ClientState{
		Client:   client,
		IsMatch:  false,
		IsFight:  false,
		RoomID:   -1,
		RoomSeat: -1,
		PlayerID: -1,
	}
	return c
}
func (c *ClientState) MatchOut() {
	RoomCache.RemovePlayer(c.PlayerID, c.Client)
}

func (c *ClientState) FightOut() {
	RoomMng[c.RoomID].PlayerLeave(c.RoomSeat)
}
