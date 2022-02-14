package grpc

import (
	"clean-architecture-beego/internal/domain"
	_productUsecaseMock "clean-architecture-beego/internal/product/mocks"
	"clean-architecture-beego/pkg/logger"
	"context"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcMetadata "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log"
	"net"
	"testing"
)
const bufSize = 1024 * 1024

var (
	lis *bufconn.Listener
	productUsecaseMock  *_productUsecaseMock.Usecase
	l = logger.NewStdOutLogger(30,"all","Local",true,"1.1.0","kreditmu","clean-architecture-beego",logger.XmodeTest)
)

type Token struct {
	AccessToken string `json:"access_token"`
}
func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	productUsecaseMock = new(_productUsecaseMock.Usecase)

	productService := NewProductService(productUsecaseMock,l)
	RegisterProductServiceServer(s, productService)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestProductService_GetProducts(t *testing.T) {
	//resultMock
	mockProduct := []domain.ProductObjectResponse{}
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	for i := 0; i < 1; i++ {
		t.Run("TestProductService_integration_GetProducts-success",func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			var connection *grpc.ClientConn
			connection, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()

			// Add token to gRPC Request.
			ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ4NjI1OTQsImlhdCI6MTY0NDg2MTA5NCwiaXNzIjoiYmFja2VuZCIsImp0aSI6IjE2NDQ4NjEwOTQxNDMxNzg5MDAiLCJ1aWQiOjMsInVzZXJuYW1lIjoicmFkeWExMjMifQ.yBiTEwEqP0_Drx80Wxa0DEkz4xwxRct0kPEKSZXvUSI")


			req := &GetProductsParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Limit:         "1",
				Offset:        "10",
			}
			client := NewProductServiceClient(connection)
			resp, err := client.GetProducts(ctx, req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})

		t.Run("TestProductService_mocks_GetProducts-success",func(t *testing.T) {
			t.Parallel()

			productUsecaseMock.On("GetProducts", mock.Anything,
				mock.AnythingOfType("int"),
				mock.AnythingOfType("int")).
				Return(mockProduct, nil,"")

			ctx := context.Background()
			var connection *grpc.ClientConn
			connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer connection.Close()
			req := &GetProductsParams{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Limit:         "1",
				Offset:        "10",
			}
			client := NewProductServiceClient(connection)
			resp, err := client.GetProducts(ctx, req)
			t.Log("Response: ", resp)
			log.Printf("Response: %+v", resp)
			assert.NoError(t, err)
		})
	}

}
