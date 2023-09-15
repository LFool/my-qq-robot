package m_dto

// Member 群成员
type Member struct {
	GuildID  string    `json:"guild_id"`
	JoinedAt Timestamp `json:"joined_at"`
	Nick     string    `json:"nick"`
	User     *User     `json:"user"`
	Roles    []string  `json:"roles"`
}

// DeleteHistoryMsgDay 消息撤回天数
type DeleteHistoryMsgDay = int

// MemberDeleteOpts 删除成员额外参数
type MemberDeleteOpts struct {
	AddBlackList         bool                `json:"add_blacklist"`
	DeleteHistoryMsgDays DeleteHistoryMsgDay `json:"delete_history_msg_days"`
}

// MemberDeleteOption 删除成员选项
type MemberDeleteOption func(*MemberDeleteOpts)
