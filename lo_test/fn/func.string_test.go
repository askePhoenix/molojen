package fn

import "testing"

func TestIntToString(t *testing.T) {
	type args struct {
		i   int
		in1 int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"it should success", args{1, 0}, "1"},
		{"it should success", args{2, 0}, "2"},
		{"it should success", args{0, 0}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(tt.args.i, tt.args.in1); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
