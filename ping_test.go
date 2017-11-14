package minepong

import (
	"fmt"
	"sync"
	"testing"
)

var testServers = map[string][]string{
	"desteria": {"Play.NirvanaMC.com:25565", "Play.NirvanaMC.com:25565"},
	"gotpvp":   {"play.gotpvp.com:25565", "283hd134d142d7h2.ddns.net.:25565"},
	"SRV-PVP":  {"ping.minecraft.syfaro.net", "play.gotpvp.com.:25565"},
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
