// main project main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func (tos *tableOfStatusType) indexHandler(w http.ResponseWriter, r *http.Request) {
	templateString := templHeader + templIndex + templFooter
	tmpl, err := template.New("index").Parse(templateString)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	date := time.Now().Format("2006_01_02")
	tos.clearCache()

	tos.readFromLogs(date)
	tos.checkactualListIP(&servers)

	tmpl.Execute(w, tos.ServersList)

}

func (tos *tableOfStatusType) addHeadersHendler(w http.ResponseWriter, r *http.Request) {
	// tos.DelHeader()
	// tos.AddHeader()
	http.Redirect(w, r, "/", 302)
}

func (s *Configuration) editHandler(w http.ResponseWriter, r *http.Request) {
	templateString := templHeader + templWrite + templFooter

	tmpl, err := template.New("data").Parse(templateString)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	IP := r.FormValue("IP")
	serverElm, ok := s.ServersList[IP]
	if !ok {
		http.NotFound(w, r)
	}

	tmpl.Execute(w, serverElm)

}

func (s *Configuration) writeHandler(w http.ResponseWriter, r *http.Request) {
	writetmpl, err := template.ParseFiles("template/write.html", "template/header.html", "template/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	err = writetmpl.ExecuteTemplate(w, "write", nil)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func (s *Configuration) addserverHandler(w http.ResponseWriter, r *http.Request) {
	var serverElm ServersAttr
	serverElm.IP = r.FormValue("IP")
	serverElm.Note = r.FormValue("Note")
	serverElm.SiteID = r.FormValue("SiteID")
	runOncePing(serverElm.IP)
	s.ServersList[serverElm.IP] = serverElm
	tos.fillShapku(servers.ServersList)
	http.Redirect(w, r, "/", 302)
}

func (s *Configuration) deleteHandler(w http.ResponseWriter, r *http.Request) {
	IP := r.FormValue("IP")
	if IP == "" {
		http.NotFound(w, r)
	}
	delete(s.ServersList, IP)
	http.Redirect(w, r, "/", 302)
}

func (s *Configuration) getIPLists() []string {
	var slice []string
	for ip := range s.ServersList {
		slice = append(slice, ip)
	}
	return slice
}

func (s *Configuration) checkNowHandler(w http.ResponseWriter, r *http.Request) {

	runPinger(s.getIPLists())
	http.Redirect(w, r, "/", 302)
}

//ReLoadDefaultConfigHandler Reload Config servers
func (s *Configuration) ReLoadDefaultConfigHandler(w http.ResponseWriter, r *http.Request) {
	conf := getConf("./default_config.json")
	s = &conf
	s.checkNowHandler(w, r)

	toLog(servers.logLevel, 1, fmt.Sprintf("Ждём минут: %d\n", servers.TimeOutSleep))
}

//ReLoadConfigHandler Reload Config servers
func (s *Configuration) ReLoadConfigHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("Читаю конфигурационный файл config.json\n")
	conf := getConf("./default_config.json")
	s = &conf
	http.Redirect(w, r, "/", 302)
}

//SaveConfigHandler Save configuration in file JSON
func (s *Configuration) SaveConfigHandler(w http.ResponseWriter, r *http.Request) {
	err := backupConfig("./config.json")
	if err != nil {
		toLog(s.logLevel, 2, "Rename file failed: ", err)
	}

	saveConf("./config.json", s)
	http.Redirect(w, r, "/", 302)
}

func runWeb() {
	http.HandleFunc("/", tos.indexHandler)
	http.HandleFunc("/getreport", tos.getreportHandler)
	http.HandleFunc("/checknow", servers.checkNowHandler)
	http.HandleFunc("/write", servers.writeHandler)
	http.HandleFunc("/addserver", servers.addserverHandler)
	http.HandleFunc("/edit", servers.editHandler)
	http.HandleFunc("/delete", servers.deleteHandler)
	http.HandleFunc("/loaddefaultconf", servers.ReLoadDefaultConfigHandler)
	http.HandleFunc("/reloadconf", servers.ReLoadConfigHandler)
	http.HandleFunc("/saveconf", servers.SaveConfigHandler)
	http.Handle("/report/", http.StripPrefix("/report/", http.FileServer(http.Dir("./report"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	toLog(servers.logLevel, 1, "Запуск локального WEB-сервера на порту :8089")

	http.ListenAndServe(":8089", nil)
}
