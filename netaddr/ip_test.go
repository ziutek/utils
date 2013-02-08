package netaddr

import (
	"net"
	"testing"
)

type example struct {
	ip     string
	offset int
	exp    string
}

const ipv6max = "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"

var test = []example{
	{"0.0.0.0", 1, "0.0.0.1"},
	{"0.0.0.0", -1, "255.255.255.255"},
	{"255.255.255.255", 1, "0.0.0.0"},
	{"255.255.255.255", -1, "255.255.255.254"},
	{"0.0.0.0", 256, "0.0.1.0"},
	{"0.0.0.0", -256, "255.255.255.0"},
	{"255.255.255.255", 256, "0.0.0.255"},
	{"255.255.255.255", -256, "255.255.254.255"},

	{"::", 1, "::1"},
	{"::", -1, ipv6max},
	{ipv6max, 1, "::"},
	{ipv6max, -1, "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fffe"},
	{"::ffff:ffff:ffff:0000", 0x10000, "::1:0000:0000:0000:0000"},
	{"::ffff:ffff:ffff:0000", -0x10000, "::ffff:ffff:fffe:0000"},
}

func TestIPAdd(t *testing.T) {
	for _, e := range test {
		a := net.ParseIP(e.ip)
		b := IPAdd(a, e.offset)
		if !b.Equal(net.ParseIP(e.exp)) {
			t.Errorf("IPAdd(%s, %d)=%s != %s", e.ip, e.offset, b, e.exp)
		}
	}
}
