package cloudwatchlogurl

import "testing"

func TestQueryEscape(t *testing.T) {
	tests := []struct {
		input, output string
	}{

		{
			input:  "?queryDetail=%7E%28end%7E0%7Estart%7E-3600%7EtimeType%7E%27RELATIVE%7Eunit%7E%27seconds%7EeditorString%7E%27fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20%27example%27*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100%7EisLiveTrail%7Efalse%7Esource%7E%28%7E%27*2Fapplication*2Fsample1%7E%27*2Fapplication*2Fsample2%29%29",
			output: "%3FqueryDetail%3D%257E%2528end%257E0%257Estart%257E-3600%257EtimeType%257E%2527RELATIVE%257Eunit%257E%2527seconds%257EeditorString%257E%2527fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20%2527example%2527*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100%257EisLiveTrail%257Efalse%257Esource%257E%2528%257E%2527*2Fapplication*2Fsample1%257E%2527*2Fapplication*2Fsample2%2529%2529",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			output := QueryEscape(tt.input)
			if output != tt.output {
				t.Fatalf("QueryEscape(%q): got %q, want %q", tt.input, output, tt.output)
			}
		})
	}
}

func TestFragmentEscape(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{
			input:  "p4-audit/tasks",
			output: "p4-audit$252Ftasks",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			output := FragmentEscape(tt.input)
			if output != tt.output {
				t.Fatalf("FragmentEscape(%q): got %q, want %q", tt.input, output, tt.output)
			}
		})
	}
}

func TestEcmaEscape(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{
			input:  "~(end~0~start~-3600~timeType~'RELATIVE~unit~'seconds~editorString~'fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20'example'*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100~isLiveTrail~false~source~(~'*2Fapplication*2Fsample1~'*2Fapplication*2Fsample2))",
			output: "%7E%28end%7E0%7Estart%7E-3600%7EtimeType%7E%27RELATIVE%7Eunit%7E%27seconds%7EeditorString%7E%27fields*20*40timestamp*2C*20*40message*0A*20*20*20*20*7C*20filter*20*40message*20not*20like*20%27example%27*0A*20*20*20*20*7C*20sort*20*40timestamp*20asc*0A*20*20*20*20*7C*20limit*20100%7EisLiveTrail%7Efalse%7Esource%7E%28%7E%27*2Fapplication*2Fsample1%7E%27*2Fapplication*2Fsample2%29%29",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			output := EcmaEscape(tt.input)
			if output != tt.output {
				t.Fatalf("EcmaEscape(%q): got %q, want %q", tt.input, output, tt.output)
			}
		})
	}
}
