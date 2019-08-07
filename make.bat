@Echo Off
Set oldipforsite=%~dp0
For /D %%a In ("%oldipforsite:~0,-1%.txt") Do Set oldipforsite=%%~na
go build --ldflags "-w -s" -o %oldipforsite%.exe
D:\apps\upx394w\upx.exe --ultra-brute .\%oldipforsite%.exe