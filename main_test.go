package main

import (
	"testing"
	"time"
)

func TestLimiter_Allow(t *testing.T) {
	tests := []struct {
		name     string
		limit    int
		window   time.Duration
		requests []struct {
			userID   string
			delay    time.Duration //  delay before making this request
			expected bool
		}
		description string
	}{
		{
			name:   "allow requests within limit",
			limit:  3,
			window: time.Second,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, true},
				{"user1", 0, true},
				{"user1", 0, true},
			},
			description: "should allow all requests when under the limit",
		},
		{
			name:   "reject requests exceeding limit",
			limit:  2,
			window: time.Second,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, true},
				{"user1", 0, true},
				{"user1", 0, false}, // exceeds limit
				{"user1", 0, false}, // still exceeds limit
			},
			description: "should reject requests when limit is exceeded",
		},
		{
			name:   "allow requests after window expires",
			limit:  2,
			window: 100 * time.Millisecond,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, true},
				{"user1", 0, true},
				{"user1", 0, false},                     // exceeds limit
				{"user1", 150 * time.Millisecond, true}, // window expired, should allow
				{"user1", 0, true},                      // still within new window
			},
			description: "should allow requests after time window expires",
		},
		{
			name:   "multiple users independent limits",
			limit:  2,
			window: time.Second,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, true},
				{"user2", 0, true},
				{"user1", 0, true},
				{"user2", 0, true},
				{"user1", 0, false}, // user1 exceeds limit
				{"user2", 0, false}, // user2 exceeds limit
				{"user3", 0, true},  // user3 first request
			},
			description: "should maintain independent rate limits per user",
		},
		{
			name:   "zero limit should always reject",
			limit:  0,
			window: time.Second,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, false},
				{"user1", 0, false},
			},
			description: "should always reject when limit is zero",
		},
		{
			name:   "single request limit",
			limit:  1,
			window: time.Second,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, true},
				{"user1", 0, false},
				{"user1", 0, false},
			},
			description: "should work correctly with limit of 1",
		},
		{
			name:   "partial window expiry",
			limit:  3,
			window: 100 * time.Millisecond,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, true},                     // t=0, count=1
				{"user1", 20 * time.Millisecond, true}, // t=20ms, count=2
				{"user1", 0, true},                     // t=20ms, count=3
				{"user1", 0, false},                    // t=20ms, exceeds limit (3)
				{"user1", 90 * time.Millisecond, true}, // t=110ms, first request (t=0) expired, count=2+1=3
				{"user1", 0, false},                    // t=110ms, still at limit
				{"user1", 30 * time.Millisecond, true}, // t=140ms, second request (t=20) expired, count=2+1=3
			},
			description: "should handle partial window expiry correctly",
		},
		{
			name:   "very short window",
			limit:  2,
			window: time.Nanosecond,
			requests: []struct {
				userID   string
				delay    time.Duration
				expected bool
			}{
				{"user1", 0, true},
				{"user1", time.Microsecond, true}, // previous request should have expired
				{"user1", time.Microsecond, true}, // previous request should have expired
			},
			description: "should work with very short time windows",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := New(tt.limit, tt.window)

			for i, req := range tt.requests {
				if req.delay > 0 {
					time.Sleep(req.delay)
				}

				result := rl.Allow(req.userID)
				if result != req.expected {
					t.Errorf("Request %d: Allow(%q) = %v, want %v. %s",
						i+1, req.userID, result, req.expected, tt.description)
				}
			}
		})
	}
}
