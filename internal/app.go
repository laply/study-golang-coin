package app

import "github.com/laply/study-golang-coin/internal/block"

type App struct {
	// Config is the configuration for the app.

}

func New() *App {
	return &App{}
}

func (a *App) Run() error {

	bc := block.GetBlockchain()
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")
	bc.ListBlocks()
	return nil
}
