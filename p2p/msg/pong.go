package msg

type Pong struct {
	Ping
}

func NewPong(height uint32) *Pong {
	pong := new(Pong)
	pong.Height = uint64(height)
	return pong
}

func (msg *Pong) CMD() string {
	return "pong"
}
