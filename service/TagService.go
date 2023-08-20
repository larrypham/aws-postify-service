package service

import (
	"aws-postify-service/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetTags() ([]string, error) {
	const maxNumTags = 20

	queryTags := dynamodb.QueryInput{
		TableName:              aws.String(TagTableName),
		IndexName:              aws.String("ArticleCount"),
		KeyConditionExpression: aws.String("Dummy=:zero"),
		Limit:                  aws.Int64(maxNumTags),
		ScanIndexForward:       aws.Bool(false),
	}

	items, err := QueryItems(&queryTags, 0, maxNumTags)
	if err != nil {
		return nil, err
	}

	tagObjects := make([]model.Tag, len(items))
	err = dynamodbattribute.UnmarshalListOfMaps(items, &tagObjects)
	if err != nil {
		return nil, err
	}

	tags := make([]string, 0, len(tagObjects))
	for _, tagObject := range tagObjects {
		if tagObject.ArticleCount > 0 {
			tags = append(tags, tagObject.Tag)
		}
	}
	return tags, nil
}
