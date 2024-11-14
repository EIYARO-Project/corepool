package ey

import (
	"math"

	"github.com/cornelk/hashmap"

	ss "corepool/stratum"
)

const serverIdOffset = uint(60)

type eySessionData struct {
	nonce    uint64
	worker   *ss.Worker
	submitId *hashmap.HashMap
}

func (s *eySessionData) GetWorker() *ss.Worker {
	return s.worker
}

func (s *eySessionData) SetWorker(worker *ss.Worker) {
	s.worker = worker
}

func (s *eySessionData) getNonce() uint64 {
	return s.nonce
}

type eySessionDataBuilder struct {
	id              uint64
	maxSessions     int
	sessionIdOffset uint
}

func NewBtmcSessionDataBuilder(serverId uint64, maxSessions int) *eySessionDataBuilder {
	sessionIdOffset := serverIdOffset - uint(math.Ceil(math.Log2(float64(maxSessions))))
	if sessionIdOffset == serverIdOffset {
		sessionIdOffset--
	}
	return &eySessionDataBuilder{
		id:              serverId,
		maxSessions:     maxSessions,
		sessionIdOffset: sessionIdOffset,
	}
}

// Build builds a btmSession
func (b *eySessionDataBuilder) Build(sessionId uint) ss.SessionData {
	return &eySessionData{
		nonce:    (b.id << serverIdOffset) | (uint64(sessionId) << b.sessionIdOffset),
		worker:   nil,
		submitId: &hashmap.HashMap{},
	}
}
