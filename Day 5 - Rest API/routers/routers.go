package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DataPost struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func Init() *echo.Echo {
	e := echo.New()

	// membuat group dengan prefix "/api"
	api := e.Group("/api/v1")

	api.GET("/posts", PostController)
	api.POST("/posts", PostStoreController)
	api.GET("/posts/:id", PostByIdController)
	api.DELETE("/posts/:id", PostDeleteByIdController)

	return e
}

func PostController(c echo.Context) error {

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		return err
	}

	defer res.Body.Close()

	var posts []DataPost

	if err := json.NewDecoder(res.Body).Decode(&posts); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, posts)
}

func PostByIdController(c echo.Context) error {

	id := c.Param("id")

	// Membuat URL dengan query parameter
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	var post DataPost

	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, post)
}

func PostStoreController(c echo.Context) error {

	// title := c.FormValue("title")
	// body := c.FormValue("body")

	// post := DataPost{
	// 	Title: title,
	// 	Body:  body,
	// }

	p := new(DataPost)
	if err := c.Bind(p); err != nil {
		return err
	}

	// Kirim permintaan POST ke API https://jsonplaceholder.typicode.com/posts
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return c.JSON(http.StatusCreated, p)
}

func PostDeleteByIdController(c echo.Context) error {
	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/"+idStr, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return c.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
	} else {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menghapus data"})
	}
}
