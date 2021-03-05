package iputil

import (
	"testing"
)

func Test_CheckIP(t *testing.T) {
	netIP := CheckValidIP("192.168.100.0")
	if netIP == nil {
		t.Fatal("wrong network IP")
	}
}
