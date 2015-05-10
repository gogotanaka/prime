package prime

import "testing"

func TestGet(t *testing.T) {
	if p := Get(1); p != 2 {
    t.Errorf("Expect: Get(1) is 2, Actual: %d", p)
  }

  if p := Get(5); p != 11 {
    t.Errorf("Expect: Get(5) is 11, Actual: %d", p)
  }

  if p := Get(10); p != 29 {
    t.Errorf("Expect: Get(10) is 29, Actual: %d", p)
  }
}
