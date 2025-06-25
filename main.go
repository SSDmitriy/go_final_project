package main

import (
	_ "encoding/json"
	"fmt"
	"go_final_project/internal/api"
	"go_final_project/internal/server"
	"go_final_project/internal/storage"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

const dbFile = "scheduler.db"

func main() {

	if err := storage.Init(dbFile); err != nil {
		fmt.Printf("ошибка инициализации БД: %s", err)
	}

	db := storage.GetDB()

	db.Close() //!

	r := chi.NewRouter()

	api.Init(r)

	// Статические файлы
	setupStaticFiles(r)

	port := server.DefaultPort
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		fmt.Printf("ошибка 001 при запуске сервера: %s", err.Error())
		return
	}

	fmt.Printf("Сервер запущен на порту :%s", port)

}

func setupStaticFiles(r *chi.Mux) {
	root := http.Dir(server.WebDir)
	fs := http.FileServer(root)

	r.Handle("/*", fs)
	r.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir(filepath.Join(server.WebDir, "css")))))
	r.Handle("/js/*", http.StripPrefix("/js/", http.FileServer(http.Dir(filepath.Join(server.WebDir, "js")))))

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(server.WebDir, "favicon.ico"))
	})

}

// func getTask(w http.ResponseWriter, r *http.Request) {
// 	resp, err := []byte{'a', 'b', 'c'}, json.RawMessage([]byte{'1', '2'})

// 	//json.Marshal(artists)

// 	if err != nil {
// 		http.Error(w, "erorroro 999", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	w.WriteHeader(http.StatusOK)

// 	w.Write(resp)
// }
