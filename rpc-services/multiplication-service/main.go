package main

import (
	management "github.com/Lortals/111cloudwego-apigateway/rpc-services/multiplication-service/kitex_gen/multiplication/management/multiplicationmanagement"
	"github.com/Lortals/111cloudwego-apigateway/utils"
)

func main() {

	servers := utils.CreateMultipleServers(
		3,
		"Multiplication",
		new(MultiplicationManagementImpl),
		management.NewServiceInfo(),
		utils.RateLimit(1000, 1000),
	)

	utils.RunServers(servers)

}
