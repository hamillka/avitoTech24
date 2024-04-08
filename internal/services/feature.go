package services

type FeatureService struct {
	repo IFeatureRepository
}

func NewFeatureService(repository IFeatureRepository) *FeatureService {
	return &FeatureService{repo: repository}
}
