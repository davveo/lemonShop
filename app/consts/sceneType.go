package consts

type SceneType string

var (
	//验证码登录
	LOGIN SceneType = "验证码登录"
	//手机注册
	REGISTER SceneType = "手机注册"
	//找回密码
	FIND_PASSWORD SceneType = "找回密码"
	//绑定手机
	BIND_MOBILE SceneType = "绑定手机"
	//修改密码
	MODIFY_PASSWORD SceneType = "修改密码"
	//添加店员
	ADD_CLERK SceneType = "添加店员"
	//手机验证
	VALIDATE_MOBILE SceneType = "验证手机"
	//绑定邮箱
	BIND_EMAIL SceneType = "绑定邮箱"
)
