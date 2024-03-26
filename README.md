## 目标


### 原理

### 快速使用

```c
git clone  https://github.com/284851828/alilite_go.git
go mod tidy
go run demo.go
```


"# alilite_go" 


```c
package main

import (
	"log"
	"time"

	c "alilite/client" // Replace with your actual package path
)

func main() {
	const URL_PATH = "/gate/liteContract/create"
	client := c.Client{
		AppID: "999999", // Replace with your actual App ID
	}
	request := c.Request{
		Timestamp: time.Now().Unix(),
		Content:   c.Content{ExtTradeNo: "orderid_" + time.Now().String(), RedirectURL: "https://to_your_successful_webpage/"},
	}
	//收款方
	request.Content.Company.ID = "compid_xxxx"
	request.Content.Company.Name = "某某有限公司"
	//付款方
	request.Content.Customer.ExtID = "userid_xxxx"
	request.Content.Customer.Name = "张某"
	request.Content.Customer.IDCard = "33100000000" //根据此号关联付款用户
	request.Content.Customer.Addr = "地址"
	//产品描述
	request.Content.Product.ExtID = "id_xxx"
	request.Content.Product.Name = "产品名称"
	request.Content.Product.Price = "1.00"
	request.Content.Product.Content = "描述"
	//扣款相关
	request.Content.Installment.Limit = 20.00
	request.Content.Installment.Num = 2
	request.Content.Installment.First = 0.01
	request.Content.Installment.Type = "SDI"

	r, err := client.Post(URL_PATH, &request)
	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
		return
	}

	log.Println("POST request successful.")
	log.Println("Resp:	", r)
}


```


### 请求与返回

```c
curl -X POST \
     -H "Content-Type: application/json" \
     -H "X-Gateway-AppId: 999999" \
     -d '{"timestamp":1711347382,"content":{"extTradeNo":"1711347382","redirectUrl":"https://geekland.cc","company":{"id":"cnogda7i2dkqvf","name":"浙江某科技有限公司"},"customer":{"extId":"user_002","name":"张某","addr":"杭州市西湖区","phone":"13958040000","idCard":"3310021981000000"},"product":{"extId":"product_002","name":"恰恰香瓜子","price":"1.00","Content":"好吃"},"installment":{"limit":1.00,"first":0.01,"num":2,"type":"SDI"}}}' \
     https://open.xiadandt.com/gate/liteContract/create


{
  "code": 0,
  "content": {
    "contractId": "co0hdgni2dkrn7or4m8g",
    "customer": {
      "extId": "user_002",
      "name": "张某",
      "addr": "杭州市西湖区",
      "phone": "13958040000",
      "idCard": "33100219810412251X"
    },
    "bill": [
      {
        "index": 1,
        "days": "2024-03-25T08:00:00+08:00",
        "amount": 0.5
      },
      {
        "index": 2,
        "days": "2024-04-25T08:00:00+08:00",
        "amount": 0.5
      }
    ],
    "signUrl": "https://u.alipay.cn/_eLriCTVod5djaQX9hEFxd",
    "status": 12
  },
  "msg": "创建成功"
}

```

申请 APPID
 
<img src="https://raw.githubusercontent.com/284851828/alilite_nodejs/main/github_8888.png" width = 250 height = 250>

联系客服

<img src="https://raw.githubusercontent.com/284851828/alilite_nodejs/main/wx.jpg" width = 300 height = 300>

 
 

 




