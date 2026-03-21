package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type InstallExecutor struct {
	Params InstallParams
}

func (e *InstallExecutor) Execute() error {
	var err error
	downloadZipPath := filepath.Join(e.Params.Path, "MetasiaEditor.zip")

	if e.Params.MetasiaAssetsUrl == "" {
		return fmt.Errorf("download url is empty")
	}

	_, err = os.Stat(e.Params.Path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(e.Params.Path, 0755)
		if err != nil {
			fmt.Println(err)
			goto Done
		}
	}

	fmt.Println(e.Params.MetasiaAssetsUrl + "からファイルをダウンロードします...")
	err = downloadFile(e.Params.MetasiaAssetsUrl, downloadZipPath)

	fmt.Println("ダウンロード完了。ファイルを展開します...")
	err = unzip(downloadZipPath, e.Params.Path)

	fmt.Println("一時ファイルを削除します...")
	os.Remove(downloadZipPath)

	// プラグイン用フォルダを作成
	os.MkdirAll(e.Params.PluginsPath, 0755)

	for plugin := range e.Params.Plugins {
		pluginInfo := e.Params.Plugins[plugin]
		fmt.Println(pluginInfo.AssetUrl + "からファイルをダウンロードします...")
		err = downloadFile(pluginInfo.AssetUrl, filepath.Join(e.Params.PluginsPath, pluginInfo.FileName))
		fmt.Println("ファイルを展開します...")
		err = unzip(filepath.Join(e.Params.PluginsPath, pluginInfo.FileName), e.Params.PluginsPath)
		fmt.Println("一時ファイルを削除します...")
		os.Remove(filepath.Join(e.Params.PluginsPath, pluginInfo.FileName))
	}

Done:
	return err
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("不正なファイルパス: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		rc.Close()
		outFile.Close()

		if err != nil {
			return err
		}

	}
	return nil

}
