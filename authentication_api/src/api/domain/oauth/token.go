package oauth

import "time"

type Token struct {
	Token   string `json:"token"`
	UserID  int64  `json:"user_id"`
	Expires int64  `json:"expires"`
}

func (t *Token) IsExpired() bool {
	return time.Unix(t.Expires, 0).UTC().Before(time.Now().UTC())
}
