package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

func handler(w http.ResponseWriter, req *http.Request) {
	graph := BuildMermaidGraph();

	w.Header().Set("Access-Control-Allow-Origin", "*");

	fmt.Fprintf(w, graph);
}

func main() {
	db := DbConn()
	conf := GetConfig();

	http.HandleFunc("/", handler)

	log.Printf("Port: %d \n", conf.Network.Port);

	log.Fatal(http.ListenAndServe(":" + strconv.FormatInt(conf.Network.Port, 10), nil))
	defer db.Close();
}
