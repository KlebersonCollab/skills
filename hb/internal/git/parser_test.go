package git

import (
	"testing"
)

func TestParseDiff(t *testing.T) {
	diff := `diff --git a/test.txt b/test.txt
index 12345..67890 100644
--- a/test.txt
+++ b/test.txt
@@ -1,3 +1,4 @@
 line 1
-line 2
+line 2 mod
 line 3
+line 4`

	hunks, err := ParseDiff(diff)
	if err != nil {
		t.Fatalf("Failed to parse diff: %v", err)
	}

	if len(hunks) != 1 {
		t.Errorf("Expected 1 hunk, got %d", len(hunks))
	}

	if hunks[0].FilePath != "test.txt" {
		t.Errorf("Expected FilePath test.txt, got %s", hunks[0].FilePath)
	}

	if hunks[0].OldStart != 1 || hunks[0].NewStart != 1 {
		t.Errorf("Incorrect hunk starts")
	}
}
