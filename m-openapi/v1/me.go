package v1

import (
	"context"
	"demo/m-botgo/errs"
	m_dto "demo/m-dto"
	"encoding/json"
)

// Me 拉取当前用户的信息
func (o *openAPI) Me(ctx context.Context) (*m_dto.User, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.User{}).
		Get(o.getURL(userMeURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.User), nil
}

// MeGuilds 拉取当前用户加入的频道列表
func (o *openAPI) MeGuilds(ctx context.Context, pager *m_dto.GuildPager) ([]*m_dto.Guild, error) {
	if pager == nil {
		return nil, errs.ErrPagerIsNil
	}
	resp, err := o.request(ctx).
		SetQueryParams(pager.QueryParams()).
		Get(o.getURL(userMeGuildsURI))
	if err != nil {
		return nil, err
	}

	guilds := make([]*m_dto.Guild, 0)
	if err := json.Unmarshal(resp.Body(), &guilds); err != nil {
		return nil, err
	}

	return guilds, nil
}
