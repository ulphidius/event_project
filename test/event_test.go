package test

import (
	"bytes"
	"encoding/json"
	"event_project/infrastructure/http/route"
	"testing"

	_ "event_project/domain/event"
	_ "event_project/infrastructure/http/validation"

	_ "goyave.dev/goyave/v3/database/dialect/mysql"
	"goyave.dev/goyave/v3/validation"

	"goyave.dev/goyave/v3"
)

type EventTestSuite struct {
	goyave.TestSuite
}

func (suite *EventTestSuite) SetupTest() {
	suite.ClearDatabase()
}

func (suite *EventTestSuite) TestCreateEvent() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"type": "Click", "timestamp": "1577836800"})
		resp, err := suite.Post("/event", headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(201, resp.StatusCode)
		}
	})
}

func (suite *EventTestSuite) TestCreateEventUnknowType() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"type": "sample", "timestamp": "1577836800"})
		resp, err := suite.Post("/event", headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			json := map[string]validation.Errors{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				errorContent, err := json["validationError"]["type"]
				suite.True(err)
				suite.Equal(422, resp.StatusCode)
				suite.Equal("The type format is invalid. An event should be equal to Impression, Click or Visible", errorContent[0])
			}
		}
	})
}

func (suite *EventTestSuite) TestCreateEventTimestampLowerThanZero() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"type": "Click", "timestamp": "-1"})
		resp, err := suite.Post("/event", headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			json := map[string]validation.Errors{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				errorContent, err := json["validationError"]["timestamp"]
				suite.True(err)
				suite.Equal(422, resp.StatusCode)
				suite.Equal("The timestamp must be at least 0.", errorContent[0])
			}
		}
	})
}

func (suite *EventTestSuite) TestCreateEventTimestampNotANumber() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"type": "Click", "timestamp": "sample"})
		resp, err := suite.Post("/event", headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			json := map[string]validation.Errors{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				errorContent, err := json["validationError"]["timestamp"]
				suite.True(err)
				suite.Equal(422, resp.StatusCode)
				suite.Equal("The timestamp must be an integer.", errorContent[0])
			}
		}
	})
}

func (suite *EventTestSuite) TestCreateEventTypeRequired() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"timestamp": "sample"})
		resp, err := suite.Post("/event", headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			json := map[string]validation.Errors{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				errorContent, err := json["validationError"]["type"]
				expected := []string{
					"The type is required.",
					"The type must be a string.",
					"The type format is invalid. An event should be equal to Impression, Click or Visible",
				}
				suite.True(err)
				suite.Equal(422, resp.StatusCode)
				suite.Equal(expected, errorContent)
			}
		}
	})
}

func (suite *EventTestSuite) TestCreateEventTimestampRequired() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		body, _ := json.Marshal(map[string]interface{}{"type": "Click"})
		resp, err := suite.Post("/event", headers, bytes.NewReader(body))

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			json := map[string]validation.Errors{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				errorContent, err := json["validationError"]["timestamp"]
				expected := []string{
					"The timestamp is required.",
					"The timestamp must be an integer.",
				}
				suite.True(err)
				suite.Equal(422, resp.StatusCode)
				suite.Equal(expected, errorContent)
			}
		}
	})
}

func TestEventSuite(t *testing.T) {
	goyave.RunTest(t, new(EventTestSuite))
}
