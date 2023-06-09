//go:build linux

package symtab

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/grafana/agent/component/pyroscope/ebpf/ebpfspy/symtab/elf"
)

func TestSelfGoSymbolComparison(t *testing.T) {
	f, err := os.Readlink("/proc/self/exe")
	require.NoError(t, err)

	expectedSymbols, err := newGoSymbols(f)
	require.NoError(t, err)

	me, err := elf.NewMMapedElfFile(f)
	require.NoError(t, err)

	goTable, err := me.ReadGoSymbols()
	require.NoError(t, err)

	require.Greater(t, len(expectedSymbols.Symbols), 1000)

	for _, symbol := range expectedSymbols.Symbols {
		name := goTable.Resolve(symbol.Start)
		require.Equal(t, symbol.Name, name)
	}

}