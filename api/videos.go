package api

import (
	"net/http"
	"strconv"

	"github.com/Ayush-Walia/Fampay-Youtube/models"
	"github.com/Ayush-Walia/Fampay-Youtube/storage"
	"github.com/Ayush-Walia/Fampay-Youtube/utils"
	"github.com/gookit/slog"
)

func videos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageNo, err := strconv.Atoi(r.URL.Query().Get("pageNo"))
	if err != nil {
		http.Error(w, "Invalid 'pageNo' parameter", http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		http.Error(w, "Invalid 'pageSize' parameter", http.StatusBadRequest)
		return
	}

	// Get videos from database
	videosDAO, err := storage.VideosDao.GetVideosByPublishTime(ctx, pageNo, pageSize)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Database query error!", http.StatusInternalServerError)
		return
	}

	var videosDTO []models.VideoDTO
	for _, videoDAO := range videosDAO {
		var videoDTO models.VideoDTO
		err = utils.CopyProperties(videoDAO, &videoDTO)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Error creating response!", http.StatusInternalServerError)
			return
		}

		videosDTO = append(videosDTO, videoDTO)
	}

	if len(videosDTO) == 0 {
		utils.RespondWithString(w, http.StatusNotFound, "No videos found!")
	} else {
		utils.RespondWithJSON(w, videosDTO)
	}
}
