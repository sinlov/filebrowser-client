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

func TestStrArrRemoveDuplicates(t *testing.T) {
	// mock StrArrRemoveDuplicates

	t.Logf("~> mock StrArrRemoveDuplicates")
	var fooArr = []string{
		"a", "b", "c", "b", "a", "d", "f",
	}
	// do StrArrRemoveDuplicates
	t.Logf("~> do StrArrRemoveDuplicates")
	rdFooArr := StrArrRemoveDuplicates(fooArr)
	// verify StrArrRemoveDuplicates
	assert.Equal(t, 5, len(rdFooArr))

	var barArr []string
	for i := 0; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 1000; i < 2000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 4000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}

	rdBarArr := StrArrRemoveDuplicates(barArr)
	// verify StrArrRemoveDuplicates
	assert.Equal(t, 5000, len(rdBarArr))
}

func BenchmarkStrArrRemoveDuplicates(b *testing.B) {
	var fooArr = []string{
		"a", "b", "c", "b", "a", "d", "f",
	}
	var barArr []string
	for i := 0; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 1000; i < 2000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 4000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 3000; i < 5000; i++ {
		barArr = append(barArr, string(rune(i)))
	}
	for i := 0; i < b.N; i++ {
		rdFooArr := StrArrRemoveDuplicates(fooArr)
		// verify StrArrRemoveDuplicates
		assert.Equal(b, 5, len(rdFooArr))

		rdBarArr := StrArrRemoveDuplicates(barArr)
		// verify StrArrRemoveDuplicates
		assert.Equal(b, 5000, len(rdBarArr))
	}
}
