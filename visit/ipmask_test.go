package visit

import "testing"

func TestIpMask(t *testing.T) {
	var ips = []string{
		"127.0.0.1",
		"::1",
		"8.8.8.8",
		"684D:1111:222:3333:4444:5555:6:77",
		"2001:4860:4860::6464",
		"2001:4860:4860:0:0:0:0:64",
	}
	for _, ip := range ips {
		t.Log(IpMask(ip))
	}
}
