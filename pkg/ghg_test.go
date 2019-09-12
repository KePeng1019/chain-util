package pkg

import (
	"encoding/json"
	"github.com/bwmarrin/snowflake"
	"log"
	"testing"
)

func TestCoin(t *testing.T) {
	coin := Coin{}
	data, err := json.Marshal(coin)
	if nil != err {
		t.Fatal(err)
	}
	log.Println(string(data))
	snowFlakeNode, _ := snowflake.NewNode(1)
	log.Println(snowFlakeNode.Generate().Int64())
}
