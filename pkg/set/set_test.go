package set

import "testing"

func TestUpdate(t *testing.T) {
	// Test the Update function
	t.Run("update weight of set", func(t *testing.T) {
		set := Set{SetNumber: 1, Weight: 10.5, Reps: 8}
		set.Update("weight", 12.0)

		var want float32 = 12.0
		got := set.Weight

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
	t.Run("update reps of set", func(t *testing.T) {
		set := Set{SetNumber: 1, Weight: 10.5, Reps: 8}
		set.Update("reps", 10)

		want := 10
		got := set.Reps

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestGet(t *testing.T) {
	// Test Get() function
	set := Set{SetNumber: 1, Weight: 10.5, Reps: 8}

	want := [3]float32{float32(1), 10.5, float32(8)}

	setNum := 1
	got := set.Get(setNum)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
