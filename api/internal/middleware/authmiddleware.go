package middleware

import "net/http"

// AuthMiddleware 中间件
type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//开始自己的中间件服务开发
		//获取redis 或者数据库中的数据进行 数据验证

		next(w, r)
	}
}
