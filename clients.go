package cwl

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	cwl "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type Clients map[string]*Client

func NewClients() Clients {

	var regions Clients = make(map[string]*Client)

	for _, e := range endpoints.DefaultResolver().(endpoints.EnumPartitions).Partitions() {

		for region := range e.Regions() {

			s := session.Must(session.NewSession(&aws.Config{
				Region: &region,
			}))

			regions[region] = &Client{
				Client: cwl.New(s),
				Region: region,
			}

		}

		break

	}

	return regions

}

func (c Clients) FindGroups(filter string) (map[string][]string, error) {

	var (
		fErr   error
		groups = make(map[string][]string, len(c))
		wg     sync.WaitGroup
	)

	for region, client := range c {

		wg.Add(1)

		go func(region string, client *Client) {

			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			req := &cwl.DescribeLogGroupsInput{}

			if filter != "" {
				req.LogGroupNamePrefix = &filter
			}

			if err := client.Client.DescribeLogGroupsPagesWithContext(ctx, req, func(res *cwl.DescribeLogGroupsOutput, more bool) bool {

				for _, e := range res.LogGroups {
					groups[region] = append(groups[region], *e.LogGroupName)
				}

				return more

			}); err != nil {
				fErr = err
			}

		}(region, client)

	}

	wg.Wait()

	return groups, fErr

}
