package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"time"

	"goji.io"
	"goji.io/pat"
)

type Report struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Article   string `json:"article"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

// Global variable (instead of Database)
var reports = []Report{
	// Sample records
	Report{Id: 1, Title: "title1", Article: "this is article1", CreatedAt: 1470495600, UpdatedAt: 1470495600},
	Report{Id: 3, Title: "title3", Article: "here is article3", CreatedAt: 1470495800, UpdatedAt: 1470495900},
}

// Errors
type ReportNotFoundError struct {
	Msg    string
	Status int
}

type WrongParameterError struct {
	Msg    string
	Status int
}

type JsonError struct {
	Msg    string
	Status int
}

func (err ReportNotFoundError) Error() string {
	return err.Msg
}

func (err WrongParameterError) Error() string {
	return err.Msg
}

func (err JsonError) Error() string {
	return err.Msg
}

// Helpers
func ParseId(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, WrongParameterError{"ERROR! Wrong parameter.", http.StatusBadRequest}
	}
	return id, nil
}

func FindReport(id int) (int, error) {
	indexAt := -1
	for i := 0; i < len(reports); i++ {
		if reports[i].Id == id {
			indexAt = i
		}
	}
	if indexAt == -1 {
		return 0, ReportNotFoundError{"ERROR! Report not found.", http.StatusNotFound}
	}
	return indexAt, nil
}

// Endpoints
func Create(w http.ResponseWriter, r *http.Request) {
	newId := reports[len(reports)-1].Id + 1
	now := int(time.Now().Unix())
	newReport := Report{Id: newId, CreatedAt: now, UpdatedAt: now}
	err := json.NewDecoder(r.Body).Decode(&newReport)
	if err != nil {
		e := JsonError{"ERROR! Decode newReport failed.", http.StatusInternalServerError}
		http.Error(w, e.Msg, e.Status)
		return
	}

	reports = append(reports, newReport)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("New report id: %v", newReport.Id)))
}

func Index(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(reports)
	if err != nil {
		e := JsonError{"ERROR! Marshal reports failed.", http.StatusInternalServerError}
		http.Error(w, e.Msg, e.Status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Show(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := ParseId(pat.Param(ctx, "id"))
	if err != nil {
		switch e := err.(type) {
		case WrongParameterError:
			http.Error(w, e.Msg, e.Status)
			return
		}
	}

	indexAt, err := FindReport(id)
	if err != nil {
		switch e := err.(type) {
		case ReportNotFoundError:
			http.Error(w, e.Msg, e.Status)
			return
		}
	}

	res, err := json.Marshal(reports[indexAt])
	if err != nil {
		e := JsonError{"ERROR! Marshal reports[indeAt] failed.", http.StatusInternalServerError}
		http.Error(w, e.Msg, e.Status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := ParseId(pat.Param(ctx, "id"))
	if err != nil {
		switch e := err.(type) {
		case WrongParameterError:
			http.Error(w, e.Msg, e.Status)
			return
		}
	}

	indexAt, err := FindReport(id)
	if err != nil {
		switch e := err.(type) {
		case ReportNotFoundError:
			http.Error(w, e.Msg, e.Status)
			return
		}
	}

	now := int(time.Now().Unix())
	updatedReport := Report{Id: reports[indexAt].Id, CreatedAt: reports[indexAt].CreatedAt, UpdatedAt: now}
	jsonErr := json.NewDecoder(r.Body).Decode(&updatedReport)
	if jsonErr != nil {
		e := JsonError{"ERROR! Decode newReport failed.", http.StatusInternalServerError}
		http.Error(w, e.Msg, e.Status)
		return
	}

	// TODO: Keep sorting...
	reports = append(reports[:indexAt], reports[indexAt+1:]...)
	reports = append(reports, updatedReport)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("Update report id: %v", updatedReport.Id)))
}

func Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := ParseId(pat.Param(ctx, "id"))
	if err != nil {
		switch e := err.(type) {
		case WrongParameterError:
			http.Error(w, e.Msg, e.Status)
			return
		}
	}

	indexAt, err := FindReport(id)
	if err != nil {
		switch e := err.(type) {
		case ReportNotFoundError:
			http.Error(w, e.Msg, e.Status)
			return
		}
	}

	reports = append(reports[:indexAt], reports[indexAt+1:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Delete report id: %v", id)))
}

func main() {
	mux := goji.NewMux()

	// Simple C-R-U-D APIs
	mux.HandleFuncC(pat.Get("/:id"), Show)
	mux.HandleFuncC(pat.Put("/:id"), Update)
	mux.HandleFuncC(pat.Delete("/:id"), Delete)
	mux.HandleFunc(pat.Post("/"), Create)
	mux.HandleFunc(pat.Get("/"), Index)

	http.ListenAndServe("localhost:9999", mux)
}
