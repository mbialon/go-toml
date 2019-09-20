package toml

import "testing"

func TestInlineTable(t *testing.T) {
	cases := []struct {
		Input    string
		Expected map[string]interface{}
	}{
		{
			`foo = { -bar = "buz"}`,
			map[string]interface{}{
				"foo": map[string]interface{}{
					"-bar": "buz",
				},
			},
		},
		{
			`foo = { "whatever" = "buz"}`,
			map[string]interface{}{
				"foo": map[string]interface{}{
					"whatever": "buz",
				},
			},
		},
		{
			`foo = { _no = "buz"}`,
			map[string]interface{}{
				"foo": map[string]interface{}{
					"_no": "buz",
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run("", func(t *testing.T) {
			tree, err := Load(tc.Input)
			assertTree(t, tree, err, tc.Expected)
		})
	}
}
