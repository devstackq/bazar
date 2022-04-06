package models

import "time"

type AccessDetails struct {
	AccessUuid string
	UserId     int64
}

type TokenDetails struct {
	AccessToken    string
	RefreshToken   string
	AccessUuid     string
	RefreshUuid    string
	AtExpires      int64
	RtExpires      int64
	SubTimeRefresh time.Duration
	SubTimeAccess  time.Duration
	UserID         int
}
