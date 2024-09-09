package app

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/laply/study-golang-coin/internal/block"
)

type Router struct {
	port string
}

func InitRouter(port string) *Router {
	return &Router{port: port}
}

func (r *Router) Run() error {
	fmt.Printf("Listening on http://localhost%s\n", r.port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/pages/index.gohtml")
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		tmpl.Execute(w, struct {
			Blocks []block.Block
		}{Blocks: block.GetBlockchain().GetBlocks()})

	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, block.GetBlockchain().ListBlocks())
	})
	http.HandleFunc("/add-block", func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("data")

		if data == "" {
			fmt.Fprint(w, "Data is empty")
			return

		} else {
			block.GetBlockchain().AddBlock(data)

			tmpl, err := template.ParseFiles("templates/pages/add.gohtml")
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			tmpl.Execute(w, data)
		}
	})

	return http.ListenAndServe(r.port, nil)
}
