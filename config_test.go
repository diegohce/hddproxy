package main

import (
	"os"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {

	os.Setenv("HDDPROXY_TCP_BIND", ":6060")
	os.Setenv("HDDPROXY_DIRS", "/home/diego/tmp/stuff:/home/diego/tmp/stuff2")
	os.Setenv("HDDPROXY_SHORTPOLL", "10")
	os.Setenv("HDDPROXY_LONGPOLL", "20")

	c := configFromEnv()

	if c.tcpBind != ":6060" {
		t.Errorf("Bad tcpBind: %v. Want: :6060", c.tcpBind)
	}

	if c.dirs[0] != "/home/diego/tmp/stuff" || c.dirs[1] != "/home/diego/tmp/stuff2" {
		t.Error("Failed on dirs array")
	}

	if c.longPoll != time.Duration(20)*time.Second {
		t.Errorf("Bad longPoll: %v. Expected 20\n", c.longPoll)
	}

	if c.shortPoll != time.Duration(10)*time.Second {
		t.Errorf("Bad shortPoll: %v. Expected 10\n", c.shortPoll)
	}
}
