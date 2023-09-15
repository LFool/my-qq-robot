package v1

import (
	"context"
	m_dto "demo/m-dto"
	"fmt"
)

func (o *openAPI) Roles(ctx context.Context, guildID string) (*m_dto.GuildRoles, error) {
	resp, err := o.request(ctx).
		SetResult(m_dto.GuildRoles{}).
		SetPathParam("guild_id", guildID).
		Get(o.getURL(rolesURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.GuildRoles), nil
}

func (o *openAPI) PostRole(ctx context.Context, guildID string, role *m_dto.Role) (*m_dto.UpdateResult, error) {
	if role.Color == 0 {
		role.Color = m_dto.DefaultColor
	}
	// openapi 上修改哪个字段，就需要传递哪个 filter
	filter := &m_dto.UpdateRoleFilter{
		Name:  1,
		Color: 1,
		Hoist: 1,
	}
	body := &m_dto.UpdateRole{
		GuildID: guildID,
		Filter:  filter,
		Update:  role,
	}
	fmt.Sprint(body)
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetResult(m_dto.UpdateResult{}).
		SetBody(body).
		Post(o.getURL(rolesURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.UpdateResult), nil
}

func (o *openAPI) PatchRole(ctx context.Context,
	guildID string, roleID m_dto.RoleID, role *m_dto.Role) (*m_dto.UpdateResult, error) {
	if role.Color == 0 {
		role.Color = m_dto.DefaultColor
	}
	filter := &m_dto.UpdateRoleFilter{
		Name:  1,
		Color: 1,
		Hoist: 1,
	}
	body := &m_dto.UpdateRole{
		GuildID: guildID,
		Filter:  filter,
		Update:  role,
	}
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		SetResult(m_dto.UpdateResult{}).
		SetBody(body).
		Patch(o.getURL(roleURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*m_dto.UpdateResult), nil
}

func (o *openAPI) DeleteRole(ctx context.Context, guildID string, roleID m_dto.RoleID) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		Delete(o.getURL(roleURI))
	return err
}
