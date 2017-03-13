package template

import (
	"encoding/json"
	"fmt"
	"github.com/hpeng526/wx/cache"
	"github.com/hpeng526/wx/context"
	"testing"
)

type inner struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type tmpData struct {
	First  inner `json:"first"`
	Send   inner `json:"send"`
	Text   inner `json:"text"`
	Time   inner `json:"time"`
	Remark inner `json:"remark"`
}

func TestTmpData(t *testing.T) {
	tInner := &inner{Color: "#173177"}
	tmp := &tmpData{*tInner, *tInner, *tInner, *tInner, *tInner}
	tmp.First.Value = "first"
	tmp.Send.Value = "send"
	tmp.Text.Value = "Text"
	tmp.Time.Value = "Time"
	tmp.Remark.Value = "Remark"
	jsonData, err := json.Marshal(tmp)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	fmt.Printf("tmp is : %v \n", string(jsonData))

	wxID := "yourid"
	tpID := "yourtempid"
	var tempMsg = &TemplateMessage{ToUser: wxID, TemplateID: tpID, TopColor: "#ff0000", URL: "http://www.bilibili.com", JSONData: jsonData}

	jsonData, err = json.Marshal(tempMsg)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	fmt.Printf("tempMsg is : %v \n", string(jsonData))

	var ctx = &context.Context{AppID: "wx9482dd922f2b6ba4", AppSecret: "8bce95b05895f73ed6a7cb912c6edb51", Cache: cache.NewRedisCache("127.0.0.1:6379")}
	token, err := ctx.GetAccessToken()
	if err != nil {
		fmt.Printf("err :%s", err)
	}
	fmt.Printf("token is %s", token)
	tempMsg.SendTemplate(token)
}
