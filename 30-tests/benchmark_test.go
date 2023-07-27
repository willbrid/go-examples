package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
*
La fonction BenchmarkSort crée une tranche avec des données aléatoires et la transmet à la fonction sortAndTotal.
*
*/
func BenchmarkSort(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	size := 250
	data := make([]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		sortAndTotal(data)
	}
}

/*
*
Le timer est réinitialisé une fois que la seed aléatoire est définie et que la tranche a été initialisée. Dans la boucle for,
la méthode StopTimer est utilisée pour arrêter le minuteur avant que la tranche ne soit remplie de données aléatoires, et la méthode StartTimer est
utilisée pour démarrer le minuteur avant l'appel de la fonction sortAndTotal.
*
*/
func BenchmarkSortWithTiming(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	size := 250
	data := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		b.StartTimer()
		sortAndTotal(data)
	}
}

func BenchmarkSortWithSubBenchmark(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	sizes := []int{10, 100, 250}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Array Size %v", size), func(subB *testing.B) {
			data := make([]int, size)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				subB.StopTimer()
				for j := 0; j < size; j++ {
					data[j] = rand.Int()
				}
				subB.StartTimer()
				sortAndTotal(data)
			}
		})
	}
}
