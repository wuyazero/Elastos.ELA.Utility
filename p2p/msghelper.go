package p2p

import (
	"errors"
	"fmt"
	"net"
	"bytes"
)

const MaxBufLen = 1024 * 16

var (
	ErrDisconnected   = errors.New("disconnected")
	ErrUnmatchedMagic = errors.New("unmatched Magic")
)

// The interface to callback message read errors, message creation and decoded message.
type MsgDecodeHandler interface {
	// When something wrong on read or decode message
	// this method will callback the error
	OnDecodeError(err error)

	// After message header decoded, this method will be
	// called to create the message instance with the CMD
	// which is the message type of the received message
	OnMakeMessage(cmd string) (Message, error)

	// After message has been successful decoded, this method
	// will be called to pass the decoded message instance
	OnMessageDecoded(msg Message)
}

type MsgHelper struct {
	buf     []byte
	len     int
	magic   uint32
	conn    net.Conn
	handler MsgDecodeHandler
}

// NewMsgHelper create a new instance of *MsgHelper
func NewMsgHelper(magic uint32, conn net.Conn, handler MsgDecodeHandler) *MsgHelper {
	reader := new(MsgHelper)
	reader.magic = magic
	reader.conn = conn
	reader.handler = handler
	return reader
}

func (reader *MsgHelper) Magic() uint32 {
	return reader.magic
}

func (reader *MsgHelper) Conn() net.Conn {
	return reader.conn
}

func (reader *MsgHelper) Build(msg Message) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := msg.Serialize(buf)
	if err != nil {
		return nil, err
	}
	hdr, err := buildHeader(reader.magic, msg.CMD(), buf.Bytes()).Serialize()
	if err != nil {
		return nil, err
	}

	return append(hdr, buf.Bytes()...), nil
}

func (reader *MsgHelper) Read() {
	var buf = make([]byte, MaxBufLen)
	for {
		len, err := reader.conn.Read(buf[0:MaxBufLen-1])
		buf[MaxBufLen-1] = 0 //Prevent overflow
		switch err {
		case nil:
			reader.unpack(buf[:len])
		default:
			goto ERROR
		}
	}
ERROR:
	reader.handler.OnDecodeError(ErrDisconnected)
}

func (reader *MsgHelper) append(msg []byte) {
	reader.buf = append(reader.buf, msg...)
}

func (reader *MsgHelper) reset() {
	reader.buf = nil
	reader.len = 0
}

func (reader *MsgHelper) unpack(buf []byte) {
	if len(buf) == 0 {
		return
	}

	if reader.len == 0 { // Buffering message header
		index := HEADERLEN - len(reader.buf)
		if index > len(buf) { // header not finished, continue read
			index = len(buf)
			reader.append(buf[0:index])
			return
		}

		reader.append(buf[0:index])

		var header header
		err := header.Deserialize(reader.buf)
		if err != nil {
			fmt.Println("Get error message header, relocate the msg header")
			reader.reset()
			return
		}

		if header.Magic != reader.magic {
			reader.handler.OnDecodeError(ErrUnmatchedMagic)
			return
		}

		reader.len = int(header.Length)
		buf = buf[index:]
	}

	msgLen := reader.len

	if len(buf) == msgLen { // Just read the full message

		reader.append(buf[:])
		go reader.decode(reader.buf)
		reader.reset()

	} else if len(buf) < msgLen { // Read part of the message

		reader.append(buf[:])
		reader.len = msgLen - len(buf)

	} else { // Read more than the message

		reader.append(buf[0:msgLen])
		go reader.decode(reader.buf)
		reader.reset()
		reader.unpack(buf[msgLen:])
	}
}

func (reader *MsgHelper) decode(buf []byte) {
	if len(buf) < HEADERLEN {
		reader.handler.OnDecodeError(errors.New("message Length is not enough"))
		return
	}

	hdr, err := verify(buf)
	if err != nil {
		reader.handler.OnDecodeError(errors.New("verify message header error: " + err.Error()))
		return
	}

	msg, err := reader.handler.OnMakeMessage(hdr.GetCMD())
	if err != nil {
		reader.handler.OnDecodeError(errors.New("make message error, " + err.Error()))
		return
	}

	err = msg.Deserialize(bytes.NewReader(buf[HEADERLEN:]))
	if err != nil {
		reader.handler.OnDecodeError(errors.New("Deserialize message " + msg.CMD() + " error: " + err.Error()))
		return
	}

	reader.handler.OnMessageDecoded(msg)
}

func verify(buf []byte) (*header, error) {
	header := new(header)
	err := header.Deserialize(buf)
	if err = header.Verify(buf[HEADERLEN:]); err != nil {
		return nil, err
	}
	return header, nil
}
