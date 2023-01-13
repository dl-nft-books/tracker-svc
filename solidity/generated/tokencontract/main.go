// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokencontract

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

// ITokenContractMintedTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type ITokenContractMintedTokenInfo struct {
	TokenId          *big.Int
	PricePerOneToken *big.Int
	TokenURI         string
}

// TokencontractMetaData contains all meta data concerning the Tokencontract contract.
var TokencontractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaidTokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structITokenContract.MintedTokenInfo\",\"name\":\"mintedTokenInfo\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidTokensAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"name\":\"SuccessfullyMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"TokenContractParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVoucherTokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVoucherTokensAmount\",\"type\":\"uint256\"}],\"name\":\"VoucherParamsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenFactoryAddr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount_\",\"type\":\"uint256\"}],\"name\":\"__TokenContract_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"existingTokenURIs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserTokenIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"contractITokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newVoucherTokenContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newVoucherTokensAmount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newTokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newTokenSymbol_\",\"type\":\"string\"}],\"name\":\"updateAllParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newTokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newTokenSymbol_\",\"type\":\"string\"}],\"name\":\"updateTokenContractParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVoucherTokenContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newVoucherTokensAmount_\",\"type\":\"uint256\"}],\"name\":\"updateVoucherParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voucherTokenContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voucherTokensAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"}],\"name\":\"withdrawPaidTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokencontractABI is the input ABI used to generate the binding from.
// Deprecated: Use TokencontractMetaData.ABI instead.
var TokencontractABI = TokencontractMetaData.ABI

// Tokencontract is an auto generated Go binding around an Ethereum contract.
type Tokencontract struct {
	TokencontractCaller     // Read-only binding to the contract
	TokencontractTransactor // Write-only binding to the contract
	TokencontractFilterer   // Log filterer for contract events
}

// TokencontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokencontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokencontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokencontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokencontractSession struct {
	Contract     *Tokencontract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokencontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokencontractCallerSession struct {
	Contract *TokencontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokencontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokencontractTransactorSession struct {
	Contract     *TokencontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokencontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokencontractRaw struct {
	Contract *Tokencontract // Generic contract binding to access the raw methods on
}

// TokencontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokencontractCallerRaw struct {
	Contract *TokencontractCaller // Generic read-only contract binding to access the raw methods on
}

// TokencontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokencontractTransactorRaw struct {
	Contract *TokencontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokencontract creates a new instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontract(address common.Address, backend bind.ContractBackend) (*Tokencontract, error) {
	contract, err := bindTokencontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokencontract{TokencontractCaller: TokencontractCaller{contract: contract}, TokencontractTransactor: TokencontractTransactor{contract: contract}, TokencontractFilterer: TokencontractFilterer{contract: contract}}, nil
}

// NewTokencontractCaller creates a new read-only instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontractCaller(address common.Address, caller bind.ContractCaller) (*TokencontractCaller, error) {
	contract, err := bindTokencontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokencontractCaller{contract: contract}, nil
}

// NewTokencontractTransactor creates a new write-only instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontractTransactor(address common.Address, transactor bind.ContractTransactor) (*TokencontractTransactor, error) {
	contract, err := bindTokencontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokencontractTransactor{contract: contract}, nil
}

// NewTokencontractFilterer creates a new log filterer instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontractFilterer(address common.Address, filterer bind.ContractFilterer) (*TokencontractFilterer, error) {
	contract, err := bindTokencontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokencontractFilterer{contract: contract}, nil
}

// bindTokencontract binds a generic wrapper to an already deployed contract.
func bindTokencontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokencontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokencontract *TokencontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokencontract.Contract.TokencontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokencontract *TokencontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.Contract.TokencontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokencontract *TokencontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokencontract.Contract.TokencontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokencontract *TokencontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokencontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokencontract *TokencontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokencontract *TokencontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokencontract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Tokencontract *TokencontractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Tokencontract *TokencontractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.BalanceOf(&_Tokencontract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.BalanceOf(&_Tokencontract.CallOpts, owner)
}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string ) view returns(bool)
func (_Tokencontract *TokencontractCaller) ExistingTokenURIs(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "existingTokenURIs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string ) view returns(bool)
func (_Tokencontract *TokencontractSession) ExistingTokenURIs(arg0 string) (bool, error) {
	return _Tokencontract.Contract.ExistingTokenURIs(&_Tokencontract.CallOpts, arg0)
}

// ExistingTokenURIs is a free data retrieval call binding the contract method 0xb8f4b821.
//
// Solidity: function existingTokenURIs(string ) view returns(bool)
func (_Tokencontract *TokencontractCallerSession) ExistingTokenURIs(arg0 string) (bool, error) {
	return _Tokencontract.Contract.ExistingTokenURIs(&_Tokencontract.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Tokencontract *TokencontractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Tokencontract *TokencontractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Tokencontract.Contract.GetApproved(&_Tokencontract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Tokencontract *TokencontractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Tokencontract.Contract.GetApproved(&_Tokencontract.CallOpts, tokenId)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_Tokencontract *TokencontractCaller) GetUserTokenIDs(opts *bind.CallOpts, userAddr_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "getUserTokenIDs", userAddr_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_Tokencontract *TokencontractSession) GetUserTokenIDs(userAddr_ common.Address) ([]*big.Int, error) {
	return _Tokencontract.Contract.GetUserTokenIDs(&_Tokencontract.CallOpts, userAddr_)
}

// GetUserTokenIDs is a free data retrieval call binding the contract method 0x595ee598.
//
// Solidity: function getUserTokenIDs(address userAddr_) view returns(uint256[] tokenIDs_)
func (_Tokencontract *TokencontractCallerSession) GetUserTokenIDs(userAddr_ common.Address) ([]*big.Int, error) {
	return _Tokencontract.Contract.GetUserTokenIDs(&_Tokencontract.CallOpts, userAddr_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Tokencontract *TokencontractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Tokencontract *TokencontractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Tokencontract.Contract.IsApprovedForAll(&_Tokencontract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Tokencontract *TokencontractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Tokencontract.Contract.IsApprovedForAll(&_Tokencontract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokencontract *TokencontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokencontract *TokencontractSession) Name() (string, error) {
	return _Tokencontract.Contract.Name(&_Tokencontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokencontract *TokencontractCallerSession) Name() (string, error) {
	return _Tokencontract.Contract.Name(&_Tokencontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokencontract *TokencontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokencontract *TokencontractSession) Owner() (common.Address, error) {
	return _Tokencontract.Contract.Owner(&_Tokencontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokencontract *TokencontractCallerSession) Owner() (common.Address, error) {
	return _Tokencontract.Contract.Owner(&_Tokencontract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Tokencontract *TokencontractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Tokencontract *TokencontractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Tokencontract.Contract.OwnerOf(&_Tokencontract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Tokencontract *TokencontractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Tokencontract.Contract.OwnerOf(&_Tokencontract.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tokencontract *TokencontractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tokencontract *TokencontractSession) Paused() (bool, error) {
	return _Tokencontract.Contract.Paused(&_Tokencontract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tokencontract *TokencontractCallerSession) Paused() (bool, error) {
	return _Tokencontract.Contract.Paused(&_Tokencontract.CallOpts)
}

// PricePerOneToken is a free data retrieval call binding the contract method 0x1935fba9.
//
// Solidity: function pricePerOneToken() view returns(uint256)
func (_Tokencontract *TokencontractCaller) PricePerOneToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "pricePerOneToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerOneToken is a free data retrieval call binding the contract method 0x1935fba9.
//
// Solidity: function pricePerOneToken() view returns(uint256)
func (_Tokencontract *TokencontractSession) PricePerOneToken() (*big.Int, error) {
	return _Tokencontract.Contract.PricePerOneToken(&_Tokencontract.CallOpts)
}

// PricePerOneToken is a free data retrieval call binding the contract method 0x1935fba9.
//
// Solidity: function pricePerOneToken() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) PricePerOneToken() (*big.Int, error) {
	return _Tokencontract.Contract.PricePerOneToken(&_Tokencontract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tokencontract *TokencontractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tokencontract *TokencontractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tokencontract.Contract.SupportsInterface(&_Tokencontract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tokencontract *TokencontractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tokencontract.Contract.SupportsInterface(&_Tokencontract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokencontract *TokencontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokencontract *TokencontractSession) Symbol() (string, error) {
	return _Tokencontract.Contract.Symbol(&_Tokencontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokencontract *TokencontractCallerSession) Symbol() (string, error) {
	return _Tokencontract.Contract.Symbol(&_Tokencontract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Tokencontract *TokencontractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Tokencontract *TokencontractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Tokencontract.Contract.TokenByIndex(&_Tokencontract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Tokencontract.Contract.TokenByIndex(&_Tokencontract.CallOpts, index)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Tokencontract *TokencontractCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Tokencontract *TokencontractSession) TokenFactory() (common.Address, error) {
	return _Tokencontract.Contract.TokenFactory(&_Tokencontract.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Tokencontract *TokencontractCallerSession) TokenFactory() (common.Address, error) {
	return _Tokencontract.Contract.TokenFactory(&_Tokencontract.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Tokencontract *TokencontractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Tokencontract *TokencontractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Tokencontract.Contract.TokenOfOwnerByIndex(&_Tokencontract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Tokencontract.Contract.TokenOfOwnerByIndex(&_Tokencontract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Tokencontract *TokencontractCaller) TokenURI(opts *bind.CallOpts, tokenId_ *big.Int) (string, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "tokenURI", tokenId_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Tokencontract *TokencontractSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Tokencontract.Contract.TokenURI(&_Tokencontract.CallOpts, tokenId_)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId_) view returns(string)
func (_Tokencontract *TokencontractCallerSession) TokenURI(tokenId_ *big.Int) (string, error) {
	return _Tokencontract.Contract.TokenURI(&_Tokencontract.CallOpts, tokenId_)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokencontract *TokencontractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokencontract *TokencontractSession) TotalSupply() (*big.Int, error) {
	return _Tokencontract.Contract.TotalSupply(&_Tokencontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) TotalSupply() (*big.Int, error) {
	return _Tokencontract.Contract.TotalSupply(&_Tokencontract.CallOpts)
}

// VoucherTokenContract is a free data retrieval call binding the contract method 0xf09ad8f2.
//
// Solidity: function voucherTokenContract() view returns(address)
func (_Tokencontract *TokencontractCaller) VoucherTokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "voucherTokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoucherTokenContract is a free data retrieval call binding the contract method 0xf09ad8f2.
//
// Solidity: function voucherTokenContract() view returns(address)
func (_Tokencontract *TokencontractSession) VoucherTokenContract() (common.Address, error) {
	return _Tokencontract.Contract.VoucherTokenContract(&_Tokencontract.CallOpts)
}

// VoucherTokenContract is a free data retrieval call binding the contract method 0xf09ad8f2.
//
// Solidity: function voucherTokenContract() view returns(address)
func (_Tokencontract *TokencontractCallerSession) VoucherTokenContract() (common.Address, error) {
	return _Tokencontract.Contract.VoucherTokenContract(&_Tokencontract.CallOpts)
}

// VoucherTokensAmount is a free data retrieval call binding the contract method 0x59843195.
//
// Solidity: function voucherTokensAmount() view returns(uint256)
func (_Tokencontract *TokencontractCaller) VoucherTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "voucherTokensAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherTokensAmount is a free data retrieval call binding the contract method 0x59843195.
//
// Solidity: function voucherTokensAmount() view returns(uint256)
func (_Tokencontract *TokencontractSession) VoucherTokensAmount() (*big.Int, error) {
	return _Tokencontract.Contract.VoucherTokensAmount(&_Tokencontract.CallOpts)
}

// VoucherTokensAmount is a free data retrieval call binding the contract method 0x59843195.
//
// Solidity: function voucherTokensAmount() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) VoucherTokensAmount() (*big.Int, error) {
	return _Tokencontract.Contract.VoucherTokensAmount(&_Tokencontract.CallOpts)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0x8b1b500b.
//
// Solidity: function __TokenContract_init(string tokenName_, string tokenSymbol_, address tokenFactoryAddr_, uint256 pricePerOneToken_, address voucherTokenContract_, uint256 voucherTokensAmount_) returns()
func (_Tokencontract *TokencontractTransactor) TokenContractInit(opts *bind.TransactOpts, tokenName_ string, tokenSymbol_ string, tokenFactoryAddr_ common.Address, pricePerOneToken_ *big.Int, voucherTokenContract_ common.Address, voucherTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "__TokenContract_init", tokenName_, tokenSymbol_, tokenFactoryAddr_, pricePerOneToken_, voucherTokenContract_, voucherTokensAmount_)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0x8b1b500b.
//
// Solidity: function __TokenContract_init(string tokenName_, string tokenSymbol_, address tokenFactoryAddr_, uint256 pricePerOneToken_, address voucherTokenContract_, uint256 voucherTokensAmount_) returns()
func (_Tokencontract *TokencontractSession) TokenContractInit(tokenName_ string, tokenSymbol_ string, tokenFactoryAddr_ common.Address, pricePerOneToken_ *big.Int, voucherTokenContract_ common.Address, voucherTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.TokenContractInit(&_Tokencontract.TransactOpts, tokenName_, tokenSymbol_, tokenFactoryAddr_, pricePerOneToken_, voucherTokenContract_, voucherTokensAmount_)
}

// TokenContractInit is a paid mutator transaction binding the contract method 0x8b1b500b.
//
// Solidity: function __TokenContract_init(string tokenName_, string tokenSymbol_, address tokenFactoryAddr_, uint256 pricePerOneToken_, address voucherTokenContract_, uint256 voucherTokensAmount_) returns()
func (_Tokencontract *TokencontractTransactorSession) TokenContractInit(tokenName_ string, tokenSymbol_ string, tokenFactoryAddr_ common.Address, pricePerOneToken_ *big.Int, voucherTokenContract_ common.Address, voucherTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.TokenContractInit(&_Tokencontract.TransactOpts, tokenName_, tokenSymbol_, tokenFactoryAddr_, pricePerOneToken_, voucherTokenContract_, voucherTokensAmount_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Approve(&_Tokencontract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Approve(&_Tokencontract.TransactOpts, to, tokenId)
}

// MintToken is a paid mutator transaction binding the contract method 0xc0e0d461.
//
// Solidity: function mintToken(address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 discount_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Tokencontract *TokencontractTransactor) MintToken(opts *bind.TransactOpts, paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, discount_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "mintToken", paymentTokenAddress_, paymentTokenPrice_, discount_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// MintToken is a paid mutator transaction binding the contract method 0xc0e0d461.
//
// Solidity: function mintToken(address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 discount_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Tokencontract *TokencontractSession) MintToken(paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, discount_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Tokencontract.Contract.MintToken(&_Tokencontract.TransactOpts, paymentTokenAddress_, paymentTokenPrice_, discount_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// MintToken is a paid mutator transaction binding the contract method 0xc0e0d461.
//
// Solidity: function mintToken(address paymentTokenAddress_, uint256 paymentTokenPrice_, uint256 discount_, uint256 endTimestamp_, string tokenURI_, bytes32 r_, bytes32 s_, uint8 v_) payable returns()
func (_Tokencontract *TokencontractTransactorSession) MintToken(paymentTokenAddress_ common.Address, paymentTokenPrice_ *big.Int, discount_ *big.Int, endTimestamp_ *big.Int, tokenURI_ string, r_ [32]byte, s_ [32]byte, v_ uint8) (*types.Transaction, error) {
	return _Tokencontract.Contract.MintToken(&_Tokencontract.TransactOpts, paymentTokenAddress_, paymentTokenPrice_, discount_, endTimestamp_, tokenURI_, r_, s_, v_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokencontract *TokencontractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokencontract *TokencontractSession) Pause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Pause(&_Tokencontract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokencontract *TokencontractTransactorSession) Pause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Pause(&_Tokencontract.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.SafeTransferFrom(&_Tokencontract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.SafeTransferFrom(&_Tokencontract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Tokencontract *TokencontractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Tokencontract *TokencontractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Tokencontract.Contract.SafeTransferFrom0(&_Tokencontract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Tokencontract *TokencontractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Tokencontract.Contract.SafeTransferFrom0(&_Tokencontract.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Tokencontract *TokencontractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Tokencontract *TokencontractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Tokencontract.Contract.SetApprovalForAll(&_Tokencontract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Tokencontract *TokencontractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Tokencontract.Contract.SetApprovalForAll(&_Tokencontract.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.TransferFrom(&_Tokencontract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Tokencontract *TokencontractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.TransferFrom(&_Tokencontract.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tokencontract *TokencontractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tokencontract *TokencontractSession) Unpause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Unpause(&_Tokencontract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tokencontract *TokencontractTransactorSession) Unpause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Unpause(&_Tokencontract.TransactOpts)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0xea527ebc.
//
// Solidity: function updateAllParams(uint256 newPrice_, address newVoucherTokenContract_, uint256 newVoucherTokensAmount_, string newTokenName_, string newTokenSymbol_) returns()
func (_Tokencontract *TokencontractTransactor) UpdateAllParams(opts *bind.TransactOpts, newPrice_ *big.Int, newVoucherTokenContract_ common.Address, newVoucherTokensAmount_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "updateAllParams", newPrice_, newVoucherTokenContract_, newVoucherTokensAmount_, newTokenName_, newTokenSymbol_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0xea527ebc.
//
// Solidity: function updateAllParams(uint256 newPrice_, address newVoucherTokenContract_, uint256 newVoucherTokensAmount_, string newTokenName_, string newTokenSymbol_) returns()
func (_Tokencontract *TokencontractSession) UpdateAllParams(newPrice_ *big.Int, newVoucherTokenContract_ common.Address, newVoucherTokensAmount_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Tokencontract.Contract.UpdateAllParams(&_Tokencontract.TransactOpts, newPrice_, newVoucherTokenContract_, newVoucherTokensAmount_, newTokenName_, newTokenSymbol_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0xea527ebc.
//
// Solidity: function updateAllParams(uint256 newPrice_, address newVoucherTokenContract_, uint256 newVoucherTokensAmount_, string newTokenName_, string newTokenSymbol_) returns()
func (_Tokencontract *TokencontractTransactorSession) UpdateAllParams(newPrice_ *big.Int, newVoucherTokenContract_ common.Address, newVoucherTokensAmount_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Tokencontract.Contract.UpdateAllParams(&_Tokencontract.TransactOpts, newPrice_, newVoucherTokenContract_, newVoucherTokensAmount_, newTokenName_, newTokenSymbol_)
}

// UpdateTokenContractParams is a paid mutator transaction binding the contract method 0x541a738c.
//
// Solidity: function updateTokenContractParams(uint256 newPrice_, string newTokenName_, string newTokenSymbol_) returns()
func (_Tokencontract *TokencontractTransactor) UpdateTokenContractParams(opts *bind.TransactOpts, newPrice_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "updateTokenContractParams", newPrice_, newTokenName_, newTokenSymbol_)
}

// UpdateTokenContractParams is a paid mutator transaction binding the contract method 0x541a738c.
//
// Solidity: function updateTokenContractParams(uint256 newPrice_, string newTokenName_, string newTokenSymbol_) returns()
func (_Tokencontract *TokencontractSession) UpdateTokenContractParams(newPrice_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Tokencontract.Contract.UpdateTokenContractParams(&_Tokencontract.TransactOpts, newPrice_, newTokenName_, newTokenSymbol_)
}

// UpdateTokenContractParams is a paid mutator transaction binding the contract method 0x541a738c.
//
// Solidity: function updateTokenContractParams(uint256 newPrice_, string newTokenName_, string newTokenSymbol_) returns()
func (_Tokencontract *TokencontractTransactorSession) UpdateTokenContractParams(newPrice_ *big.Int, newTokenName_ string, newTokenSymbol_ string) (*types.Transaction, error) {
	return _Tokencontract.Contract.UpdateTokenContractParams(&_Tokencontract.TransactOpts, newPrice_, newTokenName_, newTokenSymbol_)
}

// UpdateVoucherParams is a paid mutator transaction binding the contract method 0xfc8c2147.
//
// Solidity: function updateVoucherParams(address newVoucherTokenContract_, uint256 newVoucherTokensAmount_) returns()
func (_Tokencontract *TokencontractTransactor) UpdateVoucherParams(opts *bind.TransactOpts, newVoucherTokenContract_ common.Address, newVoucherTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "updateVoucherParams", newVoucherTokenContract_, newVoucherTokensAmount_)
}

// UpdateVoucherParams is a paid mutator transaction binding the contract method 0xfc8c2147.
//
// Solidity: function updateVoucherParams(address newVoucherTokenContract_, uint256 newVoucherTokensAmount_) returns()
func (_Tokencontract *TokencontractSession) UpdateVoucherParams(newVoucherTokenContract_ common.Address, newVoucherTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.UpdateVoucherParams(&_Tokencontract.TransactOpts, newVoucherTokenContract_, newVoucherTokensAmount_)
}

// UpdateVoucherParams is a paid mutator transaction binding the contract method 0xfc8c2147.
//
// Solidity: function updateVoucherParams(address newVoucherTokenContract_, uint256 newVoucherTokensAmount_) returns()
func (_Tokencontract *TokencontractTransactorSession) UpdateVoucherParams(newVoucherTokenContract_ common.Address, newVoucherTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.UpdateVoucherParams(&_Tokencontract.TransactOpts, newVoucherTokenContract_, newVoucherTokensAmount_)
}

// WithdrawPaidTokens is a paid mutator transaction binding the contract method 0xd3fd39aa.
//
// Solidity: function withdrawPaidTokens(address tokenAddr_, address recipient_) returns()
func (_Tokencontract *TokencontractTransactor) WithdrawPaidTokens(opts *bind.TransactOpts, tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "withdrawPaidTokens", tokenAddr_, recipient_)
}

// WithdrawPaidTokens is a paid mutator transaction binding the contract method 0xd3fd39aa.
//
// Solidity: function withdrawPaidTokens(address tokenAddr_, address recipient_) returns()
func (_Tokencontract *TokencontractSession) WithdrawPaidTokens(tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.WithdrawPaidTokens(&_Tokencontract.TransactOpts, tokenAddr_, recipient_)
}

// WithdrawPaidTokens is a paid mutator transaction binding the contract method 0xd3fd39aa.
//
// Solidity: function withdrawPaidTokens(address tokenAddr_, address recipient_) returns()
func (_Tokencontract *TokencontractTransactorSession) WithdrawPaidTokens(tokenAddr_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.WithdrawPaidTokens(&_Tokencontract.TransactOpts, tokenAddr_, recipient_)
}

// TokencontractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokencontract contract.
type TokencontractApprovalIterator struct {
	Event *TokencontractApproval // Event containing the contract specifics and raw log

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
func (it *TokencontractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractApproval)
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
		it.Event = new(TokencontractApproval)
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
func (it *TokencontractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractApproval represents a Approval event raised by the Tokencontract contract.
type TokencontractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Tokencontract *TokencontractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TokencontractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokencontractApprovalIterator{contract: _Tokencontract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Tokencontract *TokencontractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokencontractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractApproval)
				if err := _Tokencontract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Tokencontract *TokencontractFilterer) ParseApproval(log types.Log) (*TokencontractApproval, error) {
	event := new(TokencontractApproval)
	if err := _Tokencontract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Tokencontract contract.
type TokencontractApprovalForAllIterator struct {
	Event *TokencontractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokencontractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractApprovalForAll)
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
		it.Event = new(TokencontractApprovalForAll)
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
func (it *TokencontractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractApprovalForAll represents a ApprovalForAll event raised by the Tokencontract contract.
type TokencontractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Tokencontract *TokencontractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TokencontractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokencontractApprovalForAllIterator{contract: _Tokencontract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Tokencontract *TokencontractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokencontractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractApprovalForAll)
				if err := _Tokencontract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Tokencontract *TokencontractFilterer) ParseApprovalForAll(log types.Log) (*TokencontractApprovalForAll, error) {
	event := new(TokencontractApprovalForAll)
	if err := _Tokencontract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractPaidTokensWithdrawnIterator is returned from FilterPaidTokensWithdrawn and is used to iterate over the raw logs and unpacked data for PaidTokensWithdrawn events raised by the Tokencontract contract.
type TokencontractPaidTokensWithdrawnIterator struct {
	Event *TokencontractPaidTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *TokencontractPaidTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractPaidTokensWithdrawn)
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
		it.Event = new(TokencontractPaidTokensWithdrawn)
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
func (it *TokencontractPaidTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractPaidTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractPaidTokensWithdrawn represents a PaidTokensWithdrawn event raised by the Tokencontract contract.
type TokencontractPaidTokensWithdrawn struct {
	TokenAddr common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaidTokensWithdrawn is a free log retrieval operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Tokencontract *TokencontractFilterer) FilterPaidTokensWithdrawn(opts *bind.FilterOpts, tokenAddr []common.Address) (*TokencontractPaidTokensWithdrawnIterator, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &TokencontractPaidTokensWithdrawnIterator{contract: _Tokencontract.contract, event: "PaidTokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaidTokensWithdrawn is a free log subscription operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Tokencontract *TokencontractFilterer) WatchPaidTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *TokencontractPaidTokensWithdrawn, tokenAddr []common.Address) (event.Subscription, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractPaidTokensWithdrawn)
				if err := _Tokencontract.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
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

// ParsePaidTokensWithdrawn is a log parse operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Tokencontract *TokencontractFilterer) ParsePaidTokensWithdrawn(log types.Log) (*TokencontractPaidTokensWithdrawn, error) {
	event := new(TokencontractPaidTokensWithdrawn)
	if err := _Tokencontract.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Tokencontract contract.
type TokencontractPausedIterator struct {
	Event *TokencontractPaused // Event containing the contract specifics and raw log

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
func (it *TokencontractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractPaused)
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
		it.Event = new(TokencontractPaused)
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
func (it *TokencontractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractPaused represents a Paused event raised by the Tokencontract contract.
type TokencontractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tokencontract *TokencontractFilterer) FilterPaused(opts *bind.FilterOpts) (*TokencontractPausedIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokencontractPausedIterator{contract: _Tokencontract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tokencontract *TokencontractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokencontractPaused) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractPaused)
				if err := _Tokencontract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tokencontract *TokencontractFilterer) ParsePaused(log types.Log) (*TokencontractPaused, error) {
	event := new(TokencontractPaused)
	if err := _Tokencontract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractSuccessfullyMintedIterator is returned from FilterSuccessfullyMinted and is used to iterate over the raw logs and unpacked data for SuccessfullyMinted events raised by the Tokencontract contract.
type TokencontractSuccessfullyMintedIterator struct {
	Event *TokencontractSuccessfullyMinted // Event containing the contract specifics and raw log

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
func (it *TokencontractSuccessfullyMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractSuccessfullyMinted)
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
		it.Event = new(TokencontractSuccessfullyMinted)
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
func (it *TokencontractSuccessfullyMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractSuccessfullyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractSuccessfullyMinted represents a SuccessfullyMinted event raised by the Tokencontract contract.
type TokencontractSuccessfullyMinted struct {
	Recipient           common.Address
	MintedTokenInfo     ITokenContractMintedTokenInfo
	PaymentTokenAddress common.Address
	PaidTokensAmount    *big.Int
	PaymentTokenPrice   *big.Int
	Discount            *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSuccessfullyMinted is a free log retrieval operation binding the contract event 0xa0e50db80a8b39153d2af9fbbd044e9b21d5aa11c51b217f544f2f308cb1a241.
//
// Solidity: event SuccessfullyMinted(address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed paymentTokenAddress, uint256 paidTokensAmount, uint256 paymentTokenPrice, uint256 discount)
func (_Tokencontract *TokencontractFilterer) FilterSuccessfullyMinted(opts *bind.FilterOpts, recipient []common.Address, paymentTokenAddress []common.Address) (*TokencontractSuccessfullyMintedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var paymentTokenAddressRule []interface{}
	for _, paymentTokenAddressItem := range paymentTokenAddress {
		paymentTokenAddressRule = append(paymentTokenAddressRule, paymentTokenAddressItem)
	}

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "SuccessfullyMinted", recipientRule, paymentTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokencontractSuccessfullyMintedIterator{contract: _Tokencontract.contract, event: "SuccessfullyMinted", logs: logs, sub: sub}, nil
}

// WatchSuccessfullyMinted is a free log subscription operation binding the contract event 0xa0e50db80a8b39153d2af9fbbd044e9b21d5aa11c51b217f544f2f308cb1a241.
//
// Solidity: event SuccessfullyMinted(address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed paymentTokenAddress, uint256 paidTokensAmount, uint256 paymentTokenPrice, uint256 discount)
func (_Tokencontract *TokencontractFilterer) WatchSuccessfullyMinted(opts *bind.WatchOpts, sink chan<- *TokencontractSuccessfullyMinted, recipient []common.Address, paymentTokenAddress []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var paymentTokenAddressRule []interface{}
	for _, paymentTokenAddressItem := range paymentTokenAddress {
		paymentTokenAddressRule = append(paymentTokenAddressRule, paymentTokenAddressItem)
	}

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "SuccessfullyMinted", recipientRule, paymentTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractSuccessfullyMinted)
				if err := _Tokencontract.contract.UnpackLog(event, "SuccessfullyMinted", log); err != nil {
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

// ParseSuccessfullyMinted is a log parse operation binding the contract event 0xa0e50db80a8b39153d2af9fbbd044e9b21d5aa11c51b217f544f2f308cb1a241.
//
// Solidity: event SuccessfullyMinted(address indexed recipient, (uint256,uint256,string) mintedTokenInfo, address indexed paymentTokenAddress, uint256 paidTokensAmount, uint256 paymentTokenPrice, uint256 discount)
func (_Tokencontract *TokencontractFilterer) ParseSuccessfullyMinted(log types.Log) (*TokencontractSuccessfullyMinted, error) {
	event := new(TokencontractSuccessfullyMinted)
	if err := _Tokencontract.contract.UnpackLog(event, "SuccessfullyMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractTokenContractParamsUpdatedIterator is returned from FilterTokenContractParamsUpdated and is used to iterate over the raw logs and unpacked data for TokenContractParamsUpdated events raised by the Tokencontract contract.
type TokencontractTokenContractParamsUpdatedIterator struct {
	Event *TokencontractTokenContractParamsUpdated // Event containing the contract specifics and raw log

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
func (it *TokencontractTokenContractParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractTokenContractParamsUpdated)
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
		it.Event = new(TokencontractTokenContractParamsUpdated)
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
func (it *TokencontractTokenContractParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractTokenContractParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractTokenContractParamsUpdated represents a TokenContractParamsUpdated event raised by the Tokencontract contract.
type TokencontractTokenContractParamsUpdated struct {
	NewPrice    *big.Int
	TokenName   string
	TokenSymbol string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenContractParamsUpdated is a free log retrieval operation binding the contract event 0x3419ca7e094a7b9d428666d41c7b45306d09be116d969b39e10447b02c2ee50f.
//
// Solidity: event TokenContractParamsUpdated(uint256 newPrice, string tokenName, string tokenSymbol)
func (_Tokencontract *TokencontractFilterer) FilterTokenContractParamsUpdated(opts *bind.FilterOpts) (*TokencontractTokenContractParamsUpdatedIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "TokenContractParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &TokencontractTokenContractParamsUpdatedIterator{contract: _Tokencontract.contract, event: "TokenContractParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenContractParamsUpdated is a free log subscription operation binding the contract event 0x3419ca7e094a7b9d428666d41c7b45306d09be116d969b39e10447b02c2ee50f.
//
// Solidity: event TokenContractParamsUpdated(uint256 newPrice, string tokenName, string tokenSymbol)
func (_Tokencontract *TokencontractFilterer) WatchTokenContractParamsUpdated(opts *bind.WatchOpts, sink chan<- *TokencontractTokenContractParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "TokenContractParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractTokenContractParamsUpdated)
				if err := _Tokencontract.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
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

// ParseTokenContractParamsUpdated is a log parse operation binding the contract event 0x3419ca7e094a7b9d428666d41c7b45306d09be116d969b39e10447b02c2ee50f.
//
// Solidity: event TokenContractParamsUpdated(uint256 newPrice, string tokenName, string tokenSymbol)
func (_Tokencontract *TokencontractFilterer) ParseTokenContractParamsUpdated(log types.Log) (*TokencontractTokenContractParamsUpdated, error) {
	event := new(TokencontractTokenContractParamsUpdated)
	if err := _Tokencontract.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokencontract contract.
type TokencontractTransferIterator struct {
	Event *TokencontractTransfer // Event containing the contract specifics and raw log

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
func (it *TokencontractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractTransfer)
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
		it.Event = new(TokencontractTransfer)
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
func (it *TokencontractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractTransfer represents a Transfer event raised by the Tokencontract contract.
type TokencontractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Tokencontract *TokencontractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TokencontractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokencontractTransferIterator{contract: _Tokencontract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Tokencontract *TokencontractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokencontractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractTransfer)
				if err := _Tokencontract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Tokencontract *TokencontractFilterer) ParseTransfer(log types.Log) (*TokencontractTransfer, error) {
	event := new(TokencontractTransfer)
	if err := _Tokencontract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Tokencontract contract.
type TokencontractUnpausedIterator struct {
	Event *TokencontractUnpaused // Event containing the contract specifics and raw log

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
func (it *TokencontractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractUnpaused)
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
		it.Event = new(TokencontractUnpaused)
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
func (it *TokencontractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractUnpaused represents a Unpaused event raised by the Tokencontract contract.
type TokencontractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tokencontract *TokencontractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TokencontractUnpausedIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TokencontractUnpausedIterator{contract: _Tokencontract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tokencontract *TokencontractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TokencontractUnpaused) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractUnpaused)
				if err := _Tokencontract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tokencontract *TokencontractFilterer) ParseUnpaused(log types.Log) (*TokencontractUnpaused, error) {
	event := new(TokencontractUnpaused)
	if err := _Tokencontract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractVoucherParamsUpdatedIterator is returned from FilterVoucherParamsUpdated and is used to iterate over the raw logs and unpacked data for VoucherParamsUpdated events raised by the Tokencontract contract.
type TokencontractVoucherParamsUpdatedIterator struct {
	Event *TokencontractVoucherParamsUpdated // Event containing the contract specifics and raw log

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
func (it *TokencontractVoucherParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractVoucherParamsUpdated)
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
		it.Event = new(TokencontractVoucherParamsUpdated)
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
func (it *TokencontractVoucherParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractVoucherParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractVoucherParamsUpdated represents a VoucherParamsUpdated event raised by the Tokencontract contract.
type TokencontractVoucherParamsUpdated struct {
	NewVoucherTokenContract common.Address
	NewVoucherTokensAmount  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterVoucherParamsUpdated is a free log retrieval operation binding the contract event 0x526fc674430e078ab81756274a8fedd526b74df2995fe145d18f8bf48ea0a1ef.
//
// Solidity: event VoucherParamsUpdated(address newVoucherTokenContract, uint256 newVoucherTokensAmount)
func (_Tokencontract *TokencontractFilterer) FilterVoucherParamsUpdated(opts *bind.FilterOpts) (*TokencontractVoucherParamsUpdatedIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "VoucherParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &TokencontractVoucherParamsUpdatedIterator{contract: _Tokencontract.contract, event: "VoucherParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchVoucherParamsUpdated is a free log subscription operation binding the contract event 0x526fc674430e078ab81756274a8fedd526b74df2995fe145d18f8bf48ea0a1ef.
//
// Solidity: event VoucherParamsUpdated(address newVoucherTokenContract, uint256 newVoucherTokensAmount)
func (_Tokencontract *TokencontractFilterer) WatchVoucherParamsUpdated(opts *bind.WatchOpts, sink chan<- *TokencontractVoucherParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "VoucherParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractVoucherParamsUpdated)
				if err := _Tokencontract.contract.UnpackLog(event, "VoucherParamsUpdated", log); err != nil {
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

// ParseVoucherParamsUpdated is a log parse operation binding the contract event 0x526fc674430e078ab81756274a8fedd526b74df2995fe145d18f8bf48ea0a1ef.
//
// Solidity: event VoucherParamsUpdated(address newVoucherTokenContract, uint256 newVoucherTokensAmount)
func (_Tokencontract *TokencontractFilterer) ParseVoucherParamsUpdated(log types.Log) (*TokencontractVoucherParamsUpdated, error) {
	event := new(TokencontractVoucherParamsUpdated)
	if err := _Tokencontract.contract.UnpackLog(event, "VoucherParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
