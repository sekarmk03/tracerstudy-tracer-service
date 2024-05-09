package service

import (
	"bytes"
	"encoding/json"

	// "fmt"
	"io"
	"log"
	"net/http"
	"time"
	"tracerstudy-tracer-service/common/config"

	// "tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/modules/mhsbiodata/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

const (
	apiMaxRetries = 3
	sleepTime     = 500 * time.Millisecond
)

type MhsBiodataService struct {
	cfg config.Config
}

type MhsBiodataServiceUseCase interface {
	FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsBiodata, error)
}

func NewMhsBiodataService(cfg config.Config) *MhsBiodataService {
	return &MhsBiodataService{
		cfg: cfg,
	}
}

func (svc *MhsBiodataService) FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsBiodata, error) {
	payload := map[string]string{"nim": nim}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while marshalling payload:", err)
		if _, isUnsupportedTypeError := err.(*json.UnsupportedTypeError); isUnsupportedTypeError {
			return nil, status.Errorf(codes.InvalidArgument, "invalid payload: unsupported data type")
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	apiUrl := svc.cfg.SIAK_API.URL
	apiKey := svc.cfg.SIAK_API.KEY

	for attempt := 1; attempt <= apiMaxRetries; attempt++ {
		reqHttp, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while creating HTTP request:", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}

		reqHttp.Header.Set("Api-Key", apiKey)
		reqHttp.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(reqHttp)
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while sending HTTP request:", err)

			if attempt == apiMaxRetries {
				log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached:", err)
				return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
			}

			time.Sleep(sleepTime)
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if attempt == apiMaxRetries {
				log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached:", resp.StatusCode, resp.Body)
				return nil, status.Errorf(codes.Internal, "internal server error: HTTP request failed with status code: %d", resp.StatusCode)
			}

			time.Sleep(sleepTime)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while reading HTTP response body:", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}

		var apiResponse []entity.MhsBiodata
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while unmarshalling HTTP response body:", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}

		if len(apiResponse) == 0 {
			log.Println("WARNING: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Resource not found: nim", nim)
			return nil, status.Errorf(codes.NotFound, "mhs resource not found")
		}

		return &apiResponse[0], nil
	}

	log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached without success")
	return nil, status.Errorf(codes.Internal, "internal server error: maximum retries reached without success")
}
