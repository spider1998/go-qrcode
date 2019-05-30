package main

import (
	"fmt"
	q "github.com/skip2/go-qrcode"
	"github.com/tuotoo/qrcode"
	"image/color"
	"os"
)

func main() {
	err := MakeQRcode("ghfdiugsduhfldsgfiugrfbds","./spider1998.png",256,color.White,color.Black)
	if err != nil{
		fmt.Println(err)
	}
	content,err := ScanQRcode("./spider1998.png")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(content)

}

//生成
func MakeQRcode(content,fileName string,size int,backGroundColor color.Gray16,ForegroundColor color.Gray16)(err error)  {
	qr,err:=q.New(content,q.Medium)
	if err != nil {
		return
	} else {
		qr.BackgroundColor = backGroundColor
		qr.ForegroundColor = ForegroundColor
		err = qr.WriteFile(size, fileName)
		if err != nil{
			return
		}
	}
	return nil
}


//扫描
func ScanQRcode(fileName string)(content string,err error)  {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	content = qrmatrix.Content
	return
}