// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// IEtherFiNodesManagerValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEtherFiNodesManagerValidatorInfo struct {
	ValidatorIndex       uint32
	ExitRequestTimestamp uint32
	ExitTimestamp        uint32
	Phase                uint8
}

// EtherFiNodesManagerMetaData contains all meta data concerning the EtherFiNodesManager contract.
var EtherFiNodesManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEPRECATED_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPRECATED_enableNodeRecycling\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPRECATED_protocolRevenueManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProtocolRevenueManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPRECATED_protocolRevenueManagerContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPRECATED_protocolRewardsSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"treasury\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nodeOperator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tnft\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bnft\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SCALE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocateEtherFiNode\",\"inputs\":[{\"name\":\"_enableRestaking\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"withdrawalSafeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowedForwardedEigenpodCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowedForwardedExternalCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAuctionManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchFullWithdraw\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchPartialWithdraw\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchQueueRestakedWithdrawal\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchSendExitRequest\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bnft\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBNFT\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateTVL\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_beaconBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"toNodeOperator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toTnft\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toBnft\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toTreasury\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"_receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delayedWithdrawalRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelayedWithdrawalRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"etherfiNodeAddress\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwardEigenpodCall\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardEigenpodCall\",\"inputs\":[{\"name\":\"_etherfiNodes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardExternalCall\",\"inputs\":[{\"name\":\"_etherfiNodes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardExternalCall\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fullWithdraw\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateWithdrawalCredentials\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getEigenPod\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFullWithdrawalPayouts\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"toNodeOperator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toTnft\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toBnft\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toTreasury\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonExitPenalty\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"nonExitPenalty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRewardsPayouts\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnusedWithdrawalSafesLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorInfo\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEtherFiNodesManager.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"exitRequestTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"exitTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"phase\",\"type\":\"uint8\",\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalCredentials\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalSafeAddress\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incrementNumberOfValidators\",\"inputs\":[{\"name\":\"_count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_treasuryContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_auctionContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakingManagerContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tnftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bnftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delayedWithdrawalRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isExitRequested\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markBeingSlashed\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxEigenlayerWithdrawals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonExitPenaltyDailyRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonExitPenaltyPrincipal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numAssociatedValidators\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numberOfValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatingAdmin\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"partialWithdraw\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"phase\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validatorPhase\",\"type\":\"uint8\",\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processNodeExit\",\"inputs\":[{\"name\":\"_validatorIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_exitTimestamps\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerValidator\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_enableRestaking\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_withdrawalSafeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxEigenLayerWithdrawals\",\"inputs\":[{\"name\":\"_max\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNonExitPenalty\",\"inputs\":[{\"name\":\"_nonExitPenaltyDailyRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_nonExitPenaltyPrincipal\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProofSubmitter\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_newProofSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStakingRewardsSplit\",\"inputs\":[{\"name\":\"_treasury\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_nodeOperator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_tnft\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_bnft\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidatorPhase\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_phase\",\"type\":\"uint8\",\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingManagerContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingRewardsSplit\",\"inputs\":[],\"outputs\":[{\"name\":\"treasury\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nodeOperator\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tnft\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bnft\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_revertIfNoBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tnft\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTNFT\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasuryContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unPauseContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unregisterValidator\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unusedWithdrawalSafes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdmin\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isAdmin\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedForwardedEigenpodCalls\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedForwardedExternalCalls\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateEigenLayerOperatingAdmin\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isAdmin\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateEtherFiNode\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedForwardedEigenpodCallsUpdated\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedForwardedExternalCallsUpdated\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"_target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FullWithdrawal\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toOperator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toBnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTreasury\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeExitProcessed\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeExitRequested\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialWithdrawal\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toOperator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toBnft\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTreasury\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PhaseChanged\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_phase\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIEtherFiNode.VALIDATOR_PHASE\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QueuedRestakingWithdrawal\",\"inputs\":[{\"name\":\"_validatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"etherFiNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInstalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForwardedCallNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEtherFiNodeVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidForwardedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPenaltyRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInstalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotStakingManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SendFail\",\"inputs\":[]}]",
}

// EtherFiNodesManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherFiNodesManagerMetaData.ABI instead.
var EtherFiNodesManagerABI = EtherFiNodesManagerMetaData.ABI

// EtherFiNodesManager is an auto generated Go binding around an Ethereum contract.
type EtherFiNodesManager struct {
	EtherFiNodesManagerCaller     // Read-only binding to the contract
	EtherFiNodesManagerTransactor // Write-only binding to the contract
	EtherFiNodesManagerFilterer   // Log filterer for contract events
}

// EtherFiNodesManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherFiNodesManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodesManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherFiNodesManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodesManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherFiNodesManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherFiNodesManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherFiNodesManagerSession struct {
	Contract     *EtherFiNodesManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EtherFiNodesManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherFiNodesManagerCallerSession struct {
	Contract *EtherFiNodesManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EtherFiNodesManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherFiNodesManagerTransactorSession struct {
	Contract     *EtherFiNodesManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EtherFiNodesManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherFiNodesManagerRaw struct {
	Contract *EtherFiNodesManager // Generic contract binding to access the raw methods on
}

// EtherFiNodesManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherFiNodesManagerCallerRaw struct {
	Contract *EtherFiNodesManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EtherFiNodesManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherFiNodesManagerTransactorRaw struct {
	Contract *EtherFiNodesManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherFiNodesManager creates a new instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManager(address common.Address, backend bind.ContractBackend) (*EtherFiNodesManager, error) {
	contract, err := bindEtherFiNodesManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManager{EtherFiNodesManagerCaller: EtherFiNodesManagerCaller{contract: contract}, EtherFiNodesManagerTransactor: EtherFiNodesManagerTransactor{contract: contract}, EtherFiNodesManagerFilterer: EtherFiNodesManagerFilterer{contract: contract}}, nil
}

// NewEtherFiNodesManagerCaller creates a new read-only instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManagerCaller(address common.Address, caller bind.ContractCaller) (*EtherFiNodesManagerCaller, error) {
	contract, err := bindEtherFiNodesManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerCaller{contract: contract}, nil
}

// NewEtherFiNodesManagerTransactor creates a new write-only instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherFiNodesManagerTransactor, error) {
	contract, err := bindEtherFiNodesManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerTransactor{contract: contract}, nil
}

// NewEtherFiNodesManagerFilterer creates a new log filterer instance of EtherFiNodesManager, bound to a specific deployed contract.
func NewEtherFiNodesManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherFiNodesManagerFilterer, error) {
	contract, err := bindEtherFiNodesManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerFilterer{contract: contract}, nil
}

// bindEtherFiNodesManager binds a generic wrapper to an already deployed contract.
func bindEtherFiNodesManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherFiNodesManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherFiNodesManager *EtherFiNodesManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherFiNodesManager.Contract.EtherFiNodesManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherFiNodesManager *EtherFiNodesManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.EtherFiNodesManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherFiNodesManager *EtherFiNodesManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.EtherFiNodesManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherFiNodesManager *EtherFiNodesManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherFiNodesManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.contract.Transact(opts, method, params...)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) DEPRECATEDAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "DEPRECATED_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDAdmin(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDAdmin(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDEnableNodeRecycling is a free data retrieval call binding the contract method 0x792cdc9c.
//
// Solidity: function DEPRECATED_enableNodeRecycling() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) DEPRECATEDEnableNodeRecycling(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "DEPRECATED_enableNodeRecycling")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DEPRECATEDEnableNodeRecycling is a free data retrieval call binding the contract method 0x792cdc9c.
//
// Solidity: function DEPRECATED_enableNodeRecycling() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) DEPRECATEDEnableNodeRecycling() (bool, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDEnableNodeRecycling(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDEnableNodeRecycling is a free data retrieval call binding the contract method 0x792cdc9c.
//
// Solidity: function DEPRECATED_enableNodeRecycling() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) DEPRECATEDEnableNodeRecycling() (bool, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDEnableNodeRecycling(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) DEPRECATEDProtocolRevenueManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "DEPRECATED_protocolRevenueManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) DEPRECATEDProtocolRevenueManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDProtocolRevenueManager(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManager is a free data retrieval call binding the contract method 0x2f708968.
//
// Solidity: function DEPRECATED_protocolRevenueManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) DEPRECATEDProtocolRevenueManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDProtocolRevenueManager(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManagerContract is a free data retrieval call binding the contract method 0x722395d5.
//
// Solidity: function DEPRECATED_protocolRevenueManagerContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) DEPRECATEDProtocolRevenueManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "DEPRECATED_protocolRevenueManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDProtocolRevenueManagerContract is a free data retrieval call binding the contract method 0x722395d5.
//
// Solidity: function DEPRECATED_protocolRevenueManagerContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) DEPRECATEDProtocolRevenueManagerContract() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDProtocolRevenueManagerContract(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRevenueManagerContract is a free data retrieval call binding the contract method 0x722395d5.
//
// Solidity: function DEPRECATED_protocolRevenueManagerContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) DEPRECATEDProtocolRevenueManagerContract() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDProtocolRevenueManagerContract(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRewardsSplit is a free data retrieval call binding the contract method 0xca692dc7.
//
// Solidity: function DEPRECATED_protocolRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) DEPRECATEDProtocolRewardsSplit(opts *bind.CallOpts) (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "DEPRECATED_protocolRewardsSplit")

	outstruct := new(struct {
		Treasury     uint64
		NodeOperator uint64
		Tnft         uint64
		Bnft         uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Treasury = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.NodeOperator = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Tnft = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Bnft = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// DEPRECATEDProtocolRewardsSplit is a free data retrieval call binding the contract method 0xca692dc7.
//
// Solidity: function DEPRECATED_protocolRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) DEPRECATEDProtocolRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDProtocolRewardsSplit(&_EtherFiNodesManager.CallOpts)
}

// DEPRECATEDProtocolRewardsSplit is a free data retrieval call binding the contract method 0xca692dc7.
//
// Solidity: function DEPRECATED_protocolRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) DEPRECATEDProtocolRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherFiNodesManager.Contract.DEPRECATEDProtocolRewardsSplit(&_EtherFiNodesManager.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) SCALE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "SCALE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SCALE() (uint64, error) {
	return _EtherFiNodesManager.Contract.SCALE(&_EtherFiNodesManager.CallOpts)
}

// SCALE is a free data retrieval call binding the contract method 0xeced5526.
//
// Solidity: function SCALE() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) SCALE() (uint64, error) {
	return _EtherFiNodesManager.Contract.SCALE(&_EtherFiNodesManager.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.Admins(&_EtherFiNodesManager.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.Admins(&_EtherFiNodesManager.CallOpts, arg0)
}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) AllowedForwardedEigenpodCalls(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "allowedForwardedEigenpodCalls", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) AllowedForwardedEigenpodCalls(arg0 [4]byte) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedEigenpodCalls(&_EtherFiNodesManager.CallOpts, arg0)
}

// AllowedForwardedEigenpodCalls is a free data retrieval call binding the contract method 0x93b572a0.
//
// Solidity: function allowedForwardedEigenpodCalls(bytes4 ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) AllowedForwardedEigenpodCalls(arg0 [4]byte) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedEigenpodCalls(&_EtherFiNodesManager.CallOpts, arg0)
}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) AllowedForwardedExternalCalls(opts *bind.CallOpts, arg0 [4]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "allowedForwardedExternalCalls", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) AllowedForwardedExternalCalls(arg0 [4]byte, arg1 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedExternalCalls(&_EtherFiNodesManager.CallOpts, arg0, arg1)
}

// AllowedForwardedExternalCalls is a free data retrieval call binding the contract method 0xd0cece05.
//
// Solidity: function allowedForwardedExternalCalls(bytes4 , address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) AllowedForwardedExternalCalls(arg0 [4]byte, arg1 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.AllowedForwardedExternalCalls(&_EtherFiNodesManager.CallOpts, arg0, arg1)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) AuctionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "auctionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) AuctionManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.AuctionManager(&_EtherFiNodesManager.CallOpts)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) AuctionManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.AuctionManager(&_EtherFiNodesManager.CallOpts)
}

// Bnft is a free data retrieval call binding the contract method 0x7bc92fd5.
//
// Solidity: function bnft() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Bnft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "bnft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bnft is a free data retrieval call binding the contract method 0x7bc92fd5.
//
// Solidity: function bnft() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Bnft() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Bnft(&_EtherFiNodesManager.CallOpts)
}

// Bnft is a free data retrieval call binding the contract method 0x7bc92fd5.
//
// Solidity: function bnft() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Bnft() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Bnft(&_EtherFiNodesManager.CallOpts)
}

// CalculateTVL is a free data retrieval call binding the contract method 0x30068a65.
//
// Solidity: function calculateTVL(uint256 _validatorId, uint256 _beaconBalance) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) CalculateTVL(opts *bind.CallOpts, _validatorId *big.Int, _beaconBalance *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "calculateTVL", _validatorId, _beaconBalance)

	outstruct := new(struct {
		ToNodeOperator *big.Int
		ToTnft         *big.Int
		ToBnft         *big.Int
		ToTreasury     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ToNodeOperator = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ToTnft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ToBnft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ToTreasury = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateTVL is a free data retrieval call binding the contract method 0x30068a65.
//
// Solidity: function calculateTVL(uint256 _validatorId, uint256 _beaconBalance) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) CalculateTVL(_validatorId *big.Int, _beaconBalance *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherFiNodesManager.Contract.CalculateTVL(&_EtherFiNodesManager.CallOpts, _validatorId, _beaconBalance)
}

// CalculateTVL is a free data retrieval call binding the contract method 0x30068a65.
//
// Solidity: function calculateTVL(uint256 _validatorId, uint256 _beaconBalance) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) CalculateTVL(_validatorId *big.Int, _beaconBalance *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherFiNodesManager.Contract.CalculateTVL(&_EtherFiNodesManager.CallOpts, _validatorId, _beaconBalance)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) DelayedWithdrawalRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "delayedWithdrawalRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DelayedWithdrawalRouter(&_EtherFiNodesManager.CallOpts)
}

// DelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x1a5057be.
//
// Solidity: function delayedWithdrawalRouter() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) DelayedWithdrawalRouter() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DelayedWithdrawalRouter(&_EtherFiNodesManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) DelegationManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DelegationManager(&_EtherFiNodesManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) DelegationManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.DelegationManager(&_EtherFiNodesManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) EigenPodManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.EigenPodManager(&_EtherFiNodesManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _EtherFiNodesManager.Contract.EigenPodManager(&_EtherFiNodesManager.CallOpts)
}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) EtherfiNodeAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "etherfiNodeAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) EtherfiNodeAddress(arg0 *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.EtherfiNodeAddress(&_EtherFiNodesManager.CallOpts, arg0)
}

// EtherfiNodeAddress is a free data retrieval call binding the contract method 0xb165e295.
//
// Solidity: function etherfiNodeAddress(uint256 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) EtherfiNodeAddress(arg0 *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.EtherfiNodeAddress(&_EtherFiNodesManager.CallOpts, arg0)
}

// GenerateWithdrawalCredentials is a free data retrieval call binding the contract method 0x2b5cfa81.
//
// Solidity: function generateWithdrawalCredentials(address _address) pure returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GenerateWithdrawalCredentials(opts *bind.CallOpts, _address common.Address) ([]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "generateWithdrawalCredentials", _address)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateWithdrawalCredentials is a free data retrieval call binding the contract method 0x2b5cfa81.
//
// Solidity: function generateWithdrawalCredentials(address _address) pure returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GenerateWithdrawalCredentials(_address common.Address) ([]byte, error) {
	return _EtherFiNodesManager.Contract.GenerateWithdrawalCredentials(&_EtherFiNodesManager.CallOpts, _address)
}

// GenerateWithdrawalCredentials is a free data retrieval call binding the contract method 0x2b5cfa81.
//
// Solidity: function generateWithdrawalCredentials(address _address) pure returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GenerateWithdrawalCredentials(_address common.Address) ([]byte, error) {
	return _EtherFiNodesManager.Contract.GenerateWithdrawalCredentials(&_EtherFiNodesManager.CallOpts, _address)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 _validatorId) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetEigenPod(opts *bind.CallOpts, _validatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getEigenPod", _validatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 _validatorId) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetEigenPod(_validatorId *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetEigenPod(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetEigenPod is a free data retrieval call binding the contract method 0xf3c148ec.
//
// Solidity: function getEigenPod(uint256 _validatorId) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetEigenPod(_validatorId *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetEigenPod(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetFullWithdrawalPayouts is a free data retrieval call binding the contract method 0x8edb719e.
//
// Solidity: function getFullWithdrawalPayouts(uint256 _validatorId) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetFullWithdrawalPayouts(opts *bind.CallOpts, _validatorId *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getFullWithdrawalPayouts", _validatorId)

	outstruct := new(struct {
		ToNodeOperator *big.Int
		ToTnft         *big.Int
		ToBnft         *big.Int
		ToTreasury     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ToNodeOperator = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ToTnft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ToBnft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ToTreasury = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFullWithdrawalPayouts is a free data retrieval call binding the contract method 0x8edb719e.
//
// Solidity: function getFullWithdrawalPayouts(uint256 _validatorId) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetFullWithdrawalPayouts(_validatorId *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherFiNodesManager.Contract.GetFullWithdrawalPayouts(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetFullWithdrawalPayouts is a free data retrieval call binding the contract method 0x8edb719e.
//
// Solidity: function getFullWithdrawalPayouts(uint256 _validatorId) view returns(uint256 toNodeOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetFullWithdrawalPayouts(_validatorId *big.Int) (struct {
	ToNodeOperator *big.Int
	ToTnft         *big.Int
	ToBnft         *big.Int
	ToTreasury     *big.Int
}, error) {
	return _EtherFiNodesManager.Contract.GetFullWithdrawalPayouts(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetImplementation() (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetImplementation(&_EtherFiNodesManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetImplementation() (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetImplementation(&_EtherFiNodesManager.CallOpts)
}

// GetNonExitPenalty is a free data retrieval call binding the contract method 0xf0a2ae91.
//
// Solidity: function getNonExitPenalty(uint256 _validatorId) view returns(uint256 nonExitPenalty)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetNonExitPenalty(opts *bind.CallOpts, _validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getNonExitPenalty", _validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonExitPenalty is a free data retrieval call binding the contract method 0xf0a2ae91.
//
// Solidity: function getNonExitPenalty(uint256 _validatorId) view returns(uint256 nonExitPenalty)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetNonExitPenalty(_validatorId *big.Int) (*big.Int, error) {
	return _EtherFiNodesManager.Contract.GetNonExitPenalty(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetNonExitPenalty is a free data retrieval call binding the contract method 0xf0a2ae91.
//
// Solidity: function getNonExitPenalty(uint256 _validatorId) view returns(uint256 nonExitPenalty)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetNonExitPenalty(_validatorId *big.Int) (*big.Int, error) {
	return _EtherFiNodesManager.Contract.GetNonExitPenalty(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetRewardsPayouts is a free data retrieval call binding the contract method 0x02e651c6.
//
// Solidity: function getRewardsPayouts(uint256 _validatorId) view returns(uint256, uint256, uint256, uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetRewardsPayouts(opts *bind.CallOpts, _validatorId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getRewardsPayouts", _validatorId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetRewardsPayouts is a free data retrieval call binding the contract method 0x02e651c6.
//
// Solidity: function getRewardsPayouts(uint256 _validatorId) view returns(uint256, uint256, uint256, uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetRewardsPayouts(_validatorId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EtherFiNodesManager.Contract.GetRewardsPayouts(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetRewardsPayouts is a free data retrieval call binding the contract method 0x02e651c6.
//
// Solidity: function getRewardsPayouts(uint256 _validatorId) view returns(uint256, uint256, uint256, uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetRewardsPayouts(_validatorId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EtherFiNodesManager.Contract.GetRewardsPayouts(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetUnusedWithdrawalSafesLength is a free data retrieval call binding the contract method 0x9e22f949.
//
// Solidity: function getUnusedWithdrawalSafesLength() view returns(uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetUnusedWithdrawalSafesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getUnusedWithdrawalSafesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnusedWithdrawalSafesLength is a free data retrieval call binding the contract method 0x9e22f949.
//
// Solidity: function getUnusedWithdrawalSafesLength() view returns(uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetUnusedWithdrawalSafesLength() (*big.Int, error) {
	return _EtherFiNodesManager.Contract.GetUnusedWithdrawalSafesLength(&_EtherFiNodesManager.CallOpts)
}

// GetUnusedWithdrawalSafesLength is a free data retrieval call binding the contract method 0x9e22f949.
//
// Solidity: function getUnusedWithdrawalSafesLength() view returns(uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetUnusedWithdrawalSafesLength() (*big.Int, error) {
	return _EtherFiNodesManager.Contract.GetUnusedWithdrawalSafesLength(&_EtherFiNodesManager.CallOpts)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 _validatorId) view returns((uint32,uint32,uint32,uint8))
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetValidatorInfo(opts *bind.CallOpts, _validatorId *big.Int) (IEtherFiNodesManagerValidatorInfo, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getValidatorInfo", _validatorId)

	if err != nil {
		return *new(IEtherFiNodesManagerValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEtherFiNodesManagerValidatorInfo)).(*IEtherFiNodesManagerValidatorInfo)

	return out0, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 _validatorId) view returns((uint32,uint32,uint32,uint8))
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetValidatorInfo(_validatorId *big.Int) (IEtherFiNodesManagerValidatorInfo, error) {
	return _EtherFiNodesManager.Contract.GetValidatorInfo(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0xb7797537.
//
// Solidity: function getValidatorInfo(uint256 _validatorId) view returns((uint32,uint32,uint32,uint8))
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetValidatorInfo(_validatorId *big.Int) (IEtherFiNodesManagerValidatorInfo, error) {
	return _EtherFiNodesManager.Contract.GetValidatorInfo(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x61669d27.
//
// Solidity: function getWithdrawalCredentials(uint256 _validatorId) view returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetWithdrawalCredentials(opts *bind.CallOpts, _validatorId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getWithdrawalCredentials", _validatorId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x61669d27.
//
// Solidity: function getWithdrawalCredentials(uint256 _validatorId) view returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetWithdrawalCredentials(_validatorId *big.Int) ([]byte, error) {
	return _EtherFiNodesManager.Contract.GetWithdrawalCredentials(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x61669d27.
//
// Solidity: function getWithdrawalCredentials(uint256 _validatorId) view returns(bytes)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetWithdrawalCredentials(_validatorId *big.Int) ([]byte, error) {
	return _EtherFiNodesManager.Contract.GetWithdrawalCredentials(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalSafeAddress is a free data retrieval call binding the contract method 0x84e1c393.
//
// Solidity: function getWithdrawalSafeAddress(uint256 _validatorId) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) GetWithdrawalSafeAddress(opts *bind.CallOpts, _validatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "getWithdrawalSafeAddress", _validatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWithdrawalSafeAddress is a free data retrieval call binding the contract method 0x84e1c393.
//
// Solidity: function getWithdrawalSafeAddress(uint256 _validatorId) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) GetWithdrawalSafeAddress(_validatorId *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetWithdrawalSafeAddress(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// GetWithdrawalSafeAddress is a free data retrieval call binding the contract method 0x84e1c393.
//
// Solidity: function getWithdrawalSafeAddress(uint256 _validatorId) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) GetWithdrawalSafeAddress(_validatorId *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.GetWithdrawalSafeAddress(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// IsExitRequested is a free data retrieval call binding the contract method 0x1babf0bf.
//
// Solidity: function isExitRequested(uint256 _validatorId) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) IsExitRequested(opts *bind.CallOpts, _validatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "isExitRequested", _validatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExitRequested is a free data retrieval call binding the contract method 0x1babf0bf.
//
// Solidity: function isExitRequested(uint256 _validatorId) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) IsExitRequested(_validatorId *big.Int) (bool, error) {
	return _EtherFiNodesManager.Contract.IsExitRequested(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// IsExitRequested is a free data retrieval call binding the contract method 0x1babf0bf.
//
// Solidity: function isExitRequested(uint256 _validatorId) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) IsExitRequested(_validatorId *big.Int) (bool, error) {
	return _EtherFiNodesManager.Contract.IsExitRequested(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// MaxEigenlayerWithdrawals is a free data retrieval call binding the contract method 0x45401c9b.
//
// Solidity: function maxEigenlayerWithdrawals() view returns(uint8)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) MaxEigenlayerWithdrawals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "maxEigenlayerWithdrawals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxEigenlayerWithdrawals is a free data retrieval call binding the contract method 0x45401c9b.
//
// Solidity: function maxEigenlayerWithdrawals() view returns(uint8)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) MaxEigenlayerWithdrawals() (uint8, error) {
	return _EtherFiNodesManager.Contract.MaxEigenlayerWithdrawals(&_EtherFiNodesManager.CallOpts)
}

// MaxEigenlayerWithdrawals is a free data retrieval call binding the contract method 0x45401c9b.
//
// Solidity: function maxEigenlayerWithdrawals() view returns(uint8)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) MaxEigenlayerWithdrawals() (uint8, error) {
	return _EtherFiNodesManager.Contract.MaxEigenlayerWithdrawals(&_EtherFiNodesManager.CallOpts)
}

// NonExitPenaltyDailyRate is a free data retrieval call binding the contract method 0x7082994b.
//
// Solidity: function nonExitPenaltyDailyRate() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) NonExitPenaltyDailyRate(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "nonExitPenaltyDailyRate")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonExitPenaltyDailyRate is a free data retrieval call binding the contract method 0x7082994b.
//
// Solidity: function nonExitPenaltyDailyRate() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) NonExitPenaltyDailyRate() (uint64, error) {
	return _EtherFiNodesManager.Contract.NonExitPenaltyDailyRate(&_EtherFiNodesManager.CallOpts)
}

// NonExitPenaltyDailyRate is a free data retrieval call binding the contract method 0x7082994b.
//
// Solidity: function nonExitPenaltyDailyRate() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) NonExitPenaltyDailyRate() (uint64, error) {
	return _EtherFiNodesManager.Contract.NonExitPenaltyDailyRate(&_EtherFiNodesManager.CallOpts)
}

// NonExitPenaltyPrincipal is a free data retrieval call binding the contract method 0xbbe78ecd.
//
// Solidity: function nonExitPenaltyPrincipal() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) NonExitPenaltyPrincipal(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "nonExitPenaltyPrincipal")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NonExitPenaltyPrincipal is a free data retrieval call binding the contract method 0xbbe78ecd.
//
// Solidity: function nonExitPenaltyPrincipal() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) NonExitPenaltyPrincipal() (uint64, error) {
	return _EtherFiNodesManager.Contract.NonExitPenaltyPrincipal(&_EtherFiNodesManager.CallOpts)
}

// NonExitPenaltyPrincipal is a free data retrieval call binding the contract method 0xbbe78ecd.
//
// Solidity: function nonExitPenaltyPrincipal() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) NonExitPenaltyPrincipal() (uint64, error) {
	return _EtherFiNodesManager.Contract.NonExitPenaltyPrincipal(&_EtherFiNodesManager.CallOpts)
}

// NumAssociatedValidators is a free data retrieval call binding the contract method 0xbaf08700.
//
// Solidity: function numAssociatedValidators(uint256 _validatorId) view returns(uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) NumAssociatedValidators(opts *bind.CallOpts, _validatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "numAssociatedValidators", _validatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssociatedValidators is a free data retrieval call binding the contract method 0xbaf08700.
//
// Solidity: function numAssociatedValidators(uint256 _validatorId) view returns(uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) NumAssociatedValidators(_validatorId *big.Int) (*big.Int, error) {
	return _EtherFiNodesManager.Contract.NumAssociatedValidators(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// NumAssociatedValidators is a free data retrieval call binding the contract method 0xbaf08700.
//
// Solidity: function numAssociatedValidators(uint256 _validatorId) view returns(uint256)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) NumAssociatedValidators(_validatorId *big.Int) (*big.Int, error) {
	return _EtherFiNodesManager.Contract.NumAssociatedValidators(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// NumberOfValidators is a free data retrieval call binding the contract method 0xd6832ea9.
//
// Solidity: function numberOfValidators() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) NumberOfValidators(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "numberOfValidators")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NumberOfValidators is a free data retrieval call binding the contract method 0xd6832ea9.
//
// Solidity: function numberOfValidators() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) NumberOfValidators() (uint64, error) {
	return _EtherFiNodesManager.Contract.NumberOfValidators(&_EtherFiNodesManager.CallOpts)
}

// NumberOfValidators is a free data retrieval call binding the contract method 0xd6832ea9.
//
// Solidity: function numberOfValidators() view returns(uint64)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) NumberOfValidators() (uint64, error) {
	return _EtherFiNodesManager.Contract.NumberOfValidators(&_EtherFiNodesManager.CallOpts)
}

// OperatingAdmin is a free data retrieval call binding the contract method 0xf9ffd6e0.
//
// Solidity: function operatingAdmin(address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) OperatingAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "operatingAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatingAdmin is a free data retrieval call binding the contract method 0xf9ffd6e0.
//
// Solidity: function operatingAdmin(address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) OperatingAdmin(arg0 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.OperatingAdmin(&_EtherFiNodesManager.CallOpts, arg0)
}

// OperatingAdmin is a free data retrieval call binding the contract method 0xf9ffd6e0.
//
// Solidity: function operatingAdmin(address ) view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) OperatingAdmin(arg0 common.Address) (bool, error) {
	return _EtherFiNodesManager.Contract.OperatingAdmin(&_EtherFiNodesManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Owner() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Owner(&_EtherFiNodesManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Owner() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Owner(&_EtherFiNodesManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Paused() (bool, error) {
	return _EtherFiNodesManager.Contract.Paused(&_EtherFiNodesManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Paused() (bool, error) {
	return _EtherFiNodesManager.Contract.Paused(&_EtherFiNodesManager.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 _validatorId) view returns(uint8 validatorPhase)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Phase(opts *bind.CallOpts, _validatorId *big.Int) (uint8, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "phase", _validatorId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 _validatorId) view returns(uint8 validatorPhase)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Phase(_validatorId *big.Int) (uint8, error) {
	return _EtherFiNodesManager.Contract.Phase(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 _validatorId) view returns(uint8 validatorPhase)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Phase(_validatorId *big.Int) (uint8, error) {
	return _EtherFiNodesManager.Contract.Phase(&_EtherFiNodesManager.CallOpts, _validatorId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ProxiableUUID(&_EtherFiNodesManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherFiNodesManager.Contract.ProxiableUUID(&_EtherFiNodesManager.CallOpts)
}

// StakingManagerContract is a free data retrieval call binding the contract method 0x4f608156.
//
// Solidity: function stakingManagerContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) StakingManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "stakingManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingManagerContract is a free data retrieval call binding the contract method 0x4f608156.
//
// Solidity: function stakingManagerContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) StakingManagerContract() (common.Address, error) {
	return _EtherFiNodesManager.Contract.StakingManagerContract(&_EtherFiNodesManager.CallOpts)
}

// StakingManagerContract is a free data retrieval call binding the contract method 0x4f608156.
//
// Solidity: function stakingManagerContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) StakingManagerContract() (common.Address, error) {
	return _EtherFiNodesManager.Contract.StakingManagerContract(&_EtherFiNodesManager.CallOpts)
}

// StakingRewardsSplit is a free data retrieval call binding the contract method 0x62f7b332.
//
// Solidity: function stakingRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) StakingRewardsSplit(opts *bind.CallOpts) (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "stakingRewardsSplit")

	outstruct := new(struct {
		Treasury     uint64
		NodeOperator uint64
		Tnft         uint64
		Bnft         uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Treasury = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.NodeOperator = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Tnft = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Bnft = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// StakingRewardsSplit is a free data retrieval call binding the contract method 0x62f7b332.
//
// Solidity: function stakingRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) StakingRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherFiNodesManager.Contract.StakingRewardsSplit(&_EtherFiNodesManager.CallOpts)
}

// StakingRewardsSplit is a free data retrieval call binding the contract method 0x62f7b332.
//
// Solidity: function stakingRewardsSplit() view returns(uint64 treasury, uint64 nodeOperator, uint64 tnft, uint64 bnft)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) StakingRewardsSplit() (struct {
	Treasury     uint64
	NodeOperator uint64
	Tnft         uint64
	Bnft         uint64
}, error) {
	return _EtherFiNodesManager.Contract.StakingRewardsSplit(&_EtherFiNodesManager.CallOpts)
}

// Tnft is a free data retrieval call binding the contract method 0xad36cd0e.
//
// Solidity: function tnft() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) Tnft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "tnft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tnft is a free data retrieval call binding the contract method 0xad36cd0e.
//
// Solidity: function tnft() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Tnft() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Tnft(&_EtherFiNodesManager.CallOpts)
}

// Tnft is a free data retrieval call binding the contract method 0xad36cd0e.
//
// Solidity: function tnft() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) Tnft() (common.Address, error) {
	return _EtherFiNodesManager.Contract.Tnft(&_EtherFiNodesManager.CallOpts)
}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) TreasuryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "treasuryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) TreasuryContract() (common.Address, error) {
	return _EtherFiNodesManager.Contract.TreasuryContract(&_EtherFiNodesManager.CallOpts)
}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) TreasuryContract() (common.Address, error) {
	return _EtherFiNodesManager.Contract.TreasuryContract(&_EtherFiNodesManager.CallOpts)
}

// UnusedWithdrawalSafes is a free data retrieval call binding the contract method 0x4c3551bd.
//
// Solidity: function unusedWithdrawalSafes(uint256 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCaller) UnusedWithdrawalSafes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherFiNodesManager.contract.Call(opts, &out, "unusedWithdrawalSafes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnusedWithdrawalSafes is a free data retrieval call binding the contract method 0x4c3551bd.
//
// Solidity: function unusedWithdrawalSafes(uint256 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UnusedWithdrawalSafes(arg0 *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.UnusedWithdrawalSafes(&_EtherFiNodesManager.CallOpts, arg0)
}

// UnusedWithdrawalSafes is a free data retrieval call binding the contract method 0x4c3551bd.
//
// Solidity: function unusedWithdrawalSafes(uint256 ) view returns(address)
func (_EtherFiNodesManager *EtherFiNodesManagerCallerSession) UnusedWithdrawalSafes(arg0 *big.Int) (common.Address, error) {
	return _EtherFiNodesManager.Contract.UnusedWithdrawalSafes(&_EtherFiNodesManager.CallOpts, arg0)
}

// AllocateEtherFiNode is a paid mutator transaction binding the contract method 0x15ef0e5e.
//
// Solidity: function allocateEtherFiNode(bool _enableRestaking) returns(address withdrawalSafeAddress)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) AllocateEtherFiNode(opts *bind.TransactOpts, _enableRestaking bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "allocateEtherFiNode", _enableRestaking)
}

// AllocateEtherFiNode is a paid mutator transaction binding the contract method 0x15ef0e5e.
//
// Solidity: function allocateEtherFiNode(bool _enableRestaking) returns(address withdrawalSafeAddress)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) AllocateEtherFiNode(_enableRestaking bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.AllocateEtherFiNode(&_EtherFiNodesManager.TransactOpts, _enableRestaking)
}

// AllocateEtherFiNode is a paid mutator transaction binding the contract method 0x15ef0e5e.
//
// Solidity: function allocateEtherFiNode(bool _enableRestaking) returns(address withdrawalSafeAddress)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) AllocateEtherFiNode(_enableRestaking bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.AllocateEtherFiNode(&_EtherFiNodesManager.TransactOpts, _enableRestaking)
}

// BatchFullWithdraw is a paid mutator transaction binding the contract method 0x00373389.
//
// Solidity: function batchFullWithdraw(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) BatchFullWithdraw(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "batchFullWithdraw", _validatorIds)
}

// BatchFullWithdraw is a paid mutator transaction binding the contract method 0x00373389.
//
// Solidity: function batchFullWithdraw(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) BatchFullWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchFullWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// BatchFullWithdraw is a paid mutator transaction binding the contract method 0x00373389.
//
// Solidity: function batchFullWithdraw(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) BatchFullWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchFullWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// BatchPartialWithdraw is a paid mutator transaction binding the contract method 0xabb565d7.
//
// Solidity: function batchPartialWithdraw(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) BatchPartialWithdraw(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "batchPartialWithdraw", _validatorIds)
}

// BatchPartialWithdraw is a paid mutator transaction binding the contract method 0xabb565d7.
//
// Solidity: function batchPartialWithdraw(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) BatchPartialWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchPartialWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// BatchPartialWithdraw is a paid mutator transaction binding the contract method 0xabb565d7.
//
// Solidity: function batchPartialWithdraw(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) BatchPartialWithdraw(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchPartialWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// BatchQueueRestakedWithdrawal is a paid mutator transaction binding the contract method 0xad35567b.
//
// Solidity: function batchQueueRestakedWithdrawal(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) BatchQueueRestakedWithdrawal(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "batchQueueRestakedWithdrawal", _validatorIds)
}

// BatchQueueRestakedWithdrawal is a paid mutator transaction binding the contract method 0xad35567b.
//
// Solidity: function batchQueueRestakedWithdrawal(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) BatchQueueRestakedWithdrawal(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchQueueRestakedWithdrawal(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// BatchQueueRestakedWithdrawal is a paid mutator transaction binding the contract method 0xad35567b.
//
// Solidity: function batchQueueRestakedWithdrawal(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) BatchQueueRestakedWithdrawal(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchQueueRestakedWithdrawal(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// BatchSendExitRequest is a paid mutator transaction binding the contract method 0xfb63cf5c.
//
// Solidity: function batchSendExitRequest(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) BatchSendExitRequest(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "batchSendExitRequest", _validatorIds)
}

// BatchSendExitRequest is a paid mutator transaction binding the contract method 0xfb63cf5c.
//
// Solidity: function batchSendExitRequest(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) BatchSendExitRequest(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchSendExitRequest(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// BatchSendExitRequest is a paid mutator transaction binding the contract method 0xfb63cf5c.
//
// Solidity: function batchSendExitRequest(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) BatchSendExitRequest(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.BatchSendExitRequest(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xb42b396c.
//
// Solidity: function completeQueuedWithdrawals(uint256[] _validatorIds, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, bool _receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, _validatorIds []*big.Int, withdrawals []IDelegationManagerTypesWithdrawal, _receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "completeQueuedWithdrawals", _validatorIds, withdrawals, _receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xb42b396c.
//
// Solidity: function completeQueuedWithdrawals(uint256[] _validatorIds, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, bool _receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) CompleteQueuedWithdrawals(_validatorIds []*big.Int, withdrawals []IDelegationManagerTypesWithdrawal, _receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.CompleteQueuedWithdrawals(&_EtherFiNodesManager.TransactOpts, _validatorIds, withdrawals, _receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0xb42b396c.
//
// Solidity: function completeQueuedWithdrawals(uint256[] _validatorIds, (address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, bool _receiveAsTokens) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) CompleteQueuedWithdrawals(_validatorIds []*big.Int, withdrawals []IDelegationManagerTypesWithdrawal, _receiveAsTokens bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.CompleteQueuedWithdrawals(&_EtherFiNodesManager.TransactOpts, _validatorIds, withdrawals, _receiveAsTokens)
}

// ForwardEigenpodCall is a paid mutator transaction binding the contract method 0x3016ef6f.
//
// Solidity: function forwardEigenpodCall(uint256[] _validatorIds, bytes[] _data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) ForwardEigenpodCall(opts *bind.TransactOpts, _validatorIds []*big.Int, _data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "forwardEigenpodCall", _validatorIds, _data)
}

// ForwardEigenpodCall is a paid mutator transaction binding the contract method 0x3016ef6f.
//
// Solidity: function forwardEigenpodCall(uint256[] _validatorIds, bytes[] _data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ForwardEigenpodCall(_validatorIds []*big.Int, _data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardEigenpodCall(&_EtherFiNodesManager.TransactOpts, _validatorIds, _data)
}

// ForwardEigenpodCall is a paid mutator transaction binding the contract method 0x3016ef6f.
//
// Solidity: function forwardEigenpodCall(uint256[] _validatorIds, bytes[] _data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) ForwardEigenpodCall(_validatorIds []*big.Int, _data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardEigenpodCall(&_EtherFiNodesManager.TransactOpts, _validatorIds, _data)
}

// ForwardEigenpodCall0 is a paid mutator transaction binding the contract method 0xcb8d4c59.
//
// Solidity: function forwardEigenpodCall(address[] _etherfiNodes, bytes[] _data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) ForwardEigenpodCall0(opts *bind.TransactOpts, _etherfiNodes []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "forwardEigenpodCall0", _etherfiNodes, _data)
}

// ForwardEigenpodCall0 is a paid mutator transaction binding the contract method 0xcb8d4c59.
//
// Solidity: function forwardEigenpodCall(address[] _etherfiNodes, bytes[] _data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ForwardEigenpodCall0(_etherfiNodes []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardEigenpodCall0(&_EtherFiNodesManager.TransactOpts, _etherfiNodes, _data)
}

// ForwardEigenpodCall0 is a paid mutator transaction binding the contract method 0xcb8d4c59.
//
// Solidity: function forwardEigenpodCall(address[] _etherfiNodes, bytes[] _data) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) ForwardEigenpodCall0(_etherfiNodes []common.Address, _data [][]byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardEigenpodCall0(&_EtherFiNodesManager.TransactOpts, _etherfiNodes, _data)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x9fab3743.
//
// Solidity: function forwardExternalCall(address[] _etherfiNodes, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) ForwardExternalCall(opts *bind.TransactOpts, _etherfiNodes []common.Address, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "forwardExternalCall", _etherfiNodes, _data, _target)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x9fab3743.
//
// Solidity: function forwardExternalCall(address[] _etherfiNodes, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ForwardExternalCall(_etherfiNodes []common.Address, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardExternalCall(&_EtherFiNodesManager.TransactOpts, _etherfiNodes, _data, _target)
}

// ForwardExternalCall is a paid mutator transaction binding the contract method 0x9fab3743.
//
// Solidity: function forwardExternalCall(address[] _etherfiNodes, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) ForwardExternalCall(_etherfiNodes []common.Address, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardExternalCall(&_EtherFiNodesManager.TransactOpts, _etherfiNodes, _data, _target)
}

// ForwardExternalCall0 is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] _validatorIds, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) ForwardExternalCall0(opts *bind.TransactOpts, _validatorIds []*big.Int, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "forwardExternalCall0", _validatorIds, _data, _target)
}

// ForwardExternalCall0 is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] _validatorIds, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ForwardExternalCall0(_validatorIds []*big.Int, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardExternalCall0(&_EtherFiNodesManager.TransactOpts, _validatorIds, _data, _target)
}

// ForwardExternalCall0 is a paid mutator transaction binding the contract method 0xf1b76911.
//
// Solidity: function forwardExternalCall(uint256[] _validatorIds, bytes[] _data, address _target) returns(bytes[] returnData)
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) ForwardExternalCall0(_validatorIds []*big.Int, _data [][]byte, _target common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ForwardExternalCall0(&_EtherFiNodesManager.TransactOpts, _validatorIds, _data, _target)
}

// FullWithdraw is a paid mutator transaction binding the contract method 0x66e704bf.
//
// Solidity: function fullWithdraw(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) FullWithdraw(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "fullWithdraw", _validatorId)
}

// FullWithdraw is a paid mutator transaction binding the contract method 0x66e704bf.
//
// Solidity: function fullWithdraw(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) FullWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.FullWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// FullWithdraw is a paid mutator transaction binding the contract method 0x66e704bf.
//
// Solidity: function fullWithdraw(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) FullWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.FullWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// IncrementNumberOfValidators is a paid mutator transaction binding the contract method 0x790833d4.
//
// Solidity: function incrementNumberOfValidators(uint64 _count) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) IncrementNumberOfValidators(opts *bind.TransactOpts, _count uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "incrementNumberOfValidators", _count)
}

// IncrementNumberOfValidators is a paid mutator transaction binding the contract method 0x790833d4.
//
// Solidity: function incrementNumberOfValidators(uint64 _count) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) IncrementNumberOfValidators(_count uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.IncrementNumberOfValidators(&_EtherFiNodesManager.TransactOpts, _count)
}

// IncrementNumberOfValidators is a paid mutator transaction binding the contract method 0x790833d4.
//
// Solidity: function incrementNumberOfValidators(uint64 _count) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) IncrementNumberOfValidators(_count uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.IncrementNumberOfValidators(&_EtherFiNodesManager.TransactOpts, _count)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _treasuryContract, address _auctionContract, address _stakingManagerContract, address _tnftContract, address _bnftContract, address _eigenPodManager, address _delayedWithdrawalRouter, address _delegationManager) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) Initialize(opts *bind.TransactOpts, _treasuryContract common.Address, _auctionContract common.Address, _stakingManagerContract common.Address, _tnftContract common.Address, _bnftContract common.Address, _eigenPodManager common.Address, _delayedWithdrawalRouter common.Address, _delegationManager common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "initialize", _treasuryContract, _auctionContract, _stakingManagerContract, _tnftContract, _bnftContract, _eigenPodManager, _delayedWithdrawalRouter, _delegationManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _treasuryContract, address _auctionContract, address _stakingManagerContract, address _tnftContract, address _bnftContract, address _eigenPodManager, address _delayedWithdrawalRouter, address _delegationManager) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Initialize(_treasuryContract common.Address, _auctionContract common.Address, _stakingManagerContract common.Address, _tnftContract common.Address, _bnftContract common.Address, _eigenPodManager common.Address, _delayedWithdrawalRouter common.Address, _delegationManager common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.Initialize(&_EtherFiNodesManager.TransactOpts, _treasuryContract, _auctionContract, _stakingManagerContract, _tnftContract, _bnftContract, _eigenPodManager, _delayedWithdrawalRouter, _delegationManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _treasuryContract, address _auctionContract, address _stakingManagerContract, address _tnftContract, address _bnftContract, address _eigenPodManager, address _delayedWithdrawalRouter, address _delegationManager) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) Initialize(_treasuryContract common.Address, _auctionContract common.Address, _stakingManagerContract common.Address, _tnftContract common.Address, _bnftContract common.Address, _eigenPodManager common.Address, _delayedWithdrawalRouter common.Address, _delegationManager common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.Initialize(&_EtherFiNodesManager.TransactOpts, _treasuryContract, _auctionContract, _stakingManagerContract, _tnftContract, _bnftContract, _eigenPodManager, _delayedWithdrawalRouter, _delegationManager)
}

// MarkBeingSlashed is a paid mutator transaction binding the contract method 0x36017df5.
//
// Solidity: function markBeingSlashed(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) MarkBeingSlashed(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "markBeingSlashed", _validatorIds)
}

// MarkBeingSlashed is a paid mutator transaction binding the contract method 0x36017df5.
//
// Solidity: function markBeingSlashed(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) MarkBeingSlashed(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.MarkBeingSlashed(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// MarkBeingSlashed is a paid mutator transaction binding the contract method 0x36017df5.
//
// Solidity: function markBeingSlashed(uint256[] _validatorIds) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) MarkBeingSlashed(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.MarkBeingSlashed(&_EtherFiNodesManager.TransactOpts, _validatorIds)
}

// PartialWithdraw is a paid mutator transaction binding the contract method 0x71d2ee6c.
//
// Solidity: function partialWithdraw(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) PartialWithdraw(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "partialWithdraw", _validatorId)
}

// PartialWithdraw is a paid mutator transaction binding the contract method 0x71d2ee6c.
//
// Solidity: function partialWithdraw(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) PartialWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.PartialWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// PartialWithdraw is a paid mutator transaction binding the contract method 0x71d2ee6c.
//
// Solidity: function partialWithdraw(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) PartialWithdraw(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.PartialWithdraw(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) PauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.PauseContract(&_EtherFiNodesManager.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) PauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.PauseContract(&_EtherFiNodesManager.TransactOpts)
}

// ProcessNodeExit is a paid mutator transaction binding the contract method 0x53000b9b.
//
// Solidity: function processNodeExit(uint256[] _validatorIds, uint32[] _exitTimestamps) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) ProcessNodeExit(opts *bind.TransactOpts, _validatorIds []*big.Int, _exitTimestamps []uint32) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "processNodeExit", _validatorIds, _exitTimestamps)
}

// ProcessNodeExit is a paid mutator transaction binding the contract method 0x53000b9b.
//
// Solidity: function processNodeExit(uint256[] _validatorIds, uint32[] _exitTimestamps) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) ProcessNodeExit(_validatorIds []*big.Int, _exitTimestamps []uint32) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ProcessNodeExit(&_EtherFiNodesManager.TransactOpts, _validatorIds, _exitTimestamps)
}

// ProcessNodeExit is a paid mutator transaction binding the contract method 0x53000b9b.
//
// Solidity: function processNodeExit(uint256[] _validatorIds, uint32[] _exitTimestamps) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) ProcessNodeExit(_validatorIds []*big.Int, _exitTimestamps []uint32) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.ProcessNodeExit(&_EtherFiNodesManager.TransactOpts, _validatorIds, _exitTimestamps)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x387dcbc1.
//
// Solidity: function registerValidator(uint256 _validatorId, bool _enableRestaking, address _withdrawalSafeAddress) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) RegisterValidator(opts *bind.TransactOpts, _validatorId *big.Int, _enableRestaking bool, _withdrawalSafeAddress common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "registerValidator", _validatorId, _enableRestaking, _withdrawalSafeAddress)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x387dcbc1.
//
// Solidity: function registerValidator(uint256 _validatorId, bool _enableRestaking, address _withdrawalSafeAddress) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) RegisterValidator(_validatorId *big.Int, _enableRestaking bool, _withdrawalSafeAddress common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.RegisterValidator(&_EtherFiNodesManager.TransactOpts, _validatorId, _enableRestaking, _withdrawalSafeAddress)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x387dcbc1.
//
// Solidity: function registerValidator(uint256 _validatorId, bool _enableRestaking, address _withdrawalSafeAddress) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) RegisterValidator(_validatorId *big.Int, _enableRestaking bool, _withdrawalSafeAddress common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.RegisterValidator(&_EtherFiNodesManager.TransactOpts, _validatorId, _enableRestaking, _withdrawalSafeAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.RenounceOwnership(&_EtherFiNodesManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.RenounceOwnership(&_EtherFiNodesManager.TransactOpts)
}

// SetMaxEigenLayerWithdrawals is a paid mutator transaction binding the contract method 0xeeba122a.
//
// Solidity: function setMaxEigenLayerWithdrawals(uint8 _max) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) SetMaxEigenLayerWithdrawals(opts *bind.TransactOpts, _max uint8) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "setMaxEigenLayerWithdrawals", _max)
}

// SetMaxEigenLayerWithdrawals is a paid mutator transaction binding the contract method 0xeeba122a.
//
// Solidity: function setMaxEigenLayerWithdrawals(uint8 _max) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SetMaxEigenLayerWithdrawals(_max uint8) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetMaxEigenLayerWithdrawals(&_EtherFiNodesManager.TransactOpts, _max)
}

// SetMaxEigenLayerWithdrawals is a paid mutator transaction binding the contract method 0xeeba122a.
//
// Solidity: function setMaxEigenLayerWithdrawals(uint8 _max) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) SetMaxEigenLayerWithdrawals(_max uint8) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetMaxEigenLayerWithdrawals(&_EtherFiNodesManager.TransactOpts, _max)
}

// SetNonExitPenalty is a paid mutator transaction binding the contract method 0x0701d306.
//
// Solidity: function setNonExitPenalty(uint64 _nonExitPenaltyDailyRate, uint64 _nonExitPenaltyPrincipal) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) SetNonExitPenalty(opts *bind.TransactOpts, _nonExitPenaltyDailyRate uint64, _nonExitPenaltyPrincipal uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "setNonExitPenalty", _nonExitPenaltyDailyRate, _nonExitPenaltyPrincipal)
}

// SetNonExitPenalty is a paid mutator transaction binding the contract method 0x0701d306.
//
// Solidity: function setNonExitPenalty(uint64 _nonExitPenaltyDailyRate, uint64 _nonExitPenaltyPrincipal) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SetNonExitPenalty(_nonExitPenaltyDailyRate uint64, _nonExitPenaltyPrincipal uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetNonExitPenalty(&_EtherFiNodesManager.TransactOpts, _nonExitPenaltyDailyRate, _nonExitPenaltyPrincipal)
}

// SetNonExitPenalty is a paid mutator transaction binding the contract method 0x0701d306.
//
// Solidity: function setNonExitPenalty(uint64 _nonExitPenaltyDailyRate, uint64 _nonExitPenaltyPrincipal) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) SetNonExitPenalty(_nonExitPenaltyDailyRate uint64, _nonExitPenaltyPrincipal uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetNonExitPenalty(&_EtherFiNodesManager.TransactOpts, _nonExitPenaltyDailyRate, _nonExitPenaltyPrincipal)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 _validatorId, address _newProofSubmitter) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) SetProofSubmitter(opts *bind.TransactOpts, _validatorId *big.Int, _newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "setProofSubmitter", _validatorId, _newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 _validatorId, address _newProofSubmitter) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SetProofSubmitter(_validatorId *big.Int, _newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetProofSubmitter(&_EtherFiNodesManager.TransactOpts, _validatorId, _newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xf8047125.
//
// Solidity: function setProofSubmitter(uint256 _validatorId, address _newProofSubmitter) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) SetProofSubmitter(_validatorId *big.Int, _newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetProofSubmitter(&_EtherFiNodesManager.TransactOpts, _validatorId, _newProofSubmitter)
}

// SetStakingRewardsSplit is a paid mutator transaction binding the contract method 0xb1257a7b.
//
// Solidity: function setStakingRewardsSplit(uint64 _treasury, uint64 _nodeOperator, uint64 _tnft, uint64 _bnft) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) SetStakingRewardsSplit(opts *bind.TransactOpts, _treasury uint64, _nodeOperator uint64, _tnft uint64, _bnft uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "setStakingRewardsSplit", _treasury, _nodeOperator, _tnft, _bnft)
}

// SetStakingRewardsSplit is a paid mutator transaction binding the contract method 0xb1257a7b.
//
// Solidity: function setStakingRewardsSplit(uint64 _treasury, uint64 _nodeOperator, uint64 _tnft, uint64 _bnft) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SetStakingRewardsSplit(_treasury uint64, _nodeOperator uint64, _tnft uint64, _bnft uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetStakingRewardsSplit(&_EtherFiNodesManager.TransactOpts, _treasury, _nodeOperator, _tnft, _bnft)
}

// SetStakingRewardsSplit is a paid mutator transaction binding the contract method 0xb1257a7b.
//
// Solidity: function setStakingRewardsSplit(uint64 _treasury, uint64 _nodeOperator, uint64 _tnft, uint64 _bnft) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) SetStakingRewardsSplit(_treasury uint64, _nodeOperator uint64, _tnft uint64, _bnft uint64) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetStakingRewardsSplit(&_EtherFiNodesManager.TransactOpts, _treasury, _nodeOperator, _tnft, _bnft)
}

// SetValidatorPhase is a paid mutator transaction binding the contract method 0x59b65fbc.
//
// Solidity: function setValidatorPhase(uint256 _validatorId, uint8 _phase) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) SetValidatorPhase(opts *bind.TransactOpts, _validatorId *big.Int, _phase uint8) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "setValidatorPhase", _validatorId, _phase)
}

// SetValidatorPhase is a paid mutator transaction binding the contract method 0x59b65fbc.
//
// Solidity: function setValidatorPhase(uint256 _validatorId, uint8 _phase) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) SetValidatorPhase(_validatorId *big.Int, _phase uint8) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetValidatorPhase(&_EtherFiNodesManager.TransactOpts, _validatorId, _phase)
}

// SetValidatorPhase is a paid mutator transaction binding the contract method 0x59b65fbc.
//
// Solidity: function setValidatorPhase(uint256 _validatorId, uint8 _phase) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) SetValidatorPhase(_validatorId *big.Int, _phase uint8) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.SetValidatorPhase(&_EtherFiNodesManager.TransactOpts, _validatorId, _phase)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0xd4a9d4a7.
//
// Solidity: function startCheckpoint(uint256 _validatorId, bool _revertIfNoBalance) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) StartCheckpoint(opts *bind.TransactOpts, _validatorId *big.Int, _revertIfNoBalance bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "startCheckpoint", _validatorId, _revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0xd4a9d4a7.
//
// Solidity: function startCheckpoint(uint256 _validatorId, bool _revertIfNoBalance) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) StartCheckpoint(_validatorId *big.Int, _revertIfNoBalance bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.StartCheckpoint(&_EtherFiNodesManager.TransactOpts, _validatorId, _revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0xd4a9d4a7.
//
// Solidity: function startCheckpoint(uint256 _validatorId, bool _revertIfNoBalance) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) StartCheckpoint(_validatorId *big.Int, _revertIfNoBalance bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.StartCheckpoint(&_EtherFiNodesManager.TransactOpts, _validatorId, _revertIfNoBalance)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.TransferOwnership(&_EtherFiNodesManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.TransferOwnership(&_EtherFiNodesManager.TransactOpts, newOwner)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UnPauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "unPauseContract")
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UnPauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UnPauseContract(&_EtherFiNodesManager.TransactOpts)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UnPauseContract() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UnPauseContract(&_EtherFiNodesManager.TransactOpts)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0xaed18c8d.
//
// Solidity: function unregisterValidator(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UnregisterValidator(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "unregisterValidator", _validatorId)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0xaed18c8d.
//
// Solidity: function unregisterValidator(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UnregisterValidator(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UnregisterValidator(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0xaed18c8d.
//
// Solidity: function unregisterValidator(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UnregisterValidator(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UnregisterValidator(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAdmin(&_EtherFiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAdmin(&_EtherFiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 _selector, bool _allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpdateAllowedForwardedEigenpodCalls(opts *bind.TransactOpts, _selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "updateAllowedForwardedEigenpodCalls", _selector, _allowed)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 _selector, bool _allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpdateAllowedForwardedEigenpodCalls(_selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedEigenpodCalls(&_EtherFiNodesManager.TransactOpts, _selector, _allowed)
}

// UpdateAllowedForwardedEigenpodCalls is a paid mutator transaction binding the contract method 0x4cba6c74.
//
// Solidity: function updateAllowedForwardedEigenpodCalls(bytes4 _selector, bool _allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpdateAllowedForwardedEigenpodCalls(_selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedEigenpodCalls(&_EtherFiNodesManager.TransactOpts, _selector, _allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 _selector, address _target, bool _allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpdateAllowedForwardedExternalCalls(opts *bind.TransactOpts, _selector [4]byte, _target common.Address, _allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "updateAllowedForwardedExternalCalls", _selector, _target, _allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 _selector, address _target, bool _allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpdateAllowedForwardedExternalCalls(_selector [4]byte, _target common.Address, _allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedExternalCalls(&_EtherFiNodesManager.TransactOpts, _selector, _target, _allowed)
}

// UpdateAllowedForwardedExternalCalls is a paid mutator transaction binding the contract method 0xc7f61eec.
//
// Solidity: function updateAllowedForwardedExternalCalls(bytes4 _selector, address _target, bool _allowed) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpdateAllowedForwardedExternalCalls(_selector [4]byte, _target common.Address, _allowed bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateAllowedForwardedExternalCalls(&_EtherFiNodesManager.TransactOpts, _selector, _target, _allowed)
}

// UpdateEigenLayerOperatingAdmin is a paid mutator transaction binding the contract method 0xb57dade3.
//
// Solidity: function updateEigenLayerOperatingAdmin(address _address, bool _isAdmin) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpdateEigenLayerOperatingAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "updateEigenLayerOperatingAdmin", _address, _isAdmin)
}

// UpdateEigenLayerOperatingAdmin is a paid mutator transaction binding the contract method 0xb57dade3.
//
// Solidity: function updateEigenLayerOperatingAdmin(address _address, bool _isAdmin) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpdateEigenLayerOperatingAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateEigenLayerOperatingAdmin(&_EtherFiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateEigenLayerOperatingAdmin is a paid mutator transaction binding the contract method 0xb57dade3.
//
// Solidity: function updateEigenLayerOperatingAdmin(address _address, bool _isAdmin) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpdateEigenLayerOperatingAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateEigenLayerOperatingAdmin(&_EtherFiNodesManager.TransactOpts, _address, _isAdmin)
}

// UpdateEtherFiNode is a paid mutator transaction binding the contract method 0x936fb00c.
//
// Solidity: function updateEtherFiNode(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpdateEtherFiNode(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "updateEtherFiNode", _validatorId)
}

// UpdateEtherFiNode is a paid mutator transaction binding the contract method 0x936fb00c.
//
// Solidity: function updateEtherFiNode(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpdateEtherFiNode(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateEtherFiNode(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// UpdateEtherFiNode is a paid mutator transaction binding the contract method 0x936fb00c.
//
// Solidity: function updateEtherFiNode(uint256 _validatorId) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpdateEtherFiNode(_validatorId *big.Int) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpdateEtherFiNode(&_EtherFiNodesManager.TransactOpts, _validatorId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeTo(&_EtherFiNodesManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeTo(&_EtherFiNodesManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeToAndCall(&_EtherFiNodesManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.UpgradeToAndCall(&_EtherFiNodesManager.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherFiNodesManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerSession) Receive() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.Receive(&_EtherFiNodesManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherFiNodesManager *EtherFiNodesManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherFiNodesManager.Contract.Receive(&_EtherFiNodesManager.TransactOpts)
}

// EtherFiNodesManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAdminChangedIterator struct {
	Event *EtherFiNodesManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerAdminChanged)
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
		it.Event = new(EtherFiNodesManagerAdminChanged)
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
func (it *EtherFiNodesManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerAdminChanged represents a AdminChanged event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EtherFiNodesManagerAdminChangedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerAdminChangedIterator{contract: _EtherFiNodesManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerAdminChanged)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseAdminChanged(log types.Log) (*EtherFiNodesManagerAdminChanged, error) {
	event := new(EtherFiNodesManagerAdminChanged)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator is returned from FilterAllowedForwardedEigenpodCallsUpdated and is used to iterate over the raw logs and unpacked data for AllowedForwardedEigenpodCallsUpdated events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator struct {
	Event *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
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
		it.Event = new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
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
func (it *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated represents a AllowedForwardedEigenpodCallsUpdated event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated struct {
	Selector [4]byte
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedForwardedEigenpodCallsUpdated is a free log retrieval operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterAllowedForwardedEigenpodCallsUpdated(opts *bind.FilterOpts, selector [][4]byte) (*EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "AllowedForwardedEigenpodCallsUpdated", selectorRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdatedIterator{contract: _EtherFiNodesManager.contract, event: "AllowedForwardedEigenpodCallsUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedForwardedEigenpodCallsUpdated is a free log subscription operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchAllowedForwardedEigenpodCallsUpdated(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated, selector [][4]byte) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "AllowedForwardedEigenpodCallsUpdated", selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedEigenpodCallsUpdated", log); err != nil {
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

// ParseAllowedForwardedEigenpodCallsUpdated is a log parse operation binding the contract event 0x4ce3499ecaa9e9191427292e43115026001fdf36236d8527a962a1e67086021b.
//
// Solidity: event AllowedForwardedEigenpodCallsUpdated(bytes4 indexed selector, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseAllowedForwardedEigenpodCallsUpdated(log types.Log) (*EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated, error) {
	event := new(EtherFiNodesManagerAllowedForwardedEigenpodCallsUpdated)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedEigenpodCallsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator is returned from FilterAllowedForwardedExternalCallsUpdated and is used to iterate over the raw logs and unpacked data for AllowedForwardedExternalCallsUpdated events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator struct {
	Event *EtherFiNodesManagerAllowedForwardedExternalCallsUpdated // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
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
		it.Event = new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
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
func (it *EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerAllowedForwardedExternalCallsUpdated represents a AllowedForwardedExternalCallsUpdated event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerAllowedForwardedExternalCallsUpdated struct {
	Selector [4]byte
	Target   common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedForwardedExternalCallsUpdated is a free log retrieval operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterAllowedForwardedExternalCallsUpdated(opts *bind.FilterOpts, selector [][4]byte, _target []common.Address) (*EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "AllowedForwardedExternalCallsUpdated", selectorRule, _targetRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerAllowedForwardedExternalCallsUpdatedIterator{contract: _EtherFiNodesManager.contract, event: "AllowedForwardedExternalCallsUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedForwardedExternalCallsUpdated is a free log subscription operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchAllowedForwardedExternalCallsUpdated(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerAllowedForwardedExternalCallsUpdated, selector [][4]byte, _target []common.Address) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "AllowedForwardedExternalCallsUpdated", selectorRule, _targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedExternalCallsUpdated", log); err != nil {
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

// ParseAllowedForwardedExternalCallsUpdated is a log parse operation binding the contract event 0xf8fbe92f65ae1c742965709f0619c43fc13758a0d9c3b13b9afaecdfb7c50f48.
//
// Solidity: event AllowedForwardedExternalCallsUpdated(bytes4 indexed selector, address indexed _target, bool _allowed)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseAllowedForwardedExternalCallsUpdated(log types.Log) (*EtherFiNodesManagerAllowedForwardedExternalCallsUpdated, error) {
	event := new(EtherFiNodesManagerAllowedForwardedExternalCallsUpdated)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "AllowedForwardedExternalCallsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerBeaconUpgradedIterator struct {
	Event *EtherFiNodesManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerBeaconUpgraded)
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
		it.Event = new(EtherFiNodesManagerBeaconUpgraded)
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
func (it *EtherFiNodesManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerBeaconUpgraded represents a BeaconUpgraded event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EtherFiNodesManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerBeaconUpgradedIterator{contract: _EtherFiNodesManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerBeaconUpgraded)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseBeaconUpgraded(log types.Log) (*EtherFiNodesManagerBeaconUpgraded, error) {
	event := new(EtherFiNodesManagerBeaconUpgraded)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerFullWithdrawalIterator is returned from FilterFullWithdrawal and is used to iterate over the raw logs and unpacked data for FullWithdrawal events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerFullWithdrawalIterator struct {
	Event *EtherFiNodesManagerFullWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerFullWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerFullWithdrawal)
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
		it.Event = new(EtherFiNodesManagerFullWithdrawal)
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
func (it *EtherFiNodesManagerFullWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerFullWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerFullWithdrawal represents a FullWithdrawal event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerFullWithdrawal struct {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterFullWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherFiNodesManagerFullWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "FullWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerFullWithdrawalIterator{contract: _EtherFiNodesManager.contract, event: "FullWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFullWithdrawal is a free log subscription operation binding the contract event 0x23fd4a72178e02ea64b0e1b08ed6de9c7a7fb4bbb565b0917b52e0650a2d3a09.
//
// Solidity: event FullWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchFullWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerFullWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "FullWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerFullWithdrawal)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "FullWithdrawal", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseFullWithdrawal(log types.Log) (*EtherFiNodesManagerFullWithdrawal, error) {
	event := new(EtherFiNodesManagerFullWithdrawal)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "FullWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerInitializedIterator struct {
	Event *EtherFiNodesManagerInitialized // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerInitialized)
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
		it.Event = new(EtherFiNodesManagerInitialized)
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
func (it *EtherFiNodesManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerInitialized represents a Initialized event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EtherFiNodesManagerInitializedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerInitializedIterator{contract: _EtherFiNodesManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerInitialized)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseInitialized(log types.Log) (*EtherFiNodesManagerInitialized, error) {
	event := new(EtherFiNodesManagerInitialized)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerNodeExitProcessedIterator is returned from FilterNodeExitProcessed and is used to iterate over the raw logs and unpacked data for NodeExitProcessed events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerNodeExitProcessedIterator struct {
	Event *EtherFiNodesManagerNodeExitProcessed // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerNodeExitProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerNodeExitProcessed)
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
		it.Event = new(EtherFiNodesManagerNodeExitProcessed)
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
func (it *EtherFiNodesManagerNodeExitProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerNodeExitProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerNodeExitProcessed represents a NodeExitProcessed event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerNodeExitProcessed struct {
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeExitProcessed is a free log retrieval operation binding the contract event 0x0a9622219d3011f688c7de77a5e0f0f80a2ee1205429b3062b66827ee8c3b6b0.
//
// Solidity: event NodeExitProcessed(uint256 _validatorId)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterNodeExitProcessed(opts *bind.FilterOpts) (*EtherFiNodesManagerNodeExitProcessedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "NodeExitProcessed")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerNodeExitProcessedIterator{contract: _EtherFiNodesManager.contract, event: "NodeExitProcessed", logs: logs, sub: sub}, nil
}

// WatchNodeExitProcessed is a free log subscription operation binding the contract event 0x0a9622219d3011f688c7de77a5e0f0f80a2ee1205429b3062b66827ee8c3b6b0.
//
// Solidity: event NodeExitProcessed(uint256 _validatorId)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchNodeExitProcessed(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerNodeExitProcessed) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "NodeExitProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerNodeExitProcessed)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "NodeExitProcessed", log); err != nil {
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

// ParseNodeExitProcessed is a log parse operation binding the contract event 0x0a9622219d3011f688c7de77a5e0f0f80a2ee1205429b3062b66827ee8c3b6b0.
//
// Solidity: event NodeExitProcessed(uint256 _validatorId)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseNodeExitProcessed(log types.Log) (*EtherFiNodesManagerNodeExitProcessed, error) {
	event := new(EtherFiNodesManagerNodeExitProcessed)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "NodeExitProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerNodeExitRequestedIterator is returned from FilterNodeExitRequested and is used to iterate over the raw logs and unpacked data for NodeExitRequested events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerNodeExitRequestedIterator struct {
	Event *EtherFiNodesManagerNodeExitRequested // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerNodeExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerNodeExitRequested)
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
		it.Event = new(EtherFiNodesManagerNodeExitRequested)
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
func (it *EtherFiNodesManagerNodeExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerNodeExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerNodeExitRequested represents a NodeExitRequested event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerNodeExitRequested struct {
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeExitRequested is a free log retrieval operation binding the contract event 0x8f1aebefc80facd94136da81cfa288e9361156d61eddc7e0348391c7376c5c07.
//
// Solidity: event NodeExitRequested(uint256 _validatorId)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterNodeExitRequested(opts *bind.FilterOpts) (*EtherFiNodesManagerNodeExitRequestedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "NodeExitRequested")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerNodeExitRequestedIterator{contract: _EtherFiNodesManager.contract, event: "NodeExitRequested", logs: logs, sub: sub}, nil
}

// WatchNodeExitRequested is a free log subscription operation binding the contract event 0x8f1aebefc80facd94136da81cfa288e9361156d61eddc7e0348391c7376c5c07.
//
// Solidity: event NodeExitRequested(uint256 _validatorId)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchNodeExitRequested(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerNodeExitRequested) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "NodeExitRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerNodeExitRequested)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "NodeExitRequested", log); err != nil {
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

// ParseNodeExitRequested is a log parse operation binding the contract event 0x8f1aebefc80facd94136da81cfa288e9361156d61eddc7e0348391c7376c5c07.
//
// Solidity: event NodeExitRequested(uint256 _validatorId)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseNodeExitRequested(log types.Log) (*EtherFiNodesManagerNodeExitRequested, error) {
	event := new(EtherFiNodesManagerNodeExitRequested)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "NodeExitRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerOwnershipTransferredIterator struct {
	Event *EtherFiNodesManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerOwnershipTransferred)
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
		it.Event = new(EtherFiNodesManagerOwnershipTransferred)
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
func (it *EtherFiNodesManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EtherFiNodesManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerOwnershipTransferredIterator{contract: _EtherFiNodesManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerOwnershipTransferred)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EtherFiNodesManagerOwnershipTransferred, error) {
	event := new(EtherFiNodesManagerOwnershipTransferred)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerPartialWithdrawalIterator is returned from FilterPartialWithdrawal and is used to iterate over the raw logs and unpacked data for PartialWithdrawal events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPartialWithdrawalIterator struct {
	Event *EtherFiNodesManagerPartialWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerPartialWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerPartialWithdrawal)
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
		it.Event = new(EtherFiNodesManagerPartialWithdrawal)
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
func (it *EtherFiNodesManagerPartialWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerPartialWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerPartialWithdrawal represents a PartialWithdrawal event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPartialWithdrawal struct {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterPartialWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherFiNodesManagerPartialWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "PartialWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerPartialWithdrawalIterator{contract: _EtherFiNodesManager.contract, event: "PartialWithdrawal", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawal is a free log subscription operation binding the contract event 0x0c9b1112957fe7d0e2f96690e65a9122e07ca9cd19a2f99966b29b5991c3be84.
//
// Solidity: event PartialWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, uint256 toOperator, uint256 toTnft, uint256 toBnft, uint256 toTreasury)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchPartialWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerPartialWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "PartialWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerPartialWithdrawal)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "PartialWithdrawal", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParsePartialWithdrawal(log types.Log) (*EtherFiNodesManagerPartialWithdrawal, error) {
	event := new(EtherFiNodesManagerPartialWithdrawal)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "PartialWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPausedIterator struct {
	Event *EtherFiNodesManagerPaused // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerPaused)
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
		it.Event = new(EtherFiNodesManagerPaused)
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
func (it *EtherFiNodesManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerPaused represents a Paused event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EtherFiNodesManagerPausedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerPausedIterator{contract: _EtherFiNodesManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerPaused)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParsePaused(log types.Log) (*EtherFiNodesManagerPaused, error) {
	event := new(EtherFiNodesManagerPaused)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerPhaseChangedIterator is returned from FilterPhaseChanged and is used to iterate over the raw logs and unpacked data for PhaseChanged events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPhaseChangedIterator struct {
	Event *EtherFiNodesManagerPhaseChanged // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerPhaseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerPhaseChanged)
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
		it.Event = new(EtherFiNodesManagerPhaseChanged)
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
func (it *EtherFiNodesManagerPhaseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerPhaseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerPhaseChanged represents a PhaseChanged event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerPhaseChanged struct {
	ValidatorId *big.Int
	Phase       uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPhaseChanged is a free log retrieval operation binding the contract event 0x70eca82567b065893a5e6cc590178b6b320855676b6a9a066625933e0c8ebe58.
//
// Solidity: event PhaseChanged(uint256 indexed _validatorId, uint8 _phase)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterPhaseChanged(opts *bind.FilterOpts, _validatorId []*big.Int) (*EtherFiNodesManagerPhaseChangedIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "PhaseChanged", _validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerPhaseChangedIterator{contract: _EtherFiNodesManager.contract, event: "PhaseChanged", logs: logs, sub: sub}, nil
}

// WatchPhaseChanged is a free log subscription operation binding the contract event 0x70eca82567b065893a5e6cc590178b6b320855676b6a9a066625933e0c8ebe58.
//
// Solidity: event PhaseChanged(uint256 indexed _validatorId, uint8 _phase)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchPhaseChanged(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerPhaseChanged, _validatorId []*big.Int) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "PhaseChanged", _validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerPhaseChanged)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "PhaseChanged", log); err != nil {
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

// ParsePhaseChanged is a log parse operation binding the contract event 0x70eca82567b065893a5e6cc590178b6b320855676b6a9a066625933e0c8ebe58.
//
// Solidity: event PhaseChanged(uint256 indexed _validatorId, uint8 _phase)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParsePhaseChanged(log types.Log) (*EtherFiNodesManagerPhaseChanged, error) {
	event := new(EtherFiNodesManagerPhaseChanged)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "PhaseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerQueuedRestakingWithdrawalIterator is returned from FilterQueuedRestakingWithdrawal and is used to iterate over the raw logs and unpacked data for QueuedRestakingWithdrawal events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerQueuedRestakingWithdrawalIterator struct {
	Event *EtherFiNodesManagerQueuedRestakingWithdrawal // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerQueuedRestakingWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerQueuedRestakingWithdrawal)
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
		it.Event = new(EtherFiNodesManagerQueuedRestakingWithdrawal)
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
func (it *EtherFiNodesManagerQueuedRestakingWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerQueuedRestakingWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerQueuedRestakingWithdrawal represents a QueuedRestakingWithdrawal event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerQueuedRestakingWithdrawal struct {
	ValidatorId     *big.Int
	EtherFiNode     common.Address
	WithdrawalRoots [][32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterQueuedRestakingWithdrawal is a free log retrieval operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterQueuedRestakingWithdrawal(opts *bind.FilterOpts, _validatorId []*big.Int, etherFiNode []common.Address) (*EtherFiNodesManagerQueuedRestakingWithdrawalIterator, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "QueuedRestakingWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerQueuedRestakingWithdrawalIterator{contract: _EtherFiNodesManager.contract, event: "QueuedRestakingWithdrawal", logs: logs, sub: sub}, nil
}

// WatchQueuedRestakingWithdrawal is a free log subscription operation binding the contract event 0x8c9fe3546da789766f4f5cd07e17b8e68c0e46e494b3a60a798a8f493283263a.
//
// Solidity: event QueuedRestakingWithdrawal(uint256 indexed _validatorId, address indexed etherFiNode, bytes32[] withdrawalRoots)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchQueuedRestakingWithdrawal(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerQueuedRestakingWithdrawal, _validatorId []*big.Int, etherFiNode []common.Address) (event.Subscription, error) {

	var _validatorIdRule []interface{}
	for _, _validatorIdItem := range _validatorId {
		_validatorIdRule = append(_validatorIdRule, _validatorIdItem)
	}
	var etherFiNodeRule []interface{}
	for _, etherFiNodeItem := range etherFiNode {
		etherFiNodeRule = append(etherFiNodeRule, etherFiNodeItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "QueuedRestakingWithdrawal", _validatorIdRule, etherFiNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerQueuedRestakingWithdrawal)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "QueuedRestakingWithdrawal", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseQueuedRestakingWithdrawal(log types.Log) (*EtherFiNodesManagerQueuedRestakingWithdrawal, error) {
	event := new(EtherFiNodesManagerQueuedRestakingWithdrawal)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "QueuedRestakingWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUnpausedIterator struct {
	Event *EtherFiNodesManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerUnpaused)
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
		it.Event = new(EtherFiNodesManagerUnpaused)
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
func (it *EtherFiNodesManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerUnpaused represents a Unpaused event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EtherFiNodesManagerUnpausedIterator, error) {

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerUnpausedIterator{contract: _EtherFiNodesManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerUnpaused)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseUnpaused(log types.Log) (*EtherFiNodesManagerUnpaused, error) {
	event := new(EtherFiNodesManagerUnpaused)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherFiNodesManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUpgradedIterator struct {
	Event *EtherFiNodesManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherFiNodesManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherFiNodesManagerUpgraded)
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
		it.Event = new(EtherFiNodesManagerUpgraded)
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
func (it *EtherFiNodesManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherFiNodesManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherFiNodesManagerUpgraded represents a Upgraded event raised by the EtherFiNodesManager contract.
type EtherFiNodesManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EtherFiNodesManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EtherFiNodesManagerUpgradedIterator{contract: _EtherFiNodesManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EtherFiNodesManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherFiNodesManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherFiNodesManagerUpgraded)
				if err := _EtherFiNodesManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherFiNodesManager *EtherFiNodesManagerFilterer) ParseUpgraded(log types.Log) (*EtherFiNodesManagerUpgraded, error) {
	event := new(EtherFiNodesManagerUpgraded)
	if err := _EtherFiNodesManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
