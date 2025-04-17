package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fenglangyj/gewechat-go/gewechat_client"
	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	// 初始化配置
	initConfig()
}

// 初始化配置文件
func initConfig() {
	// 配置文件名（不带扩展名）
	viper.SetConfigName("gewechat_config")
	// 配置文件类型
	viper.SetConfigType("yaml")
	// 配置文件搜索路径（当前目录）
	viper.AddConfigPath(".")
	// 自动读取环境变量
	viper.AutomaticEnv()

	// 如果配置文件不存在则创建默认配置
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// 设置默认值
			viper.Set("base_url", "http://127.0.0.1:2531/v2/api")
			viper.Set("token", "f5dcf2b61bb740668fa625ee40f7cb7a")
			//设备ID，首次登录后系统生成
			viper.Set("app_id", "wx_6wDsviA5iBsJTdgcoNu8a")
			//需要聊天的昵称
			viper.Set("send_msg_nickname", "太吉")

			// 写入配置文件
			if err := viper.SafeWriteConfigAs("gewechat_config.yaml"); err != nil {
				fmt.Println("创建配置文件失败:", err)
				os.Exit(1)
			}
			fmt.Println("已创建默认配置文件，请修改后重新运行")
			//os.Exit(0)
		}
	}
}

// 在控制台显示二维码
func showQR(str string) {
	qrterminal.GenerateWithConfig(
		//fmt.Sprintf("http://weixin.qq.com/x/%s", uuid),
		str,
		qrterminal.Config{
			Level:     qrterminal.L, // 纠错等级
			Writer:    os.Stdout,    // 输出到控制台
			BlackChar: qrterminal.BLACK,
			WhiteChar: qrterminal.WHITE,
			QuietZone: 1, // 二维码边距
		},
	)
}

// Login 微信扫描登录逻辑
func Login(appID string, gewechatClient *gewechat_client.GewechatClient) (string, error) {
	// 1. 检查是否已经登录
	inputAppID := appID
	// 如果已经登录，检测是否在线
	if inputAppID != "" {
		result, err := gewechatClient.LoginApi.CheckOnline(inputAppID)
		if err != nil {
			return "", err
		}
		if result["ret"] == float64(200) && result["data"] == true {
			fmt.Printf("AppID: %s 已在线，无需登录\n", inputAppID)
			return inputAppID, nil
		} else {
			fmt.Printf("AppID: %s 未在线，执行登录流程\n", inputAppID)
		}
	}

	// 2. 获取登录二维码
	appID, uuid, err := gewechatClient.LoginApi.GetAndValidateQR(appID)
	if err != nil {
		return "获取二维码失败", err
	}

	if inputAppID == "" {
		//保存appid, 下次登录时继续使用
		viper.Set("app_id", appID)
		if err := viper.WriteConfig(); err != nil {
			fmt.Println("保存配置失败:", err)
		} else {
			fmt.Println("已更新配置文件")
		}
		fmt.Printf("AppID: %s, 请保存此app_id，下次登录时继续使用!\n", appID)
		fmt.Println("新设备登录平台，次日凌晨会掉线一次，重新登录时需使用原来的app_id取码，否则新app_id仍然会掉线，登录成功后则可以长期在线")
	}

	// 这里可以添加生成和打印二维码的逻辑
	//fmt.Printf("二维码链接: http://weixin.qq.com/x/%s\n", uuid)

	fmt.Println("请使用手机微信扫码登录", "uuid=", uuid)
	showQR("http://weixin.qq.com/x/" + uuid)
	// 3. 轮询检查登录状态
	retryCount := 0
	maxRetries := 100 // 最大重试100次
	for retryCount < maxRetries {
		fmt.Println(retryCount, "次轮询检测是否登录成功")
		//检查登录状态
		result, err := gewechatClient.LoginApi.CheckQR(appID, uuid, "")
		if err != nil {
			return appID, fmt.Errorf("检查登录状态失败: %v", err)
		}
		if result["ret"] != float64(200) {
			return appID, fmt.Errorf("检查登录状态失败: %v", result)
		}

		data, ok := result["data"].(map[string]interface{})
		if !ok {
			return appID, fmt.Errorf("无法解析登录数据: %v", result)
		}

		status := int(data["status"].(float64))
		expiredTime := int(data["expiredTime"].(float64))

		// 检查二维码是否过期，提前5秒重新获取
		if expiredTime <= 5 {
			fmt.Println("二维码即将过期，正在重新获取...")
			_, uuid, err = gewechatClient.LoginApi.GetAndValidateQR(appID)
			if err != nil {
				return appID, fmt.Errorf("重新获取二维码失败: %v", err)
			}
			// 这里可以添加生成和打印新二维码的逻辑
			// fmt.Printf("新二维码链接: http://weixin.qq.com/x/%s\n", uuid)
			fmt.Println("请使用手机微信扫码登录", "uuid=", uuid)
			showQR("http://weixin.qq.com/x/" + uuid)
			continue
		}
		if status == 2 { // 登录成功
			nickName := data["nickName"].(string)
			if nickName == "" {
				nickName = "未知用户"
			}
			fmt.Printf("\n登录成功！用户昵称: %s\n", nickName)
			return appID, nil
		}
		//间隔5秒，重新检测是否登录成功。

		time.Sleep(5 * time.Second)
		retryCount++
	}
	fmt.Println("登录超时，请重新尝试")
	return "登录超时，请重新尝试", fmt.Errorf("%d 次重试后登录超时", maxRetries)
}
func httpServer() {
	fmt.Println("启动HTTP服务器，监听8080端口...")
	http.HandleFunc("/gewechat/callback", callbackHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP服务启动失败:", err)
		return
	}
}

func main() {
	//启动HTTP服务器
	go httpServer()
	// 创建 GewechatClient 实例:
	Client := gewechat_client.NewGewechatClient(viper.GetString("base_url"), viper.GetString("token"))
	//TokenResult, err := Client.LoginApi.GetToken()
	//fmt.Println("刚刚拿到的token", TokenResult)
	// 查看设备列表
	DeviceResult, err := Client.LoginApi.DeviceList()
	fmt.Println("已登录设备列表", DeviceResult)

	//设置消息回调地址
	callbackRes, _ := Client.LoginApi.SetCallback("http://192.168.31.165:8080/gewechat/callback")
	fmt.Println("设置消息回调地址：", callbackRes)

	appID := viper.GetString("app_id")
	appID, err = Login(appID, Client)
	if err != nil {
		fmt.Println("登录失败:", err)
		return
	}
	fmt.Println("登录后得到的设备id：", appID)

	// 获取个人信息(持久化保存)
	ProfileRes, err := Client.PersonalApi.GetProfile(appID)
	fmt.Println("获取个人信息：", ProfileRes)
	// 获取好友列表（持久化存储）
	FetchContactsListResult, _ := Client.ContactApi.FetchContactsList(appID)
	fmt.Println("获取好友数据：", FetchContactsListResult)
	data, ok := FetchContactsListResult["data"].(map[string]interface{})
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

	/*chatrooms, ok := data["chatrooms"].([]interface{})
	if !ok || len(chatrooms) == 0 {
		fmt.Println("获取到的群列表为空")
		return
	}
	fmt.Println("获取到的群列表:", chatrooms)

	ghs, ok := data["ghs"].([]interface{})
	if !ok || len(ghs) == 0 {
		fmt.Println("获取到的ghs列表为空")
		return
	}
	fmt.Println("获取到的ghs列表:", ghs)*/

	// 获取好友简要信息(包含群信息)（一次最多100条）

	//todo 给好友发消息
	//todo 添加好友
	//todo 发布朋友圈消息

	// 获取好友的简要信息
	var wxids []string
	for _, friend := range friends {
		wxids = append(wxids, friend.(string))
	}

	var dataList []interface{}
	wxidsChunks := chunkSlice(wxids, 100) // 分块函数见下方

	for _, chunk := range wxidsChunks {
		// 获取好友简要信息（一次最多100条）
		chunkFriendsInfo, err := Client.ContactApi.GetBriefInfo(appID, chunk)
		if err != nil {
			fmt.Println("获取好友简要信息失败:", err)
			return
		}
		ret, ok := chunkFriendsInfo["ret"].(float64)
		if !ok || ret != 200 {
			fmt.Println("获取好友简要信息失败:", chunkFriendsInfo)
			return
		}
		chunkData, ok := chunkFriendsInfo["data"].([]interface{})
		if !ok || len(chunkData) == 0 {
			fmt.Println("当前分块获取到的好友简要信息列表为空")
			continue
		}
		fmt.Println("拿到最多100条好友信息，可以在这儿加数据持久化功能:", chunkData)
		dataList = append(dataList, chunkData...)
	}

	// 找对目标好友的wxid
	wxid := ""
	for _, friendInfo := range dataList {
		info, ok := friendInfo.(map[string]interface{})
		if !ok {
			continue
		}
		if info["nickName"] == viper.GetString("send_msg_nickname") {
			fmt.Println("找到好友:", info)
			wxid = info["userName"].(string)
			break
		}
	}
	if wxid == "" {
		fmt.Printf("没有找到好友: %s 的wxid\n", viper.GetString("send_msg_nickname"))
		return
	}
	fmt.Println("找到好友:", viper.GetString("send_msg_nickname"), "", wxid)

	// 发送消息
	sendMsgResult, err := Client.MessageApi.PostText(appID, wxid, "你好啊", nil)
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
	ret, ok := sendMsgResult["ret"].(float64)
	if !ok || ret != 200 {
		fmt.Println("发送消息失败:", sendMsgResult)
		return
	}
	fmt.Println("发送消息成功:", sendMsgResult)

	//循环 接收终端输入，给好友“wxid”发消息
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入要发送的消息内容: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		// 发送消息
		_, err := Client.MessageApi.PostText(appID, wxid, input, nil)
		if err != nil {
			fmt.Printf("发送失败: %v\n", err)
			continue
		}
		// 处理发送结果
		/*if ret, ok := sendResult["ret"].(float64); ok && ret == 200 {
			fmt.Println("✓ 消息发送成功")
		} else {
			fmt.Printf("× 消息发送失败: %+v\n", sendResult)
		}*/
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	// 限制请求体大小
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB
	// 只处理POST请求
	if r.Method != http.MethodPost {
		http.Error(w, `{"ret": 405, "msg": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}
	// 解析JSON请求体
	var payload map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf(`{"ret": 400, "msg": "Invalid request body: %v"}`, err), http.StatusBadRequest)
		return
	}
	// 这里添加你的业务逻辑处理
	fmt.Printf("收到回调数据：%+v\n", payload)
	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, `{"ret": 200,"msg": "success"}`)
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
