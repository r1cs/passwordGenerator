package key_reflect

import (
	"fmt"
	"math/big"
)

type SymbolMap interface {
	Size() *big.Int
	IndexAt(index int64) string
}


func BigToSymbol(m SymbolMap,key *big.Int)string{
	base := m.Size()
	n := new(big.Int).Set(key)
	indexs := []int64{}
	for n.Cmp(base)>0{
		mod := new(big.Int).Mod(n,base)
		indexs=append(indexs,mod.Int64())
		n.Div(n,base)
	}
	indexs=append(indexs,n.Int64())
	 ret := ""
	 for i := len(indexs)-1;i>=0;i--{
	 	ret+=m.IndexAt(indexs[i])
	 }
	 return ret
}

type HexKey struct {}

func(h *HexKey)Size()*big.Int{
	return big.NewInt(16)
}

func(h *HexKey)IndexAt(Index int64)string{
	return fmt.Sprintf("%x",Index)
}