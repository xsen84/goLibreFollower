APP=goLibreFollower
VERSION = 1.0.0
ICON = "appicon.ico"

windows:
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" ~/go/bin/./fyne package -os windows --icon ${ICON} --appVersion ${VERSION} --name ${APP}-WINDOWS
	zip ${APP}-WINDOWS_${VERSION}.zip ${APP}-WINDOWS.exe
	rm -rf ${APP}-WINDOWS.exe
darwin:
	env GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 ~/go/bin/./fyne package -os darwin --icon ${ICON} --appVersion ${VERSION} --name ${APP}-MAC
	zip -r ${APP}-MAC_${VERSION}.zip ${APP}-MAC.app
	rm -rf ${APP}-MAC.app
clean:
	rm -rf ${APP}-WINDOWS_${VERSION}.zip ${APP}-MAC_${VERSION}.zip
all: windows darwin 