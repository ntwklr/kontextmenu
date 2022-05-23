clean:
	go clean
	rm -rf build/

build:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o build/kontextmenu-mac_arm64
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o build/kontextmenu-mac_amd64
	lipo -create -output build/kontextmenu-mac build/kontextmenu-mac_amd64 build/kontextmenu-mac_arm64

build-mac-app:
	mkdir -p build/app
	cd Kontextmenu/ && xcodebuild
	cp -r Kontextmenu/build/Release/Kontextmenu.app build/app/

copy-binary-to-mac:
	cp build/kontextmenu-mac build/app/Kontextmenu.app/Contents/MacOS/Kontextmenu

create-dmg:
	create-dmg \
	  --volname "Kontextmenu Installer" \
	  --window-pos 200 120 \
	  --window-size 800 400 \
	  --icon-size 100 \
	  --icon "Kontextmenu.app" 200 190 \
	  --hide-extension "Kontextmenu.app" \
	  --app-drop-link 600 185 \
	  "build/Kontextmenu.dmg" \
	  "build/app/"

dist-mac-app: clean build build-mac-app copy-binary-to-mac

dist-mac-dmg: dist-mac-app create-dmg

.PHONY: build