package manager

import uuid "github.com/satori/go.uuid"

// UUIDTokenManager 生成uuid的token管理者
type UUIDTokenManager struct{}

// NewToken 生成token
func (*UUIDTokenManager) NewToken() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
