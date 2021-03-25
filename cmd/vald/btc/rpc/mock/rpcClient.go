// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/cmd/vald/btc/rpc"
	"github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"sync"
)

// Ensure, that ClientMock does implement rpc.Client.
// If this is not the case, regenerate this file with moq.
var _ rpc.Client = &ClientMock{}

// ClientMock is a mock implementation of rpc.Client.
//
// 	func TestSomethingThatUsesClient(t *testing.T) {
//
// 		// make and configure a mocked rpc.Client
// 		mockedClient := &ClientMock{
// 			GetTxOutFunc: func(txHash *chainhash.Hash, voutIdx uint32, mempool bool) (*btcjson.GetTxOutResult, error) {
// 				panic("mock out the GetTxOut method")
// 			},
// 			NetworkFunc: func() types.Network {
// 				panic("mock out the Network method")
// 			},
// 			SendRawTransactionFunc: func(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
// 				panic("mock out the SendRawTransaction method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires rpc.Client
// 		// and then make assertions.
//
// 	}
type ClientMock struct {
	// GetTxOutFunc mocks the GetTxOut method.
	GetTxOutFunc func(txHash *chainhash.Hash, voutIdx uint32, mempool bool) (*btcjson.GetTxOutResult, error)

	// NetworkFunc mocks the Network method.
	NetworkFunc func() types.Network

	// SendRawTransactionFunc mocks the SendRawTransaction method.
	SendRawTransactionFunc func(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetTxOut holds details about calls to the GetTxOut method.
		GetTxOut []struct {
			// TxHash is the txHash argument value.
			TxHash *chainhash.Hash
			// VoutIdx is the voutIdx argument value.
			VoutIdx uint32
			// Mempool is the mempool argument value.
			Mempool bool
		}
		// Network holds details about calls to the Network method.
		Network []struct {
		}
		// SendRawTransaction holds details about calls to the SendRawTransaction method.
		SendRawTransaction []struct {
			// Tx is the tx argument value.
			Tx *wire.MsgTx
			// AllowHighFees is the allowHighFees argument value.
			AllowHighFees bool
		}
	}
	lockGetTxOut           sync.RWMutex
	lockNetwork            sync.RWMutex
	lockSendRawTransaction sync.RWMutex
}

// GetTxOut calls GetTxOutFunc.
func (mock *ClientMock) GetTxOut(txHash *chainhash.Hash, voutIdx uint32, mempool bool) (*btcjson.GetTxOutResult, error) {
	if mock.GetTxOutFunc == nil {
		panic("ClientMock.GetTxOutFunc: method is nil but Client.GetTxOut was just called")
	}
	callInfo := struct {
		TxHash  *chainhash.Hash
		VoutIdx uint32
		Mempool bool
	}{
		TxHash:  txHash,
		VoutIdx: voutIdx,
		Mempool: mempool,
	}
	mock.lockGetTxOut.Lock()
	mock.calls.GetTxOut = append(mock.calls.GetTxOut, callInfo)
	mock.lockGetTxOut.Unlock()
	return mock.GetTxOutFunc(txHash, voutIdx, mempool)
}

// GetTxOutCalls gets all the calls that were made to GetTxOut.
// Check the length with:
//     len(mockedClient.GetTxOutCalls())
func (mock *ClientMock) GetTxOutCalls() []struct {
	TxHash  *chainhash.Hash
	VoutIdx uint32
	Mempool bool
} {
	var calls []struct {
		TxHash  *chainhash.Hash
		VoutIdx uint32
		Mempool bool
	}
	mock.lockGetTxOut.RLock()
	calls = mock.calls.GetTxOut
	mock.lockGetTxOut.RUnlock()
	return calls
}

// Network calls NetworkFunc.
func (mock *ClientMock) Network() types.Network {
	if mock.NetworkFunc == nil {
		panic("ClientMock.NetworkFunc: method is nil but Client.Network was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNetwork.Lock()
	mock.calls.Network = append(mock.calls.Network, callInfo)
	mock.lockNetwork.Unlock()
	return mock.NetworkFunc()
}

// NetworkCalls gets all the calls that were made to Network.
// Check the length with:
//     len(mockedClient.NetworkCalls())
func (mock *ClientMock) NetworkCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNetwork.RLock()
	calls = mock.calls.Network
	mock.lockNetwork.RUnlock()
	return calls
}

// SendRawTransaction calls SendRawTransactionFunc.
func (mock *ClientMock) SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
	if mock.SendRawTransactionFunc == nil {
		panic("ClientMock.SendRawTransactionFunc: method is nil but Client.SendRawTransaction was just called")
	}
	callInfo := struct {
		Tx            *wire.MsgTx
		AllowHighFees bool
	}{
		Tx:            tx,
		AllowHighFees: allowHighFees,
	}
	mock.lockSendRawTransaction.Lock()
	mock.calls.SendRawTransaction = append(mock.calls.SendRawTransaction, callInfo)
	mock.lockSendRawTransaction.Unlock()
	return mock.SendRawTransactionFunc(tx, allowHighFees)
}

// SendRawTransactionCalls gets all the calls that were made to SendRawTransaction.
// Check the length with:
//     len(mockedClient.SendRawTransactionCalls())
func (mock *ClientMock) SendRawTransactionCalls() []struct {
	Tx            *wire.MsgTx
	AllowHighFees bool
} {
	var calls []struct {
		Tx            *wire.MsgTx
		AllowHighFees bool
	}
	mock.lockSendRawTransaction.RLock()
	calls = mock.calls.SendRawTransaction
	mock.lockSendRawTransaction.RUnlock()
	return calls
}