package v1

import (
	"context"
	m_dto "demo/m-dto"
)

// WS 获取带分片 WSS 接入点
func (o *openAPI) WS(ctx context.Context, _ map[string]string, _ string) (*m_dto.WebsocketAP, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.WebsocketAP{}).
		Get(o.getURL(gatewayBotURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.WebsocketAP), nil
}
