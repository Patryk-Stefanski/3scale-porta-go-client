package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestListProductMethods(t *testing.T) {
	var (
		productID int64 = 97
		hitsID    int64 = 98
		endpoint        = fmt.Sprintf(productMethodListResourceEndpoint, productID, hitsID)
	)

	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		if req.URL.Path != endpoint {
			t.Fatalf("Path does not match. Expected [%s]; got [%s]", endpoint, req.URL.Path)
		}

		if req.Method != http.MethodGet {
			t.Fatalf("Method does not match. Expected [%s]; got [%s]", http.MethodGet, req.Method)
		}

		list := MethodList{
			Methods: []Method{
				{
					Element: MethodItem{
						ID:         1,
						Name:       "method01",
						SystemName: "desc01",
					},
				},
				{
					Element: MethodItem{
						ID:         2,
						Name:       "method02",
						SystemName: "desc02",
					},
				},
			},
		}

		responseBodyBytes, err := json.Marshal(list)
		if err != nil {
			t.Fatal(err)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(responseBodyBytes)),
			Header:     make(http.Header),
		}
	})

	credential := "someAccessToken"
	c := NewThreeScale(NewTestAdminPortal(t), credential, httpClient)
	list, err := c.ListProductMethods(productID, hitsID)
	if err != nil {
		t.Fatal(err)
	}

	if list == nil {
		t.Fatal("list returned nil")
	}

	if len(list.Methods) != 2 {
		t.Fatalf("Then number of list items does not match. Expected [%d]; got [%d]", 2, len(list.Methods))
	}
}

func TestCreateProductMethod(t *testing.T) {
	var (
		productID int64 = 97
		hitsID    int64 = 98
		params          = Params{"friendly_name": "method5"}
		endpoint        = fmt.Sprintf(productMethodListResourceEndpoint, productID, hitsID)
	)

	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		if req.URL.Path != endpoint {
			t.Fatalf("Path does not match. Expected [%s]; got [%s]", endpoint, req.URL.Path)
		}

		if req.Method != http.MethodPost {
			t.Fatalf("Method does not match. Expected [%s]; got [%s]", http.MethodPost, req.Method)
		}

		item := &Method{
			Element: MethodItem{
				ID:         10,
				Name:       params["friendly_name"],
				SystemName: params["friendly_name"],
			},
		}

		responseBodyBytes, err := json.Marshal(item)
		if err != nil {
			t.Fatal(err)
		}

		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(bytes.NewBuffer(responseBodyBytes)),
			Header:     make(http.Header),
		}
	})

	credential := "someAccessToken"
	c := NewThreeScale(NewTestAdminPortal(t), credential, httpClient)
	obj, err := c.CreateProductMethod(productID, hitsID, params)
	if err != nil {
		t.Fatal(err)
	}

	if obj == nil {
		t.Fatal("returned nil")
	}

	if obj.Element.ID != 10 {
		t.Fatalf("method ID does not match. Expected [%d]; got [%d]", 10, obj.Element.ID)
	}

	if obj.Element.Name != params["friendly_name"] {
		t.Fatalf("method name does not match. Expected [%s]; got [%s]", params["friendly_name"], obj.Element.Name)
	}
}

func TestDeleteProductMethod(t *testing.T) {
	var (
		productID int64 = 97
		hitsID    int64 = 98
		methodID  int64 = 123325
		endpoint        = fmt.Sprintf(productMethodResourceEndpoint, productID, hitsID, methodID)
	)

	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		if req.URL.Path != endpoint {
			t.Fatalf("Path does not match. Expected [%s]; got [%s]", endpoint, req.URL.Path)
		}

		if req.Method != http.MethodDelete {
			t.Fatalf("Method does not match. Expected [%s]; got [%s]", http.MethodDelete, req.Method)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("")),
			Header:     make(http.Header),
		}
	})

	credential := "someAccessToken"
	c := NewThreeScale(NewTestAdminPortal(t), credential, httpClient)
	err := c.DeleteProductMethod(productID, hitsID, methodID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadProductMethod(t *testing.T) {
	var (
		productID int64 = 97
		hitsID    int64 = 98
		methodID  int64 = 123325
		endpoint        = fmt.Sprintf(productMethodResourceEndpoint, productID, hitsID, methodID)
	)

	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		if req.URL.Path != endpoint {
			t.Fatalf("Path does not match. Expected [%s]; got [%s]", endpoint, req.URL.Path)
		}

		if req.Method != http.MethodGet {
			t.Fatalf("Method does not match. Expected [%s]; got [%s]", http.MethodGet, req.Method)
		}

		item := &Method{
			Element: MethodItem{
				ID:         methodID,
				Name:       "someName",
				SystemName: "someName2",
			},
		}

		responseBodyBytes, err := json.Marshal(item)
		if err != nil {
			t.Fatal(err)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(responseBodyBytes)),
			Header:     make(http.Header),
		}
	})

	credential := "someAccessToken"
	c := NewThreeScale(NewTestAdminPortal(t), credential, httpClient)
	obj, err := c.ProductMethod(productID, hitsID, methodID)
	if err != nil {
		t.Fatal(err)
	}

	if obj == nil {
		t.Fatal("method returned nil")
	}

	if obj.Element.ID != methodID {
		t.Fatalf("ID does not match. Expected [%d]; got [%d]", methodID, obj.Element.ID)
	}
}

func TestUpdateProductMethod(t *testing.T) {
	var (
		productID int64 = 98765
		hitsID    int64 = 1
		methodID  int64 = 123325
		endpoint        = fmt.Sprintf(productMethodResourceEndpoint, productID, hitsID, methodID)
		params          = Params{"description": "newDescr"}
	)

	httpClient := NewTestClient(func(req *http.Request) *http.Response {
		if req.URL.Path != endpoint {
			t.Fatalf("Path does not match. Expected [%s]; got [%s]", endpoint, req.URL.Path)
		}

		if req.Method != http.MethodPut {
			t.Fatalf("Method does not match. Expected [%s]; got [%s]", http.MethodPut, req.Method)
		}

		method := &Method{
			Element: MethodItem{
				ID:          methodID,
				Name:        "someName",
				SystemName:  "someName2",
				Description: params["description"],
			},
		}

		responseBodyBytes, err := json.Marshal(method)
		if err != nil {
			t.Fatal(err)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(responseBodyBytes)),
			Header:     make(http.Header),
		}
	})

	credential := "someAccessToken"
	c := NewThreeScale(NewTestAdminPortal(t), credential, httpClient)
	obj, err := c.UpdateProductMethod(productID, hitsID, methodID, params)
	if err != nil {
		t.Fatal(err)
	}

	if obj == nil {
		t.Fatal("method returned nil")
	}

	if obj.Element.ID != methodID {
		t.Fatalf("method ID does not match. Expected [%d]; got [%d]", methodID, obj.Element.ID)
	}

	if obj.Element.Description != params["description"] {
		t.Fatalf("method description does not match. Expected [%s]; got [%s]", params["description"], obj.Element.Description)
	}
}
