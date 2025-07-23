package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/namsh70747/Rest_API/internal/storage"
	"github.com/namsh70747/Rest_API/internal/types"
	"github.com/namsh70747/Rest_API/internal/utils/response"
)

func New(strg storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		student := types.Student{}

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		validate := validator.New()
		err = validate.Struct(student)
		if err != nil {
			// Handle validation errors
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		// Use the student variable here if needed, e.g., save to DB
		lastId, err := strg.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		student.Id = int(lastId)
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
