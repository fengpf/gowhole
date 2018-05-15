package watermark

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/math/fixed"
)

var (
	wm       = flag.String("wm", "./img/wm.png", "最终效果图片")
	fontfile = flag.String("fontfile", "./img/msyhb.ttf", "filename of the ttf font")
)

// Watermark  create watermark info.
type Watermark struct {
	canvasWidth  int
	canvasHeight int
	txtWidth     int
	size         int
	text         string
	imgPath      string
	srcImg       image.Image
	canvas       *image.NRGBA
	c            *Context
	f            *truetype.Font
}

// NewWatermark create watermark.
func NewWatermark(path, text string, size int) (w *Watermark) {
	w = &Watermark{
		text:    text,
		size:    size,
		imgPath: path,
	}
	w.readFont(*fontfile)
	w.readSrcImg(w.imgPath)
	w.textWidth()
	w.canvasWidth = w.newImgWidth()
	w.canvasHeight = w.imgHeight()
	w.canvas = w.newCanvas(w.canvasWidth, w.canvasHeight)
	draw.Draw(w.canvas, w.canvas.Bounds(), w.canvas, image.ZP, draw.Src)
	return
}

func (w *Watermark) readSrcImg(path string) (err error) {
	w.srcImg, err = Open(path)
	defer debug.FreeOSMemory()
	if err != nil {
		log.Println(err)
	}
	return
}

func (w *Watermark) readFont(path string) (err error) {
	fbs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return
	}
	w.f, err = ParseFont(fbs)
	if err != nil {
		log.Println(err)
		return
	}
	w.c = NewContext()
	w.c.SetFont(w.f)
	w.c.SetFontSize(float64(w.size))
	return
}

func (w *Watermark) imgWidth() int {
	return w.srcImg.Bounds().Max.X
}

func (w *Watermark) imgHeight() int {
	return w.srcImg.Bounds().Max.Y
}

func (w *Watermark) newImgWidth() int {
	return w.txtWidth + w.imgWidth()
}

func (w *Watermark) newCanvas(width, height int) *image.NRGBA {
	return New(width, height, color.RGBA{255, 0, 0, 0})
}

func (w *Watermark) fillColor(r, g, b, a int32) image.Image {
	return &image.Uniform{color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}}
}

func (w *Watermark) textWidth() (err error) {
	box, err := w.c.FontBox(w.text, w.pt(3, 5))
	if err != nil {
		log.Println(err)
		return
	}
	widthStr := strings.Split(Int26_6ToString(box.X), ":")[0]
	wid, err := strconv.ParseInt(widthStr, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	w.txtWidth = int(wid)
	return
}

func (w *Watermark) pt(x, y int) fixed.Point26_6 {
	return Pt(x, y+int(w.c.PointToFixed(float64(w.size))>>6))
}

func (w *Watermark) setFont(dstRgba *image.NRGBA, fsrc image.Image, pt fixed.Point26_6) (err error) {
	w.c.SetClip(w.canvas.Bounds())
	w.c.SetDst(dstRgba)
	w.c.SetSrc(fsrc)
	_, err = w.c.DrawString(w.text, pt)
	if err != nil {
		log.Println(err)
	}
	return
}

func (w *Watermark) composite(dstCanvas *image.NRGBA, src image.Image, isLeft bool) {
	var p image.Point
	if isLeft {
		p = image.Point{-int(w.txtWidth), 0}
	} else {
		p = image.ZP
	}
	draw.Draw(dstCanvas, image.Rect(0, 0, w.newImgWidth(), w.imgHeight()), src, p, draw.Over)
}

// Draw  write text to the left or right of img.
func (w *Watermark) Draw(isLeft bool) (err error) {
	var pt fixed.Point26_6
	if isLeft {
		pt = w.pt(3, 5)
	} else {
		pt = w.pt(w.imgWidth(), 8)
	}
	black := w.fillColor(0, 0, 0, 125)
	w.setFont(w.canvas, black, pt)
	blurRgba := Blur(w.canvas, 6, 3.5)
	white := w.fillColor(255, 255, 255, 180)
	w.setFont(blurRgba, white, pt)
	w.composite(blurRgba, w.srcImg, isLeft)
	Save(blurRgba, *wm)
	return
}
