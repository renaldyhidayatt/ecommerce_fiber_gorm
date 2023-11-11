package utilscloudinary

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/spf13/viper"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cldSecret := viper.GetString("CLOUDINARY_SECRET_KEY")
	cldName := viper.GetString("CLOUDINARY_CLOUD_NAME")
	cldKey := viper.GetString("CLOUDINARY_API_KEY")
	cldfolder := viper.GetString("CLOUDINARY_UPLOAD_FOLDER")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: cldfolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
