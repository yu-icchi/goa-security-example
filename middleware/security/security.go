package security

import (
	"github.com/goadesign/goa"
)

// Security 認証後に処理したいミドルウェアを管理する
type Security struct {
	main       goa.Middleware
	middleware []goa.Middleware
}

// Use 認証後に処理したいミドルウェアを設定する
func (sec *Security) Use(m goa.Middleware) {
	sec.middleware = append(sec.middleware, m)
}

// Middleware 追加されたミドルウェアも含めてセキュリティミドルウェアとして使用する
func (sec *Security) Middleware() goa.Middleware {
	return func(nextHandler goa.Handler) goa.Handler {
		ml := len(sec.middleware)
		for i := range sec.middleware {
			nextHandler = sec.middleware[ml-i-1](nextHandler)
		}
		return sec.main(nextHandler)
	}
}

// New Security middleware
func New(security goa.Middleware) *Security {
	return &Security{
		main:       security,
		middleware: []goa.Middleware{},
	}
}
