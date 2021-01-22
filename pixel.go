package pixel

import(
	"math/rand"
	"image/color"
	"time"
	"image"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomColor() color.NRGBA{
	return color.NRGBA{R : uint8(rand.Intn(255)), G : uint8(rand.Intn(255)), B :  uint8(rand.Intn(255)), A: 0xff}
}

func RandPixelsImg(width, height int) image.Image{
	upleft := image.Point{0,0}
	downright := image.Point{width,height}

	img := image.NewRGBA(image.Rectangle{upleft, downright})
	
	for x:= 0; x <width; x++{
		for y:=0; y<height; y++{
			r := randomColor()
			img.Set(x,y,r)
		}
	}

	return img
}

func RandFixedSizePixelsImg(width, height, pixSize int) image.Image{
	upleft := image.Point{0,0}
	downright := image.Point{width,height}

	img := image.NewRGBA(image.Rectangle{upleft, downright})
	
	for x:= 0; x <width; x = x + pixSize{
		for y:=0; y<height; y = y + pixSize{
			r := randomColor()
			for i:=0;i<pixSize;i++{
				for j:=0;j<pixSize;j++{
					img.Set(x+i,y+j,r)
				}
			}
		}
	}

	return img

}
