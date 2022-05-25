// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	abcitypes "github.com/tendermint/tendermint/abci/types"

	mempool "github.com/tendermint/tendermint/internal/mempool"

	mock "github.com/stretchr/testify/mock"

	types "github.com/tendermint/tendermint/types"
)

// Mempool is an autogenerated mock type for the Mempool type
type Mempool struct {
	mock.Mock
}

// CheckTx provides a mock function with given fields: ctx, tx, callback, txInfo
func (_m *Mempool) CheckTx(ctx context.Context, tx types.Tx, callback func(*abcitypes.ResponseCheckTx), txInfo mempool.TxInfo) error {
	ret := _m.Called(ctx, tx, callback, txInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx, func(*abcitypes.ResponseCheckTx), mempool.TxInfo) error); ok {
		r0 = rf(ctx, tx, callback, txInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableTxsAvailable provides a mock function with given fields:
func (_m *Mempool) EnableTxsAvailable() {
	_m.Called()
}

// Flush provides a mock function with given fields:
func (_m *Mempool) Flush() {
	_m.Called()
}

// FlushAppConn provides a mock function with given fields: _a0
func (_m *Mempool) FlushAppConn(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Lock provides a mock function with given fields:
func (_m *Mempool) Lock() {
	_m.Called()
}

// ReapMaxBytesMaxGas provides a mock function with given fields: maxBytes, maxGas
func (_m *Mempool) ReapMaxBytesMaxGas(maxBytes int64, maxGas int64) types.Txs {
	ret := _m.Called(maxBytes, maxGas)

	var r0 types.Txs
	if rf, ok := ret.Get(0).(func(int64, int64) types.Txs); ok {
		r0 = rf(maxBytes, maxGas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Txs)
		}
	}

	return r0
}

// ReapMaxTxs provides a mock function with given fields: max
func (_m *Mempool) ReapMaxTxs(max int) types.Txs {
	ret := _m.Called(max)

	var r0 types.Txs
	if rf, ok := ret.Get(0).(func(int) types.Txs); ok {
		r0 = rf(max)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Txs)
		}
	}

	return r0
}

// RemoveTxByKey provides a mock function with given fields: txKey
func (_m *Mempool) RemoveTxByKey(txKey types.TxKey) error {
	ret := _m.Called(txKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.TxKey) error); ok {
		r0 = rf(txKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Size provides a mock function with given fields:
func (_m *Mempool) Size() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SizeBytes provides a mock function with given fields:
func (_m *Mempool) SizeBytes() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// TxsAvailable provides a mock function with given fields:
func (_m *Mempool) TxsAvailable() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Unlock provides a mock function with given fields:
func (_m *Mempool) Unlock() {
	_m.Called()
}

// Update provides a mock function with given fields: ctx, blockHeight, blockTxs, txResults, newPreFn, newPostFn, recheck
func (_m *Mempool) Update(ctx context.Context, blockHeight int64, blockTxs types.Txs, txResults []*abcitypes.ExecTxResult, newPreFn mempool.PreCheckFunc, newPostFn mempool.PostCheckFunc, recheck bool) error {
	ret := _m.Called(ctx, blockHeight, blockTxs, txResults, newPreFn, newPostFn, recheck)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, types.Txs, []*abcitypes.ExecTxResult, mempool.PreCheckFunc, mempool.PostCheckFunc, bool) error); ok {
		r0 = rf(ctx, blockHeight, blockTxs, txResults, newPreFn, newPostFn, recheck)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewMempoolT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMempool creates a new instance of Mempool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMempool(t NewMempoolT) *Mempool {
	mock := &Mempool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
