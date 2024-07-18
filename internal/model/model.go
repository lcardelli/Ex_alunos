package model

type ExStudent struct {
    Name                 string `json:"name"`
    RA                   string `json:"ra"`
    Serie                string `json:"serie"`
    Year                 string `json:"year"`
    Course               string `json:"course"`
    Description          string `json:"description"`
    Branch               int    `json:"branch"`
}

