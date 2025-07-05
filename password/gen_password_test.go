package password

import (
	"testing"
	"unicode"
)

func TestGenPassLength(t *testing.T) {
	pwSize := 12
	pass, _ := GenPass(true, true, true, pwSize)
	if len(pass) != pwSize{
		t.Errorf("fail: expected password of size %d, got %d", pwSize, len(pass))
	}
}

func TestGenPassNoUppers(t *testing.T) {
	pwSize := 12
	pass, _ := GenPass(false, true, true, pwSize)
	hasUpper := false
	for _, x := range pass{
		if unicode.IsUpper(x){
			hasUpper = true
			break
		}
	}
	if hasUpper{
		t.Errorf("fail: requested to generate password without upper-case chars, but password has upper-case chars")
	}
}

func TestGenPassNoLowers(t *testing.T) {
	pwSize := 12
	pass, _ := GenPass(true, false, true, pwSize)
	hasLower := false
	for _, x := range pass{
		if unicode.IsLower(x){
			hasLower = true
			break
		}
	}
	if hasLower{
		t.Errorf("fail: requested to generate password without lower-case chars, but password has lower-case chars")
	}
}

func TestGenPassNoDigits(t *testing.T) {
	pwSize := 12
	pass, _ := GenPass(true, true, false, pwSize)
	hasDigits := false
	for _, x := range pass{
		if x <= '9' && x >= '0'{
			hasDigits = true
			break
		}
	}
	if hasDigits{
		t.Errorf("fail: requested to generate password without numeric digits, but password has numeric digits")
	}
}