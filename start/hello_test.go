package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Zhanserik"},
			result: "Hello, Zhanserik!",
		},
		{
			items:  []string{"Zhanserik", "Bekarys", "Alikhan"},
			result: "Hello, Zhanserik, Bekarys, Alikhan!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}
