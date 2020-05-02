package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// AllStores - All stores data
type AllStores struct {
	GroceryStore GroceryStore `json:"grocerystore"`
	FruitStore   FruitStore   `json:"fruitstore"`
}

// GroceryStore - Grocery stores data
type GroceryStore struct {
	StoreID int      `json:"storeid"`
	Type    string   `json:"type"`
	Items   []string `json:"items"`
}

// FruitStore - Fruit stores data
type FruitStore struct {
	StoreID int      `json:"storeid"`
	Type    string   `json:"type"`
	Items   []string `json:"items"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	Gs := GroceryStore{
		StoreID: 1,
		Type:    "Grocery",
		Items:   []string{"wheat", "rice", "pulse"},
	}

	Fs := FruitStore{
		StoreID: 2,
		Type:    "Fruits",
		Items:   []string{"banana", "apple", "mango"},
	}

	showAll := AllStores{
		Gs,
		Fs,
	}

	jsonData, err := json.Marshal(showAll)
	groceryJSON, err := json.Marshal(Gs)
	fruitsJSON, err := json.Marshal(Fs)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid input",
			StatusCode: 400,
		}, err
	}

	// /stores/store?type=grocery

	ID := request.PathParameters["id"]
	ST := request.QueryStringParameters["type"]
	fmt.Println(ID, ST)

	switch {

	case ID == "store" && ST == "grocery":
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:       string(groceryJSON),
			StatusCode: 200,
		}, nil

	case ID == "store" && ST == "fruit":
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:       string(fruitsJSON),
			StatusCode: 200,
		}, nil
	case ID == "store" && ST == "all":
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:       string(jsonData),
			StatusCode: 200,
		}, nil
	default:
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:       fmt.Sprintf("%s, store not present", ST),
			StatusCode: 404,
		}, nil
	}
}
