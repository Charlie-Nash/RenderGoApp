package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
)

type Vendedor struct {
    ID     int
    Nombre string
}

func main() {
    dbURL := os.Getenv("DATABASE_URL")
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Error al conectar a la BD:", err)
    }
    defer db.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT id, nombre FROM vendedores")
        if err != nil {
            http.Error(w, "Error en la consulta", 500)
            return
        }
        defer rows.Close()

        var vendedores []Vendedor
        for rows.Next() {
            var v Vendedor
            if err := rows.Scan(&v.ID, &v.Nombre); err != nil {
                log.Println(err)
                continue
            }
            vendedores = append(vendedores, v)
        }

        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, vendedores)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Println("Servidor corriendo en el puerto", port)
    http.ListenAndServe(":"+port, nil)
}
