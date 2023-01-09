package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/william22913/chat-api/config"
	"github.com/william22913/common/dto/out"
	srvErr "github.com/william22913/common/error"
)

type Service struct {
	Config config.Configuration
}

func (s *Service) UnmarshalMessage(
	r *http.Request,
	data interface{},
) (
	err error,
) {

	b, err := io.ReadAll(r.Body)

	if err != nil {
		log.Error().
			Str("path", r.URL.Path).
			Err(err).
			Caller().
			Msg("Error aquired when read body request")

		return err
	}

	err = json.Unmarshal(b, data)

	if err != nil {
		log.Error().
			Str("path", r.URL.Path).
			Err(err).
			Caller().
			Msg("Error aquired when unmarshaling body request")

		return err
	}

	return
}

func (s *Service) AfterServiceProcess(
	action string,
	response interface{},
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {

	apiResponse := out.DefaultResponse{
		Success: true,
		Payload: response,
	}

	w.Header().Add("Content-Type", "application/json")

	if err == nil {
		log.Info().
			Str("path", r.URL.Path).
			Str("action", action).
			Msg("Success")

		w.WriteHeader(http.StatusOK)

	} else {

		log.Error().
			Str("path", r.URL.Path).
			Str("action", action).
			Err(err).
			Msg("Error aquired when handle path")

		apiResponse.Success = false
		errs := srvErr.ReformatErrorMessage(err)
		apiResponse.Payload = errs
		w.WriteHeader(errs.Error.Code)
	}

	data, errs := json.Marshal(apiResponse)

	if errs != nil {

		w.WriteHeader(500)
		log.Error().
			Str("path", r.URL.Path).
			Str("action", action).
			Err(errs).
			Msg("Error aquired when marshaling response")
	}

	_, errs = w.Write(data)

	if errs != nil {
		log.Error().
			Str("path", r.URL.Path).
			Str("action", action).
			Err(errs).
			Msg("Error aquired when write response")
	}

}
