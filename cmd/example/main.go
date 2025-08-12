package main

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/tursodatabase/go-libsql"
	"github.com/vlence/configman"
	sqlstore "github.com/vlence/configman/stores/sql"
	"github.com/vlence/gossert"
)

//go:embed templates/base.html templates/index.html
var indexTemplates embed.FS

//go:embed scripts
var scriptsDir embed.FS

//go:embed styles
var stylesDir embed.FS

var indexTmpl = template.Must(template.ParseFS(indexTemplates, "templates/*.html"))

func main() {
        var db *sql.DB
        var err error
        var store configman.Store

        addr := "127.0.0.1:8080"

        db, err = sql.Open("libsql", "file:db/test.db")
        gossert.Ok(err == nil, "failed to open db")

        store, err = sqlstore.NewSqlStore(db)
        gossert.Ok(err == nil, "failed to create config store")

        http.Handle("GET /styles/", http.FileServer(http.FS(stylesDir)))
        http.Handle("GET /scripts/", http.FileServer(http.FS(scriptsDir)))
        
        http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
                var configs []configman.Config

                if configs, err = store.GetConfigs(); err != nil {
                        log.Println(err)
                        w.WriteHeader(http.StatusInternalServerError)
                        return
                }

                pageData := make(map[string]any)
                pageData["Title"] = "Configman"
                pageData["Configs"] = configs

                w.WriteHeader(http.StatusOK)

                if err = indexTmpl.ExecuteTemplate(w, "base", pageData); err != nil {
                        log.Println(err)
                }
        })

        http.HandleFunc("POST /configs/", func(w http.ResponseWriter, r *http.Request) {
                var err error
                var name string
                var config configman.Config
                var configs []configman.Config

                name = strings.TrimSpace(r.FormValue("name"))

                if config, err = store.GetConfig(name); err != nil {
                        log.Println(err)
                        w.WriteHeader(http.StatusInternalServerError)
                        return
                }

                if config == nil {
                        // create a new config if one doesn't already exist with the same name
                        if config, err = store.CreateConfig(name, ""); err != nil {
                                log.Println(err)
                                w.WriteHeader(http.StatusInternalServerError)
                                return
                        }
                }

                gossert.Ok(config != nil, "config created without error but got nil")

                if configs, err = store.GetConfigs(); err != nil {
                        log.Println(err)
                        w.WriteHeader(http.StatusInternalServerError)
                        return
                }

                pageData := make(map[string]any)
                pageData["Configs"] = configs

                w.WriteHeader(http.StatusCreated)

                if err = indexTmpl.ExecuteTemplate(w, "configs", configs); err != nil {
                        log.Println(err)
                }
        })

        http.HandleFunc("GET /configs/{name}/", func(w http.ResponseWriter, r *http.Request) {
                var err error
                var config configman.Config

                name := r.PathValue("name")

                if config, err = store.GetConfig(name); err != nil {
                        log.Println(err)
                        w.WriteHeader(http.StatusInternalServerError)
                        return
                }

                if config == nil {
                        w.WriteHeader(http.StatusNotFound)
                        return
                }

                w.WriteHeader(http.StatusOK)

                if err = indexTmpl.ExecuteTemplate(w, "config", config); err != nil {
                        log.Println(err)
                }
        })

        http.HandleFunc("PATCH /configs/{name}/", func(w http.ResponseWriter, r *http.Request) {
                var err error
                var done bool
                var name, desc string
                var config configman.Config

                name = r.PathValue("name")
                desc = r.FormValue("desc")

                if config, err = store.GetConfig(name); err != nil {
                        log.Println(err)
                        w.WriteHeader(http.StatusInternalServerError)
                        return
                }

                if config == nil {
                        w.WriteHeader(http.StatusNotFound)
                        return
                }

                if done, err = config.SetDesc(desc); err != nil {
                        log.Println(err)
                        w.WriteHeader(http.StatusInternalServerError)
                        return
                }

                if done {
                        gossert.Ok(config.Desc() == desc, "config desc not updated correctly")
                }

                if !done {
                        log.Printf("warn: desc of config %s not updated", name)
                }

                w.WriteHeader(http.StatusOK)

                if err = indexTmpl.ExecuteTemplate(w, "config-desc", config); err != nil {
                        log.Println(err)
                }
        })

        http.HandleFunc("POST /configs/{name}/settings/", func(w http.ResponseWriter, r *http.Request) {
                
        })

        log.Printf("Listening on %s\n", addr)
        log.Fatal(http.ListenAndServe(addr, logger(http.DefaultServeMux)))
}
