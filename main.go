package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"net/http"

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
	// TODO
}

func Index(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(reports)
	if err != nil {
		fmt.Printf("Error!: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Show(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// TODO
}

func Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// TODO
}

func main() {
	mux := goji.NewMux()

	// Simple C-R-U-D APIs
	mux.HandleFunc(pat.Post("/"), Create)
	mux.HandleFunc(pat.Get("/"), Index)
	mux.HandleFuncC(pat.Get("/:id"), Show)
	mux.HandleFuncC(pat.Put("/:id"), Update)
	mux.HandleFuncC(pat.Delete("/:id"), Delete)

	http.ListenAndServe("localhost:9999", mux)
}
