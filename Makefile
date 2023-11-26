VERSION = 1.0.0
ICON = "appicon.ico"

windows:
	env GOOS=windows  GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" ~/go/bin/./fyne package -os windows --icon ${ICON} --appVersion ${VERSION}
darwin:
	~/go/bin/./fyne package -os darwin --icon ${ICON} --appVersion ${VERSION}
clean:
	rm -rf goLibreFollower goLibreFollower.exe 
all: windows darwin 