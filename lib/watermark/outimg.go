package watermark

import (
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"time"
)

func printImg(w http.ResponseWriter, req *http.Request) {
	logo, err := Open(*wm)
	if err != nil {
		return
	}
	bounds := logo.Bounds()
	rgb := image.NewRGBA(bounds)
	draw.Draw(rgb, bounds, logo, image.ZP, draw.Over)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, rgb)
}

// HTTPPrint print image stream.
func HTTPPrint() {
	http.HandleFunc("/", printImg)
	s := &http.Server{
		Addr:           ":9000",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
