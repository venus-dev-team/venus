// FETCHED FROM LOTUS: builtin/multisig/message.go.template

package multisig

import (
	"math"

	"golang.org/x/crypto/blake2b"

	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	init0 "github.com/filecoin-project/specs-actors/actors/builtin/init"
	multisig0 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"

	builtintypes "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/venus/venus-shared/actors"
	init_ "github.com/filecoin-project/venus/venus-shared/actors/builtin/init"
	"github.com/filecoin-project/venus/venus-shared/actors/types"
)

type message0 struct{ from address.Address }

func (m message0) Create(
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {

	lenAddrs := uint64(len(signers))

	if lenAddrs < threshold {
		return nil, fmt.Errorf("cannot require signing of more addresses than provided for multisig")
	}

	if threshold == 0 {
		threshold = lenAddrs
	}

	if m.from == address.Undef {
		return nil, fmt.Errorf("must provide source address")
	}

	if unlockStart != 0 {
		return nil, fmt.Errorf("actors v0 does not support a non-zero vesting start time")
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig0.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init0.ExecParams{
		CodeCID:           builtin0.MultisigActorCodeID,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Method: builtintypes.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}

func (m message0) Propose(msig, to address.Address, amt abi.TokenAmount,
	method abi.MethodNum, params []byte) (*types.Message, error) {

	if msig == address.Undef {
		return nil, fmt.Errorf("must provide a multisig address for proposal")
	}

	if to == address.Undef {
		return nil, fmt.Errorf("must provide a target address for proposal")
	}

	if amt.Sign() == -1 {
		return nil, fmt.Errorf("must provide a non-negative amount for proposed send")
	}

	if m.from == address.Undef {
		return nil, fmt.Errorf("must provide source address")
	}

	enc, actErr := actors.SerializeParams(&multisig0.ProposeParams{
		To:     to,
		Value:  amt,
		Method: method,
		Params: params,
	})
	if actErr != nil {
		return nil, fmt.Errorf("failed to serialize parameters: %w", actErr)
	}

	return &types.Message{
		To:     msig,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin0.MethodsMultisig.Propose,
		Params: enc,
	}, nil
}

func (m message0) Approve(msig address.Address, txID uint64, hashData *ProposalHashData) (*types.Message, error) {
	enc, err := txnParams(txID, hashData)
	if err != nil {
		return nil, err
	}

	return &types.Message{
		To:     msig,
		From:   m.from,
		Value:  types.NewInt(0),
		Method: builtin0.MethodsMultisig.Approve,
		Params: enc,
	}, nil
}

func (m message0) Cancel(msig address.Address, txID uint64, hashData *ProposalHashData) (*types.Message, error) {
	enc, err := txnParams(txID, hashData)
	if err != nil {
		return nil, err
	}

	return &types.Message{
		To:     msig,
		From:   m.from,
		Value:  types.NewInt(0),
		Method: builtin0.MethodsMultisig.Cancel,
		Params: enc,
	}, nil
}

func txnParams(id uint64, data *ProposalHashData) ([]byte, error) {
	if id > uint64(math.MaxInt64) {
		return nil, fmt.Errorf("transaction ID %d is out of range for int64", id)
	}
	params := multisig0.TxnIDParams{ID: multisig0.TxnID(id)}
	if data != nil {
		if data.Requester.Protocol() != address.ID {
			return nil, fmt.Errorf("proposer address must be an ID address, was %s", data.Requester)
		}
		if data.Value.Sign() == -1 {
			return nil, fmt.Errorf("proposal value must be non-negative, was %s", data.Value)
		}
		if data.To == address.Undef {
			return nil, fmt.Errorf("proposed destination address must be set")
		}
		pser, err := data.Serialize()
		if err != nil {
			return nil, err
		}
		hash := blake2b.Sum256(pser)
		params.ProposalHash = hash[:]
	}

	return actors.SerializeParams(&params)
}
