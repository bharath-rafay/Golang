package sport

import (
	"testing"
)

func TestCreate(t *testing.T) {
	var tmp1 Details

	tmp := make([]Details, 0)
	// tmp2 := make([]Details.Sport,0)

	tmp1.Name = "Bharath"
	tmp1.Age = 22
	tmp1.Gender = "M"

	tmp1.Sport = []string{"cricket", "hockey"}

	tmp1.Height = 5.10
	tmp1.Weight = 77

	tmp = append(tmp, tmp1)
	var js MetaJson
	js.Json = tmp
	got := CreateFile(&js, "testingfile.json")

	if got != nil {
		t.Errorf("want:%v,got:%v", nil, got)
	}

}

func TestGet(t *testing.T) {

	tmp := make([]Details, 0)
	file := "/home/bharat/Downloads/in.txt"

	_, err := Getinfo(&tmp, file)

	if err != nil {
		t.Errorf("want:%v,got:%v", nil, err)
	}

}
