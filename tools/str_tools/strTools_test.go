package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrInArr(t *testing.T) {
	// mock StrInArr
	var testArr = []string{
		"a",
		"b",
		"c",
	}
	t.Logf("~> mock StrInArr")
	// do StrInArr
	t.Logf("~> do StrInArr")
	// verify StrInArr
	assert.True(t, StrInArr("c", testArr))
	assert.False(t, StrInArr("d", testArr))
}

func BenchmarkStrInArr(b *testing.B) {
	var testArr = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
	}
	for i := 0; i < b.N; i++ {
		assert.True(b, StrInArr("f", testArr))
	}
}

func Test_Str2LineRaw(t *testing.T) {
	// mock _Str2LineRaw
	mockEnvDroneCommitMessage := "mock message commit\nmore line\nand more line\r\n"
	commitMessage := mockEnvDroneCommitMessage
	t.Logf("~> mock _Str2LineRaw")
	// do _Str2LineRaw
	t.Logf("~> do _Str2LineRaw")
	lineRaw := Str2LineRaw(commitMessage)
	//commitMessage = strings.Replace(commitMessage, "\n", `\\n`, -1)
	t.Logf("lineRaw: %v", lineRaw)
	assert.Equal(t, "mock message commit\\nmore line\\nand more line\\n", lineRaw)
	// verify _Str2LineRaw
}
