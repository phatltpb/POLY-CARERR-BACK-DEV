package dto

import (
	"time"

	"github.com/tuongnguyen1209/poly-career-back/apis/models"
)

type ApplyJobWithNewCv struct {
	PostId   int    `json:"job_id"`
	Letter   string `json:"letter"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
}

func (dto *ApplyJobWithNewCv) ConvertToApplyJob(cvId int) *models.ApplyJob {
	return &models.ApplyJob{
		PostId:    dto.PostId,
		Letter:    dto.Letter,
		CvId:      cvId,
		DateApply: time.Now(),
	}
}
