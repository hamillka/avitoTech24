package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hamillka/avitoTech24/internal/handlers/dto"
	"github.com/hamillka/avitoTech24/internal/services/models"
)

type IBannerService interface {
	GetBannersByFeature(featureID, limit, offset int64) ([]*models.BannerWithTagIDs, error)
	GetBannersByTag(tagID, limit, offset int64) ([]*models.BannerWithTagIDs, error)
	GetBannersByFeatureAndTag(featureID, tagID, limit, offset int64) ([]*models.BannerWithTagIDs, error)
}

type BannerHandler struct {
	service IBannerService
}

func (bh *BannerHandler) GetBanners(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	featureID, err := getQueryParam(params, "feature_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Возникла внутренняя ошибка",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	tagID, err := getQueryParam(params, "tag_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Возникла внутренняя ошибка",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	limit, err := getQueryParam(params, "limit")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Возникла внутренняя ошибка",
		}
		err = json.NewEncoder(w).Encode(errorDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	offset, err := getQueryParam(params, "offset")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorDto := &dto.ErrorDto{
			Error: "Возникла внутренняя ошибка",
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
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Возникла внутренняя ошибка",
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
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Возникла внутренняя ошибка",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	} else if featureID != 0 && tagID != 0 {
		bannersWithTags, err = bh.service.GetBannersByFeatureAndTag(featureID, tagID, limit, offset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorDto := &dto.ErrorDto{
				Error: "Возникла внутренняя ошибка",
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

func getQueryParam(params map[string]string, key string) (int64, error) {
	var result int64
	var err error

	if param, exist := params[key]; exist {
		result, err = strconv.ParseInt(param, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("internal server error") // TODO: custom error
		}
	}

	return result, nil
}
