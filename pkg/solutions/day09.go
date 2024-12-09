package solutions

import (
	"strconv"
)

type Day09DiskBlocks struct {
	Id         int
	FileBlocks int
	EmptyBlock int
}
type Day09Part01 struct {
	DiskBlocks []Day09DiskBlocks
}
type Day09Part02 struct {
	DiskBlocks []Day09DiskBlocks
}

func (data Day09Part01) blocksValues() []int64 {
	blocksValues := make([]int64, 0)
	for _, diskBlock := range data.DiskBlocks {
		for i := 0; i < diskBlock.FileBlocks; i++ {
			blocksValues = append(blocksValues, int64(diskBlock.Id))
		}
		for i := 0; i < diskBlock.EmptyBlock; i++ {
			blocksValues = append(blocksValues, -1)
		}
	}
	return blocksValues
}

type day09Blocks struct {
	blockValue int64
	blocks     int64
	visited    bool
	next       *day09Blocks
	previous   *day09Blocks
}

func (blocks day09Blocks) isFileBlock() bool {
	return blocks.blockValue != -1
}

func (data Day09Part02) blocksValues() (day09Blocks, day09Blocks) {
	firstBlock := day09Blocks{}
	lastBlock := day09Blocks{}
	var previousBlock *day09Blocks = nil
	for index, diskBlock := range data.DiskBlocks {
		fileBlock := day09Blocks{blockValue: int64(diskBlock.Id), blocks: int64(diskBlock.FileBlocks)}
		emptyBlock := day09Blocks{blockValue: -1, blocks: int64(diskBlock.EmptyBlock)}
		emptyBlock.previous = &fileBlock
		fileBlock.next = &emptyBlock
		fileBlock.previous = previousBlock
		if previousBlock != nil {
			previousBlock.next = &fileBlock
		}
		previousBlock = &emptyBlock
		if index == 0 {
			firstBlock = fileBlock
		} else if index == len(data.DiskBlocks)-1 {
			lastBlock = emptyBlock
		}
	}
	return firstBlock, lastBlock
}
func (data *Day09Part01) Exec() (string, error) {
	diskSpace := data.blocksValues()
	right := len(diskSpace) - 1
	for left, diskBlock := range diskSpace {
		if diskBlock == -1 {
			for ; right > left; right-- {
				if diskSpace[right] != -1 {
					diskSpace[left] = diskSpace[right]
					diskSpace[right] = -1
					break
				}
			}
		}
		if left >= right {
			break
		}
	}
	var result int64 = 0
	var id int64 = 0
	for index := 0; index < len(diskSpace); index++ {
		if diskSpace[index] == -1 {
			break
		}
		result += diskSpace[index] * id
		id++
	}
	return strconv.FormatInt(result, 10), nil
}

func (data *Day09Part02) Exec() (string, error) {
	firstBlock, lastBlock := data.blocksValues()
	currentBlock := &lastBlock
	for {
		if currentBlock == nil {
			break
		}
		if currentBlock.isFileBlock() && !currentBlock.visited {
			fulfillBlock := findFirstBlockToFulfill(currentBlock, &firstBlock)
			if fulfillBlock == nil {
				currentBlock = currentBlock.previous
			} else {
				freeBlocks := fulfillBlock.blocks - currentBlock.blocks
				fulfillBlock.blockValue = currentBlock.blockValue
				fulfillBlock.blocks = currentBlock.blocks
				fulfillBlock.visited = true
				if freeBlocks > 0 {
					nextBlock := fulfillBlock.next
					newEmptyBlock := day09Blocks{blockValue: -1, blocks: freeBlocks}
					newEmptyBlock.previous = fulfillBlock
					newEmptyBlock.next = nextBlock
					fulfillBlock.next = &newEmptyBlock
					nextBlock.previous = &newEmptyBlock
				}
				currentBlock.blockValue = -1
			}
		} else {
			currentBlock = currentBlock.previous
		}
	}
	var result int64 = 0
	var index int = 0
	block := &firstBlock
	for {
		if block == nil {
			break
		}
		if block.isFileBlock() {
			for i := 0; i < int(block.blocks); i++ {
				idx := index + i
				result += int64(idx) * block.blockValue
			}
			index += int(block.blocks)
		} else {
			index += int(block.blocks)
		}
		block = block.next
	}
	return strconv.FormatInt(result, 10), nil
}

func findFirstBlockToFulfill(currentBlock *day09Blocks, firstBlock *day09Blocks) *day09Blocks {
	searchingBlock := firstBlock
	for {
		if searchingBlock == nil {
			return nil
		}
		if searchingBlock.blockValue == currentBlock.blockValue { //means that we are at position of searching block
			return nil
		}
		if !searchingBlock.isFileBlock() && searchingBlock.blocks >= currentBlock.blocks {
			return searchingBlock
		}
		searchingBlock = searchingBlock.next
	}
}
