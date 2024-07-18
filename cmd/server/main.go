package main

import (
    "log"
    "net/http"
    
    "ex_alunos/internal/db"
    "ex_alunos/internal/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializa a conexão com o banco de dados
    dbConn, err := db.InitDB()
    if err != nil {
        log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
    }
    defer dbConn.Close()

    // Inicializa o router do Gin
    r := gin.Default()

    // Define o diretório para arquivos estáticos
    r.Static("/static", "./web/static")

    // Define o diretório para templates HTML
    r.LoadHTMLGlob("web/templates/*")

    // Define as rotas
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    r.GET("/login", handler.HandleGoogleLogin)
    r.GET("/callback", handler.HandleGoogleCallback)
    r.GET("/search", handler.SearchHandler(dbConn))

    // Inicia o servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}