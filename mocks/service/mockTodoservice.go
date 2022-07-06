package services

import (
	dto "Fiber_Api_/dto"
	models "Fiber_Api_/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

func (m *MockTodoService) TodoDelete(arg0 primitive.ObjectID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoDelete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoServiceMockRecorder) TodoDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoDelete", reflect.TypeOf((*MockTodoService)(nil).TodoDelete), arg0)
}

func (m *MockTodoService) TodoEdit(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoEdit", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockTodoServiceMockRecorder) TodoEdit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoEdit", reflect.TypeOf((*MockTodoService)(nil).TodoEdit), arg0)
}

// TodoGetAll mocks base method.
func (m *MockTodoService) TodoGetAll() ([]models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoGetAll")
	ret0, _ := ret[0].([]models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoServiceMockRecorder) TodoGetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoGetAll", reflect.TypeOf((*MockTodoService)(nil).TodoGetAll))
}

func (m *MockTodoService) TodoInsert(arg0 models.Todo) (*dto.TodoDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoInsert", arg0)
	ret0, _ := ret[0].(*dto.TodoDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTodoServiceMockRecorder) TodoInsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoInsert", reflect.TypeOf((*MockTodoService)(nil).TodoInsert), arg0)
}
