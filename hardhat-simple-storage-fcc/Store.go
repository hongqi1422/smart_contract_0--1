// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_favoriteNumber\",\"type\":\"uint256\"}],\"name\":\"addPerson\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"favoriteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToFavoriteNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"people\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"favoriteNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_favoriteNumber\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506109068061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80632e64cec114610064578063471f7cdf146100825780636057361d146100a05780636f760f41146100bc5780638bab8dd5146100d85780639e7a13ad14610108575b5f80fd5b61006c610139565b60405161007991906102cd565b60405180910390f35b61008a610141565b60405161009791906102cd565b60405180910390f35b6100ba60048036038101906100b59190610321565b610146565b005b6100d660048036038101906100d19190610488565b61014f565b005b6100f260048036038101906100ed91906104e2565b6101d3565b6040516100ff91906102cd565b60405180910390f35b610122600480360381019061011d9190610321565b610200565b604051610130929190610589565b60405180910390f35b5f8054905090565b5f5481565b805f8190555050565b6002604051806040016040528083815260200184815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f015560208201518160010190816101a991906107b1565b505050806001836040516101bd91906108ba565b9081526020016040518091039020819055505050565b6001818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b6002818154811061020f575f80fd5b905f5260205f2090600202015f91509050805f015490806001018054610234906105e4565b80601f0160208091040260200160405190810160405280929190818152602001828054610260906105e4565b80156102ab5780601f10610282576101008083540402835291602001916102ab565b820191905f5260205f20905b81548152906001019060200180831161028e57829003601f168201915b5050505050905082565b5f819050919050565b6102c7816102b5565b82525050565b5f6020820190506102e05f8301846102be565b92915050565b5f604051905090565b5f80fd5b5f80fd5b610300816102b5565b811461030a575f80fd5b50565b5f8135905061031b816102f7565b92915050565b5f60208284031215610336576103356102ef565b5b5f6103438482850161030d565b91505092915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61039a82610354565b810181811067ffffffffffffffff821117156103b9576103b8610364565b5b80604052505050565b5f6103cb6102e6565b90506103d78282610391565b919050565b5f67ffffffffffffffff8211156103f6576103f5610364565b5b6103ff82610354565b9050602081019050919050565b828183375f83830152505050565b5f61042c610427846103dc565b6103c2565b90508281526020810184848401111561044857610447610350565b5b61045384828561040c565b509392505050565b5f82601f83011261046f5761046e61034c565b5b813561047f84826020860161041a565b91505092915050565b5f806040838503121561049e5761049d6102ef565b5b5f83013567ffffffffffffffff8111156104bb576104ba6102f3565b5b6104c78582860161045b565b92505060206104d88582860161030d565b9150509250929050565b5f602082840312156104f7576104f66102ef565b5b5f82013567ffffffffffffffff811115610514576105136102f3565b5b6105208482850161045b565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61055b82610529565b6105658185610533565b9350610575818560208601610543565b61057e81610354565b840191505092915050565b5f60408201905061059c5f8301856102be565b81810360208301526105ae8184610551565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806105fb57607f821691505b60208210810361060e5761060d6105b7565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106707fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610635565b61067a8683610635565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6106b56106b06106ab846102b5565b610692565b6102b5565b9050919050565b5f819050919050565b6106ce8361069b565b6106e26106da826106bc565b848454610641565b825550505050565b5f90565b6106f66106ea565b6107018184846106c5565b505050565b5b81811015610724576107195f826106ee565b600181019050610707565b5050565b601f8211156107695761073a81610614565b61074384610626565b81016020851015610752578190505b61076661075e85610626565b830182610706565b50505b505050565b5f82821c905092915050565b5f6107895f198460080261076e565b1980831691505092915050565b5f6107a1838361077a565b9150826002028217905092915050565b6107ba82610529565b67ffffffffffffffff8111156107d3576107d2610364565b5b6107dd82546105e4565b6107e8828285610728565b5f60209050601f831160018114610819575f8415610807578287015190505b6108118582610796565b865550610878565b601f19841661082786610614565b5f5b8281101561084e57848901518255600182019150602085019450602081019050610829565b8683101561086b5784890151610867601f89168261077a565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f61089482610529565b61089e8185610880565b93506108ae818560208601610543565b80840191505092915050565b5f6108c5828461088a565b91508190509291505056fea2646970667358221220af980864c14189db23dbbd198966455109f0ae3e61acc0dec095f9d463a818c264736f6c63430008190033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// FavoriteNumber is a free data retrieval call binding the contract method 0x471f7cdf.
//
// Solidity: function favoriteNumber() view returns(uint256)
func (_Store *StoreCaller) FavoriteNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "favoriteNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FavoriteNumber is a free data retrieval call binding the contract method 0x471f7cdf.
//
// Solidity: function favoriteNumber() view returns(uint256)
func (_Store *StoreSession) FavoriteNumber() (*big.Int, error) {
	return _Store.Contract.FavoriteNumber(&_Store.CallOpts)
}

// FavoriteNumber is a free data retrieval call binding the contract method 0x471f7cdf.
//
// Solidity: function favoriteNumber() view returns(uint256)
func (_Store *StoreCallerSession) FavoriteNumber() (*big.Int, error) {
	return _Store.Contract.FavoriteNumber(&_Store.CallOpts)
}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_Store *StoreCaller) NameToFavoriteNumber(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "nameToFavoriteNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_Store *StoreSession) NameToFavoriteNumber(arg0 string) (*big.Int, error) {
	return _Store.Contract.NameToFavoriteNumber(&_Store.CallOpts, arg0)
}

// NameToFavoriteNumber is a free data retrieval call binding the contract method 0x8bab8dd5.
//
// Solidity: function nameToFavoriteNumber(string ) view returns(uint256)
func (_Store *StoreCallerSession) NameToFavoriteNumber(arg0 string) (*big.Int, error) {
	return _Store.Contract.NameToFavoriteNumber(&_Store.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_Store *StoreCaller) People(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "people", arg0)

	outstruct := new(struct {
		FavoriteNumber *big.Int
		Name           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FavoriteNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_Store *StoreSession) People(arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	return _Store.Contract.People(&_Store.CallOpts, arg0)
}

// People is a free data retrieval call binding the contract method 0x9e7a13ad.
//
// Solidity: function people(uint256 ) view returns(uint256 favoriteNumber, string name)
func (_Store *StoreCallerSession) People(arg0 *big.Int) (struct {
	FavoriteNumber *big.Int
	Name           string
}, error) {
	return _Store.Contract.People(&_Store.CallOpts, arg0)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Store *StoreCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Store *StoreSession) Retrieve() (*big.Int, error) {
	return _Store.Contract.Retrieve(&_Store.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Store *StoreCallerSession) Retrieve() (*big.Int, error) {
	return _Store.Contract.Retrieve(&_Store.CallOpts)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_Store *StoreTransactor) AddPerson(opts *bind.TransactOpts, _name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addPerson", _name, _favoriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_Store *StoreSession) AddPerson(_name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Store.Contract.AddPerson(&_Store.TransactOpts, _name, _favoriteNumber)
}

// AddPerson is a paid mutator transaction binding the contract method 0x6f760f41.
//
// Solidity: function addPerson(string _name, uint256 _favoriteNumber) returns()
func (_Store *StoreTransactorSession) AddPerson(_name string, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Store.Contract.AddPerson(&_Store.TransactOpts, _name, _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Store *StoreTransactor) Store(opts *bind.TransactOpts, _favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "store", _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Store *StoreSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Store(&_Store.TransactOpts, _favoriteNumber)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 _favoriteNumber) returns()
func (_Store *StoreTransactorSession) Store(_favoriteNumber *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Store(&_Store.TransactOpts, _favoriteNumber)
}
