package test

import (
	_ "event_project/domain/event"
	"event_project/infrastructure/http/route"
	_ "event_project/infrastructure/http/validation"
	"testing"

	"goyave.dev/goyave/v3"
	_ "goyave.dev/goyave/v3/database/dialect/mysql"
)

type StatisticsTestSuite struct {
	goyave.TestSuite
}

func (suite *StatisticsTestSuite) SetupTest() {
	suite.ClearDatabase()
}

func (suite *StatisticsTestSuite) TestOsStatistics() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&min=1420070400&max=1577836800", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "{\"number_of_linux\":0,\"number_of_mac\":0,\"number_of_windows\":0,\"number_of_phone\":0,\"other\":0}\n"
			suite.Equal(200, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestTypeStatistics() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=type&min=1420070400&max=1577836800", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			suite.Equal("{\"number_of_impression\":0,\"number_of_click\":0,\"number_of_visible\":0}\n", string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsIllType() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=sample&min=1420070400&max=1577836800", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"For type argument only the values 'os' and 'type' are allowed\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMinNotANumber() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&min=sample&max=1577836800", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"Min must be integer\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMaxNotANumber() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&min=1420070400&max=sample", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"Max must be integer\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMinUpperThanMax() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&min=1577836800&max=1420070400", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"Min must be lower or equal to max\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMinLowerThanZero() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&min=-1&max=1577836800", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"Min and max must be upper than 0\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMaxLowerThanZero() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&min=-1&max=1577836800", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"Min and max must be upper than 0\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMissingMin() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&max=1577836800", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"The arguments type, min and max are required\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMissingMax() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?type=os&min=1577836800&", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"The arguments type, min and max are required\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func (suite *StatisticsTestSuite) TestStatisticsMissingType() {
	suite.RunServer(route.Register, func() {
		resp, err := suite.Get("/statistics?min=1577836800&max=1420070400", nil)

		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			expected := "\"The arguments type, min and max are required\"\n"

			suite.Equal(422, resp.StatusCode)
			suite.Equal(expected, string(suite.GetBody(resp)))
		}
	})
}

func TestStatisticsSuite(t *testing.T) {
	goyave.RunTest(t, new(StatisticsTestSuite))
}
