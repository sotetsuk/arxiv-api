package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	version := "arxiv-api 0.0.1"
	usage := `Call arXiv api. http://arxiv.org/help/api/user-manual

Usage:
  arxiv-api <query> [--id_list=<id_list>] [--start=<s>] [--max_results=<n>]
  arxiv-api -h | --help
  arxiv-api --version

Options:
  --id_list=<id_list>    Id list (comma-delimited)
  --start=<s>            Index of start [default: 0]
  --max_results=<n>      Number of results [default: 10]
  -h --help              Show this screen.
  --version              Show version.`

	query, id_list, start, max_results := parse(usage, version)
	resp := call(query, id_list, start, max_results)
	defer resp.Body.Close()
	execute(resp)
}

func parse(usage, version string) (query, id_list, start, max_results string) {
	arguments, _ := docopt.Parse(usage, os.Args[1:], true, version, false)

	var casted = make(map[string]string)
	for key, val := range arguments {
		casted[key], _ = val.(string)
	}

	return casted["<query>"], casted["--id_list"], casted["--start"], casted["--max_results"]
}

func call(query, id_list, start, max_results string) *http.Response {
	values := url.Values{}
	values.Add("search_query", query)
	values.Add("id_list", id_list)
	values.Add("start", start)
	values.Add("max_results", max_results)

	resp, err := http.Get("http://export.arxiv.org/api/query" + "?" + values.Encode())

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return resp
}

func execute(response *http.Response) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
