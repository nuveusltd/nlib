package nlib

import (
	//"os"
	"encoding/gob"
	//"io"
	"fmt"
	"bytes"
	"io/ioutil"
	"reflect"
)


func SaveToNSF(filename string, obj interface{} ) {
	bin := new(bytes.Buffer)
	encoder := gob.NewEncoder(bin)
	encoder.Encode(obj)
	encoded:=Encrypt(bin.Bytes())
	ioutil.WriteFile(filename,encoded,0644)
}


func LoadFromNSF(filename string, o interface {} ) interface {} {

	v:=reflect.New(reflect.TypeOf(o))

	bin,_:=ioutil.ReadFile(filename)
	dec:=Decrypt(bin)
	buf:=new(bytes.Buffer)
	buf.Write(dec)
	decoder := gob.NewDecoder(buf)
	err:=decoder.Decode(v.Interface())
	fmt.Printf("Val : %v\n ERR : %v",v,err)
	return v.Interface()

}
