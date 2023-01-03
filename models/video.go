package models

import "time"

type VideoDTO struct {
	VideoID      string    `json:"video_id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	ThumbnailURL string    `json:"thumbnail_url"`
	PublishTime  time.Time `json:"publish_time"`
}
