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
		LogGroupName:  aws.String("/ecs/staging-migration-task"),
		LogStreamName: aws.String("ecs/staging-migration-ecs-container/"),
		StartTime:     aws.Int64(time.Now().Unix()),
	}

	ctx := context.Background()
	sess := session.Must(session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	))
	cwl := cloudwatchlogs.New(sess)

	go func() {
		tick := time.Tick(5 * time.Second)
		for {
			select {
			case <-ctx.Done():
				log.Println("done!!!")
				return
			case <-tick:
				out, err := cwl.GetLogEventsWithContext(ctx, input)
				if err != nil {
					panic(err)
				}
				log.Printf("LogEvents: %+v\n", out.String())
			}
		}
	}()
	//out, err := cwl.GetLogEventsWithContext(ctx, input)
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("LogEvents: %+v\n", out.String())
	time.Sleep(30 * time.Second)
}
