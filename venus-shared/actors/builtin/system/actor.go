// FETCHED FROM LOTUS: builtin/system/actor.go.template

package system

import (
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/venus/venus-shared/actors"
	"github.com/filecoin-project/venus/venus-shared/actors/adt"
	"github.com/filecoin-project/venus/venus-shared/actors/types"
	"github.com/ipfs/go-cid"

	"fmt"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"

	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"

	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	builtin15 "github.com/filecoin-project/go-state-types/builtin"
)

var (
	Address = builtin15.SystemActorAddr
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := actors.GetActorMetaByCode(act.Code); ok {
		if name != manifest.SystemKey {
			return nil, fmt.Errorf("actor code is not system: %s", name)
		}

		switch av {

		case actorstypes.Version8:
			return load8(store, act.Head)

		case actorstypes.Version9:
			return load9(store, act.Head)

		case actorstypes.Version10:
			return load10(store, act.Head)

		case actorstypes.Version11:
			return load11(store, act.Head)

		case actorstypes.Version12:
			return load12(store, act.Head)

		case actorstypes.Version13:
			return load13(store, act.Head)

		case actorstypes.Version14:
			return load14(store, act.Head)

		case actorstypes.Version15:
			return load15(store, act.Head)

		}
	}

	switch act.Code {

	case builtin0.SystemActorCodeID:
		return load0(store, act.Head)

	case builtin2.SystemActorCodeID:
		return load2(store, act.Head)

	case builtin3.SystemActorCodeID:
		return load3(store, act.Head)

	case builtin4.SystemActorCodeID:
		return load4(store, act.Head)

	case builtin5.SystemActorCodeID:
		return load5(store, act.Head)

	case builtin6.SystemActorCodeID:
		return load6(store, act.Head)

	case builtin7.SystemActorCodeID:
		return load7(store, act.Head)

	}

	return nil, fmt.Errorf("unknown actor code %s", act.Code)
}

func MakeState(store adt.Store, av actorstypes.Version, builtinActors cid.Cid) (State, error) {
	switch av {

	case actorstypes.Version0:
		return make0(store)

	case actorstypes.Version2:
		return make2(store)

	case actorstypes.Version3:
		return make3(store)

	case actorstypes.Version4:
		return make4(store)

	case actorstypes.Version5:
		return make5(store)

	case actorstypes.Version6:
		return make6(store)

	case actorstypes.Version7:
		return make7(store)

	case actorstypes.Version8:
		return make8(store, builtinActors)

	case actorstypes.Version9:
		return make9(store, builtinActors)

	case actorstypes.Version10:
		return make10(store, builtinActors)

	case actorstypes.Version11:
		return make11(store, builtinActors)

	case actorstypes.Version12:
		return make12(store, builtinActors)

	case actorstypes.Version13:
		return make13(store, builtinActors)

	case actorstypes.Version14:
		return make14(store, builtinActors)

	case actorstypes.Version15:
		return make15(store, builtinActors)

	}
	return nil, fmt.Errorf("unknown actor version %d", av)
}

type State interface {
	Code() cid.Cid
	ActorKey() string
	ActorVersion() actorstypes.Version

	GetState() interface{}
	GetBuiltinActors() cid.Cid
	SetBuiltinActors(cid.Cid) error
}

func AllCodes() []cid.Cid {
	return []cid.Cid{
		(&state0{}).Code(),
		(&state2{}).Code(),
		(&state3{}).Code(),
		(&state4{}).Code(),
		(&state5{}).Code(),
		(&state6{}).Code(),
		(&state7{}).Code(),
		(&state8{}).Code(),
		(&state9{}).Code(),
		(&state10{}).Code(),
		(&state11{}).Code(),
		(&state12{}).Code(),
		(&state13{}).Code(),
		(&state14{}).Code(),
		(&state15{}).Code(),
	}
}
