package ey

import (
	"errors"
	"math/big"

	"corepool/core/protocol/bc"
	"corepool/core/protocol/bc/types"

	"corepool/common/logger"
	ss "corepool/stratum"
)

const (
	verMagicNum   = uint64(1)
	defaultReward = uint64(41250000000)
	defaultFee    = uint64(300)
)

type eyShare struct {
	job    *eyJob
	worker *ss.Worker

	nonce     uint64
	result    string
	header    *types.BlockHeader
	blockHash *bc.Hash
	netDiff   *big.Int

	state  ss.ShareState
	reason ss.RejectReason
}

// build block from the share for node submission
func (s *eyShare) BuildBlock() (ss.BlockTemplate, error) {
	// not implemented
	logger.Fatal("BuildBlock not implemented")
	return nil, nil
}

// build pb sharelog from the share for logging
func (s *eyShare) BuildLog(port uint64) ([]byte, error) {
	return nil, errors.New("not support")
}

// update share state
func (s *eyShare) UpdateState(state ss.ShareState, reason ss.RejectReason) error {
	s.state = state
	s.reason = reason
	return nil
}

func (s *eyShare) GetState() ss.ShareState {
	return s.state
}

func (s *eyShare) GetReason() ss.RejectReason {
	return s.reason
}

func (s *eyShare) GetWorker() *ss.Worker {
	return s.worker
}
