package slackbot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Message struct {
	ResponseType string       `json:"response_type"`
	Text         string       `json:"text"`
	Attachments  []attachment `json:"attachments"`
}

type attachment struct {
	Color     string `json:"color"`
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
	ImageURL  string `json:"image_url"`
}

func DayCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST requests are accepted", 405)
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Couldn't parse form", 400)
		log.Fatalf("ParseForm: %v", err)
	}
	if err := verifyWebHook(r.Form); err != nil {
		log.Fatalf("verifyWebhook: %v", err)
	}

	message := formatMessage()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		log.Fatalf("json.Marshal: %v", err)
	}
}

func verifyWebHook(form url.Values) error {
	t := form.Get("token")
	if len(t) == 0 {
		return fmt.Errorf("empty form token")
	}
	if t != os.Getenv("token") {
		return fmt.Errorf("invalid request/credentials: %q", t[0])
	}
	return nil
}

type Company struct {
	Title string `yaml:"title"`
	Date  string `yaml:"date"`
}

func (company Company) GetDate() time.Time {
	date, _ := time.Parse("2006/01/02", company.Date)
	return date
}

func loadList() []Company {
	var companies []Company
	launches := strings.Split(os.Getenv("launch"), ",")
	for _, launch := range launches {
		details := strings.Split(launch, "=>")
		companies = append(companies, Company{
			Title: details[0],
			Date:  details[1],
		})
	}
	log.Printf("loadList %v\n", companies)
	return companies
}

func formatMessage() *Message {
	companies := loadList()
	attachments := make([]attachment, 0)
	for _, company := range companies {
		attachments = append(attachments, attachment{Color: "#d6334b", Text: message(company, time.Now())})
	}
	return &Message{
		ResponseType: "in_channel",
		Text:         dateMessage(time.Now()),
		Attachments:  attachments,
	}
}

func dateMessage(t time.Time) string {
	date := weekdays[t.Weekday()]
	month := monthTitle[t.Month()-1]
	year := t.Year() + 543
	return fmt.Sprintf("%sที่ %d %s พ.ศ. %d", date, t.Day(), month, year)
}

func yearMonthDayCounter(launchedDate, now time.Time) string {
	var output string
	launchedDate = launchedDate.AddDate(0, 0, -1)
	year := now.Year() - launchedDate.Year()
	month := int(now.Month() - launchedDate.Month())
	if int(month) < 0 {
		year--
		month += 12
	}
	day := now.Day() - launchedDate.Day()
	if day < 0 {
		month--
		day += 30
	}
	if year > 0 {
		output = fmt.Sprintf("%d ปี ", year)
	}
	if month > 0 {
		output += fmt.Sprintf("%d เดือน ", month)
	}
	if day > 0 {
		output += fmt.Sprintf("%d วัน", day)
	}
	return output
}

func message(company Company, t time.Time) string {
	return fmt.Sprintf("%s %s", company.Title, yearMonthDayCounter(company.GetDate(), t))
}
