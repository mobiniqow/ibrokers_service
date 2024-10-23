package manager

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
	"ibrokers_service/pkg/configs"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type FileManager struct {
	Path        string
	MinioClient configs.Minio
	PathUtils   PathUtils
}

func NewFileManager(path string, config configs.Minio) *FileManager {
	return &FileManager{
		Path:        path,
		MinioClient: config,
		PathUtils:   PathUtils{Path: path},
	}
}

func (c *FileManager) IsExist(file *multipart.FileHeader) bool {
	var path = c.PathUtils.AbsolutePath(file.Filename)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
	}
	return false
}

func (c *FileManager) SaveFile(bucketName string, file *multipart.FileHeader) (string, error) {

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	d, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	extensions := strings.Split(file.Filename, ".")
	filename := fmt.Sprintf("%s.%s", d, extensions[len(extensions)-1])
	fileContent, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	err = c.MinioClient.UploadFile(bucketName, filename, bytes.NewReader(fileContent),
		int64(len(fileContent)), minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {

		return "", err
	}

	return fmt.Sprintf("%s/%s", bucketName, filename), nil
}

func (c *FileManager) DeleteFile(fileName string) error {
	return os.Remove(c.PathUtils.AbsolutePath(fileName))
}
