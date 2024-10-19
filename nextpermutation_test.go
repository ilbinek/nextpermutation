package nextpermutation

import "testing"

func TestNextPerm(t *testing.T) {
	perm := []int{1, 2, 3}
	Next(&perm, 0, 2)
	if perm[0] != 1 || perm[1] != 3 || perm[2] != 2 {
		t.Errorf("Expected 1 3 2, got %v", perm)
	}

	perm = []int{3, 2, 1}
	r := Next(&perm, 0, 1)
	if r {
		t.Errorf("Expected false, got %v", r)
	}
}

func TestNextString(t *testing.T) {
	str := "abc"
	NextString(&str, 0, 2)
	if str != "acb" {
		t.Errorf("Expected abc, got %v", str)
	}

	str = "cba"
	r := NextString(&str, 0, 1)
	if r {
		t.Errorf("Expected false, got %v", r)
	}
}


