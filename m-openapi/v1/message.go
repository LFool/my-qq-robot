package v1

import (
	"context"
	"demo/m-botgo/errs"
	m_dto "demo/m-dto"
	"encoding/json"
)

// Message 拉取单条消息
func (o *openAPI) Message(ctx context.Context, channelID string, messageID string) (*m_dto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.Message{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Get(o.getURL(messageURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.Message), nil
}

// Messages 拉取消息列表
func (o *openAPI) Messages(ctx context.Context, channelID string, pager *m_dto.MessagesPager) ([]*m_dto.Message, error) {
	if pager == nil {
		return nil, errs.ErrPagerIsNil
	}
	resp, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetQueryParams(pager.QueryParams()).
		Get(o.getURL(messagesURI))
	if err != nil {
		return nil, err
	}

	messages := make([]*m_dto.Message, 0)
	if err := json.Unmarshal(resp.Body(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// PostMessage 发消息
func (o *openAPI) PostMessage(ctx context.Context, channelID string, msg *m_dto.MessageToCreate) (*m_dto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.Message{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Post(o.getURL(messagesURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.Message), nil
}

// RetractMessage 撤回消息
func (o *openAPI) RetractMessage(ctx context.Context, channelID, msgID string) error {
	_, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", string(msgID)).
		Delete(o.getURL(messageURI))
	return err
}
