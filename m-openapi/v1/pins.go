package v1

import (
	"context"
	m_dto "demo/m-dto"
)

// AddPins 添加精华消息
func (o *openAPI) AddPins(ctx context.Context, channelID string, messageID string) (*m_dto.PinsMessage, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Put(o.getURL(pinURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*m_dto.PinsMessage), nil
}

// DeletePins 删除精华消息
func (o *openAPI) DeletePins(ctx context.Context, channelID, messageID string) error {
	_, err := o.request(ctx).
		SetResult(m_dto.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Delete(o.getURL(pinURI))
	return err
}

// GetPins 获取精华消息
func (o *openAPI) GetPins(ctx context.Context, channelID string) (*m_dto.PinsMessage, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		Get(o.getURL(pinsURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*m_dto.PinsMessage), nil
}

// CleanPins 清除全部精华消息
func (o *openAPI) CleanPins(ctx context.Context, channelID string) error {
	_, err := o.request(ctx).
		SetResult(m_dto.PinsMessage{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", "all").
		Delete(o.getURL(pinURI))
	return err
}
