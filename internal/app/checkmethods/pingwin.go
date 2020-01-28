package checkmethods

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func pingWin(ip string) string {
	var pingTime string

	// tm := time.Now().Format("2006/01/02 15:04:05")
	tm := time.Now().Format("15:04:05")
	out, err := exec.Command("c:\\Windows\\System32\\ping.exe ", ip).Output()
	if err != nil {
		pingTime = "down"
	} else {
		// stringDecodeOut := cp866ToUTF8(out)
		stringDecodeOut, _, err := transform.String(charmap.CodePage866.NewDecoder(), string(out))
		if err != nil {
			stringDecodeOut = ""
		}

		pingTime = parseStringToTime(stringDecodeOut)
	}

	result := fmt.Sprintf("%s,%s,%s\n", tm, ip, pingTime)

	return result
}

// func cp866ToUTF8(out []byte) string {
// 	//Инициализируем декодирование с указанным типом CodePage866
// 	d := charmap.CodePage866.NewDecoder()
// 	//Обрабатываем вывод
// 	decodeOut, _ := d.Bytes(out)
// 	//Возвращаем обработанный ответ
// 	stringDecodeOut := string(decodeOut)

// 	return stringDecodeOut
// }

func parseStringToTime(stringDecodeOut string) string {
	// Русская вверсия винды
	if strings.Contains(stringDecodeOut, "Среднее") {
		indexStr := strings.Index(stringDecodeOut, "Среднее")
		pingTime := stringDecodeOut[indexStr+17 : len(stringDecodeOut)-11]

		return pingTime

		// Английская версия винды
	} else if strings.Contains(stringDecodeOut, "Average") {
		indexStr := strings.Index(stringDecodeOut, "Average")
		pingTime := stringDecodeOut[indexStr+10 : len(stringDecodeOut)-4]

		return pingTime
	}
	return "down"

}
