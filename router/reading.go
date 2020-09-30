package router

import (
	"fmt"
	"os"
	"strings"
	//"encoding/json"
	//. "fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
)



func GetClassBookList(c *gin.Context)  {
	//返回data的格式
	type Data struct{
		Id string `json:"id"`
		Name string `json:"name"`
		Image string `json:"image"`
		//content string
		Writer string `json:"writer"`
	}
	bookclass :=c.Query("bookclass")
	page :=c.Query("page")
	url := "http://www.mingzw.net/mzwlist/"+bookclass+"_"+page+".html"
	fmt.Fprint(os.Stdout,url)
	doc, err := goquery.NewDocument(url)
	if(err!=nil){
		if(err==nil) {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":-1,
				"msg":"failed",
				"data":make([]Data,0),
			})
		}

	}


	respondData :=[]Data{}
	doc.Find("div[class=bd]").Find("div[class='figure-horizontal figure-1']").EachWithBreak(func(i int, selection *goquery.Selection)bool {

		pic,picerr:= selection.Find("div[class=pic]").Find("img").Eq(0).Attr("src")
		if(picerr){
			fmt.Fprint(os.Stdout,picerr)
		}

		pic="http://www.mingzw.net"+strings.Replace(pic,"http://www.mingzw.net","",1)

		writer :=selection.Find("div.cont").Find("dl").Eq(0).Find("dd").Text()
		name :=selection.Find("div.cont").Find("h3").Text()
		id,iderr :=selection.Find("div.cont").Find("a").Attr("href")
		if(iderr) {
		}
		id=strings.Replace(strings.Replace(id,"/mzwbook/","",1),".html","",1)

		tempdata :=Data{
			id,
			name,
			pic,
			writer,
		}
		respondData=append(respondData,tempdata)
		return true
	})







	c.JSON(http.StatusBadRequest,gin.H{
		"status":0,
		"msg":"succeed",
		"data":respondData,
	})


}



func GetBookClass(c *gin.Context) {

	doc, err := goquery.NewDocument("http://www.mingzw.net/")
	classlist :=[]string{}
	if(err==nil) {
		doc.Find("#categories").Find("div.bd").Find("a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
			classlist= append(classlist, selection.Text())
			return true
		})

		c.JSON(http.StatusOK,gin.H{
			"status":0,
			"msg":"succeed",
			"data":classlist,
		})
	}
}



func GetCharther(c *gin.Context) {
	type Data struct{
		Page string `json:"page"`
		Title string `json:"title"`
	}
	respondata := []Data{}
	id :=c.Query("bookid")
	url := "http://www.mingzw.net/mzwchapter/"+id+".html"
	doc, err := goquery.NewDocument(url)
	if(err!=nil) {
		fmt.Fprint(os.Stdout,err)
	}
	doc.Find("div[class='content gclearfix']").Find("ul[class=gclearfix]").Find("li[id!=addbookshelf_1]").Each(func(i int, selection *goquery.Selection) {

		page,status :=selection.Find("a").Attr("href")
		if(status==true) {
			page = strings.Replace(strings.Replace(page, "/mzwread/", "", 1), ".html", "", 1)
		}
		title:=selection.Find("a").Text()
		respondata=append(respondata,Data{page,title})

	})
	c.JSON(http.StatusOK, gin.H{
		"msg":"succeed",
		"staus":0,
		"bookid":id,
		"data":respondata,
	})


}




func GetBookContent(c *gin.Context) {
	type Data struct{
		Content string `json:"content"`
		NextID string 	`json:"next_id"`
		PerivousID string `json:"perivous_id"`
		BookID string `json:"book_id"`
	}

	id :=c.Query("pageid")
	BookID :=strings.Split(id,"_")[0]
	url := "http://www.mingzw.net/mzwread/"+id+".html"
	fmt.Fprint(os.Stdout,url)

	doc, err := goquery.NewDocument(url)
	if(err!=nil) {
		fmt.Fprint(os.Stdout,err)
	}
	Next,Nextstatus:=doc.Find("a:contains('下一章')").Attr("href")
	if(Nextstatus==true){
		Next=strings.Replace( strings.Replace(Next,"/mzwread/","",1),".html","",1)
	}
	Perivous,Pervioustatus:=doc.Find("a:contains('下一章')").Attr("href")

	doc.Find("div[class=content]").ChildrenFiltered("div").Remove()
	Content :=doc.Find("div[class=content]").Text()

	if(Pervioustatus==true){
		Perivous=strings.Replace( strings.Replace(Perivous,"/mzwread/","",1),".html","",1)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":"succeed",
		"staus":0,
		"data":Data{
			Content,
			Next,
			Perivous,
			BookID},
	})


}



