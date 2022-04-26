// FETCHED FROM LOTUS: builtin/reward/state.go.template

package reward

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/venus/venus-shared/actors/adt"
	"github.com/filecoin-project/venus/venus-shared/actors/builtin"

	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"
	reward8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/reward"
	smoothing8 "github.com/filecoin-project/specs-actors/v8/actors/util/smoothing"
)

var _ State = (*state8)(nil)

func load8(store adt.Store, root cid.Cid) (State, error) {
	out := state8{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make8(store adt.Store, currRealizedPower abi.StoragePower) (State, error) {
	out := state8{store: store}
	out.State = *reward8.ConstructState(currRealizedPower)
	return &out, nil
}

type state8 struct {
	reward8.State
	store adt.Store
}

func (s *state8) ThisEpochReward() (abi.TokenAmount, error) {
	return s.State.ThisEpochReward, nil
}

func (s *state8) ThisEpochRewardSmoothed() (builtin.FilterEstimate, error) {

	return builtin.FilterEstimate{
		PositionEstimate: s.State.ThisEpochRewardSmoothed.PositionEstimate,
		VelocityEstimate: s.State.ThisEpochRewardSmoothed.VelocityEstimate,
	}, nil

}

func (s *state8) ThisEpochBaselinePower() (abi.StoragePower, error) {
	return s.State.ThisEpochBaselinePower, nil
}

func (s *state8) TotalStoragePowerReward() (abi.TokenAmount, error) {
	return s.State.TotalStoragePowerReward, nil
}

func (s *state8) EffectiveBaselinePower() (abi.StoragePower, error) {
	return s.State.EffectiveBaselinePower, nil
}

func (s *state8) EffectiveNetworkTime() (abi.ChainEpoch, error) {
	return s.State.EffectiveNetworkTime, nil
}

func (s *state8) CumsumBaseline() (reward8.Spacetime, error) {
	return s.State.CumsumBaseline, nil
}

func (s *state8) CumsumRealized() (reward8.Spacetime, error) {
	return s.State.CumsumRealized, nil
}

func (s *state8) InitialPledgeForPower(qaPower abi.StoragePower, networkTotalPledge abi.TokenAmount, networkQAPower *builtin.FilterEstimate, circSupply abi.TokenAmount) (abi.TokenAmount, error) {
	return miner8.InitialPledgeForPower(
		qaPower,
		s.State.ThisEpochBaselinePower,
		s.State.ThisEpochRewardSmoothed,
		smoothing8.FilterEstimate{
			PositionEstimate: networkQAPower.PositionEstimate,
			VelocityEstimate: networkQAPower.VelocityEstimate,
		},
		circSupply,
	), nil
}

func (s *state8) PreCommitDepositForPower(networkQAPower builtin.FilterEstimate, sectorWeight abi.StoragePower) (abi.TokenAmount, error) {
	return miner8.PreCommitDepositForPower(s.State.ThisEpochRewardSmoothed,
		smoothing8.FilterEstimate{
			PositionEstimate: networkQAPower.PositionEstimate,
			VelocityEstimate: networkQAPower.VelocityEstimate,
		},
		sectorWeight), nil
}

func (s *state8) GetState() interface{} {
	return &s.State
}
