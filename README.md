# go-fedifinger

Package **fedifinger** provides tools for working with the **WebFinger** protocol as it is used by the **Fediverse**, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fedifinger

[![GoDoc](https://godoc.org/github.com/reiver/go-fedifinger?status.svg)](https://godoc.org/github.com/reiver/go-fedifinger)

## Examples

To resolve a **Fediverse-ID** to an HTTPS URL, do something similar to:

```golang
import "github.com/reiver/go-fedifinger"

// ...

url, err := fedifinger.Resolve("@reiver@mastodon.social")
```

To get the activity-JSON for a **Fediverse-ID**, do something similar to:

```golang
import "github.com/reiver/go-fedifinger"

// ...

bytes, err := fedifinger.Get("@reiver@mastodon.social")
```

## Import

To import package **fedifinger** use `import` code like the follownig:
```
import "github.com/reiver/go-fedifinger"
```

## Installation

To install package **fedifinger** do the following:
```
GOPROXY=direct go get github.com/reiver/go-fedifinger
```

## Author

Package **fedifinger** was written by [Charles Iliya Krempeaux](http://reiver.link)
