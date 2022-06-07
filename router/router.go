package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pundixtest/constant"
	"pundixtest/controller"
)

type Router struct {
	engine *gin.Engine
}

func (router *Router) Init() {
	router.engine = new(gin.Engine)
	router.engine = gin.Default()
}

func (router *Router) ServeRoutes(controller *controller.Controller) {
	queryRoutes := router.engine.Group(fmt.Sprintf("/%s", constant.QueryCommand))
	{
		bankRoutes := queryRoutes.Group(fmt.Sprintf("/%s", constant.BankCommand))
		{
			bankRoutes.GET(fmt.Sprintf("/%s", constant.TotalCommand), controller.GetQueryBankTotal)
			bankRoutes.GET(fmt.Sprintf("/%s/:%s", constant.BalancesCommand, constant.AddressParam),
				controller.ValidateAddress(), controller.GetQueryBankBalances)
			bankRoutes.GET(fmt.Sprintf("/%s", constant.DenomMetadataCommand), controller.GetQueryBankDenomMetadata)
		}

		distributionRoutes := queryRoutes.Group(fmt.Sprintf("/%s", constant.DistributionCommand))
		{
			distributionRoutes.GET(fmt.Sprintf("/%s/:%s", constant.CommissionCommand, constant.ValidatorParam),
				controller.ValidateValidator(), controller.GetQueryDistributionCommission)
			distributionRoutes.GET(fmt.Sprintf("/%s", constant.CommunityPoolCommand), controller.GetQueryDistributionCommunityPool)
			distributionRoutes.GET(fmt.Sprintf("/%s", constant.ParamsCommand), controller.GetQueryDistributionParams)
			distributionRoutes.GET(fmt.Sprintf("/%s/:%s/:%s", constant.RewardsCommand, constant.AddressParam, constant.ValidatorParam),
				controller.ValidateAddress(), controller.ValidateValidator(), controller.GetQueryDistributionRewards)
			distributionRoutes.GET(fmt.Sprintf("/%s/:%s/:%s/:%s", constant.SlashesCommand, constant.ValidatorParam, constant.StartHeightParam, constant.EndHeightParam),
				controller.ValidateValidator(), controller.ValidateHeights(), controller.GetQueryDistributionSlashes)
			distributionRoutes.GET(fmt.Sprintf("/%s/:%s", constant.ValidatorOutstandingRewardsCommand, constant.ValidatorParam),
				controller.ValidateValidator(), controller.GetQueryDistributionValidatorOutstandingRewards)
		}
	}
}

// Run on localhost by default
func (router *Router) Run(port int) {
	router.engine.Run(fmt.Sprintf(":%d", port))
}