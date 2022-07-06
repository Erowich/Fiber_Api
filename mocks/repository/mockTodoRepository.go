package repository

import (
	models "Fiber_Api/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

func (m *MockTodoRepository) Delete(arg0 primitive.ObjectID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoRepository)(nil).Delete), arg0)
}

func (m *MockTodoRepository) Edit(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) Edit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockTodoRepository)(nil).Edit), arg0)
}

func (m *MockTodoRepository) GetAll() ([]models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoRepository)(nil).GetAll))
}

func (m *MockTodoRepository) Insert(arg0 models.Todo) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoRepositoryMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTodoRepository)(nil).Insert), arg0)
}
