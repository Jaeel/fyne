// auto-generated
// Code generated by '$ fyne bundle'. DO NOT EDIT.

package templates

import "github.com/Jaleel/fyne"

var resourceInfoPlist = &fyne.StaticResource{
	StaticName: "Info.plist",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple Computer//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n<dict>\n\t<key>CFBundleName</key>\n\t<string>{{.Name}}</string>\n\t<key>CFBundleExecutable</key>\n\t<string>{{.ExeName}}</string>\n\t<key>CFBundleIdentifier</key>\n\t<string>{{.AppID}}</string>\n\t<key>CFBundleIconFile</key>\n\t<string>icon.icns</string>\n\t<key>CFBundleShortVersionString</key>\n\t<string>{{.Version}}</string>\n\t<key>CFBundleSupportedPlatforms</key>\n\t<array>\n\t\t<string>MacOSX</string>\n\t</array>\n\t<key>CFBundleVersion</key>\n\t<string>{{.Build}}</string>\n\t<key>NSHighResolutionCapable</key>\n\t<true/>\n\t<key>NSSupportsAutomaticGraphicsSwitching</key>\n\t<true/>\n\t<key>CFBundleInfoDictionaryVersion</key>\n\t<string>6.0</string>\n\t<key>CFBundlePackageType</key>\n\t<string>APPL</string>\n\t<key>LSApplicationCategoryType</key>\n\t<string>public.app-category.{{.Category}}</string>\n\t<key>LSMinimumSystemVersion</key>\n\t<string>10.11</string>\n</dict>\n</plist>"),
}
var resourceMakefile = &fyne.StaticResource{
	StaticName: "Makefile",
	StaticContent: []byte(
		"# If PREFIX isn't provided, we check for $(DESTDIR)/usr/local and use that if it exists.\n# Otherwice we fall back to using /usr.\n\nLOCAL != test -d $(DESTDIR)/usr/local && echo -n \"/local\" || echo -n \"\"\nLOCAL ?= $(shell test -d $(DESTDIR)/usr/local && echo \"/local\" || echo \"\")\nPREFIX ?= /usr$(LOCAL)\n\ndefault:\n\t# Run \"sudo make install\" to install the application.\n\t# Run \"sudo make uninstall\" to uninstall the application.\n\ninstall:\n\tinstall -Dm00644 usr/{{.Local}}share/applications/{{.Name}}.desktop $(DESTDIR)$(PREFIX)/share/applications/{{.Name}}.desktop\n\tinstall -Dm00755 usr/{{.Local}}bin/{{.Exec}} $(DESTDIR)$(PREFIX)/bin/{{.Exec}}\n\tinstall -Dm00644 usr/{{.Local}}share/pixmaps/{{.Icon}} $(DESTDIR)$(PREFIX)/share/pixmaps/{{.Icon}}\nuninstall:\n\t-rm $(DESTDIR)$(PREFIX)/share/applications/{{.Name}}.desktop\n\t-rm $(DESTDIR)$(PREFIX)/bin/{{.Exec}}\n\t-rm $(DESTDIR)$(PREFIX)/share/pixmaps/{{.Icon}}"),
}
var resourceAppDesktop = &fyne.StaticResource{
	StaticName: "app.desktop",
	StaticContent: []byte(
		"[Desktop Entry]\nType=Application\nName={{.Name}}\nExec={{.Exec}}\nIcon={{.Name}}\n"),
}
var resourceAppManifest = &fyne.StaticResource{
	StaticName: "app.manifest",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n<assembly xmlns=\"urn:schemas-microsoft-com:asm.v1\" manifestVersion=\"1.0\" xmlns:asmv3=\"urn:schemas-microsoft-com:asm.v3\">\n<assemblyIdentity version=\"{{.CombinedVersion}}\" processorArchitecture=\"*\" name=\"{{.Name}}\" type=\"win32\"/>\n</assembly>"),
}
var resourceAppxmanifestXML = &fyne.StaticResource{
	StaticName: "appxmanifest.XML",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<Package xmlns=\"http://schemas.microsoft.com/appx/2010/manifest\">\n  <Identity Name=\"{{.AppID}}\"\n            Version=\"{{.Version}}\"\n            Publisher=\"{{.Developer}}\" />\n  <Properties>\n    <DisplayName>{{.Name}}</DisplayName>\n    <PublisherDisplayName>{{.DeveloperName}}</PublisherDisplayName>\n    <Logo>Icon.png</Logo>\n  </Properties>\n  <Prerequisites>\n    <OSMinVersion>6.2</OSMinVersion>\n    <OSMaxVersionTested>10.0</OSMaxVersionTested>\n  </Prerequisites>\n  <Resources>\n    <Resource Language=\"en-us\" />\n  </Resources>\n\n  <!-- TODO <Applications> -->\n</Package>"),
}
var resourceEntitlementsDarwinPlist = &fyne.StaticResource{
	StaticName: "entitlements-darwin.plist",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n<dict>\n    <key>com.apple.security.app-sandbox</key>\n    <true/>\n</dict>\n</plist>"),
}
var resourceEntitlementsIosPlist = &fyne.StaticResource{
	StaticName: "entitlements-ios.plist",
	StaticContent: []byte(
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n<dict>\n    <key>application-identifier</key>\n    <string>{{.TeamID}}.{{.AppID}}</string>\n</dict>\n</plist>"),
}
var resourceXcassetsJSON = &fyne.StaticResource{
	StaticName: "xcassets.JSON",
	StaticContent: []byte(
		"{\n  \"images\" : [\n    {\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"2x\",\n      \"size\" : \"20x20\"\n    },\n    {\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"3x\",\n      \"size\" : \"20x20\"\n    },\n    {\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"2x\",\n      \"size\" : \"29x29\"\n    },\n    {\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"3x\",\n      \"size\" : \"29x29\"\n    },\n    {\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"2x\",\n      \"size\" : \"40x40\"\n    },\n    {\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"3x\",\n      \"size\" : \"40x40\"\n    },\n    {\n      \"filename\" : \"Icon_120.png\",\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"2x\",\n      \"size\" : \"60x60\"\n    },\n    {\n      \"filename\" : \"Icon_180.png\",\n      \"idiom\" : \"iphone\",\n      \"scale\" : \"3x\",\n      \"size\" : \"60x60\"\n    },\n    {\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"1x\",\n      \"size\" : \"20x20\"\n    },\n    {\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"2x\",\n      \"size\" : \"20x20\"\n    },\n    {\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"1x\",\n      \"size\" : \"29x29\"\n    },\n    {\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"2x\",\n      \"size\" : \"29x29\"\n    },\n    {\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"1x\",\n      \"size\" : \"40x40\"\n    },\n    {\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"2x\",\n      \"size\" : \"40x40\"\n    },\n    {\n      \"filename\" : \"Icon_76.png\",\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"1x\",\n      \"size\" : \"76x76\"\n    },\n    {\n      \"filename\" : \"Icon_152.png\",\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"2x\",\n      \"size\" : \"76x76\"\n    },\n    {\n      \"idiom\" : \"ipad\",\n      \"scale\" : \"2x\",\n      \"size\" : \"83.5x83.5\"\n    },\n    {\n      \"filename\" : \"Icon_1024.png\",\n      \"idiom\" : \"ios-marketing\",\n      \"scale\" : \"1x\",\n      \"size\" : \"1024x1024\"\n    }\n  ],\n  \"info\" : {\n    \"author\" : \"xcode\",\n    \"version\" : 1\n  }\n}"),
}
