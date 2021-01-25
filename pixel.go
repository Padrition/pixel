package pixel

import(
	"math/rand"
	"image/color"
	"time"
	"image"
)

type PixelImage struct {
	Width 		uint8
	Height 		uint8
	PixelBlockSize 	uint8
	NumOfColors 	uint8
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomColor() color.NRGBA{
	return color.NRGBA{R : uint8(rand.Intn(255)), G : uint8(rand.Intn(255)), B :  uint8(rand.Intn(255)), A: 0xff}
}

func CreatePixeledImg(pixImg PixelImage) image.Image{
	
	width 		:= int(pixImg.Width)
	height		:= int(pixImg.Height)
	blockSize	:= int(pixImg.PixelBlockSize)
	numOfColors	:= int(pixImg.NumOfColors)

	colors := make([]color.NRGBA, numOfColors)
	var infColor bool

	upleft := image.Point{0,0}
	downright := image.Point{width, height}
	
	img := image.NewRGBA(image.Rectangle{upleft, downright})

	if numOfColors > 1 {
		for i := 0; i < numOfColors; i++{
			colors[i] = randomColor()
		}
		infColor = false
	}else{
		infColor = true
	}

	for x := 0; x < width; x = x + blockSize{
		for y := 0; y < height; y = y + blockSize{
			if infColor == true{
				for ix := 0; ix < blockSize; ix ++{
					for iy := 0; iy < blockSize; iy ++ {
						img.Set(x + ix, y + iy, randomColor())
					}
				}
			}else {
				c := colors[rand.Intn(numOfColors)]
				for ix := 0; ix < blockSize; ix ++{
					for iy := 0; iy < blockSize; iy ++ {
						img.Set(x + ix, y + iy, c)
					}
				}
			}
		}
	}

	return img
}
