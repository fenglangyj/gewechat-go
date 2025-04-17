package gewechat_client

type MessageApi struct {
	client *GewechatClient
}

func NewMessageApi(c *GewechatClient) *MessageApi {
	return &MessageApi{client: c}

}

// ForwardMiniApp 转发小程序
func (m *MessageApi) ForwardMiniApp(appID, toWxid, xml, coverImgURL string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":       appID,
		"toWxid":      toWxid,
		"xml":         xml,
		"coverImgUrl": coverImgURL,
	}
	return m.client.PostJson("/tools/forwardMiniApp", param)
}

// ForwardURL 转发链接
func (m *MessageApi) ForwardURL(appID, toWxid, xml string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"toWxid": toWxid,
		"xml":    xml,
	}
	return m.client.PostJson("/message/forwardUrl", param)
}

// RevokeMsg 撤回消息
func (m *MessageApi) RevokeMsg(appID, toWxid, msgID, newMsgID string, createTime int64) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"toWxid":     toWxid,
		"msgId":      msgID,
		"newMsgId":   newMsgID,
		"createTime": createTime,
	}
	return m.client.PostJson("/message/revokeMsg", param)
}

// PostText 发送文字消息
func (m *MessageApi) PostText(appID, toWxid, content string, ats []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":   appID,
		"toWxid":  toWxid,
		"content": content,
		"ats":     ats,
	}
	return m.client.PostJson("/message/postText", param)
}

// PostFile 发送文件消息
func (m *MessageApi) PostFile(appID, toWxid, fileURL, fileName string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":    appID,
		"toWxid":   toWxid,
		"fileUrl":  fileURL,
		"fileName": fileName,
	}
	return m.client.PostJson("/message/postFile", param)
}

// PostImage 发送图片消息
func (m *MessageApi) PostImage(appID, toWxid, imgURL string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"toWxid": toWxid,
		"imgUrl": imgURL,
	}
	return m.client.PostJson("/message/postImage", param)
}

// PostVoice 发送语音消息
func (m *MessageApi) PostVoice(appID, toWxid, voiceURL string, voiceDuration int64) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":         appID,
		"toWxid":        toWxid,
		"voiceUrl":      voiceURL,
		"voiceDuration": voiceDuration,
	}
	return m.client.PostJson("/message/postVoice", param)
}

// PostVideo 发送视频消息
func (m *MessageApi) PostVideo(appID, toWxid, videoURL, thumbURL string, videoDuration int64) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":         appID,
		"toWxid":        toWxid,
		"videoUrl":      videoURL,
		"thumbUrl":      thumbURL,
		"videoDuration": videoDuration,
	}
	return m.client.PostJson("/message/postVideo", param)
}

// PostLink 发送链接消息
func (m *MessageApi) PostLink(appID, toWxid, title, desc, linkURL, thumbURL string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":    appID,
		"toWxid":   toWxid,
		"title":    title,
		"desc":     desc,
		"linkUrl":  linkURL,
		"thumbUrl": thumbURL,
	}
	return m.client.PostJson("/message/postLink", param)
}

// PostNameCard 发送名片消息
func (m *MessageApi) PostNameCard(appID, toWxid, nickName, nameCardWxid string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":        appID,
		"toWxid":       toWxid,
		"nickName":     nickName,
		"nameCardWxid": nameCardWxid,
	}
	return m.client.PostJson("/message/postNameCard", param)
}

// PostEmoji 发送emoji消息
func (m *MessageApi) PostEmoji(appID, toWxid, emojiMD5 string, emojiSize int64) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":     appID,
		"toWxid":    toWxid,
		"emojiMd5":  emojiMD5,
		"emojiSize": emojiSize,
	}
	return m.client.PostJson("/message/postEmoji", param)
}

// PostAppMsg 发送appmsg消息
func (m *MessageApi) PostAppMsg(appID, toWxid string, appmsg interface{}) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"toWxid": toWxid,
		"appmsg": appmsg,
	}
	return m.client.PostJson("/message/postAppMsg", param)
}

// PostMiniApp 发送小程序消息
func (m *MessageApi) PostMiniApp(appID, toWxid, miniAppID, displayName, pagePath, coverImgURL, title, userName string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":       appID,
		"toWxid":      toWxid,
		"miniAppId":   miniAppID,
		"displayName": displayName,
		"pagePath":    pagePath,
		"coverImgUrl": coverImgURL,
		"title":       title,
		"userName":    userName,
	}
	return m.client.PostJson("/message/postMiniApp", param)
}

// ForwardFile 转发文件
func (m *MessageApi) ForwardFile(appID, toWxid, xml string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"toWxid": toWxid,
		"xml":    xml,
	}
	return m.client.PostJson("/message/forwardFile", param)
}

// ForwardImage 转发图片
func (m *MessageApi) ForwardImage(appID, toWxid, xml string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"toWxid": toWxid,
		"xml":    xml,
	}
	return m.client.PostJson("/message/forwardImage", param)
}

// ForwardVideo 转发视频
func (m *MessageApi) ForwardVideo(appID, toWxid, xml string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"toWxid": toWxid,
		"xml":    xml,
	}
	return m.client.PostJson("/message/forwardVideo", param)
}

// ForwardUrl 转发链接
func (m *MessageApi) ForwardUrl(appID, toWxid, xml string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":  appID,
		"toWxid": toWxid,
		"xml":    xml,
	}
	return m.client.PostJson("/message/forwardUrl", param)
}
