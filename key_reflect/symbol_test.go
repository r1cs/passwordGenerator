package key_reflect

import (
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestSymbol(t *testing.T){
	a := crypto.Keccak256(nil)
	l := new(big.Int).SetBytes(a)
	k := BigToSymbol(UrlKeyMap,l)
	t.Log(k)

	kk := BigToSymbol(&HexKey{},l)
	t.Log(kk)
	if kk!=l.Text(16){
		t.Fatal(1)
	}
}
