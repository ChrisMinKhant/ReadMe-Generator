package generatereadmefileservice

import dto "github.com/ChrisMinKhant/megoyougo_framework/dto/request"

type GenerateReadmeFileService interface {
	GenerateReadmeFile(*dto.GenerateReadmeFileRequest) bool
}
