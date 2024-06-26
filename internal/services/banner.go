package services

import (
	repoModels "github.com/hamillka/avitoTech24/internal/repositories/models"
	serviceModels "github.com/hamillka/avitoTech24/internal/services/models"
)

type BannerRepository interface {
	GetBannersByFeature(featureID, limit, offset, role int64) ([]*repoModels.Banner, map[int64][]int64, error)
	GetBannersByTag(tagID, limit, offset, role int64) ([]*repoModels.Banner, map[int64][]int64, error)
	GetBannerByFeatureAndTag(featureID, tagID, role int64) (*repoModels.Banner, error)
	CreateBanner(featureID int64, content string, isActive bool) (int64, error)
	UpdateBanner(bannerID, featureID int64, content string, isActive *bool) (int64, error)
	DeleteBanner(bannerID int64) error
}

type FeatureRepository interface {
	GetFeatureByID(featureID int64) (*repoModels.Feature, error)
}

type TagRepository interface {
	GetTagByID(tagID int64) (*repoModels.Tag, error)
}

type BannerTagRepository interface {
	CreateBannerTag(bannerID, tagID int64) error
	DeleteRecordsByBannerID(bannerID int64) error
}

type BannerService struct {
	bannerRepo    BannerRepository
	featureRepo   FeatureRepository
	tagRepo       TagRepository
	bannerTagRepo BannerTagRepository
}

func NewBannerService(
	bannerRepository BannerRepository,
	featureRepository FeatureRepository,
	tagRepository TagRepository,
	bannerTagRepository BannerTagRepository,
) *BannerService {
	return &BannerService{
		bannerRepo:    bannerRepository,
		featureRepo:   featureRepository,
		tagRepo:       tagRepository,
		bannerTagRepo: bannerTagRepository,
	}
}

func (bs *BannerService) GetBannersByFeature(featureID, limit, offset, role int64) ([]*serviceModels.BannerWithTagIDs, error) {
	banners, tagIDs, err := bs.bannerRepo.GetBannersByFeature(featureID, limit, offset, role)
	if err != nil {
		return nil, err
	}

	bannersWithTags := make([]*serviceModels.BannerWithTagIDs, 0)

	for _, banner := range banners {
		bannersWithTags = append(bannersWithTags, serviceModels.ConvertToBL(banner, tagIDs))
	}

	return bannersWithTags, nil
}

func (bs *BannerService) GetBannersByTag(tagID, limit, offset, role int64) ([]*serviceModels.BannerWithTagIDs, error) {
	banners, tagIDs, err := bs.bannerRepo.GetBannersByTag(tagID, limit, offset, role)
	if err != nil {
		return nil, err
	}

	bannersWithTags := make([]*serviceModels.BannerWithTagIDs, 0)

	for _, banner := range banners {
		bannersWithTags = append(bannersWithTags, serviceModels.ConvertToBL(banner, tagIDs))
	}

	return bannersWithTags, nil
}

func (bs *BannerService) GetBannersByFeatureAndTag(featureID, tagID, role int64) ([]*serviceModels.BannerWithTagIDs, error) {
	banner, err := bs.bannerRepo.GetBannerByFeatureAndTag(featureID, tagID, role)
	if err != nil {
		return nil, err
	}

	bannersWithTags := make([]*serviceModels.BannerWithTagIDs, 1)

	bannersWithTags[0] = serviceModels.ConvertToBL(banner, map[int64][]int64{
		banner.BannerID: {tagID},
	})

	return bannersWithTags, nil
}

func (bs *BannerService) CreateBanner(tagIDs []int64, featureID int64, content string, isActive bool) (int64, error) {
	if featureID != 0 {
		_, err := bs.featureRepo.GetFeatureByID(featureID)
		if err != nil {
			return 0, err
		}
	}

	for _, tagID := range tagIDs {
		_, err := bs.tagRepo.GetTagByID(tagID)
		if err != nil {
			return 0, err
		}
	}

	bannerID, err := bs.bannerRepo.CreateBanner(featureID, content, isActive)
	if err != nil {
		return 0, err
	}

	for _, tagID := range tagIDs {
		err := bs.bannerTagRepo.CreateBannerTag(bannerID, tagID)
		if err != nil {
			return 0, err
		}
	}

	return bannerID, nil
}

func (bs *BannerService) UpdateBanner(bannerID int64, tagIDs []int64, featureID int64, content string, isActive *bool) (int64, error) {
	if featureID != 0 {
		_, err := bs.featureRepo.GetFeatureByID(featureID)
		if err != nil {
			return 0, err
		}
	}

	for _, tagID := range tagIDs {
		_, err := bs.tagRepo.GetTagByID(tagID)
		if err != nil {
			return 0, err
		}
	}

	id, err := bs.bannerRepo.UpdateBanner(bannerID, featureID, content, isActive)
	if err != nil {
		return 0, err
	}

	err = bs.bannerTagRepo.DeleteRecordsByBannerID(bannerID)
	if err != nil {
		return 0, err
	}

	for _, tagID := range tagIDs {
		err := bs.bannerTagRepo.CreateBannerTag(bannerID, tagID)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (bs *BannerService) DeleteBanner(bannerID int64) error {
	err := bs.bannerRepo.DeleteBanner(bannerID)
	if err != nil {
		return err
	}

	return nil
}

func (bs *BannerService) GetBannerContentByFeatureAndTag(featureID, tagID, role int64) (string, error) {
	banner, err := bs.bannerRepo.GetBannerByFeatureAndTag(featureID, tagID, role)
	if err != nil {
		return "", err
	}

	return banner.Content, nil
}
