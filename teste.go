package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/denisenkom/go-mssqldb"
    "github.com/joho/godotenv"
	_"github.com/go-sql-driver/mysql"
)

func main() {
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
    connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    // Conectar ao banco de dados MySQL
    db, err := sql.Open("mysql", connString)
    if err != nil {
        log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
    }
    defer db.Close()

   // Testar a conexão
   err = db.Ping()
   if err != nil {
	   log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
   }

   fmt.Println("Conexão com o banco de dados MySQL bem-sucedida!")

   // Executar uma consulta simples
   var version string
   err = db.QueryRow("SELECT version()").Scan(&version)
   if err != nil {
	   log.Fatalf("Erro ao executar consulta: %v", err)
   }

   fmt.Println("Versão do MySQL:", version)
}