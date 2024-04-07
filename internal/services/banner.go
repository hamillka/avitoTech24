package services

import (
	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	serviceModels "github.com/hamillka/avitoTech24/internal/services/models"
)

type IBannerRepository interface {
	GetBannersByFeature(featureID, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error)
	GetBannersByTag(tagID, limit, offset int64) ([]*repoModels.Banner, map[int64][]int64, error)
	GetBannerByFeatureAndTag(featureID, tagID, limit, offset int64) (*repoModels.Banner, error)
}

type BannerService struct {
	repo IBannerRepository
}

func NewBannerService(repository IBannerRepository) *BannerService {
	return &BannerService{repo: repository}
}

func (bs *BannerService) GetBannersByFeature(featureID, limit, offset int64) ([]*serviceModels.BannerWithTagIDs, error) {
	banners, tagIDs, err := bs.repo.GetBannersByFeature(featureID, limit, offset)
	if err != nil {
		return nil, err
	}

	bannersWithTags := make([]*serviceModels.BannerWithTagIDs, 0)

	for _, banner := range banners {
		bannersWithTags = append(bannersWithTags, serviceModels.ConvertToBL(*banner, tagIDs))
	}

	return bannersWithTags, nil
}

func (bs *BannerService) GetBannersByTag(tagID, limit, offset int64) ([]*serviceModels.BannerWithTagIDs, error) {
	banners, tagIDs, err := bs.repo.GetBannersByTag(tagID, limit, offset)
	if err != nil {
		return nil, err
	}

	bannersWithTags := make([]*serviceModels.BannerWithTagIDs, 0)

	for _, banner := range banners {
		bannersWithTags = append(bannersWithTags, serviceModels.ConvertToBL(*banner, tagIDs))
	}

	return bannersWithTags, nil
}

func (bs *BannerService) GetBannersByFeatureAndTag(featureID, tagID, limit, offset int64) ([]*serviceModels.BannerWithTagIDs, error) {
	banner, err := bs.repo.GetBannerByFeatureAndTag(featureID, tagID, limit, offset)
	if err != nil {
		return nil, err
	}

	bannersWithTags := make([]*serviceModels.BannerWithTagIDs, 1)

	bannersWithTags[0] = serviceModels.ConvertToBL(*banner, map[int64][]int64{
		banner.BannerID: {tagID},
	})

	return bannersWithTags, nil
}
