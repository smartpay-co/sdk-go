package integration

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/smartpay-co/sdk-go"
	"time"
)

func (suite *IntegrationTestSuite) TestCreateACoupon() {
	//suite.T().Skip()
	payload := CreateACouponJSONRequestBody{
		Currency:       Ptr(CurrencyJPY),
		Name:           fmt.Sprintf("%s %d", "sdk-go-test", time.Now().Unix()),
		DiscountType:   DiscountTypeAmount,
		DiscountAmount: Ptr(int32(50)),
	}
	result, err := suite.client.CreateACouponWithResponse(suite.ctx, payload)

	suite.Nil(err)
	suite.NotNil(result.Body)
	suite.NotNil(result.JSON200)
	suite.NotNil(result.JSON200.Id)
	couponId := string(*result.JSON200.Id)

	suite.Run("TestRetrieveACoupon", func() {
		result, err := suite.client.RetrieveACouponWithResponse(suite.ctx, couponId)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		suite.EqualValues(string(*result.JSON200.Id), couponId)
	})

	suite.Run("TestUpdateACoupon", func() {
		payload := UpdateACouponJSONRequestBody{
			Metadata: Ptr(Metadata{
				"testKey1": "testValue1",
				"testKey2": "testValue2",
			}),
		}
		result, err := suite.client.UpdateACouponWithResponse(suite.ctx, couponId, payload)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		suite.EqualValues(string(*result.JSON200.Id), couponId)
	})

	suite.Run("TestListAllCoupons", func() {
		params := ListAllCouponsParams{}
		result, err := suite.client.ListAllCouponsWithResponse(suite.ctx, &params)
		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		obj, _ := ConvertToStruct[Coupon]((*result.JSON200.Data)[0])
		suite.NotNil(obj.Id)
	})

	suite.Run("TestCreateAPromotionCode", func() {
		payload := CreateAPromotionCodeJSONRequestBody{
			Coupon:   CouponId(couponId),
			Currency: Ptr(CurrencyJPY),
			Code:     uuid.NewString(),
		}
		result, err := suite.client.CreateAPromotionCodeWithResponse(suite.ctx, payload)

		suite.Nil(err)
		suite.NotNil(result.Body)
		suite.NotNil(result.JSON200)
		suite.NotNil(result.JSON200.Id)
		promotionCodeId := string(*result.JSON200.Id)

		suite.Run("TestRetrieveAPromotionCode", func() {
			result, err := suite.client.RetrieveAPromotionCodeWithResponse(suite.ctx, promotionCodeId)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			suite.EqualValues(string(*result.JSON200.Id), promotionCodeId)
		})

		suite.Run("TestUpdateAPromotionCode", func() {
			payload := UpdateAPromotionCodeJSONRequestBody{
				Metadata: Ptr(Metadata{
					"testKey1": "testValue1",
					"testKey2": "testValue2",
				}),
			}
			result, err := suite.client.UpdateAPromotionCodeWithResponse(suite.ctx, promotionCodeId, payload)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			suite.EqualValues(string(*result.JSON200.Id), promotionCodeId)
		})

		suite.Run("TestListAllPromotionCodes", func() {
			params := ListAllPromotionCodesParams{}
			result, err := suite.client.ListAllPromotionCodesWithResponse(suite.ctx, &params)
			suite.Nil(err)
			suite.NotNil(result.Body)
			suite.NotNil(result.JSON200)
			obj, _ := ConvertToStruct[PromotionCode]((*result.JSON200.Data)[0])
			suite.NotNil(obj.Id)
		})
	})
}
