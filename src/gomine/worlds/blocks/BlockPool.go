package blocks

import (
	"gomine/interfaces"
)

var blocks = map[int]func(byte) interfaces.IBlock{}

func InitBlockPool() {
	//RegisterBlock(Air, func(data byte) interfaces.IBlock { return transparent.NewAir(data) })
	//RegisterBlock(Stone, func(data byte) interfaces.IBlock { return full.NewStone(data) })
}

/**
 * Registers a new block with a function that creates it.
 */
func RegisterBlock(id int, block func(byte) interfaces.IBlock) {
	blocks[id] = block
}

/**
 * Returns a new block with the given ID.
 */
func GetBlock(id int, data byte) interfaces.IBlock {
	return blocks[id](data)
}