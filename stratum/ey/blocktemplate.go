package ey

import (
	"time"

	"corepool/core/protocol/bc"

	"corepool/common/logger"
	ss "corepool/stratum"
)

type eyBlockTemplate struct {
	version                uint64
	height                 uint64
	previousBlockHash      *bc.Hash
	timestamp              time.Time
	transactionsMerkleRoot *bc.Hash
	transactionStatusHash  *bc.Hash
	nonce                  uint64
	bits                   uint64
	seed                   *bc.Hash
}

func (b *eyBlockTemplate) CreateJob(session *ss.TcpSession) (ss.Job, error) {
	data := session.GetSessionData().(*eySessionData)
	job := &eyJob{
		id:                     ss.AllocJobId(),
		version:                b.version,
		height:                 b.height,
		previousBlockHash:      b.previousBlockHash,
		timestamp:              b.timestamp,
		transactionsMerkleRoot: b.transactionsMerkleRoot,
		transactionStatusHash:  b.transactionStatusHash,
		bits:                   b.bits,
		seed:                   b.seed,
		nonce:                  data.getNonce(),
		diff:                   session.GetDiff(),
	}
	logger.Info("generate new job",
		"session_id", session.GetId(),
		"session_ip", session.GetIp(),
		"job_id", job.GetId(),
		"job_diff", job.diff,
		"height", b.height)
	return job, nil
}

// compare with another block template
// 1 : newer than the other
// 0 : same as the other
// -1 : older than the other, update
func (b *eyBlockTemplate) Compare(template ss.BlockTemplate) int {
	// TODO: compare when height info is available
	eyBT := template.(*eyBlockTemplate)
	if b != nil && b.previousBlockHash.String() == eyBT.previousBlockHash.String() {
		return 0
	}
	if eyBT.height <= b.height {
		logger.Warn("bt diff height",
			"old_height", b.height,
			"new_height", eyBT.height,
			"old_prevhash", b.previousBlockHash,
			"new_prevhash", eyBT.previousBlockHash,
		)
	}
	// update block template when previous block hash not the same, no matter newer or older
	return -1
}
