package gewechat_client

// GroupApi 定义 GroupApi 结构体
type GroupApi struct {
	client *GewechatClient
}

// NewGroupApi 创建一个新的 GroupApi 实例
func NewGroupApi(c *GewechatClient) *GroupApi {
	return &GroupApi{client: c}
}

// CreateChatroom 创建微信群
func (g *GroupApi) CreateChatroom(appID string, wxids []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
		"wxids": wxids,
	}
	return g.client.PostJson("/group/createChatroom", param)
}

// ModifyChatroomName 修改群名称
func (g *GroupApi) ModifyChatroomName(appID, chatroomName, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":        appID,
		"chatroomName": chatroomName,
		"chatroomId":   chatroomID,
	}
	return g.client.PostJson("/group/modifyChatroomName", param)
}

// ModifyChatroomRemark 修改群备注
func (g *GroupApi) ModifyChatroomRemark(appID, chatroomRemark, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":          appID,
		"chatroomRemark": chatroomRemark,
		"chatroomId":     chatroomID,
	}
	return g.client.PostJson("/group/modifyChatroomRemark", param)
}

// ModifyChatroomNicknameForSelf 修改我在群内的昵称
func (g *GroupApi) ModifyChatroomNicknameForSelf(appID, nickName, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"nickName":   nickName,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/modifyChatroomNickNameForSelf", param)
}

// InviteMember 邀请/添加 进群
func (g *GroupApi) InviteMember(appID string, wxids []string, chatroomID, reason string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"wxids":      wxids,
		"reason":     reason,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/inviteMember", param)
}

// RemoveMember 删除群成员
func (g *GroupApi) RemoveMember(appID string, wxids []string, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"wxids":      wxids,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/removeMember", param)
}

// QuitChatroom 退出群聊
func (g *GroupApi) QuitChatroom(appID, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/quitChatroom", param)
}

// DisbandChatroom 解散群聊
func (g *GroupApi) DisbandChatroom(appID, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/disbandChatroom", param)
}

// GetChatroomInfo 获取群信息
func (g *GroupApi) GetChatroomInfo(appID, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/getChatroomInfo", param)
}

// GetChatroomMemberList 获取群成员列表
func (g *GroupApi) GetChatroomMemberList(appID, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/getChatroomMemberList", param)
}

// GetChatroomMemberDetail 获取群成员详情
func (g *GroupApi) GetChatroomMemberDetail(appID, chatroomID string, memberWxids []string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":       appID,
		"memberWxids": memberWxids,
		"chatroomId":  chatroomID,
	}
	return g.client.PostJson("/group/getChatroomMemberDetail", param)
}

// GetChatroomAnnouncement 获取群公告
func (g *GroupApi) GetChatroomAnnouncement(appID, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/getChatroomAnnouncement", param)
}

// SetChatroomAnnouncement 设置群公告
func (g *GroupApi) SetChatroomAnnouncement(appID, chatroomID, content string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"chatroomId": chatroomID,
		"content":    content,
	}
	return g.client.PostJson("/group/setChatroomAnnouncement", param)
}

// AgreeJoinRoom 同意进群
func (g *GroupApi) AgreeJoinRoom(appID, url string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":        appID,
		"chatroomName": url,
	}
	return g.client.PostJson("/group/agreeJoinRoom", param)
}

// AddGroupMemberAsFriend 添加群成员为好友
func (g *GroupApi) AddGroupMemberAsFriend(appID, memberWxid, chatroomID, content string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"memberWxid": memberWxid,
		"content":    content,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/addGroupMemberAsFriend", param)
}

// GetChatroomQrCode 获取群二维码
func (g *GroupApi) GetChatroomQrCode(appID, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/getChatroomQrCode", param)
}

// SaveContractList 群保存到通讯录或从通讯录移除
func (g *GroupApi) SaveContractList(appID, operType, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"operType":   operType,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/saveContractList", param)
}

// AdminOperate 管理员操作
func (g *GroupApi) AdminOperate(appID, chatroomID string, wxids []string, operType string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"wxids":      wxids,
		"operType":   operType,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/adminOperate", param)
}

// PinChat 聊天置顶
func (g *GroupApi) PinChat(appID, top, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"top":        top,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/pinChat", param)
}

// SetMsgSilence 设置消息免打扰
func (g *GroupApi) SetMsgSilence(appID, silence, chatroomID string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"silence":    silence,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/setMsgSilence", param)
}

// JoinRoomUsingQrCode 扫码进群
func (g *GroupApi) JoinRoomUsingQrCode(appID, qrURL string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId": appID,
		"qrUrl": qrURL,
	}
	return g.client.PostJson("/group/joinRoomUsingQRCode", param)
}

// RoomAccessApplyCheckApprove 确认进群申请
func (g *GroupApi) RoomAccessApplyCheckApprove(appID, newMsgID, chatroomID, msgContent string) (map[string]interface{}, error) {
	param := map[string]interface{}{
		"appId":      appID,
		"newMsgId":   newMsgID,
		"msgContent": msgContent,
		"chatroomId": chatroomID,
	}
	return g.client.PostJson("/group/roomAccessApplyCheckApprove", param)
}
