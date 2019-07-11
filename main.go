package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {
	input := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String("/ecs/redish-staging-migration-task"),
		LogStreamName: aws.String("ecs/redish-staging-migration-ecs-container/92012b12-ebbd-4e46-bc19-36ea026c57c2"),
		StartTime:     aws.Int64(time.Now().Unix()),
	}

	ctx := context.Background()
	sess := session.Must(session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	))
	cwl := cloudwatchlogs.New(sess)
	out, err := cwl.GetLogEventsWithContext(ctx, input)
	if err != nil {
		panic(err)
	}
	log.Printf("LogEvents: %+v\n", out.String())
}
