package internal

import "testing"

func TestStreamWrapper(t *testing.T) {
	streamWrapperConfig := StreamWrapperConfig{
		apiKey:            "ajpcpacmp",
		apiSecretKey:      "nosdnvosdnvonsdovnosd",
		accessToken:       "snidviosdnvionsdfviondfiovndfiovnosdfnvoi",
		accessTokenSecret: "vcjsopvmpvmpsdmvpsdmvpdspovmopsdmvopmsdopmvpmsdpvm",
	}
	filter := Filter{
		track: []string{"ETH,BTC"},
	}
	Start(streamWrapperConfig, filter)
}
