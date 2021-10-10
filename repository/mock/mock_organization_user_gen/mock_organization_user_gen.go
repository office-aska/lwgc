// Code generated by MockGen. DO NOT EDIT.
// Source: organization_user_gen.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	firestore "cloud.google.com/go/firestore"
	gomock "github.com/golang/mock/gomock"
	model "github.com/office-aska/lwgc/domain/model"
	repository "github.com/office-aska/lwgc/repository"
)

// MockOrganizationUserRepository is a mock of OrganizationUserRepository interface.
type MockOrganizationUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationUserRepositoryMockRecorder
}

// MockOrganizationUserRepositoryMockRecorder is the mock recorder for MockOrganizationUserRepository.
type MockOrganizationUserRepositoryMockRecorder struct {
	mock *MockOrganizationUserRepository
}

// NewMockOrganizationUserRepository creates a new mock instance.
func NewMockOrganizationUserRepository(ctrl *gomock.Controller) *MockOrganizationUserRepository {
	mock := &MockOrganizationUserRepository{ctrl: ctrl}
	mock.recorder = &MockOrganizationUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationUserRepository) EXPECT() *MockOrganizationUserRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockOrganizationUserRepository) Delete(ctx context.Context, subject *model.OrganizationUser, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationUserRepositoryMockRecorder) Delete(ctx, subject interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationUserRepository)(nil).Delete), varargs...)
}

// DeleteByID mocks base method.
func (m *MockOrganizationUserRepository) DeleteByID(ctx context.Context, id string, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteByID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockOrganizationUserRepositoryMockRecorder) DeleteByID(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockOrganizationUserRepository)(nil).DeleteByID), varargs...)
}

// DeleteByIDWithTx mocks base method.
func (m *MockOrganizationUserRepository) DeleteByIDWithTx(ctx context.Context, tx *firestore.Transaction, id string, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteByIDWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByIDWithTx indicates an expected call of DeleteByIDWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) DeleteByIDWithTx(ctx, tx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByIDWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).DeleteByIDWithTx), varargs...)
}

// DeleteMulti mocks base method.
func (m *MockOrganizationUserRepository) DeleteMulti(ctx context.Context, subjects []*model.OrganizationUser, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subjects}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMulti", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMulti indicates an expected call of DeleteMulti.
func (mr *MockOrganizationUserRepositoryMockRecorder) DeleteMulti(ctx, subjects interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMulti", reflect.TypeOf((*MockOrganizationUserRepository)(nil).DeleteMulti), varargs...)
}

// DeleteMultiByIDs mocks base method.
func (m *MockOrganizationUserRepository) DeleteMultiByIDs(ctx context.Context, ids []string, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiByIDs", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiByIDs indicates an expected call of DeleteMultiByIDs.
func (mr *MockOrganizationUserRepositoryMockRecorder) DeleteMultiByIDs(ctx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDs", reflect.TypeOf((*MockOrganizationUserRepository)(nil).DeleteMultiByIDs), varargs...)
}

// DeleteMultiByIDsWithTx mocks base method.
func (m *MockOrganizationUserRepository) DeleteMultiByIDsWithTx(ctx context.Context, tx *firestore.Transaction, ids []string, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiByIDsWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiByIDsWithTx indicates an expected call of DeleteMultiByIDsWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) DeleteMultiByIDsWithTx(ctx, tx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiByIDsWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).DeleteMultiByIDsWithTx), varargs...)
}

// DeleteMultiWithTx mocks base method.
func (m *MockOrganizationUserRepository) DeleteMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.OrganizationUser, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, subjects}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMultiWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMultiWithTx indicates an expected call of DeleteMultiWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) DeleteMultiWithTx(ctx, tx, subjects interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, subjects}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).DeleteMultiWithTx), varargs...)
}

// DeleteWithTx mocks base method.
func (m *MockOrganizationUserRepository) DeleteWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.OrganizationUser, opts ...repository.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWithTx indicates an expected call of DeleteWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) DeleteWithTx(ctx, tx, subject interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).DeleteWithTx), varargs...)
}

// Free mocks base method.
func (m *MockOrganizationUserRepository) Free() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Free")
}

// Free indicates an expected call of Free.
func (mr *MockOrganizationUserRepositoryMockRecorder) Free() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockOrganizationUserRepository)(nil).Free))
}

// Get mocks base method.
func (m *MockOrganizationUserRepository) Get(ctx context.Context, id string, opts ...repository.GetOption) (*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrganizationUserRepositoryMockRecorder) Get(ctx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrganizationUserRepository)(nil).Get), varargs...)
}

// GetCollection mocks base method.
func (m *MockOrganizationUserRepository) GetCollection() *firestore.CollectionRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection")
	ret0, _ := ret[0].(*firestore.CollectionRef)
	return ret0
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetCollection))
}

// GetCollectionName mocks base method.
func (m *MockOrganizationUserRepository) GetCollectionName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCollectionName indicates an expected call of GetCollectionName.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetCollectionName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionName", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetCollectionName))
}

// GetDocRef mocks base method.
func (m *MockOrganizationUserRepository) GetDocRef(id string) *firestore.DocumentRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocRef", id)
	ret0, _ := ret[0].(*firestore.DocumentRef)
	return ret0
}

// GetDocRef indicates an expected call of GetDocRef.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetDocRef(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocRef", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetDocRef), id)
}

// GetMulti mocks base method.
func (m *MockOrganizationUserRepository) GetMulti(ctx context.Context, ids []string, opts ...repository.GetOption) ([]*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMulti", varargs...)
	ret0, _ := ret[0].([]*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetMulti(ctx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetMulti), varargs...)
}

// GetMultiWithTx mocks base method.
func (m *MockOrganizationUserRepository) GetMultiWithTx(tx *firestore.Transaction, ids []string, opts ...repository.GetOption) ([]*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMultiWithTx", varargs...)
	ret0, _ := ret[0].([]*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultiWithTx indicates an expected call of GetMultiWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetMultiWithTx(tx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultiWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetMultiWithTx), varargs...)
}

// GetWithDoc mocks base method.
func (m *MockOrganizationUserRepository) GetWithDoc(ctx context.Context, doc *firestore.DocumentRef, opts ...repository.GetOption) (*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDoc", varargs...)
	ret0, _ := ret[0].(*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDoc indicates an expected call of GetWithDoc.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetWithDoc(ctx, doc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDoc", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetWithDoc), varargs...)
}

// GetWithDocWithTx mocks base method.
func (m *MockOrganizationUserRepository) GetWithDocWithTx(tx *firestore.Transaction, doc *firestore.DocumentRef, opts ...repository.GetOption) (*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithDocWithTx", varargs...)
	ret0, _ := ret[0].(*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDocWithTx indicates an expected call of GetWithDocWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetWithDocWithTx(tx, doc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDocWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetWithDocWithTx), varargs...)
}

// GetWithTx mocks base method.
func (m *MockOrganizationUserRepository) GetWithTx(tx *firestore.Transaction, id string, opts ...repository.GetOption) (*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWithTx", varargs...)
	ret0, _ := ret[0].(*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithTx indicates an expected call of GetWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) GetWithTx(tx, id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).GetWithTx), varargs...)
}

// Insert mocks base method.
func (m *MockOrganizationUserRepository) Insert(ctx context.Context, subject *model.OrganizationUser) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockOrganizationUserRepositoryMockRecorder) Insert(ctx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockOrganizationUserRepository)(nil).Insert), ctx, subject)
}

// InsertMulti mocks base method.
func (m *MockOrganizationUserRepository) InsertMulti(ctx context.Context, subjects []*model.OrganizationUser) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMulti", ctx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMulti indicates an expected call of InsertMulti.
func (mr *MockOrganizationUserRepositoryMockRecorder) InsertMulti(ctx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMulti", reflect.TypeOf((*MockOrganizationUserRepository)(nil).InsertMulti), ctx, subjects)
}

// InsertMultiWithTx mocks base method.
func (m *MockOrganizationUserRepository) InsertMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.OrganizationUser) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMultiWithTx indicates an expected call of InsertMultiWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) InsertMultiWithTx(ctx, tx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMultiWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).InsertMultiWithTx), ctx, tx, subjects)
}

// InsertWithTx mocks base method.
func (m *MockOrganizationUserRepository) InsertWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.OrganizationUser) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertWithTx indicates an expected call of InsertWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) InsertWithTx(ctx, tx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).InsertWithTx), ctx, tx, subject)
}

// RunInTransaction mocks base method.
func (m *MockOrganizationUserRepository) RunInTransaction() func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInTransaction")
	ret0, _ := ret[0].(func(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error)
	return ret0
}

// RunInTransaction indicates an expected call of RunInTransaction.
func (mr *MockOrganizationUserRepositoryMockRecorder) RunInTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInTransaction", reflect.TypeOf((*MockOrganizationUserRepository)(nil).RunInTransaction))
}

// Search mocks base method.
func (m *MockOrganizationUserRepository) Search(ctx context.Context, param *repository.OrganizationUserSearchParam, q *firestore.Query) ([]*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, param, q)
	ret0, _ := ret[0].([]*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockOrganizationUserRepositoryMockRecorder) Search(ctx, param, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockOrganizationUserRepository)(nil).Search), ctx, param, q)
}

// SearchWithTx mocks base method.
func (m *MockOrganizationUserRepository) SearchWithTx(tx *firestore.Transaction, param *repository.OrganizationUserSearchParam, q *firestore.Query) ([]*model.OrganizationUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithTx", tx, param, q)
	ret0, _ := ret[0].([]*model.OrganizationUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithTx indicates an expected call of SearchWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) SearchWithTx(tx, param, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).SearchWithTx), tx, param, q)
}

// SetParentDoc mocks base method.
func (m *MockOrganizationUserRepository) SetParentDoc(doc *firestore.DocumentRef) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParentDoc", doc)
}

// SetParentDoc indicates an expected call of SetParentDoc.
func (mr *MockOrganizationUserRepositoryMockRecorder) SetParentDoc(doc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParentDoc", reflect.TypeOf((*MockOrganizationUserRepository)(nil).SetParentDoc), doc)
}

// StrictUpdate mocks base method.
func (m *MockOrganizationUserRepository) StrictUpdate(ctx context.Context, id string, param *repository.OrganizationUserUpdateParam, opts ...firestore.Precondition) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, param}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StrictUpdate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StrictUpdate indicates an expected call of StrictUpdate.
func (mr *MockOrganizationUserRepositoryMockRecorder) StrictUpdate(ctx, id, param interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdate", reflect.TypeOf((*MockOrganizationUserRepository)(nil).StrictUpdate), varargs...)
}

// StrictUpdateWithTx mocks base method.
func (m *MockOrganizationUserRepository) StrictUpdateWithTx(tx *firestore.Transaction, id string, param *repository.OrganizationUserUpdateParam, opts ...firestore.Precondition) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, id, param}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StrictUpdateWithTx", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StrictUpdateWithTx indicates an expected call of StrictUpdateWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) StrictUpdateWithTx(tx, id, param interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, id, param}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrictUpdateWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).StrictUpdateWithTx), varargs...)
}

// Update mocks base method.
func (m *MockOrganizationUserRepository) Update(ctx context.Context, subject *model.OrganizationUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationUserRepositoryMockRecorder) Update(ctx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationUserRepository)(nil).Update), ctx, subject)
}

// UpdateMulti mocks base method.
func (m *MockOrganizationUserRepository) UpdateMulti(ctx context.Context, subjects []*model.OrganizationUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMulti", ctx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMulti indicates an expected call of UpdateMulti.
func (mr *MockOrganizationUserRepositoryMockRecorder) UpdateMulti(ctx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMulti", reflect.TypeOf((*MockOrganizationUserRepository)(nil).UpdateMulti), ctx, subjects)
}

// UpdateMultiWithTx mocks base method.
func (m *MockOrganizationUserRepository) UpdateMultiWithTx(ctx context.Context, tx *firestore.Transaction, subjects []*model.OrganizationUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMultiWithTx", ctx, tx, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMultiWithTx indicates an expected call of UpdateMultiWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) UpdateMultiWithTx(ctx, tx, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMultiWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).UpdateMultiWithTx), ctx, tx, subjects)
}

// UpdateWithTx mocks base method.
func (m *MockOrganizationUserRepository) UpdateWithTx(ctx context.Context, tx *firestore.Transaction, subject *model.OrganizationUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWithTx", ctx, tx, subject)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWithTx indicates an expected call of UpdateWithTx.
func (mr *MockOrganizationUserRepositoryMockRecorder) UpdateWithTx(ctx, tx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithTx", reflect.TypeOf((*MockOrganizationUserRepository)(nil).UpdateWithTx), ctx, tx, subject)
}

// MockOrganizationUserRepositoryMiddleware is a mock of OrganizationUserRepositoryMiddleware interface.
type MockOrganizationUserRepositoryMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationUserRepositoryMiddlewareMockRecorder
}

// MockOrganizationUserRepositoryMiddlewareMockRecorder is the mock recorder for MockOrganizationUserRepositoryMiddleware.
type MockOrganizationUserRepositoryMiddlewareMockRecorder struct {
	mock *MockOrganizationUserRepositoryMiddleware
}

// NewMockOrganizationUserRepositoryMiddleware creates a new mock instance.
func NewMockOrganizationUserRepositoryMiddleware(ctrl *gomock.Controller) *MockOrganizationUserRepositoryMiddleware {
	mock := &MockOrganizationUserRepositoryMiddleware{ctrl: ctrl}
	mock.recorder = &MockOrganizationUserRepositoryMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationUserRepositoryMiddleware) EXPECT() *MockOrganizationUserRepositoryMiddlewareMockRecorder {
	return m.recorder
}

// BeforeDelete mocks base method.
func (m *MockOrganizationUserRepositoryMiddleware) BeforeDelete(ctx context.Context, subject *model.OrganizationUser, opts ...repository.DeleteOption) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subject}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeforeDelete", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeDelete indicates an expected call of BeforeDelete.
func (mr *MockOrganizationUserRepositoryMiddlewareMockRecorder) BeforeDelete(ctx, subject interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subject}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelete", reflect.TypeOf((*MockOrganizationUserRepositoryMiddleware)(nil).BeforeDelete), varargs...)
}

// BeforeDeleteByID mocks base method.
func (m *MockOrganizationUserRepositoryMiddleware) BeforeDeleteByID(ctx context.Context, ids []string, opts ...repository.DeleteOption) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeforeDeleteByID", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeDeleteByID indicates an expected call of BeforeDeleteByID.
func (mr *MockOrganizationUserRepositoryMiddlewareMockRecorder) BeforeDeleteByID(ctx, ids interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDeleteByID", reflect.TypeOf((*MockOrganizationUserRepositoryMiddleware)(nil).BeforeDeleteByID), varargs...)
}

// BeforeInsert mocks base method.
func (m *MockOrganizationUserRepositoryMiddleware) BeforeInsert(ctx context.Context, subject *model.OrganizationUser) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeInsert", ctx, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeInsert indicates an expected call of BeforeInsert.
func (mr *MockOrganizationUserRepositoryMiddlewareMockRecorder) BeforeInsert(ctx, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeInsert", reflect.TypeOf((*MockOrganizationUserRepositoryMiddleware)(nil).BeforeInsert), ctx, subject)
}

// BeforeUpdate mocks base method.
func (m *MockOrganizationUserRepositoryMiddleware) BeforeUpdate(ctx context.Context, old, subject *model.OrganizationUser) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeUpdate", ctx, old, subject)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeforeUpdate indicates an expected call of BeforeUpdate.
func (mr *MockOrganizationUserRepositoryMiddlewareMockRecorder) BeforeUpdate(ctx, old, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeUpdate", reflect.TypeOf((*MockOrganizationUserRepositoryMiddleware)(nil).BeforeUpdate), ctx, old, subject)
}
