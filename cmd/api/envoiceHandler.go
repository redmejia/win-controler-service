package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"win/controler/internal/models"
	"win/controler/utils"
)

func (a *ApiConfig) EnvoiceHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8081/api/env?company=1238878-89883hdsj-2197ejds

	if r.Method == http.MethodGet {
		companyUID := r.URL.Query().Get("company")

		url := fmt.Sprintf("http://localhost:8089/api/env/all?c_uid=%s", companyUID)

		resp, err := http.Get(url)
		if err != nil {
			a.Errorlog.Fatal(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusFound {
			a.Errorlog.Fatalf("bad status code expect %d but %d was recived insted", http.StatusFound, resp.StatusCode)

		}

		var envoices []models.CompanyEnvoice

		decode := json.NewDecoder(resp.Body)
		err = decode.Decode(&envoices)
		if err != nil {
			a.Errorlog.Fatal(err)
		}

		e, _ := json.MarshalIndent(&envoices, "", " ")

		a.Infolog.Println(string(e))

		utils.WriteJSON(w, http.StatusFound, envoices)

	}
}
