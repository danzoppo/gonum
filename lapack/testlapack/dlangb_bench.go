// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testlapack

import (
	"fmt"
	"testing"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/lapack"
)

func DlangbBenchmark(b *testing.B, impl Dlangber) {
	rnd := rand.New(rand.NewSource(1))
	for _, bm := range []struct {
		n, k int
	}{
		{n: 1000, k: 0},
		{n: 1000, k: 1},
		{n: 1000, k: 2},
		{n: 1000, k: 10},
		{n: 1000, k: 20},
		{n: 1000, k: 30},
		{n: 10000, k: 0},
		{n: 10000, k: 1},
		{n: 10000, k: 2},
		{n: 10000, k: 10},
		{n: 10000, k: 30},
		{n: 10000, k: 60},
		{n: 10000, k: 100},
	} {
		n := bm.n
		k := bm.k
		lda := 2*k + 1
		a := make([]float64, n*lda)
		for i := range a {
			a[i] = rnd.NormFloat64()
		}
		for _, norm := range []lapack.MatrixNorm{lapack.MaxAbs, lapack.MaxRowSum, lapack.MaxColumnSum, lapack.Frobenius} {
			name := fmt.Sprintf("%v_N=%v_K=%v", normToString(norm), n, k)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					result = impl.Dlangb(norm, n, n, k, k, a, lda)
				}
			})
		}
	}
}
