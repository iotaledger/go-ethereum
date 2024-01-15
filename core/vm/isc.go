package vm

import "github.com/holiman/uint256"

type ISCMagicContract interface {
	Run(evm *EVM, caller ContractRef, input []byte, value *uint256.Int, gas uint64, readOnly bool) ([]byte, uint64, error)
}
