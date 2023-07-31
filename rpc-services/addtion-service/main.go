package main

import (
	management "github.com/Lortals/111cloudwego-apigateway/rpc-services/addition-service/kitex_gen/addition/management/additionmanagement"
	"github.com/Lortals/111cloudwego-apigateway/utils"
)

func main() {

	servers := utils.CreateMultipleServers(
		3,
		"Addition",
		new(AdditionManagementImpl),
		management.NewServiceInfo(),
		utils.RateLimit(1000, 1000),
	)

	utils.RunServers(servers)
}
