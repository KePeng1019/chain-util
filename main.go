package main

import (
	"github.com/bwmarrin/snowflake"
	"github.com/Kepeng1019/chain-util/pkg"
	"log"
)

func main() {
	node, _ := snowflake.NewNode(1)
	log.Println(node.Generate().Int64())
	log.Println(pkg.Coin{})
}
