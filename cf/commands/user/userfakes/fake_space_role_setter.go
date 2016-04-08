// This file was generated by counterfeiter
package userfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/command_registry"
	"github.com/cloudfoundry/cli/cf/commands/user"
	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/flags"
)

type FakeSpaceRoleSetter struct {
	MetaDataStub        func() command_registry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct{}
	metaDataReturns     struct {
		result1 command_registry.CommandMetadata
	}
	SetDependencyStub        func(deps command_registry.Dependency, pluginCall bool) command_registry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		deps       command_registry.Dependency
		pluginCall bool
	}
	setDependencyReturns struct {
		result1 command_registry.Command
	}
	RequirementsStub        func(requirementsFactory requirements.Factory, context flags.FlagContext) []requirements.Requirement
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
	}
	ExecuteStub        func(context flags.FlagContext)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		context flags.FlagContext
	}
	SetSpaceRoleStub        func(space models.Space, role, userGuid, userName string) (err error)
	setSpaceRoleMutex       sync.RWMutex
	setSpaceRoleArgsForCall []struct {
		space    models.Space
		role     string
		userGuid string
		userName string
	}
	setSpaceRoleReturns struct {
		result1 error
	}
}

func (fake *FakeSpaceRoleSetter) MetaData() command_registry.CommandMetadata {
	fake.metaDataMutex.Lock()
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	} else {
		return fake.metaDataReturns.result1
	}
}

func (fake *FakeSpaceRoleSetter) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeSpaceRoleSetter) MetaDataReturns(result1 command_registry.CommandMetadata) {
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 command_registry.CommandMetadata
	}{result1}
}

func (fake *FakeSpaceRoleSetter) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	fake.setDependencyMutex.Lock()
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		deps       command_registry.Dependency
		pluginCall bool
	}{deps, pluginCall})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(deps, pluginCall)
	} else {
		return fake.setDependencyReturns.result1
	}
}

func (fake *FakeSpaceRoleSetter) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeSpaceRoleSetter) SetDependencyArgsForCall(i int) (command_registry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return fake.setDependencyArgsForCall[i].deps, fake.setDependencyArgsForCall[i].pluginCall
}

func (fake *FakeSpaceRoleSetter) SetDependencyReturns(result1 command_registry.Command) {
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 command_registry.Command
	}{result1}
}

func (fake *FakeSpaceRoleSetter) Requirements(requirementsFactory requirements.Factory, context flags.FlagContext) []requirements.Requirement {
	fake.requirementsMutex.Lock()
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}{requirementsFactory, context})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(requirementsFactory, context)
	} else {
		return fake.requirementsReturns.result1
	}
}

func (fake *FakeSpaceRoleSetter) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeSpaceRoleSetter) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return fake.requirementsArgsForCall[i].requirementsFactory, fake.requirementsArgsForCall[i].context
}

func (fake *FakeSpaceRoleSetter) RequirementsReturns(result1 []requirements.Requirement) {
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
	}{result1}
}

func (fake *FakeSpaceRoleSetter) Execute(context flags.FlagContext) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		context flags.FlagContext
	}{context})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		fake.ExecuteStub(context)
	}
}

func (fake *FakeSpaceRoleSetter) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeSpaceRoleSetter) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].context
}

func (fake *FakeSpaceRoleSetter) SetSpaceRole(space models.Space, role string, userGuid string, userName string) (err error) {
	fake.setSpaceRoleMutex.Lock()
	fake.setSpaceRoleArgsForCall = append(fake.setSpaceRoleArgsForCall, struct {
		space    models.Space
		role     string
		userGuid string
		userName string
	}{space, role, userGuid, userName})
	fake.setSpaceRoleMutex.Unlock()
	if fake.SetSpaceRoleStub != nil {
		return fake.SetSpaceRoleStub(space, role, userGuid, userName)
	} else {
		return fake.setSpaceRoleReturns.result1
	}
}

func (fake *FakeSpaceRoleSetter) SetSpaceRoleCallCount() int {
	fake.setSpaceRoleMutex.RLock()
	defer fake.setSpaceRoleMutex.RUnlock()
	return len(fake.setSpaceRoleArgsForCall)
}

func (fake *FakeSpaceRoleSetter) SetSpaceRoleArgsForCall(i int) (models.Space, string, string, string) {
	fake.setSpaceRoleMutex.RLock()
	defer fake.setSpaceRoleMutex.RUnlock()
	return fake.setSpaceRoleArgsForCall[i].space, fake.setSpaceRoleArgsForCall[i].role, fake.setSpaceRoleArgsForCall[i].userGuid, fake.setSpaceRoleArgsForCall[i].userName
}

func (fake *FakeSpaceRoleSetter) SetSpaceRoleReturns(result1 error) {
	fake.SetSpaceRoleStub = nil
	fake.setSpaceRoleReturns = struct {
		result1 error
	}{result1}
}

var _ user.SpaceRoleSetter = new(FakeSpaceRoleSetter)
