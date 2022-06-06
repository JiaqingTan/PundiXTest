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
	sshClient *ssh.Client
}

func (controller *Controller) Init(appConfig *config.AppConfig) error {
	sshHost := "13.212.254.102"
	sshPort := 22
	sshUser := "root"
	sshPassword := "5uJr_H{F"

	config := &ssh.ClientConfig {
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	// dial get SSH client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Printf("Failed to dial remote EC2 instance via TCP")
		return err
	}

	controller.sshClient = sshClient

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

	cmd := fmt.Sprintf("go/bin/fxcored %s %s %s --node https://fx-json.functionx.io:26657", command1, command2, command3)
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
			controller.NotFound(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (controller *Controller) NotFound(ctx *gin.Context) {
	ctx.JSON(404, "404 Not Found")
}