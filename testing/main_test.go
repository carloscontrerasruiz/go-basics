package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)

	// if total != 10 {
	// 	t.Errorf("Expected %d got %d", 10, total)
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)

		if total != item.n {
			t.Errorf("Expected %d got %d", item.n, total)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{5, 2, 5},
		{22, 26, 26},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)

		if max != item.n {
			t.Errorf("Expected %d got %d", item.n, max)
		}
	}
}
