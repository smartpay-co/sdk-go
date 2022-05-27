package integration

import (
	"context"
	. "github.com/smartpay-co/sdk-go"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type OrderTestSuite struct {
	client *ClientWithResponses
	suite.Suite
	ctx context.Context
}

func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}

func (suite *OrderTestSuite) SetupSuite() {
	secretKey := os.Getenv("SMARTPAY_SECRET_KEY")
	publicKey := os.Getenv("SMARTPAY_PUBLIC_KEY")
	apiBase := os.Getenv("API_BASE")

	suite.ctx = context.TODO()

	var err error
	suite.client, err = NewClientWithResponses(secretKey, publicKey, WithBaseURL(apiBase))
	if err != nil {
		panic(err)
	}
}

func (suite *OrderTestSuite) TestListAllOrders() {
	//suite.T().Skip()
	params := ListAllOrdersParams{}

	result, err := suite.client.ListAllOrdersWithResponse(suite.ctx, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	order, _ := ConvertToStruct[Order]((*result.JSON200.Data)[0])
	suite.NotNil(order.Id)
}

func (suite *OrderTestSuite) TestListAllOrdersExpanded() {
	//suite.T().Skip()
	params := ListAllOrdersParams{
		Expand: Ptr(ListAllOrdersParamsExpand(ExpandAll)),
	}

	result, err := suite.client.ListAllOrdersWithResponse(suite.ctx, &params)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)

	order, _ := ConvertToStruct[Order]((*result.JSON200.Data)[0])
	suite.NotNil(order.Id)
}
