package m_token

import "fmt"

const (
	TypeBot    string = "Bot"
	TypeNormal string = "Bearer"
)

type MToken struct {
	AppID       uint64
	AccessToken string
	Type        string
}

func BotToken(appID uint64, accessToken string) *MToken {
	return &MToken{
		AppID:       appID,
		AccessToken: accessToken,
		Type:        TypeBot,
	}
}

func (t *MToken) GetString() string {
	if t.Type == TypeNormal {
		return t.AccessToken
	}
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}
