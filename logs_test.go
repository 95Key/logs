package logs

import "testing"

func Test_logName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "日志文件名",
			want: "data/log/wangchenyang.log",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logName(); got != tt.want {
				t.Errorf("logName() = %v, want %v", got, tt.want)
			}
		})
	}
}
