package controllers

import (
	"corona/api/responses"
	"encoding/json"
	"net/http"
)

type Author struct {
	Name    string	`json:"name"`
	Jobs   	string	`json:"jobs"`
	Url_Prov   	string	`json:"provinsi"`
	Url_Dunia   	string	`json:"dunia"`
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request)  {
	author := Author{
		Name:	"Adji Pangestu",
		Jobs:	"Software Enginerr at detikcom",
		Url_Prov:	"/provinsi",
		Url_Dunia:	"/all",
	}

	var data []byte
	data, err := json.Marshal(author)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
