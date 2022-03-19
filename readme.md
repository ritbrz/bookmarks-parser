# Bookmarks Parser in Go

A Go parser for Chrome/Edge HTML Bookmarks Files

## Input

Chrome/Edge HTML Bookmarks File

## Output

`Bookmarks` struct

### Bookmarks

```go
type Bookmarks struct {
	Title    string
	URL      string
	IsDir    bool
	Children []*Bookmarks
}
```