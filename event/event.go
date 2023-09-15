package event

import (
	m_dto "demo/m-dto"
	"encoding/json"
	"github.com/tidwall/gjson"
)

var eventParseFuncMap = map[m_dto.OPCode]map[m_dto.EventType]eventParseFunc{
	m_dto.WSDispatchEvent: {
		m_dto.EventGuildCreate: guildHandler,
		m_dto.EventGuildUpdate: guildHandler,
		m_dto.EventGuildDelete: guildHandler,

		m_dto.EventChannelCreate: channelHandler,
		m_dto.EventChannelUpdate: channelHandler,
		m_dto.EventChannelDelete: channelHandler,

		m_dto.EventGuildMemberAdd:    guildMemberHandler,
		m_dto.EventGuildMemberUpdate: guildMemberHandler,
		m_dto.EventGuildMemberRemove: guildMemberHandler,

		m_dto.EventMessageCreate: messageHandler,

		m_dto.EventMessageReactionAdd:    messageReactionHandler,
		m_dto.EventMessageReactionRemove: messageReactionHandler,

		m_dto.EventAtMessageCreate:     atMessageHandler,
		m_dto.EventDirectMessageCreate: directMessageHandler,

		m_dto.EventAudioStart:  audioHandler,
		m_dto.EventAudioFinish: audioHandler,
		m_dto.EventAudioOnMic:  audioHandler,
		m_dto.EventAudioOffMic: audioHandler,

		m_dto.EventMessageAuditPass:   messageAuditHandler,
		m_dto.EventMessageAuditReject: messageAuditHandler,
	},
}

type eventParseFunc func(event *m_dto.WSPayload, message []byte) error

func ParseAndHandle(event *m_dto.WSPayload) error {
	// 指定类型的 handler
	if h, ok := eventParseFuncMap[event.OPCode][event.Type]; ok {
		return h(event, event.RawMessage)
	}
	// 透传handler，如果未注册具体类型的 handler，会统一投递到这个 handler
	if DefaultHandlers.Plain != nil {
		return DefaultHandlers.Plain(event, event.RawMessage)
	}
	return nil
}

func ParseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	return json.Unmarshal([]byte(data.String()), target)
}

func guildHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSGuildData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Guild != nil {
		return DefaultHandlers.Guild(event, data)
	}
	return nil
}

func channelHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSChannelData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Channel != nil {
		return DefaultHandlers.Channel(event, data)
	}
	return nil
}

func guildMemberHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSGuildMemberData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GuildMember != nil {
		return DefaultHandlers.GuildMember(event, data)
	}
	return nil
}

func messageHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Message != nil {
		return DefaultHandlers.Message(event, data)
	}
	return nil
}

func messageReactionHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSMessageReactionData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageReaction != nil {
		return DefaultHandlers.MessageReaction(event, data)
	}
	return nil
}

func atMessageHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSATMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.ATMessage != nil {
		return DefaultHandlers.ATMessage(event, data)
	}
	return nil
}

func directMessageHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSDirectMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.DirectMessage != nil {
		return DefaultHandlers.DirectMessage(event, data)
	}
	return nil
}

func audioHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSAudioData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.Audio != nil {
		return DefaultHandlers.Audio(event, data)
	}
	return nil
}

func messageAuditHandler(event *m_dto.WSPayload, message []byte) error {
	data := &m_dto.WSMessageAuditData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.MessageAudit != nil {
		return DefaultHandlers.MessageAudit(event, data)
	}
	return nil
}
