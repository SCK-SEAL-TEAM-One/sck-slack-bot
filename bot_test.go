package slackbot

import (
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func Test_YearMonthDayCounter_Input_20130307_20200413_Should_Be_7_Years_1_Months_7_days(t *testing.T) {
	expected := `7 ปี 1 เดือน 7 วัน`
	endOfYearDate, _ := time.Parse("20060102", "20200412")
	launchDate, _ := time.Parse("20060102", "20130307")

	actual := yearMonthDayCounter(launchDate, endOfYearDate)

	assert.Equal(t, expected, actual)
}

func Test_YearMonthDayCounter_Input_20130307_Should_Be_6_Years_9_Months_25_days(t *testing.T) {
	expected := `6 ปี 9 เดือน 26 วัน`
	endOfYearDate, _ := time.Parse("20060102", "20191231")
	launchDate, _ := time.Parse("20060102", "20130307")

	actual := yearMonthDayCounter(launchDate, endOfYearDate)

	assert.Equal(t, expected, actual)
}
func Test_YearMonthDayCounter_Input_20180104_Should_Be_1_Years_11_Months_28_days(t *testing.T) {
	expected := `1 ปี 11 เดือน 29 วัน`
	endOfYearDate, _ := time.Parse("20060102", "20191231")
	launchDate, _ := time.Parse("20060102", "20180104")

	actual := yearMonthDayCounter(launchDate, endOfYearDate)

	assert.Equal(t, expected, actual)
}

func Test_YearMonthDayCounter_Input_20130307_Should_Be_1_Years_7_Months_11_days(t *testing.T) {
	expected := `1 ปี 6 เดือน 12 วัน`
	endOfYearDate, _ := time.Parse("20060102", "20191231")
	launchDate, _ := time.Parse("20060102", "20180621")

	actual := yearMonthDayCounter(launchDate, endOfYearDate)

	assert.Equal(t, expected, actual)
}

func Test_YearMonthDayCounter_Input_20130307_Should_Be_7_Months_11_days(t *testing.T) {
	expected := `6 เดือน 12 วัน`
	endOfYearDate, _ := time.Parse("20060102", "20191231")
	launchDate, _ := time.Parse("20060102", "20190621")

	actual := yearMonthDayCounter(launchDate, endOfYearDate)

	assert.Equal(t, expected, actual)
}

func Test_dateMessage_Input_20191231_Should_Be_Tuesday_31_December_2562(t *testing.T) {
	expected := `วันอังคารที่ 31 ธันวาคม พ.ศ. 2562`
	endOfYearDate, _ := time.Parse("20060102", "20191231")

	actual := dateMessage(endOfYearDate)

	assert.Equal(t, expected, actual)
}

func Test_message_Input_20130307_20200413_Should_Be_7_Years_1_Months_7_days(t *testing.T) {
	expected := `สยามชำนาญกิจ วันพฤหัสที่ 7 มีนาคม พ.ศ. 2556 6 ปี 9 เดือน 26 วัน`
	endOfYearDate, _ := time.Parse("20060102", "20191231")

	company := Company{
		Title: "สยามชำนาญกิจ",
		Date:  "2013/03/07",
	}
	actual := message(company, endOfYearDate)

	assert.Equal(t, expected, actual)
}
