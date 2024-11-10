@echo off

set /a counter=0

for %%f in (*.jpg) do (
	C:\Users\MSI\Downloads\libwebp-1.4.0-windows-x64\libwebp-1.4.0-windows-x64\bin\cwebp.exe -q 100 %%f -o %%~nf.webp
	set /A counter=counter+1
)

echo ************************************
echo %counter% files have been processed.
echo ************************************