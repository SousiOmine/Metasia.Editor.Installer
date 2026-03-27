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
	downloadZipPath := filepath.Join(e.Params.Path, "MetasiaEditor.zip")

	if e.Params.MetasiaAssetsUrl == "" {
		return fmt.Errorf("download url is empty")
	}

	_, err := os.Stat(e.Params.Path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(e.Params.Path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create install directory: %w", err)
		}
	}

	fmt.Println(e.Params.MetasiaAssetsUrl + "からファイルをダウンロードします...")
	err = downloadFile(e.Params.MetasiaAssetsUrl, downloadZipPath)
	if err != nil {
		return fmt.Errorf("failed to download main file: %w", err)
	}

	fmt.Println("ダウンロード完了。ファイルを展開します...")
	err = unzip(downloadZipPath, e.Params.Path)
	if err != nil {
		return fmt.Errorf("failed to unzip main file: %w", err)
	}

	fmt.Println("一時ファイルを削除します...")
	err = os.Remove(downloadZipPath)
	if err != nil {
		return fmt.Errorf("failed to remove temp file: %w", err)
	}

	// プラグイン用フォルダを作成
	err = os.MkdirAll(e.Params.PluginsPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create plugins directory: %w", err)
	}

	for _, pluginInfo := range e.Params.Plugins {
		pluginZipPath := filepath.Join(e.Params.PluginsPath, pluginInfo.FileName)
		fmt.Println(pluginInfo.AssetUrl + "からファイルをダウンロードします...")
		err = downloadFile(pluginInfo.AssetUrl, pluginZipPath)
		if err != nil {
			return fmt.Errorf("failed to download plugin %s: %w", pluginInfo.FileName, err)
		}

		fmt.Println("ファイルを展開します...")
		err = unzip(pluginZipPath, e.Params.PluginsPath)
		if err != nil {
			return fmt.Errorf("failed to unzip plugin %s: %w", pluginInfo.FileName, err)
		}

		fmt.Println("一時ファイルを削除します...")
		err = os.Remove(pluginZipPath)
		if err != nil {
			return fmt.Errorf("failed to remove plugin temp file %s: %w", pluginInfo.FileName, err)
		}
	}

	return nil
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
