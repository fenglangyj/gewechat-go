package gewechat_client

// LabelApi 结构体用于封装标签相关的 API 操作
type LabelApi struct {
	client *GewechatClient
}

// NewLabelApi 创建一个新的 LabelApi 实例
func NewLabelApi(c *GewechatClient) *LabelApi {
	return &LabelApi{client: c}
}

// Add 添加标签
func (l *LabelApi) Add(appID, labelName string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId":     appID,
		"labelName": labelName,
	}
	return l.client.PostJson("/label/add", param)
}

// Delete 删除标签
func (l *LabelApi) Delete(appID string, labelIDs []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":    appID,
		"labelIds": labelIDs,
	}
	return l.client.PostJson("/label/delete", param)
}

// List 获取标签列表
func (l *LabelApi) List(appID string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
	}
	return l.client.PostJson("/label/list", param)
}

// ModifyMemberList 修改标签成员列表
func (l *LabelApi) ModifyMemberList(appID string, labelIDs, wxIDs []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":    appID,
		"labelIds": labelIDs,
		"wxIds":    wxIDs,
	}
	return l.client.PostJson("/label/modifyMemberList", param)
}
