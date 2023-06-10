package myduo

import (
	"bytes"
	"context"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"my-duo/internal/consts"
	"my-duo/internal/service"
	"my-duo/internal/utils"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type sMyDuo struct {
	ImgSize         image.Rectangle
	FontUnicode     *truetype.Font
	FontAsciiBold   *truetype.Font
	BackgroundImage image.Image
	RoundMaskImage  image.Image
}

func prepareFont(path string) *truetype.Font {
	tt, err := freetype.ParseFont(utils.GetResource("resource/public/resource/font/" + path))
	if err != nil {
		panic(err)
	}
	return tt
}

func prepareImage(path string) image.Image {
	im, _, err := image.Decode(
		bytes.NewReader(utils.GetResource(
			"resource/public/resource/image/" + path)),
	)
	if err != nil {
		panic(err)
	}
	return im
}

func init() {
	service.RegisterMyDuo(New())
}
func New() *sMyDuo {
	return &sMyDuo{
		ImgSize:         image.Rect(0, 0, 793, 793),
		FontUnicode:     prepareFont("OPPOSans-R.ttf"),
		FontAsciiBold:   prepareFont("DIN Next Rounded LT W05 Bold.ttf"),
		BackgroundImage: prepareImage("background-image.png"),
		RoundMaskImage:  prepareImage("round-mask.png"),
	}
}

func (sv *sMyDuo) Draw(ctx context.Context, elem consts.MyDuoElements) []byte {
	img := image.NewRGBA(sv.ImgSize)
	sv.drawBackground(img)
	sv.drawCharacter(img, elem.Character)

	posX := 80
	posY := 290
	bg := image.NewUniform(color.RGBA{75, 75, 75, 255})
	point := fixed.Point26_6{X: fixed.Int26_6(posX * 64), Y: fixed.Int26_6(posY * 64)}
	drawDst := image.NewRGBA(img.Bounds())
	drawer := &font.Drawer{
		Dst: drawDst,
		Src: bg,
		Face: truetype.NewFace(
			sv.FontAsciiBold,
			&truetype.Options{Size: 45}),
		Dot: point,
	}
	drawer.DrawString("1234")
	draw.Draw(img, img.Bounds(), drawDst, drawDst.Bounds().Min, draw.Over)

	// encode to bytes
	buff := new(bytes.Buffer)
	png.Encode(buff, img)
	return buff.Bytes()
}

func (sv *sMyDuo) drawBackground(img *image.RGBA) {
	draw.Draw(img, img.Bounds(), sv.BackgroundImage, sv.BackgroundImage.Bounds().Min, draw.Src)
}

func (sv *sMyDuo) drawCharacter(img *image.RGBA, character consts.MyDuoCharacters) {
	c := prepareImage("characters/" + character.ToString() + ".png")
	draw.Draw(img, img.Bounds(), c, c.Bounds().Min, draw.Over)
}
