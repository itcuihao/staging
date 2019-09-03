package confs

import "testing"

func TestRconf(t *testing.T) {
	NewConfig("conf.json")
	t.Logf("conf:%+v", conf)
	t.Logf("conf.mysql:%+v", conf.MySQL)
	select {}
}
