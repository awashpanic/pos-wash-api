package aws

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"

	aws_config "github.com/aws/aws-sdk-go-v2/config"
	aws_manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	aws_s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ffajarpratama/pos-wash-api/config"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
)

type IFaceAWS interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (*UploadResponse, error)
}

type AWS struct {
	cnf      config.AWSConfig
	uploader *aws_manager.Uploader
}

type UploadResponse struct {
	Size     int
	Mimetype string
	Name     string
	Location string
}

func NewAWSClient(cnf config.AWSConfig) (IFaceAWS, error) {
	cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion(cnf.Region))
	if err != nil {
		return nil, err
	}

	client := aws_s3.NewFromConfig(cfg)
	uploader := aws_manager.NewUploader(client)
	return &AWS{cnf: cnf, uploader: uploader}, nil
}

// UploadFile implements IFaceAWS.
func (a *AWS) UploadFile(ctx context.Context, file *multipart.FileHeader) (*UploadResponse, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	ext := filepath.Ext(file.Filename)
	n := strings.LastIndexByte(file.Filename, '.')

	filename := util.RemoveSpecialCharacters(file.Filename[:n])
	filename = filename + "-" + strconv.FormatInt(util.TimeNow().Unix(), 10) + ext
	mimetype := file.Header.Get("Content-Type")

	fileObject := &aws_s3.PutObjectInput{
		Key:          aws.String(filename),
		Bucket:       aws.String(a.cnf.Bucket),
		Body:         f,
		ContentType:  aws.String(mimetype),
		CacheControl: aws.String(fmt.Sprintf("max-age=%d", constant.FILE_UPLOAD_MAX_AGE)),
	}

	result, err := a.uploader.Upload(context.TODO(), fileObject)
	if err != nil {
		return nil, err
	}

	res := &UploadResponse{
		Size:     int(file.Size),
		Mimetype: mimetype,
		Name:     filename,
		Location: result.Location,
	}

	return res, nil
}
