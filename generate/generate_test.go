package generate

import "testing"

func TestGenerate(t *testing.T) {
	cases := []struct {
		files []string
		err   bool
	}{
		{
			[]string{"README.md"},
			false,
		},
		{
			[]string{""},
			false,
		},
		{
			[]string{"NOTHING"},
			true,
		},
		{
			[]string{"NOTHING", "README.md"},
			true,
		},
	}

	for _, test := range cases {
		err := Generate(test.files)
		if err != nil {
			if !test.err {
				t.Error("Failed with files", test.files, ":", err)
				continue
			}

			continue
		}

		if test.err {
			t.Error("Did not return error with files", test.files)
		}
	}
}
