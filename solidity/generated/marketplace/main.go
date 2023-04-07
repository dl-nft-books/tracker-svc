// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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
)

// IMarketplaceMintedTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceMintedTokenInfo struct {
	TokenId          *big.Int
	MintedTokenPrice *big.Int
	TokenURI         string
}

// IMarketplaceTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceTokenParams struct {
	PricePerOneToken     *big.Int
	MinNFTFloorPrice     *big.Int
	VoucherTokensAmount  *big.Int
	IsNFTBuyable         bool
	VoucherTokenContract common.Address
	FundsRecipient       common.Address
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newBaseTokenContractsURI\",\"type\":\"string\"}],\"name\":\"BaseTokenContractsURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaidTokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintedTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.MintedTokenInfo\",\"name\":\"mintedTokenInfo\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidTokensAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"name\":\"SuccessfullyMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintedTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.MintedTokenInfo\",\"name\":\"mintedTokenInfo\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftFloorPrice\",\"type\":\"uint256\"}],\"name\":\"SuccessfullyMintedByNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractParamsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"__Marketplace_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams_\",\"type\":\"tuple\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenProxy_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenContractsURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentTokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftFloorPrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"buyTokenByNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTokenContractsPart\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"}],\"name\":\"getTokenParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserTokenIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenContractsURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractsRegistry_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"newTokenParams_\",\"type\":\"tuple\"}],\"name\":\"updateAllParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceCaller) BaseTokenContractsURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "baseTokenContractsURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceSession) BaseTokenContractsURI() (string, error) {
	return _Marketplace.Contract.BaseTokenContractsURI(&_Marketplace.CallOpts)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceCallerSession) BaseTokenContractsURI() (string, error) {
	return _Marketplace.Contract.BaseTokenContractsURI(&_Marketplace.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Marketplace *MarketplaceCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Marketplace *MarketplaceSession) GetInjector() (common.Address, error) {
	return _Marketplace.Contract.GetInjector(&_Marketplace.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Marketplace *MarketplaceCallerSession) GetInjector() (common.Address, error) {
	return _Marketplace.Contract.GetInjector(&_Marketplace.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetTokenContractsCount(&_Marketplace.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetTokenContractsCount(&_Marketplace.CallOpts)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCaller) GetTokenContractsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getTokenContractsPart", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetTokenContractsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCallerSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetTokenContractsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetTokenParams is a free data retrieval call binding the contract method 0x70aa6543.
//
// Solidity: function getTokenParams(address tokenContract_) view returns((uint256,uint256,uint256,bool,address,address))
func (_Marketplace *MarketplaceCaller) GetTokenParams(opts *bind.CallOpts, tokenContract_ common.Address) (IMarketplaceTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getTokenParams", tokenContract_)

	if err != nil {
		return *new(IMarketplaceTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IMarketplaceTokenParams)).(*IMarketplaceTokenParams)

	return out0, err

}

// GetTokenParams is a free data retrieval call binding the contract method 0x70aa6543.
//
// Solidity: function getTokenParams(address tokenContract_) view returns((uint256,uint256,uint256,bool,address,address))
func (_Marketplace *MarketplaceSession) GetTokenParams(tokenContract_ common.Address) (IMarketplaceTokenParams, error) {
	return _Marketplace.Contract.GetTokenParams(&_Marketplace.CallOpts, tokenContract_)
}

// GetTokenParams is a free data retrieval call binding the contract method 0x70aa6543.
//
// Solidity: function getTokenParams(address tokenContract_) view returns((uint256,uint256,uint256,bool,address,address))
func (_Marketplace *MarketplaceCallerSession) GetTokenParams(tokenContract_ common.Address) (IMarketplaceTokenParams, error) {
	return _Marketplace.Contract.GetTokenParams(&_Marketplace.CallOpts, tokenContract_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0xe420e11f.
//
// Solidity: function getUserTokenIDs(address tokenContract_, address userAddr_) view returns(uint256[] tokenIDs_)
func (_Marketplace *MarketplaceCaller) GetUserTokenIDs(opts *bind.CallOpts, tokenContract_ common.Address, userAddr_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getUserTokenIDs", tokenContract_, userAddr_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0xe420e11f.
//
// Solidity: function getUserTokenIDs(address tokenContract_, address userAddr_) view returns(uint256[] tokenIDs_)
func (_Marketplace *MarketplaceSession) GetUserTokenIDs(tokenContract_ common.Address, userAddr_ common.Address) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserTokenIDs(&_Marketplace.CallOpts, tokenContract_, userAddr_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0xe420e11f.
//
// Solidity: function getUserTokenIDs(address tokenContract_, address userAddr_) view returns(uint256[] tokenIDs_)
func (_Marketplace *MarketplaceCallerSession) GetUserTokenIDs(tokenContract_ common.Address, userAddr_ common.Address) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserTokenIDs(&_Marketplace.CallOpts, tokenContract_, userAddr_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactor) MarketplaceInit(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "__Marketplace_init", baseTokenContractsURI_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceSession) MarketplaceInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactorSession) MarketplaceInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// AddToken is a paid mutator transaction binding the contract method 0x62b43e96.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,bool,address,address) tokenParams_) returns(address tokenProxy_)
func (_Marketplace *MarketplaceTransactor) AddToken(opts *bind.TransactOpts, name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "addToken", name_, symbol_, tokenParams_)
}

// AddToken is a paid mutator transaction binding the contract method 0x62b43e96.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,bool,address,address) tokenParams_) returns(address tokenProxy_)
func (_Marketplace *MarketplaceSession) AddToken(name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.AddToken(&_Marketplace.TransactOpts, name_, symbol_, tokenParams_)
}

// AddToken is a paid mutator transaction binding the contract method 0x62b43e96.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,bool,address,address) tokenParams_) returns(address tokenProxy_)
func (_Marketplace *MarketplaceTransactorSession) AddToken(name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.AddToken(&_Marketplace.TransactOpts, name_, symbol_, tokenParams_)
}

// BuyToken is a paid mutator transaction binding the contract method 0xdefcaca3.
//
// Solidity: function buyToken(address tokenContract_, uint256 futureTokenId_, address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 discount_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Marketplace *MarketplaceTransactor) BuyToken(opts *bind.TransactOpts, tokenContract_ common.Address, futureTokenId_ *big.Int, paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, discount_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyToken", tokenContract_, futureTokenId_, paymentTokenAddress_, paymentTokenPrice_, discount_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// BuyToken is a paid mutator transaction binding the contract method 0xdefcaca3.
//
// Solidity: function buyToken(address tokenContract_, uint256 futureTokenId_, address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 discount_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Marketplace *MarketplaceSession) BuyToken(tokenContract_ common.Address, futureTokenId_ *big.Int, paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, discount_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyToken(&_Marketplace.TransactOpts, tokenContract_, futureTokenId_, paymentTokenAddress_, paymentTokenPrice_, discount_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// BuyToken is a paid mutator transaction binding the contract method 0xdefcaca3.
//
// Solidity: function buyToken(address tokenContract_, uint256 futureTokenId_, address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 discount_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Marketplace *MarketplaceTransactorSession) BuyToken(tokenContract_ common.Address, futureTokenId_ *big.Int, paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, discount_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyToken(&_Marketplace.TransactOpts, tokenContract_, futureTokenId_, paymentTokenAddress_, paymentTokenPrice_, discount_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// BuyTokenByNFT is a paid mutator transaction binding the contract method 0x40c4fc47.
//
// Solidity: function buyTokenByNFT(address tokenContract_, uint256 futureTokenId_, address nftAddress_, uint256 nftFloorPrice_, uint256 tokenId_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenByNFT(opts *bind.TransactOpts, tokenContract_ common.Address, futureTokenId_ *big.Int, nftAddress_ common.Address, nftFloorPrice_ *big.Int, tokenId_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenByNFT", tokenContract_, futureTokenId_, nftAddress_, nftFloorPrice_, tokenId_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// BuyTokenByNFT is a paid mutator transaction binding the contract method 0x40c4fc47.
//
// Solidity: function buyTokenByNFT(address tokenContract_, uint256 futureTokenId_, address nftAddress_, uint256 nftFloorPrice_, uint256 tokenId_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenByNFT(tokenContract_ common.Address, futureTokenId_ *big.Int, nftAddress_ common.Address, nftFloorPrice_ *big.Int, tokenId_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenByNFT(&_Marketplace.TransactOpts, tokenContract_, futureTokenId_, nftAddress_, nftFloorPrice_, tokenId_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// BuyTokenByNFT is a paid mutator transaction binding the contract method 0x40c4fc47.
//
// Solidity: function buyTokenByNFT(address tokenContract_, uint256 futureTokenId_, address nftAddress_, uint256 nftFloorPrice_, uint256 tokenId_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenByNFT(tokenContract_ common.Address, futureTokenId_ *big.Int, nftAddress_ common.Address, nftFloorPrice_ *big.Int, tokenId_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenByNFT(&_Marketplace.TransactOpts, tokenContract_, futureTokenId_, nftAddress_, nftFloorPrice_, tokenId_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactor) SetBaseTokenContractsURI(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setBaseTokenContractsURI", baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBaseTokenContractsURI(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactorSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBaseTokenContractsURI(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes ) returns()
func (_Marketplace *MarketplaceTransactor) SetDependencies(opts *bind.TransactOpts, contractsRegistry_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setDependencies", contractsRegistry_, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes ) returns()
func (_Marketplace *MarketplaceSession) SetDependencies(contractsRegistry_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SetDependencies(&_Marketplace.TransactOpts, contractsRegistry_, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes ) returns()
func (_Marketplace *MarketplaceTransactorSession) SetDependencies(contractsRegistry_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SetDependencies(&_Marketplace.TransactOpts, contractsRegistry_, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Marketplace *MarketplaceTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Marketplace *MarketplaceSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetInjector(&_Marketplace.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Marketplace *MarketplaceTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetInjector(&_Marketplace.TransactOpts, injector_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x6f579449.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,bool,address,address) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactor) UpdateAllParams(opts *bind.TransactOpts, tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateAllParams", tokenContract_, name_, symbol_, newTokenParams_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x6f579449.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,bool,address,address) newTokenParams_) returns()
func (_Marketplace *MarketplaceSession) UpdateAllParams(tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateAllParams(&_Marketplace.TransactOpts, tokenContract_, name_, symbol_, newTokenParams_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x6f579449.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,bool,address,address) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateAllParams(tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateAllParams(&_Marketplace.TransactOpts, tokenContract_, name_, symbol_, newTokenParams_)
}

// MarketplaceBaseTokenContractsURIUpdatedIterator is returned from FilterBaseTokenContractsURIUpdated and is used to iterate over the raw logs and unpacked data for BaseTokenContractsURIUpdated events raised by the Marketplace contract.
type MarketplaceBaseTokenContractsURIUpdatedIterator struct {
	Event *MarketplaceBaseTokenContractsURIUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBaseTokenContractsURIUpdated)
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
		it.Event = new(MarketplaceBaseTokenContractsURIUpdated)
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
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBaseTokenContractsURIUpdated represents a BaseTokenContractsURIUpdated event raised by the Marketplace contract.
type MarketplaceBaseTokenContractsURIUpdated struct {
	NewBaseTokenContractsURI string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenContractsURIUpdated is a free log retrieval operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) FilterBaseTokenContractsURIUpdated(opts *bind.FilterOpts) (*MarketplaceBaseTokenContractsURIUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceBaseTokenContractsURIUpdatedIterator{contract: _Marketplace.contract, event: "BaseTokenContractsURIUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseTokenContractsURIUpdated is a free log subscription operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) WatchBaseTokenContractsURIUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceBaseTokenContractsURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBaseTokenContractsURIUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
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

// ParseBaseTokenContractsURIUpdated is a log parse operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) ParseBaseTokenContractsURIUpdated(log types.Log) (*MarketplaceBaseTokenContractsURIUpdated, error) {
	event := new(MarketplaceBaseTokenContractsURIUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePaidTokensWithdrawnIterator is returned from FilterPaidTokensWithdrawn and is used to iterate over the raw logs and unpacked data for PaidTokensWithdrawn events raised by the Marketplace contract.
type MarketplacePaidTokensWithdrawnIterator struct {
	Event *MarketplacePaidTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *MarketplacePaidTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePaidTokensWithdrawn)
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
		it.Event = new(MarketplacePaidTokensWithdrawn)
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
func (it *MarketplacePaidTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePaidTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePaidTokensWithdrawn represents a PaidTokensWithdrawn event raised by the Marketplace contract.
type MarketplacePaidTokensWithdrawn struct {
	TokenContract common.Address
	TokenAddr     common.Address
	Recipient     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaidTokensWithdrawn is a free log retrieval operation binding the contract event 0x7f583ae9ce267b66e7f454cadaa0dc67f8262d94d34f68739a53a75ae417a3ea.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenContract, address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) FilterPaidTokensWithdrawn(opts *bind.FilterOpts, tokenContract []common.Address, tokenAddr []common.Address) (*MarketplacePaidTokensWithdrawnIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "PaidTokensWithdrawn", tokenContractRule, tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePaidTokensWithdrawnIterator{contract: _Marketplace.contract, event: "PaidTokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaidTokensWithdrawn is a free log subscription operation binding the contract event 0x7f583ae9ce267b66e7f454cadaa0dc67f8262d94d34f68739a53a75ae417a3ea.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenContract, address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) WatchPaidTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MarketplacePaidTokensWithdrawn, tokenContract []common.Address, tokenAddr []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "PaidTokensWithdrawn", tokenContractRule, tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePaidTokensWithdrawn)
				if err := _Marketplace.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
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

// ParsePaidTokensWithdrawn is a log parse operation binding the contract event 0x7f583ae9ce267b66e7f454cadaa0dc67f8262d94d34f68739a53a75ae417a3ea.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenContract, address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) ParsePaidTokensWithdrawn(log types.Log) (*MarketplacePaidTokensWithdrawn, error) {
	event := new(MarketplacePaidTokensWithdrawn)
	if err := _Marketplace.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceSuccessfullyMintedIterator is returned from FilterSuccessfullyMinted and is used to iterate over the raw logs and unpacked data for SuccessfullyMinted events raised by the Marketplace contract.
type MarketplaceSuccessfullyMintedIterator struct {
	Event *MarketplaceSuccessfullyMinted // Event containing the contract specifics and raw log

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
func (it *MarketplaceSuccessfullyMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceSuccessfullyMinted)
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
		it.Event = new(MarketplaceSuccessfullyMinted)
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
func (it *MarketplaceSuccessfullyMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceSuccessfullyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceSuccessfullyMinted represents a SuccessfullyMinted event raised by the Marketplace contract.
type MarketplaceSuccessfullyMinted struct {
	TokenContract       common.Address
	Recipient           common.Address
	MintedTokenInfo     IMarketplaceMintedTokenInfo
	PaymentTokenAddress common.Address
	PaidTokensAmount    *big.Int
	PaymentTokenPrice   *big.Int
	Discount            *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSuccessfullyMinted is a free log retrieval operation binding the contract event 0xa7f4a6bb14175bd3d0dccff4b599a6ca4b0359311a7f65b6a97b720169cadfe9.
//
// Solidity: event SuccessfullyMinted(address indexed tokenContract, address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed paymentTokenAddress, uint256 paidTokensAmount, uint256 paymentTokenPrice, uint256 discount)
func (_Marketplace *MarketplaceFilterer) FilterSuccessfullyMinted(opts *bind.FilterOpts, tokenContract []common.Address, recipient []common.Address, paymentTokenAddress []common.Address) (*MarketplaceSuccessfullyMintedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var paymentTokenAddressRule []interface{}
	for _, paymentTokenAddressItem := range paymentTokenAddress {
		paymentTokenAddressRule = append(paymentTokenAddressRule, paymentTokenAddressItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "SuccessfullyMinted", tokenContractRule, recipientRule, paymentTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceSuccessfullyMintedIterator{contract: _Marketplace.contract, event: "SuccessfullyMinted", logs: logs, sub: sub}, nil
}

// WatchSuccessfullyMinted is a free log subscription operation binding the contract event 0xa7f4a6bb14175bd3d0dccff4b599a6ca4b0359311a7f65b6a97b720169cadfe9.
//
// Solidity: event SuccessfullyMinted(address indexed tokenContract, address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed paymentTokenAddress, uint256 paidTokensAmount, uint256 paymentTokenPrice, uint256 discount)
func (_Marketplace *MarketplaceFilterer) WatchSuccessfullyMinted(opts *bind.WatchOpts, sink chan<- *MarketplaceSuccessfullyMinted, tokenContract []common.Address, recipient []common.Address, paymentTokenAddress []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var paymentTokenAddressRule []interface{}
	for _, paymentTokenAddressItem := range paymentTokenAddress {
		paymentTokenAddressRule = append(paymentTokenAddressRule, paymentTokenAddressItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "SuccessfullyMinted", tokenContractRule, recipientRule, paymentTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceSuccessfullyMinted)
				if err := _Marketplace.contract.UnpackLog(event, "SuccessfullyMinted", log); err != nil {
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

// ParseSuccessfullyMinted is a log parse operation binding the contract event 0xa7f4a6bb14175bd3d0dccff4b599a6ca4b0359311a7f65b6a97b720169cadfe9.
//
// Solidity: event SuccessfullyMinted(address indexed tokenContract, address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed paymentTokenAddress, uint256 paidTokensAmount, uint256 paymentTokenPrice, uint256 discount)
func (_Marketplace *MarketplaceFilterer) ParseSuccessfullyMinted(log types.Log) (*MarketplaceSuccessfullyMinted, error) {
	event := new(MarketplaceSuccessfullyMinted)
	if err := _Marketplace.contract.UnpackLog(event, "SuccessfullyMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceSuccessfullyMintedByNFTIterator is returned from FilterSuccessfullyMintedByNFT and is used to iterate over the raw logs and unpacked data for SuccessfullyMintedByNFT events raised by the Marketplace contract.
type MarketplaceSuccessfullyMintedByNFTIterator struct {
	Event *MarketplaceSuccessfullyMintedByNFT // Event containing the contract specifics and raw log

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
func (it *MarketplaceSuccessfullyMintedByNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceSuccessfullyMintedByNFT)
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
		it.Event = new(MarketplaceSuccessfullyMintedByNFT)
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
func (it *MarketplaceSuccessfullyMintedByNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceSuccessfullyMintedByNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceSuccessfullyMintedByNFT represents a SuccessfullyMintedByNFT event raised by the Marketplace contract.
type MarketplaceSuccessfullyMintedByNFT struct {
	TokenContract   common.Address
	Recipient       common.Address
	MintedTokenInfo IMarketplaceMintedTokenInfo
	NftAddress      common.Address
	TokenId         *big.Int
	NftFloorPrice   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSuccessfullyMintedByNFT is a free log retrieval operation binding the contract event 0xa4cc88cf83ebd5f0651b1a21f60a45f9a9128452fc1ac5239df062d2c0a6f2cc.
//
// Solidity: event SuccessfullyMintedByNFT(address indexed tokenContract, address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed nftAddress, uint256 tokenId, uint256 nftFloorPrice)
func (_Marketplace *MarketplaceFilterer) FilterSuccessfullyMintedByNFT(opts *bind.FilterOpts, tokenContract []common.Address, recipient []common.Address, nftAddress []common.Address) (*MarketplaceSuccessfullyMintedByNFTIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "SuccessfullyMintedByNFT", tokenContractRule, recipientRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceSuccessfullyMintedByNFTIterator{contract: _Marketplace.contract, event: "SuccessfullyMintedByNFT", logs: logs, sub: sub}, nil
}

// WatchSuccessfullyMintedByNFT is a free log subscription operation binding the contract event 0xa4cc88cf83ebd5f0651b1a21f60a45f9a9128452fc1ac5239df062d2c0a6f2cc.
//
// Solidity: event SuccessfullyMintedByNFT(address indexed tokenContract, address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed nftAddress, uint256 tokenId, uint256 nftFloorPrice)
func (_Marketplace *MarketplaceFilterer) WatchSuccessfullyMintedByNFT(opts *bind.WatchOpts, sink chan<- *MarketplaceSuccessfullyMintedByNFT, tokenContract []common.Address, recipient []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "SuccessfullyMintedByNFT", tokenContractRule, recipientRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceSuccessfullyMintedByNFT)
				if err := _Marketplace.contract.UnpackLog(event, "SuccessfullyMintedByNFT", log); err != nil {
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

// ParseSuccessfullyMintedByNFT is a log parse operation binding the contract event 0xa4cc88cf83ebd5f0651b1a21f60a45f9a9128452fc1ac5239df062d2c0a6f2cc.
//
// Solidity: event SuccessfullyMintedByNFT(address indexed tokenContract, address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed nftAddress, uint256 tokenId, uint256 nftFloorPrice)
func (_Marketplace *MarketplaceFilterer) ParseSuccessfullyMintedByNFT(log types.Log) (*MarketplaceSuccessfullyMintedByNFT, error) {
	event := new(MarketplaceSuccessfullyMintedByNFT)
	if err := _Marketplace.contract.UnpackLog(event, "SuccessfullyMintedByNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenContractDeployedIterator is returned from FilterTokenContractDeployed and is used to iterate over the raw logs and unpacked data for TokenContractDeployed events raised by the Marketplace contract.
type MarketplaceTokenContractDeployedIterator struct {
	Event *MarketplaceTokenContractDeployed // Event containing the contract specifics and raw log

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
func (it *MarketplaceTokenContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenContractDeployed)
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
		it.Event = new(MarketplaceTokenContractDeployed)
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
func (it *MarketplaceTokenContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenContractDeployed represents a TokenContractDeployed event raised by the Marketplace contract.
type MarketplaceTokenContractDeployed struct {
	TokenContract common.Address
	TokenName     string
	TokenSymbol   string
	TokenParams   IMarketplaceTokenParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenContractDeployed is a free log retrieval operation binding the contract event 0x3214c519377ff477fa9e0a8113896e1cef1dbf827e16f5aabf7f4fd26ad5dc6b.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,bool,address,address) tokenParams)
func (_Marketplace *MarketplaceFilterer) FilterTokenContractDeployed(opts *bind.FilterOpts, tokenContract []common.Address) (*MarketplaceTokenContractDeployedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenContractDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenContractDeployedIterator{contract: _Marketplace.contract, event: "TokenContractDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenContractDeployed is a free log subscription operation binding the contract event 0x3214c519377ff477fa9e0a8113896e1cef1dbf827e16f5aabf7f4fd26ad5dc6b.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,bool,address,address) tokenParams)
func (_Marketplace *MarketplaceFilterer) WatchTokenContractDeployed(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenContractDeployed, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenContractDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenContractDeployed)
				if err := _Marketplace.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
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

// ParseTokenContractDeployed is a log parse operation binding the contract event 0x3214c519377ff477fa9e0a8113896e1cef1dbf827e16f5aabf7f4fd26ad5dc6b.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,bool,address,address) tokenParams)
func (_Marketplace *MarketplaceFilterer) ParseTokenContractDeployed(log types.Log) (*MarketplaceTokenContractDeployed, error) {
	event := new(MarketplaceTokenContractDeployed)
	if err := _Marketplace.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenContractParamsUpdatedIterator is returned from FilterTokenContractParamsUpdated and is used to iterate over the raw logs and unpacked data for TokenContractParamsUpdated events raised by the Marketplace contract.
type MarketplaceTokenContractParamsUpdatedIterator struct {
	Event *MarketplaceTokenContractParamsUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceTokenContractParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenContractParamsUpdated)
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
		it.Event = new(MarketplaceTokenContractParamsUpdated)
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
func (it *MarketplaceTokenContractParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenContractParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenContractParamsUpdated represents a TokenContractParamsUpdated event raised by the Marketplace contract.
type MarketplaceTokenContractParamsUpdated struct {
	TokenContract common.Address
	TokenName     string
	TokenSymbol   string
	TokenParams   IMarketplaceTokenParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenContractParamsUpdated is a free log retrieval operation binding the contract event 0x8d8b216c2733ee4b1dac28eff4ceb68c44feb38dc0f75f4e1e6f2a349411be64.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,bool,address,address) tokenParams)
func (_Marketplace *MarketplaceFilterer) FilterTokenContractParamsUpdated(opts *bind.FilterOpts, tokenContract []common.Address) (*MarketplaceTokenContractParamsUpdatedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenContractParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenContractParamsUpdatedIterator{contract: _Marketplace.contract, event: "TokenContractParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenContractParamsUpdated is a free log subscription operation binding the contract event 0x8d8b216c2733ee4b1dac28eff4ceb68c44feb38dc0f75f4e1e6f2a349411be64.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,bool,address,address) tokenParams)
func (_Marketplace *MarketplaceFilterer) WatchTokenContractParamsUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenContractParamsUpdated, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenContractParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenContractParamsUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
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

// ParseTokenContractParamsUpdated is a log parse operation binding the contract event 0x8d8b216c2733ee4b1dac28eff4ceb68c44feb38dc0f75f4e1e6f2a349411be64.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,bool,address,address) tokenParams)
func (_Marketplace *MarketplaceFilterer) ParseTokenContractParamsUpdated(log types.Log) (*MarketplaceTokenContractParamsUpdated, error) {
	event := new(MarketplaceTokenContractParamsUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
