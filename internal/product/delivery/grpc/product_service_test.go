package grpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log"
	"testing"
)

func TestProductService_GetProducts(t *testing.T) {
	for i := 0; i < 1; i++ {
		t.Run("TestProductService_GetProducts-success",func(t *testing.T) {
			t.Parallel()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &GetProductsRequest{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Limit:         "1",
				Offset:        "10",
			}
			client := NewProductServiceClient(connection)
			resp, err := client.GetProducts(context.Background(), req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}

}
