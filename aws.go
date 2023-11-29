package decimal

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (this *Decimal) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if av.N == nil {
		return nil
	}
	var err error
	*this, err = NewFromString(*av.N)
	if err != nil {
		return err
	}
	return nil
}

func (this *Decimal) MarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	n := fmt.Sprintf("%s", this)
	av.N = &n
	return nil
}
