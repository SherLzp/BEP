package service

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/BEP/bep_backend/models"
)

type EncryptedData struct {
	Ciphertext string `json:ciphertext`
	Capsule string `json:capsule`
}

const (
	pre_url string = "http://127.0.0.1:5000"
	content_type string= "application/json;charset=utf-8"
)

// generate keypair
func generateKeypair() (*models.Keypair, error) {
	url := pre_url + "/generateKey";

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("generate keypair failed")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed")
	}

	var keypair models.Keypair
	err = json.Unmarshal(body, &keypair)
	if err != nil {
		return nil, fmt.Errorf("response decode error")
	}

	return &keypair, nil
}

// encrypt msg, return (ciphertext, capsule)
func encryptMsg(msg string, privkey string) (*EncryptedData, error) {
	type Req struct {
		Msg string `json:msg`
		PrivateKey string `json:private_key`
	}

	b, err := json.Marshal(Req{Msg: msg, PrivateKey: privkey})
	if err != nil {
		return nil, fmt.Errorf("json format error")
	}

	url := pre_url + "/encryptData"
	body := bytes.NewBuffer(b)
	res, err := http.Post(url, content_type, body)
	if err != nil {
		return nil, fmt.Errorf("post request failed")
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error")
	}

	var enc EncryptedData
	err = json.Unmarshal(content, &enc)
	if err != nil {
		return nil, fmt.Errorf("json decode error")
	}
	return &enc, nil
}

// return a list of kfrags
func generateKfrags(privkey string, signkey string, recv_pubkey string, threshold int, n int) ([]string, error) {
	type Req struct {
		PrivateKey string `json:private_key`
		SigningKey string `json:signing_key`
		ReceivingPubKey string `json:receiving_pubkey`
		Threshold int `json:threshold`
		N int `json:n`
	}

	b, err := json.Marshal(Req{PrivateKey: privkey, SigningKey: signkey, ReceivingPubKey: recv_pubkey, Threshold: threshold, N: n})
	if err != nil {
		return nil, fmt.Errorf("json format error")
	}

	url := pre_url + "/generateKfrags"
	body := bytes.NewBuffer(b)
	res, err := http.Post(url, content_type, body)
	if err != nil {
		return nil, fmt.Errorf("post request error")
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error")
	}

	var kfrags_list []string
	err = json.Unmarshal([]byte(content), &kfrags_list)
	if err != nil {
		return nil, fmt.Errorf("json decode error")
	}
	return kfrags_list, nil
}

// return a list of cfrags
func reencrypt(pubkey string, verikey string, recv_pubkey string, kfrags_list []string, capsule string) ([]string, error) {
	type Req struct {
		PublicKey string `json:public_key`
		VerifyingKey string `json:verifying_key`
		ReceivingPubKey string `json:receiving_pubkey`
		KfragsList string `json:kfrags`
		Capsule string `json:capsule`
	}

	b, err := json.Marshal(Req{PublicKey: pubkey, VerifyingKey: verikey, ReceivingPubKey: recv_pubkey, KfragsList: kfrags_list, Capsule: capsule})
	if err != nil {
		return nil, fmt.Errorf("json format error")
	}

	url := pre_url + "/reencrypt"
	body := bytes.NewBuffer(b)
	res, err := http.Post(url, content_type, body)
	if err != nil {
		return nil, fmt.Errorf("post request error")
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error")
	}

	var cfrags []string
	err = json.Unmarshal([]byte(content), &cfrags)
	if err != nil {
		return nil, fmt.Errorf("json decode error")
	}
	return cfrags, nil
}

// return decrypted msg
func decrypt(pubkey string, verikey string, recv_pubkey string, ciphertext string, cfrags []string, capsule string) (string, error) {
	type Req struct {
		PublicKey string `json:public_key`
		VerifyingKey string `json:verifying_key`
		ReceivingPubKey string `json:receiving_pubkey`
		Ciphertext string `json:ciphertext`
		Capsule string `json:capsule`
	}

	b, err := json.Marshal(Req{PublicKey: pubkey, VerifyingKey: verikey, ReceivingPubKey: recv_pubkey, Ciphertext: ciphertext, Capsule: capsule})
	if err != nil {
		return nil, fmt.Errorf("json format error")
	}

	url := pre_url + "/decrypt"
	body := bytes.NewBuffer(b)
	res, err := http.Post(url, content_type, body)
	if err != nil {
		return nil, fmt.Errorf("post request error")
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error")
	}
	return content, nil
}
