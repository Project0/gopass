package manifest

import "testing"

func TestManifest(t *testing.T) {
	if _, err := GetLocation("foobar", "", false); err == nil {
		t.Error("browser should not exist")
	}
}
