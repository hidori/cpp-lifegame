package lifegame

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		width  int
		height int
		cells  []bool
	}
	tests := []struct {
		name string
		args args
		want *Lifegame
	}{
		{
			name: "",
			args: args{
				width:  2,
				height: 3,
				cells:  make([]bool, 6),
			},
			want: &Lifegame{
				Width:  2,
				Height: 3,
				Cells:  make([]bool, 6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.width, tt.args.height, tt.args.cells); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLifegame_CountNeighbors(t *testing.T) {
	type fields struct {
		Width  int
		Height int
		Cells  []bool
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Lifegame{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
				Cells:  tt.fields.Cells,
			}
			if got := g.CountNeighbors(tt.args.i); got != tt.want {
				t.Errorf("Lifegame.CountNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}
