package imgprocessing

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/golang/freetype"
)

const (
	typeName = 1
	uidMark  = "/Users/fpf/gowork/src/gostudy/img/uid_mark.png" // uid图片
)

var (
	unamemark = flag.String("unamemark", "/Users/fpf/gowork/src/gostudy/img/mark.png", "用户名图片")
	wmblur    = flag.String("wmblur", "/Users/fpf/gowork/src/gostudy/img/wmblur.png", "模糊图片")
	wmimg     = flag.String("wmimg", "/Users/fpf/gowork/src/gostudy/img/wmimg.png", "最终效果图片")
	size      = flag.Float64("size", 32, "font size in points")
	dpi       = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch") // 屏幕每英寸的分辨率 DPI是Dots Per Inch（每英寸所打印的点数）
	fontfile  = flag.String("fontfile", "/Users/fpf/gowork/src/gostudy/img/msyhb.ttf", "filename of the ttf font")
	spacing   = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb      = flag.Bool("whiteonblack", false, "white text on a black background")
)

func createWaterMark(characters string) (finalRgba *image.NRGBA, err error) {
	// new rgb.
	rgba := New(300, 50, color.RGBA{255, 0, 0, 0})
	draw.Draw(rgba, rgba.Bounds(), rgba, image.ZP, draw.Src)
	setfont(characters, rgba, &image.Uniform{color.NRGBA{0, 0, 0, 125}})
	// blur rgb.
	blurRgba := Blur(rgba, 6, 3.5)
	finalRgba = blurRgba
	if IsTooBright(blurRgba) {
		finalRgba = AdjustBrightness(blurRgba, -10)
	}
	//Save(finalRgba, *wmblur) // save blur rgb
	setfont(characters, blurRgba, &image.Uniform{color.NRGBA{255, 255, 255, 180}})
	return
}

func setfont(characters string, rgba *image.NRGBA, fillColor image.Image) (err error) {
	// 读字体数据
	fbs, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		return
	}
	f, err := freetype.ParseFont(fbs)
	if err != nil {
		return
	}
	c := freetype.NewContext()
	//c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	//println(fillColor.Opaque())// 检查不透明度
	c.SetSrc(fillColor)
	pt := freetype.Pt(3, 5+int(c.PointToFixed(*size)>>6)) // 字出现的位置
	//fmt.Printf("%v\n", pt)
	_, err = c.DrawString(characters, pt)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// WaterMark get the img.
func WaterMark(uname string) (err error) {
	rga, err := createWaterMark(uname)
	if err != nil {
		log.Println(err)
		return
	}
	dx := rga.Bounds().Dx()
	dy := rga.Bounds().Dy()
	px := -185
	py := 0
	logo, err := Open(*unamemark)
	if err != nil {
		return
	}
	defer debug.FreeOSMemory()
	draw.Draw(rga, image.Rect(0, 0, dx, dy), logo, image.Point{px, py}, draw.Over)
	Save(rga, *wmimg)
	println("composite ...")
	return
}

func printImg(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	rga, err := createWaterMark("图片流输出")
	if err != nil {
		log.Println(err)
		return
	}
	png.Encode(w, rga)
}

// HTTPPrint print image stream.
func HTTPPrint() {
	http.HandleFunc("/", printImg)
	s := &http.Server{
		Addr:           ":88",
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
