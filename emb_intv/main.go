package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// 1. Define route "/product". It should respond with details about a product specified by a query parameter (e.g. /product?category and sort=asc/desc).

// Implement conditional responses based on the presence and values of query parameters. (E.g. http 404)

// Please organize your code to handle different types of parameters in a modular and scalable way.

// category queryParam
// sort queryParam
// GET product endpoint

type product struct {
	ProductID   int     `json:"product_id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var products = make([]product, 0)

const (
	HOST = ""
	PORT = 9090
)

func main() {
	fmt.Println("Starting server")

	err := json.Unmarshal(productsByte, &products)
	if err != nil {
		panic("error loading products into memory")
	}

	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/products", productsHandler)

	host := fmt.Sprintf("%s:%v", HOST, PORT)
	err = http.ListenAndServe(host, nil)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Listening on localhost:%d\n", PORT)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	responseJson, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"internal error parsing products"}`))
		return
	}
	w.Header().Add("content-type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 0)
	r.Body.Read(body)
	fmt.Println(body)

	w.Header().Add("content-type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

var productsByte = []byte(`[
	{
		"product_id": 123,
		"name": "Premium Noise-Canceling Headphones",
		"category": "Electronics",
		"description": "High-quality headphones with noise-canceling features.",
		"price": 129.99,
		"stock": 50
	},
	{
		"product_id": 456,
		"name": "Fitness Tracking Smartwatch",
		"category": "Wearables",
		"description": "Smartwatch with fitness tracking and heart rate monitoring.",
		"price": 199.99,
		"stock": 30
	},
	{
		"product_id": 789,
		"name": "Compact Digital Camera",
		"category": "Cameras",
		"description": "Compact digital camera with zoom capabilities.",
		"price": 299.99,
		"stock": 20
	},
	{
		"product_id": 101,
		"name": "Gaming Laptop",
		"category": "Laptops",
		"description": "Laptop with powerful specifications for gaming and productivity.",
		"price": 1499.99,
		"stock": 10
	},
	{
		"product_id": 202,
		"name": "Wireless Bluetooth Speakers",
		"category": "Audio",
		"description": "Wireless Bluetooth speakers for home entertainment.",
		"price": 79.99,
		"stock": 40
	},
	{
		"product_id": 303,
		"name": "Advanced Robotic Vacuum Cleaner",
		"category": "Home Appliances",
		"description": "Robotic vacuum cleaner with advanced cleaning features.",
		"price": 349.99,
		"stock": 15
	},
	{
		"product_id": 404,
		"name": "Waterproof Hiking Boots",
		"category": "Outdoor Gear",
		"description": "Waterproof hiking boots for all-terrain adventures.",
		"price": 129.99,
		"stock": 25
	},
	{
		"product_id": 505,
		"name": "Smart Coffee Maker",
		"category": "Kitchen Appliances",
		"description": "Smart coffee maker with programmable brewing options.",
		"price": 89.99,
		"stock": 35
	},
	{
		"product_id": 606,
		"name": "Professional Tennis Racket",
		"category": "Sports Equipment",
		"description": "Professional-grade tennis racket for competitive players.",
		"price": 159.99,
		"stock": 20
	},
	{
		"product_id": 707,
		"name": "Fitness Tracker",
		"category": "Health and Fitness",
		"description": "Fitness tracker with heart rate monitoring and sleep analysis.",
		"price": 79.99,
		"stock": 30
	},
	{
		"product_id": 808,
		"name": "Modern Sectional Sofa",
		"category": "Furniture",
		"description": "Modern sectional sofa with adjustable seating configurations.",
		"price": 799.99,
		"stock": 10
	},
	{
		"product_id": 909,
		"name": "Ergonomic Office Chair",
		"category": "Office Supplies",
		"description": "Ergonomic office chair with lumbar support and adjustable armrests.",
		"price": 199.99,
		"stock": 15
	},
	{
		"product_id": 1010,
		"name": "Winter Jacket with Thermal Insulation",
		"category": "Clothing",
		"description": "Winter jacket with water-resistant material and thermal insulation.",
		"price": 149.99,
		"stock": 20
	},
	{
		"product_id": 1111,
		"name": "Canvas Wall Art",
		"category": "Home Decor",
		"description": "Canvas wall art depicting a serene landscape.",
		"price": 59.99,
		"stock": 25
	},
	{
		"product_id": 1212,
		"name": "HD Camera Drone",
		"category": "Toys and Games",
		"description": "Remote-controlled drone with HD camera for aerial photography.",
		"price": 129.99,
		"stock": 12
	},
	{
		"product_id": 1313,
		"name": "Interactive Cat Toy",
		"category": "Pet Supplies",
		"description": "Interactive cat toy with laser pointer and feather attachments.",
		"price": 19.99,
		"stock": 30
	},
	{
		"product_id": 1414,
		"name": "Electric Hair Straightener",
		"category": "Beauty and Personal Care",
		"description": "Electric hair straightener with ceramic plates.",
		"price": 49.99,
		"stock": 18
	},
	{
		"product_id": 1515,
		"name": "Dashcam with GPS Tracking",
		"category": "Automotive",
		"description": "Dashcam with GPS tracking and wide-angle lens.",
		"price": 89.99,
		"stock": 22
	},
	{
		"product_id": 1616,
		"name": "Solar-Powered Outdoor Lights",
		"category": "Gardening",
		"description": "Solar-powered outdoor lights for garden pathways.",
		"price": 29.99,
		"stock": 40
	},
	{
		"product_id": 1717,
		"name": "Carry-On Luggage with USB Charger",
		"category": "Travel Accessories",
		"description": "Lightweight and durable carry-on luggage with built-in USB charger.",
		"price": 129.99,
		"stock": 15
	}
]`)
