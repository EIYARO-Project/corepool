package ey

import (
	"errors"
	"sync"

	"github.com/segmentio/encoding/json"

	"corepool/core/api"
	"corepool/common/logger"
	ss "corepool/stratum"
	"corepool/stratum/ey/rpc"
)

type eyNodeSyncer struct {
	client *rpc.BtmcClient
	bt     *api.GetWorkResp
	btLock sync.RWMutex

	latestHeight uint64
}

func NewBtmcNodeSyncer(service string, nodeURL string) (*eyNodeSyncer, error) {
	return &eyNodeSyncer{
		client:       rpc.NewBtmcClient(service, nodeURL),
		latestHeight: 0,
	}, nil
}

func (n *eyNodeSyncer) fetchBlockTemplate() (ss.BlockTemplate, error) {
	reply, err := n.client.GetWork()
	if err != nil {
		return nil, err
	}

	header := reply.BlockHeader
	if header == nil {
		return nil, ErrNullBlockHeader
	}

	return &eyBlockTemplate{
		version:                header.Version,
		height:                 header.Height,
		previousBlockHash:      &header.PreviousBlockHash,
		timestamp:              header.Time(),
		transactionsMerkleRoot: &header.TransactionsMerkleRoot,
		transactionStatusHash:  &header.TransactionStatusHash,
		nonce:                  header.Nonce,
		bits:                   header.Bits,
		seed:                   reply.Seed,
	}, nil
}

func (n *eyNodeSyncer) Pull() (ss.BlockTemplate, error) {
	return n.fetchBlockTemplate()
}

func (n *eyNodeSyncer) Submit(share ss.Share) error {
	eyShare := share.(*eyShare)
	rawdata, err := n.client.SubmitBlock(&api.SubmitWorkReq{BlockHeader: eyShare.header})
	if err != nil {
		return err
	}

	resultrawdata, err := json.Marshal(rawdata)
	if err != nil {
		return err
	}
	var result bool
	if err := json.Unmarshal(resultrawdata, &result); err != nil {
		return err
	}
	if !result {
		logger.Error("block rejected", "nonce", eyShare.nonce, "hash", eyShare.blockHash)
		return nil
	}
	logger.Info("send nonce success", "nonce", eyShare.nonce)
	return nil
}

func (n *eyNodeSyncer) GetBt() (*api.GetWorkResp, error) {
	n.btLock.RLock()
	defer n.btLock.RUnlock()
	if n.bt == nil {
		return nil, errors.New("getting blocktemplate")
	}
	return n.bt, nil
}
