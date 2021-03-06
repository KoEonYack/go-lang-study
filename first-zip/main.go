package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Person struct {
	Id     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	ImgUrl string `json:"img_url" form:"img_url"`
}

func (p Person) get() (person Person, err error) {

	row := db.QueryRow("SELECT id, name, img_url FROM person WHERE id=?", p.Id)
	err = row.Scan(&person.Id, &person.Name, &person.ImgUrl)
	if err != nil {
		return
	}
	return
}

func (p Person) getAll() (persons []Person, err error) {
	rows, err := db.Query("SELECT id, name, img_url FROM person")
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.Name, &person.ImgUrl)
		persons = append(persons, person)
	}
	defer rows.Close()
	return
}

func (p Person) add() (Id int, err error) {
	stmt, err := db.Prepare("INSERT INTO person(name, img_url) VALUES (?, ?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Name, p.ImgUrl)
	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	Id = int(id)
	defer stmt.Close()
	return
}

func (p Person) update() (rows int, err error) {
	stmt, err := db.Prepare("UPDATE person SET name=?, img_url=? WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(p.Name, p.ImgUrl, p.Id)
	if err != nil {
		log.Fatalln(err)
	}

	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	rows = int(row)
	defer stmt.Close()
	return
}

func (p Person) del() (rows int, err error) {
	stmt, err := db.Prepare("DELETE FROM person WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}

	rs, err := stmt.Exec(p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows = int(row)
	return
}

func restData() {
	delete, err := db.Query("TRUNCATE person")
	if err != nil {
		log.Fatalln(err)
	}
	defer delete.Close()

	i1, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Carl', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2F52f8E%2FbtqE9vgymwO%2FUrXBJGcVfUajthMssv9An1%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i1.Close()

	i2, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Cindy', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2FbYuUjq%2FbtqE8wNYd51%2FIEP4auHXWa67rRuZp82tR0%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i2.Close()

	i3, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Ehan', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2FQf0yC%2FbtqE8xF9f4U%2FFFjHZPrgXnKkRJDHp5n5h0%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i3.Close()

	i4, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Nathan', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2FdoeXPC%2FbtqFa7k0OH4%2F2ozvAmrURBmIUI5LitrRLk%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i4.Close()

	i5, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Noah', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2Fbmbl2t%2FbtqE81NIKdc%2FMOisXuCIpE92mH9PPlNu91%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i5.Close()

	i6, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Raccoon', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2FtZYyL%2FbtqE8ySBGN9%2FNcoIak0zDpASkgubrXKApK%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i6.Close()

	i7, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Woody', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2FbvIPsK%2FbtqE82Tt3Qy%2FKjOWKquFpJ7pZdwFC54VF1%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i7.Close()

	i8, err := db.Query("INSERT INTO person(name, img_url) VALUES ('Yozi', 'https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fk.kakaocdn.net%2Fdn%2Fbmg6qu%2FbtqE8wNYd7m%2F2gRGi5JeLTohKXBr6tChqk%2Fimg.png')")
	if err != nil {
		log.Fatalln(err)
	}
	defer i8.Close()

	return
}

func main() {

	/* 난수 생성 */
	var answer = rand.Intn(8)

	var err error
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/olim_db?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	/* Test - Template */
	router.GET("/", func(c *gin.Context) {
		restData()
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Root Page",
		})
	})

	/* Test Life bits */
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/* 참여자 모두 불러옴 */
	router.GET("/persons", func(c *gin.Context) {

		p := Person{}
		persons, err := p.getAll()
		if err != nil {
			log.Fatalln(err)
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"result": persons,
		// 	"count":  len(persons),
		// })

		fmt.Println(persons)

		byteArray, err := json.MarshalIndent(persons, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(answer)
		fmt.Println(string(byteArray))
		fmt.Println()

		c.HTML(http.StatusOK, "persons.tmpl", gin.H{
			// "result":  string(byteArray),
			"result":  persons,
			"count":   len(persons),
			"message": "한 명을 지목해 주세요",
		})

	})

	/* 한명 제거 */
	router.POST("/next", func(c *gin.Context) {

		SelectedId := c.PostForm("member")
		fmt.Println(SelectedId)

		intID, err := strconv.ParseInt(SelectedId, 10, 10)
		if err != nil {
			log.Fatalln(err)
		}

		// p := Person{Id: int(SelectedId)}
		// rows, err := p.del()
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// fmt.Println("delete rows ", rows)

		ap := Person{}
		persons, err := ap.getAll()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("=========================")
		fmt.Println(answer, SelectedId)
		fmt.Println("=========================")

		res := "persons.tmpl"
		if answer == int(intID) {
			res = "success.tmpl"
			fmt.Println("Debug", answer, SelectedId)
		}
		c.HTML(http.StatusOK, res, gin.H{
			"result":  persons,
			"count":   len(persons),
			"message": fmt.Sprintf("선택한 사람은 범인이 아닙니다."),
		})

	})

	router.DELETE("/person/:id", func(c *gin.Context) {
		id := c.Param("id")

		Id, err := strconv.ParseInt(id, 10, 10)
		if err != nil {
			log.Fatalln(err)
		}
		p := Person{Id: int(Id)}
		rows, err := p.del()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("delete rows ", rows)

		// c.JSON(http.StatusOK, gin.H{
		// 	"message": fmt.Sprintf("Successfully deleted user: %s", id),
		// })

		persons, err := p.getAll()
		if err != nil {
			log.Fatalln(err)
		}

		c.HTML(http.StatusOK, "persons.tmpl", gin.H{
			"result":  persons,
			"count":   len(persons),
			"message": "한명을 성공적으로 제거하였습니다.",
		})

	})

	// router.GET("/person/:id", func(c *gin.Context) {
	// 	var result gin.H
	// 	id := c.Param("id")

	// 	Id, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	p := Person{
	// 		Id: Id,
	// 	}
	// 	person, err := p.get()
	// 	if err != nil {
	// 		result = gin.H{
	// 			"result": nil,
	// 			"count":  0,
	// 		}
	// 	} else {
	// 		result = gin.H{
	// 			"result": person,
	// 			"count":  1,
	// 		}
	// 	}

	// 	c.JSON(http.StatusOK, result)
	// })

	// // curl http://127.0.0.1:8000/person -X POST -d '{"first_name": "rsj", "last_name": "你好"}' -H "Content-Type: application/json"
	// router.POST("/person", func(c *gin.Context) {

	// 	var p Person
	// 	err := c.Bind(&p)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	Id, err := p.add()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println(Id)
	// 	name := p.FirstName + " " + p.LastName
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf(" %s successfully created", name),
	// 	})

	// })

	// //  curl http://127.0.0.1:8000/person/1 -X PUT -d "first_name=admin&last_name=reg"
	// router.PUT("/person/:id", func(c *gin.Context) {
	// 	var (
	// 		p      Person
	// 		buffer bytes.Buffer
	// 	)

	// 	id := c.Param("id")
	// 	Id, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	err = c.Bind(&p)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	p.Id = Id
	// 	rows, err := p.update()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println(rows)
	// 	buffer.WriteString(p.FirstName)
	// 	buffer.WriteString(" ")
	// 	buffer.WriteString(p.LastName)
	// 	name := buffer.String()

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("Successfully update to %s", name),
	// 	})
	// })

	router.Run(":3004")
}
