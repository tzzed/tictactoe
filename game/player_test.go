package game

import "testing"

func TestPlayer_FindInDiag(t *testing.T) {
	type fields struct {
		ID     int
		Mat    [][]int
		Isfull bool
		Wins   int
		Op     byte
	}
	tests := []struct {
		name   string
		fields fields
	}{

		// TODO: Add test cases.

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				ID:     1,
				Mat:    tt.fields.Mat,
				Isfull: false,
				Wins:   0,
				Op:     'X',
			}
			p.Mat = [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}

			p.FindInDiag()
			if !p.Isfull {
				t.Errorf("FindInDiag() Failed, expected %v,  and receive %v", true, p.Isfull)
			}

			p.Mat = [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			}
			if !p.Isfull {
				t.Logf("FindInDiag() success the Mat is not full in Diag, expected %v,  and receive %v", true, p.Isfull)
			}
		})
	}
}
