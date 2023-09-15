package v1

import (
	"context"
	"demo/m-botgo/errs"
	m_dto "demo/m-dto"
	"encoding/json"
)

// MemberAddRole 添加成员角色
func (o *openAPI) MemberAddRole(
	ctx context.Context, guildID string, roleID m_dto.RoleID, userID string,
	value *m_dto.MemberAddRoleBody,
) error {
	if value == nil {
		value = new(m_dto.MemberAddRoleBody)
	}
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		SetPathParam("user_id", userID).
		SetBody(value).
		Put(o.getURL(memberRoleURI))
	return err
}

// MemberDeleteRole 删除成员角色
func (o *openAPI) MemberDeleteRole(
	ctx context.Context, guildID string, roleID m_dto.RoleID, userID string,
	value *m_dto.MemberAddRoleBody,
) error {
	if value == nil {
		value = new(m_dto.MemberAddRoleBody)
	}
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		SetPathParam("user_id", userID).
		SetBody(value).
		Delete(o.getURL(memberRoleURI))
	return err
}

// GuildMember 拉取频道指定成员
func (o *openAPI) GuildMember(ctx context.Context, guildID, userID string) (*m_dto.Member, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.Member{}).
		SetPathParam("guild_id", guildID).
		SetPathParam("user_id", userID).
		Get(o.getURL(guildMemberURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.Member), nil
}

// GuildMembers 分页拉取频道内成员列表
func (o *openAPI) GuildMembers(
	ctx context.Context,
	guildID string, pager *m_dto.GuildMembersPager,
) ([]*m_dto.Member, error) {
	if pager == nil {
		return nil, errs.ErrPagerIsNil
	}
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetQueryParams(pager.QueryParams()).
		Get(o.getURL(guildMembersURI))
	if err != nil {
		return nil, err
	}

	members := make([]*m_dto.Member, 0)
	if err := json.Unmarshal(resp.Body(), &members); err != nil {
		return nil, err
	}

	return members, nil
}

// DeleteGuildMember 将指定成员踢出频道
func (o *openAPI) DeleteGuildMember(ctx context.Context, guildID, userID string, opts ...m_dto.MemberDeleteOption) error {
	opt := &m_dto.MemberDeleteOpts{}
	for _, o := range opts {
		o(opt)
	}
	_, err := o.request(ctx).
		SetResult(m_dto.Member{}).
		SetPathParam("guild_id", guildID).
		SetPathParam("user_id", userID).
		SetBody(opt).
		Delete(o.getURL(guildMemberURI))
	return err
}
