package api


// _LoginResponse 登陆接口响应数据
type _LoginResponse struct {
	Code    int64                   `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    _JWTData `json:"data"`    // 数据
}


type _JWTData struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}