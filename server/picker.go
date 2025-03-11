package main

import "net/http"

// these may change, but i think they'll cover the necessary parts of the google photos picker API
func (apiConf *apiConfig) showQRCode(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("QR Code"))
}
func (apiConf *apiConfig) fetchImages(w http.ResponseWriter, r *http.Request)      {}
func (apiConf *apiConfig) loadImageIntoImg(w http.ResponseWriter, r *http.Request) {}
