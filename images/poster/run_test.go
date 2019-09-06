package q2

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestImageURL(t *testing.T) {
	s := []string{
		"Hello,world!",
		"I am 葱葱.",
		"快乐的葱葱",
		"",
		"葱葱给猴子♥♥",
	}
	url := "https://preview.qiantucdn.com/element_origin_pic/18/04/12/auto_11f3def5ed73b1160561d210050b9ce5_PIC2018.jpg!qt324new_nowater"
	img := getImgUrl(url)
	m := textCornerMark(s, img)
	outFile("2.jpg", m)
}
func TestImage(t *testing.T) {
	s := []string{
		"Hello,world!",
		"I am 葱葱.",
		"快乐的葱葱",
		"",
		"葱葱给猴子♥♥",
		"λ⺪☺あいえたせちみほふかきひ【pic】(ˆoˆԅ)(【pic】˙0˙)",
	}
	img := getImg("2.png")
	m := textCornerMark(s, img)
	outFile("2.jpg", m)
}
func TestHaibao(t *testing.T) {
	haibao()
}

func BenchmarkHaibao(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		haibao()
	}
}
func BenchmarkHaibaottf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []string{
			"Hello,world!",
			"I am 葱葱.",
			"快乐的葱葱",
			"",
			"葱葱给猴子♥♥",
		}
		url := "https://preview.qiantucdn.com/element_origin_pic/18/04/12/auto_11f3def5ed73b1160561d210050b9ce5_PIC2018.jpg!qt324new_nowater"
		img := getImgUrl(url)
		m := textCornerMark(s, img)
		outFile("2.jpg", m)
	}
}
func BenchmarkHaibaol(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []string{
			"Hello,world!",
			"I am 葱葱.",
			"快乐的葱葱",
			"",
			"葱葱给猴子♥♥",
		}
		img := getImg("2.png")
		m := textCornerMark(s, img)
		outFile("2.jpg", m)
	}
}

func TestDrawCircle(t *testing.T) {
	img := getImg("2.png")
	drawCircle(img, 40, 40, 30, color.RGBA{255, 0, 0, 255})
	outFile("3.jpg", img)
}
func TestCircle(t *testing.T) {
	img, _ := openImage("4.jpg")
	base, _ := openImage("1.jpg")
	ccImg := cropCircle(base, img, image.Point{125, 125}, 100)
	outFile("5.jpg", ccImg)
}

func TestThumb(t *testing.T) {
	img, _ := openImage("4.jpg")
	thumb := thumbnail(img, 50, 50, 100)
	outFile("5t.jpg", thumb)
}

func TestCircleThumb(t *testing.T) {
	// 等待画圆的图
	img, _ := openImage("4.jpg")
	// 等待画圆的缩略图
	thumb := thumbnail(img, 120, 120, 100)
	thx := thumb.Bounds().Dx()
	thy := thumb.Bounds().Dy()
	// 要画的圆的半径
	var r int
	if thx > thy {
		r = thy / 2
	} else {
		r = thx / 2
	}
	// 要画的圆心的坐标
	point := image.Point{thx / 2, thy / 2}
	outFile("6t.jpg", thumb)
	fmt.Println(thumb.Bounds().Dx(), thumb.Bounds().Dy())
	// 在base上画图
	base, _ := openImage("1.jpg")

	ccImg := cropCircle(base, thumb, point, r)
	outFile("6.jpg", ccImg)
}
