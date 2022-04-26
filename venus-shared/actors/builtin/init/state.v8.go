// FETCHED FROM LOTUS: builtin/init/state.go.template

package init

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/venus/venus-shared/actors/adt"

	builtin8 "github.com/filecoin-project/specs-actors/v8/actors/builtin"

	init8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/init"
	adt8 "github.com/filecoin-project/specs-actors/v8/actors/util/adt"
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

func make8(store adt.Store, networkName string) (State, error) {
	out := state8{store: store}

	s, err := init8.ConstructState(store, networkName)
	if err != nil {
		return nil, err
	}

	out.State = *s

	return &out, nil
}

type state8 struct {
	init8.State
	store adt.Store
}

func (s *state8) ResolveAddress(address address.Address) (address.Address, bool, error) {
	return s.State.ResolveAddress(s.store, address)
}

func (s *state8) MapAddressToNewID(address address.Address) (address.Address, error) {
	return s.State.MapAddressToNewID(s.store, address)
}

func (s *state8) ForEachActor(cb func(id abi.ActorID, address address.Address) error) error {
	addrs, err := adt8.AsMap(s.store, s.State.AddressMap, builtin8.DefaultHamtBitwidth)
	if err != nil {
		return err
	}
	var actorID cbg.CborInt
	return addrs.ForEach(&actorID, func(key string) error {
		addr, err := address.NewFromBytes([]byte(key))
		if err != nil {
			return err
		}
		return cb(abi.ActorID(actorID), addr)
	})
}

func (s *state8) NetworkName() (string, error) {
	return string(s.State.NetworkName), nil
}

func (s *state8) SetNetworkName(name string) error {
	s.State.NetworkName = name
	return nil
}

func (s *state8) SetNextID(id abi.ActorID) error {
	s.State.NextID = id
	return nil
}

func (s *state8) Remove(addrs ...address.Address) (err error) {
	m, err := adt8.AsMap(s.store, s.State.AddressMap, builtin8.DefaultHamtBitwidth)
	if err != nil {
		return err
	}
	for _, addr := range addrs {
		if err = m.Delete(abi.AddrKey(addr)); err != nil {
			return xerrors.Errorf("failed to delete entry for address: %s; err: %w", addr, err)
		}
	}
	amr, err := m.Root()
	if err != nil {
		return xerrors.Errorf("failed to get address map root: %w", err)
	}
	s.State.AddressMap = amr
	return nil
}

func (s *state8) SetAddressMap(mcid cid.Cid) error {
	s.State.AddressMap = mcid
	return nil
}

func (s *state8) AddressMap() (adt.Map, error) {
	return adt8.AsMap(s.store, s.State.AddressMap, builtin8.DefaultHamtBitwidth)
}

func (s *state8) GetState() interface{} {
	return &s.State
}
