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

var reports = []Report{
	// Sample records
	Report{Id: 1, Title: "title1", Article: "this is article1", CreatedAt: 1470495600, UpdatedAt: 1470495600},
	Report{Id: 3, Title: "title3", Article: "here is article3", CreatedAt: 1470495800, UpdatedAt: 1470495900},
}

func Create(w http.ResponseWriter, r *http.Request) {
	newId := reports[len(reports)-1].Id + 1
	now := int(time.Now().Unix())
	report := Report{Id: newId, CreatedAt: now, UpdatedAt: now}

	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		fmt.Printf("Error!: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reports = append(reports, report)
	http.Error(w, fmt.Sprintf("Report Created! id: %v", newId), http.StatusCreated)
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(reports)
	if err != nil {
		fmt.Printf("Error!: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func Show(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(pat.Param(ctx, "id"))
	if err != nil {
		fmt.Printf("Error!: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	indexAt := -1
	for i := 0; i < len(reports); i++ {
		if reports[i].Id == id {
			indexAt = i
		}
	}

	if indexAt == -1 {
		fmt.Printf("Error!: Report Not Found\n")
		http.Error(w, "Report Not Found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(reports[indexAt])
	if err != nil {
		fmt.Printf("Error!: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(pat.Param(ctx, "id"))
	if err != nil {
		fmt.Printf("Error!: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	indexAt := -1
	for i := 0; i < len(reports); i++ {
		if reports[i].Id == id {
			indexAt = i
		}
	}

	if indexAt == -1 {
		fmt.Printf("Error!: Report Not Found\n")
		http.Error(w, "Report Not Found", http.StatusNotFound)
		return
	}

	now := int(time.Now().Unix())
	report := Report{Id: reports[indexAt].Id, CreatedAt: reports[indexAt].CreatedAt, UpdatedAt: now}
	err2 := json.NewDecoder(r.Body).Decode(&report)
	if err2 != nil {
		fmt.Printf("Error!: %v\n", err2.Error())
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Keep sorting...
	reports = append(reports[:indexAt], reports[indexAt+1:]...)
	reports = append(reports, report)
	http.Error(w, fmt.Sprintf("Report Updated! id: %v", id), http.StatusCreated)
	return
}

func Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(pat.Param(ctx, "id"))
	if err != nil {
		fmt.Printf("Error!: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	indexAt := -1
	for i := 0; i < len(reports); i++ {
		if reports[i].Id == id {
			indexAt = i
		}
	}

	if indexAt == -1 {
		fmt.Printf("Error!: Report Not Found\n")
		http.Error(w, "Report Not Found", http.StatusNotFound)
		return
	}

	reports = append(reports[:indexAt], reports[indexAt+1:]...)
	http.Error(w, fmt.Sprintf("Report Deleted! id: %v", id), http.StatusNoContent)
	return
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
