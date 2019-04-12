package main

import (

	//parse JSON

	"fmt" //file work
	"io"
	"os"
	_ "path"

	"github.com/fogleman/gg"
	geojson "github.com/paulmach/go.geojson"

	//http work
	//convert string to int

	_ "mime/multipart"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//file, err := os.Open("map2.geojson")
	file, err := os.Open("map.geojson")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	data := make([]byte, stat.Size())
	if err != nil {
		return
	}
	for {
		/* n */ _, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		//	fmt.Print(string(data[:n]))
	}

	fc, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	const S = 100
	dc := gg.NewContext(1366, 1024)
	for i := 0; i < len(fc.Features); i++ {

		p := fc.Features[i].Geometry.Polygon[0]
		dc.Push()

		dc.Scale(5, 5)
		for i := 0; i < len(p); i++ {
			dc.LineTo(p[i][0], p[i][1]+85)
		}
		dc.SetLineWidth(10)
		fmt.Println(fc.Features[i].Properties["color"])
		switch fc.Features[i].Properties["color"] {
		case "green":
			dc.SetRGBA255(91, 255, 15, 255)
			dc.StrokePreserve()
			dc.SetRGBA255(91, 155, 15, 255)

		case "orange":
			dc.SetRGBA255(255, 184, 5, 255)
			dc.StrokePreserve()
			dc.SetRGBA255(200, 184, 5, 255)
		default:
			dc.SetRGBA255(255, 255, 255, 255)
			dc.StrokePreserve()
			dc.SetRGBA255(0, 0, 0, 255)
		}

		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG("out.png")

}
