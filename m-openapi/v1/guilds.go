package v1

import (
	"context"
	m_dto "demo/m-dto"
)

// Guild 拉取频道信息
func (o *openAPI) Guild(ctx context.Context, guildID string) (*m_dto.Guild, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.Guild{}).
		SetPathParam("guild_id", guildID).
		Get(o.getURL(guildURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.Guild), nil
}
