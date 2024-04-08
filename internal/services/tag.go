package services

type TagService struct {
	repo ITagRepository
}

func NewTagService(repository ITagRepository) *TagService {
	return &TagService{repo: repository}
}
