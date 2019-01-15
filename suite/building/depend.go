package building

import (
	"go-interface/tools"
)

type create struct {
	Name          string             `json:"name"`
	Loc           map[string]string  `json:"loc"`
	Area          float64            `json:"area"`
	LayerNum      int32              `json:"layer_num"`
	UnderLayerNum int32              `json:"underlayer_num"`
	Coord         map[string]float64 `json:"coord"`
}

var bName = "go随机建筑" + tools.RandInt()

func newCreate(name string, loc map[string]string, area float64, layerNum, underlayerNum int32, coord map[string]float64) *create {
	return &create{
		Name:          name,
		Loc:           loc,
		Area:          area,
		LayerNum:      layerNum,
		UnderLayerNum: underlayerNum,
		Coord:         coord,
	}
}

func newDefaultCreate() *create {

	return &create{
		Name: bName,
		Loc: map[string]string{
			"province": "上海市",
			"city":     "上海市",
			"county":   "静安区",
			"addr":     "恒丰路329号",
		},
		Area:          100,
		LayerNum:      31,
		UnderLayerNum: 3,
		Coord: map[string]float64{
			"altitude":  122,
			"latitude":  32,
			"longitude": 0,
			"angle":     0,
		},
	}
}
