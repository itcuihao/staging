package q2

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

func getBGImage(url string) image.Image {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return image.NewNRGBA(image.Rect(0, 0, 600, 400))
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	contentType := http.DetectContentType(data)
	var imgBG image.Image
	if contentType == "image/png" {
		imgBG, err = png.Decode(bytes.NewReader(data))
	} else if contentType == "image/jpeg" {
		imgBG, err = jpeg.Decode(bytes.NewReader(data))
	} else if contentType == "image/gif" {
		imgBG, err = gif.Decode(bytes.NewReader(data))
	}
	if err != nil {
		log.Println(err)
		return image.NewNRGBA(image.Rect(0, 0, 600, 400))
	}
	return imgBG
}

// 初始化字体文件
func InitSty() {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	fontStyle, err = freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
}

var (
	// fontFile = "config/luxisr.ttf"
	// fontFile = "config/seguiemj.ttf"
	fontFile  = "config/PingFang_Medium.ttf"
	fontStyle *truetype.Font
	fontSize  float64 = 14
	fontDPI   float64 = 72
)

func textCornerMark(text []string, img *image.NRGBA) *image.NRGBA {
	InitSty()

	// fg := image.Black
	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(fontStyle)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	// c.SetSrc(fg)
	// 设置字体颜色
	c.SetSrc(image.NewUniform(color.RGBA{R: 240, G: 240, B: 245, A: 180}))

	pt := freetype.Pt(100, 10+int(c.PointToFixed(fontSize)>>6))
	for _, s := range text {
		_, err := c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
		}
		pt.Y += c.PointToFixed(fontSize * 1.5)
	}

	return img
}

func getImgUrl(url string) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, 250, 250))
	textCornerMarkCorner := getBGImage(url)
	draw.Draw(img, textCornerMarkCorner.Bounds(), textCornerMarkCorner, image.ZP, draw.Src)
	return img
}

func getImg(fname string) *image.NRGBA {
	flag.Parse()
	if fname == "" {
		fname = out
	}
	// 打开图片
	logoImage, err := openImage(fname)
	if err != nil {
		return nil
	}
	// 图片大小
	img := image.NewNRGBA(image.Rect(0, 0, 250, 250))
	draw.Draw(img, logoImage.Bounds(), logoImage, image.ZP, draw.Src)
	return img
}

func outFile(out string, srcImage image.Image) {
	outAbs, err := filepath.Abs(out)
	if err != nil {
		log.Fatalf("获取输出文件绝对路径发生错误：%s", err.Error())
	}

	os.MkdirAll(filepath.Dir(outAbs), 0777)
	outFile, err := os.Create(outAbs)
	if err != nil {
		log.Fatalf("创建输出文件发生错误：%s", err.Error())
	}
	defer outFile.Close()

	jpeg.Encode(outFile, srcImage, &jpeg.Options{Quality: 100})
	// png.Encode(outFile, srcImage)
	log.Printf("二维码生成成功，文件路径：%s", outAbs)
}

var (
	text    string
	logo    string
	percent int
	size    int
	out     string
)

func init() {
	flag.StringVar(&text, "t", "https://www.putaoabc.com/onlinepages/activity/invite_friend/invite_message.html?fromPutaoId=21581", "二维码内容")
	flag.StringVar(&logo, "l", "1.jpg", "二维码Logo(png)")
	flag.IntVar(&percent, "p", 15, "二维码Logo的显示比例(默认15%)")
	flag.IntVar(&size, "s", 250, "二维码的大小(默认256)")
	flag.StringVar(&out, "o", "images/1.jpg", "输出文件")
}

func haibao() {
	flag.Parse()

	if text == "" {
		log.Fatalf("请指定二维码的生成内容")
	}

	if out == "" {
		log.Fatalf("请指定输出文件")
	}

	code, err := qrcode.New(text, qrcode.Highest)
	if err != nil {
		log.Fatalf("创建二维码发生错误：%s", err.Error())
	}

	srcImage := code.Image(size)
	if logo != "" {
		// logoSize := float64(size) * float64(percent) / 100

		srcImage, err = addQrcode(srcImage, logo, 256)
		// srcImage, err = addLogo(srcImage, logo, int(logoSize))
		if err != nil {
			log.Fatalf("增加Logo发生错误：%s", err.Error())
		}
	}

	outFile(out, srcImage)
}

func checkFile(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func openImage(name string) (image.Image, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func addQrcode(srcImage image.Image, logo string, size int) (image.Image, error) {
	// logoImage, err := resizeLogo(logo, uint(size))
	// 要添加的logo图
	logoImage, err := openImage(logo)
	if err != nil {
		return nil, err
	}

	// 坐标起点，进行覆盖
	x := 2*((logoImage.Bounds().Dx())/3) - 20
	y := 4 * ((logoImage.Bounds().Dy()) / 5)
	offset := image.Pt(x, y)

	b := logoImage.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, logoImage, image.ZP, draw.Src)
	draw.Draw(m, srcImage.Bounds().Add(offset), srcImage, image.ZP, draw.Over)
	return m, nil
}

func drawCircle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// 在base上画img的原型图像
func cropCircle(base, img image.Image, p image.Point, r int) image.Image {
	// 坐标起点，进行覆盖
	bx := base.Bounds().Dx()
	by := base.Bounds().Dy()
	log.Println(bx, by)
	x := (base.Bounds().Dx()) / 30
	y := (base.Bounds().Dy()) / 40
	offset := image.Pt(x, y)
	fmt.Println(offset)
	// 底图
	ib := base.Bounds()
	in := image.NewRGBA(ib)
	draw.Draw(in, ib, base, image.ZP, draw.Src)

	// 在底图上画图
	draw.DrawMask(in, base.Bounds().Add(offset), img, image.ZP, &circle{p, r}, image.ZP, draw.Over)
	return in
}

// 缩略图
func thumbnail(origin image.Image, width, height, quality int) image.Image {
	if width == 0 || height == 0 {
		width = origin.Bounds().Max.X
		height = origin.Bounds().Max.Y
	}
	if quality == 0 {
		quality = 100
	}
	thumb := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)
	return thumb
}
