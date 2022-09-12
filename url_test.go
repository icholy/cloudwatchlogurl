package cloudwatchlogurl

import "testing"

func TestGroupURL(t *testing.T) {
	tests := []struct {
		region, group, url string
	}{
		{
			region: "us-east-1",
			group:  "/aws/lambda/cdl-email-staging-email_catchall",
			url:    "https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcdl-email-staging-email_catchall",
		},
	}
	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			got := GroupURL(tt.region, tt.group)
			if got != tt.url {
				t.Fatalf("GroupURL(%q, %q): got %q, want %q", tt.region, tt.group, got, tt.url)
			}
		})
	}
}
