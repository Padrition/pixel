package pixel

import(
	"math/rand"
	"image/color"
	"time"
	"image"
	"os"
	"image/png"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomColor() color.NRGBA{
	return color.NRGBA{R : uint8(rand.Intn(255)), G : uint8(rand.Intn(255)), B :  uint8(rand.Intn(255)), A: 0xff}
}

func CreateImgWithRandFilling(width, height int){
	upleft := image.Point{0,0}
	downright := image.Point{width,height}

	img := image.NewRGBA(image.Rectangle{upleft, downright})
	
	for x:= 0; x <width; x++{
		for y:=0; y<height; y++{
			r := randomColor()
			img.Set(x,y,r)
		}
	}

	f, _ := os.Create("test.png")
	png.Encode(f, img)
}
