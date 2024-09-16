package password

import "testing"

func Test_GIVEN__GeneratePassword__WHEN__Ok(t *testing.T) {
	passwd := GeneratePassword(16)
	if len(passwd) != 16 {
		t.Fatalf("generated password: %s should have length equals 16", passwd)
	}
}
