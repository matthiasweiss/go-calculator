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
	s.logger.Println("JWT service verifies token")

	// return map[string]string{}, fmt.Errorf("Error verifying JWT")
	return map[string]string{}, nil
}
