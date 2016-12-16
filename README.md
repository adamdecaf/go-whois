# go-whois

> A simple whois client

![](https://travis-ci.org/adamdecaf/go-whois.svg?branch=master)

## Usage

```go
resp, err := WhoisQuery("google.com")

if err != nil {
	t.Errorf("error when finding whois server = %s", err)
}

fmt.Printf("resp = %s\n", resp)
```

## Docs

- [godoc](https://godoc.org/github.com/adamdecaf/go-whois)

## Dependencies

We use [zonedb/zonedb](https://github.com/zonedb/zonedb) to hold the record of whois servers.
