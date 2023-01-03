package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Ayush-Walia/Fampay-Youtube/config"
	"github.com/Ayush-Walia/Fampay-Youtube/storage"
	"github.com/gookit/slog"
	"github.com/robfig/cron/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeService struct {
	APIKeys map[string]bool
	query   string
}

var ErrQuotaExpired = errors.New("Quota Exceeded")

func NewYoutubeService() *YoutubeService {
	y := new(YoutubeService)
	y.APIKeys = make(map[string]bool)
	return y
}

func (y *YoutubeService) Init(conf *config.AppConfig) {
	for _, key := range strings.Split(conf.APIKeys, ",") {
		y.APIKeys[key] = true
	}
	y.query = conf.YoutubeQuery

	y.registerCron(conf.YoutubeCron)
	y.syncYoutube()
}

func (y *YoutubeService) registerCron(cronStr string) {
	c := cron.New()
	_, err := c.AddFunc(cronStr, y.syncYoutube)
	if err != nil {
		slog.Error(err)
	}

	c.Start()
}

// searchYoutube searches YouTube for videos matching the given query.
func (y *YoutubeService) searchYoutube() (*youtube.SearchListResponse, error) {
	ctx := context.Background()
	APIKey, err := y.getValidAPIKey()
	if err != nil {
		return nil, err
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(APIKey))
	if err != nil {
		return nil, fmt.Errorf("youtube.NewService: %v", err)
	}

	call := service.Search.List([]string{"id,snippet"}).
		Q(y.query).
		Type("video").
		Order("date").
		PublishedAfter(time.Now().AddDate(-1, 0, 0).Format(time.RFC3339)).
		MaxResults(100)
	response, err := call.Do()

	// If API quota has expired mark it as false
	if response.HTTPStatusCode == http.StatusForbidden {
		slog.Error(ErrQuotaExpired)
		y.APIKeys[APIKey] = false
		return nil, ErrQuotaExpired
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (y *YoutubeService) syncYoutube() {
	slog.Info("SyncYoutube Cron triggered")
	resp, err := y.searchYoutube()

	// Keep retrying for other API keys
	for {
		if errors.Is(err, ErrQuotaExpired) {
			resp, err = y.searchYoutube()
		} else {
			break
		}
	}

	if err != nil {
		slog.Error(err)
		return
	}

	var videos []storage.Video
	for _, item := range resp.Items {
		snippet := item.Snippet

		// Youtube API uses time in RFC3339 format.
		publishTime, err := time.Parse(time.RFC3339, snippet.PublishedAt)
		if err != nil {
			slog.Error(err)
			return
		}

		video := storage.Video{
			VideoID:      item.Id.VideoId,
			Title:        snippet.Title,
			Description:  &snippet.Description,
			ThumbnailURL: snippet.Thumbnails.High.Url,
			PublishTime:  publishTime,
		}
		videos = append(videos, video)
	}

	// Save videos to database
	err = storage.VideosDao.SaveVideos(context.Background(), videos)
	if err != nil {
		slog.Error(err)
	}

	slog.Info("SyncYoutube Cron finished")
}

// returns first API key whose quota has not expired.
func (y *YoutubeService) getValidAPIKey() (string, error) {
	for key, val := range y.APIKeys {
		if val == true {
			return key, nil
		}
	}

	return "", errors.New("all keys quota expired")
}
