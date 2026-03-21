package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type InstallParams struct {
	IsCompleted          bool
	Path                 string
	MetasiaAssetsUrl     string
	MetasiaAssetFileName string
	PluginsPath          string
	Plugins              []PluginInfo
}

type PluginInfo struct {
	AssetUrl string
	FileName string
}

func (p *InstallParams) SetDefault() {
	version := "v0.2.2"
	switch runtime.GOOS {
	case "windows":
		p.Path = filepath.Join(os.Getenv("ProgramFiles"), "Metasia")
		p.MetasiaAssetFileName = "MetasiaEditor-windows-x64.zip"
		p.PluginsPath = filepath.Join(os.Getenv("APPDATA"), "Metasia", "Plugins")

		mediaFoundationPlugin := PluginInfo{
			FileName: "MediaFoundationPlugin-windows-x64.zip",
		}
		mediaFoundationPlugin.AssetUrl, _ = ResolveGithubAssetUrl("SousiOmine", "Metasia.Editor.MediaFoundationPlugin", "v0.2.2", mediaFoundationPlugin.FileName)
		p.Plugins = append(p.Plugins, mediaFoundationPlugin)
	default:
		p.Path = ""
		p.MetasiaAssetFileName = ""
		userConfigDir, _ := os.UserConfigDir()
		p.PluginsPath = filepath.Join(userConfigDir, "Metasia", "Plugins")
	}

	var err error
	p.MetasiaAssetsUrl, err = ResolveGithubAssetUrl("SousiOmine", "Metasia", version, p.MetasiaAssetFileName)
	p.PluginsPath = filepath.Join(p.Path, "Plugins")
	if err != nil {
		fmt.Println("Failed to resolve asset URL:", err)
		p.MetasiaAssetsUrl = ""
	}
}

func ResolveGithubAssetUrl(owner string, repo string, tag string, assetName string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/tags/%s", owner, repo, tag), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Metasia-Installer")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status: %d", resp.StatusCode)
	}

	type Asset struct {
		Name               string `json:"name"`
		BrowserDownloadUrl string `json:"browser_download_url"`
	}

	type Release struct {
		Assets []Asset `json:"assets"`
	}

	body, _ := io.ReadAll(resp.Body)
	var release Release
	json.Unmarshal(body, &release)

	for _, asset := range release.Assets {
		if asset.Name == assetName {
			return asset.BrowserDownloadUrl, nil
		}
	}
	return "", fmt.Errorf("asset not found: %s", assetName)
}
