package audit

import (
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type AuditHandler interface {
	Write(msg string) error
	Get(limit int64) (*cloudwatchlogs.GetLogEventsOutput, error)
}

type audit struct {
	session *cloudwatchlogs.CloudWatchLogs
	m       *sync.Mutex
	token   *string
	group   string
	stream  string
	region  string
}

func New() (AuditHandler, error) {
	var err error

	var cwGroup string = "api"
	cwGroupValue, cwGroupPresent := os.LookupEnv("AWS_CW_GROUP")
	if cwGroupPresent {
		cwGroup = cwGroupValue
	}

	var cwStream string = "api"
	cwStreamValue, cwStreamPresent := os.LookupEnv("AWS_CW_STREAM")
	if cwStreamPresent {
		cwStream = cwStreamValue
	}

	var cwRegion string = "eu-west-1"
	cwRegionValue, cwRegionPresent := os.LookupEnv("AWS_CW_REGION")
	if cwRegionPresent {
		cwRegion = cwRegionValue
	}

	session, err := session.NewSession(&aws.Config{Region: aws.String(cwRegion)})
	if err != nil {
		return nil, err
	}

	a := &audit{session: cloudwatchlogs.New(session), m: &sync.Mutex{}, token: nil, group: cwGroup, stream: cwStream, region: cwRegion}

	err = a.cwGroup()
	if err != nil {
		return nil, err
	}

	err = a.cwStream()
	if err != nil {
		return nil, err
	}

	resp, err := a.cwNextToken()
	if err != nil {
		return nil, err
	}
	a.token = resp

	return a, nil
}

func (a *audit) cwGroup() error {
	_, err := a.session.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: &a.group,
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

// createLogStream will make a new logStream with a random uuid as its name.
func (a *audit) cwStream() error {
	_, err := a.session.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &a.group,
		LogStreamName: &a.stream,
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

func (a *audit) cwNextToken() (*string, error) {
	resp, err := a.session.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &a.group,
		LogStreamNamePrefix: &a.stream,
	})
	if err != nil {
		return nil, err
	}

	return resp.LogStreams[0].UploadSequenceToken, nil
}

// CwWriteLog will process the log queue
func (a *audit) Write(msg string) error {
	var err error
	var logEvent []*cloudwatchlogs.InputLogEvent

	if a.token == nil {
		resp, err := a.cwNextToken()
		if err != nil {
			return err
		}
		a.token = resp
	}

	logEvent = append(logEvent, &cloudwatchlogs.InputLogEvent{
		Message:   &msg,
		Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
	})

	a.m.Lock()
	resp, err := a.session.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  &a.group,
		LogStreamName: &a.stream,
		LogEvents:     logEvent,
		SequenceToken: a.token,
	})

	if err != nil {
		var err2 error
		a.token, err2 = a.cwNextToken()
		a.m.Unlock()
		if err2 != nil {
			return err2
		}
		return err
	}

	a.token = resp.NextSequenceToken
	a.m.Unlock()
	return nil
}

func (a *audit) Get(limit int64) (*cloudwatchlogs.GetLogEventsOutput, error) {
	resp, err := a.session.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
		Limit:         &limit,
		LogGroupName:  &a.group,
		LogStreamName: &a.stream,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
