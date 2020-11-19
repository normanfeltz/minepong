# minepong

Golang Minecraft server pinger.

# Usage

The `minepong.Ping` method takes care of resolving SRV records, connecting to
the server, and retrieving data.

```go
pong, err := minepong.Ping(host)
if err != nil {
    fmt.Println(err)
    return
}
```

A `minepong.Pong` is returned with information from the ping.

```go
type Pong struct {
    Version struct {
        Name     string
        Protocol int
    } `json:"version"`
    Players struct {
        Max    int `json:"max"`
        Online int `json:"online"`
        Sample []map[string]string
    } `json:"players"`
    Description  interface{} `json:"description"`
    FavIcon      string      `json:"favicon"`
    ResolvedHost string      `json:"resolved_host"`
}
```
