package event

import m_dto "demo/m-dto"

// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready       ReadyHandler
	ErrorNotify ErrorNotifyHandler
	Plain       PlainEventHandler

	Guild       GuildEventHandler
	GuildMember GuildMemberEventHandler
	Channel     ChannelEventHandler

	Message         MessageEventHandler
	MessageReaction MessageReactionEventHandler
	ATMessage       ATMessageEventHandler
	DirectMessage   DirectMessageEventHandler
	MessageAudit    MessageAuditEventHandler
	Audio           AudioEventHandler
}

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *m_dto.WSPayload, data *m_dto.WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *m_dto.WSPayload, message []byte) error

// GuildEventHandler 频道事件handler
type GuildEventHandler func(event *m_dto.WSPayload, data *m_dto.WSGuildData) error

// GuildMemberEventHandler 频道成员事件 handler
type GuildMemberEventHandler func(event *m_dto.WSPayload, data *m_dto.WSGuildMemberData) error

// ChannelEventHandler 子频道事件 handler
type ChannelEventHandler func(event *m_dto.WSPayload, data *m_dto.WSChannelData) error

// MessageEventHandler 消息事件 handler
type MessageEventHandler func(event *m_dto.WSPayload, data *m_dto.WSMessageData) error

// MessageReactionEventHandler 表情表态事件 handler
type MessageReactionEventHandler func(event *m_dto.WSPayload, data *m_dto.WSMessageReactionData) error

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *m_dto.WSPayload, data *m_dto.WSATMessageData) error

// DirectMessageEventHandler 私信消息事件 handler
type DirectMessageEventHandler func(event *m_dto.WSPayload, data *m_dto.WSDirectMessageData) error

// AudioEventHandler 音频机器人事件 handler
type AudioEventHandler func(event *m_dto.WSPayload, data *m_dto.WSAudioData) error

// MessageAuditEventHandler 消息审核事件 handler
type MessageAuditEventHandler func(event *m_dto.WSPayload, data *m_dto.WSMessageAuditData) error

// RegisterHandlers 注册事件回调，并返回 intent 用于 websocket 的鉴权
func RegisterHandlers(handlers ...interface{}) m_dto.Intent {
	var i m_dto.Intent
	for _, h := range handlers {
		switch handle := h.(type) {
		case ReadyHandler:
			DefaultHandlers.Ready = handle
		case ErrorNotifyHandler:
			DefaultHandlers.ErrorNotify = handle
		case PlainEventHandler:
			DefaultHandlers.Plain = handle
		case AudioEventHandler:
			DefaultHandlers.Audio = handle
			i = i | m_dto.EventToIntent(
				m_dto.EventAudioStart, m_dto.EventAudioFinish,
				m_dto.EventAudioOnMic, m_dto.EventAudioOffMic,
			)
		default:
		}
	}
	i = i | registerRelationHandlers(i, handlers...)
	i = i | registerMessageHandlers(i, handlers...)
	return i
}

// registerRelationHandlers 注册频道关系链相关handlers
func registerRelationHandlers(i m_dto.Intent, handlers ...interface{}) m_dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case GuildEventHandler:
			DefaultHandlers.Guild = handle
			i = i | m_dto.EventToIntent(m_dto.EventGuildCreate, m_dto.EventGuildDelete, m_dto.EventGuildUpdate)
		case GuildMemberEventHandler:
			DefaultHandlers.GuildMember = handle
			i = i | m_dto.EventToIntent(m_dto.EventGuildMemberAdd, m_dto.EventGuildMemberRemove, m_dto.EventGuildMemberUpdate)
		case ChannelEventHandler:
			DefaultHandlers.Channel = handle
			i = i | m_dto.EventToIntent(m_dto.EventChannelCreate, m_dto.EventChannelDelete, m_dto.EventChannelUpdate)
		default:
		}
	}
	return i
}

// registerMessageHandlers 注册消息相关的 handler
func registerMessageHandlers(i m_dto.Intent, handlers ...interface{}) m_dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case MessageEventHandler:
			DefaultHandlers.Message = handle
			i = i | m_dto.EventToIntent(m_dto.EventMessageCreate)
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handle
			i = i | m_dto.EventToIntent(m_dto.EventAtMessageCreate)
		case DirectMessageEventHandler:
			DefaultHandlers.DirectMessage = handle
			i = i | m_dto.EventToIntent(m_dto.EventDirectMessageCreate)
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
			i = i | m_dto.EventToIntent(m_dto.EventMessageReactionAdd, m_dto.EventMessageReactionRemove)
		case MessageAuditEventHandler:
			DefaultHandlers.MessageAudit = handle
			i = i | m_dto.EventToIntent(m_dto.EventMessageAuditPass, m_dto.EventMessageAuditReject)
		default:
		}
	}
	return i
}
