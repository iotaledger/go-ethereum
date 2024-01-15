package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type ISCMagicContract interface {
	Run(evm *EVM, caller common.Address, input []byte, value *uint256.Int, gas uint64, readOnly bool) ([]byte, uint64, error)
}
