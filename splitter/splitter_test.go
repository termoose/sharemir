package splitter

import (
	"testing"
)

func BenchmarkFileSplit(b *testing.B) {
	var parts []Part

	b.Run("Splitting", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			parts, _ = Split("splitter_test.go", 4, 3)
		}
	})

	b.Run("ToFiles", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = PartsToFiles(parts, "file_%d.part")
		}
	})

	b.Run("Combine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = CombineFiles("file_0.part", "file_1.part",
				"file_3.part", "file_4.part")
		}
	})

	b.Run("CombineReader", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = CombineFilesReader("file_0.part", "file_1.part",
				"file_3.part", "file_4.part")
		}
	})
}
