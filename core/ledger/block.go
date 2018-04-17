package ledger

import (
	"io"
	"bytes"
	"errors"

	"github.com/elastos/Elastos.ELA.Utility/crypto"
	. "github.com/elastos/Elastos.ELA.Utility/common"
	tx "github.com/elastos/Elastos.ELA.Utility/core/transaction"
	"github.com/elastos/Elastos.ELA.Utility/common/serialize"
)

const (
	BlockVersion     uint32 = 0
	GenesisNonce     uint32 = 2083236893
	InvalidBlockSize int    = -1
)

type Block struct {
	Header       *Header
	Transactions []*tx.Transaction
}

func (b *Block) Serialize(w io.Writer) error {
	b.Header.Serialize(w)
	err := serialize.WriteUint32(w, uint32(len(b.Transactions)))
	if err != nil {
		return errors.New("Block item Transactions length serialization failed.")
	}

	for _, transaction := range b.Transactions {
		transaction.Serialize(w)
	}
	return nil
}

func (b *Block) Deserialize(r io.Reader) error {
	if b.Header == nil {
		b.Header = new(Header)
	}
	b.Header.Deserialize(r)

	//Transactions
	var i uint32
	Len, err := serialize.ReadUint32(r)
	if err != nil {
		return err
	}
	var txhash Uint256
	var tharray []Uint256
	for i = 0; i < Len; i++ {
		transaction := new(tx.Transaction)
		transaction.Deserialize(r)
		txhash = transaction.Hash()
		b.Transactions = append(b.Transactions, transaction)
		tharray = append(tharray, txhash)
	}

	merkleRoot, err := crypto.ComputeRoot(tharray)
	if err != nil {
		return errors.New("Block Deserialize merkleTree compute failed")
	}
	b.Header.MerkleRoot = merkleRoot

	return nil
}

func (b *Block) Trim(w io.Writer) error {
	b.Header.Serialize(w)
	err := serialize.WriteUint32(w, uint32(len(b.Transactions)))
	if err != nil {
		return errors.New("Block item Transactions length serialization failed.")
	}
	for _, transaction := range b.Transactions {
		temp := *transaction
		hash := temp.Hash()
		hash.Serialize(w)
	}
	return nil
}

func (b *Block) FromTrimmedData(r io.Reader) error {
	if b.Header == nil {
		b.Header = new(Header)
	}
	b.Header.Deserialize(r)

	//Transactions
	var i uint32
	Len, err := serialize.ReadUint32(r)
	if err != nil {
		return err
	}
	var txhash Uint256
	var tharray []Uint256
	for i = 0; i < Len; i++ {
		txhash.Deserialize(r)
		b.Transactions = append(b.Transactions, tx.NewTrimmed(txhash))
		tharray = append(tharray, txhash)
	}

	b.Header.MerkleRoot, err = crypto.ComputeRoot(tharray)
	if err != nil {
		return errors.New("Block Deserialize merkleTree compute failed")
	}

	return nil
}

func (tx *Block) GetSize() int {
	var buffer bytes.Buffer
	if err := tx.Serialize(&buffer); err != nil {
		return InvalidBlockSize
	}

	return buffer.Len()
}

func (b *Block) Hash() Uint256 {
	return b.Header.Hash()
}

func (b *Block) RebuildMerkleRoot() error {
	txs := b.Transactions
	transactionHashes := []Uint256{}
	for _, tx := range txs {
		transactionHashes = append(transactionHashes, tx.Hash())
	}
	hash, err := crypto.ComputeRoot(transactionHashes)
	if err != nil {
		return errors.New("[Block] , RebuildMerkleRoot ComputeRoot failed.")
	}
	b.Header.MerkleRoot = hash
	return nil
}
