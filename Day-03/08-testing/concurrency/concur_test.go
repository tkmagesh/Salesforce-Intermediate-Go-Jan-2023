package concurrency

import "testing"

func Test_Concurrency(t *testing.T) {
	counter := Counter{}

	t.Run("Test_1", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 100; i++ {
			counter.Increment()
		}
	})

	t.Run("Test_2", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 100; i++ {
			counter.Increment()
		}
	})

	t.Run("Test_3", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			counter.Increment()
		}
	})

	t.Run("Test_4", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			counter.Increment()
		}
	})

	t.Run("Test_5", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 100; i++ {
			counter.Increment()
		}
	})

}
