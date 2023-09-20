package logger

import (
	"fmt"
	"os"
	"time"
)

func CreateLog(path, log string) {

	day := time.Now().Day()
	mounth := time.Now().Month()
	year := time.Now().Year()

	sdate := fmt.Sprintf("%d-%d-%d", day, mounth, year)

	file, err := os.OpenFile("logs/"+path+"/"+sdate+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error open log file", err)
		return
	}

	defer file.Close()

	_, err2 := file.WriteString(log)
	if err2 != nil {
		fmt.Println("Error added log on file", err2)
	}

}
