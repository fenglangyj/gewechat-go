package gewechat_client

// ContactApi 结构体表示联系人 API
type ContactApi struct {
	client *GewechatClient
}

// NewContactApi 创建一个新的 ContactApi 实例
func NewContactApi(c *GewechatClient) *ContactApi {
	return &ContactApi{client: c}
}

// FetchContactsList 获取通讯录列表
func (c *ContactApi) FetchContactsList(appID string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
	}
	return c.client.PostJson("/contacts/fetchContactsList", param)
}

// GetBriefInfo 获取群/好友简要信息
func (c *ContactApi) GetBriefInfo(appID string, wxids []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
		"wxids": wxids,
	}
	return c.client.PostJson("/contacts/getBriefInfo", param)
}

// GetDetailInfo 获取群/好友详细信息
func (c *ContactApi) GetDetailInfo(appID string, wxids []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
		"wxids": wxids,
	}
	return c.client.PostJson("/contacts/getDetailInfo", param)
}

// Search 搜索好友
func (c *ContactApi) Search(appID string, contactsInfo interface{}) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":        appID,
		"contactsInfo": contactsInfo,
	}
	return c.client.PostJson("/contacts/search", param)
}

// AddContacts 添加联系人/同意添加好友
func (c *ContactApi) AddContacts(appID, scene, option, v3, v4, content string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId":   appID,
		"scene":   scene,
		"option":  option,
		"v3":      v3,
		"v4":      v4,
		"content": content,
	}
	return c.client.PostJson("/contacts/addContacts", param)
}

// DeleteFriend 删除好友
func (c *ContactApi) DeleteFriend(appID, wxid string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
		"wxid":  wxid,
	}
	return c.client.PostJson("/contacts/deleteFriend", param)
}

// SetFriendPermissions 设置好友仅聊天
func (c *ContactApi) SetFriendPermissions(appID, wxid string, onlyChat bool) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":    appID,
		"wxid":     wxid,
		"onlyChat": onlyChat,
	}
	return c.client.PostJson("/contacts/setFriendPermissions", param)
}

// SetFriendRemark 设置好友备注
func (c *ContactApi) SetFriendRemark(appID, wxid, remark string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId":    appID,
		"wxid":     wxid,
		"onlyChat": remark,
	}
	return c.client.PostJson("/contacts/setFriendRemark", param)
}

// GetPhoneAddressList 获取手机通讯录
func (c *ContactApi) GetPhoneAddressList(appID string, phones []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
		"wxid":  phones,
	}
	return c.client.PostJson("/contacts/getPhoneAddressList", param)
}

// UploadPhoneAddressList 上传手机通讯录
func (c *ContactApi) UploadPhoneAddressList(appID string, phones []string, opType int) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"wxid":   phones,
		"opType": opType,
	}
	return c.client.PostJson("/contacts/uploadPhoneAddressList", param)
}
