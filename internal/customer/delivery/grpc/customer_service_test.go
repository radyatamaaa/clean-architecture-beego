package grpc

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func TestCustomerService_GetCustomers(t *testing.T) {
	for i := 0; i < 1; i++ {
		t.Run("TestCustomerService_GetCustomer-success", func(t *testing.T) {
			t.Parallel()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &GetCustomersParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Limit:         "1",
				Offset:        "10",
			}
			client := NewCustomerServiceClient(connection)
			resp, err := client.GetCustomers(context.Background(), req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}

}
