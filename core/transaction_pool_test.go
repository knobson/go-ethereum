package core

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/core/state"
)

// State query interface
type stateQuery struct{ db common.Database }

func SQ() stateQuery {
	db, _ := ethdb.NewMemDatabase()
	return stateQuery{db: db}
}

func (self stateQuery) GetAccount(addr []byte) *state.StateObject {
	return state.NewStateObject(common.BytesToAddress(addr), self.db)
}

func transaction() *types.Transaction {
	return types.NewTransactionMessage(common.Address{}, common.Big0, common.Big0, common.Big0, nil)
}

func setup() (*TxPool, *ecdsa.PrivateKey) {
	var m event.TypeMux
	key, _ := crypto.GenerateKey()
	return NewTxPool(&m), key
}

func TestTxAdding(t *testing.T) {
	pool, key := setup()
	tx1 := transaction()
	tx1.SignECDSA(key)
	err := pool.Add(tx1)
	if err != nil {
		t.Error(err)
	}

	err = pool.Add(tx1)
	if err == nil {
		t.Error("added tx twice")
	}
}

func TestAddInvalidTx(t *testing.T) {
	pool, _ := setup()
	tx1 := transaction()
	err := pool.Add(tx1)
	if err == nil {
		t.Error("expected error")
	}
}

func TestRemoveSet(t *testing.T) {
	pool, _ := setup()
	tx1 := transaction()
	pool.addTx(tx1)
	pool.RemoveSet(types.Transactions{tx1})
	if pool.Size() > 0 {
		t.Error("expected pool size to be 0")
	}
}

func TestRemoveInvalid(t *testing.T) {
	pool, key := setup()
	tx1 := transaction()
	pool.addTx(tx1)
	pool.RemoveInvalid(SQ())
	if pool.Size() > 0 {
		t.Error("expected pool size to be 0")
	}

	tx1.SetNonce(1)
	tx1.SignECDSA(key)
	pool.addTx(tx1)
	pool.RemoveInvalid(SQ())
	if pool.Size() != 1 {
		t.Error("expected pool size to be 1, is", pool.Size())
	}
}

func TestInvalidSender(t *testing.T) {
	pool, _ := setup()
	err := pool.ValidateTransaction(new(types.Transaction))
	if err != ErrInvalidSender {
		t.Errorf("expected %v, got %v", ErrInvalidSender, err)
	}
}
