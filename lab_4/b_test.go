package lab4

// import (
// 	"testing"
// )

// // insertXPreallocIntMap - для добавления X элементов в Map[int]int
// func insertXPreallocIntMap(x int, b *testing.B) {
// 	// Инициализируем Map, вставляем X элементов и предустанавливаем размер X.
// 	testmap := make(map[int]int, x)
// 	// Сбрасываем таймер после инициализации Map.
// 	b.ResetTimer()
// 	for i := 0; i < x; i++ {
// 		// Insert value of I into I key.
// 		testmap[i] = i
// 	}
// }

// // BenchmarkInsertIntMapPreAlloc1000000 тестирует скорость вставки 1000000 целых чисел в карту.
// func BenchmarkInsertIntMapPreAlloc1000000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		insertXPreallocIntMap(1000000, b)
// 	}
// }
