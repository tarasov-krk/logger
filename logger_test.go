package logger

import (
	"testing"
)

func TestLevel_IsValid(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want bool
	}{
		{
			name: "valid level Error",
			l:    LevelError,
			want: true,
		},
		{
			name: "not valid level",
			l:    10,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetLevel(t *testing.T) {
	t.Run("Check Set Level", func(t *testing.T) {
		l := NewStd()
		SetLogger(l)
		SetLevel(LevelInfo)

		if l.GetLevel() != LevelInfo {
			t.Errorf("Need level = %v, want = %v", LevelInfo, l.GetLevel())
		}
	})

	t.Run("Set Incorrect Level", func(t *testing.T) {
		l := NewStd()
		SetLogger(l)
		SetLevel(55)

		if l.GetLevel() != 0 {
			t.Errorf("Need level = %v, want = %v", 0, l.GetLevel())
		}
	})
}
