package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("Failed to initialize snowflake node: %v", err)
	}
}

func GenerateSnowflakeID() snowflake.ID {
	return node.Generate()
}
