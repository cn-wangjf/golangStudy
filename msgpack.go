package main

import (
	"log"
	"reflect"

	"github.com/ugorji/go/codec"
)

type BuyItem struct {
	ItemId int
	Num    int
}

func main() {
	var mh codec.MsgpackHandle
	mh.StructToArray = true
	v := BuyItem{
		ItemId: 387,
		Num:    1,
	}

	/*
		v := make(map[string]int)
		v["ItemId"] = 3
		v["Num"] = 1
	*/

	log.Printf("编码前：%v, %v", reflect.TypeOf(v), v)
	b := make([]byte, 32)
	enc := codec.NewEncoderBytes(&b, &mh)
	err := enc.Encode(v)
	if err != nil {
		log.Printf("error encoding %v to MessagePack: %v", v, err)
	}
	log.Printf("编码后：%v, %v,长度：%d", reflect.TypeOf(b), b, len(b))

	dec := codec.NewDecoderBytes(b, &mh)
	var v2 BuyItem
	err2 := dec.Decode(&v2)
	if err2 != nil {
		log.Printf("error dencoding: %v from MessagePack: %v ", v, err2)
	}

	log.Printf("解码后：%v, %v", reflect.TypeOf(v2), v2)
}
