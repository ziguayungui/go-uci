package uci

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockTree implements the Tree interface and an instance replaces the
// global defaultTree (see TestMain).
type mockTree struct {
	mock.Mock
}

func (m *mockTree) LoadConfig(name string, forceReload bool) error {
	args := m.Called(name, forceReload)
	return args.Error(0)
}

func (m *mockTree) Commit(configs ...string) error {
	args := m.Called(configs)
	return args.Error(0)
}

func (m *mockTree) Revert(configs ...string) {
	m.Called(configs)
}

func (m *mockTree) GetSections(config string, secType string) ([]string, error) {
	args := m.Called(config, secType)
	return []string{args.String(0)}, args.Error(1)
}

func (m *mockTree) GetSectionOptions(config string, section string) ([]SectionOption, error) {
	args := m.Called(config, section)
	return []SectionOption{{
		Name: args.String(0),
		Type: OptionType(args.Int(1)),
	}}, args.Error(2)
}

func (m *mockTree) Get(config, section, option string) ([]string, bool) {
	args := m.Called(config, section, option)
	var vals []string
	if arg0, ok := args.Get(0).([]string); ok {
		vals = arg0
	} else {
		vals = []string{args.String(0)}
	}
	return vals, args.Bool(1)
}

func (m *mockTree) GetLast(config, section, option string) (string, bool) {
	args := m.Called(config, section, option)
	return args.String(0), args.Bool(1)
}

func (m *mockTree) GetBool(config, section, option string) (bool, bool) {
	args := m.Called(config, section, option)
	return args.Bool(0), args.Bool(1)
}

func (m *mockTree) SetType(config, section, option string, typ OptionType, values ...string) error {
	args := m.Called(config, section, option, typ, values)
	return args.Error(0)
}

func (m *mockTree) Del(config, section, option string) error {
	m.Called(config, section, option)
	return nil
}

func (m *mockTree) AddSection(config, section, typ string) error {
	args := m.Called(config, section, typ)
	return args.Error(0)
}

func (m *mockTree) DelSection(config, section string) error {
	m.Called(config, section)
	return nil
}

func (m *mockTree) AddAnonymousSection(config, typ string) (string, error) {
	args := m.Called(config, typ)
	return args.String(0), args.Error(1)
}

func (m *mockTree) AddList(config, section, option, value string) error {
	args := m.Called(config, section, option, value)
	return args.Error(0)
}

func (m *mockTree) DelList(config, section, option, value string) (bool, error) {
	args := m.Called(config, section, option, value)
	return args.Bool(0), args.Error(1)
}

func (m *mockTree) RenameSection(config, oldName, newName string) error {
	args := m.Called(config, oldName, newName)
	return args.Error(0)
}

func (m *mockTree) Show(config, section, option string) (string, error) {
	args := m.Called(config, section, option)
	return args.String(0), args.Error(1)
}

func TestMain(m *testing.M) {
	defaultTree = &mockTree{}
	os.Exit(m.Run())
}

func TestConvenienceLoadConfig(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("LoadConfig", "foo", true).Return(nil)
	m.On("LoadConfig", "bar", false).Return(io.ErrUnexpectedEOF)
	assert.NoError(LoadConfig("foo", true))
	err := LoadConfig("bar", false)
	assert.Error(err, io.ErrUnexpectedEOF)
	m.AssertExpectations(t)
}

func TestConvenienceCommit(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)

	m.On("Commit", mock.Anything).Return(nil)
	assert.NoError(Commit())
	assert.NoError(Commit("network"))
	assert.NoError(Commit("network", "system"))

	m.AssertExpectations(t)
}

func TestConvenienceRevert(t *testing.T) {
	m := defaultTree.(*mockTree)
	m.On("Revert", []string{"foo", "bar"}).Return()
	Revert("foo", "bar")
	m.AssertExpectations(t)
}

func TestConvenienceGetSections(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("GetSections", "foo", "bar").Return("sec1", nil)
	list, err := GetSections("foo", "bar")
	assert.NoError(err)
	assert.EqualValues([]string{"sec1"}, list)
	m.AssertExpectations(t)
}

func TestConvenienceGetSectionOptions(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("GetSectionOptions", "foo", "bar").Return("opt1", 1, nil)
	list, err := GetSectionOptions("foo", "bar")
	assert.NoError(err)
	assert.EqualValues([]SectionOption{{Name: "opt1", Type: TypeList}}, list)
	m.AssertExpectations(t)
}

func TestConvenienceGet(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("Get", "foo", "bar", "opt").Return("ok", true)
	list, ok := Get("foo", "bar", "opt")
	assert.True(ok)
	assert.EqualValues([]string{"ok"}, list)
	m.AssertExpectations(t)
}

func TestConvenienceGetLast(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("GetLast", "foo", "bar", "opt").Return("ok", true)
	list, ok := GetLast("foo", "bar", "opt")
	assert.True(ok)
	assert.EqualValues("ok", list)
	m.AssertExpectations(t)
}

func TestConvenienceGetBool(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("GetBool", "foo", "bar", "opt").Return(true, true)
	value, ok := GetBool("foo", "bar", "opt")
	assert.True(ok)
	assert.True(value)
	m.AssertExpectations(t)
}

func TestConvenienceDel(t *testing.T) {
	m := defaultTree.(*mockTree)
	m.On("Del", "foo", "bar", "opt").Return()
	err := Del("foo", "bar", "opt")
	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestConvenienceAddSection(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	addSectionErr := errors.New("invalid")
	m.On("AddSection", "foo", "bar", "system").Return(nil)
	m.On("AddSection", "foo", "bar", "interface").Return(addSectionErr) //nolint:goerr113
	err := AddSection("foo", "bar", "interface")
	assert.Error(err)
	assert.EqualError(err, addSectionErr.Error())
	assert.NoError(AddSection("foo", "bar", "system"))
	m.AssertExpectations(t)
}

func TestConvenienceDelSection(t *testing.T) {
	m := defaultTree.(*mockTree)
	m.On("DelSection", "foo", "bar").Return()
	err := DelSection("foo", "bar")
	assert.NoError(t, err)
	m.AssertExpectations(t)
}

func TestConvenienceAddAnonymousSection(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("AddAnonymousSection", "network", "interface").Return("@interface[0]", nil)
	secName, err := AddAnonymousSection("network", "interface")
	assert.NoError(err)
	assert.Equal("@interface[0]", secName)
	m.AssertExpectations(t)
}

func TestConvenienceAddList(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("AddList", "network", "lan", "dns", "8.8.8.8").Return(nil)
	err := AddList("network", "lan", "dns", "8.8.8.8")
	assert.NoError(err)
	m.AssertExpectations(t)
}

func TestConvenienceDelList(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("DelList", "network", "lan", "dns", "8.8.8.8").Return(true, nil)
	removed, err := DelList("network", "lan", "dns", "8.8.8.8")
	assert.NoError(err)
	assert.True(removed)
	m.AssertExpectations(t)
}

func TestConvenienceRenameSection(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("RenameSection", "network", "old_lan", "new_lan").Return(nil)
	err := RenameSection("network", "old_lan", "new_lan")
	assert.NoError(err)
	m.AssertExpectations(t)
}

func TestConvenienceShow(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	expectedOutput := "config interface 'lan'\n\toption ifname 'eth0'\n"
	m.On("Show", "network", "lan", "").Return(expectedOutput, nil)
	output, err := Show("network", "lan", "")
	assert.NoError(err)
	assert.Equal(expectedOutput, output)
	m.AssertExpectations(t)
}

func TestConvenienceSet(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("SetType", "network", "lan", "ipaddr", TypeOption, []string{"192.168.1.1"}).Return(nil)
	err := Set("network", "lan", "ipaddr", "192.168.1.1")
	assert.NoError(err)
	m.AssertExpectations(t)
}

func TestConvenienceShowPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	expectedOutput := "config interface 'lan'\n\toption ifname 'eth0'\n"
	m.On("Show", "network", "lan", "").Return(expectedOutput, nil)
	output, err := ShowPath("network.lan")
	assert.NoError(err)
	assert.Equal(expectedOutput, output)
	m.AssertExpectations(t)
}

func TestConvenienceGetPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("Get", "network", "lan", "ipaddr").Return([]string{"192.168.1.1"}, true)
	values, ok := GetPath("network.lan.ipaddr")
	assert.True(ok)
	assert.Equal([]string{"192.168.1.1"}, values)
	m.AssertExpectations(t)
}

func TestConvenienceGetLastPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("GetLast", "network", "lan", "ipaddr").Return("192.168.1.1", true)
	value, ok := GetLastPath("network.lan.ipaddr")
	assert.True(ok)
	assert.Equal("192.168.1.1", value)
	m.AssertExpectations(t)
}

func TestConvenienceGetBoolPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("GetBool", "network", "lan", "enabled").Return(true, true)
	value, ok := GetBoolPath("network.lan.enabled")
	assert.True(ok)
	assert.True(value)
	m.AssertExpectations(t)
}

func TestConvenienceSetPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("SetType", "network", "lan", "ipaddr", TypeOption, []string{"192.168.1.1"}).Return(nil)
	err := SetPath("network.lan.ipaddr", "192.168.1.1")
	assert.NoError(err)
	m.AssertExpectations(t)
}

func TestConvenienceDelPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("Del", "network", "lan", "ipaddr").Return(nil)
	err := DelPath("network.lan.ipaddr")
	assert.NoError(err)
	m.AssertExpectations(t)
}

func TestConvenienceAddListPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("AddList", "network", "lan", "dns", "8.8.8.8").Return(nil)
	err := AddListPath("network.lan.dns", "8.8.8.8")
	assert.NoError(err)
	m.AssertExpectations(t)
}

func TestConvenienceDelListPath(t *testing.T) {
	assert := assert.New(t)
	m := defaultTree.(*mockTree)
	m.On("DelList", "network", "lan", "dns", "8.8.8.8").Return(true, nil)
	removed, err := DelListPath("network.lan.dns", "8.8.8.8")
	assert.NoError(err)
	assert.True(removed)
	m.AssertExpectations(t)
}
