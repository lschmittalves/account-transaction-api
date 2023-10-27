package tests

import (
	"account-transaction-api/internal/models"
	"account-transaction-api/internal/repositories"
	transacionService "account-transaction-api/internal/services/transactions"
	mock_cache "account-transaction-api/tests/mocks/clients"
	mock_repositories "account-transaction-api/tests/mocks/repositories"
	mock_account "account-transaction-api/tests/mocks/services/account"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestTransactionRegister_Success(t *testing.T) {
	accountId, _ := uuid.Parse("b18b834e-04d2-4361-8cd7-efa65021b9d8")
	operationId, _ := uuid.Parse("f37ad389-2e88-4927-85fe-047ea6677c31")
	transactionId, _ := uuid.Parse("4febfb40-8b34-41c6-90ad-7f797503d7fe")

	newTransaction := &models.Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          10,
	}

	ctrl := gomock.NewController(t)

	accReader := mock_repositories.NewMockAccountReader(ctrl)
	opReader := mock_repositories.NewMockOperationTypeReader(ctrl)
	tWriter := mock_repositories.NewMockTransactionWriter(ctrl)
	accService := mock_account.NewMockService(ctrl)

	accReader.EXPECT().ExistsId(gomock.Eq(accountId)).Return(true).Times(1)

	opReader.EXPECT().GetById(gomock.Any(), gomock.Eq(operationId)).DoAndReturn(
		func(op *models.OperationType, id uuid.UUID) error {
			op.IsDebit = false
			op.Id = operationId
			return nil
		}).Times(1)

	tWriter.EXPECT().Add(gomock.Eq(newTransaction)).DoAndReturn(func(a *models.Transaction) error {
		a.Id = transactionId
		return nil
	}).Times(1)

	accService.EXPECT().UpdateCreditLimit(gomock.Eq(accountId), gomock.Eq(int64(10))).Return(nil).Times(1)

	service := transacionService.NewTransactionService(accReader, accService, opReader, tWriter)

	tr, err := service.Create(newTransaction)

	assert.Nil(t, err)
	assert.Equal(t, transactionId, tr.Id)
	assert.Equal(t, accountId, tr.AccountId)
	assert.Equal(t, operationId, tr.OperationTypeId)
	assert.Equal(t, int64(10), tr.Amount)
}

func TestTransactionRegister_Debit(t *testing.T) {
	accountId, _ := uuid.Parse("b18b834e-04d2-4361-8cd7-efa65021b9d8")
	operationId, _ := uuid.Parse("f37ad389-2e88-4927-85fe-047ea6677c31")

	newTransaction := &models.Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          10,
	}

	ctrl := gomock.NewController(t)

	accReader := mock_repositories.NewMockAccountReader(ctrl)
	opReader := mock_repositories.NewMockOperationTypeReader(ctrl)
	tWriter := mock_repositories.NewMockTransactionWriter(ctrl)
	accService := mock_account.NewMockService(ctrl)

	accReader.EXPECT().ExistsId(gomock.Any()).Return(true).Times(1)

	opReader.EXPECT().GetById(gomock.Any(), gomock.Eq(operationId)).DoAndReturn(
		func(op *models.OperationType, id uuid.UUID) error {
			op.IsDebit = true
			op.Id = operationId
			return nil
		}).Times(1)

	tWriter.EXPECT().Add(gomock.Eq(newTransaction)).Return(nil).Times(1)

	accService.EXPECT().UpdateCreditLimit(gomock.Eq(accountId), gomock.Eq(int64(-10))).Return(nil).Times(1)

	service := transacionService.NewTransactionService(accReader, accService, opReader, tWriter)

	tr, err := service.Create(newTransaction)

	assert.Nil(t, err)
	assert.Equal(t, int64(-10), tr.Amount)
}

func TestTransactionRegister_AccountNotFound(t *testing.T) {
	accountId, _ := uuid.Parse("b18b834e-04d2-4361-8cd7-efa65021b9d8")
	operationId, _ := uuid.Parse("f37ad389-2e88-4927-85fe-047ea6677c31")

	newTransaction := &models.Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          10,
	}

	ctrl := gomock.NewController(t)

	accReader := mock_repositories.NewMockAccountReader(ctrl)
	opReader := mock_repositories.NewMockOperationTypeReader(ctrl)
	tWriter := mock_repositories.NewMockTransactionWriter(ctrl)
	accService := mock_account.NewMockService(ctrl)

	accReader.EXPECT().ExistsId(gomock.Eq(accountId)).Return(false).Times(1)
	opReader.EXPECT().GetById(gomock.Any(), gomock.Any()).Times(0)
	tWriter.EXPECT().Add(gomock.Eq(newTransaction)).Times(0)
	accService.EXPECT().UpdateCreditLimit(gomock.Any(), gomock.Any()).Times(0)

	service := transacionService.NewTransactionService(accReader, accService, opReader, tWriter)

	tr, err := service.Create(newTransaction)

	assert.Nil(t, tr)
	assert.Errorf(t, err, fmt.Sprintf("account with id %s does not exists", accountId))

}

func TestTransactionRegister_OperationNotFound(t *testing.T) {
	accountId, _ := uuid.Parse("b18b834e-04d2-4361-8cd7-efa65021b9d8")
	operationId, _ := uuid.Parse("f37ad389-2e88-4927-85fe-047ea6677c31")

	newTransaction := &models.Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          10,
	}

	ctrl := gomock.NewController(t)

	accReader := mock_repositories.NewMockAccountReader(ctrl)
	opReader := mock_repositories.NewMockOperationTypeReader(ctrl)
	tWriter := mock_repositories.NewMockTransactionWriter(ctrl)
	accService := mock_account.NewMockService(ctrl)

	accReader.EXPECT().ExistsId(gomock.Any()).Return(true).Times(1)
	opReader.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(errors.New("")).Times(1)
	tWriter.EXPECT().Add(gomock.Eq(newTransaction)).Times(0)
	accService.EXPECT().UpdateCreditLimit(gomock.Any(), gomock.Any()).Times(0)

	service := transacionService.NewTransactionService(accReader, accService, opReader, tWriter)

	tr, err := service.Create(newTransaction)

	assert.Nil(t, tr)
	assert.Errorf(t, err, fmt.Sprintf("operation type with id %s does not exists", accountId))
}

func TestTransactionRegister_TransactionIsLocked(t *testing.T) {
	accountId, _ := uuid.Parse("b18b834e-04d2-4361-8cd7-efa65021b9d8")
	operationId, _ := uuid.Parse("f37ad389-2e88-4927-85fe-047ea6677c31")

	newTransaction := &models.Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          10,
	}
	cacheId := fmt.Sprintf("%s/%s/%d", accountId, operationId, 10)

	ctrl := gomock.NewController(t)

	cache := mock_cache.NewMockCache(ctrl)
	cache.EXPECT().SetNx(gomock.Any(), gomock.Any(), cacheId).Return(false).Times(1)

	r := repositories.NewTransactionRepository(nil, cache)

	err := r.Add(newTransaction)
	assert.Errorf(t, err, fmt.Sprintf("transaction %s has been sent allready", cacheId))
}
