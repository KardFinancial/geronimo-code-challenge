package main

import "time"

type Limiter struct {
}

func New(limit int, window time.Duration) *Limiter {
	return &Limiter{}
}

// Allow should return true if the request is allowed, or false if it's rate limited
func (l *Limiter) Allow(userID string) bool {
	return false
}
