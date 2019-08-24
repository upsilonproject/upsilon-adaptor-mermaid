package main;

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	uam "github.com/upsilonproject/upsilon-adaptor-mermaid/pkg/upsilonproject/upsilon-adaptor-mermaid"
)

func handler(w http.ResponseWriter, req *http.Request) {
	graph := uam.BuildMermaidGraph();

	w.Header().Set("Access-Control-Allow-Origin", "*");

	fmt.Fprintf(w, graph);
}

func main() {
	db := uam.DbConn()
	conf := uam.GetConfig();

	http.HandleFunc("/", handler)

	log.Printf("Port: %d \n", conf.Network.Port);

	log.Fatal(http.ListenAndServe(":" + strconv.FormatInt(conf.Network.Port, 10), nil))
	defer db.Close();
}
