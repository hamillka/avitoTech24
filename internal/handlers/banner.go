package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/hamillka/avitoTech24/internal/handlers/dto"
	"github.com/hamillka/avitoTech24/internal/repositories"
	"github.com/hamillka/avitoTech24/internal/services/models"
	"go.uber.org/zap"
)

var ErrValidate = errors.New("validation error")

const (
	ADMIN = 0
)

type IBannerService interface {
	GetBannersByFeature(featureID, limit, offset int64) ([]*models.BannerWithTagIDs, error)
	GetBannersByTag(tagID, limit, offset int64) ([]*models.BannerWithTagIDs, error)
	GetBannersByFeatureAndTag(featureID, tagID int64) ([]*models.BannerWithTagIDs, error)
	GetBannerContentByFeatureAndTag(featureID, tagID int64) (string, error)
	CreateBanner(tagIDs []int64, featureID int64, content string, isActive bool) (int64, error)
	UpdateBanner(bannerID int64, tagIDs []int64, featureID int64, content string, isActive *bool) (int64, error)
	DeleteBanner(bannerID int64) error
}

type BannerHandler struct {
	service IBannerService
	logger  *zap.SugaredLogger
}

func NewBannerHandler(s IBannerService, logger *zap.SugaredLogger) *BannerHandler {
	return &BannerHandler{
		service: s,
		logger:  logger,
	}
}

func (bh *BannerHandler) GetBanners(w http.ResponseWriter, r *http.Request) {
	featureID, err := getQueryParam(r, "feature_id")
	if err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	tagID, err := getQueryParam(r, "tag_id")
	if err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	limit, err := getQueryParam(r, "limit")
	if err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	offset, err := getQueryParam(r, "offset")
	if err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	var bannersWithTags []*models.BannerWithTagIDs

	if featureID != 0 && tagID == 0 {
		bannersWithTags, err = bh.service.GetBannersByFeature(featureID, limit, offset)
		if err != nil {
			bh.logger.Errorf("banner handler: GetBannersByFeature %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Внутренняя ошибка сервера",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	} else if featureID == 0 && tagID != 0 {
		bannersWithTags, err = bh.service.GetBannersByTag(tagID, limit, offset)
		if err != nil {
			bh.logger.Errorf("banner handler: GetBannersByTag %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Внутренняя ошибка сервера",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	} else if featureID != 0 && tagID != 0 {
		bannersWithTags, err = bh.service.GetBannersByFeatureAndTag(featureID, tagID)
		if err != nil {
			bh.logger.Errorf("banner handler: GetBannersByFeatureAndTag %s", err)
			var errorDto *dto.ErrorDto
			if errors.Is(err, repositories.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
				errorDto = &dto.ErrorDto{
					Error: "Баннер не найден",
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				errorDto = &dto.ErrorDto{
					Error: "Внутренняя ошибка сервера",
				}
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}

	getBannersResponseDto := make([]*dto.GetBannersResponseDto, 0)

	for _, bannerWithTag := range bannersWithTags {
		getBannersResponseDto = append(getBannersResponseDto, dto.ConvertToDto(*bannerWithTag))
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getBannersResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (bh *BannerHandler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	var banner dto.CreateOrUpdateBannerRequestDto

	ctx := r.Context()
	claims := ctx.Value("props").(jwt.MapClaims)
	role, _ := claims["role"].(int)
	if role != ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&banner)
	if err != nil {
		bh.logger.Errorf("banner handler: json body decoding %s", err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	content, err := json.Marshal(banner.Content)
	id, err := bh.service.CreateBanner(banner.TagIDs, banner.FeatureID, string(content), *banner.IsActive)
	if err != nil {
		bh.logger.Errorf("banner handler: CreateBanner %s", err)
		var errorDto *dto.ErrorDto
		if errors.Is(err, repositories.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			errorDto = &dto.ErrorDto{
				Error: "Запись с данным идентификатором фичи и/или тега не найдена",
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			errorDto = &dto.ErrorDto{
				Error: "Внутренняя ошибка сервера",
			}
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	createOrUpdateBannerResponseDto := dto.CreateOrUpdateBannerResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createOrUpdateBannerResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (bh *BannerHandler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	var banner dto.CreateOrUpdateBannerRequestDto

	ctx := r.Context()
	claims := ctx.Value("props").(jwt.MapClaims)
	role, _ := claims["role"].(int)
	if role != ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	param, ok := mux.Vars(r)["id"]
	bannerID, err := strToInt64(param)
	if !ok || err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewDecoder(r.Body).Decode(&banner)
	if err != nil {
		bh.logger.Errorf("banner handler: json body decoding %s", err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	content, err := json.Marshal(banner.Content)
	id, err := bh.service.UpdateBanner(bannerID, banner.TagIDs, banner.FeatureID, string(content), banner.IsActive)
	if err != nil {
		bh.logger.Errorf("banner handler: UpdateBanner %s", err)
		var errorDto *dto.ErrorDto
		if errors.Is(err, repositories.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			errorDto = &dto.ErrorDto{
				Error: "Баннер с данным идентификатором не найден",
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			errorDto = &dto.ErrorDto{
				Error: "Внутренняя ошибка сервера",
			}
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	createOrUpdateBannerResponseDto := dto.CreateOrUpdateBannerResponseDto{
		ID: id,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createOrUpdateBannerResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (bh *BannerHandler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value("props").(jwt.MapClaims)
	role, _ := claims["role"].(int)
	if role != ADMIN {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	param, ok := mux.Vars(r)["id"]
	bannerID, err := strToInt64(param)
	if !ok || err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = bh.service.DeleteBanner(bannerID)
	if err != nil {
		bh.logger.Errorf("banner handler: DeleteBanner %s", err)
		var errorDto *dto.ErrorDto
		if errors.Is(err, repositories.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			errorDto = &dto.ErrorDto{
				Error: "Баннер для тэга не найден",
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			errorDto = &dto.ErrorDto{
				Error: "Внутренняя ошибка сервера",
			}
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (bh *BannerHandler) GetUserBanner(w http.ResponseWriter, r *http.Request) {
	featureID, err := getQueryParam(r, "feature_id")
	if err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	tagID, err := getQueryParam(r, "tag_id")
	if err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	_, err = getQueryParam(r, "use_last_revision") // useLastRevision
	if err != nil {
		bh.logger.Warn(err)
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Некорректные данные",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	content, err := bh.service.GetBannerContentByFeatureAndTag(featureID, tagID)
	if err != nil {
		bh.logger.Errorf("banner handler: GetBannerContentByFeatureAndTag %s", err)
		var errorDto *dto.ErrorDto
		if errors.Is(err, repositories.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			errorDto = &dto.ErrorDto{
				Error: "Баннер для пользователя не найден",
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			errorDto = &dto.ErrorDto{
				Error: "Внутренняя ошибка сервера",
			}
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	getUserBannerResponseDto := &dto.GetUserBannerResponseDto{
		Content: content,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getUserBannerResponseDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getQueryParam(r *http.Request, key string) (int64, error) {
	var val string
	if val = r.URL.Query().Get(key); val == "" {
		return 0, nil
	}

	result, err := strToInt64(val)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func strToInt64(val string) (int64, error) {
	result, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, ErrValidate
	}

	return result, nil
}
