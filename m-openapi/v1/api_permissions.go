package v1

import (
	"context"
	m_dto "demo/m-dto"
)

// GetAPIPermissions 获取频道可用权限列表
func (o *openAPI) GetAPIPermissions(ctx context.Context, guildID string) (*m_dto.APIPermissions, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.APIPermissions{}).
		SetPathParam("guild_id", guildID).
		Get(o.getURL(apiPermissionURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*m_dto.APIPermissions), nil
}

// RequireAPIPermissions 创建频道 API 接口权限授权链接
func (o *openAPI) RequireAPIPermissions(ctx context.Context,
	guildID string, demand *m_dto.APIPermissionDemandToCreate) (*m_dto.APIPermissionDemand, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.APIPermissionDemand{}).
		SetPathParam("guild_id", guildID).
		SetBody(demand).
		Post(o.getURL(apiPermissionDemandURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*m_dto.APIPermissionDemand), nil
}
