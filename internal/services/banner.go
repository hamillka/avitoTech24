package services

import (
	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	serviceModels "github.com/hamillka/avitoTech24/internal/services/models"
)

type IBannerRepository interface {
	GetBannersByFeature(featureId, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error)
	GetBannersByTag(tagId, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error)
	GetBannerByFeatureAndTag(featureId, tagId, limit, offset int64) (*repoModels.Banner, error)
}

type BannerService struct {
	repo IBannerRepository
}

func NewBannerService(repository IBannerRepository) *BannerService {
	return &BannerService{repo: repository}
}

func (bs *BannerService) GetBannersByFeature(featureID, limit, offset int64) ([]*serviceModels.BannerWithTagIDs, error) {
	banners, tagIds, err := bs.repo.GetBannersByFeature(featureID, limit, offset)
	if err != nil {
		//
	}

	bannersWithTags := make([]*serviceModels.BannerWithTagIDs, 0)

	for _, banner := range banners {
		bannersWithTags = append(bannersWithTags, serviceModels.ConvertToBL(*banner, tagIds))

	}

	return bannersWithTags, nil
}

func (bs *BannerService) GetBannersByTag(tagID, limit, offset int64) ([]*serviceModels.BannerWithTagIDs, error) {
	return nil, nil
}

func (bs *BannerService) GetBannersByFeatureAndTag(featureID, tagID, limit, offset int64) ([]*serviceModels.BannerWithTagIDs, error) {
	return nil, nil
}
