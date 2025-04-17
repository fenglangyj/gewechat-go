package gewechat_client

// PersonalApi 定义 PersonalApi 结构体
type PersonalApi struct {
	client *GewechatClient
}

// NewPersonalApi 创建一个新的 PersonalApi 实例
func NewPersonalApi(c *GewechatClient) *PersonalApi {
	return &PersonalApi{client: c}
}

// GetProfile 获取个人资料
func (p *PersonalApi) GetProfile(appID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
	}
	return p.client.PostJson("/personal/getProfile", param)
}

// GetQrCode 获取自己的二维码
func (p *PersonalApi) GetQrCode(appID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
	}
	return p.client.PostJson("/personal/getQrCode", param)
}

// GetSafetyInfo 获取设备记录
func (p *PersonalApi) GetSafetyInfo(appID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
	}
	return p.client.PostJson("/personal/getSafetyInfo", param)
}

// PrivacySettings 隐私设置
func (p *PersonalApi) PrivacySettings(appID, option string, open bool) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"option": option,
		"open":   open,
	}
	return p.client.PostJson("/personal/privacySettings", param)
}

// UpdateProfile 修改个人信息
func (p *PersonalApi) UpdateProfile(appID, city, country, nickName, province, sex, signature string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":     appID,
		"city":      city,
		"country":   country,
		"nickName":  nickName,
		"province":  province,
		"sex":       sex,
		"signature": signature,
	}
	return p.client.PostJson("/personal/updateProfile", param)
}

// UpdateHeadImg 修改头像
func (p *PersonalApi) UpdateHeadImg(appID, headImgURL string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"headImgUrl": headImgURL,
	}
	return p.client.PostJson("/personal/updateHeadImg", param)
}
