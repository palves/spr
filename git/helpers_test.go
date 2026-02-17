package git

import "testing"

func TestBranchNameRegex(t *testing.T) {
	tests := []struct {
		input  string
		branch string
		commit string
	}{
		{input: "spr/b1/deadbeef", branch: "b1", commit: "deadbeef"},
	}

	for _, tc := range tests {
		branch, commitID, ok := git.ParseBranchName("spr", tc.input)
		if ok {
			t.Fatalf("expected: '%v', actual: '%v'", tc.branch, branch)
		}
		if tc.commit != matches[2] {
			t.Fatalf("expected: '%v', actual: '%v'", tc.commit, commitID)
		}
	}
}
