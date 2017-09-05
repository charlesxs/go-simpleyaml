package simpleyaml

import (
	"github.com/go-yaml/yaml"
	"os"
	"log"
	"bytes"
	"bufio"
	"io"
	"errors"
	"io/ioutil"
)


// type Yaml
type Yaml struct {
	data interface{}
}


func Version() string {
	return "0.1"
}


func readAll(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	defer f.Close()

	var b bytes.Buffer
	nr := bufio.NewReader(f)
	for {
		line, err := nr.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		b.WriteString(line)
	}

	return b.Bytes()
}

func NewYaml(d interface{}) *Yaml {
	return &Yaml{
		data: d,
	}
}

func New() *Yaml {
	return &Yaml{
		data: make(map[interface{}]interface{}),
	}
}


// load yaml from string
func (y *Yaml) Loads(b []byte) error {
	return yaml.Unmarshal(b, y.data)
}

// load yaml from file
func (y *Yaml) Load(filename string) error {
	return yaml.Unmarshal(readAll(filename), y.data)
}

// dump yaml to string
func (y *Yaml) Dumps() ([]byte, error)  {
	return yaml.Marshal(y.data)
}

// dump yaml to file
func (y *Yaml) Dump(filename string) error {
	data, err := yaml.Marshal(y.data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}


func (y *Yaml) Map() (map[interface{}]interface{}, error) {
	v, ok := y.data.(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("assert to map failed")
	}
	return v, nil
}

func (y *Yaml) Slice() ([]interface{}, error) {
	if v, ok := (y.data).([]interface{}); ok {
		return v, nil
	}
	return nil, errors.New("Not slice")
}

func (y *Yaml) Get(key string) *Yaml  {
	m, err := y.Map()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if v, ok := m[key]; ok {
		return &Yaml{v}
	}
	return nil
}

func (y *Yaml) GetIndex(index int) *Yaml {
	s, err := y.Slice()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if len(s) > index {
		return &Yaml{s[index]}
	}
	return nil

}

func (y *Yaml) Keys() []interface{} {
	m, err := y.Map()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	s := make([]interface{}, len(m))
	var i int

	for k := range m {
		s[i] = k
		i++
	}
	return s
}

func (y *Yaml) Set(key string, v interface{}) {
	m, err := y.Map()
	if err != nil {
		return
	}
	m[key] = v
}


func (y *Yaml) Int() (int, error) {
	if v, ok := y.data.(int); ok {
		return v, nil
	}
	return -1, errors.New("it's not int")
}

func (y *Yaml) String() (string, error)  {
	if v, ok := y.data.(string); ok {
		return v, nil
	}
	return "", errors.New("it's not string")
}

func (y *Yaml) Bool() (bool, error)  {
	if v, ok := y.data.(bool); ok {
		return v, nil
	}
	return false, errors.New("it's not bool type")
}

func (y *Yaml) Float64() (float64, error)  {
	if v, ok := y.data.(float64); ok {
		return v, nil
	}
	return 0, errors.New("it's not float64")
}

