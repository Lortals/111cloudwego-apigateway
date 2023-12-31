package main

import (
	management "github.com/Lortals/111cloudwego-apigateway/rpc-services/division-service/kitex_gen/division/management/divisionmanagement"
	"github.com/Lortals/111cloudwego-apigateway/utils"
)

func main() {

	servers := utils.CreateMultipleServers(
		3,                           // Number of servers you want to create to handle requests made to this microservice
		"Division",                  // name of microservice (servers will be registered under this name)
		new(DivisionManagementImpl), // the handler file that you defined in the previous step
		management.NewServiceInfo(), // serviceInfo file containing generated details unique to your microservice
		utils.RateLimit(1000, 1000), // optional rate limit if you wish to include it
	)

	utils.RunServers(servers) // runs all servers simultaneously
}
