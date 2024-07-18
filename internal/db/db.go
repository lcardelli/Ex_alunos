package db

import (
    "database/sql"
    "log"
    "ex_alunos/internal/model"
    _ "github.com/denisenkom/go-mssqldb"
)

func InitDB() (*sql.DB, error) {
    // Atualize a string de conexão para o formato MySQL
    connString := "username:password@tcp(servername:port)/dbname"

    db, err := sql.Open("mysql", connString)
    if err != nil {
        log.Fatal("Erro ao abrir conexão com o banco de dados:", err)
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Erro ao conectar ao banco de dados:", err)
        return nil, err
    }

    return db, nil
}

func SearchExStudent(db *sql.DB, nameOrCPF string) ([]model.ExStudent, error) {
    query := `SELECT Name, CPF, EntryYear FROM ex_students WHERE Name LIKE ? OR CPF = ?`
    rows, err := db.Query(query, "%"+nameOrCPF+"%", nameOrCPF)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []model.ExStudent
    for rows.Next() {
        var student model.ExStudent
        if err := rows.Scan(&student.Name, &student.RA, &student.EntryYear); err != nil {
            return nil, err
        }
        students = append(students, student)
    }
    return students, nil
}
