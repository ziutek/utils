package netaddr

import (
	"fmt"
	"strconv"
	"strings"
)

type MAC uint64

// ParseMAC returns 0 MAC if error
func ParseMAC(s string) MAC {
	s = strings.Map(
		func(r rune) rune {
			switch r {
			case '-', '.', ':':
				return -1
			}
			return r
		},
		s,
	)
	if len(s) != 12 {
		return 0
	}
	m, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0
	}
	return MAC(m)
}

func (m MAC) String() string {
	return fmt.Sprintf(
		"%02x-%02x-%02x-%02x-%02x-%02x",
		byte(m>>40), byte(m>>32), byte(m>>24), byte(m>>16), byte(m>>8), byte(m),
	)
}

func (m MAC) ColonStr() string {
	return fmt.Sprintf(
		"%02x:%02x:%02x:%02x:%02x:%02x",
		byte(m>>40), byte(m>>32), byte(m>>24), byte(m>>16), byte(m>>8), byte(m),
	)
}
