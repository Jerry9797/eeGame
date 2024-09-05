package consts

const (
	// 登录相关
	TokenType          = "Bearer"
	CacheModeRedis     = 2
	BackendServerName  = "开源电商系统"
	MultiLogin         = true
	FrontendMultiLogin = true
	GTokenExpireIn     = 10 * 24 * 60 * 60

	//统一管理错误提示
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginFaulMsg         = "登录失败，账号或密码错误"
	ErrSecretAnswerMsg      = "密保问题不正确"
	ResourcePermissionFail  = "没有权限操作"

	//for user
	CtxUserId         = "CtxUserId"
	CtxUserName       = "CtxUserName"
	CtxUserAvatar     = "CtxUserAvatar"
	CtxAddress        = "CtxAddress"
	CtxUserSex        = "CtxUserSex"
	CtxUserSign       = "CtxUserSign"
	CtxUserStatus     = "CtxUserStatus"
	GTokenAdminPrefix = "EeGame:"
)
