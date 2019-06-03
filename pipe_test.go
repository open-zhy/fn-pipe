package fnpipe

import "testing"

func TestPipe_ExecWith(t *testing.T) {
	t.Run("fn(a) -> fn(x, y) -> N", func(t *testing.T) {
		pp, err := NewPipeline(
			func(a int) (int, int) { return a + 5, a - 5 },
			func(a int, b int) int { return a * b },
		)
		if err != nil {
			t.Fatalf("Error while creating pipeline -> %s", err)
		}

		err, res := pp.ExecWith(0)
		if err != nil {
			t.Errorf("cannot execute pipeline: %s", err)
		}

		r, ok := res.([]interface{})
		if !ok {
			t.Fatalf("expected result to slice, got %d, type=%T", res, res)
		}

		t.Logf("r=%+d, type=%T", r, r)
		t.Logf("r[0]=%+d, type=%T", r[0], r[0])
		if r[0].(int) != -25 {
			t.Errorf("expected result to -25, got %d", r[0])
		}
	})

	t.Run("fn(a) -> fn(x, y)", func(t *testing.T) {
		pp, err := NewPipeline(
			func(a int) (int, int) { return a + 3, a - 3 },
		)
		if err != nil {
			t.Fatalf("Error while creating pipeline -> %s", err)
		}

		err, _ = pp.ExecWith(0)
		if err != nil {
			t.Errorf("cannot execute pipeline: %s", err)
		}
	})
}

func TestPipe_ErrorCases(t *testing.T) {
	t.Run("mismatch arguments", func(t *testing.T) {
		pp, err := NewPipeline(
			func(a int) (int, int) { return a + 3, a - 3 },
			func(a int) int { return a },
		)

		if err != nil {
			t.Fatalf("Error while creating pipeline -> %s", err)
		}

		err, _ = pp.ExecWith(0)
		if err == nil {
			t.Fatalf("expect 'argument missmatch on pipe #1' error")
		}
	})

	t.Run("output mismatch", func(t *testing.T) {
		pp, err := NewPipeline(
			func(a int) (int, int) { return a + 3, a - 3 },
			func(a int) int { return a },
		)

		if err != nil {
			t.Fatalf("Error while creating pipeline -> %s", err)
		}

		err, _ = pp.ExecWith(0)
		if err == nil {
			t.Fatalf("expect 'argument missmatch on pipe #1' error")
		}
	})

	t.Run("pipe func should be type Func", func(t *testing.T) {
		_, err := NewPipeline(
			0,
			func(a int) int { return a },
		)

		if err == nil {
			t.Fatalf("expected error while defining pipeline")
		}
	})
}
