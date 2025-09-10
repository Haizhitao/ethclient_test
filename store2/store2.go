// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Store2MetaData contains all meta data concerning the Store2 contract.
var Store2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Store2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Store2MetaData.ABI instead.
var Store2ABI = Store2MetaData.ABI

// Store2 is an auto generated Go binding around an Ethereum contract.
type Store2 struct {
	Store2Caller     // Read-only binding to the contract
	Store2Transactor // Write-only binding to the contract
	Store2Filterer   // Log filterer for contract events
}

// Store2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Store2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Store2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Store2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Store2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Store2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Store2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Store2Session struct {
	Contract     *Store2           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Store2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Store2CallerSession struct {
	Contract *Store2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Store2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Store2TransactorSession struct {
	Contract     *Store2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Store2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Store2Raw struct {
	Contract *Store2 // Generic contract binding to access the raw methods on
}

// Store2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Store2CallerRaw struct {
	Contract *Store2Caller // Generic read-only contract binding to access the raw methods on
}

// Store2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Store2TransactorRaw struct {
	Contract *Store2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStore2 creates a new instance of Store2, bound to a specific deployed contract.
func NewStore2(address common.Address, backend bind.ContractBackend) (*Store2, error) {
	contract, err := bindStore2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store2{Store2Caller: Store2Caller{contract: contract}, Store2Transactor: Store2Transactor{contract: contract}, Store2Filterer: Store2Filterer{contract: contract}}, nil
}

// NewStore2Caller creates a new read-only instance of Store2, bound to a specific deployed contract.
func NewStore2Caller(address common.Address, caller bind.ContractCaller) (*Store2Caller, error) {
	contract, err := bindStore2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Store2Caller{contract: contract}, nil
}

// NewStore2Transactor creates a new write-only instance of Store2, bound to a specific deployed contract.
func NewStore2Transactor(address common.Address, transactor bind.ContractTransactor) (*Store2Transactor, error) {
	contract, err := bindStore2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Store2Transactor{contract: contract}, nil
}

// NewStore2Filterer creates a new log filterer instance of Store2, bound to a specific deployed contract.
func NewStore2Filterer(address common.Address, filterer bind.ContractFilterer) (*Store2Filterer, error) {
	contract, err := bindStore2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Store2Filterer{contract: contract}, nil
}

// bindStore2 binds a generic wrapper to an already deployed contract.
func bindStore2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Store2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store2 *Store2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store2.Contract.Store2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store2 *Store2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store2.Contract.Store2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store2 *Store2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store2.Contract.Store2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store2 *Store2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store2 *Store2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store2 *Store2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store2.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store2 *Store2Caller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Store2.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store2 *Store2Session) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store2.Contract.Items(&_Store2.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store2 *Store2CallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store2.Contract.Items(&_Store2.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store2 *Store2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store2 *Store2Session) Version() (string, error) {
	return _Store2.Contract.Version(&_Store2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store2 *Store2CallerSession) Version() (string, error) {
	return _Store2.Contract.Version(&_Store2.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store2 *Store2Transactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store2.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store2 *Store2Session) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store2.Contract.SetItem(&_Store2.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store2 *Store2TransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store2.Contract.SetItem(&_Store2.TransactOpts, key, value)
}

// Store2ItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Store2 contract.
type Store2ItemSetIterator struct {
	Event *Store2ItemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Store2ItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Store2ItemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Store2ItemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Store2ItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Store2ItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Store2ItemSet represents a ItemSet event raised by the Store2 contract.
type Store2ItemSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store2 *Store2Filterer) FilterItemSet(opts *bind.FilterOpts) (*Store2ItemSetIterator, error) {

	logs, sub, err := _Store2.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &Store2ItemSetIterator{contract: _Store2.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store2 *Store2Filterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *Store2ItemSet) (event.Subscription, error) {

	logs, sub, err := _Store2.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Store2ItemSet)
				if err := _Store2.contract.UnpackLog(event, "ItemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseItemSet is a log parse operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store2 *Store2Filterer) ParseItemSet(log types.Log) (*Store2ItemSet, error) {
	event := new(Store2ItemSet)
	if err := _Store2.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
