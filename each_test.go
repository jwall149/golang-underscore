package underscore

import (
	"testing"
	"time"
)

func TestEach(t *testing.T) {
	arr := []int{1, 2, 3}
	Each(arr, func(n, i int) {
		if n != arr[i] {
			t.Error("wrong")
		}
	})
}

func BenchmarkFor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s := 0
		for i := 0; i < len(arr); i++ {
			s = s + arr[i]
		}
	}
}

func BenchmarkEach(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s := 0
		Each(arr, func(num, i int) {
			s = s + num
		})
	}
}

func TestChain_Each(t *testing.T) {
	arr := []int{1, 2, 3}
	Chain(arr).Each(func(n, i int) {
		if n != arr[i] {
			t.Error("wrong")
		}
	})
}

func TestChain_Parallel_Each(t *testing.T) {
	arr := []int{1, 2, 3}
	beginUnix := time.Now().Unix()
	Chain(arr).AsParallel().Each(func(n, i int) {
		time.Sleep(time.Second)
	}).Value()
	endUnix := time.Now().Unix()
	if int(endUnix-beginUnix) > len(arr) {
		t.Error("wrong")
	}
}
