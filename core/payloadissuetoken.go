package core

import (
	"io"
)

const IssueTokenPayloadVersion byte = 0x00

type PayloadIssueToken struct{}

func (t *PayloadIssueToken) Data(version byte) []byte {
	return []byte{0}
}

func (t *PayloadIssueToken) Serialize(w io.Writer, version byte) error {
	return nil
}

func (t *PayloadIssueToken) Deserialize(r io.Reader, version byte) error {
	return nil
}
