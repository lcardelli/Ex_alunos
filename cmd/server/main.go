package main

import (
    "log"
    "net/http"
    "ex_alunos/internal/db"
    "github.com/gin-gonic/gin"
)


func main() {
    // Inicializar a conexão com o banco de dados
    database, err := db.InitDB()
    if err != nil {
        log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
    }
    defer database.Close()

    r := gin.Default()
    r.Static("/static", "./web/static")
    r.LoadHTMLGlob("web/templates/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    r.GET("/search", func(c *gin.Context) {
        nameOrCPF := c.Query("q")
        students, err := db.SearchExStudent(database, nameOrCPF)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, students)
    })

    r.Run(":8080")
}
