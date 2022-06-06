package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"pundixtest/config"
	"pundixtest/constant"
	"pundixtest/util"
	"time"
)

type Controller struct {
	sshClient      *ssh.Client
	fxcoredCommand string
}

func (controller *Controller) Init(appConfig *config.AppConfig) error {
	config := &ssh.ClientConfig {
		Timeout:         time.Second,
		User:            appConfig.SSHUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(appConfig.SSHPassword)}

	addr := fmt.Sprintf("%s:%d", appConfig.SSHHost, appConfig.SSHPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Printf("Failed to dial remote EC2 instance via TCP")
		return err
	}

	controller.sshClient = sshClient
	controller.fxcoredCommand = appConfig.FXCoredCommand

	return nil
}

func (controller *Controller) ExecuteCommand(ctx *gin.Context) {
	command1 := ctx.Param(constant.Command1)
	command2 := ctx.Param(constant.Command2)
	command3 := ctx.Param(constant.Command3)

	session, err := controller.sshClient.NewSession()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to connect to create SSH session")
		return
	}
	defer session.Close()

	cmd := fmt.Sprintf(controller.fxcoredCommand, command1, command2, command3)
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to execute fxcored command")
		return
	}
	ctx.JSON(http.StatusOK, string(output))
}

func (controller *Controller) ValidateExecuteCommandParams(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		command1Param := ctx.Param(constant.Command1)
		command2Param := ctx.Param(constant.Command2)

		if !util.SliceContains(appConfig.AllowedCommands[constant.Command1], command1Param) ||
			!util.SliceContains(appConfig.AllowedCommands[constant.Command2], command2Param) {
			controller.BadRequest(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (controller *Controller) BadRequest(ctx *gin.Context) {
	ctx.JSON(400, "400 Bad Request")
}