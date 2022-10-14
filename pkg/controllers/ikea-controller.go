package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tbpathirana/go-ikea/pkg/models"
	"github.com/tbpathirana/go-ikea/pkg/utils"
	"net/http"
	"strconv"
)

var newProduct models.Product

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	newProduct := models.GetProduct()
	// json marshal (byte to string)
	res, _ := json.Marshal(newProduct)
	setHeader(w)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["ProductId"]
	if ID, err := strconv.ParseInt(productId, 10, 64); err == nil {
		bookDetails, _ := models.GetProductById(ID)
		res, _ := json.Marshal(bookDetails)
		setHeader(w)
		w.Write(res)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := &models.Product{}
	utils.ParseBody(r, newProduct)
	p := newProduct.CreateProduct()
	res, _ := json.Marshal(p)
	setHeader(w)
	w.Write(res)
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["ProductId"]
	if ID, err := strconv.ParseInt(productId, 10, 64); err == nil {
		productToDelete := models.DeleteProductById(ID)
		res, _ := json.Marshal(productToDelete)
		setHeader(w)
		w.Write(res)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := &models.Product{}
	utils.ParseBody(r, newProduct)
	vars := mux.Vars(r)
	productID := vars["ProductId"]

	ID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	productDetails, db := models.GetProductById(ID)

	productDetails.Name = newProduct.Name
	productDetails.Colour = newProduct.Colour
	productDetails.Location = newProduct.Location
	productDetails.Size = newProduct.Size
	productDetails.Price = newProduct.Size
	productDetails.Availability = newProduct.Availability

	//saving the db
	db.Save(&newProduct)

	//sending json response
	res, _ := json.Marshal(newProduct)
	setHeader(w)
	w.Write(res)
}
