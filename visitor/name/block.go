package name

import "fmt"

var blockCount = 0

func BlockName() string {
	name := fmt.Sprintf("block-%d", blockCount)
	blockCount++
	return name
}