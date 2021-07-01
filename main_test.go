package main

import "testing"

func TestTotalLending(t *testing.T) {
	pedro := Student{
		subjects: []Subject{
			{ academicCredits: 5 },
			{ academicCredits: 1 },
			{ academicCredits: 3 },
			{ academicCredits: 3 },
			{ academicCredits: 3 },
		},
	}

	result := 15
	total := TotalLending(&pedro)
	if total != result {
		t.Errorf("sum academic credits was incorrect %d, expected %d", total, result)
	}

}

func TestParcialSum(t *testing.T) {
	juan := Student{
		subjects: []Subject{
			{ academicCredits: 5, grade: 3.0 },
			{ academicCredits: 1, grade: 5.0 },
			{ academicCredits: 3, grade: 4.0 },
			{ academicCredits: 3, grade: 3.0 },
			{ academicCredits: 3, grade: 5.0 },
		},
	}

	result := 56
	total := ParcialSum(&juan)
	if total != result {
		t.Errorf("sum academic credits was incorrect %d, expected %d", total, result)
	}
}

func TestAverage(t *testing.T) {
	tables := []struct {
		a int
		b int
		result float32
	} {
		{56, 15, 3.7333333},
		{86, 16, 5.375},
	}

	for _, item := range tables {
		total := Average(item.a, item.b)
		if total != item.result {
			t.Errorf("Sum was incorrect, got %v, expected %v", total, item.result)
		}
	}
}