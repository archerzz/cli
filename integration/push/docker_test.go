package push

import (
	"os"

	"code.cloudfoundry.org/cli/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("pushing a docker image", func() {
	var (
		appName string
	)

	BeforeEach(func() {
		appName = helpers.NewAppName()
	})

	Describe("a public docker image", func() {
		Describe("app existence", func() {
			Context("when the app does not exist", func() {
				It("creates the app", func() {
					session := helpers.CF(PushCommandName, appName, "-o", PublicDockerImage)
					Eventually(session).Should(Say("Getting app info\\.\\.\\."))
					Eventually(session).Should(Say("Creating app with these attributes\\.\\.\\."))
					Eventually(session).Should(Say("\\+\\s+name:\\s+%s", appName))
					Eventually(session).Should(Say("\\s+docker image:\\s+%s", PublicDockerImage))
					Eventually(session).Should(Say("Mapping routes\\.\\.\\."))
					Eventually(session).Should(Say("Staging Complete"))
					Eventually(session).Should(Say("Waiting for app to start\\.\\.\\."))
					Eventually(session).Should(Say("requested state:\\s+started"))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the app exists", func() {
				BeforeEach(func() {
					Eventually(helpers.CF(PushCommandName, appName, "-o", PublicDockerImage)).Should(Exit(0))
				})

				It("updates the app", func() {
					session := helpers.CF(PushCommandName, appName, "-o", PublicDockerImage)
					Eventually(session).Should(Say("Getting app info\\.\\.\\."))
					Eventually(session).Should(Say("Updating app with these attributes\\.\\.\\."))
					Eventually(session).Should(Say("\\s+name:\\s+%s", appName))
					Eventually(session).Should(Say("\\s+docker image:\\s+%s", PublicDockerImage))
					Eventually(session).Should(Say("Mapping routes\\.\\.\\."))
					Eventually(session).Should(Say("Waiting for app to start\\.\\.\\."))
					Eventually(session).Should(Say("requested state:\\s+started"))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Exit(0))
				})
			})
		})
	})

	Describe("private docker image", func() {
		var (
			privateDockerImage    string
			privateDockerUsername string
			privateDockerPassword string
		)

		BeforeEach(func() {
			privateDockerImage = os.Getenv("CF_INT_DOCKER_IMAGE")
			privateDockerUsername = os.Getenv("CF_INT_DOCKER_USERNAME")
			privateDockerPassword = os.Getenv("CF_INT_DOCKER_PASSWORD")

			if privateDockerImage == "" || privateDockerUsername == "" || privateDockerPassword == "" {
				Skip("CF_INT_DOCKER_IMAGE, CF_INT_DOCKER_USERNAME, or CF_INT_DOCKER_PASSWORD is not set")
			}
		})

		Context("when CF_DOCKER_PASSWORD is set", func() {
			It("push the docker image with those credentials", func() {
				session := helpers.CustomCF(
					helpers.CFEnv{
						EnvVars: map[string]string{"CF_DOCKER_PASSWORD": privateDockerPassword},
					},
					PushCommandName, "--docker-username", privateDockerUsername, "--docker-image", privateDockerImage, appName,
				)
				Eventually(session).Should(Say("Getting app info\\.\\.\\."))
				Eventually(session).Should(Say("Creating app with these attributes\\.\\.\\."))
				Eventually(session).Should(Say("\\+\\s+name:\\s+%s", appName))
				Eventually(session).Should(Say("\\s+docker image:\\s+%s", privateDockerImage))
				Eventually(session).Should(Say("Mapping routes\\.\\.\\."))
				Eventually(session).Should(Say("Staging Complete"))
				Eventually(session).Should(Say("Waiting for app to start\\.\\.\\."))
				Eventually(session).Should(Say("requested state:\\s+started"))
				Eventually(session).Should(Exit(0))
			})
		})

		Context("when CF_DOCKER_PASSWORD is *not* set", func() {
			It("errors with usage", func() {
				session := helpers.CF(PushCommandName, "--docker-username", privateDockerUsername, "--docker-image", privateDockerImage, appName)
				Eventually(session.Err).Should(Say("Environment variable CF_DOCKER_PASSWORD not set."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the --docker-username is provided without an image", func() {
			It("errors with usage", func() {
				session := helpers.CF(PushCommandName, "--docker-username", privateDockerUsername, appName)
				Eventually(session.Err).Should(Say("Incorrect Usage: '--docker-image, -o' and '--docker-username' must be used together."))
				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("Push a new app or sync changes to an existing app"))
				Eventually(session).Should(Exit(1))
			})
		})
	})
})
