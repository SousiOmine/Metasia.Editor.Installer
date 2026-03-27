package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("こんにちは %s!, It's show time!", name)
}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
}

func (a *App) GetDefaultInstallParams() (InstallParams, error) {
	var params InstallParams
	params.SetDefault()
	return params, nil
}

func (a *App) SelectDirectoryDialog(defaultPath string) (string, error) {
	options := runtime.OpenDialogOptions{}

	selectedDir, err := runtime.OpenDirectoryDialog(a.ctx, options)
	if err != nil {
		return "", err
	}
	return selectedDir, nil
}

func (a *App) InstallExecute(params InstallParams) error {
	executor := InstallExecutor{
		Params: params,
	}
	return executor.Execute()
}
