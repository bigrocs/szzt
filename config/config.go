package config

// 服务器URL ： https://proxy.szzt.com.cn/cs-api
// accessId : ISKJF459JD9FGU34
// accessKey : KLSKF3OD4RU3CDU9M3VOM39384VD35
// productKey : a1Z0BXAK0jS

type Config struct {
	AccessId   string `json:"access_id"`   // 开发者ID
	AccessKey  string `json:"access_key"`  // 开发者密钥
	ProductKey string `json:"product_key"` // 产品KEY
	Sandbox    bool   `json:"sandbox"`     // 沙盒
}
