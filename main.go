package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "syscall/js"
    "math/rand"
    "time"
)


type Size struct {
    Width float64 `json:"width"`
    Height float64 `json:"height"`
}

type TileFormat struct {
    Size Size `json:"size"`
}

type Top struct {
//     Version string `json:"version"`
    Tiles []Tile `json:"tiles"`
    TileFormat TileFormat `json:"tileFormat"`
}

type Tile struct {
    X int64 `json:"x"`
    Y int64 `json:"y"`
    Image string `json:"image"`
}

func getRandomNum() float32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := float32(rand.Intn(10000))
	return  n / 10000.0
}

func main() {
    fmt.Println("Hello, WebAssembly!")

    req, _ := http.NewRequest("GET", "http://mock.puddle-sketch.com/sample_paper_puddly/", nil)
    req.Header.Set("Cache-Control", "no-cache")

    client := new(http.Client)
    resp, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    var top Top
    unmarshallErr := json.Unmarshal(body, &top)
    if unmarshallErr != nil {
        fmt.Println("error:", unmarshallErr)
        return
    }

    for _, tile := range top.Tiles {
        fmt.Printf("x: %i, y: %i, image: %s\n",
            tile.X, tile.Y, tile.Image)
    }

    var width = 400
    var height = 800
// 	var width  = js.Global().Get("innerWidth")
// 	var height = js.Global().Get("innerHeight")
    doc := js.Global().Get("document")
    canvasEl := doc.Call("getElementById", "mycanvas")
	canvasEl.Call("setAttribute", "width", width)
	canvasEl.Call("setAttribute", "height", height)
    ctx := canvasEl.Call("getContext", "2d")

	ctx.Call("clearRect", 0, 0, width, height)

	for i := 0; i < 1; i ++ {
		ctx.Call("beginPath")
		ctx.Call("moveTo", 200, 400)
		ctx.Call("lineTo", 300, 500)
		ctx.Call("stroke")
	}

//     image := doc.Call("getElementById", "myImage")
    image := doc.Call("createElement", "img")
    image.Call("setAttribute", "src", "https://pbs.twimg.com/media/EB6QRvsX4AE_YMO.jpg")
    fmt.Println("image")
    fmt.Println(image)

//     ctx.Call("drawImage", image, 0, 0, 100, 100, 0, 0, 100, 100)
    ctx.Call("drawImage", image, 0, 0, nil, nil, 0, 0, nil, nil)
}
