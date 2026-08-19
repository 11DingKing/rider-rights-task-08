package domain

import (
	"testing"
	"time"
)

func TestEscalationDeadlineExtendsExistingDeadline(t *testing.T) {
	old := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	now := old.Add(2 * time.Hour)
	got := NextEscalationDeadline(old, now, 30*time.Minute)
	if want := old.Add(30 * time.Minute); !got.Equal(want) {
		t.Fatalf("deadline moved from scan time: got %s want %s", got, want)
	}
}
