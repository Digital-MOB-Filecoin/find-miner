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
	TargetURL string
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

	postJson := `{"jsonrpc": "2.0", "method": "miners.find", "id": 1, "params": [`
	if w.size <= 0 {
		postJson += "null,"
	} else {
		postJson += strconv.FormatInt(w.size, 10)
		postJson += ","
	}
	if len(w.region) == 0 {
		postJson += "null,"
	} else {
		postJson += "\"" + w.region + "\""
		postJson += ","
	}
	if w.verifiedSPL < 0 {
		postJson += "null,"
	} else {
		postJson += strconv.FormatInt(w.verifiedSPL, 10)
		postJson += ","
	}
	if w.skip <= 0 {
		postJson += "null"
	} else {
		postJson += strconv.FormatInt(w.skip, 10)
	}
	postJson += "]}"

	request := gorequest.New()
	resp, body, errs := request.Post(w.config.TargetURL).
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
	fmt.Println(data["result"])

	return nil
}
