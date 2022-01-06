package main

import "testing"

func TestCheckLang(t *testing.T) {
	cases := []struct {
		text   string
		expect Param
	}{
		{
			text: "日本語を含むtext",
			expect: Param{
				source: "ja",
				target: "en",
			},
		},
		{
			text: "Text only english.",
			expect: Param{
				source: "en",
				target: "ja",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.text, func(t *testing.T) {
			result := checkLang(c.text)
			if result.source != c.expect.source || result.target != c.expect.target {
				t.Errorf("- expected\n%v\n=====\n- actual\n%v", c.expect, result)
			}
		})
	}
}
