package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"encoding/base64"
	"encoding/json"
)
type ResponseStruct struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
} 
type ResponseData struct{
	Value string `json:"value"`
}
var tempValue string;

func saveTempValue(temp string){
	tempValue = temp;
}
func getTempValue() string {

	return tempValue
}

func updateTempValue(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println("path",r.URL.Path)
	if r.Method == "GET"{
		var sendValue string 
		sendValue = getTempValue()
		fmt.Println("temp control value ：",sendValue)
		d := ResponseData{
			Value: sendValue,
		}
		s := ResponseStruct{
			Code : 200,
			Message : "recieve sucess ",
			Data : d,
		}
		b , _ := json.Marshal(s)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	}else{
		//fmt.Println("Heard:",r.Header)
		//fmt.Println("Body:",r.Body)
		tempValueBuff := r.Form["temp"]

		tempValue = string(tempValueBuff[0])
		fmt.Println("temp value:",tempValue)
		saveTempValue(tempValue)
		
		fmt.Fprintf(w,"submit sucess ,temp value : %s",tempValue)
	}
	

}

func recieveImage(w http.ResponseWriter , r *http.Request){
	r.ParseForm()
	fmt.Println("path",r.URL.Path)
	imagebuff := r.Form["image"]
	formatTimeStr:=time.Now().Format("2006-01-02-15-04-05")
	fmt.Println(formatTimeStr)   //打印结果：2017-04-11 13:30:39	
	dataTimeString := "static/"+string(formatTimeStr)+".png"
	fmt.Println(dataTimeString)
	imageString := string(imagebuff[0])
	//fmt.Println(imageString)
	base2Img(imageString,dataTimeString)
	fmt.Fprintf(w,"submit sucess!")
}
func base2Img(imageString string,filename string){
	ddd, _ := base64.StdEncoding.DecodeString(imageString) 
	err :=ioutil.WriteFile(filename, ddd, 0666) 
	fmt.Println(err)
}
func welcomeUse(w http.ResponseWriter , r *http.Request){
	r.ParseForm()
	fmt.Println("path",r.URL.Path)
	fmt.Fprintf(w,"http server!")
}
func main(){
	fmt.Println(" \t welcome to use  http server \t")
	fmt.Println("\t\t\t\t\t深圳市德沃仪器\t")
	http.HandleFunc("/",welcomeUse)
	http.HandleFunc("/image/upload",recieveImage)
	http.HandleFunc("/standard_value/update",updateTempValue)
	http.HandleFunc("/standard_value/low",updateTempValue)
	http.HandleFunc("/standard_value/mediume",updateTempValue)
	http.HandleFunc("/standard_value/high",updateTempValue)
	err := http.ListenAndServe(":8000",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}