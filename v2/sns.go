package notificationsdk

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

type SnsClient interface {
}

type snsClient struct {
	snsClient *sns.Client
	topic     string
}

func NewSnsClient(client *sns.Client, topic string) Client {
	return &snsClient{client, topic}
}

func (s *snsClient) send(ctx context.Context, payload any, groupID string, _type string) error {
	bPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	stringPayload := string(bPayload)

	topicList, err := s.snsClient.ListTopics(ctx, &sns.ListTopicsInput{})
	if err != nil {
		return err
	}

	var topicArn string
	for _, t := range topicList.Topics {
		topicName := strings.Split(*t.TopicArn, ":")[5]
		if topicName == s.topic {
			topicArn = *t.TopicArn
			break
		}
	}

	_, err = s.snsClient.Publish(ctx, &sns.PublishInput{
		Message:        &stringPayload,
		TopicArn:       &topicArn,
		MessageGroupId: &groupID,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"type": {
				DataType:    aws.String("String.Array"),
				StringValue: aws.String(_type),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
