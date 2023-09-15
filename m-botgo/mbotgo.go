package m_botgo

import (
	m_openapi "demo/m-openapi"
	v1 "demo/m-openapi/v1"
	"demo/m-token"
	"demo/websocket/client"
)

func init() {
	v1.Setup()     // 注册 v1 接口
	client.Setup() // 注册 websocket client 实现
}

// NewSessionManager 获得 session manager 实例
func NewSessionManager() SessionManager {
	return defaultSessionManager
}

// NewOpenAPI 创建新的 openapi 实例，会返回当前的 openapi 实现的实例
// 如果需要使用其他版本的实现，需要在调用这个方法之前调用 SelectOpenAPIVersion 方法
func NewOpenAPI(token *m_token.MToken) m_openapi.OpenAPI {
	return m_openapi.DefaultImpl.Setup(token, false)
}
