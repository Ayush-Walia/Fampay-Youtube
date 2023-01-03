package storage

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/gookit/slog"
)

type VideosDAOImpl struct {
	DB *sql.DB
}

type Video struct {
	ID           int64     `db:"id"`
	VideoID      string    `db:"video_id"`
	Title        string    `db:"title"`
	Description  *string   `db:"description"`
	ThumbnailURL string    `db:"thumbnail_url"`
	PublishTime  time.Time `db:"publish_time"`
}

var VideosDao *VideosDAOImpl

func newVideosDAO(DB *sql.DB) *VideosDAOImpl {
	videos := new(VideosDAOImpl)
	videos.DB = DB
	return videos
}

func (v *VideosDAOImpl) GetVideosByPublishTime(ctx context.Context, pageNo int, pageSize int) ([]Video, error) {
	offset := (pageNo - 1) * pageSize
	sqlQuery := `SELECT id, video_id, title, description, thumbnail_url, publish_time FROM videos ORDER BY publish_time DESC LIMIT ? OFFSET ?`

	rows, err := v.DB.QueryContext(ctx, sqlQuery, pageSize, offset)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			slog.Error(err)
		}
	}(rows)

	var videos []Video
	for rows.Next() {
		var video Video
		err = rows.Scan(&video.ID, &video.VideoID, &video.Title, &video.Description, &video.ThumbnailURL, &video.PublishTime)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (v *VideosDAOImpl) SearchVideos(ctx context.Context, query string, pageNo int, pageSize int) ([]Video, error) {
	offset := (pageNo - 1) * pageSize
	sqlQuery := `SELECT id, video_id, title, description, thumbnail_url, publish_time FROM videos WHERE MATCH(title, description) AGAINST(? IN BOOLEAN MODE) LIMIT ? OFFSET ?`

	keywords := strings.ReplaceAll("+"+query, " ", " +")
	rows, err := v.DB.QueryContext(ctx, sqlQuery, keywords, pageSize, offset)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			slog.Error(err)
		}
	}(rows)

	var videos []Video
	for rows.Next() {
		var video Video
		err = rows.Scan(&video.ID, &video.VideoID, &video.Title, &video.Description, &video.ThumbnailURL, &video.PublishTime)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (v *VideosDAOImpl) SaveVideos(ctx context.Context, videos []Video) error {
	sqlQuery := "INSERT IGNORE INTO videos (video_id, title, description, thumbnail_url, publish_time) VALUES (?, ?, ?, ?, ?)"
	stmt, err := v.DB.Prepare(sqlQuery)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			slog.Error(err)
		}
	}(stmt)

	for _, video := range videos {
		_, err = stmt.ExecContext(ctx, video.VideoID, video.Title, video.Description, video.ThumbnailURL, video.PublishTime)
		if err != nil {
			return err
		}
	}

	return nil
}
