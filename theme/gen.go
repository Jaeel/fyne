// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Jaleel/fyne"
)

const fontFace = "NotoSans"

const fileHeader = "// auto-generated\n" + // to exclude this file in goreportcard (it has to be first)
	"// Code generated by '$ fyne bundle'. DO NOT EDIT." // idiomatic mark, see https://golang.org/s/generatedcode

func formatVariable(name string) string {
	str := strings.Replace(name, "-", "", -1)
	return strings.Replace(str, "_", "", -1)
}

func bundleFile(name string, filepath string, f io.Writer) {
	res, err := fyne.LoadResourceFromPath(filepath)
	if err != nil {
		fyne.LogError("Unable to load file "+filepath, err)
		return
	}
	staticRes, ok := res.(*fyne.StaticResource)
	if !ok {
		fyne.LogError("Unable to format resource", fmt.Errorf("unexpected resource type %T", res))
		return
	}
	v := fmt.Sprintf("var %s = &fyne.StaticResource{\n\tStaticName: %q,\n\tStaticContent: []byte(%q),\n}\n\n",
		formatVariable(name), staticRes.StaticName, staticRes.StaticContent)
	_, err = f.Write([]byte(v))
	if err != nil {
		fyne.LogError("Unable to write to bundled file", err)
	}
}

func bundleFont(font, name string, f io.Writer) {
	_, dirname, _, _ := runtime.Caller(0)
	path := path.Join(path.Dir(dirname), "font", fmt.Sprintf("%s-%s.ttf", font, name))

	if name == "Powerline" && font != fontFace {
		name = "Monospace"
	}

	bundleFile(strings.ToLower(name), path, f)
}

func iconDir() string {
	_, dirname, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(dirname), "icons")
}

func bundleIcon(name string, f io.Writer) {
	path := path.Join(iconDir(), fmt.Sprintf("%s.svg", name))

	formatted := fmt.Sprintf("%sIconRes", strings.ToLower(name))
	bundleFile(formatted, path, f)
}

func writeFile(filename string, contents []byte) error {
	formatted, err := format.Source(contents)
	if err != nil {
		return err
	}
	_, dirname, _, _ := runtime.Caller(0)
	return ioutil.WriteFile(filepath.Join(filepath.Dir(dirname), filename), formatted, 0644)
}

func main() {
	f := &bytes.Buffer{}
	f.WriteString(fileHeader + "\n\npackage theme\n\nimport \"github.com/Jaleel/fyne\"\n\n")
	bundleFont(fontFace, "Regular", f)
	bundleFont(fontFace, "Bold", f)
	bundleFont(fontFace, "Italic", f)
	bundleFont(fontFace, "BoldItalic", f)
	bundleFont("DejaVuSansMono", "Powerline", f)
	err := writeFile("bundled-fonts.go", f.Bytes())
	if err != nil {
		fyne.LogError("Unable to write file:", err)
		return
	}

	f = &bytes.Buffer{}
	f.WriteString(fileHeader + "\n\npackage theme\n\nimport \"github.com/Jaleel/fyne\"\n\n")
	icon := path.Join(iconDir(), "fyne.png")
	bundleFile("fyne-logo", icon, f)

	bundleIcon("cancel", f)
	bundleIcon("check", f)
	bundleIcon("delete", f)
	bundleIcon("search", f)
	bundleIcon("search-replace", f)
	bundleIcon("menu", f)
	bundleIcon("menu-expand", f)

	bundleIcon("check-box", f)
	bundleIcon("check-box-blank", f)
	bundleIcon("radio-button", f)
	bundleIcon("radio-button-checked", f)

	bundleIcon("content-add", f)
	bundleIcon("content-remove", f)
	bundleIcon("content-cut", f)
	bundleIcon("content-copy", f)
	bundleIcon("content-paste", f)
	bundleIcon("content-redo", f)
	bundleIcon("content-undo", f)

	bundleIcon("color-achromatic", f)
	bundleIcon("color-chromatic", f)
	bundleIcon("color-palette", f)

	bundleIcon("document", f)
	bundleIcon("document-create", f)
	bundleIcon("document-print", f)
	bundleIcon("document-save", f)

	bundleIcon("more-horizontal", f)
	bundleIcon("more-vertical", f)

	bundleIcon("info", f)
	bundleIcon("question", f)
	bundleIcon("warning", f)
	bundleIcon("error", f)

	bundleIcon("arrow-back", f)
	bundleIcon("arrow-down", f)
	bundleIcon("arrow-forward", f)
	bundleIcon("arrow-up", f)
	bundleIcon("arrow-drop-down", f)
	bundleIcon("arrow-drop-up", f)

	bundleIcon("file", f)
	bundleIcon("file-application", f)
	bundleIcon("file-audio", f)
	bundleIcon("file-image", f)
	bundleIcon("file-text", f)
	bundleIcon("file-video", f)
	bundleIcon("folder", f)
	bundleIcon("folder-new", f)
	bundleIcon("folder-open", f)
	bundleIcon("help", f)
	bundleIcon("history", f)
	bundleIcon("home", f)
	bundleIcon("settings", f)

	bundleIcon("mail-attachment", f)
	bundleIcon("mail-compose", f)
	bundleIcon("mail-forward", f)
	bundleIcon("mail-reply", f)
	bundleIcon("mail-reply_all", f)
	bundleIcon("mail-send", f)

	bundleIcon("media-music", f)
	bundleIcon("media-photo", f)
	bundleIcon("media-video", f)
	bundleIcon("media-fast-forward", f)
	bundleIcon("media-fast-rewind", f)
	bundleIcon("media-pause", f)
	bundleIcon("media-play", f)
	bundleIcon("media-record", f)
	bundleIcon("media-replay", f)
	bundleIcon("media-skip-next", f)
	bundleIcon("media-skip-previous", f)
	bundleIcon("media-stop", f)

	bundleIcon("view-fullscreen", f)
	bundleIcon("view-refresh", f)
	bundleIcon("view-zoom-fit", f)
	bundleIcon("view-zoom-in", f)
	bundleIcon("view-zoom-out", f)

	bundleIcon("volume-down", f)
	bundleIcon("volume-mute", f)
	bundleIcon("volume-up", f)

	bundleIcon("visibility", f)
	bundleIcon("visibility-off", f)

	bundleIcon("download", f)
	bundleIcon("computer", f)
	bundleIcon("storage", f)
	bundleIcon("upload", f)

	bundleIcon("account", f)
	bundleIcon("login", f)
	bundleIcon("logout", f)

	bundleIcon("list", f)
	bundleIcon("grid", f)

	err = writeFile("bundled-icons.go", f.Bytes())
	if err != nil {
		fyne.LogError("Unable to write file: ", err)
		return
	}
}
