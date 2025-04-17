package main

import (
	"fmt"
	"os"
)
import "github.com/fenglangyj/gewechat-go/gewechat_client"

func main() {
	// 配置参数
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:2531/v2/api"
	}
	token := os.Getenv("GEWECHAT_TOKE")
	if token == "" {
		token = "f5dcf2b61bb740668fa625ee40f7cb7a"
	}
	fmt.Println("baseURL:", baseURL, "token", token)

	var appID = os.Getenv("APP_ID")
	if appID == "" {
		// 我的小号
		appID = "wx_6wDsviA5iBsJTdgcoNu8a"
		fmt.Println("环境变量内没有appID，使用我的小号:", appID)
	}

	sendMsgNickname := "太吉"

	// 创建 GewechatClient 实例
	newClient := gewechat_client.NewGewechatClient(baseURL, token)

	// 登录, 自动创建二维码，扫码后自动登录
	var err error
	appID, _, err = newClient.LoginApi.Login(appID)
	if err != nil {
		fmt.Println("登录失败:", err)
		return
	}
	//临时改的环境变量，当前窗口有用，重启后就消失了
	err = os.Setenv("APP_ID", appID)
	if err != nil {
		fmt.Println("写入环境变量失败:", err, "appID:", appID)
		return
	}

	// 获取好友列表
	fetchContactsListResult, err := newClient.ContactApi.FetchContactsList(appID)
	if err != nil {
		fmt.Println("获取通讯录列表失败:", err)
		return
	}
	ret, ok := fetchContactsListResult["ret"].(float64)
	if !ok || ret != 200 {
		fmt.Println("获取通讯录列表失败:", fetchContactsListResult)
		return
	}
	data, ok := fetchContactsListResult["data"].(map[string]interface{})
	if !ok {
		fmt.Println("获取通讯录列表失败: 数据格式错误")
		return
	}
	friends, ok := data["friends"].([]interface{})
	if !ok || len(friends) == 0 {
		fmt.Println("获取到的好友列表为空")
		return
	}
	fmt.Println("获取到的好友列表:", friends)

	// 获取好友的简要信息
	var wxids []string
	for _, friend := range friends {
		wxids = append(wxids, friend.(string))
	}

	var dataList []interface{}
	wxidsChunks := chunkSlice(wxids, 100) // 分块函数见下方

	for _, chunk := range wxidsChunks {
		chunkFriendsInfo, err := newClient.ContactApi.GetBriefInfo(appID, chunk)
		if err != nil {
			fmt.Println("获取好友简要信息失败:", err)
			return
		}
		ret, ok = chunkFriendsInfo["ret"].(float64)
		if !ok || ret != 200 {
			fmt.Println("获取好友简要信息失败:", chunkFriendsInfo)
			return
		}
		chunkData, ok := chunkFriendsInfo["data"].([]interface{})
		if !ok || len(chunkData) == 0 {
			fmt.Println("当前分块获取到的好友简要信息列表为空")
			continue
		}
		dataList = append(dataList, chunkData...)
	}

	/*friendsInfo, err := newClient.ContactApi.GetBriefInfo(appID, wxids)
	if err != nil {
		fmt.Println("获取好友简要信息失败:", err)
		return
	}
	ret, ok = friendsInfo["ret"].(float64)
	if !ok || ret != 200 {
		fmt.Println("获取好友简要信息失败:", friendsInfo)
		return
	}
	dataList, ok := friendsInfo["data"].([]interface{})
	if !ok || len(dataList) == 0 {
		fmt.Println("获取到的好友简要信息列表为空")
		return
	}*/

	// 找对目标好友的wxid
	wxid := ""
	for _, friendInfo := range dataList {
		info, ok := friendInfo.(map[string]interface{})
		if !ok {
			continue
		}
		if info["nickName"] == sendMsgNickname {
			fmt.Println("找到好友:", info)
			wxid = info["userName"].(string)
			break
		}
	}
	if wxid == "" {
		fmt.Printf("没有找到好友: %s 的wxid\n", sendMsgNickname)
		return
	}
	fmt.Println("找到好友:", wxid)

	// 发送消息
	sendMsgResult, err := newClient.MessageApi.PostText(appID, wxid, "你好啊", nil)
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
	ret, ok = sendMsgResult["ret"].(float64)
	if !ok || ret != 200 {
		fmt.Println("发送消息失败:", sendMsgResult)
		return
	}
	fmt.Println("发送消息成功:", sendMsgResult)

}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
