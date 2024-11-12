package data

import "github.com/laply/study-golang-coin/internal/block"

type Page struct {
	PageTitle string
	Blocks    []block.Block
}
