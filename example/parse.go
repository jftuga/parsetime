package main

import (
	"fmt"
	"github.com/crackcomm/go-clitable"
	"github.com/jftuga/parsetime"
	"log"
)

var iso8601Times = []string{
	"2006-01-02 15:04",
	"2006-01-02 15:04-07:00",
	"2006-01-02 15:04 -07:00",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05-07:00",
	"2006-01-02 15:04:05 -07:00",
	"2006-01-02 15:04:05-07:00 MST",
	"2006-01-02 15:04:05 -07:00 MST",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02 15:04:05.999999-07:00 MST",
	"2006-01-02 15:04:05.9-07:00 MST",
	"2006-01-02 15:04:05.9 -07:00 MST",
	"2006-01-02 15:04:05.999-07:00 MST",
	"2006-01-02 15:04:05.999 -07:00 MST",
	"2006-01-02 15:04:05.999999-07:00 MST",
	"2006-01-02 15:04:05.999999 -07:00 MST",
	"2006-01-02 15:04:05.999999999-07:00 MST",
	"2006-01-02 15:04:05.999999999 -07:00 MST",
	"2006-01-02T15:04",
	"2006-01-02T15:04-07:00",
	"2006-01-02T15:04 -07:00",
	"2006-01-02T15:04:05",
	"2006-01-02T15:04:05-07:00",
	"2006-01-02T15:04:05 -07:00",
	"2006-01-02T15:04:05-07:00 MST",
	"2006-01-02T15:04:05 -07:00 MST",
	"2006-01-02T15:04:05.999999999",
	"2006-01-02T15:04:05.999999999-07:00 MST",
	"2006-01-02T15:04:05.999999999 -07:00 MST",
	"2006-01-02T15:04:05.999999-07:00 MST",
	"2006-01-02T15:04:05.999999 -07:00 MST",
	"2006-01-02T15:04:05.9-07:00 MST",
	"2006-01-02T15:04:05.9 -07:00 MST",
	"2006-01-02",
	"20060102",
	"20060102150405",
	"20060102 150405",
	"20060102T150405",
	"15:04:05",
	"15:04:05-07:00 MST",
	"15:04:05 -07:00 MST",
	"15:04:05.9-07:00 MST",
	"15:04:05.9 -07:00 MST",
	"15:04:05.999-07:00 MST",
	"15:04:05.999 -07:00 MST",
	"15:04:05.999999-07:00 MST",
	"15:04:05.999999 -07:00 MST",
	"15:04:05.999999999-07:00 MST",
	"15:04:05.999999999 -07:00 MST",
	"150405-07:00 MST",
	"150405 -07:00 MST",
	"150405.9-07:00 MST",
	"150405.9 -07:00 MST",
	"150405.999-07:00 MST",
	"150405.999 -07:00 MST",
	"150405.999999-07:00 MST",
	"150405.999999 -07:00 MST",
	"150405.999999999-07:00 MST",
	"150405.999999999 -07:00 MST",
	"2006-01-02 15:04:05Z",
	"2006-01-02T15:04:05Z",
	"2006-01-02 15:04:05.9Z",
	"2006-01-02T15:04:05.9Z",
	"2006-01-02 15:04:05.999Z",
	"2006-01-02T15:04:05.999Z",
	"2006-01-02 15:04:05.999999Z",
	"2006-01-02T15:04:05.999999Z",
	"2006-01-02 15:04:05.999999999Z",
	"2006-01-02T15:04:05.999999999Z",
}

var rfc8xx1123Times = []string{
	"02-Jan-06 1504 MST",
	"02-Jan-06 15:04 MST",
	"02-Jan-06 150405 MST",
	"02-Jan-06 15:04:05 MST",
	"02-Jan-06 1504-0700",
	"02-Jan-06 15:04-0700",
	"02-Jan-06 150405-0700",
	"02-Jan-06 15:04:05-0700",
	"02-Jan-06 15:04 -0700",
	"02-Jan-06 15:04:05 -0700",
	"Monday, 02-Jan-06 15:04 MST",
	"Monday, 02-Jan-06 15:04:05 MST",
	"Mon, 02-Jan-06 15:04 MST",
	"Mon, 02-Jan-06 15:04:05 MST",
	"Mon, 02-Jan-06 15:04-07:00",
	"Mon, 02-Jan-06 15:04:05-07:00",
	"Mon, 02-Jan-06 15:04 -07:00",
	"Mon, 02-Jan-06 15:04:05 -07:00",
	"Mon, 02-Jan-2006 15:04-07:00",
	"Mon, 02-Jan-2006 15:04:05-07:00",
	"Mon, 02-Jan-2006 15:04 -07:00",
	"Mon, 02-Jan-2006 15:04:05 -07:00",
	"Mon, 02-Jan-70 15:04-07:00",
	"Mon, 02-Jan-70 15:04:05-07:00",
	"Mon, 02-Jan-70 15:04 -07:00",
	"Mon, 02-Jan-70 15:04:05 -07:00",
	"Mon, 02-Jan-99 15:04-07:00",
	"Mon, 02-Jan-99 15:04:05-07:00",
	"Mon, 02-Jan-99 15:04:05 -07:00",
	"Mon, 02-Jan-00 15:04-07:00",
	"Mon, 02-Jan-00 15:04:05-07:00",
	"Mon, 02-Jan-00 15:04:05 -07:00",
	"Mon, 02-Jan-00 15:04:05.9-07:00",
	"Mon, 02-Jan-00 15:04:05.9 -07:00",
	"Mon, 02-Jan-00 15:04:05.999-07:00",
	"Mon, 02-Jan-00 15:04:05.999 -07:00",
	"Mon, 02-Jan-00 15:04:05.999999-07:00",
	"Mon, 02-Jan-00 15:04:05.999999 -07:00",
	"Mon, 02-Jan-00 15:04:05.999999999-07:00",
	"Mon, 02-Jan-00 15:04:05.999999999 -07:00",
}

var ansicTimes = []string{
	"Mon Jan 02 150405 2006",
	"Mon Jan 02 15:04:05 2006",
	"Mon Jan 02 150405 MST 2006",
	"Mon Jan 02 15:04:05 MST 2006",
	"Mon Jan 02 1504-07:00 2006",
	"Mon Jan 02 15:04-07:00 2006",
	"Mon Jan 02 1504 -07:00 2006",
	"Mon Jan 02 15:04 -07:00 2006",
	"Mon Jan 02 150405-07:00 2006",
	"Mon Jan 02 15:04:05-07:00 2006",
	"Mon Jan 02 150405 -07:00 2006",
	"Mon Jan 02 15:04:05 -07:00 2006",
	"Jan 02 150405",
	"Jan 02 15:04:05",
	"Jan 02 150405.9",
	"Jan 02 15:04:05.9",
	"Jan 02 150405.999",
	"Jan 02 15:04:05.999",
	"Jan 02 150405.999999",
	"Jan 02 15:04:05.999999",
	"Jan 02 150405.999999999",
	"Jan 02 15:04:05.999999999",
}

var usTimes = []string{
	"11:04AM",
	"11:04PM",
	"11:04 AM",
	"11:04 PM",
	"11:04:05 AM",
	"11:04:05 PM",
	"11:04:05.9AM",
	"11:04:05.9 AM",
	"11:04:05.9PM",
	"11:04:05.9 PM",
	"11:04:05.999AM",
	"11:04:05.999 AM",
	"11:04:05.999PM",
	"11:04:05.999 PM",
	"11:04:05.999999AM",
	"11:04:05.999999 AM",
	"11:04:05.999999PM",
	"11:04:05.999999 PM",
	"11:04:05.999999999AM",
	"11:04:05.999999999 AM",
	"11:04:05.999999999PM",
	"11:04:05.999999999 PM",
	"01-02-06 3:04AM",
	"01-02-06 3:04 AM",
	"01-02-06 3:04PM",
	"01-02-06 3:04 PM",
	"01-02-06 03:04:05AM",
	"01-02-06 03:04:05 AM",
	"01-02-06 03:04:05PM",
	"01-02-06 03:04:05 PM",
	"01-02-06 03:04:05.9AM",
	"01-02-06 03:04:05.9 AM",
	"01-02-06 03:04:05.9PM",
	"01-02-06 03:04:05.9 PM",
	"01-02-06 03:04:05.999AM",
	"01-02-06 03:04:05.999 AM",
	"01-02-06 03:04:05.999PM",
	"01-02-06 03:04:05.999 PM",
	"01-02-06 03:04:05.999999AM",
	"01-02-06 03:04:05.999999 AM",
	"01-02-06 03:04:05.999999PM",
	"01-02-06 03:04:05.999999 PM",
	"01-02-06 03:04:05.999999999AM",
	"01-02-06 03:04:05.999999999 AM",
	"01-02-06 03:04:05.999999999PM",
	"01-02-06 03:04:05.999999999 PM",
	"Jan 2, 2006",
	"Jan 2, 2006 at 3:04am (MST)",
	"Jan 2, 2006 at 03:04am (MST)",
	"Jan 2, 2006 at 3:04pm (MST)",
	"Jan 2, 2006 at 03:04pm (MST)",
	"Jan 2, 2006 at 3:04 am (MST)",
	"Jan 2, 2006 at 03:04 am (MST)",
	"Jan 2, 2006 at 3:04 pm (MST)",
	"Jan 2, 2006 at 03:04 pm (MST)",
	"Jan 2, 2006 at 3:04:05am (MST)",
	"Jan 2, 2006 at 3:04:05pm (MST)",
	"Jan 2, 2006 at 3:04:05 am (MST)",
	"Jan 2, 2006 at 3:04:05 pm (MST)",
	"Jan 2, 2006 at 3:04:05.9am (MST)",
	"Jan 2, 2006 at 3:04:05.9pm (MST)",
	"Jan 2, 2006 at 3:04:05.999am (MST)",
	"Jan 2, 2006 at 3:04:05.999pm (MST)",
	"Jan 2, 2006 at 3:04:05.999999am (MST)",
	"Jan 2, 2006 at 3:04:05.999999pm (MST)",
	"Jan 2, 2006 at 3:04:05.999999999am (MST)",
	"Jan 2, 2006 at 3:04:05.999999999pm (MST)",
	"Jan 2, 2006 at 3:04am MST",
	"Jan 2, 2006 at 3:04pm MST",
	"Jan 2, 2006 at 3:04am -07:00",
	"Jan 2, 2006 at 3:04pm -07:00",
	"Jan 2, 2006 at 3:04:05am MST",
	"Jan 2, 2006 at 3:04:05pm MST",
	"Jan 2, 2006 at 3:04:05am -07:00",
	"Jan 2, 2006 at 3:04:05pm -07:00",
}

func printTable(p parsetime.ParseTime, header string, times []string) {
	fmt.Printf("### %s", header)
	fmt.Println("")
	fmt.Println("")

	table := clitable.New([]string{"Input string", "_time.Time"})

	for _, v := range times {
		t, err := p.Parse(v)
		if err != nil {
			log.Printf("%s : %s", err, t.String())
			fmt.Println()
		}

		table.AddRow(map[string]interface{}{"Input string": v, "_time.Time": t.String()})
	}

	table.Markdown = true

	table.Print()
}

func mergeSlice(times ...[]string) []string {
	var mergedTimes []string
	mergedTimes = make([]string, 0)

	for _, t := range times {
		mergedTimes = append(mergedTimes, t...)
	}

	return mergedTimes
}

func main() {
	p, _ := parsetime.NewParseTime()

	printTable(p, "ISO8601", iso8601Times)
	fmt.Println("")
	printTable(p, "RFC8xx1123", rfc8xx1123Times)
	fmt.Println("")
	printTable(p, "ANSIC", ansicTimes)
	fmt.Println("")
	printTable(p, "US", usTimes)
	fmt.Println("")

	parseTimes := mergeSlice(iso8601Times, rfc8xx1123Times, ansicTimes, usTimes)
	printTable(p, "Parse", parseTimes)
	fmt.Println("")
}
