package services

type BannerTagService struct {
	repo IBannerTagRepository
}

func NewBannerTagService(repository IBannerTagRepository) *BannerTagService {
	return &BannerTagService{repo: repository}
}
