package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"win/controler/internal/models"
	"win/controler/utils"
)

func (a *ApiConfig) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

// TxControler Transaction Controler
func (a *ApiConfig) TxHandler(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction

	err := utils.ReadJSON(r, &tx)
	if err != nil {
		a.Errorlog.Fatal(err)
	}

	txStatus := checkTx(a, &tx)
	// status code 2 is accepted
	// create envoice with accepted status
	if txStatus.Proceed && txStatus.TxStatusCode == 2 {
		tx.TxDate = time.Now()
		txData := models.TransactionData{
			TxAccepted:   txStatus.Proceed,
			MessageState: txStatus.TxMessage,
			Date:         time.Now(),
			Transaction:  tx,
		}

		envoInfo := createEnvoice(a, txData)

		a.Infolog.Println(envoInfo)

		err := utils.WriteJSON(w, http.StatusOK, envoInfo)
		if err != nil {
			a.Errorlog.Println(err)
		}

	}
	// status code 4 is decline status
	// create envoice with decline status
	if txStatus.TxStatusCode == 4 {
		tx.TxDate = time.Now()
		txData := models.TransactionData{
			TxAccepted:   txStatus.Proceed,
			MessageState: txStatus.TxMessage,
			Date:         time.Now(),
			Transaction:  tx,
		}

		envoInfo := createEnvoice(a, txData)

		a.Infolog.Println(envoInfo)

		err := utils.WriteJSON(w, http.StatusOK, envoInfo)
		if err != nil {
			a.Errorlog.Println(err)
		}
		// declinePending() no envoice was created but trasanction decline is save for record
		// err := utils.WriteJSON(w, http.StatusOK, txData)
		// if err != nil {
		// 	a.Errorlog.Println(err)
		// }

	}

}

// createEnvoice create envoice if the transaction was accepted
func createEnvoice(a *ApiConfig, txData models.TransactionData) models.EnvoiceInfo {

	url := "http://localhost:8089/api/env"

	data, err := json.Marshal(txData)
	if err != nil {
		a.Errorlog.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		a.Errorlog.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// must response with the envoice information
	resp, err := client.Do(req)
	if err != nil {
		a.Errorlog.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		a.Errorlog.Fatalf("bad response expected %d status code but recived %d ",
			http.StatusCreated, resp.StatusCode)
	}

	var envoiceInfo models.EnvoiceInfo
	json.NewDecoder(resp.Body).Decode(&envoiceInfo)

	return envoiceInfo
}

// declinePending transaction was declined
func declinedPending() {
	return
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
		a.Errorlog.Fatalf("bad status code expect %d but %d was recived insted ",
			http.StatusAccepted, resp.StatusCode)
	}

	var txStatus models.TransactionStatus
	json.NewDecoder(resp.Body).Decode(&txStatus)

	return txStatus
}
