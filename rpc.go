package omnilayer

import (
	"github.com/iuouiyiuty/omnilayer-go/omnijson"
	"github.com/iuouiyiuty/omnilayer-go/btcjson"
)

func (c *Client) GetBlockChainInfo() (btcjson.GetBlockChainInfoResult, error) {
	return futureGetBlockChainInfo(c.do(btcjson.GetBlockChainInfoCommand{})).Receive()
}

func (c *Client) OmniListBlockTransactions(block int64) (omnijson.OmniListBlockTransactionsResult, error) {
	return futureOmniListBlockTransactions(c.do(omnijson.OmniListBlockTransactionsCommand{
		Block: block,
	})).Receive()
}

func (c *Client) GetInfo() (omnijson.OmniGetInfoResult, error) {
	return futureGetInfo(c.do(omnijson.OmniGetInfoCommand{})).Receive()
}

func (c *Client) OmniGetTransaction(hash string) (omnijson.OmniGettransactionResult, error) {
	return futureOmniGetTransaction(c.do(omnijson.OmniGetTransactionCommand{
		Hash: hash,
	})).Receive()
}

func (c *Client) ListUnspent(cmd btcjson.ListUnspentCommand) (btcjson.ListUnspentResult, error) {
	return futureListUnspent(c.do(cmd)).Receive()
}

func (c *Client) OmniCreatePayloadSimpleSend(cmd omnijson.OmniCreatePayloadSimpleSendCommand) (omnijson.OmniCreatePayloadSimpleSendResult, error) {
	return futureOmniCreatePayloadSimpleSend(c.do(cmd)).Receive()
}

func (c *Client) CreateRawTransaction(cmd btcjson.CreateRawTransactionCommand) (btcjson.CreateRawTransactionResult, error) {
	return futureCreateRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxOpReturn(cmd omnijson.OmniCreateRawTxOpReturnCommand) (omnijson.OmniCreateRawTxOpReturnResult, error) {
	return futureOmniCreateRawTxOpReturn(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxReference(cmd omnijson.OmniCreateRawTxReferenceCommand) (omnijson.OmniCreateRawTxReferenceResult, error) {
	return futureOmniCreateRawTxReference(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxChange(cmd omnijson.OmniCreateRawTxChangeCommand) (omnijson.OmniCreateRawTxChangeResult, error) {
	return futureOmniCreateRawTxChange(c.do(cmd)).Receive()
}

func (c *Client) ImportAddress(address string, rescan bool) error {
	return futureImportAddress(c.do(btcjson.ImportAddressCommand{Adress: address, Rescan: rescan})).Receive()
}

func (c *Client) SendRawTransaction(cmd btcjson.SendRawTransactionCommand) (btcjson.SendRawTransactionResult, error) {
	return futureSendRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransaction(cmd btcjson.SignRawTransactionCommand) (btcjson.SignRawTransactionResult, error) {
	return futureSignRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransactionWithKey(cmd btcjson.SignRawTransactionWithKeyCommand) (btcjson.SignRawTransactionWithKeyResult, error) {
	return futureSignRawTransactionWithKey(c.do(cmd)).Receive()
}

func (c *Client) OmniGetBalance(cmd omnijson.OmniGetBalanceCommand) (omnijson.OmniGetBalanceResult, error) {
	return futureOmniGetBalance(c.do(cmd)).Receive()
}
