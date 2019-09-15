package getone

import (
	"testing"
)

func TestHandler(t *testing.T) {

	t.Run("Should return a 200 status", func(t *testing.T) {
		h := NewHandler()
		req := Request{}

		res, err := h.Handle(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		const exp = 200
		if act := res.StatusCode; act != 200 {
			t.Fatalf("Expected %d, got %d", exp, act)
		}
	})

}
