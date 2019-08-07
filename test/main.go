package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// runPinger("10.61.133.146")
	// runPinger("10.61.133.144")
	ips := []string{"10.61.133.146", "10.61.133.144"}
	runPingers(ips)
}

func runPinger(ip string) {
	var str string
	str += fmt.Sprintf("'%v'", ip)
	strRun := "$ns=" + str + `;$O= @();foreach ($n in $ns){if (Test-Connection -ComputerName $n -Count 1 -ErrorAction SilentlyContinue){$s="up" }else{ $s="down"};$t = Get-Date -Format "HH:mm:ss";$O += "$t,$n,$s";Write-Host = "$t,$n,$s";};$d = Get-Date -Format "yyyy_MM_dd";$f = "./logs/$d.csv";Add-Content $f $O;`
	fmt.Printf("str=%v\n%v\v", str, strRun)
	_ = exec.Command("powershell", strRun).Run()
}

func runPingers(ips []string) {
	var str string
	for i, ip := range ips {
		str += fmt.Sprintf("'%v'", ip)
		if i < len(ips)-1 {
			str += ","
		}
	}
	strRun := "$ns=" + str + `;$O= @();foreach ($n in $ns){if (Test-Connection -ComputerName $n -Count 1 -ErrorAction SilentlyContinue){$s="up" }else{ $s="down"};$t = Get-Date -Format "HH:mm:ss";$O += "$t,$n,$s";Write-Host = "$t,$n,$s";};$d = Get-Date -Format "yyyy_MM_dd";$f = "./logs/$d.csv";Add-Content $f $O;`
	fmt.Printf("str=%v\n%v\v", str, strRun)
	_ = exec.Command("powershell", strRun).Run()
}
