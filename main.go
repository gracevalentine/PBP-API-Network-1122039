package main

import (
	"FrameworkAPI/Controller"

	"github.com/gin-gonic/gin" // import gin
)

func main() {
	//utk inisialisasi routerGIN
	routerGIN := gin.Default()

	// route.(GET/POST/PUT/DELETE)("(nama_endpoint)", (nama_function))
	routerGIN.POST("/product/insert", Controller.InsertProductGIN)       //aman
	routerGIN.GET("/product/select", Controller.SelectProductGIN)        // aman
	routerGIN.DELETE("/product/delete/:id", Controller.DeleteProductGIN) // aman
	routerGIN.PUT("/product/update/:id", Controller.UpdateProductGIN)    //aman

	routerGIN.Run(":8080")

}
