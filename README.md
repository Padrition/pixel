# pixel

Pixel is a go module for creating pixeled images with various settings that allow you to choose parameters like:
* Size of an image
* Size of a space that will be filled with a color
* Numbers of random colors
* Particular colors that you want to randomly place on a canvas

## Instalation

```bash
git clone https://github.com/Padrition/pixel.git
```

## Usage 

After instaling the module you need to create an instance of PixelImage :

```go

	pix := pixel.PixelImage{width, height, sizeOfColorBlock, numberOfColors}

```
* width - sets the width of a picture
* height - sets the height of a picture
* sizeOfColorBlock - sets the area size wich single color will occupy(seted to 10 will make single color occupy 10 by 10 pixels area)
* numberOfColors - sets the limit on number of colors you want to be present on a canvas(seted to 0 or 1 will set no limit on number of colors)

After that use the PixelImage instace as a param for the CreatePixeledImg function :

```go

	img := pixel.CreatePixeledImg(pix)

```

This function returns image.Image.
