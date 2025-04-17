package gewechat_client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// GewechatClient 主客户端结构体
type GewechatClient struct {
	baseURL    string
	token      string
	httpClient *http.Client

	LoginApi    *LoginApi
	ContactApi  *ContactApi
	MessageApi  *MessageApi
	GroupApi    *GroupApi
	DownloadApi *DownloadApi
	FavorApi    *FavorApi
	LabelApi    *LabelApi
	PersonalApi *PersonalApi
}

// NewGewechatClient 初始化客户端
func NewGewechatClient(baseURL, token string) *GewechatClient {
	c := &GewechatClient{
		baseURL:    baseURL,
		token:      token,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
	// 初始化各子模块
	c.LoginApi = NewLoginApi(c)
	c.ContactApi = NewContactApi(c)
	c.MessageApi = NewMessageApi(c)
	c.GroupApi = NewGroupApi(c)
	c.DownloadApi = NewDownloadApi(c)
	c.FavorApi = NewFavorApi(c)
	c.LabelApi = NewLabelApi(c)
	c.PersonalApi = NewPersonalApi(c)
	return c
}

// PostJson 基础请求方法
func (c *GewechatClient) PostJson(path string, data interface{}) (map[string]interface{}, error) {
	var urlStr = c.baseURL + path
	return c.HttpRequest(urlStr, "POST", nil, nil, data)
}

func (c *GewechatClient) HttpRequest(urlStr string, method string, headers map[string]string, params map[string]string, data any) (map[string]interface{}, error) {
	// 创建URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	query := u.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()

	// 将数据编码为JSON
	buf := new(bytes.Buffer)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	// 创建请求
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("X-GEWE-TOKEN", c.token)

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
