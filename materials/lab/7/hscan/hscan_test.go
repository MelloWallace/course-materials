// Optional Todo

package hscan

import (
	"testing"
)

// func TestGuessSingle(t *testing.T) {
// 	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f","/home/cabox/workspace/materials/lab/7/main/wordlist.txt") // Currently function returns only number of open ports
// 	want := "Nickelback4life"
// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
//
// }

// func TestGenHashMapsMostSec(t *testing.T){
// 	 GenHashMaps("/home/cabox/workspace/materials/lab/7/main/most_security.txt")
// }

// func TestGenHashMapsTop304(t *testing.T){
// 	GenHashMaps("/home/cabox/workspace/materials/lab/7/main/Top304Thousand.txt")
// }

func TestGenHashMaps10mil(t *testing.T){
	GenHashMaps("/home/cabox/workspace/materials/lab/7/main/10_million_password_list_top_100000.txt")
}