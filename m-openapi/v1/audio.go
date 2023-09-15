package v1

import (
	"context"
	m_dto "demo/m-dto"
)

// PostAudio AudioAPI 接口实现
func (o openAPI) PostAudio(ctx context.Context, channelID string, value *m_dto.AudioControl) (*m_dto.AudioControl, error) {
	// 目前服务端成功不回包
	_, err := o.request(ctx).
		SetResult(m_dto.Channel{}).
		SetPathParam("channel_id", channelID).
		SetBody(value).
		Post(o.getURL(audioControlURI))
	if err != nil {
		return nil, err
	}

	return value, nil
}
