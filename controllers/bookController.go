package controllers

import (
	"strconv"

	"github.com/atorahmad/go/crud/initializers"
	"github.com/atorahmad/go/crud/models"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	//request ke body
	var body struct {
		Judul        string
		Nama_penulis string
		Tahun        string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	//create a books
	book := models.Books{Judul: body.Judul, Nama_penulis: body.Nama_penulis, Tahun: body.Tahun}

	if err := initializers.Db.Create(&book).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	//kalo createnya berhasil statusnya 200 dan pesannya created
	c.JSON(200, gin.H{"message": "Created"})
}

func GetBooks(c *gin.Context) {
	//strconv.Atoi sama fungsinya seperti strconv.ParseInt mengkonversi string menjadi int
	//Namun, strconv.Atoi adalah cara yang paling umum dan mudah digunakan untuk mengonversi string menjadi integer di Go.
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid page",
		})
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid limit",
		})
		return
	}
	//search dan sort
	search := c.Query("search")
	sort := c.Query("sort")
	//set default page mejadi 1
	if page <= 0 {
		page = 1
	}
	//set limit data yg akan ditampilkan setiap page menjadi 10 setiap pagenya
	if limit <= 0 {
		limit = 10
	}
	//deklarasi variabel untuk menampung data models.Books yaitu totalData
	var books []models.Books
	var totalData int64
	//buat query untuk mencari data
	query := initializers.Db.Model(&models.Books{})
	if search != "" {
		query = query.Where("judul LIKE ?", "%"+search+"%")
	}
	// sort dengan query di bawah, akan menampilkan data buku dengan tahun terbit yang dicari
	if sort == "asc" {
		query = query.Order("tahun asc")
	} else if sort == "desc" {
		query = query.Order("tahun desc")
	} else if sort == "year" {
		year, err := strconv.Atoi(c.Query("year"))
		if err != nil {
			c.Status(400)
			return
		}
		query = query.Where("tahun = ?", year)
	}
	//hitung total data yang sesuai dengan query
	query.Count(&totalData)

	//ambil data dengan limit dan offset
	//(page -1) * limit itu maksudnya setiap page dikali dengan limit
	//contoh page = 1 dan limit = 10 maka 1x10 maka hasilnya tetap 10 kan nah begitu pun dgn page 2 dan seterusnya
	offset := (page - 1) * limit
	query.Limit(limit).Offset(offset).Find(&books)

	//response
	c.JSON(200, gin.H{
		"books":     books,
		"page":      page,
		"limit":     limit,
		"totalData": totalData,
	})
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	//cari buku bedasarkan id
	var book models.Books
	result := initializers.Db.Where("id = ?", id).First(&book)
	if result.Error != nil {
		// jika buku tidak ditemukan maka respon...
		c.JSON(400, gin.H{
			"message": "Id tidak ditemukan!",
		})
		return
	}
	//kalo berhasil responnya 200
	c.JSON(200, gin.H{
		"book": book,
	})
}

func UpdateBook(c *gin.Context) {
	//get by id
	id := c.Param("id")

	//reqeust ke body
	var body struct {
		Judul        string
		Nama_penulis string
		Tahun        string
	}

	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}

	//find a book id
	var book models.Books
	initializers.Db.First(&book, id)

	//update a book
	result := initializers.Db.Model(&book).Updates(models.Books{
		Judul:        body.Judul,
		Nama_penulis: body.Nama_penulis,
		Tahun:        body.Tahun,
	})

	//kalo updatenya gagal statusnya 402
	if result.Error != nil {
		c.Status(402)
		return
	}

	c.JSON(200, gin.H{
		"book": book,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	//cari buku bedasarkan id
	var book models.Books
	result := initializers.Db.Where("id = ?", id).First(&book)
	if result.Error != nil {
		// jika buku tidak ditemukan maka respon...
		c.JSON(400, gin.H{
			"message": "Buku tidak ditemukan!",
		})
		return
	}
	//hapus buku dari database
	if err := initializers.Db.Delete(&book).Error; err != nil {
		//jika terjadi error saat menghapus maka respon...
		c.JSON(402, gin.H{
			"message": "Gagal menghapus buku"})
		return
	}
	//kirim repon 200 kalo berhasil dihapus
	c.JSON(200, gin.H{"message": "Deleted"})
}
