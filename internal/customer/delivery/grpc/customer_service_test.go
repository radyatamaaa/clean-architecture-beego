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

func TestCustomerService_GetCustomerById(t *testing.T) {
	for i := 0; i < 1; i++ {
		t.Run("TestCustomerService_GetCustomerById-success", func(t *testing.T) {
			t.Parallel()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &GetCustomerByIdParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Id:            1,
			}
			client := NewCustomerServiceClient(connection)
			resp, err := client.GetCustomerById(context.Background(), req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}
}

func TestCustomerService_StoreCustomer(t *testing.T) {
	for i := 0; i < 1; i++ {
		t.Run("TestCustomerService_StoreCustomer-success", func(t *testing.T) {
			t.Parallel()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &CustomerStoreRequest{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				CustomerName:  "Test Name",
				Phone:         "0811223344556",
				Email:         "test@mail.com",
				Address:       "Jl. test, Jakarta",
			}
			client := NewCustomerServiceClient(connection)
			resp, err := client.StoreCustomer(context.Background(), req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}
}

func TestCustomerService_UpdateCustomer(t *testing.T) {
	for i := 0; i < 1; i++ {
		t.Run("TestCustomerService_UpdateCustomer-success", func(t *testing.T) {
			t.Parallel()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &CustomerUpdateRequest{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Id:            1,
				CustomerName:  "Test Update",
				Phone:         "0811223344556",
				Email:         "testupdate@mail.com",
				Address:       "Jl. testupdate, Jakarta",
			}
			client := NewCustomerServiceClient(connection)
			resp, err := client.UpdateCustomer(context.Background(), req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}
}

func TestCustomerService_DeleteCustomer(t *testing.T) {
	for i := 0; i < 1; i++ {
		t.Run("TestCustomerService_DeleteCustomer-success", func(t *testing.T) {
			t.Parallel()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &GetCustomerByIdParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Id:            1,
			}
			client := NewCustomerServiceClient(connection)
			resp, err := client.DeleteCustomer(context.Background(), req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}
}
