package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

type Images struct {
	Origin string `json:"origin"`
	Crop   string `json:"crop"`
}

func main() {
	img := &Images{Origin: "wqwqw", Crop: "dsdssdsd"}
	img1 := &Images{Origin: "wewewe", Crop: "ffddfdf"}

	imgs := make([]*Images, 0, 2)
	imgs = append(imgs, img)
	imgs = append(imgs, img1)

	urls, err := json.Marshal(imgs)
	if err != nil {
		panic(err)
	}
	// spew.Dump(string(urls))

	imgURLs := string(urls)
	var ims []*Images
	err = json.Unmarshal([]byte(imgURLs), &ims)
	if err != nil {
		panic(err)
	}
	// spew.Dump(ims)

	mergeImageURL()

}

// type Images struct {
// 	Origin string `json:"origin"` -> OriginImageURLs
// 	Crop   string `json:"crop"` -> ImageURLs
// }

func mergeImageURL() {
	origin := []string{"assassin", "dsdssd", "gfgfgfgf", "更换合格合格"}
	crop := []string{"哈哈哈", "大大方方的", "对方答复"}
	imgs := make([]*Images, 0)

	if len(origin) >= len(crop) {
		for i, o := range origin {
			img := &Images{Origin: o}
			if i >= len(crop) {
				img.Crop = ""
			} else {
				img.Crop = crop[i]
			}
			imgs = append(imgs, img)
		}
	} else {
		for i, c := range crop {
			img := &Images{Crop: c}
			if i >= len(origin) {
				img.Origin = ""
			} else {
				img.Origin = origin[i]
			}
			imgs = append(imgs, img)
		}
	}

	spew.Dump(imgs)
	return
}
