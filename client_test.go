package szzt

import (
	"fmt"
	"testing"

	"github.com/bigrocs/szzt/requests"
	uuid "github.com/satori/go.uuid"
)

func TestPlay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AccessId = "ISKJF459JD9FGU34"
	client.Config.AccessKey = "KLSKF3OD4RU3CDU9M3VOM39384VD35"
	// client.Config.ProductKey = "a1Z0BXAK0jS"
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "play"
	request.BizContent = map[string]interface{}{
		"requestId":  uuid.NewV4().String(),
		"productKey": "a1Z0BXAK0jS",
		"deviceId":   "ZS3191300125",
		"volume":     "100",
		"amount":     "125.6",
		"template":   "{微信}{收款}${元}",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	r, err := response.GetVerifySignDataMap()
	fmt.Println("TestPlay", r, err)
	t.Log(r, err, "|||")
}
