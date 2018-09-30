package main

import (
	"testing"
)

func TestCalcXY(t *testing.T) {

	// Отсчёт от 0 так что макс значение будет меньше на еденицу
	config = Config{
		Width:    66, // max y, 65 max
		Height:   73, // max x, 72 max
	}

	size := config.Height * config.Width

	// 66x73-1=4817
	// 0	-> (0,0)
	// 1	-> (0,1)
	// 3	-> (0,3)
	// 88	-> (1,15)
	// 4817	-> (65,72)
	// 8888	-> fail, out of range

	if x, y := calcXY(0); x != 0 && y != 0 {
		t.Fatalf("calcXY(0): %d != 0, %d != 0", x, y)
	}

	if x, y := calcXY(1); x != 0 && y != 1 {
		t.Fatalf("calcXY(1): %d != 0, %d != 1", x, y)
	}

	if x, y := calcXY(3); x != 0 && y != 3 {
		t.Fatalf("calcXY(3): %d != 0, %d != 3", x, y)
	}

	if x, y := calcXY(88); x != 1 && y != 15 {
		t.Fatalf("calcXY(88): %d != 1, %d != 15", x, y)
	}

	if x, y := calcXY(size-1); x != config.Width - 1 && y != config.Height - 1 {
		t.Fatalf("calcXY(size-1): %d != %d, %d != %d", x, config.Width - 1, y, config.Height - 1)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic: out of range")
		}
	}()

	calcXY(8888)
}

func TestResolveXY(t *testing.T) {

	config = Config{
		Width:    66,
		Height:   73,
	}

	// 66x73-1=4817
	// 0	-> (0,0)
	// 1	-> (0,1)
	// 3	-> (0,3)
	// 88	-> (1,15)
	// 4817	-> (65,72)

	if i := resolveXY(0, 0); i != 0 {
		t.Fatalf("resolveXY(0, 0): %d != 0", i)
	}

	if i := resolveXY(0, 1); i != 1 {
		t.Fatalf("resolveXY(0, 1): %d != 1", i)
	}

	if i := resolveXY(0, 3); i != 3 {
		t.Fatalf("resolveXY(0, 3): %d != 3", i)
	}

	if i := resolveXY(1, 15); i != 88 {
		t.Fatalf("resolveXY(1, 15): %d != 8", i)
	}

	if i := resolveXY(65, 72); i != 4817 {
		t.Fatalf("resolveXY(0, 1): %d != 4817", i)
	}
}

func TestResolveXY_OutOfRangeX(t *testing.T)  {

	config = Config{
		Width:    66,
		Height:   73,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic: out of range X")
		}
	}()

	if i := resolveXY(65, 999); i != 4817 {
		t.Fatalf("resolveXY(0, 1): %d != 4817", i)
	}
}

func TestResolveXY_OutOfRangeY(t *testing.T)  {

	config = Config{
		Width:    66,
		Height:   73,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic: out of range Y")
		}
	}()

	if i := resolveXY(999, 72); i != 4817 {
		t.Fatalf("resolveXY(0, 1): %d != 4817", i)
	}
}