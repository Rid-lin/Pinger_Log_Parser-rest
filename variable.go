package main

//ServersAttr - attribute and status of server
type ServersAttr struct {
	IP     string `json:"IP"`
	Note   string `json:"Note"`
	SiteID string `json:"SiteID"`
}

// Configuration type for list of servers for check
type Configuration struct {
	logLevel      int
	FileLog       string
	NumberOfCheck int
	TimeOutSleep  int
	ServersList   map[string]ServersAttr
}

//Servers - list of servers for check
var servers Configuration

var tos tableOfStatusType

// StatusType type for status servers
type StatusType struct {
	Code       string
	Background string
	NumPass    int
	NumFail    int
}

//StatusNowType type for dysplay status last check
type StatusNowType struct {
	Code string
}

//StatusOfHourType type for calculate of check
type StatusOfHourType struct {
	NumPass int
	NumFail int
}

//LineOfStatusTableType type for display of table
type LineOfStatusTableType struct {
	IP           string
	Note         string
	SiteID       string
	StatusNow    StatusNowType
	StatusOfHour [24]StatusOfHourType
}

type tableOfStatusType struct {
	ServersList map[string]LineOfStatusTableType
}
