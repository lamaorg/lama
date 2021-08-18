package common

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

/*
	FileUtils expose function to ease writing files to disk in a common way
*/

type FileUtils interface {
	EncodeStruct(encoder *gob.Encoder, s struct{}) ([]byte, error)
	WriteEncodedStructToLocalFile(enc []byte, filename string)
	LoadStructFromFile(filename string) *struct{}
	DecodeStruct(decoder *gob.Decoder, s []byte) *struct{}
	LoadJsonFile(filename string) []byte
	WriteJsonFile(filename string, data []byte)
	WriteNewReceiptToDisk(filename string, data []byte)
}

type FU struct {
	Encoder *gob.Encoder
	Decoder *gob.Decoder
	Buf     bytes.Buffer
	FileUtils
}

func (f *FU) InitEncodeDecode() {
	f.Encoder = gob.NewEncoder(&f.Buf)
	f.Decoder = gob.NewDecoder(&f.Buf)

}

func (f *FU) EncodeStruct(s interface{}) ([]byte, error) {

	f.InitEncodeDecode()
	err := f.Encoder.Encode(s)
	if err != nil {
		return nil, err
	}
	return f.Buf.Bytes(), nil
}

func (f *FU) DecodeStruct(i interface{}) interface{} {

	err := f.Decoder.Decode(&i)
	if err != nil {
		return nil
	}

	return f.Buf.Bytes()

}

func (f *FU) WriteEncodedStructToLocalFile(i interface{}, filename string) bool {

	b, _ := f.EncodeStruct(i)
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return false
	}
	return true
}
