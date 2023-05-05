package simulation

import (
	// NewOperationMsg (in this package) requires all the types to be imported in order to dynamically create a protoreflect
	// capable proto message for aminojson.Encoder to marshal into amino json.  Since we need a conrete instance of the
	// type/proto message, and simulations (theoretically) can be run against any module, we need to import all the types.
	_ "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	_ "cosmossdk.io/api/cosmos/app/v1alpha1"
	_ "cosmossdk.io/api/cosmos/auth/module/v1"
	_ "cosmossdk.io/api/cosmos/auth/v1beta1"
	_ "cosmossdk.io/api/cosmos/authz/module/v1"
	_ "cosmossdk.io/api/cosmos/authz/v1beta1"
	_ "cosmossdk.io/api/cosmos/autocli/v1"
	_ "cosmossdk.io/api/cosmos/bank/module/v1"
	_ "cosmossdk.io/api/cosmos/bank/v1beta1"
	_ "cosmossdk.io/api/cosmos/base/abci/v1beta1"
	_ "cosmossdk.io/api/cosmos/base/node/v1beta1"
	_ "cosmossdk.io/api/cosmos/base/query/v1beta1"
	_ "cosmossdk.io/api/cosmos/base/reflection/v1beta1"
	_ "cosmossdk.io/api/cosmos/base/reflection/v2alpha1"
	_ "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	_ "cosmossdk.io/api/cosmos/base/v1beta1"
	_ "cosmossdk.io/api/cosmos/consensus/module/v1"
	_ "cosmossdk.io/api/cosmos/consensus/v1"
	_ "cosmossdk.io/api/cosmos/crisis/module/v1"
	_ "cosmossdk.io/api/cosmos/crisis/v1beta1"
	_ "cosmossdk.io/api/cosmos/crypto/ed25519"
	_ "cosmossdk.io/api/cosmos/crypto/hd/v1"
	_ "cosmossdk.io/api/cosmos/crypto/multisig"
	_ "cosmossdk.io/api/cosmos/crypto/multisig/v1beta1"
	_ "cosmossdk.io/api/cosmos/crypto/secp256k1"
	_ "cosmossdk.io/api/cosmos/crypto/secp256r1"
	_ "cosmossdk.io/api/cosmos/distribution/module/v1"
	_ "cosmossdk.io/api/cosmos/distribution/v1beta1"
	_ "cosmossdk.io/api/cosmos/evidence/module/v1"
	_ "cosmossdk.io/api/cosmos/evidence/v1beta1"
	_ "cosmossdk.io/api/cosmos/feegrant/module/v1"
	_ "cosmossdk.io/api/cosmos/feegrant/v1beta1"
	_ "cosmossdk.io/api/cosmos/genutil/module/v1"
	_ "cosmossdk.io/api/cosmos/genutil/v1beta1"
	_ "cosmossdk.io/api/cosmos/gov/module/v1"
	_ "cosmossdk.io/api/cosmos/gov/v1"
	_ "cosmossdk.io/api/cosmos/gov/v1beta1"
	_ "cosmossdk.io/api/cosmos/group/module/v1"
	_ "cosmossdk.io/api/cosmos/group/v1"
	_ "cosmossdk.io/api/cosmos/mint/module/v1"
	_ "cosmossdk.io/api/cosmos/mint/v1beta1"
	_ "cosmossdk.io/api/cosmos/msg/textual/v1"
	_ "cosmossdk.io/api/cosmos/msg/v1"
	_ "cosmossdk.io/api/cosmos/nft/module/v1"
	_ "cosmossdk.io/api/cosmos/nft/v1beta1"
	_ "cosmossdk.io/api/cosmos/orm/module/v1alpha1"
	_ "cosmossdk.io/api/cosmos/orm/query/v1alpha1"
	_ "cosmossdk.io/api/cosmos/orm/v1"
	_ "cosmossdk.io/api/cosmos/orm/v1alpha1"
	_ "cosmossdk.io/api/cosmos/params/module/v1"
	_ "cosmossdk.io/api/cosmos/params/v1beta1"
	_ "cosmossdk.io/api/cosmos/query/v1"
	_ "cosmossdk.io/api/cosmos/reflection/v1"
	_ "cosmossdk.io/api/cosmos/slashing/module/v1"
	_ "cosmossdk.io/api/cosmos/slashing/v1beta1"
	_ "cosmossdk.io/api/cosmos/staking/module/v1"
	_ "cosmossdk.io/api/cosmos/staking/v1beta1"
	_ "cosmossdk.io/api/cosmos/store/snapshots/v1"
	_ "cosmossdk.io/api/cosmos/store/streaming/abci"
	_ "cosmossdk.io/api/cosmos/store/v1beta1"
	_ "cosmossdk.io/api/cosmos/tx/config/v1"
	_ "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	_ "cosmossdk.io/api/cosmos/tx/v1beta1"
	_ "cosmossdk.io/api/cosmos/upgrade/module/v1"
	_ "cosmossdk.io/api/cosmos/upgrade/v1beta1"
	_ "cosmossdk.io/api/cosmos/vesting/module/v1"
	_ "cosmossdk.io/api/cosmos/vesting/v1beta1"
)
