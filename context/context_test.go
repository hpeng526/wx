package context

import (
	"fmt"
	"github.com/hpeng526/wx/cache"
	"testing"
)

func TestContext_GetAccessToken(t *testing.T) {
	var ctx = &Context{AppID: "wx9482dd922f2b6ba4", AppSecret: "8bce95b05895f73ed6a7cb912c6edb51", Cache: cache.NewRedisCache("127.0.0.1:6379")}
	token, err := ctx.GetAccessToken()
	if err != nil {
		fmt.Printf("err :%s", err)
	}
	fmt.Printf("token is %s", token)

}
