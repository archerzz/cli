package v3_test

import (
	"errors"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/cli/command/v3"
	"code.cloudfoundry.org/cli/command/v3/v3fakes"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("v3-restart-app-instance Command", func() {
	var (
		cmd             v3.V3RestartAppInstanceCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v3fakes.FakeV3RestartAppInstanceActor
		binaryName      string
		processType     string
		executeErr      error
		app             string
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v3fakes.FakeV3RestartAppInstanceActor)

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)
		app = "some-app"
		processType = "some-special-type"

		cmd = v3.V3RestartAppInstanceCommand{
			RequiredArgs: flag.AppInstance{AppName: app, Index: 6},
			ProcessType:  processType,

			UI:          testUI,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
			Actor:       fakeActor,
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when checking target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(sharedaction.NoOrganizationTargetedError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(translatableerror.NoOrganizationTargetedError{BinaryName: binaryName}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			_, checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrg).To(BeTrue())
			Expect(checkTargetedSpace).To(BeTrue())
		})
	})

	Context("when the user is not logged in", func() {
		var expectedErr error

		BeforeEach(func() {
			expectedErr = errors.New("some current user error")
			fakeConfig.CurrentUserReturns(configv3.User{}, expectedErr)
		})

		It("return an error", func() {
			Expect(executeErr).To(Equal(expectedErr))
		})
	})

	Context("when the user is logged in", func() {
		BeforeEach(func() {
			fakeConfig.TargetedOrganizationReturns(configv3.Organization{
				Name: "some-org",
			})
			fakeConfig.TargetedSpaceReturns(configv3.Space{
				Name: "some-space",
				GUID: "some-space-guid",
			})
			fakeConfig.CurrentUserReturns(configv3.User{Name: "steve"}, nil)
		})

		Context("when restarting the specified instance returns an error", func() {
			BeforeEach(func() {
				fakeActor.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns(v3action.Warnings{"some-warning"}, errors.New("some-error"))
			})

			It("displays all warnings and returns the error", func() {
				Expect(executeErr).To(MatchError("some-error"))

				Expect(testUI.Out).To(Say("Restarting instance 6 of process some-special-type of app some-app in org some-org / space some-space as steve"))
				Expect(testUI.Err).To(Say("some-warning"))
			})
		})

		Context("when restarting the specified instance succeeds", func() {
			BeforeEach(func() {
				fakeActor.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexReturns(v3action.Warnings{"some-warning"}, nil)
			})

			It("deletes application process instance", func() {
				Expect(fakeActor.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexCallCount()).To(Equal(1))
				appName, spaceGUID, pType, index := fakeActor.DeleteInstanceByApplicationNameSpaceProcessTypeAndIndexArgsForCall(0)
				Expect(appName).To(Equal(app))
				Expect(spaceGUID).To(Equal("some-space-guid"))
				Expect(pType).To(Equal("some-special-type"))
				Expect(index).To(Equal(6))
			})

			It("displays all warnings and OK", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(testUI.Out).To(Say("Restarting instance 6 of process some-special-type of app some-app in org some-org / space some-space as steve"))
				Expect(testUI.Out).To(Say("OK"))
				Expect(testUI.Err).To(Say("some-warning"))
			})
		})
	})
})
