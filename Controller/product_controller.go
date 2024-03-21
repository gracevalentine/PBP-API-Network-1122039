package Controller

import (
	m "FrameworkAPI/Model"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

// untuk mux, hanya membuat codenya aja

func SelectProductMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponse(w, 500, "Failed to establish a connection to the database")
		return
	}

	var products []m.Products
	queryResult := db.Last(&products)
	if queryResult.Error != nil {
		SendErrorResponse(w, 500, "Error retrieving product data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response m.ProductsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = products
	json.NewEncoder(w).Encode(response)
}

func SelectProductGIN(c *gin.Context) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponseGIN(c, 500, "Error connecting to the database")
		return
	}

	var products []m.Products
	queryResult := db.Find(&products)
	if queryResult.Error != nil {
		SendErrorResponseGIN(c, 500, "Error retrieving product data")
		return
	}

	SendDataResponseGIN(c, 200, "Success", products)
}

func InsertProductMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponse(w, 500, "Error connecting to database")
		return
	}

	product_name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")
	price := r.URL.Query().Get("price")
	qty := r.URL.Query().Get("qty")

	priceInt, err := strconv.Atoi(price)
	if err != nil {
		SendErrorResponse(w, 505, "Failed to convert")
		return
	}

	qtystr, err := strconv.Atoi(qty)
	if err != nil {
		SendErrorResponse(w, 505, "Failed to convert")
		return
	}

	products := m.Products{
		Product_name: product_name,
		Category:     category,
		Price:        priceInt,
		Quantity:     qtystr,
	}

	result := db.Create(&products)
	err = result.Error
	if err != nil {
		SendErrorResponse(w, 404, "Data not found")
	}
	SendSuccessResponse(w, 200, "User inserted successfully!")
}

func InsertProductGIN(c *gin.Context) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponseGIN(c, 500, "Error connecting to the database")
		return
	}

	// untuk get data by query di gin
	product_name := c.Query("name")
	category := c.Query("category")
	price := c.Query("price")
	qty := c.Query("qty")

	// string to int
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		SendErrorResponseGIN(c, 505, "Failed to convert price!")
		return
	}

	// string to int
	qtyInt, err := strconv.Atoi(qty)
	if err != nil {
		SendErrorResponseGIN(c, 505, "Failed to convert quantity!")
		return
	}

	products := m.Products{
		Product_name: product_name,
		Category:     category,
		Price:        priceInt,
		Quantity:     qtyInt,
	}

	result := db.Create(&products)
	err = result.Error
	if err != nil {
		SendErrorResponseGIN(c, 404, "Data not found")
	}
	SendSuccessResponseGIN(c, 200, "Product inserted successfully!")
}

func UpdateProductMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponse(w, 500, "Error connecting to database")
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	price := r.URL.Query().Get("price")
	if id == "" {
		SendErrorResponse(w, 505, "Bad request: Missing ID")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		SendErrorResponse(w, 505, "Bad request: Invalid ID")
		return
	}
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		SendErrorResponse(w, 505, "Bad request: Invalid price")
		return
	}

	var product m.Products
	db.Find(&product, &idInt)
	product.ID = idInt
	product.Price = priceInt

	if err := db.Save(&product).Error; err != nil {
		SendErrorResponse(w, 500, "Failed to update data")
		return
	}
	SendSuccessResponse(w, 200, "Products updated successfully!")
}

func UpdateProductGIN(c *gin.Context) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponseGIN(c, 500, "Error connecting to database")
		return
	}

	id := c.Param("id")
	price := c.PostForm("price") // gunain PostForm untuk get data dari body request
	if id == "" {
		SendErrorResponseGIN(c, 400, "Bad request: Missing ID")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		SendErrorResponseGIN(c, 400, "Bad request: Invalid ID")
		return
	}
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		SendErrorResponseGIN(c, 400, "Bad request: Invalid price")
		return
	}

	var product m.Products
	if err := db.Model(&product).Where("product_id = ?", idInt).Update("price", priceInt).Error; err != nil {
		SendErrorResponseGIN(c, 404, "Product not found")
		return
	}

	product.Price = priceInt

	SendSuccessResponseGIN(c, 200, "Product updated successfully!")
}

func DeleteProductMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponse(w, 500, "Error connecting to database")
		return
	}

	id := mux.Vars(r)["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		SendErrorResponse(w, 505, "Bad request: Invalid ID")
		return
	}
	var user m.Products
	result := db.Where("id = ?", &idInt).Delete(&user)
	err = result.Error
	if err != nil {
		SendErrorResponse(w, 505, "Failed to delete data")
		return
	}

	SendSuccessResponse(w, 200, "User deleted successfully!")
}

func DeleteProductGIN(c *gin.Context) {
	db, err := connectForGorm()
	if err != nil {
		SendErrorResponseGIN(c, 500, "Error connecting to the database")
		return
	}

	//untuk get id product
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		SendErrorResponseGIN(c, 505, "Bad request: Invalid ID")
		return
	}

	var product m.Products
	result := db.Where("product_id = ?", idInt).Delete(&product)
	err = result.Error
	if err != nil {
		SendErrorResponseGIN(c, 505, "Failed to delete data")
		return
	}

	SendSuccessResponseGIN(c, 200, "Product deleted successfully!")
}
