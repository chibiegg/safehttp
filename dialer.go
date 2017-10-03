package safehttp

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type FilterDialer struct {
	Dialer *net.Dialer
	Filter Filter
}

func (d *FilterDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	s := strings.Split(addr, ":")
	hostname := s[0]
	port, _ := strconv.Atoi(s[1])

	ipAddrList, err := net.LookupIP(hostname)
	if err != nil {
		return nil, err
	}

	var ip net.IP
	err = nil

	for _, tip := range ipAddrList {
		err = d.Filter.Check(tip, port)
		if err == nil {
			ip = tip
			break
		}
	}

	if err != nil {
		return nil, err
	}

	addr = fmt.Sprintf("%s:%d", ip.String(), port)
	return d.Dialer.DialContext(ctx, network, addr)
}
