package main

func main() {
	// intro()

	// Подготавливаю мапу для заливки
	tos.ServersList = make(map[string]LineOfStatusTableType)

	//Подготавливаю мапу для заливки конфига
	servers.ServersList = make(map[string]ServersAttr)
	// Загружаю конфиг
	servers = getConf("./config.json")

	// Переношу загруженный конфиг во временную мапу для отображения
	tos.fillShapku(servers.ServersList)
	// tos.AddHeader()

	go servers.checkLoop()
	runWeb()
	// go runSpa()
	//runSpa()
}
