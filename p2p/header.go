package p2p

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

const (
	CMDLEN      = 12
	CMDOFFSET   = 4
	CHECKSUMLEN = 4
	HEADERLEN   = 24
)

var Magic uint32

type header struct {
	Magic    uint32
	CMD      [CMDLEN]byte
	Length   uint32
	Checksum [CHECKSUMLEN]byte
}

func buildHeader(cmd string, body []byte) *header {
	// Calculate Checksum
	checksum := Sha256D(body)

	header := new(header)
	// Write Magic
	header.Magic = Magic
	// Write CMD
	copy(header.CMD[:len(cmd)], cmd)
	// Write Length
	header.Length = uint32(len(body))
	// Write Checksum
	copy(header.Checksum[:], checksum[:CHECKSUMLEN])

	return header
}

func (header *header) Verify(buf []byte) error {
	// Verify Magic
	if header.Magic != Magic {
		return errors.New(fmt.Sprint("Unmatched Magic number ", header.Magic))
	}

	// Verify Checksum
	sum := Sha256D(buf)
	checksum := sum[:CHECKSUMLEN]
	if !bytes.Equal(header.Checksum[:], checksum) {
		return errors.New(
			fmt.Sprintf("Unmatched Checksum, expecting %s get $s",
				hex.EncodeToString(checksum),
				hex.EncodeToString(header.Checksum[:])))
	}

	return nil
}

func (header *header) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (header *header) Deserialize(buf []byte) error {
	cmd := buf[CMDOFFSET:CMDOFFSET+CMDLEN]
	end := bytes.IndexByte(cmd, 0)
	if end < 0 || end >= CMDLEN {
		return errors.New("Unexpected Length of CMD")
	}

	hdr := bytes.NewReader(buf[:HEADERLEN])
	return binary.Read(hdr, binary.LittleEndian, header)
}

func (header *header) GetCMD() string {
	end := bytes.IndexByte(header.CMD[:], 0)
	return string(header.CMD[:end])
}
