package v3

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v3/shared"
)

//go:generate counterfeiter . V3StartActor

type V3StartActor interface {
	GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	StartApplication(appGUID string) (v3action.Application, v3action.Warnings, error)
}

type V3StartCommand struct {
	RequiredArgs flag.AppName `positional-args:"yes"`
	usage        interface{}  `usage:"CF_NAME v3-start APP_NAME"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       V3StartActor
}

func (cmd *V3StartCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor()

	ccClient, _, err := shared.NewClients(config, ui, true)
	if err != nil {
		return err
	}
	cmd.Actor = v3action.NewActor(ccClient, config)

	return nil
}

func (cmd V3StartCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(cmd.Config, true, true)
	if err != nil {
		return shared.HandleError(err)
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return shared.HandleError(err)
	}

	app, warnings, err := cmd.Actor.GetApplicationByNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	if app.Started() {
		cmd.UI.DisplayWarning("App {{.AppName}} is already started.",
			map[string]interface{}{
				"AppName": cmd.RequiredArgs.AppName,
			})
		cmd.UI.DisplayOK()
		return nil
	}

	cmd.UI.DisplayTextWithFlavor("Starting app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":   cmd.RequiredArgs.AppName,
		"OrgName":   cmd.Config.TargetedOrganization().Name,
		"SpaceName": cmd.Config.TargetedSpace().Name,
		"Username":  user.Name,
	})

	_, warnings, err = cmd.Actor.StartApplication(app.GUID)

	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	cmd.UI.DisplayOK()

	return nil
}
