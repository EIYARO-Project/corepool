package ey

import (
	"math/big"
	"time"

	"corepool/core/protocol/bc"

	"corepool/common/mining/utils"
	ss "corepool/stratum"
	"corepool/stratum/ey/util"
)

type eyJob struct {
	id                     ss.JobId
	version                uint64
	height                 uint64
	previousBlockHash      *bc.Hash
	timestamp              time.Time
	transactionsMerkleRoot *bc.Hash
	transactionStatusHash  *bc.Hash
	bits                   uint64
	seed                   *bc.Hash
	nonce                  uint64
	diff                   *big.Int
}

func (j *eyJob) GetId() ss.JobId {
	return j.id
}

func (j *eyJob) GetDiff() uint64 {
	return j.diff.Uint64()
}

func (j *eyJob) GetTarget() (string, bool, bool) {
	return "", false, false
}

func (j *eyJob) Encode() (interface{}, error) {
	return ss.StratumJSONRpcNotify{
		Version: "2.0",
		Method:  "job",
		Params:  j.genReplyData(),
	}, nil
}

func (j *eyJob) encodeLogin(login string) *jobReply {
	return &jobReply{
		Id:     login,
		Job:    j.genReplyData(),
		Status: "OK",
	}
}

func (j *eyJob) genReplyData() *jobReplyData {
	return &jobReplyData{
		JobId:                  j.GetId().String(),
		Version:                utils.ToLittleEndianHex(j.version),
		Height:                 utils.ToLittleEndianHex(j.height),
		PreviousBlockHash:      j.previousBlockHash.String(),
		Timestamp:              utils.ToLittleEndianHex(uint64(j.timestamp.Unix())),
		TransactionsMerkleRoot: j.transactionsMerkleRoot.String(),
		TransactionStatusHash:  j.transactionStatusHash.String(),
		Nonce:                  utils.ToLittleEndianHex(uint64(j.nonce)),
		Bits:                   utils.ToLittleEndianHex(j.bits),
		Seed:                   j.seed.String(),
		Target:                 util.GetTargetHex(j.diff),
	}
}
