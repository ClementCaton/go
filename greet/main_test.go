package main

import "testing"

func TestPrinter(t *testing.T) {

	cases := []struct {
		name string
		lang string
		want string
	}{
		{"Algosup", "french", "Bonjour Algosup"},
		{"Lucien", "spanish", "Hola Lucien"},
		{"Guillaume", "english", "Hello Guillaume"},
	}

	for _, test := range cases {
		t.Run("Greet Algosup", func(t *testing.T) {
			var a = greeter(test.name, test.lang)
			if a != test.want {
				t.Errorf("Expected %s, got %s", test.want, a)
			}
		})
	}
}
