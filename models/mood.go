package models

import "time"

type Mood struct {
    Id        int64            `json:"id"`
    Type      int64            `json:"type"`
    Title     string           `json:"title,omitempty"`
    CreatedAt time.Time        `json:"created_at"`
    Owner     *User            `json:"owner"`
}

/*
 * Mood Types
 *   0 - Bad
 *   1 - Good
 * Could it be a bool instead? yes.
 * We might add more moods soon, who knows?
 */
