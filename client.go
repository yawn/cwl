package cwl

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	cwl "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type Client struct {
	Client    *cwl.CloudWatchLogs
	NextToken *string
	Region    string
	StartTime *int64
}

func (c *Client) FindEvents(callback func(*Event), group string) error {

	for {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		more, err := c.findEvents(ctx, callback, group)

		if err != nil {
			return err
		}

		if !more {
			time.Sleep(1 * time.Second)
		}

	}

}

func (c *Client) findEvents(ctx context.Context, callback func(*Event), group string) (bool, error) {

	var lastTimestamp int64

	if c.StartTime != nil {
		lastTimestamp = *c.StartTime
	}

	req := &cwl.FilterLogEventsInput{
		Interleaved:  aws.Bool(true),
		Limit:        aws.Int64(500),
		LogGroupName: &group,
		NextToken:    c.NextToken,
		StartTime:    c.StartTime,
	}

	res, err := c.Client.FilterLogEventsWithContext(ctx, req)

	if err != nil {
		return false, err
	}

	for _, event := range res.Events {

		callback(&Event{
			Group:     group,
			Message:   strings.TrimRight(*event.Message, "\n"),
			Region:    c.Region,
			Stream:    *event.LogStreamName,
			Timestamp: time.Unix((*event.Timestamp)/1000, 0).Format(time.RFC3339),
		})

		if timestamp := *event.Timestamp; timestamp > lastTimestamp {
			timestamp++
			c.StartTime = &timestamp
		}

	}

	c.NextToken = res.NextToken

	return res.NextToken != nil, nil

}
