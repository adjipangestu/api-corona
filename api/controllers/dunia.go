package controllers

import (
	"corona/api/responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Negara struct {
	Name    	string    `json:"Country_Region"`
	Positif     int32    `json:"Confirmed"`
	Sembuh    	int32    `json:"Recovered"`
	Meninggal   int32    `json:"Deaths"`
}

type DataNegara struct {
	Name    	string    `json:"name"`
	Positif     int32    `json:"positif"`
	Sembuh    	int32    `json:"sembuh"`
	Meninggal   int32    `json:"meninggal"`
}

type DataDunia struct {
	Dunia   []*DataNegara `json:"data"`
}

type AttributesNegara struct {
	Attributes Negara   `json:"attributes"`
}

func (list AttributesNegara) ToItemNegara() *DataNegara {
	return &DataNegara {
		Name  : list.Attributes.Name,
		Positif  : list.Attributes.Positif,
		Meninggal  : list.Attributes.Meninggal,
		Sembuh  : list.Attributes.Sembuh,
	}
}

func (server *Server) GetDataNegara(w http.ResponseWriter, r *http.Request)  {
	url := "https://api.kawalcorona.com/"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	var negara []AttributesNegara
	jsonErr:= json.Unmarshal([]byte(responseData), &negara)
	if jsonErr != nil {
		panic(jsonErr)
	}

	list := make([]*DataNegara, len(negara))
	for i, _ := range negara {
		list[i] = negara[i].ToItemNegara()
	}

	data, err := json.Marshal(&DataDunia{list})
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
