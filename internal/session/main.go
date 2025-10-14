package session

import "github.com/google/uuid"

type SessionStorage struct {
	Token uuid.UUID
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{}
}

// Replace old token with new one
func (s *SessionStorage) NewToken() {
	id := uuid.New()

	s.Token = id
}

func (s *SessionStorage) RevokeToken() {
	s.Token = uuid.Nil
}
