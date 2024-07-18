package handler

import (
    "database/sql"
    "net/http"
    "ex_alunos/internal/auth"
    "ex_alunos/internal/db"
    "github.com/gin-gonic/gin"
)

func HandleMain(c *gin.Context) {
    var html = `<html><body><a href="/login">Google Login</a></body></html>`
    c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func HandleGoogleLogin(c *gin.Context) {
    auth.HandleGoogleLogin(c.Writer, c.Request)
}

func HandleGoogleCallback(c *gin.Context) {
    auth.HandleGoogleCallback(c.Writer, c.Request)
}

func SearchHandler(dbConn *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        nameOrCPF := c.Query("query")
        students, err := db.SearchExStudent(dbConn, nameOrCPF)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, students)
    }
}