# arxiv-api
Call arXiv api. http://arxiv.org/help/api/user-manual

# Install & Try it

Assume that `$GOPATH` is set and `$PATH` includes `$GOPATH/bin`.

```
$ go get github.com/sotetsuk/arxiv-api
$ arxiv-api all:electron
```

# Usage

```
$ arxiv-api --help
Call arXiv api. http://arxiv.org/help/api/user-manual

Usage:
  arxiv-api <query> [--id_list=<id_list>] [--start=<s>] [--max_results=<n>]
  arxiv-api -h | --help
  arxiv-api --version

Options:
  --id_list=<id_list>    Id list (comma-delimited)
  --start=<s>            Index of start [default: 0]
  --max_results=<n>      Number of results [default: 10]
  -h --help              Show this screen.
  --version              Show version.
```
