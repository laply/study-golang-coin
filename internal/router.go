package app

import (
	"fmt"
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

	bc := block.GetBlockchain()

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, bc.ListBlocks())
	})
	http.HandleFunc("/add-block", func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("data")

		if data == "" {
			fmt.Fprint(w, "Data is empty")
			return

		} else {
			bc.AddBlock(data)
			fmt.Fprint(w, "Add Block [", data, "]")
		}
	})

	return http.ListenAndServe(r.port, nil)
}
