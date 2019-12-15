package minepong

import (
	"fmt"
	"sync"
	"testing"
)

var testServers = map[string][]string{
	"desteria":  {"Play.NirvanaMC.com:25565", "nmctemporary.ddns.net.:25564"},
	"skybounds": {"play.skybounds.com:25565", "ns105195.ip-147-135-69.us.:25565"},
}

func TestPing(t *testing.T) {
	wg := &sync.WaitGroup{}

	for name, host := range testServers {
		wg.Add(1)

		go func(name string, data []string) {
			fmt.Printf("Checking %s: %s\n", name, data[0])
			defer wg.Done()

			pong, err := Ping(data[0])
			if err != nil {
				fmt.Println(err)
				t.Fail()
				return
			}

			if pong.ResolvedHost != data[1] {
				t.Errorf("SRV lookup did not complete, got %s, expected %s\n", pong.ResolvedHost, data[1])
				t.Fail()
			}

			fmt.Printf("Got %s: %d/%d\n", name, pong.Players.Online, pong.Players.Max)
		}(name, host)
	}

	wg.Wait()
}
