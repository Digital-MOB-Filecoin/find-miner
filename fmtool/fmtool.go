package fmtool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type Config struct {
	RsvAPI string
}

type WorkerLib struct {
	size        int64
	region      string
	verifiedSPL int64
	skip        int64
	config      Config
}

func NewWorkerLib(size int64, region string, verifiedSPL int64, skip int64, config Config) *WorkerLib {
	return &WorkerLib{
		size,
		region,
		verifiedSPL,
		skip,
		config,
	}
}

func (w *WorkerLib) Run() error {
	// params:
	//   --size=N
	//   --region=[ap|cn|na|eu]
	//   --verified-retrieval-price-limit
	//   --skip-miners=N

	var size string
	var region string
	var verifiedSPL string
	var skip string

	if w.size <= 0 {
		size = "null"
	} else {
		size = strconv.FormatInt(w.size, 10)
	}

	if len(w.region) == 0 {
		region = "null"
	} else {
		region = "\"" + w.region + "\""
	}

	if w.verifiedSPL < 0 {
		verifiedSPL = "null"
	} else {
		verifiedSPL = strconv.FormatInt(w.verifiedSPL, 10)
	}

	if w.skip <= 0 {
		skip = "null"
	} else {
		skip = strconv.FormatInt(w.skip, 10)
	}

	postJson := fmt.Sprintf(`{"jsonrpc": "2.0", "method": "miners.find", "id": 1, "params": [%s,%s,%s,%s]}`,
		size,
		region,
		verifiedSPL,
		skip)

	request := gorequest.New()
	resp, body, errs := request.Post(w.config.RsvAPI).
		Set("Content-Type", "application/json").
		Send(postJson).
		End()

	if errs != nil || (resp.StatusCode != http.StatusOK) {
		return errors.New("target returned response code != 200")
	}

	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	if data["result"] != nil {
		fmt.Println(data["result"])
	}

	return nil
}
