package common

import (
	"encoding/base64"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/bigrocs/szzt/config"
	"github.com/bigrocs/szzt/requests"
	"github.com/bigrocs/szzt/responses"
	"github.com/bigrocs/szzt/util"
)

// Common 公共封装
type Common struct {
	Config   *config.Config
	Requests *requests.CommonRequest
}

// Action 创建新的公共连接
func (c *Common) Action(response *responses.CommonResponse) (err error) {
	return c.Request(response)
}

// APIBaseURL 默认 API 网关
func (c *Common) APIBaseURL() string { // TODO(): 后期做容灾功能
	con := c.Config
	if con.Sandbox { // 沙盒模式
		return "https://proxy.szzt.com.cn/cs-api"
	}
	return "http://iot-proxy.szzt.com.cn:33321/api/v1"
}

// Request 执行请求
// AppCode           string `json:"app_code"`             //API编码
// AppId             string `json:"app_id"`               //应用ID
// UniqueNo          string `json:"unique_no"`            //私钥
// PrivateKey        string `json:"private_key"`          //私钥
// szztPublicKey string `json:"lin_shang_public_key"` //临商银行公钥
// MsgId             string `json:"msg_id"`               //消息通讯唯一编号，每次调用独立生成，APP级唯一
// Signature         string `json:"Signature"`            //签名值
// Timestamp         string `json:"timestamp"`            //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
// NotifyUrl         string `json:"notify_url"`           //工商银行服务器主动通知商户服务器里指定的页面http/https路径。
// BizContent        string `json:"biz_content"`          //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
// Sandbox           bool   `json:"sandbox"`              // 沙盒
func (c *Common) Request(response *responses.CommonResponse) (err error) {
	con := c.Config
	req := c.Requests
	// 构建配置参数
	params := map[string]interface{}{
		"accessId":      con.AccessId,
		"productKey":    con.ProductKey,
		"timestamp":     time.Now().Format("20060102150405"),
		"signatureVer":  "1",
		"signatureRand": strings.Replace(uuid.NewV4().String(), "-", "", -1),
		"action":        req.ApiName,
	}
	if v, ok := req.BizContent["template"]; ok {
		req.BizContent["template"] = base64.StdEncoding.EncodeToString([]byte(v.(string)))
	}
	for k, v := range req.BizContent {
		params[k] = v
	}
	sign := util.HmacSha1(util.EncodeSignParams(params), con.AccessKey) // 开发签名
	if err != nil {
		return err
	}
	params["signatureString"] = sign
	urlParam := util.FormatURLParam(params)
	res, err := util.HTTPGet(c.APIBaseURL() + "?" + urlParam)
	if err != nil {
		return err
	}
	response.SetHttpContent(res, "string")
	return
}
