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
	"github.com/susheel7783/students-api/internal/storage"
	"github.com/susheel7783/students-api/internal/types"
	"github.com/susheel7783/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
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

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, err)
			return
		}

		// w.Write([]byte("Welcome to students-api"))
		response.WriteJSON(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student by ID", slog.String("id", id))
		// Implementation for getting student by ID will go here
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("failed to get student by ID", slog.String("id", id))
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, student)

	}
}
