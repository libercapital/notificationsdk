package notificationsdk

import (
	"context"
	"encoding/json"

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

func (s *snsClient) send(ctx context.Context, payload any, snsParams SNSParams, _type string) error {
	bPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	stringPayload := string(bPayload)

	_, err = s.snsClient.Publish(ctx, &sns.PublishInput{
		Message:                &stringPayload,
		TopicArn:               &s.topic,
		MessageGroupId:         &snsParams.GroupID,
		MessageDeduplicationId: &snsParams.DeduplicationID,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"type": {
				DataType:    aws.String("String"),
				StringValue: aws.String(_type),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
