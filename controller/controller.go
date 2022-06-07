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
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	sshClient      *ssh.Client
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

	return nil
}

/* /query/bank Handler Functions */

// total
func (controller *Controller) GetQueryBankTotal(ctx *gin.Context) {
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.BankCommand, constant.TotalCommand})
	controller.ExecuteCommand(cmd, ctx)
}

// balances
func (controller *Controller) GetQueryBankBalances(ctx *gin.Context) {
	address := ctx.Param(constant.AddressParam)
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.BankCommand, constant.BalancesCommand, address})
	controller.ExecuteCommand(cmd, ctx)
}

// denom-metadata
func (controller *Controller) GetQueryBankDenomMetadata(ctx *gin.Context) {
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.BankCommand, constant.DenomMetadataCommand})
	controller.ExecuteCommand(cmd, ctx)
}

/* /query/distribution Handler Functions */

// commission
func (controller *Controller) GetQueryDistributionCommission(ctx *gin.Context) {
	validator := ctx.Param(constant.ValidatorParam)
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.DistributionCommand, constant.CommissionCommand, validator})
	controller.ExecuteCommand(cmd, ctx)
}

// community-pool
func (controller *Controller) GetQueryDistributionCommunityPool(ctx *gin.Context) {
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.DistributionCommand, constant.CommunityPoolCommand})
	controller.ExecuteCommand(cmd, ctx)
}

// params
func (controller *Controller) GetQueryDistributionParams(ctx *gin.Context) {
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.DistributionCommand, constant.ParamsCommand})
	controller.ExecuteCommand(cmd, ctx)
}

// rewards
func (controller *Controller) GetQueryDistributionRewards(ctx *gin.Context) {
	address := ctx.Param(constant.AddressParam)
	validator := ctx.Param(constant.ValidatorParam)
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.DistributionCommand, constant.RewardsCommand, address, validator})
	controller.ExecuteCommand(cmd, ctx)
}

// slashes
func (controller *Controller) GetQueryDistributionSlashes(ctx *gin.Context) {
	validator := ctx.Param(constant.ValidatorParam)
	startHeight := ctx.Param(constant.StartHeightParam)
	endHeight := ctx.Param(constant.EndHeightParam)
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.DistributionCommand, constant.SlashesCommand, validator, startHeight, endHeight})
	controller.ExecuteCommand(cmd, ctx)
}

// validator-outstanding-rewards
func (controller *Controller) GetQueryDistributionValidatorOutstandingRewards(ctx *gin.Context) {
	validator := ctx.Param(constant.ValidatorParam)
	cmd := util.GetFormattedFxcoredCommand([]string{constant.QueryCommand, constant.DistributionCommand, constant.ValidatorOutstandingRewardsCommand, validator})
	controller.ExecuteCommand(cmd, ctx)
}

func (controller *Controller) ExecuteCommand(cmd string, ctx *gin.Context) {
	session, err := controller.sshClient.NewSession()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to connect to create SSH session")
		return
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to execute fxcored command - %s", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, string(output))
}

func (controller *Controller) ValidateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		address := ctx.Param(constant.AddressParam)
		if address == "" {
			controller.BadRequest(ctx, "Address cannot be empty '%s")
			ctx.Abort()
			return
		}

		hasPrefix := strings.HasPrefix(address, constant.AddressPrefix)
		if !hasPrefix {
			controller.BadRequest(ctx, fmt.Sprintf("Address has no prefix '%s'", constant.AddressPrefix))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (controller *Controller) ValidateValidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validator := ctx.Param(constant.ValidatorParam)
		if validator == "" {
			controller.BadRequest(ctx, "Validator cannot be empty '%s")
			ctx.Abort()
			return
		}

		hasPrefix := strings.HasPrefix(validator, constant.ValidatorPrefix)
		if !hasPrefix {
			controller.BadRequest(ctx, fmt.Sprintf("Validator has no prefix '%s'", constant.ValidatorPrefix))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (controller *Controller) ValidateHeights() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startHeight := ctx.Param(constant.StartHeightParam)
		endHeight := ctx.Param(constant.EndHeightParam)

		i, err := strconv.Atoi(startHeight)
		if err != nil {
			controller.BadRequest(ctx, fmt.Sprintf("Start height is not an integer"))
			ctx.Abort()
			return
		}
		if i < 0 {
			controller.BadRequest(ctx, fmt.Sprintf("Start height is negative"))
			ctx.Abort()
			return
		}

		i, err = strconv.Atoi(endHeight)
		if err != nil {
			controller.BadRequest(ctx, fmt.Sprintf("End height is not an integer"))
			ctx.Abort()
			return
		}
		if i < 0 {
			controller.BadRequest(ctx, fmt.Sprintf("End height is negative"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}


func (controller *Controller) BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(400, fmt.Sprintf("400 Bad Request - %s", message))
}