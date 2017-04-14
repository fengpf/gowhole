package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	blurRadius = 1
	filepath   = "./image/wm.png"
)

func main() {
	getBluredImg()
}

// x=blurRadius+i y=blurRadius+j
func getColorMatrix(imgMatrix [][][]uint8, x, y, length, whichColor int) (colorMat [][]float64) {
	var (
		col []float64
	)
	colorMat = make([][]float64, length)
	for i := x; i < length; i++ {
		col = make([]float64, length)
		for j := y; j < length; j++ {
			if whichColor == 0 {
				col[j] = float64(imgMatrix[i][j][0])
			} else if whichColor == 1 {
				col[j] = float64(imgMatrix[i][j][1])
			} else if whichColor == 2 {
				col[j] = float64(imgMatrix[i][j][2])
			}
		}
		colorMat[i] = col
	}
	fmt.Printf("%v", colorMat)

	return
}

func getBlurColor(imgMatrix [][][]uint8, x, y, whichColor int) (blurGray float64) {
	var (
		length int
	)
	blurGray = 0
	length = blurRadius*2 + 1
	//colorMat := getColorMatrix(imgMatrix, x, y, length, whichColor)
	weightArr := getWeightMatrix(length)
	//fmt.Printf("%v", imgMatrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			//fmt.Printf("%v\t", float64(imgMatrix[x][y][3]))
			if whichColor == 0 {
				blurGray = blurGray + weightArr[i][j]*float64(imgMatrix[x][y][0])
			} else if whichColor == 1 {
				blurGray = blurGray + weightArr[i][j]*float64(imgMatrix[x][y][1])
			} else if whichColor == 2 {
				blurGray = blurGray + weightArr[i][j]*float64(imgMatrix[x][y][2])
			}
		}
	}
	fmt.Printf("%v\t", blurGray)
	return
}

func getBluredImg() {
	length := blurRadius*2 + 1
	weightArr := getWeightMatrix(length)
	img := MustRead(filepath)
	height := len(img)
	width := len(img[0])
	imgMatrix := NewRGBAMatrix(height, width)
	copy(imgMatrix, img)
	println(imgMatrix, weightArr)
	//fmt.Printf("%v", imgMatrix)
	fmt.Printf("%v\t", weightArr)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			imgMatrix[i][j][0] = uint8(float64(imgMatrix[i][j][0]) * weightArr[i][j])
			imgMatrix[i][j][1] = uint8(float64(imgMatrix[i][j][1]) * weightArr[i][j])
			imgMatrix[i][j][2] = uint8(float64(imgMatrix[i][j][2]) * weightArr[i][j])

			// imgMatrix[i][j][0] = uint8(getBlurColor(imgMatrix, i, j, 0))
			// imgMatrix[i][j][1] = uint8(getBlurColor(imgMatrix, i, j, 1))
			// imgMatrix[i][j][2] = uint8(getBlurColor(imgMatrix, i, j, 2))
		}
	}
	// 以PNG格式保存文件
	err := SaveAsPng("./image/blur.png", imgMatrix)
	if err != nil {
		panic(err)
	}
	println("blur success..")
	return
}

// SaveAsPng save a image matrix as a jpeg,if unsuccessful it will return a error,quality must be 1 to 100.
func SaveAsPng(filepath string, imgMatrix [][][]uint8) error {
	height := len(imgMatrix)
	width := len(imgMatrix[0])
	if height == 0 || width == 0 {
		return errors.New("The input of matrix is illegal!")
	}
	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nrgba.SetNRGBA(j, i, color.NRGBA{imgMatrix[i][j][0], imgMatrix[i][j][1], imgMatrix[i][j][2], imgMatrix[i][j][3]})
		}
	}
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()
	png.Encode(outfile, nrgba)
	return nil
}

func getWeight(x, y float64) (weight float64) {
	var sigma float64
	sigma = (blurRadius*2 + 1) / 2
	weight = (1 / (2 * math.Pi * sigma * sigma)) * math.Pow(math.E, ((-(x*x + y*y))/((2*sigma)*(2*sigma))))
	return
}

func getWeightMatrix(length int) (weightArr [][]float64) {
	var (
		weightSum  float64
		weightArr2 []float64
	)
	weightArr = make([][]float64, length)
	for i := 0; i < length; i++ {
		weightArr2 = make([]float64, length)
		for j := 0; j < length; j++ {
			weightArr2[j] = getWeight(float64(j-blurRadius), float64(blurRadius-i))
			//fmt.Printf("%v\t", weightArr2[j])
		}
		weightArr[i] = weightArr2
		//fmt.Printf("%v\n", weightArr[i])
	}
	weightSum = 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			weightSum += weightArr[i][j]
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			weightArr[i][j] = weightArr[i][j] / weightSum
		}
	}
	return
}

// DecodeImage decode a image and retrun golang image interface。
func DecodeImage(filePath string) (img image.Image, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, _, err = image.Decode(reader)
	return
}

// MustRead read a image return a image matrix by path or image.
func MustRead(filepath string) (imgMatrix [][][]uint8) {
	img, decodeErr := DecodeImage(filepath)
	if decodeErr != nil {
		panic(decodeErr)
	}

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	src := convertToNRGBA(img)
	imgMatrix = NewRGBAMatrix(height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := src.At(j, i)
			r, g, b, a := c.RGBA()
			imgMatrix[i][j][0] = uint8(r)
			imgMatrix[i][j][1] = uint8(g)
			imgMatrix[i][j][2] = uint8(b)
			imgMatrix[i][j][3] = uint8(a)

		}
	}
	return
}

// New3DSlice get New3DSlice.
func New3DSlice(x int, y int, z int) (theSlice [][][]uint8) {
	theSlice = make([][][]uint8, x, x)
	for i := 0; i < x; i++ {
		s2 := make([][]uint8, y, y)
		for j := 0; j < y; j++ {
			s3 := make([]uint8, z, z)
			s2[j] = s3
		}
		theSlice[i] = s2
	}
	return
}

// NewRGBAMatrix get NewRGBAMatrix.
func NewRGBAMatrix(height int, width int) (rgbaMatrix [][][]uint8) {
	rgbaMatrix = New3DSlice(height, width, 4)
	return
}

func convertToNRGBA(src image.Image) *image.NRGBA {
	srcBounds := src.Bounds()
	dstBounds := srcBounds.Sub(srcBounds.Min)
	dst := image.NewNRGBA(dstBounds)
	dstMinX := dstBounds.Min.X
	dstMinY := dstBounds.Min.Y
	srcMinX := srcBounds.Min.X
	srcMinY := srcBounds.Min.Y
	srcMaxX := srcBounds.Max.X
	srcMaxY := srcBounds.Max.Y

	switch src0 := src.(type) {
	case *image.NRGBA:
		rowSize := srcBounds.Dx() * 4
		numRows := srcBounds.Dy()
		i0 := dst.PixOffset(dstMinX, dstMinY)
		j0 := src0.PixOffset(srcMinX, srcMinY)
		di := dst.Stride
		dj := src0.Stride
		for row := 0; row < numRows; row++ {
			copy(dst.Pix[i0:i0+rowSize], src0.Pix[j0:j0+rowSize])
			i0 += di
			j0 += dj
		}
	case *image.RGBA:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {
				j := src0.PixOffset(x, y)
				a := src0.Pix[j+3]
				dst.Pix[i+3] = a
				switch a {
				case 0:
					dst.Pix[i+0] = 0
					dst.Pix[i+1] = 0
					dst.Pix[i+2] = 0
				case 0xff:
					dst.Pix[i+0] = src0.Pix[j+0]
					dst.Pix[i+1] = src0.Pix[j+1]
					dst.Pix[i+2] = src0.Pix[j+2]
				default:
					dst.Pix[i+0] = uint8(uint16(src0.Pix[j+0]) * 0xff / uint16(a))
					dst.Pix[i+1] = uint8(uint16(src0.Pix[j+1]) * 0xff / uint16(a))
					dst.Pix[i+2] = uint8(uint16(src0.Pix[j+2]) * 0xff / uint16(a))
				}
			}
		}
	default:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {
				c := color.NRGBAModel.Convert(src.At(x, y)).(color.NRGBA)
				dst.Pix[i+0] = c.R
				dst.Pix[i+1] = c.G
				dst.Pix[i+2] = c.B
				dst.Pix[i+3] = c.A
			}
		}
	}
	return dst
}
