package checkmethods

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
)

func pingWin(ip string) string {
	var pingTime string

	// tm := time.Now().Format("2006/01/02 15:04:05")
	tm := time.Now().Format("15:04:05")
	out, err := exec.Command("c:\\Windows\\System32\\ping.exe ", ip).Output()
	if err != nil {
		pingTime = "down"
	} else {
		stringDecodeOut := cp866ToUTF8(out)
		pingTime = parseStringToTime(stringDecodeOut)
	}

	result := fmt.Sprintf("%s,%s,%s", tm, ip, pingTime)

	return result
}

func cp866ToUTF8(out []byte) string {
	//Инициализируем декодирование с указанным типом CodePage866
	d := charmap.CodePage866.NewDecoder()
	//Обрабатываем вывод
	decodeOut, _ := d.Bytes(out)
	//Возвращаем обработанный ответ
	stringDecodeOut := string(decodeOut)

	return stringDecodeOut
}

func parseStringToTime(stringDecodeOut string) string {
	fmt.Println(stringDecodeOut)

	// Русская вверсия винды
	if strings.Contains(stringDecodeOut, "Среднее") {
		indexStr := strings.Index(stringDecodeOut, "Среднее")
		pingTime := stringDecodeOut[indexStr+17 : len(stringDecodeOut)-10]

		return pingTime

		// Английская версия винды
	} else if strings.Contains(stringDecodeOut, "Average") {
		indexStr := strings.Index(stringDecodeOut, "Average")
		pingTime := stringDecodeOut[indexStr+10 : len(stringDecodeOut)-4]

		return pingTime
	}
	return "down"

}
