package payment_s1

import (
	"challenge/api"
	"challenge/configs"
	"context"
	"log"
)

/*grpc server*/
func main() {
	var ctx = context.Background()
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Print(err)
	}
	v1API, err := api.NewPaymentServer(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	go func() {
		log.Fatal(RunServerGRPC(ctx, v1API, "8080"))
		return
	}()

	log.Fatal(RunServerREST(ctx, "8080", "8082"))
}
