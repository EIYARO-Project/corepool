package protocol

import (
	"testing"
	"time"

	"corepool/core/protocol/bc"
	"corepool/core/protocol/bc/types"
	"corepool/core/testutil"
)

var testBlocks = []*types.Block{
	&types.Block{BlockHeader: types.BlockHeader{
		PreviousBlockHash: bc.Hash{V0: 1},
		Nonce:             0,
	}},
	&types.Block{BlockHeader: types.BlockHeader{
		PreviousBlockHash: bc.Hash{V0: 1},
		Nonce:             1,
	}},
	&types.Block{BlockHeader: types.BlockHeader{
		PreviousBlockHash: bc.Hash{V0: 2},
		Nonce:             3,
	}},
}

var blockHashes = []bc.Hash{}

func init() {
	for _, block := range testBlocks {
		blockHashes = append(blockHashes, block.Hash())
	}
}

func TestOrphanManageAdd(t *testing.T) {
	cases := []struct {
		before    *OrphanManage
		after     *OrphanManage
		addOrphan *types.Block
	}{
		{
			before: &OrphanManage{
				orphan:      map[bc.Hash]*orphanBlock{},
				prevOrphans: map[bc.Hash][]*bc.Hash{},
			},
			after: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			addOrphan: testBlocks[0],
		},
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			after: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			addOrphan: testBlocks[0],
		},
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			after: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
					blockHashes[1]: &orphanBlock{testBlocks[1], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0], &blockHashes[1]},
				},
			},
			addOrphan: testBlocks[1],
		},
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			after: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
					blockHashes[2]: &orphanBlock{testBlocks[2], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
					bc.Hash{V0: 2}: []*bc.Hash{&blockHashes[2]},
				},
			},
			addOrphan: testBlocks[2],
		},
	}

	for i, c := range cases {
		c.before.Add(c.addOrphan)
		for _, orphan := range c.before.orphan {
			orphan.expiration = time.Time{}
		}
		if !testutil.DeepEqual(c.before, c.after) {
			t.Errorf("case %d: got %v want %v", i, c.before, c.after)
		}
	}
}

func TestOrphanManageDelete(t *testing.T) {
	cases := []struct {
		before *OrphanManage
		after  *OrphanManage
		remove *bc.Hash
	}{
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			after: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			remove: &blockHashes[1],
		},
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			after: &OrphanManage{
				orphan:      map[bc.Hash]*orphanBlock{},
				prevOrphans: map[bc.Hash][]*bc.Hash{},
			},
			remove: &blockHashes[0],
		},
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
					blockHashes[1]: &orphanBlock{testBlocks[1], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0], &blockHashes[1]},
				},
			},
			after: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{testBlocks[0], time.Time{}},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			remove: &blockHashes[1],
		},
	}

	for i, c := range cases {
		c.before.delete(c.remove)
		if !testutil.DeepEqual(c.before, c.after) {
			t.Errorf("case %d: got %v want %v", i, c.before, c.after)
		}
	}
}

func TestOrphanManageExpire(t *testing.T) {
	cases := []struct {
		before *OrphanManage
		after  *OrphanManage
	}{
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{
						testBlocks[0],
						time.Unix(1633479700, 0),
					},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			after: &OrphanManage{
				orphan:      map[bc.Hash]*orphanBlock{},
				prevOrphans: map[bc.Hash][]*bc.Hash{},
			},
		},
		{
			before: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{
						testBlocks[0],
						time.Unix(1633479702, 0),
					},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
			after: &OrphanManage{
				orphan: map[bc.Hash]*orphanBlock{
					blockHashes[0]: &orphanBlock{
						testBlocks[0],
						time.Unix(1633479702, 0),
					},
				},
				prevOrphans: map[bc.Hash][]*bc.Hash{
					bc.Hash{V0: 1}: []*bc.Hash{&blockHashes[0]},
				},
			},
		},
	}

	for i, c := range cases {
		c.before.orphanExpire(time.Unix(1633479701, 0))
		if !testutil.DeepEqual(c.before, c.after) {
			t.Errorf("case %d: got %v want %v", i, c.before, c.after)
		}
	}
}
