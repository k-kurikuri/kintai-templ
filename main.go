package main

import (
	"fmt"
	"log"
	"time"
)

// layout is dateFormat
const layout = "2006-01-02 15:04"

// main
func main()  {
	now := time.Now().Format(layout)
	t, err := time.Parse(layout, now)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	printContent(t)
}

// printContent
func printContent(t time.Time)  {
	weeks := []string{"日", "月", "火", "水", "木", "金", "土"}
	strFormat := `%02d/%02d(%s) %02d:%02d - xx:xx`

	content := fmt.Sprintf(strFormat, int(t.Month()),t.Day(),weeks[int(t.Weekday())], t.Hour(), t.Minute())

	fmt.Println("```")
	fmt.Println("[出退勤]")
	fmt.Println(content)
	fmt.Print("```")
}
