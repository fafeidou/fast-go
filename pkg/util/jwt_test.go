package util

import (
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	claims, e := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA5OGY2YmNkNDYyMWQzNzNjYWRlNGU4MzI2MjdiNGY2IiwicGFzc3dvcmQiOiJjYzAzZTc0N2E2YWZiYmNiZjhiZTc2NjhhY2ZlYmVlNSIsImV4cCI6MTU2MzQ0OTQ0MywiaXNzIjoiZmFzdC1nbyJ9.6oTiM2BJKJbtbgZ7SLAjWkzsxivm12k8kAy-7ivbCUY")
	fmt.Println(claims,e)
}