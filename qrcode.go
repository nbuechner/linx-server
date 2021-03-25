package main

import (
	"image/png"
	"net/http"
	"github.com/zenazn/goji/web"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func fileQRCodeHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fileName := c.URLParams["name"]
	dataString := getSiteURL(r) + fileName
        qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 180, 180)
	png.Encode(w, qrCode)
}
