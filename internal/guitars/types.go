package guitars

import "time"

type Guitar struct {
    ID              int64     `json:"id"`
    Brand           string    `json:"brand"`
    Model           string    `json:"model"`
    Year            *int      `json:"year,omitempty"`
    Notes           *string   `json:"notes,omitempty"`

    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
