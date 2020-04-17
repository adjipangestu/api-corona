package controllers

import (
	"corona/api/responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Provinsi struct {
	Name   			string   `json:"Provinsi"`
	Kasus_Posi     	int32    `json:"Kasus_Posi"`
	Kasus_Semb    	int32    `json:"Kasus_Semb"`
	Kasus_Meni   	int32    `json:"Kasus_Meni"`
}

type DataProvinsi struct {
	Name    	string    `json:"name"`
	Positif     int32    `json:"positif"`
	Sembuh    	int32    `json:"sembuh"`
	Meninggal   int32    `json:"meninggal"`
}

type Data struct {
	Provinsi   []*DataProvinsi `json:"data"`
}

type Attributes struct {
	Attributes Provinsi   `json:"attributes"`
}

func (list Attributes) ToItem() *DataProvinsi {
	return &DataProvinsi {
		Name  : list.Attributes.Name,
		Positif  : list.Attributes.Kasus_Posi,
		Meninggal  : list.Attributes.Kasus_Meni,
		Sembuh  : list.Attributes.Kasus_Semb,
	}
}

func (server *Server) GetDataProvinsi(w http.ResponseWriter, r *http.Request)  {
	url_prov := "https://api.kawalcorona.com/indonesia/provinsi/"
	response_prov, err := http.Get(url_prov)
	if err != nil {
		log.Fatal(err)
	}
	defer response_prov.Body.Close()

	responseData_prov, err := ioutil.ReadAll(response_prov.Body)

	var provinsi []Attributes
	jsonErr_prov := json.Unmarshal([]byte(responseData_prov), &provinsi)
	if jsonErr_prov != nil {
		panic(jsonErr_prov)
	}

	list := make([]*DataProvinsi, len(provinsi))
	for i, _ := range provinsi {
		list[i] = provinsi[i].ToItem()
	}

	data, err := json.Marshal(&Data{list})
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
