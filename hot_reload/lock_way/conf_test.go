package lock_way

import "testing"

func TestRconf(t *testing.T) {
	c := NewConfig("conf.json")
	t.Log(c.MySQL.Host)
	select {}
}
