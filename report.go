// main project report.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func (s *tableOfStatusType) getreportHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("выполняется getreport")
	outputFileName, err := s.makereport()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	templateString := templHeader + templReport + templFooter

	tmpl, err := template.New("report").Parse(templateString)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	tmpl.Execute(w, outputFileName)

}

func (s *tableOfStatusType) makereport() (string, error) {
	// fmt.Println("выполняется makereport")
	wSheet := "Отчет"
	templateFile := "./Template.xlsx"
	index := 1
	outputFileName := time.Now().Format("2006-01-02") + ".xlsx"
	xlsx, err := excelize.OpenFile(templateFile)
	if err != nil {
		fmt.Println(err)
		return outputFileName, err
	}
	for key, value := range s.ServersList {
		xlsx.SetCellValue(wSheet, "B"+strconv.Itoa(index+6), index)
		xlsx.SetCellValue(wSheet, "C"+strconv.Itoa(index+6), value.Note)
		xlsx.SetCellValue(wSheet, "E"+strconv.Itoa(index+6), key)
		xlsx.SetCellValue(wSheet, "F"+strconv.Itoa(index+6), value.SiteID)
		if value.StatusNow.Code == "√" {
			xlsx.SetCellValue(wSheet, "D"+strconv.Itoa(index+6), "В сети")
		} else {
			xlsx.SetCellValue(wSheet, "D"+strconv.Itoa(index+6), "Не в сети")
		}
		index++
	}
	xlsx.SetCellValue(wSheet, "D3", time.Now().Format("02.01.2006"))
	// Save xlsx file by the given path.
	err = xlsx.SaveAs("./report/" + outputFileName)
	if err != nil {
		fmt.Println(err)
	}
	return outputFileName, err
}
