package cache

import (
	"testing"
	"time"
)

func TestRedisCache(t *testing.T) {
	rc := NewRedisCache("127.0.0.1:6379")
	timeoutDuration := 10 * time.Second
	var err error
	if err = rc.Set("username", "uu", timeoutDuration); err != nil {
		t.Error("Set Error", err)
	}
	if !rc.IsExist("username") {
		t.Error("IsExist Error")
	}
	name := rc.Get("username").(string)
	if name != "uu" {
		t.Errorf("Get Error, want uu but %s", name)
	}
	if err = rc.Delete("username"); err != nil {
		t.Errorf("Delete Error, err=%v", err)
	}
}
