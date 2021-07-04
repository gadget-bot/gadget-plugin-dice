package dice

import "testing"

func TestGetMentionRoutes(t *testing.T) {
	ans := GetMentionRoutes()
	if len(ans) != 1 {
		t.Errorf("len(GetMentionRoutes()) = %d; want 1", len(ans))
	}
}

func TestRollDie(t *testing.T) {
	ans := rollDie([]int{1})
	if ans != 1 {
		t.Errorf("rollDie([]int{1}) = %d; want 1", ans)
	}

	ans = rollDie([]int{4, 5, 6})
	if ans != 4 && ans != 5 && ans != 6 {
		t.Errorf("rollDie([]int{4, 5, 6}) = %d; want one of [4, 5, 6]", ans)
	}
}
