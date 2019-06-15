package main

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(1)
	want := "1"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
