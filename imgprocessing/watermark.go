package imgprocessing

import (
	"flag"
	"gostudy/fontprocessing"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/math/fixed"
)

const (
	typeName = 1
)

var (
	unameMark = flag.String("unameMark", "./img/mark.png", "用户名图片")
	midMark   = flag.String("midMark", "./img/uid_mark.png", "用户mid图片")
	wmBlur    = flag.String("wmBlur", "./img/wmBlur.png", "模糊图片")
	wmUname   = flag.String("wmUname", "./img/wmUname.png", "最终uname效果图片")
	wmMid     = flag.String("wmMid", "./img/wmMid.png", "最终mid效果图片")
	testimg   = flag.String("testimg", "./img/testimg.png", "最终效果图片")
	size      = flag.Float64("size", 32, "font size in points")
	dpi       = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch") // 屏幕每英寸的分辨率 DPI是Dots Per Inch（每英寸所打印的点数）
	fontfile  = flag.String("fontfile", "./img/msyhb.ttf", "filename of the ttf font")
	spacing   = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb      = flag.Bool("whiteonblack", false, "white text on a black background")
)

// CreateByUname get uname watermark.
func CreateByUname(ch string) (err error) {
	c, err := openfont(ch)
	if err != nil {
		log.Println(err)
		return
	}
	box, err := textBox(c, ch)
	if err != nil {
		log.Println(err)
		return
	}
	widthStr := strings.Split(fontprocessing.Int26_6ToString(box.X), ":")[0]
	width, err := strconv.ParseInt(widthStr, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	// open  logo.
	logo, err := Open(*unameMark)
	defer debug.FreeOSMemory()
	if err != nil {
		return
	}
	lx := logo.Bounds().Dx()
	rgbaWidth := lx + int(width)
	// new rgb.
	rgba := New(rgbaWidth, 50, color.RGBA{255, 0, 0, 0})
	draw.Draw(rgba, rgba.Bounds(), rgba, image.ZP, draw.Src)
	pt := fontprocessing.Pt(3, 5+int(c.PointToFixed(*size)>>6)) // 字出现的位置
	setfont(c, pt, ch, rgba, &image.Uniform{color.NRGBA{0, 0, 0, 125}})
	// blur rgb.
	blurRgba := Blur(rgba, 6, 3.5)
	finalRgba := blurRgba
	if IsTooBright(blurRgba) {
		finalRgba = AdjustBrightness(blurRgba, -10)
	}
	//Save(finalRgba, *wmblur) // save blur rgb
	setfont(c, pt, ch, blurRgba, &image.Uniform{color.NRGBA{255, 255, 255, 180}})
	//composite img
	dx := finalRgba.Bounds().Dx()
	dy := finalRgba.Bounds().Dy()
	println(rgbaWidth)
	draw.Draw(finalRgba, image.Rect(0, 0, dx+lx, dy), logo, image.Point{-int(width), 0}, draw.Over)
	Save(finalRgba, *wmUname)
	println("uname mark composite ...")
	return
}

// CreateByMid get mid watermark.
func CreateByMid(mid int64) (err error) {
	ch := strconv.FormatInt(mid, 10)
	c, err := openfont(ch)
	if err != nil {
		log.Println(err)
		return
	}
	box, err := textBox(c, ch)
	if err != nil {
		log.Println(err)
		return
	}
	widthStr := strings.Split(fontprocessing.Int26_6ToString(box.X), ":")[0]
	width, err := strconv.ParseInt(widthStr, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	// open  logo.
	logo, err := Open(*midMark)
	defer debug.FreeOSMemory()
	if err != nil {
		return
	}
	lx := logo.Bounds().Dx()
	rgbaWidth := int(width) + lx
	// new rgb.
	rgba := New(rgbaWidth, 50, color.RGBA{255, 0, 0, 0})
	draw.Draw(rgba, rgba.Bounds(), rgba, image.ZP, draw.Src)
	pt := fontprocessing.Pt(lx, 8+int(c.PointToFixed(*size)>>6)) // 字出现的位置
	setfont(c, pt, ch, rgba, &image.Uniform{color.NRGBA{0, 0, 0, 125}})
	// blur rgb.
	blurRgba := Blur(rgba, 6, 3.5)
	finalRgba := blurRgba
	if IsTooBright(blurRgba) {
		finalRgba = AdjustBrightness(blurRgba, -10)
	}
	//Save(finalRgba, *wmblur) // save blur rgb
	setfont(c, pt, ch, blurRgba, &image.Uniform{color.NRGBA{255, 255, 255, 180}})
	//composite img
	dx := finalRgba.Bounds().Dx()
	dy := finalRgba.Bounds().Dy()
	println(rgbaWidth) //
	draw.Draw(finalRgba, image.Rect(0, 0, dx+lx, dy), logo, image.ZP, draw.Over)
	Save(finalRgba, *wmMid)
	println("mid mark composite ...")
	return
}

func openfont(characters string) (c *fontprocessing.Context, err error) {
	// 读字体数据
	fbs, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := fontprocessing.ParseFont(fbs)
	if err != nil {
		log.Println(err)
		return
	}
	c = fontprocessing.NewContext()
	c.SetFont(f)
	c.SetFontSize(*size)
	return
}

func textBox(c *fontprocessing.Context, ch string) (box fixed.Point26_6, err error) {
	pt := fontprocessing.Pt(3, 5+int(c.PointToFixed(*size)>>6)) // 字出现的位置
	box, err = c.FontBox(ch, pt)
	if err != nil {
		log.Println(err)
	}
	return
}

func setfont(c *fontprocessing.Context, pt fixed.Point26_6, ch string, rgba *image.NRGBA, fillColor image.Image) (err error) {
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fillColor)
	_, err = c.DrawString(ch, pt)
	if err != nil {
		log.Println(err)
	}
	return
}

func printImg(w http.ResponseWriter, req *http.Request) {
	logo, err := Open(*wmUname)
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
