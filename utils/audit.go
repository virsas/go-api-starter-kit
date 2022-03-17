package utils

import (
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type Audit struct {
	cw        *cloudwatchlogs.CloudWatchLogs
	mutex     *sync.Mutex
	nextToken *string
}

var cwLogGroup string = "api"
var cwLogStream string = "stream"
var cwRegion string = "eu-west-1"

func InitAudit() (*Audit, error) {
	audit, err := initCW()
	return audit, err
}

func initCW() (*Audit, error) {
	var err error
	var cw Audit

	cwLogGroupValue, cwLogGroupPresent := os.LookupEnv("AWS_CW_GROUP")
	if cwLogGroupPresent {
		cwLogGroup = cwLogGroupValue
	}

	cwLogStreamValue, cwLogStreamPresent := os.LookupEnv("AWS_CW_STREAM")
	if cwLogStreamPresent {
		cwLogStream = cwLogStreamValue
	}

	cwRegionValue, cwRegionPresent := os.LookupEnv("AWS_CW_REGION")
	if cwRegionPresent {
		cwRegion = cwRegionValue
	}

	session, err := session.NewSession(&aws.Config{Region: aws.String(cwRegion)})
	if err != nil {
		return &cw, err
	}

	cw.cw = cloudwatchlogs.New(session)
	cw.mutex = &sync.Mutex{}

	err = cwGroup(cw.cw)
	if err != nil {
		return &cw, err
	}

	err = cwStream(cw.cw)
	if err != nil {
		return &cw, err
	}

	resp, err := cwNextToken(cw.cw)
	if err != nil {
		return &cw, err
	}

	cw.nextToken = resp

	return &cw, nil
}

func cwGroup(cw *cloudwatchlogs.CloudWatchLogs) error {
	_, err := cw.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: &cwLogGroup,
	})
	if err != nil {
		aerr, ok := err.(awserr.Error)
		if ok {
			if aerr.Code() == "ResourceAlreadyExistsException" {
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func cwStream(cw *cloudwatchlogs.CloudWatchLogs) error {
	_, err := cw.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &cwLogGroup,
		LogStreamName: &cwLogStream,
	})
	if err != nil {
		aerr, ok := err.(awserr.Error)
		if ok {
			if aerr.Code() == "ResourceAlreadyExistsException" {
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func cwNextToken(cw *cloudwatchlogs.CloudWatchLogs) (*string, error) {
	resp, err := cw.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &cwLogGroup,
		LogStreamNamePrefix: &cwLogStream,
	})
	if err != nil {
		return nil, err
	}

	return resp.LogStreams[0].UploadSequenceToken, nil
}

func CwWriteLog(cw *Audit, msg string) error {
	var err error
	var logEvent []*cloudwatchlogs.InputLogEvent

	logEvent = append(logEvent, &cloudwatchlogs.InputLogEvent{
		Message:   &msg,
		Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
	})

	cw.mutex.Lock()
	resp, err := cw.cw.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  &cwLogGroup,
		LogStreamName: &cwLogStream,
		LogEvents:     logEvent,
		SequenceToken: cw.nextToken,
	})
	cw.nextToken = resp.NextSequenceToken
	cw.mutex.Unlock()
	if err != nil {
		return err
	}

	return nil
}
