// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lad

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeToMillis(t *testing.T) {
	tests := []struct {
		t     time.Time
		stamp int64
	}{
		{t: time.Unix(0, 0), stamp: 0},
		{t: time.Unix(1, 0), stamp: 1000},
		{t: time.Unix(1, int64(500*time.Millisecond)), stamp: 1500},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.stamp, timeToMillis(tt.t), "Unexpected timestamp for time %v.", tt.t)
	}
}
