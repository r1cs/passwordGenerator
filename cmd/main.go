package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/r1cs/passwordGenerator/key_reflect"
	cli "github.com/urfave/cli/v2"
	"log"
	"math/big"
	"os"
)

func main() {
	app := cli.App{
		Name:   "passwordGenerator",
		Usage:  "one key generate various password",
		Action: run,
		Flags: []cli.Flag{
			UrlFlag,
			NameFlag,
			SecretFlag,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func run(ctx *cli.Context) error {

	url, name, key := ctx.String(UrlFlag.Name), ctx.String(NameFlag.Name), ctx.String(SecretFlag.Name)
	private := crypto.Keccak256Hash([]byte(url + name + key))

	p := MixHash(private)
	k := key_reflect.BigToSymbol(key_reflect.UrlKeyMap, new(big.Int).SetBytes(p.Bytes()))

	fmt.Println("generated key:  ", k)
	return nil
}

func MixHash(k common.Hash)common.Hash{
	return crypto.Keccak256Hash(crypto.Keccak256(append(k.Bytes(),[]byte("r1cs")...)))
}
