package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "ex_alunos/internal/model"
    "github.com/joho/godotenv"
	_"github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
    // Carregar variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Obter credenciais do banco de dados das variáveis de ambiente
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Montar a string de conexão
    connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

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

func SearchExStudent(db *sql.DB, nameOrRA string) ([]model.ExStudent, error) {
    query := `
    SELECT
        PPESSOA.NOME,
        SALUNO.RA,
        SHABILITACAO.NOME AS SÉRIE,
        SPLETIVO.CODPERLET AS ANO,
        SCURSO.NOME AS CURSO,
        'Ex_aluno' AS DESCRICAO_CODTIPOALUNO,
        SMATRICPL.CODFILIAL
    FROM
        Corpore.dbo.SALUNO (nolock)
        INNER JOIN Corpore.dbo.PPESSOA ON SALUNO.CODPESSOA = PPESSOA.CODIGO
        INNER JOIN Corpore.dbo.SMATRICPL ON SMATRICPL.CODCOLIGADA = SALUNO.CODCOLIGADA AND SMATRICPL.RA = SALUNO.RA
        INNER JOIN Corpore.dbo.SPLETIVO ON SMATRICPL.CODCOLIGADA = SPLETIVO.CODCOLIGADA AND SMATRICPL.IDPERLET = SPLETIVO.IDPERLET
        INNER JOIN Corpore.dbo.SHABILITACAOFILIAL ON SHABILITACAOFILIAL.IDHABILITACAOFILIAL = SMATRICPL.IDHABILITACAOFILIAL
        INNER JOIN Corpore.dbo.STURNO ON SHABILITACAOFILIAL.CODCOLIGADA = STURNO.CODCOLIGADA AND SHABILITACAOFILIAL.CODTURNO = STURNO.CODTURNO
        INNER JOIN Corpore.dbo.SHABILITACAO ON SHABILITACAOFILIAL.CODHABILITACAO = SHABILITACAO.CODHABILITACAO AND SHABILITACAOFILIAL.CODCURSO = SHABILITACAO.CODCURSO
        INNER JOIN Corpore.dbo.SCURSO ON SHABILITACAO.CODCOLIGADA = SCURSO.CODCOLIGADA AND SHABILITACAO.CODCURSO = SCURSO.CODCURSO AND SHABILITACAOFILIAL.CODCURSO = SCURSO.CODCURSO
        INNER JOIN Corpore.dbo.SSTATUS ON SMATRICPL.CODCOLIGADA = SSTATUS.CODCOLIGADA AND SMATRICPL.CODSTATUS = SSTATUS.CODSTATUS
    WHERE
        SALUNO.CODCOLIGADA = 1
        AND SMATRICPL.CODFILIAL IN (1)
        AND SALUNO.CODTIPOALUNO = '2'
        AND SCURSO.CODCURSO IN ('2', '3', '4')
        AND SPLETIVO.CODPERLET BETWEEN '2000' AND '2024'
        AND SPLETIVO.CODPERLET = (
            SELECT MAX(SP1.CODPERLET)
            FROM Corpore.dbo.SPLETIVO SP1
            INNER JOIN Corpore.dbo.SMATRICPL SM1 ON SM1.CODCOLIGADA = SP1.CODCOLIGADA AND SM1.IDPERLET = SP1.IDPERLET
            WHERE SM1.RA = SALUNO.RA
            AND SM1.CODCOLIGADA = SALUNO.CODCOLIGADA
        )
        AND (PPESSOA.NOME LIKE ? OR SALUNO.RA = ?)
    `

    rows, err := db.Query(query, "%"+nameOrRA+"%", nameOrRA)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []model.ExStudent
    for rows.Next() {
        var student model.ExStudent
        if err := rows.Scan(&student.Name, &student.RA, &student.Serie, &student.Year, &student.Course, &student.Description, &student.Branch); err != nil {
            return nil, err
        }
        students = append(students, student)
    }
    return students, nil
}
