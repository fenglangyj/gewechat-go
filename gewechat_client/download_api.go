package gewechat_client

// DownloadApi 结构体用于封装下载相关的 API 操作
type DownloadApi struct {
	client *GewechatClient
}

// NewDownloadApi 创建一个新的 DownloadApi 实例
func NewDownloadApi(c *GewechatClient) *DownloadApi {
	return &DownloadApi{client: c}
}

// DownloadImage 下载图片
func (d *DownloadApi) DownloadImage(appID, xml, imgType string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
		"xml":   xml,
		"type":  imgType,
	}
	return d.client.PostJson("/message/downloadImage", param)
}

// DownloadVoice 下载语音
func (d *DownloadApi) DownloadVoice(appID, xml, msgID string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
		"xml":   xml,
		"msgId": msgID,
	}
	return d.client.PostJson("/message/downloadVoice", param)
}

// DownloadVideo 下载视频
func (d *DownloadApi) DownloadVideo(appID, xml string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId": appID,
		"xml":   xml,
	}
	return d.client.PostJson("/message/downloadVideo", param)
}

// DownloadEmojiMd5 下载emoji
func (d *DownloadApi) DownloadEmojiMd5(appID, emojiMd5 string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId":    appID,
		"emojiMd5": emojiMd5,
	}
	return d.client.PostJson("/message/downloadEmojiMd5", param)
}

// DownloadCdn cdn下载
func (d *DownloadApi) DownloadCdn(appID, aesKey, fileID, fileType, totalSize, suffix string) (map[string]interface{}, error) {
	param := map[string]string{
		"appId":     appID,
		"aesKey":    aesKey,
		"fileId":    fileID,
		"totalSize": totalSize,
		"type":      fileType,
		"suffix":    suffix,
	}
	return d.client.PostJson("/message/downloadCdn", param)
}
