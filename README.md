# go-whois

> A simple whois client

![](https://api.travis-ci.org/adamdecaf/go-whois.svg)

## Usage

```go
resp, err := WhoisQuery("google.com")

if err != nil {
	t.Errorf("error when finding whois server = %s", err)
}

fmt.Printf("resp = %s\n", resp)
```

## Dependencies

We use [zonedb/zonedb](https://github.com/zonedb/zonedb) to hold the record of whois servers.

## TODO

- [ ] Automatic parsing of WHOIS responses
