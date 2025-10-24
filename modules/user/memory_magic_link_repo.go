package user

import (
	"errors"
	"sync"
	"time"
)

type memoryMagicLinkRepo struct {
	mu            sync.RWMutex
	magicLinks    map[string]magicLinkData
	verifications map[string]bool
}

type magicLinkData struct {
	email     string
	code      string
	createdAt time.Time
	expiresAt time.Time
}

func newMemoryMagicLinkRepo() magicLinkRepo {
	return &memoryMagicLinkRepo{
		magicLinks:    make(map[string]magicLinkData),
		verifications: make(map[string]bool),
	}
}

func (r *memoryMagicLinkRepo) sendMagicLink(email string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	code := "123456"
	now := time.Now()
	expiresAt := now.Add(15 * time.Minute)

	r.magicLinks[email] = magicLinkData{
		email:     email,
		code:      code,
		createdAt: now,
		expiresAt: expiresAt,
	}

	return nil
}

func (r *memoryMagicLinkRepo) verifyMagicLink(params VerifyMagicLinkParamsDto) (AuthTokensDto, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data, exists := r.magicLinks[params.Email]
	if !exists {
		return AuthTokensDto{}, errors.New("magic link not found")
	}

	if data.code != params.Code {
		return AuthTokensDto{}, errors.New("invalid code")
	}

	if time.Now().After(data.expiresAt) {
		return AuthTokensDto{}, errors.New("magic link expired")
	}

	r.verifications[params.Email] = true

	return AuthTokensDto{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil
}
