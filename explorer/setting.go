package explorer

import (
	"net/http"

	"github.com/labstack/gommon/random"
	"github.com/laply/study-golang-coin/explorer/data"
	"github.com/laply/study-golang-coin/internal/block"
)

func (router *Router) add(w http.ResponseWriter, r *http.Request) {
	newData := r.URL.Query().Get("data")
	if newData == "" {
		newData = random.String(20)
		block.GetBlockchain().AddBlock(newData)
	} else {
		block.GetBlockchain().AddBlock(newData)
	}
	templates.ExecuteTemplate(w, "add", data.Page{
		Blocks: block.GetBlockchain().GetBlocks(),
	})
}

func (router *Router) list(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home", data.Page{
		PageTitle: "COIN",
		Blocks:    block.GetBlockchain().GetBlocks(),
	})
}
