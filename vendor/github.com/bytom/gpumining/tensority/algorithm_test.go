package tensority

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"corepool/core/protocol/bc"
)

var tests = []struct {
	blockHeader [32]byte
	seed        [32]byte
	hash        [32]byte
}{
	{
		blockHeader: [32]byte{
			0xd0, 0xda, 0xd7, 0x3f, 0xb2, 0xda, 0xbf, 0x33,
			0x53, 0xfd, 0xa1, 0x55, 0x71, 0xb4, 0xe5, 0xf6,
			0xac, 0x62, 0xff, 0x18, 0x7b, 0x35, 0x4f, 0xad,
			0xd4, 0x84, 0x0d, 0x9f, 0xf2, 0xf1, 0xaf, 0xdf,
		},
		seed: [32]byte{
			0x07, 0x37, 0x52, 0x07, 0x81, 0x34, 0x5b, 0x11,
			0xb7, 0xbd, 0x0f, 0x84, 0x3c, 0x1b, 0xdd, 0x9a,
			0xea, 0x81, 0xb6, 0xda, 0x94, 0xfd, 0x14, 0x1c,
			0xc9, 0xf2, 0xdf, 0x53, 0xac, 0x67, 0x44, 0xd2,
		},
		hash: [32]byte{
			0xe3, 0x5d, 0xa5, 0x47, 0x95, 0xd8, 0x2f, 0x85,
			0x49, 0xc0, 0xe5, 0x80, 0xcb, 0xf2, 0xe3, 0x75,
			0x7a, 0xb5, 0xef, 0x8f, 0xed, 0x1b, 0xdb, 0xe4,
			0x39, 0x41, 0x6c, 0x7e, 0x6f, 0x8d, 0xf2, 0x27,
		},
	},
	{
		blockHeader: [32]byte{
			0xd0, 0xda, 0xd7, 0x3f, 0xb2, 0xda, 0xbf, 0x33,
			0x53, 0xfd, 0xa1, 0x55, 0x71, 0xb4, 0xe5, 0xf6,
			0xac, 0x62, 0xff, 0x18, 0x7b, 0x35, 0x4f, 0xad,
			0xd4, 0x84, 0x0d, 0x9f, 0xf2, 0xf1, 0xaf, 0xdf,
		},
		seed: [32]byte{
			0x07, 0x37, 0x52, 0x07, 0x81, 0x34, 0x5b, 0x11,
			0xb7, 0xbd, 0x0f, 0x84, 0x3c, 0x1b, 0xdd, 0x9a,
			0xea, 0x81, 0xb6, 0xda, 0x94, 0xfd, 0x14, 0x1c,
			0xc9, 0xf2, 0xdf, 0x53, 0xac, 0x67, 0x44, 0xd2,
		},
		hash: [32]byte{
			0xe3, 0x5d, 0xa5, 0x47, 0x95, 0xd8, 0x2f, 0x85,
			0x49, 0xc0, 0xe5, 0x80, 0xcb, 0xf2, 0xe3, 0x75,
			0x7a, 0xb5, 0xef, 0x8f, 0xed, 0x1b, 0xdb, 0xe4,
			0x39, 0x41, 0x6c, 0x7e, 0x6f, 0x8d, 0xf2, 0x27,
		},
	},
	{
		blockHeader: [32]byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
		seed: [32]byte{
			0x48, 0xdd, 0xa5, 0xbb, 0xe9, 0x17, 0x1a, 0x66,
			0x56, 0x20, 0x6e, 0xc5, 0x6c, 0x59, 0x5c, 0x58,
			0x34, 0xb6, 0xcf, 0x38, 0xc5, 0xfe, 0x71, 0xbc,
			0xb4, 0x4f, 0xe4, 0x38, 0x33, 0xae, 0xe9, 0xdf,
		},
		hash: [32]byte{
			0x26, 0xdb, 0x94, 0xef, 0xa4, 0x22, 0xd7, 0x6c,
			0x40, 0x2a, 0x54, 0xee, 0xb6, 0x1d, 0xd5, 0xf5,
			0x32, 0x82, 0xcd, 0x3c, 0xe1, 0xa0, 0xac, 0x67,
			0x7e, 0x17, 0x70, 0x51, 0xed, 0xaa, 0x98, 0xc1,
		},
	},
	{
		blockHeader: [32]byte{
			0x8d, 0x96, 0x9e, 0xef, 0x6e, 0xca, 0xd3, 0xc2,
			0x9a, 0x3a, 0x62, 0x92, 0x80, 0xe6, 0x86, 0xcf,
			0x0c, 0x3f, 0x5d, 0x5a, 0x86, 0xaf, 0xf3, 0xca,
			0x12, 0x02, 0x0c, 0x92, 0x3a, 0xdc, 0x6c, 0x92,
		},
		seed: [32]byte{
			0x0e, 0x3b, 0x78, 0xd8, 0x38, 0x08, 0x44, 0xb0,
			0xf6, 0x97, 0xbb, 0x91, 0x2d, 0xa7, 0xf4, 0xd2,
			0x10, 0x38, 0x2c, 0x67, 0x14, 0x19, 0x4f, 0xd1,
			0x60, 0x39, 0xef, 0x2a, 0xcd, 0x92, 0x4d, 0xcf,
		},
		hash: [32]byte{
			0xfe, 0xce, 0xc3, 0x36, 0x69, 0x73, 0x75, 0x92,
			0xf7, 0x75, 0x4b, 0x21, 0x5b, 0x20, 0xba, 0xce,
			0xfb, 0xa6, 0x4d, 0x2e, 0x4c, 0xa1, 0x65, 0x6f,
			0x85, 0xea, 0x1d, 0x3d, 0xbe, 0x16, 0x28, 0x39,
		},
	},
	{
		blockHeader: [32]byte{
			0x2f, 0x01, 0x43, 0x11, 0xe0, 0x92, 0x6f, 0xa8,
			0xb3, 0xd6, 0xe6, 0xde, 0x20, 0x51, 0xbf, 0x69,
			0x33, 0x21, 0x23, 0xba, 0xad, 0xfe, 0x52, 0x2b,
			0x62, 0xf4, 0x64, 0x56, 0x55, 0x85, 0x9e, 0x7a,
		},
		seed: [32]byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
		hash: [32]byte{
			0xc1, 0xc3, 0xcf, 0x4c, 0x76, 0x96, 0x8e, 0x29,
			0x67, 0xf0, 0x05, 0x3c, 0x76, 0xf2, 0x08, 0x4c,
			0xc0, 0x1e, 0xd0, 0xfe, 0x97, 0x66, 0x42, 0x8d,
			0xb9, 0x9c, 0x45, 0xbe, 0xdf, 0x0c, 0xdb, 0xe2,
		},
	},
	{
		blockHeader: [32]byte{
			0xe0, 0xe3, 0xc4, 0x31, 0x78, 0xa1, 0x26, 0xd0,
			0x48, 0x71, 0xb9, 0xc5, 0xd0, 0xc6, 0x42, 0xe5,
			0xe0, 0x8b, 0x96, 0x79, 0xa5, 0xf6, 0x6b, 0x82,
			0x1b, 0xd9, 0xa0, 0x30, 0xef, 0xf0, 0x2c, 0xe7,
		},
		seed: [32]byte{
			0x6a, 0xb2, 0x1e, 0x13, 0x01, 0xf5, 0x75, 0x2c,
			0x2f, 0xca, 0x1b, 0x55, 0x98, 0xf4, 0x9d, 0x37,
			0x69, 0x48, 0x2e, 0x07, 0x3c, 0x1f, 0x26, 0xe3,
			0xb8, 0x36, 0x5f, 0x40, 0x55, 0x53, 0xea, 0x31,
		},
		hash: [32]byte{
			0xab, 0xbc, 0x2c, 0xb3, 0x96, 0x38, 0xf6, 0x84,
			0x23, 0x5f, 0xbc, 0x1b, 0x3f, 0xf1, 0x07, 0x94,
			0x59, 0x48, 0xc5, 0x81, 0xb6, 0x92, 0x9b, 0xae,
			0x2c, 0xd6, 0x81, 0x88, 0x9f, 0xf2, 0xd8, 0x24,
		},
	},
	{
		blockHeader: [32]byte{
			0x88, 0x6a, 0x8e, 0x85, 0xb2, 0x75, 0xe7, 0xd6,
			0x5b, 0x56, 0x9b, 0xa5, 0x10, 0x87, 0x5c, 0x0e,
			0x63, 0xde, 0xce, 0x1a, 0x94, 0x56, 0x99, 0x14,
			0xd7, 0x62, 0x4c, 0x0d, 0xac, 0x80, 0x02, 0xf9,
		},
		seed: [32]byte{
			0x9e, 0x62, 0x91, 0x97, 0x0c, 0xb4, 0x4d, 0xd9,
			0x40, 0x08, 0xc7, 0x9b, 0xca, 0xf9, 0xd8, 0x6f,
			0x18, 0xb4, 0xb4, 0x9b, 0xa5, 0xb2, 0xa0, 0x47,
			0x81, 0xdb, 0x71, 0x99, 0xed, 0x3b, 0x9e, 0x4e,
		},
		hash: [32]byte{
			0x6a, 0x1f, 0x27, 0x6c, 0xac, 0x6f, 0x9e, 0x8b,
			0x42, 0x6e, 0xab, 0x46, 0xb5, 0x33, 0xf8, 0x2e,
			0x82, 0xa1, 0x48, 0xc0, 0x3f, 0xb0, 0xa8, 0x69,
			0x34, 0xa8, 0xe5, 0x48, 0x3b, 0x39, 0xda, 0x5e,
		},
	},
	{
		blockHeader: [32]byte{
			0x2f, 0x86, 0xfe, 0x50, 0x8f, 0xaa, 0x7d, 0x68,
			0x69, 0x2a, 0x20, 0x89, 0x27, 0x1b, 0x69, 0x01,
			0x38, 0x5d, 0x90, 0x58, 0xcd, 0x47, 0xe8, 0x4f,
			0xb4, 0x02, 0xb5, 0x08, 0x5f, 0x9a, 0x83, 0x60,
		},
		seed: [32]byte{
			0x9e, 0x62, 0x91, 0x97, 0x0c, 0xb4, 0x4d, 0xd9,
			0x40, 0x08, 0xc7, 0x9b, 0xca, 0xf9, 0xd8, 0x6f,
			0x18, 0xb4, 0xb4, 0x9b, 0xa5, 0xb2, 0xa0, 0x47,
			0x81, 0xdb, 0x71, 0x99, 0xed, 0x3b, 0x9e, 0x4e,
		},
		hash: [32]byte{
			0x81, 0x75, 0x6f, 0xda, 0xb3, 0x9a, 0x17, 0x16,
			0x3b, 0x0c, 0xe5, 0x82, 0xee, 0x4e, 0xe2, 0x56,
			0xfb, 0x4d, 0x1e, 0x15, 0x6c, 0x69, 0x2b, 0x99,
			0x7d, 0x60, 0x8a, 0x42, 0xec, 0xb3, 0x8d, 0x47,
		},
	},
	{
		blockHeader: [32]byte{
			0xfe, 0x97, 0x91, 0xd7, 0x1b, 0x67, 0xee, 0x62,
			0x51, 0x5e, 0x08, 0x72, 0x3c, 0x06, 0x1b, 0x5c,
			0xcb, 0x95, 0x2a, 0x80, 0xd8, 0x04, 0x41, 0x7c,
			0x8a, 0xee, 0xdf, 0x7f, 0x63, 0x3c, 0x52, 0x4a,
		},
		seed: [32]byte{
			0x9e, 0x62, 0x91, 0x97, 0x0c, 0xb4, 0x4d, 0xd9,
			0x40, 0x08, 0xc7, 0x9b, 0xca, 0xf9, 0xd8, 0x6f,
			0x18, 0xb4, 0xb4, 0x9b, 0xa5, 0xb2, 0xa0, 0x47,
			0x81, 0xdb, 0x71, 0x99, 0xed, 0x3b, 0x9e, 0x4e,
		},
		hash: [32]byte{
			0xfa, 0xd5, 0x19, 0x5a, 0x0c, 0x8e, 0x3b, 0x59,
			0x0b, 0x86, 0xa3, 0xc0, 0xa9, 0x5e, 0x75, 0x29,
			0x56, 0x58, 0x88, 0x50, 0x8a, 0xec, 0xca, 0x96,
			0xe9, 0xae, 0xda, 0x63, 0x30, 0x02, 0xf4, 0x09,
		},
	},
	{
		blockHeader: [32]byte{
			0xef, 0x30, 0x67, 0x87, 0xa0, 0x87, 0xc1, 0x18,
			0xfc, 0xb6, 0xd3, 0x51, 0xf0, 0x19, 0x9d, 0xca,
			0x98, 0x05, 0x58, 0x98, 0xe2, 0x08, 0xfb, 0xf1,
			0xa9, 0x34, 0xc9, 0xd7, 0x0b, 0x58, 0xee, 0x4b,
		},
		seed: [32]byte{
			0x9e, 0x62, 0x91, 0x97, 0x0c, 0xb4, 0x4d, 0xd9,
			0x40, 0x08, 0xc7, 0x9b, 0xca, 0xf9, 0xd8, 0x6f,
			0x18, 0xb4, 0xb4, 0x9b, 0xa5, 0xb2, 0xa0, 0x47,
			0x81, 0xdb, 0x71, 0x99, 0xed, 0x3b, 0x9e, 0x4e,
		},
		hash: [32]byte{
			0xff, 0xb0, 0xca, 0xda, 0xb3, 0x14, 0x67, 0x3d,
			0x22, 0x8f, 0x8f, 0xe7, 0x4d, 0x84, 0xa4, 0x65,
			0x2e, 0x01, 0xc8, 0x2c, 0x26, 0x41, 0xd3, 0xe2,
			0xfa, 0x91, 0x48, 0xaf, 0xea, 0xb0, 0xd6, 0x06,
		},
	},
	{
		blockHeader: [32]byte{
			0x10, 0xd2, 0xf1, 0xb2, 0xf4, 0x8e, 0x38, 0x9c,
			0x97, 0xdd, 0xe1, 0xb1, 0xa6, 0x3b, 0xcd, 0x74,
			0x3c, 0x23, 0x40, 0x93, 0x5d, 0x71, 0xe2, 0xc3,
			0x58, 0xba, 0x10, 0xe5, 0x84, 0x69, 0x61, 0x43,
		},
		seed: [32]byte{
			0x9e, 0x62, 0x91, 0x97, 0x0c, 0xb4, 0x4d, 0xd9,
			0x40, 0x08, 0xc7, 0x9b, 0xca, 0xf9, 0xd8, 0x6f,
			0x18, 0xb4, 0xb4, 0x9b, 0xa5, 0xb2, 0xa0, 0x47,
			0x81, 0xdb, 0x71, 0x99, 0xed, 0x3b, 0x9e, 0x4e,
		},
		hash: [32]byte{
			0x6a, 0xdd, 0x5a, 0xf1, 0x82, 0x53, 0xd0, 0x72,
			0x2c, 0x54, 0x2f, 0x7e, 0x71, 0xf4, 0x7d, 0x9c,
			0xb1, 0xa4, 0xd7, 0xfb, 0x7c, 0x7d, 0x26, 0xae,
			0xfe, 0x3c, 0x83, 0x1c, 0xb8, 0x54, 0xf1, 0x7e,
		},
	},
}

// Tests that tensority hash result is correct.
func TestAlgorithm(t *testing.T) {
	startT := time.Now()
	for i, tt := range tests {
		sT := time.Now()
		bhhash := bc.NewHash(tt.blockHeader)
		sdhash := bc.NewHash(tt.seed)
		result := algorithm(&bhhash, &sdhash).Bytes()
		var resArr [32]byte
		copy(resArr[:], result)
		eT := time.Now()

		if !reflect.DeepEqual(resArr, tt.hash) {
			t.Errorf("Test case %d:\n", i+1)
			t.Errorf("Gets\t%x\n", resArr)
			t.Errorf("Expects\t%x\n", tt.hash)
			t.Errorf("FAIL\n\n")
		} else {
			t.Logf("Test case %d:\n", i+1)
			t.Log("Total verification time:", eT.Sub(sT))
			t.Log("PASS\n")
		}
	}
	endT := time.Now()
	t.Log("Avg time:", time.Duration(int(endT.Sub(startT))/len(tests)))
}

func BenchmarkAlgorithm(b *testing.B) {
	bhhash := bc.NewHash(tests[0].blockHeader)
	sdhash := bc.NewHash(tests[0].seed)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		algorithm(&bhhash, &sdhash)
	}
}


func BenchmarkAlgorithmParallel(b *testing.B) {
	bhhash := bc.NewHash(tests[0].blockHeader)
	sdhash := bc.NewHash(tests[0].seed)

	b.SetParallelism(runtime.NumCPU())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			algorithm(&bhhash, &sdhash)
		}
	})
}