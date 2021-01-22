# pixel

Pixel is a go module for creating pixeled images with various settings that allow you to choose parameters like:
*Size of an image
*Size of a space that will be filled with a color
*Numbers of random colors
*Particular colors that you want to randomly place on a canvas

##Instalation

```bash
git clone https://github.com/Padrition/pixel.git
```

##Usage 

To create a fixed size image with a pixels of random RGB color in it use:
```go
	img := pixel.CreateImgWithRandFill(width, height)
```
