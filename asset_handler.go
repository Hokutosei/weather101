package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleAssets(assets ...string) {
	for _, asset := range assets {
		go func(asset string) {
			start := time.Now()
			asset_url_path := fmt.Sprintf("/%s/", asset)
			asset_dir := fmt.Sprintf("public/%s", asset)
			http.Handle(asset_url_path, http.StripPrefix(asset_url_path, http.FileServer(http.Dir(asset_dir))))
			fmt.Println(asset, " -->> served!", time.Since(start))
		}(asset)

	}
}

func initializeAssets() {
	assetsToHandle := []string{"images", "css", "js", "fonts"}
	handleAssets(assetsToHandle...)
}
