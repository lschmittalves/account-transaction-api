package tests

import (
	"account-transaction-api/internal/models"
	accountService "account-transaction-api/internal/services/account"
	mock_repositories "account-transaction-api/tests/mocks/repositories"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestAccountCreation_Success(t *testing.T) {

	accountDoc := "123456"
	accountId, _ := uuid.Parse("b18b834e-04d2-4361-8cd7-efa65021b9d8")
	accToSend := &models.Account{
		TaxDocument: accountDoc,
		Name:        "John Doe",
	}

	ctrl := gomock.NewController(t)

	accReader := mock_repositories.NewMockAccountReader(ctrl)
	accWriter := mock_repositories.NewMockAccountWriter(ctrl)

	accReader.EXPECT().ExistsDocument(gomock.Eq(accountDoc)).Return(false).Times(1)
	accWriter.EXPECT().Add(gomock.Eq(accToSend)).DoAndReturn(func(a *models.Account) error {
		a.Id = accountId
		return nil
	}).Times(1)

	service := accountService.NewAccountService(accReader, accWriter)

	acc, err := service.Create(accToSend)

	assert.Nil(t, err)
	assert.Equal(t, accountId, acc.Id)
	assert.Equal(t, accountDoc, acc.TaxDocument)
	assert.Equal(t, "John Doe", acc.Name)
}

func TestAccountCreation_DuplicatedDocument(t *testing.T) {

	accountDoc := "123456"
	accToSend := &models.Account{
		TaxDocument: accountDoc,
		Name:        "John Doe",
	}

	ctrl := gomock.NewController(t)

	accReader := mock_repositories.NewMockAccountReader(ctrl)
	accWriter := mock_repositories.NewMockAccountWriter(ctrl)

	accReader.EXPECT().ExistsDocument(gomock.Eq(accountDoc)).Return(true).Times(1)
	accWriter.EXPECT().Add(gomock.Any()).Times(0)

	service := accountService.NewAccountService(accReader, accWriter)

	acc, err := service.Create(accToSend)

	assert.Errorf(t, err, fmt.Sprintf("account with document %s allready exists", accountDoc))
	assert.Nil(t, acc)
}
