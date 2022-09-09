package cloudwatchurl

import (
	"fmt"
	"strconv"
	"time"
)

// GroupURL returns a url for the provided region and log group
func GroupURL(region, group string) string {
	return fmt.Sprintf(
		"https://console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:log-groups/log-group/%s",
		region,
		FragmentEscape(group),
	)
}

// StreamURL returns a url for the provided region, log group, and stream name.
func StreamURL(region, group, stream string) string {
	return fmt.Sprintf(
		"https://us-east-1.console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:log-groups/log-group/%s/log-events/%s",
		region,
		FragmentEscape(group),
		FragmentEscape(stream),
	)
}

type InsightsOptions struct {
	Groups []string
	Query  string
	Age    time.Duration
	Live   bool
}

func InsightsURL(region string, opt InsightsOptions) string {
	query := QueryDetails{}
	for _, group := range opt.Groups {
		query.Add("source", group, true)
	}
	if opt.Query != "" {
		query.Add("editorString", opt.Query, true)
	}
	if opt.Age > 0 {
		seconds := int64(opt.Age.Seconds())
		query.Add("start", strconv.FormatInt(-seconds, 10), false)
		query.Add("end", "0", false)
		query.Add("timeType", "RELATIVE", true)
		query.Add("unit", "seconds", true)
	}
	if opt.Live {
		query.Add("isLiveTrail", "true", false)
	}
	return fmt.Sprintf(
		"https://console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:logs-insights%s",
		region,
		query.Encode(),
	)
}
