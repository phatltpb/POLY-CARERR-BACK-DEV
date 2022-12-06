package uploadRepository

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/tuongnguyen1209/poly-career-back/config"
)

type UploadRepository struct {
}

func Init() *UploadRepository {
	return &UploadRepository{}
}

func (up UploadRepository) UploadImage(file *multipart.FileHeader, folder string) (string, error) {
	// Source
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	var (
		config      = config.GetConfig()
		CloudName   = config.Cloud.Name
		CloudApiKey = config.Cloud.ApiKey
		CloudSecret = config.Cloud.SecretKey
	)
	// defer src.Close()
	cld, _ := cloudinary.NewFromParams(CloudName, CloudApiKey, CloudSecret)
	// cld.Config.URL.Secure = true
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	println(file.Header)
	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		Folder: folder,
	})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
