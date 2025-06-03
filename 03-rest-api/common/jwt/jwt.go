package jwt

import (
	"log"
)

type JwtService interface {
	verify(jwt string) (map[string]string{}, error)
}

type jwtService struct {
	secret string
	logger *log.Logger
}

func (s *jwtService) verify(jwt string) (map[string]string{}, error) {
	s.logger.Println("JWT verification")
	return map[string]string{}{}, nil
}

func NewJwtService(secret string, logger *log.Logger) JwtService {
	return &jwtService{
		secret: secret,
		logger: logger,
	}
}
