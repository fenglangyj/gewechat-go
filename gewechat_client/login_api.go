package gewechat_client

import (
	"fmt"
	"time"
)

// LoginApi 登录模块
type LoginApi struct {
	client *GewechatClient
}

func NewLoginApi(c *GewechatClient) *LoginApi {
	return &LoginApi{client: c}
}

// GetToken 获取tokenId
func (loginApi *LoginApi) GetToken() (map[string]interface{}, error) {
	return loginApi.client.PostJson("/tools/getTokenId", nil)
}

// SetCallback 设置微信消息的回调地址
func (loginApi *LoginApi) SetCallback(callbackURL string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"token":       loginApi.client.token,
		"callbackUrl": callbackURL,
	}
	return loginApi.client.PostJson("/tools/setCallback", param)
}

// GetQR 获取登录二维码
func (loginApi *LoginApi) GetQR(appID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
	}
	return loginApi.client.PostJson("/login/getLoginQrCode", param)
}

// CheckQR 确认登陆
func (loginApi *LoginApi) CheckQR(appID, uuid, captchCode string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"uuid":       uuid,
		"captchCode": captchCode,
	}
	return loginApi.client.PostJson("/login/checkLogin", param)
}

// Logout 退出微信
func (loginApi *LoginApi) Logout(appID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
	}
	return loginApi.client.PostJson("/login/logout", param)
}

// DialogLogin 弹框登录
func (loginApi *LoginApi) DialogLogin(appID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
	}
	return loginApi.client.PostJson("/login/dialogLogin", param)
}

// DeviceList 查看设备列表
func (loginApi *LoginApi) DeviceList() (map[string]interface{}, error) {
	param := map[string]interface{}{}
	return loginApi.client.PostJson("/login/deviceList", param)
}

// CheckOnline 检查是否在线
func (loginApi *LoginApi) CheckOnline(appID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
	}
	return loginApi.client.PostJson("/login/checkOnline", param)
}

// GetAndValidateQR 获取并验证二维码数据
func (loginApi *LoginApi) GetAndValidateQR(appID string) (string, string, error) {
	result, err := loginApi.GetQR(appID)
	if err != nil {
		return "", "", err
	}

	if result["ret"] != float64(200) {
		return "", "", fmt.Errorf("获取二维码失败: %v", result)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("无法解析二维码数据: %v", result)
	}

	appID, ok = data["appId"].(string)
	if !ok {
		return "", "", fmt.Errorf("无法获取appId: %v", result)
	}

	uuid, ok := data["uuid"].(string)
	if !ok {
		return "", "", fmt.Errorf("无法获取uuid: %v", result)
	}

	return appID, uuid, nil
}

// Login 执行完整的登录流程
func (loginApi *LoginApi) Login(appID string) (string, string, error) {
	// 1. 检查是否已经登录
	inputAppID := appID
	if inputAppID != "" {
		result, err := loginApi.CheckOnline(inputAppID)
		if err != nil {
			return "", "", err
		}
		if result["ret"] == float64(200) && result["data"] == true {
			return inputAppID, "", nil
		}
	}

	// 2. 获取登录二维码
	appID, uuid, err := loginApi.GetAndValidateQR(appID)
	if err != nil {
		return "", "获取二维码失败", err
	}

	// 这里可以添加生成和打印二维码的逻辑
	fmt.Printf("二维码链接: http://weixin.qq.com/x/%s\n", uuid)

	// 3. 轮询检查登录状态
	retryCount := 0
	maxRetries := 100 // 最大重试100次

	for retryCount < maxRetries {
		result, err := loginApi.CheckQR(appID, uuid, "")
		if err != nil {
			return appID, fmt.Sprintf("检查登录状态失败: %v", err), err
		}
		if result["ret"] != float64(200) {
			return appID, fmt.Sprintf("检查登录状态失败: %v", result), nil
		}

		data, ok := result["data"].(map[string]interface{})
		if !ok {
			return appID, fmt.Sprintf("无法解析登录数据: %v", result), nil
		}

		status := int(data["status"].(float64))
		expiredTime := int(data["expiredTime"].(float64))

		// 检查二维码是否过期，提前5秒重新获取
		if expiredTime <= 5 {
			fmt.Println("二维码即将过期，正在重新获取...")
			_, uuid, err = loginApi.GetAndValidateQR(appID)
			if err != nil {
				return appID, "重新获取二维码失败", err
			}

			// 这里可以添加生成和打印新二维码的逻辑
			fmt.Printf("新二维码链接: http://weixin.qq.com/x/%s\n", uuid)
			continue
		}

		if status == 2 { // 登录成功
			nickName := data["nickName"].(string)
			if nickName == "" {
				nickName = "未知用户"
			}
			fmt.Printf("\n登录成功！用户昵称: %s\n", nickName)
			return appID, "", nil
		} else {
			retryCount++
			if retryCount >= maxRetries {
				fmt.Println("登录超时，请重新尝试")
				return appID, "登录超时，请重新尝试", nil
			}
			time.Sleep(5 * time.Second)
		}
		retryCount++
	}
	return appID, "登录超时，请重新尝试", nil
}
