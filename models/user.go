package models

import "time"

type User struct {
    Id             int64     `json:"id"`
    GoogleId       string    `json:"google_id"`
    CreatedAt      time.Time `json:"created_at"`
    DisplayName    string    `json:"display_name"`
    ProfilePicture string    `json:"profile_picture"`
    Token          string    `json:"token,omitempty"`
}
