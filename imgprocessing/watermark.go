package imgprocessing

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"
	"runtime/debug"

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

func isTooBright(img image.Image) bool {
	var (
		pixCount, totalBrightness float64
	)
	pixCount = 0
	totalBrightness = 0
	AdjustFunc(img, func(c color.NRGBA) color.NRGBA {
		brightness := 0.2126*float64(c.R) + 0.7152*float64(c.G) + 0.0722*float64(c.B)
		totalBrightness += brightness
		pixCount++
		return c
	})
	averBrightness := totalBrightness / pixCount
	// assume that average brightness higher than 100 is too bright.
	return averBrightness > 100
}

// DrawWaterMark deal font&img.
func DrawWaterMark(uname string) (err error) {
	var (
		finalRgba *image.NRGBA
	)
	// 读字体数据
	fbs, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		return
	}
	f, err := freetype.ParseFont(fbs)
	if err != nil {
		return
	}
	// 新建一个指定大小的 RGBA位图
	rgba := New(300, 50, color.RGBA{255, 0, 0, 0})
	draw.Draw(rgba, rgba.Bounds(), rgba, image.ZP, draw.Src)
	// 填充模糊本体文字.
	c := freetype.NewContext()
	//c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	fillColor := &image.Uniform{color.NRGBA{0, 0, 0, 125}}
	//println(fillColor.Opaque())//检查不透明度
	c.SetSrc(fillColor)
	pt := freetype.Pt(3, 5+int(c.PointToFixed(*size)>>6)) // 字出现的位置
	//fmt.Printf("%v\n", pt)
	_, err = c.DrawString(uname, pt)
	if err != nil {
		log.Println(err)
		return
	}
	// blur rgb
	blurRgba := Blur(rgba, 6, 3.5)
	finalRgba = blurRgba
	if isTooBright(blurRgba) {
		println("isTooBright")
		finalRgba = AdjustBrightness(blurRgba, -10)
	}
	//Save(finalRgba, *wmblur) //保存模糊图片
	// 重新画本体文字.
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(blurRgba.Bounds())
	c.SetDst(blurRgba)
	fillColor2 := &image.Uniform{color.NRGBA{255, 255, 255, 180}}
	c.SetSrc(fillColor2)
	_, err = c.DrawString(uname, pt)
	if err != nil {
		log.Println(err)
		return
	}
	composite(finalRgba, *wmimg)
	return
}

// composite  composite  logo & wmimg.
func composite(rga *image.NRGBA, dest string) (err error) {
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
	Save(rga, dest)
	println("composite ...")
	return
}
