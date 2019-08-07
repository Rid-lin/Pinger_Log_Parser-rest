// main project ui.go
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func (tos *tableOfStatusType) fillShapku(source map[string]ServersAttr) {
	for IP, value := range source {
		var shapkaLine LineOfStatusTableType
		shapkaLine.IP = value.IP
		shapkaLine.Note = value.Note
		shapkaLine.SiteID = value.SiteID
		tos.ServersList[IP] = shapkaLine
	}
}

func (tos *tableOfStatusType) readFromLogs(date string) {
	filename := "./logs/" + date + ".csv"
	f, err := os.Open(filename)
	if err != nil {
		runPinger(servers.getIPLists())
		f, _ = os.Open(filename)
	}
	r := bufio.NewReader(f)
	s, _, _ := r.ReadLine()
	i := 1
	for len(s) > 0 {
		tos.parseLogLineCSV(s)
		s, _, _ = r.ReadLine()
		i++
	}
}

func (tos *tableOfStatusType) parseLogLineCSV(b []byte) {
	line := string(b)
	slice := strings.Split(line, ",")
	timeFromLine := slice[0]
	IP := slice[1]
	rttStr := slice[2]
	hourstr := strings.Split(timeFromLine, ":")[0]
	hour, _ := strconv.Atoi(hourstr)
	status := tos.ServersList[IP]
	if rttStr == "down" {
		status.StatusOfHour[hour].NumFail++
		status.StatusNow.Code = "X"
	} else {
		status.StatusOfHour[hour].NumPass++
		status.StatusNow.Code = "√"
	}
	tos.ServersList[IP] = status
}

func (tos *tableOfStatusType) clearCache() {
	for IP, line := range tos.ServersList {
		if IP == "IP адрес" {
			continue
		}
		delete(tos.ServersList, IP)
		var status LineOfStatusTableType
		status.IP = line.IP
		status.Note = line.Note
		status.SiteID = line.SiteID
		tos.ServersList[IP] = status
	}

}

func (tos *tableOfStatusType) checkactualListIP(servers *Configuration) {
	for IP := range tos.ServersList {
		if _, ok := servers.ServersList[IP]; !ok && IP != "IP адрес" {
			delete(tos.ServersList, IP)
		}
	}
}
