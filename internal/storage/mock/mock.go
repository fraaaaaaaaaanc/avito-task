package mock_storage

import (
        hlModel "avito-tech/internal/models/hanlders_models"
        storageModels "avito-tech/internal/models/storage_model"
        context "context"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockStorageBanner is a mock of StorageBanner interface.
type MockStorageBanner struct {
        ctrl     *gomock.Controller
        recorder *MockStorageBannerMockRecorder
}

// MockStorageBannerMockRecorder is the mock recorder for MockStorageBanner.
type MockStorageBannerMockRecorder struct {
        mock *MockStorageBanner
}

// NewMockStorageBanner creates a new mock instance.
func NewMockStorageBanner(ctrl *gomock.Controller) *MockStorageBanner {
        mock := &MockStorageBanner{ctrl: ctrl}
        mock.recorder = &MockStorageBannerMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageBanner) EXPECT() *MockStorageBannerMockRecorder {
        return m.recorder
}

// CloseDB mocks base method.
func (m *MockStorageBanner) CloseDB() {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "CloseDB")
}

// CloseDB indicates an expected call of CloseDB.
func (mr *MockStorageBannerMockRecorder) CloseDB() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseDB", reflect.TypeOf((*MockStorageBanner)(nil).CloseDB))
}

// DelBanner mocks base method.
func (m *MockStorageBanner) DelBanner(arg0 context.Context, arg1 int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DelBanner", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DelBanner indicates an expected call of DelBanner.
func (mr *MockStorageBannerMockRecorder) DelBanner(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelBanner", reflect.TypeOf((*MockStorageBanner)(nil).DelBanner), arg0, arg1)
}

// DelBannerFeatureOrTag mocks base method.
func (m *MockStorageBanner) DelBannerFeatureOrTag(arg0 context.Context, arg1 []storageModels.DelFeatureOrTagChan) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DelBannerFeatureOrTag", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DelBannerFeatureOrTag indicates an expected call of DelBannerFeatureOrTag.
func (mr *MockStorageBannerMockRecorder) DelBannerFeatureOrTag(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelBannerFeatureOrTag", reflect.TypeOf((*MockStorageBanner)(nil).DelBannerFeatureOrTag), arg0, arg1)
}

// GetBanner mocks base method.
func (m *MockStorageBanner) GetBanner(arg0 context.Context, arg1 hlModel.GetBannerModel) (*hlModel.ResponseBannerModel, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetBanner", arg0, arg1)
        ret0, _ := ret[0].(*hlModel.ResponseBannerModel)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetBanner indicates an expected call of GetBanner.
func (mr *MockStorageBannerMockRecorder) GetBanner(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanner", reflect.TypeOf((*MockStorageBanner)(nil).GetBanner), arg0, arg1)
}

// GetUserBanner mocks base method.
func (m *MockStorageBanner) GetUserBanner(arg0 context.Context, arg1 hlModel.GetUserBannerModel) (*hlModel.BannerContentModel, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserBanner", arg0, arg1)
        ret0, _ := ret[0].(*hlModel.BannerContentModel)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserBanner indicates an expected call of GetUserBanner.
func (mr *MockStorageBannerMockRecorder) GetUserBanner(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBanner", reflect.TypeOf((*MockStorageBanner)(nil).GetUserBanner), arg0, arg1)
}

// GetVersionBanner mocks base method.
func (m *MockStorageBanner) GetVersionBanner(arg0 context.Context, arg1 int) (*[]hlModel.ResponseBannerModel, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetVersionBanner", arg0, arg1)
        ret0, _ := ret[0].(*[]hlModel.ResponseBannerModel)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetVersionBanner indicates an expected call of GetVersionBanner.
func (mr *MockStorageBannerMockRecorder) GetVersionBanner(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionBanner", reflect.TypeOf((*MockStorageBanner)(nil).GetVersionBanner), arg0, arg1)
}

// PatchBanner mocks base method.
func (m *MockStorageBanner) PatchBanner(arg0 context.Context, arg1 *hlModel.PatchBannerModel) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "PatchBanner", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// PatchBanner indicates an expected call of PatchBanner.
func (mr *MockStorageBannerMockRecorder) PatchBanner(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBanner", reflect.TypeOf((*MockStorageBanner)(nil).PatchBanner), arg0, arg1)
}

// SetBanner mocks base method.
func (m *MockStorageBanner) SetBanner(arg0 context.Context, arg1 *hlModel.PostBannerModel) (*hlModel.ResponsePostBannerModel, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SetBanner", arg0, arg1)
        ret0, _ := ret[0].(*hlModel.ResponsePostBannerModel)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SetBanner indicates an expected call of SetBanner.
func (mr *MockStorageBannerMockRecorder) SetBanner(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBanner", reflect.TypeOf((*MockStorageBanner)(nil).SetBanner), arg0, arg1)
}