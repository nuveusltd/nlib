package nlib

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"reflect"
)

// SaveToNSF Encrypted bicimde Nuveus formatinda objeti kaydeder.
func SaveToNSF(filename string, obj interface{}) {
	bin := new(bytes.Buffer)
	encoder := gob.NewEncoder(bin)
	encoder.Encode(obj)
	encoded := Encrypt(bin.Bytes(), []byte(GetMD5Hash("NSF.Object.2021")))
	ioutil.WriteFile(filename, encoded, 0644)
}

// LoadFromNSF Encrypted bicimde Nuveus formatinda kaydedilmis objeyi doner
func LoadFromNSF(filename string, o interface{}) interface{} {

	v := reflect.New(reflect.TypeOf(o))

	bin, _ := ioutil.ReadFile(filename)
	dec := Decrypt(bin, []byte(GetMD5Hash("NSF.Object.2021")))
	buf := new(bytes.Buffer)
	buf.Write(dec)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(v.Interface())
	return v.Interface()
}
