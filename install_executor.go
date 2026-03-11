package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
