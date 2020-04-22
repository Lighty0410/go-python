package client

//

// import (
// 	"fmt"
// 	"testing"
// )
//
// func TestNewGrpcClient(t *testing.T) {
// 	tt := []struct {
// 		name        string
// 		address     string
// 		expectedErr string
// 	}{
// 		{
// 			name:        "cannot connect",
// 			address:     "localhost:50051",
// 			expectedErr: "",
// 		},
// 	}
// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			client, err := NewGrpcClient(tc.address) // TODO refactor
// 			fmt.Println(client)
// 			fmt.Println(err)
// 		})
// 	}
// }
