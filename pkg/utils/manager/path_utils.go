package manager

import (
	"fmt"
	"ibrokers_service/pkg/configs"
)

type PathUtils struct {
	Path string
}

func (c *PathUtils) AbsolutePath(fileName string) string {
	return fmt.Sprintf("%s%s%s/%s", configs.BASE_URL, configs.MEDIA_ROOT, c.Path, fileName)
}

func (c *PathUtils) RelativePath(fileName string) string {
	return fmt.Sprintf("%s%s/%s", configs.MEDIA_ROOT, c.Path, fileName)
}
