package sport

import (
	"Project/logger"
	"bufio"
	"encoding/json"
	"errors"

	//"fmt"

	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	//"github.com/spf13/viper"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Details struct {
	Name   string
	Age    int
	Gender string
	Sport  []string
	Height float64
	Weight float64
}

type MetaJson struct {
	Json []Details
}
type MetaYaml struct {
	Yaml []Details
}

type DetailWriter interface {
	WriteToFile(file string) error
}

func (m *MetaJson) WriteToFile(file string) error {

	b, err := json.MarshalIndent(m.Json, "", "")
	if err != nil {
		logger.Error("Error in Json Marshalling ", zap.Error(err))
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		logger.Error("Error in writing a json file", zap.String("fileName", file))
		return err
	}
	logger.Info("Successfully Created json file", zap.String("fileName", file))
	return nil
}

func (m *MetaYaml) WriteToFile(file string) error {

	d, err := yaml.Marshal(m.Yaml)
	if err != nil {
		log.Fatalf("error: %v", err)
		logger.Error("Error in Marshalling yaml", zap.Error(err))
		return err
	}

	err = ioutil.WriteFile(file, d, 0644)
	if err != nil {
		logger.Error("Error in Writing yaml", zap.String("fileName", file))
		return err
	}
	logger.Info("Successfully Created yaml file", zap.String("fileName", file))
	return nil

}

func CreateFile(item DetailWriter, file string) error {
	err := item.WriteToFile(file)
	return err
}

func Getinfo(Slice *[]Details, filename string) ([]Details, error) {
	logger.Info("Inside GetInfo()")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		logger.Error("Error in opening file", zap.String("fileName", filename))
		return *Slice, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	if scanner.Scan() == true {

		for scanner.Scan() {

			//count = count + 1
			//fmt.Println(count)

			s := strings.Split(scanner.Text(), "[")
			//fmt.Println("lin", s)
			k := strings.Split(s[0], ",")
			l := strings.Split(s[1], "]")
			m := strings.Split(l[1], ",")
			p := strings.Split(l[0], ",")
			//fmt.Println(k)
			//fmt.Println(l[0])
			//fmt.Println(m)
			var detail Details
			detail.Name, detail.Gender = k[0], k[2]
			detail.Age, err = strconv.Atoi(k[1])
			if err != nil {
				logger.Error("Error in Assigning age", zap.Error(err))
				return *Slice, err
			}
			detail.Sport = p
			//fmt.Println(p)
			detail.Height, err = strconv.ParseFloat(m[1], 32)
			if err != nil {
				logger.Error("Error in Assigning Height", zap.Error(err))
				return *Slice, err
			}
			detail.Weight, err = strconv.ParseFloat(m[2], 32)
			if err != nil {
				logger.Error("Error in Assigning Weight", zap.Error(err))
				return *Slice, err
			}

			*Slice = append(*Slice, detail)

		}
	} else {
		return *Slice, errors.New("Scanner not true")
	}

	return *Slice, nil

}
