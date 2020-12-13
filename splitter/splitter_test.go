package splitter

import (
	"testing"
)

func BenchmarkFileSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts, _ := Split("splitter_test.go", 4, 3)
		_ = PartsToFiles(parts, "file_%d.part")
		_, _ = CombineFiles("file_0.part", "file_1.part",
			"file_3.part", "file_4.part")
	}
}
