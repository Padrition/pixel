package pixel

import(
	"math/rand"
	"image/color"
	"time"
	"image"
)

type PixelImage struct {
	Width int
	Height int
	PixelBlockSize int
	NumOfColors int
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomColor() color.NRGBA{
	return color.NRGBA{R : uint8(rand.Intn(255)), G : uint8(rand.Intn(255)), B :  uint8(rand.Intn(255)), A: 0xff}
}

func CreatePixeledImg(pixImg PixelImage) image.Image{
	
	colors := make([]color.NRGBA, pixImg.NumOfColors)
	var infColor bool

	upleft := image.Point{0,0}
	downright := image.Point{pixImg.Width, pixImg.Height}
	
	img := image.NewRGBA(image.Rectangle{upleft, downright})

	if pixImg.NumOfColors > 1 {
		for i := 0; i <  pixImg.NumOfColors; i++{
			colors[i] = randomColor()
		}
		infColor = false
	}else{
		infColor = true
	}

	for x := 0; x < pixImg.Width; x = x + pixImg.PixelBlockSize{
		for y := 0; y < pixImg.Height; y = y + pixImg.PixelBlockSize{
			if infColor == true{
				for ix := 0; ix < pixImg.PixelBlockSize; ix ++{
					for iy := 0; iy < pixImg.PixelBlockSize; iy ++ {
						img.Set(x + ix, y + iy, randomColor())
					}
				}
			}else {
				c := colors[rand.Intn(pixImg.NumOfColors)]
				for ix := 0; ix < pixImg.PixelBlockSize; ix ++{
					for iy := 0; iy < pixImg.PixelBlockSize; iy ++ {
						img.Set(x + ix, y + iy, c)
					}
				}
			}
		}
	}

	return img
}
