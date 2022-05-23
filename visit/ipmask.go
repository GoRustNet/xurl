package visit

import (
	"strings"

	"github.com/GoRustNet/xurl/str"
)

const invalidIP = "INVALID_IP"

func IpMask(ip string) string {
	ss, isIpv6 := split(ip)
	return mask(ss, isIpv6)
}

func mask(ss []string, isIpv6 bool) string {
	if ss == nil {
		return invalidIP
	}
	ssLen := len(ss)
	if ssLen < 3 {
		return invalidIP
	}
	newSS := []string{}
	newSS = append(newSS, ss[:ssLen-2]...)
	newSS = append(newSS, []string{"*", "*"}...)
	s := "."
	if isIpv6 {
		s = ":"
	}
	return strings.Join(newSS, s)
}
func split(ip string) ([]string, bool) {
	if str.IsEmpty(ip) {
		return nil, false
	}
	ipv4s := strings.Split(ip, ".")
	if len(ipv4s) > 1 {
		return ipv4s, false
	}
	ipv6s := strings.Split(ip, ":")
	if len(ipv6s) > 1 {
		return ipv6s, true
	}
	return nil, false
}
