package dice

import "testing"

func TestGetMentionRoutes(t *testing.T) {
	ans := GetMentionRoutes()
	if len(ans) != 1 {
		t.Errorf("len(GetMentionRoutes()) = %d; want 1", len(ans))
	}
}
