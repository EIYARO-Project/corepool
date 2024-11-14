package ey

import (
	"math/big"

	"corepool/core/consensus/difficulty"
	algorithm "corepool/core/mining/tensority/go_algorithm"
	"corepool/core/protocol/bc/types"

	ss "corepool/stratum"
)

type eyVerifier struct {
	serverState *ss.ServerState
}

func NewBtmcVerifier(state *ss.ServerState) (*eyVerifier, error) {
	return &eyVerifier{
		serverState: state,
	}, nil
}

func (v *eyVerifier) Verify(share ss.Share) error {
	eyShare := share.(*eyShare)
	eyJob := eyShare.job
	eyShare.header = &types.BlockHeader{
		Version:           eyJob.version,
		Height:            eyJob.height,
		PreviousBlockHash: *eyJob.previousBlockHash,
		Timestamp:         uint64(eyJob.timestamp.Unix()),
		BlockCommitment: types.BlockCommitment{
			TransactionsMerkleRoot: *eyJob.transactionsMerkleRoot,
			TransactionStatusHash:  *eyJob.transactionStatusHash,
		},
		Nonce: eyShare.nonce,
		Bits:  eyJob.bits,
	}
	shareHeader := eyShare.header
	headerHash := shareHeader.Hash()
	cmpHash := algorithm.LegacyAlgorithm(&headerHash, eyJob.seed)
	if cmpHash == nil {
		share.UpdateState(ss.ShareStateRejected, ss.RejectReasonUndefined)
		return nil
	}

	eyShare.blockHash = &headerHash
	bMax := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil)
	bits := difficulty.BigToCompact(big.NewInt(0).Div(bMax, eyShare.netDiff))
	if difficulty.HashToBig(cmpHash).Cmp(difficulty.CompactToBig(bits)) <= 0 {
		share.UpdateState(ss.ShareStateBlock, ss.RejectReasonPass)
		return nil
	}

	shareBits := difficulty.BigToCompact(big.NewInt(0).Div(bMax, eyJob.diff))
	if difficulty.HashToBig(cmpHash).Cmp(difficulty.CompactToBig(shareBits)) > 0 {
		share.UpdateState(ss.ShareStateRejected, ss.RejectReasonLowDiff)
		return nil
	}
	share.UpdateState(ss.ShareStateAccepted, ss.RejectReasonPass)
	return nil
}
