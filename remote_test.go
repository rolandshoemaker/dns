package dns

import (
	"net"
	"testing"
)

const LinodeAddr = "176.58.119.54:53"

func TestClientRemote(t *testing.T) {
	m := new(Msg)
	m.SetQuestion("go.dns.miek.nl.", TypeTXT)

	c := new(Client)
	c.Dialer = &net.Dialer{}
	r, _, err := c.Exchange(m, LinodeAddr)
	if err != nil {
		t.Errorf("failed to exchange: %v", err)
	}
	if r != nil && r.Rcode != RcodeSuccess {
		t.Errorf("failed to get an valid answer\n%v", r)
	}
}
