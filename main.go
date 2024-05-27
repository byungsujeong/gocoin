package main

import (
	"github.com/byungsujeong/gocoin/cli"
	"github.com/byungsujeong/gocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
