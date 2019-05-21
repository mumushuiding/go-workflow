package flow

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/mumushuiding/util"
)

// func TestSome(t *testing.T) {
// 	var step int
// 	fmt.Println(step)
// 	step++
// 	fmt.Println(step)
// }

func TestNode(t *testing.T) {
	var node = Node{}
	node.GetProcessConfigFromJSONFile()
	result, _ := util.ToJSONStr(node)
	fmt.Println(result)
	x, _ := hex.DecodeString(result)
	fmt.Printf("----%b", x)
}

// func TestNodeGenerateNodeInfos(t *testing.T) {
// 	var node = Node{}
// 	node.GetProcessConfigFromJSONFile()
// 	// result, _ := util.ToJSONStr(node)
// 	// fmt.Println(result)
// 	maps := make(map[string]string)
// 	maps["DDHolidayField-J2BWEN12__duration"] = "8"
// 	maps["DDHolidayField-J2BWEN12__options"] = "年假"
// 	list, err := ParseProcessConfig(&node, &maps)
// 	if err != nil {
// 		log.Printf("err:%v", err)
// 	}
// 	str, _ := util.ToJSONStr(util.List2Array(list))
// 	fmt.Println(str)
// }

// func BenchmarkTest(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		var node = Node{}
// 		node.GetProcessConfigFromJSONFile()
// 		// result, _ := util.ToJSONStr(node)
// 		// fmt.Println(result)
// 		maps := make(map[string]string)
// 		maps["DDHolidayField-J2BWEN12__duration"] = "8"
// 		maps["DDHolidayField-J2BWEN12__options"] = "年假"
// 		list, err := ParseProcessConfig(&node, &maps)
// 		if err != nil {
// 			log.Printf("err:%v", err)
// 		}
// 		str, _ := util.ToJSONStr(util.List2Array(list))
// 		fmt.Println(str)
// 	}
// }
