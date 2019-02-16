package model

import (
	"github.com/jinzhu/copier"
	"library/repository"
)

type Book struct {
	Name        string
	Publication string
	Author      string
	Price       int
	Content     string
	Cover       string
}

func (b *Book) Add() {
	book := &repository.Book{}
	copier.Copy(&book,&b)
	repository.Add(book)
}

func (b *Book) ReadAll() (bLice []Book) {
	b1 := Book{
		"精通 Python｜運用簡單的套件進行現代運算 (Introducing Python: Modern Computing in Simple Packages)",
		"歐萊禮",
		"Bill Lubanovic 著、賴屹民 譯",
		780,
		"本書是 Bill Lubanovic 的傑作，先為你紮下深厚的程式設計基礎，再教你使用大量 Python 工具箱處理現實生活中的問題，本書絕對適合學習如何運用 Python 來解決問題。",
		"/assets/static/image/b1.jpg",
	}
	b2 := Book{
		"演算法技術手冊, 2/e (Algorithms in a Nutshell: A Practical Guide, 2/e)",
		"歐萊禮",
		"George T. Heineman等 陳仁和 譯",
		580,
		"「這本書有三項值得閱讀的理由：針對書中演算法與資料結構，以視覺化圖表展示其特性；內容以會話方式陳述而不是生硬的學術語調；以及始終不斷強調的演算法效能基準。如果您正處於演算法領域的現實世界中，本書勢必會改變您對資料結構的使用方式。」",
		"/assets/static/image/b2.jpg",
	}
	b3 := Book{
		"JavaScript 大全, 6/e (JavaScript: The Definitive Guide: Activate Your Web Pages, 6/e)",
		"歐萊禮",
		"David Flanagan 著、黃銘偉 譯 ",
		1200,
		"這本書是程式設計師指南，也是JavaScript核心語言以及瀏覽器定義的客戶端JavaScript API之綜合參考。 ",
		"/assets/static/image/b3.jpg",
	}
	bLice = append(bLice, b1)
	bLice = append(bLice, b2)
	bLice = append(bLice, b3)
	return bLice
}
