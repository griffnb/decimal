package decimal

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

/* AWS v1
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
*/
//AWS V2
func (this *Decimal) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	avN, ok := av.(*types.AttributeValueMemberN)
	if !ok {
		return fmt.Errorf("expected AttributeValueMemberN, got %T", av)
	}
	var err error
	*this, err = NewFromString(avN.Value)
	if err != nil {
		return err
	}
	return nil
}

func (this *Decimal) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	n := fmt.Sprintf("%s", this)
	return &types.AttributeValueMemberN{Value: n}, nil
}
