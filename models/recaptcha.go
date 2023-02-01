package models

import "time"

type Recaptcha struct {
	Success     bool      `json:"success"`
	ErrorCodes  []string  `json:"error-codes,omitempty"`
	ChallengeTs time.Time `json:"challenge_ts,omitempty"`
	Hostname    string    `json:"hostname,omitempty"`
}
