package manager

// 会话相关的抽象接口
type (
	// Session 会话
	Session interface {
		Key() string
		Token() string
		Get(key string) (value string, err error)
		Set(key, value string) error
	}

	// SessionManager 会话管理者
	SessionManager interface {
		CreateSession(key string) (Session, error)
		GetSession(key string) (Session, error)
		DelSession(key string) error
		ExpireSession(key string, ttl int) error
	}

	TokenManager interface {
		NewToken() (string, error)
	}
)
