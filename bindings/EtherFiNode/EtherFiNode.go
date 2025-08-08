// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EtherFiNode

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

// BeaconChainProofsBalanceContainerProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceContainerProof struct {
	BalanceContainerRoot [32]byte
	Proof                []byte
}

// BeaconChainProofsBalanceProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceProof struct {
	PubkeyHash  [32]byte
	BalanceRoot [32]byte
	Proof       []byte
}

// IDelegationManagerTypesQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesQueuedWithdrawalParams struct {
	Strategies           []common.Address
	DepositShares        []*big.Int
	DeprecatedWithdrawer common.Address
}

// IDelegationManagerTypesWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesWithdrawal struct {
	Staker       common.Address
	DelegatedTo  common.Address
	Withdrawer   common.Address
	Nonce        *big.Int
	StartBlock   uint32
	Strategies   []common.Address
	ScaledShares []*big.Int
}

// EtherFiNodeMetaData contains all meta data concerning the EtherFiNode contract.
var EtherFiNodeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_liquidityPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_etherFiNodesManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_roleRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"EIGENLAYER_WITHDRAWAL_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETHERFI_NODE_CALL_FORWARDER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETHERFI_NODE_EIGENLAYER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeQueuedETHWithdrawals\",\"inputs\":[{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createEigenPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"etherFiNodesManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEtherFiNodesManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwardEigenPodCall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardExternalCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEigenPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidityPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILiquidityPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueETHWithdrawal\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"__deprecated_withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roleRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRoleRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setProofSubmitter\",\"inputs\":[{\"name\":\"_newProofSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweepFunds\",\"inputs\":[],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCheckpointProofs\",\"inputs\":[{\"name\":\"balanceContainerProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.BalanceContainerProof\",\"components\":[{\"name\":\"balanceContainerRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.BalanceProof[]\",\"components\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"balanceRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FullWithdrawal\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toOperator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toBnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTreasury\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsTransferred\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialWithdrawal\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toOperator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toBnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTreasury\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QueuedRestakingWithdrawal\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ForwardedCallNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectRole\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidForwardedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]}]",
}

// EtherFiNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherFiNodeMetaData.ABI instead.
var EtherFiNodeABI = EtherFiNodeMetaData.ABI

// EtherFiNode is an auto generated Go binding around an Ethereum contract.
type EtherFiNode struct {
	EtherFiNodeCaller     // Read-only binding to the contract
	EtherFiNodeTransactor // Write-only binding to the contract
	EtherFiNodeFilterer   // Log filterer for contract events
}

// EtherFiNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherFiNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherFiNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherFiNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherFiNodeSession struct {
	Contract     *EtherFiNode      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherFiNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherFiNodeCallerSession struct {
	Contract *EtherFiNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EtherFiNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherFiNodeTransactorSession struct {
	Contract     *EtherFiNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EtherFiNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherFiNodeRaw struct {
	Contract *EtherFiNode // Generic contract binding to access the raw methods on
}

// EtherFiNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherFiNodeCallerRaw struct {
	Contract *EtherFiNodeCaller // Generic read-only contract binding to access the raw methods on
}

// EtherFiNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherFiNodeTransactorRaw struct {
	Contract *EtherFiNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherFiNode creates a new instance of EtherFiNode, bound to a specific deployed contract.
func NewEtherFiNode(address common.Address, backend bind.ContractBackend) (*EtherFiNode, error) {
	contract, err := bindEtherFiNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherFiNode{EtherFiNodeCaller: EtherFiNodeCaller{contract: contract}, EtherFiNodeTransactor: EtherFiNodeTransactor{contract: contract}, EtherFiNodeFilterer: EtherFiNodeFilterer{contract: contract}}, nil
}

// NewEtherFiNodeCaller creates a new read-only instance of EtherFiNode, bound to a specific deployed contract.
func NewEtherFiNodeCaller(address common.Address, caller bind.ContractCaller) (*EtherFiNodeCaller, error) {
	contract, err := bindEtherFiNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodeCaller{contract: contract}, nil
}

// NewEtherFiNodeTransactor creates a new write-only instance of EtherFiNode, bound to a specific deployed contract.
func NewEtherFiNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherFiNodeTransactor, error) {
	contract, err := bindEtherFiNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodeTransactor{contract: contract}, nil
}

// NewEtherFiNodeFilterer creates a new log filterer instance of EtherFiNode, bound to a specific deployed contract.
func NewEtherFiNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherFiNodeFilterer, error) {
	contract, err := bindEtherFiNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodeFilterer{contract: contract}, nil
}

// bindEtherFiNode binds a generic wrapper to an already deployed contract.
func bindEtherFiNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherFiNodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherFiNode *EtherFiNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherFiNode.Contract.EtherFiNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherFiNode *EtherFiNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNode.Contract.EtherFiNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherFiNode *EtherFiNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherFiNode.Contract.EtherFiNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherFiNode *EtherFiNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherFiNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherFiNode *EtherFiNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherFiNode *EtherFiNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherFiNode.Contract.contract.Transact(opts, method, params...)
}

// EIGENLAYERWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xb9da1262.
//
// Solidity: function EIGENLAYER_WITHDRAWAL_DELAY_BLOCKS() view returns(uint32)
func (_EtherFiNode *EtherFiNodeCaller) EIGENLAYERWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "EIGENLAYER_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// EIGENLAYERWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xb9da1262.
//
// Solidity: function EIGENLAYER_WITHDRAWAL_DELAY_BLOCKS() view returns(uint32)
func (_EtherFiNode *EtherFiNodeSession) EIGENLAYERWITHDRAWALDELAYBLOCKS() (uint32, error) {
	return _EtherFiNode.Contract.EIGENLAYERWITHDRAWALDELAYBLOCKS(&_EtherFiNode.CallOpts)
}

// EIGENLAYERWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xb9da1262.
//
// Solidity: function EIGENLAYER_WITHDRAWAL_DELAY_BLOCKS() view returns(uint32)
func (_EtherFiNode *EtherFiNodeCallerSession) EIGENLAYERWITHDRAWALDELAYBLOCKS() (uint32, error) {
	return _EtherFiNode.Contract.EIGENLAYERWITHDRAWALDELAYBLOCKS(&_EtherFiNode.CallOpts)
}

// ETHERFINODECALLFORWARDERROLE is a free data retrieval call binding the contract method 0xb62ab426.
//
// Solidity: function ETHERFI_NODE_CALL_FORWARDER_ROLE() view returns(bytes32)
func (_EtherFiNode *EtherFiNodeCaller) ETHERFINODECALLFORWARDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "ETHERFI_NODE_CALL_FORWARDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHERFINODECALLFORWARDERROLE is a free data retrieval call binding the contract method 0xb62ab426.
//
// Solidity: function ETHERFI_NODE_CALL_FORWARDER_ROLE() view returns(bytes32)
func (_EtherFiNode *EtherFiNodeSession) ETHERFINODECALLFORWARDERROLE() ([32]byte, error) {
	return _EtherFiNode.Contract.ETHERFINODECALLFORWARDERROLE(&_EtherFiNode.CallOpts)
}

// ETHERFINODECALLFORWARDERROLE is a free data retrieval call binding the contract method 0xb62ab426.
//
// Solidity: function ETHERFI_NODE_CALL_FORWARDER_ROLE() view returns(bytes32)
func (_EtherFiNode *EtherFiNodeCallerSession) ETHERFINODECALLFORWARDERROLE() ([32]byte, error) {
	return _EtherFiNode.Contract.ETHERFINODECALLFORWARDERROLE(&_EtherFiNode.CallOpts)
}

// ETHERFINODEEIGENLAYERADMINROLE is a free data retrieval call binding the contract method 0xc587d7ab.
//
// Solidity: function ETHERFI_NODE_EIGENLAYER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNode *EtherFiNodeCaller) ETHERFINODEEIGENLAYERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "ETHERFI_NODE_EIGENLAYER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHERFINODEEIGENLAYERADMINROLE is a free data retrieval call binding the contract method 0xc587d7ab.
//
// Solidity: function ETHERFI_NODE_EIGENLAYER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNode *EtherFiNodeSession) ETHERFINODEEIGENLAYERADMINROLE() ([32]byte, error) {
	return _EtherFiNode.Contract.ETHERFINODEEIGENLAYERADMINROLE(&_EtherFiNode.CallOpts)
}

// ETHERFINODEEIGENLAYERADMINROLE is a free data retrieval call binding the contract method 0xc587d7ab.
//
// Solidity: function ETHERFI_NODE_EIGENLAYER_ADMIN_ROLE() view returns(bytes32)
func (_EtherFiNode *EtherFiNodeCallerSession) ETHERFINODEEIGENLAYERADMINROLE() ([32]byte, error) {
	return _EtherFiNode.Contract.ETHERFINODEEIGENLAYERADMINROLE(&_EtherFiNode.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherFiNode *EtherFiNodeCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherFiNode *EtherFiNodeSession) DelegationManager() (common.Address, error) {
	return _EtherFiNode.Contract.DelegationManager(&_EtherFiNode.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherFiNode *EtherFiNodeCallerSession) DelegationManager() (common.Address, error) {
	return _EtherFiNode.Contract.DelegationManager(&_EtherFiNode.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherFiNode *EtherFiNodeCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherFiNode *EtherFiNodeSession) EigenPodManager() (common.Address, error) {
	return _EtherFiNode.Contract.EigenPodManager(&_EtherFiNode.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherFiNode *EtherFiNodeCallerSession) EigenPodManager() (common.Address, error) {
	return _EtherFiNode.Contract.EigenPodManager(&_EtherFiNode.CallOpts)
}

// EtherFiNodesManager is a free data retrieval call binding the contract method 0x089acd98.
//
// Solidity: function etherFiNodesManager() view returns(address)
func (_EtherFiNode *EtherFiNodeCaller) EtherFiNodesManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "etherFiNodesManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EtherFiNodesManager is a free data retrieval call binding the contract method 0x089acd98.
//
// Solidity: function etherFiNodesManager() view returns(address)
func (_EtherFiNode *EtherFiNodeSession) EtherFiNodesManager() (common.Address, error) {
	return _EtherFiNode.Contract.EtherFiNodesManager(&_EtherFiNode.CallOpts)
}

// EtherFiNodesManager is a free data retrieval call binding the contract method 0x089acd98.
//
// Solidity: function etherFiNodesManager() view returns(address)
func (_EtherFiNode *EtherFiNodeCallerSession) EtherFiNodesManager() (common.Address, error) {
	return _EtherFiNode.Contract.EtherFiNodesManager(&_EtherFiNode.CallOpts)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xbcbb073a.
//
// Solidity: function getEigenPod() view returns(address)
func (_EtherFiNode *EtherFiNodeCaller) GetEigenPod(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "getEigenPod")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEigenPod is a free data retrieval call binding the contract method 0xbcbb073a.
//
// Solidity: function getEigenPod() view returns(address)
func (_EtherFiNode *EtherFiNodeSession) GetEigenPod() (common.Address, error) {
	return _EtherFiNode.Contract.GetEigenPod(&_EtherFiNode.CallOpts)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xbcbb073a.
//
// Solidity: function getEigenPod() view returns(address)
func (_EtherFiNode *EtherFiNodeCallerSession) GetEigenPod() (common.Address, error) {
	return _EtherFiNode.Contract.GetEigenPod(&_EtherFiNode.CallOpts)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_EtherFiNode *EtherFiNodeCaller) LiquidityPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "liquidityPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_EtherFiNode *EtherFiNodeSession) LiquidityPool() (common.Address, error) {
	return _EtherFiNode.Contract.LiquidityPool(&_EtherFiNode.CallOpts)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_EtherFiNode *EtherFiNodeCallerSession) LiquidityPool() (common.Address, error) {
	return _EtherFiNode.Contract.LiquidityPool(&_EtherFiNode.CallOpts)
}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_EtherFiNode *EtherFiNodeCaller) RoleRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNode.contract.Call(opts, &out, "roleRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_EtherFiNode *EtherFiNodeSession) RoleRegistry() (common.Address, error) {
	return _EtherFiNode.Contract.RoleRegistry(&_EtherFiNode.CallOpts)
}

// RoleRegistry is a free data retrieval call binding the contract method 0x08c73259.
//
// Solidity: function roleRegistry() view returns(address)
func (_EtherFiNode *EtherFiNodeCallerSession) RoleRegistry() (common.Address, error) {
	return _EtherFiNode.Contract.RoleRegistry(&_EtherFiNode.CallOpts)
}

// CompleteQueuedETHWithdrawals is a paid mutator transaction binding the contract method 0x5d089f6f.
//
// Solidity: function completeQueuedETHWithdrawals(bool receiveAsTokens) returns(uint256 balance)
func (_EtherFiNode *EtherFiNodeTransactor) CompleteQueuedETHWithdrawals(opts *bind.TransactOpts, receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "completeQueuedETHWithdrawals", receiveAsTokens)
}

// CompleteQueuedETHWithdrawals is a paid mutator transaction binding the contract method 0x5d089f6f.
//
// Solidity: function completeQueuedETHWithdrawals(bool receiveAsTokens) returns(uint256 balance)
func (_EtherFiNode *EtherFiNodeSession) CompleteQueuedETHWithdrawals(receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNode.Contract.CompleteQueuedETHWithdrawals(&_EtherFiNode.TransactOpts, receiveAsTokens)
}

// CompleteQueuedETHWithdrawals is a paid mutator transaction binding the contract method 0x5d089f6f.
//
// Solidity: function completeQueuedETHWithdrawals(bool receiveAsTokens) returns(uint256 balance)
func (_EtherFiNode *EtherFiNodeTransactorSession) CompleteQueuedETHWithdrawals(receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNode.Contract.CompleteQueuedETHWithdrawals(&_EtherFiNode.TransactOpts, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_EtherFiNode *EtherFiNodeTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "completeQueuedWithdrawals", withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_EtherFiNode *EtherFiNodeSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _EtherFiNode.Contract.CompleteQueuedWithdrawals(&_EtherFiNode.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_EtherFiNode *EtherFiNodeTransactorSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _EtherFiNode.Contract.CompleteQueuedWithdrawals(&_EtherFiNode.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// CreateEigenPod is a paid mutator transaction binding the contract method 0x0b10b201.
//
// Solidity: function createEigenPod() returns(address)
func (_EtherFiNode *EtherFiNodeTransactor) CreateEigenPod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "createEigenPod")
}

// CreateEigenPod is a paid mutator transaction binding the contract method 0x0b10b201.
//
// Solidity: function createEigenPod() returns(address)
func (_EtherFiNode *EtherFiNodeSession) CreateEigenPod() (*types.Transaction, error) {
	return _EtherFiNode.Contract.CreateEigenPod(&_EtherFiNode.TransactOpts)
}

// CreateEigenPod is a paid mutator transaction binding the contract method 0x0b10b201.
//
// Solidity: function createEigenPod() returns(address)
func (_EtherFiNode *EtherFiNodeTransactorSession) CreateEigenPod() (*types.Transaction, error) {
	return _EtherFiNode.Contract.CreateEigenPod(&_EtherFiNode.TransactOpts)
}

// ForwardEigenPodCall is a paid mutator transaction binding the contract method 0xc1bceb8c.
//
// Solidity: function forwardEigenPodCall(bytes data) returns(bytes)
func (_EtherFiNode *EtherFiNodeTransactor) ForwardEigenPodCall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "forwardEigenPodCall", data)
}

// ForwardEigenPodCall is a paid mutator transaction binding the contract method 0xc1bceb8c.
//
// Solidity: function forwardEigenPodCall(bytes data) returns(bytes)
func (_EtherFiNode *EtherFiNodeSession) ForwardEigenPodCall(data []byte) (*types.Transaction, error) {
	return _EtherFiNode.Contract.ForwardEigenPodCall(&_EtherFiNode.TransactOpts, data)
}

// ForwardEigenPodCall is a paid mutator transaction binding the contract method 0xc1bceb8c.
//
// Solidity: function forwardEigenPodCall(bytes data) returns(bytes)
func (_EtherFiNode *EtherFiNodeTransactorSession) ForwardEigenPodCall(data []byte) (*types.Transaction, error) {
	return _EtherFiNode.Contract.ForwardEigenPodCall(&_EtherFiNode.TransactOpts, data)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x5726bd51.
//
// Solidity: function forwardExternalCall(address to, bytes data) returns(bytes)
func (_EtherFiNode *EtherFiNodeTransactor) ForwardExternalCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "forwardExternalCall", to, data)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x5726bd51.
//
// Solidity: function forwardExternalCall(address to, bytes data) returns(bytes)
func (_EtherFiNode *EtherFiNodeSession) ForwardExternalCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNode.Contract.ForwardExternalCall(&_EtherFiNode.TransactOpts, to, data)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x5726bd51.
//
// Solidity: function forwardExternalCall(address to, bytes data) returns(bytes)
func (_EtherFiNode *EtherFiNodeTransactorSession) ForwardExternalCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNode.Contract.ForwardExternalCall(&_EtherFiNode.TransactOpts, to, data)
}

// QueueETHWithdrawal is a paid mutator transaction binding the contract method 0xa39945d0.
//
// Solidity: function queueETHWithdrawal(uint256 amount) returns(bytes32 withdrawalRoot)
func (_EtherFiNode *EtherFiNodeTransactor) QueueETHWithdrawal(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "queueETHWithdrawal", amount)
}

// QueueETHWithdrawal is a paid mutator transaction binding the contract method 0xa39945d0.
//
// Solidity: function queueETHWithdrawal(uint256 amount) returns(bytes32 withdrawalRoot)
func (_EtherFiNode *EtherFiNodeSession) QueueETHWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _EtherFiNode.Contract.QueueETHWithdrawal(&_EtherFiNode.TransactOpts, amount)
}

// QueueETHWithdrawal is a paid mutator transaction binding the contract method 0xa39945d0.
//
// Solidity: function queueETHWithdrawal(uint256 amount) returns(bytes32 withdrawalRoot)
func (_EtherFiNode *EtherFiNodeTransactorSession) QueueETHWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _EtherFiNode.Contract.QueueETHWithdrawal(&_EtherFiNode.TransactOpts, amount)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[] withdrawalRoots)
func (_EtherFiNode *EtherFiNodeTransactor) QueueWithdrawals(opts *bind.TransactOpts, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "queueWithdrawals", params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[] withdrawalRoots)
func (_EtherFiNode *EtherFiNodeSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _EtherFiNode.Contract.QueueWithdrawals(&_EtherFiNode.TransactOpts, params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[] withdrawalRoots)
func (_EtherFiNode *EtherFiNodeTransactorSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _EtherFiNode.Contract.QueueWithdrawals(&_EtherFiNode.TransactOpts, params)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address _newProofSubmitter) returns()
func (_EtherFiNode *EtherFiNodeTransactor) SetProofSubmitter(opts *bind.TransactOpts, _newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "setProofSubmitter", _newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address _newProofSubmitter) returns()
func (_EtherFiNode *EtherFiNodeSession) SetProofSubmitter(_newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNode.Contract.SetProofSubmitter(&_EtherFiNode.TransactOpts, _newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address _newProofSubmitter) returns()
func (_EtherFiNode *EtherFiNodeTransactorSession) SetProofSubmitter(_newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNode.Contract.SetProofSubmitter(&_EtherFiNode.TransactOpts, _newProofSubmitter)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x90b51625.
//
// Solidity: function startCheckpoint() returns()
func (_EtherFiNode *EtherFiNodeTransactor) StartCheckpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "startCheckpoint")
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x90b51625.
//
// Solidity: function startCheckpoint() returns()
func (_EtherFiNode *EtherFiNodeSession) StartCheckpoint() (*types.Transaction, error) {
	return _EtherFiNode.Contract.StartCheckpoint(&_EtherFiNode.TransactOpts)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x90b51625.
//
// Solidity: function startCheckpoint() returns()
func (_EtherFiNode *EtherFiNodeTransactorSession) StartCheckpoint() (*types.Transaction, error) {
	return _EtherFiNode.Contract.StartCheckpoint(&_EtherFiNode.TransactOpts)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns(uint256 balance)
func (_EtherFiNode *EtherFiNodeTransactor) SweepFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "sweepFunds")
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns(uint256 balance)
func (_EtherFiNode *EtherFiNodeSession) SweepFunds() (*types.Transaction, error) {
	return _EtherFiNode.Contract.SweepFunds(&_EtherFiNode.TransactOpts)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns(uint256 balance)
func (_EtherFiNode *EtherFiNodeTransactorSession) SweepFunds() (*types.Transaction, error) {
	return _EtherFiNode.Contract.SweepFunds(&_EtherFiNode.TransactOpts)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EtherFiNode *EtherFiNodeTransactor) VerifyCheckpointProofs(opts *bind.TransactOpts, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EtherFiNode.contract.Transact(opts, "verifyCheckpointProofs", balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EtherFiNode *EtherFiNodeSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EtherFiNode.Contract.VerifyCheckpointProofs(&_EtherFiNode.TransactOpts, balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EtherFiNode *EtherFiNodeTransactorSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EtherFiNode.Contract.VerifyCheckpointProofs(&_EtherFiNode.TransactOpts, balanceContainerProof, proofs)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EtherFiNode *EtherFiNodeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _EtherFiNode.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EtherFiNode *EtherFiNodeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _EtherFiNode.Contract.Fallback(&_EtherFiNode.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_EtherFiNode *EtherFiNodeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _EtherFiNode.Contract.Fallback(&_EtherFiNode.TransactOpts, calldata)
}

// EtherFiNodeFullWithdrawalIterator is returned from FilterFullWithdrawal and is used to iterate over the raw logs and unpacked data for FullWithdrawal events raised by the EtherFiNode contract.
type EtherFiNodeFullWithdrawalIterator struct {
	Event *EtherFiNodeFullWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherFiNodeFullWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodeFullWithdrawal)
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
		it.Event = new(EtherFiNodeFullWithdrawal)
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
func (it *EtherFiNodeFullWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodeFullWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodeFullWithdrawal represents a FullWithdrawal event raised by the EtherFiNode contract.
type EtherFiNodeFullWithdrawal struct {
	ValidatorId *big.Int
	EtherFiNode common.Address
	ToOperator  *big.Int
	ToTnft      *big.Int
	ToBnft      *big.Int
	ToTreasury  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFullWithdrawal is a free log retrieval operation binding the contract event 0x23fd4a72178e02ea64b0e1b08ed6de9c7a7fb4bbb565b0917b52e0650a2d3a09.
//
// Solidity: event FullWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNode *EtherFiNodeFilterer) FilterFullWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherFiNodeFullWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNode.contract.FilterLogs(opts, "FullWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodeFullWithdrawalIterator{contract: _EtherFiNode.contract, event: "FullWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFullWithdrawal is a free log subscription operation binding the contract event 0x23fd4a72178e02ea64b0e1b08ed6de9c7a7fb4bbb565b0917b52e0650a2d3a09.
//
// Solidity: event FullWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNode *EtherFiNodeFilterer) WatchFullWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherFiNodeFullWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNode.contract.WatchLogs(opts, "FullWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodeFullWithdrawal)
				if err := _EtherFiNode.contract.UnpackLog(event, "FullWithdrawal", log); err != nil {
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

// ParseFullWithdrawal is a log parse operation binding the contract event 0x23fd4a72178e02ea64b0e1b08ed6de9c7a7fb4bbb565b0917b52e0650a2d3a09.
//
// Solidity: event FullWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNode *EtherFiNodeFilterer) ParseFullWithdrawal(log types.Log) (*EtherFiNodeFullWithdrawal, error) {
	event := new(EtherFiNodeFullWithdrawal)
	if err := _EtherFiNode.contract.UnpackLog(event, "FullWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodeFundsTransferredIterator is returned from FilterFundsTransferred and is used to iterate over the raw logs and unpacked data for FundsTransferred events raised by the EtherFiNode contract.
type EtherFiNodeFundsTransferredIterator struct {
	Event *EtherFiNodeFundsTransferred // Event containing the contract specifics and raw log

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
func (it *EtherFiNodeFundsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodeFundsTransferred)
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
		it.Event = new(EtherFiNodeFundsTransferred)
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
func (it *EtherFiNodeFundsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodeFundsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodeFundsTransferred represents a FundsTransferred event raised by the EtherFiNode contract.
type EtherFiNodeFundsTransferred struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferred is a free log retrieval operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed recipient, uint256 amount)
func (_EtherFiNode *EtherFiNodeFilterer) FilterFundsTransferred(opts *bind.FilterOpts, recipient []common.Address) (*EtherFiNodeFundsTransferredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EtherFiNode.contract.FilterLogs(opts, "FundsTransferred", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodeFundsTransferredIterator{contract: _EtherFiNode.contract, event: "FundsTransferred", logs: logs, sub: sub}, nil
}

// WatchFundsTransferred is a free log subscription operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed recipient, uint256 amount)
func (_EtherFiNode *EtherFiNodeFilterer) WatchFundsTransferred(opts *bind.WatchOpts, sink chan<- *EtherFiNodeFundsTransferred, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EtherFiNode.contract.WatchLogs(opts, "FundsTransferred", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodeFundsTransferred)
				if err := _EtherFiNode.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
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

// ParseFundsTransferred is a log parse operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed recipient, uint256 amount)
func (_EtherFiNode *EtherFiNodeFilterer) ParseFundsTransferred(log types.Log) (*EtherFiNodeFundsTransferred, error) {
	event := new(EtherFiNodeFundsTransferred)
	if err := _EtherFiNode.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodePartialWithdrawalIterator is returned from FilterPartialWithdrawal and is used to iterate over the raw logs and unpacked data for PartialWithdrawal events raised by the EtherFiNode contract.
type EtherFiNodePartialWithdrawalIterator struct {
	Event *EtherFiNodePartialWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherFiNodePartialWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodePartialWithdrawal)
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
		it.Event = new(EtherFiNodePartialWithdrawal)
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
func (it *EtherFiNodePartialWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodePartialWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodePartialWithdrawal represents a PartialWithdrawal event raised by the EtherFiNode contract.
type EtherFiNodePartialWithdrawal struct {
	ValidatorId *big.Int
	EtherFiNode common.Address
	ToOperator  *big.Int
	ToTnft      *big.Int
	ToBnft      *big.Int
	ToTreasury  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPartialWithdrawal is a free log retrieval operation binding the contract event 0x0c9b1112957fe7d0e2f96690e65a9122e07ca9cd19a2f99966b29b5991c3be84.
//
// Solidity: event PartialWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNode *EtherFiNodeFilterer) FilterPartialWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherFiNodePartialWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNode.contract.FilterLogs(opts, "PartialWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodePartialWithdrawalIterator{contract: _EtherFiNode.contract, event: "PartialWithdrawal", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawal is a free log subscription operation binding the contract event 0x0c9b1112957fe7d0e2f96690e65a9122e07ca9cd19a2f99966b29b5991c3be84.
//
// Solidity: event PartialWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNode *EtherFiNodeFilterer) WatchPartialWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherFiNodePartialWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNode.contract.WatchLogs(opts, "PartialWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodePartialWithdrawal)
				if err := _EtherFiNode.contract.UnpackLog(event, "PartialWithdrawal", log); err != nil {
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

// ParsePartialWithdrawal is a log parse operation binding the contract event 0x0c9b1112957fe7d0e2f96690e65a9122e07ca9cd19a2f99966b29b5991c3be84.
//
// Solidity: event PartialWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNode *EtherFiNodeFilterer) ParsePartialWithdrawal(log types.Log) (*EtherFiNodePartialWithdrawal, error) {
	event := new(EtherFiNodePartialWithdrawal)
	if err := _EtherFiNode.contract.UnpackLog(event, "PartialWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodeQueuedRestakingWithdrawalIterator is returned from FilterQueuedRestakingWithdrawal and is used to iterate over the raw logs and unpacked data for QueuedRestakingWithdrawal events raised by the EtherFiNode contract.
type EtherFiNodeQueuedRestakingWithdrawalIterator struct {
	Event *EtherFiNodeQueuedRestakingWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherFiNodeQueuedRestakingWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodeQueuedRestakingWithdrawal)
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
		it.Event = new(EtherFiNodeQueuedRestakingWithdrawal)
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
func (it *EtherFiNodeQueuedRestakingWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodeQueuedRestakingWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodeQueuedRestakingWithdrawal represents a QueuedRestakingWithdrawal event raised by the EtherFiNode contract.
type EtherFiNodeQueuedRestakingWithdrawal struct {
	ValidatorId     *big.Int
	EtherFiNode     common.Address
	WithdrawalRoots [][32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterQueuedRestakingWithdrawal is a free log retrieval operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherFiNode *EtherFiNodeFilterer) FilterQueuedRestakingWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherFiNodeQueuedRestakingWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNode.contract.FilterLogs(opts, "QueuedRestakingWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodeQueuedRestakingWithdrawalIterator{contract: _EtherFiNode.contract, event: "QueuedRestakingWithdrawal", logs: logs, sub: sub}, nil
}

// WatchQueuedRestakingWithdrawal is a free log subscription operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherFiNode *EtherFiNodeFilterer) WatchQueuedRestakingWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherFiNodeQueuedRestakingWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNode.contract.WatchLogs(opts, "QueuedRestakingWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodeQueuedRestakingWithdrawal)
				if err := _EtherFiNode.contract.UnpackLog(event, "QueuedRestakingWithdrawal", log); err != nil {
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

// ParseQueuedRestakingWithdrawal is a log parse operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherFiNode *EtherFiNodeFilterer) ParseQueuedRestakingWithdrawal(log types.Log) (*EtherFiNodeQueuedRestakingWithdrawal, error) {
	event := new(EtherFiNodeQueuedRestakingWithdrawal)
	if err := _EtherFiNode.contract.UnpackLog(event, "QueuedRestakingWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
