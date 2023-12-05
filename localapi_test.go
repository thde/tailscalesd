package tailscalesd

import (
	"net"
	"reflect"
	"testing"
)

func TestTranslatePeerToDevice(t *testing.T) {
	want := Device{
		Addresses: []string{
			"100.2.3.4",
			"fd7a::1234",
		},
		API:        "localhost",
		Authorized: true,
		Hostname:   "somethingclever",
		ID:         "id",
		OS:         "beos",
		Tags: []string{
			"tag:foo",
			"tag:bar",
		},
	}
	var got Device
	translatePeerToDevice(&interestingPeerStatusSubset{
		ID:       "id",
		HostName: "somethingclever",
		DNSName:  "this is currently ignored",
		OS:       "beos",
		TailscaleIPs: []net.IP{
			net.ParseIP("100.2.3.4"),
			net.ParseIP("fd7a::1234"),
		},
		Tags: []string{
			"tag:foo",
			"tag:bar",
		},
	}, &got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("translatePeerToDevice() = %v, want %v", got, want)
	}
}
