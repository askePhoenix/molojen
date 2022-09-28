package fn

import "testing"

func TestFuncMultiple2(t *testing.T) {
	type args struct {
		i   int
		in1 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"it should success", args{i: 2, in1: 0}, 4},
		{"it should fail, when 0", args{i: 0, in1: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuncMultiple2(tt.args.i, tt.args.in1); got != tt.want {
				t.Errorf("FuncMultiple2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuncPrintInt(t *testing.T) {
	type args struct {
		i   int
		in1 int
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
