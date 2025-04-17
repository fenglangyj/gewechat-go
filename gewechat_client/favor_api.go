package gewechat_client

// FavorApi 结构体用于封装收藏夹相关的 API 操作
type FavorApi struct {
	client *GewechatClient
}

// NewFavorApi 创建一个新的 FavorApi 实例
func NewFavorApi(c *GewechatClient) *FavorApi {
	return &FavorApi{client: c}
}

// Sync 同步收藏夹
func (f *FavorApi) Sync(appID, syncKey string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId":   appID,
		"syncKey": syncKey,
	}
	return f.client.PostJson("/favor/sync", param)
}

// GetContent 获取收藏夹内容
func (f *FavorApi) GetContent(appID, favID string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
		"favId": favID,
	}
	return f.client.PostJson("/favor/getContent", param)
}

// Delete 删除收藏夹
func (f *FavorApi) Delete(appID, favID string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
		"favId": favID,
	}
	return f.client.PostJson("/favor/delete", param)
}
