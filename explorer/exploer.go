package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/laply/study-golang-coin/internal/block"
)

type Router struct {
	port string
}

var templates *template.Template

func InitRouter(port string) *Router {
	return &Router{port: port}
}

func (r *Router) Run(port, templateDir string) {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	fmt.Printf("Listening on http://localhost%s\n", r.port)

	http.HandleFunc("/", r.list)

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, block.GetBlockchain().ListBlocks())
	})

	http.HandleFunc("/add", r.add)
	log.Fatal(http.ListenAndServe(r.port, nil))
}
