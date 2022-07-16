package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"win/controler/internal/models"
)

// TxControler Transaction Controler
func (a *ApiConfig) TxControler(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction

	json.NewDecoder(r.Body).Decode(&tx)

	a.Infolog.Println("trying ", tx)

	txStatus := checkTx(a, &tx)

	a.Infolog.Println(txStatus)

}

// checkTx check transaction
func checkTx(a *ApiConfig, inProcessTx *models.Transaction) models.TransactionStatus {

	url := fmt.Sprintf("http://localhost:8083/api/txintent?card=%s&cv=%s&amount=%d",
		inProcessTx.TxCardNumber, inProcessTx.TxCardCv, inProcessTx.TxAmount)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		a.Errorlog.Fatalf("bad status code expect %d but %d was recived insted ", http.StatusAccepted, resp.StatusCode)
	}

	var txStatus models.TransactionStatus
	json.NewDecoder(resp.Body).Decode(&txStatus)

	return txStatus
}
