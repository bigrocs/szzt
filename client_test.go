package szzt

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigrocs/szzt/requests"
	uuid "github.com/satori/go.uuid"
)

func TestPlay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AccessId = os.Getenv("SZZT_AccessId")
	client.Config.AccessKey = os.Getenv("SZZT_AccessKey")
	client.Config.ProductKey = os.Getenv("SZZT_ProductKey")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "play"
	request.BizContent = map[string]interface{}{
		"requestId": uuid.NewV4().String(),
		"deviceId":  "ZS1213603374",
		"volume":    "100",
		"amount":    "0.06",
		"template":  "{测试2}${元}",
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
