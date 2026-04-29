package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vinitborad/students-api/internal/types"
	"github.com/vinitborad/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GeneralError(fmt.Errorf("empty body")),
			)
			return
		}

		if err != nil {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GeneralError(err),
			)
			return
		}

		err = validator.New().Struct(student)
		if err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(
			w,
			http.StatusCreated,
			map[string]string{
				"success": "OK",
			},
		)
	}
}
