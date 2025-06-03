package jwt

import (
	"log"
)

type JwtService interface {
	Verify(jwt string) (map[string]string, error)
}

type jwtService struct {
	secret string
	logger *log.Logger
}

func NewJwtService(secret string, logger *log.Logger) JwtService {
	return &jwtService{
		secret: secret,
		logger: logger,
	}
}

func (s *jwtService) Verify(jwt string) (map[string]string, error) {
	s.logger.Println("JWT verification")
	return map[string]string{}, nil
}
