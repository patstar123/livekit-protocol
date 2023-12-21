package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestDeferredLogger(t *testing.T) {
	c := &testCore{Core: zap.NewExample().Core()}
	dc, resolve := newDeferredValueCore(c)
	s := zap.New(dc).Sugar()

	s.Infow("test")
	require.Equal(t, 0, c.writeCount)

	resolve("foo", "bar")
	require.Equal(t, 1, c.writeCount)

	s.Infow("test")
	require.Equal(t, 2, c.writeCount)
}
