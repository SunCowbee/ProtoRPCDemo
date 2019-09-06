package main

import (
	"ProtoRPCDemo/prototext"
	"fmt"
	"github.com/golang/protobuf/proto"
)
func main() {
	text := &prototext.Test{
		Name:"panda",
		Tizhong:[]int32{120,125,198,180,150,195},
		Shengao:180,
		Motto:"天行健，君子以自强不息，地势坤，君子以厚德载物",
	}

	fmt.Println(text)
	//proto编码
	data ,err :=proto.Marshal(text)
	if err !=nil{
		fmt.Println("编码失败")
	}
	//编码后的打印
	fmt.Println(data)

	newtext := &prototext.Test{}

	//proto的解码
	err =proto.Unmarshal(data,newtext)
	if err !=nil{
		fmt.Println("解码失败")
	}
	//解码后的打印
	fmt.Println(newtext)
	fmt.Println(newtext.String())

	fmt.Println("名字:",newtext.Name);
	fmt.Println("身高:",newtext.Shengao);
	fmt.Println("体重:",newtext.Tizhong);
	fmt.Println("格言:",newtext.Motto);





}
