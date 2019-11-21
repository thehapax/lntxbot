package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func registerBluewalletMethods() {
	router.Path("/getinfo").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorBadAuth(w)
	})

	router.Path("/auth").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Login        string `json:"login"`
			Password     string `json:"password"`
			RefreshToken string `json:"refresh_token"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			errorInvalidParams(w)
			return
		}
		log.Debug().
			Str("login", params.Login).Str("password", params.Password).Str("token", params.RefreshToken).
			Msg("bluewallet /auth")

		var token string
		if params.Password == "" {
			token = params.RefreshToken
		} else {
			token = base64.StdEncoding.EncodeToString([]byte(params.Login + ":" + params.Password))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			RefreshToken string `json:"refresh_token"`
			AccessToken  string `json:"access_token"`
		}{token, token})
	})

	router.Path("/addinvoice").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, permission, err := loadUserFromAPICall(r)
		if err != nil {
			errorBadAuth(w)
			return
		}
		if permission < InvoicePermissions {
			errorInsufficientPermissions(w)
			return
		}

		var params struct {
			Amount string `json:"amt"`
			Memo   string `json:"memo"`
		}
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			errorInvalidParams(w)
			return
		}
		msatoshi, err := strconv.Atoi(params.Amount)
		if err != nil {
			errorInvalidParams(w)
			return
		}

		log.Debug().Str("amount", params.Amount).Str("memo", params.Memo).Msg("bluewallet /addinvoice")

		bolt11, hash, _, err := user.makeInvoice(msatoshi, params.Memo, "", nil, nil, "", "", true)
		if err != nil {
			errorInternal(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			PayReq         string `json:"pay_req"`
			PaymentRequest string `json:"payment_request"`
			AddIndex       string `json:"add_index"`
			RHash          Buffer `json:"r_hash"`
			Hash           string `json:"payment_hash"`
		}{bolt11, bolt11, "1000", Buffer(hash), hash})
	})

	router.Path("/payinvoice").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, permission, err := loadUserFromAPICall(r)
		if err != nil {
			errorBadAuth(w)
			return
		}
		if permission < FullPermissions {
			errorInsufficientPermissions(w)
			return
		}

		var params struct {
			Invoice string `json:"invoice"`
		}
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			errorInvalidParams(w)
			return
		}

		log.Debug().Str("bolt11", params.Invoice).Msg("bluewallet /payinvoice")

		err = user.payInvoice(0, params.Invoice)
		if err != nil {
			errorPaymentFailed(w, err)
			return
		}

		decoded, _ := decodeInvoiceAsLndHub(params.Invoice)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			PaymentError    string                 `json:"payment_error"`
			PaymentPreimage Buffer                 `json:"payment_preimage"`
			PaymentRoute    map[string]interface{} `json:"route"`
			PaymentHash     Buffer                 `json:"payment_hash"`
			Decoded         Decoded                `json:"decoded"`
		}{"", "", make(map[string]interface{}), "", decoded})
	})

	router.Path("/balance").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, permission, err := loadUserFromAPICall(r)
		if err != nil {
			errorBadAuth(w)
			return
		}
		if permission < ReadOnlyPermissions {
			errorInsufficientPermissions(w)
			return
		}

		info, err := user.getInfo()
		if err != nil {
			errorInternal(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]map[string]int64{
			"BTC": {
				"AvailableBalance": int64(info.Balance),
			},
		})
	})

	router.Path("/gettxs").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, permission, err := loadUserFromAPICall(r)
		if err != nil {
			errorBadAuth(w)
			return
		}
		if permission < ReadOnlyPermissions {
			errorInsufficientPermissions(w)
			return
		}

		limit, offset := getLimitAndOffset(r)
		txns, err := user.listTransactions(limit, offset, 120, Out)
		if err != nil {
			errorInternal(w)
			return
		}

		type Payment struct {
			PaymentPreimage string  `json:"payment_preimage"`
			Type            string  `json:"type"`
			Fee             float64 `json:"fee"`
			Value           float64 `json:"value"`
			Timestamp       int64   `json:"timestamp"`
			Memo            string  `json:"memo"`
		}

		payments := make([]Payment, len(txns))
		for i, txn := range txns {
			preimage := txn.Preimage.String
			if preimage == "" {
				preimage = "0000000000000000000000000000000000000000000000000000000000000000"
			}

			payments[i] = Payment{
				preimage,
				"paid_invoice",
				txn.Fees,
				-txn.Amount,
				txn.Time.Unix(),
				txn.Description + " " + txn.PeerActionDescription(),
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payments)
	})

	router.Path("/getpending").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]interface{}{})
	})

	router.Path("/getuserinvoices").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, permission, err := loadUserFromAPICall(r)
		if err != nil {
			errorBadAuth(w)
			return
		}
		if permission < ReadOnlyPermissions {
			errorInsufficientPermissions(w)
			return
		}

		limit, offset := getLimitAndOffset(r)
		txns, err := user.listTransactions(limit, offset, 120, In)
		if err != nil {
			errorInternal(w)
			return
		}

		type Inv struct {
			RHash          Buffer  `json:"r_hash"`
			PaymentRequest string  `json:"payment_request"`
			PayReq         string  `json:"pay_req"`
			AddIndex       string  `json:"add_index"`
			Description    string  `json:"description"`
			PaymentHash    string  `json:"payment_hash"`
			IsPaid         bool    `json:"ispaid"`
			Amount         float64 `json:"amt"`
			ExpireTime     float64 `json:"expire_time"`
			Timestamp      int64   `json:"timestamp"`
			Type           string  `json:"type"`
		}

		invs := make([]Inv, len(txns))
		for i, txn := range txns {
			invs[i] = Inv{
				Buffer(txn.Hash),
				"",
				"",
				"1000",
				txn.PeerActionDescription() + txn.Description,
				txn.Hash,
				true,
				txn.Amount,
				float64(s.InvoiceTimeout.Seconds()),
				txn.Time.Unix(),
				"user_invoice",
			}
		}

		iinv, err := rds.Get("justcreatedbluewalletinvoice:" + strconv.Itoa(user.Id)).Result()
		if err == nil {
			var inv map[string]interface{}
			json.Unmarshal([]byte(iinv), &inv)
			invs = append(invs, Inv{
				Buffer(inv["hash"].(string)),
				inv["bolt11"].(string),
				inv["bolt11"].(string),
				"1000",
				inv["desc"].(string),
				inv["hash"].(string),
				false,
				inv["amount"].(float64),
				inv["expiry"].(float64),
				time.Now().Unix(),
				"user_invoice",
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invs)
	})

	router.Path("/decodeinvoice").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bolt11 := r.URL.Query().Get("invoice")

		decoded, err := decodeInvoiceAsLndHub(bolt11)
		if err != nil {
			errorInternal(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(decoded)
	})
}

type Buffer string

func (b Buffer) MarshalJSON() ([]byte, error) {
	arrayBytes, err := hex.DecodeString(string(b))
	if err != nil {
		return nil, err
	}
	arr := make([]int, len(arrayBytes))
	for i, b := range arrayBytes {
		arr[i] = int(b)
	}
	return json.Marshal(map[string]interface{}{
		"type": "Buffer",
		"data": arr,
	})
}

type Decoded struct {
	Destination     string      `json:"destination"`
	PaymentHash     string      `json:"payment_hash"`
	NumSatoshis     string      `json:"num_satoshis"`
	Timestamp       string      `json:"timestamp"`
	Expiry          string      `json:"expiry"`
	Description     string      `json:"description"`
	DescriptionHash string      `json:"description_hash"`
	FallbackAddr    string      `json:"fallback_addr"`
	CLTVExpiry      string      `json:"cltv_expiry"`
	RouteHints      interface{} `json:"route_hints"`
}

func decodeInvoiceAsLndHub(bolt11 string) (Decoded, error) {
	inv, err := ln.Call("decodepay", bolt11)
	if err != nil {
		return Decoded{}, err
	}

	return Decoded{
		Destination:     inv.Get("payee").String(),
		PaymentHash:     inv.Get("payment_hash").String(),
		NumSatoshis:     strconv.Itoa(int(inv.Get("msatoshi").Float() / 1000.0)),
		Timestamp:       inv.Get("created_at").String(),
		Expiry:          inv.Get("expiry").String(),
		Description:     inv.Get("description").String(),
		DescriptionHash: inv.Get("description_hash").String(),
		FallbackAddr:    inv.Get("fallbacks.0.addr").String(),
		CLTVExpiry:      inv.Get("min_final_cltv_expiry").String(),
		RouteHints:      inv.Get("routes").Value(),
	}, nil
}

func getLimitAndOffset(r *http.Request) (limit int, offset int) {
	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 50
	}

	offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))

	return
}
