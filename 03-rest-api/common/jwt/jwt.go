package jwt

import (
	"log"
)

type JwtService interface {
	verify(jwt string) (map[string]interface{}, error)
}

type jwtService struct {
	secret string
	logger *log.Logger
}

func (s *jwtService) verify(jwt string) (map[string]interface{}, error) {
	s.logger.Println("JWT verification done")
	return map[string]interface{}{}, nil
}

func NewJwtService(secret string, logger *log.Logger) JwtService {
	return &jwtService{
		secret: secret,
		logger: logger,
	}
}
