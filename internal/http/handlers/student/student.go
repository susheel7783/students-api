package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/susheel7783/students-api/internal/types"
	"github.com/susheel7783/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student...") //now we will create student logic here and in postman we will test it we will give json body name email age

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation for this we will install a package go get github.com/go-playground/validator/v10 we have to run this command in terminal

		if err := validator.New().Struct(student); err != nil {
			validateErrors := err.(validator.ValidationErrors) //type casting
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrors))
			return
		}

		// w.Write([]byte("Welcome to students-api"))
		response.WriteJSON(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
