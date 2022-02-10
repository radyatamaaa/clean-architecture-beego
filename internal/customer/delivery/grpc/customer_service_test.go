package grpc

import (
	"context"
	"log"
	"net"
	"testing"

	_customerUseCaseMock "clean-architecture-beego/internal/customer/mocks"
	"clean-architecture-beego/internal/domain"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const bufSize = 1024 * 1024

var (
	lis                  *bufconn.Listener
	custtomerUseCaseMock *_customerUseCaseMock.Usecase
)

type Token struct {
	AccessToken string `json:"access_token"`
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	custtomerUseCaseMock = new(_customerUseCaseMock.Usecase)

	customerService := NewCustomerService(custtomerUseCaseMock)
	RegisterCustomerServiceServer(s, customerService)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCustomerService_GetCustomers(t *testing.T) {
	//resultMock
	mockCustomer := []domain.CustomerObjectResponse{}
	err := faker.FakeData(&mockCustomer)
	assert.NoError(t, err)

	for i := 0; i < 1; i++ {
		t.Run("TestProductService_integration_GetCustomer-success", func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()

			// Add token to gRPC Request.
			ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer eyJasdhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ0MTYzMjksImlhdCI6MTY0NDQxNDgyOSwiaXNzIjoiYmFja2VuZCIsImp0aSI6IjE2NDQ0MTQ4MjkwNjk1NjM5MDAiLCJ1aWQiOjMsInVzZXJuYW1lIjoicmFkeWExMjMifQ.ARQx7B_rAYe9sb9hL0Eaq4ChMT58kMgJPrFW3ft8QcA")

			req := &GetCustomersParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Limit:         "1",
				Offset:        "10",
			}
			client := NewCustomerServiceClient(connection)
			resp, err := client.GetCustomers(ctx, req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})

		t.Run("TestCustomerService_mocks_GetCustomers-success", func(t *testing.T) {
			t.Parallel()

			custtomerUseCaseMock.On("GetCustomers", mock.Anything,
				mock.AnythingOfType("int"),
				mock.AnythingOfType("int")).
				Return(mockCustomer, nil)

			ctx := context.Background()
			var connection *grpc.ClientConn
			connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
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
			resp, err := client.GetCustomers(ctx, req)
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
