package main

import (
	cli "github.com/urfave/cli/v2"
)

var UrlFlag = &cli.StringFlag{
	Usage: "net url to use the password",
	Name: "url",
	Aliases: []string{"u"},
	Required: true,
}


var NameFlag = &cli.StringFlag{
	Usage: "account name",
	Name: "name",
	Aliases: []string{"n"},
	Required: true,
}

var SecretFlag = &cli.StringFlag{
	Usage: "secret key to generate password",
	Name: "key",
	Aliases: []string{"s"},
	Required: true,
}



