package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/vinitborad/students-api/internal/storage"
	"github.com/vinitborad/students-api/internal/types"
	"github.com/vinitborad/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
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

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(
			w,
			http.StatusCreated,
			map[string]string{
				"success": "OK",
				"id":      strconv.FormatInt(lastId, 10),
			},
		)
	}
}
