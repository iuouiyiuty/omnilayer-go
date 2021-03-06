package omnilayer

import (
	"encoding/json"

	"github.com/iuouiyiuty/omnilayer-go/btcjson"
)

type futureCreateRawTransaction chan *response

func (f futureCreateRawTransaction) Receive() (btcjson.CreateRawTransactionResult, error) {
	var result btcjson.CreateRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetBlockChainInfo chan *response

func (f futureGetBlockChainInfo) Receive() (btcjson.GetBlockChainInfoResult, error) {
	var result btcjson.GetBlockChainInfoResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureListUnspent chan *response

func (f futureListUnspent) Receive() (btcjson.ListUnspentResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make(btcjson.ListUnspentResult, 0)

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureImportAddress chan *response

func (f futureImportAddress) Receive() error {
	_, err := receive(f)
	return err
}

type futureSendRawTransaction chan *response

func (f futureSendRawTransaction) Receive() (btcjson.SendRawTransactionResult, error) {
	var result btcjson.SendRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureSignRawTransaction chan *response

func (f futureSignRawTransaction) Receive() (btcjson.SignRawTransactionResult, error) {
	var result btcjson.SignRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureSignRawTransactionWithKey chan *response

func (f futureSignRawTransactionWithKey) Receive() (btcjson.SignRawTransactionWithKeyResult, error) {
	var result btcjson.SignRawTransactionWithKeyResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}
