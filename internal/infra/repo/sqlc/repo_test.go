package sqlc_test

import (
	"context"
	"testing"

	mockdb "github.com/foorester/cook/internal/infra/db/pgx/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/infra/repo/sqlc"
)

func TestCreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mockdb.NewMockDB(ctrl)           // WIP: Mock this if required
	mockConn := mockdb.NewMockConn(ctrl)       // WIP: Mock this if required
	mockQueries := mockdb.NewMockQueries(ctrl) // WIP: Mock this if required

	cookRepo := sqlc.CookRepo{
		db: mockDB, // WIP: Mock database interface (db is private, see how to solve this)
	}

	book := model.Book{
		ID:          model.NewID(uuid.New()),
		Name:        "Sample Book",
		Description: "This is a sample book",
		Owner: model.User{
			ID:       uuid.New(),
			Username: "johndoe",
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
		},
	}

	// Set up the mock behavior
	mockDB.EXPECT().Connect(gomock.Any()).Return(nil)
	mockDB.EXPECT().Conn(gomock.Any()).Return(mockConn)
	mockConn.EXPECT().Queries().Return(mockQueries)
	mockQueries.EXPECT().InsertBook(gomock.Any(), gomock.Any()).Return(nil, nil)

	err := cookRepo.CreateBook(context.Background(), book)

	if err != nil {
		t.Errorf("CreateBook returned an error: %v", err)
	}

	// Add more real use-case assertions
}
