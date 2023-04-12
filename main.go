package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/vasuvanka/machinemax-assignment/configuration"
	"github.com/vasuvanka/machinemax-assignment/definitions"
	"github.com/vasuvanka/machinemax-assignment/services"
	"github.com/vasuvanka/machinemax-assignment/utils"
)

const (
	DevEUIBatchSize       = 100
	MaxConcurrentRequests = 10
)

func main() {

	var appConfig configuration.Config
	appConfig.Load()

	ctx, cancelFunc := context.WithCancel(context.Background())
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		fmt.Println("\nInterrupt signal received. Cancelling registration...")
		cancelFunc()
	}()

	loRaWANService := services.NewLoRaWANService()

	devEUIs, err := utils.GenerateDevEUIBatch(DevEUIBatchSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while generating DevEUI Batch")
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, MaxConcurrentRequests)
	for _, devEUI := range devEUIs {
		select {
		case <-ctx.Done():
			break

		case sem <- struct{}{}:
			wg.Add(1)
			go func(devEUI definitions.DevEUI) {
				defer func() {
					<-sem
					wg.Done()
				}()

				responseDTO, err := loRaWANService.Register(ctx, devEUI)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error registering DevEUI - %s: %s\n", devEUI, err)
				}

				if responseDTO.State == definitions.RegistrationStateFailure {
					fmt.Fprintf(os.Stderr, "Error registering DevEUI - %s: %s\n", devEUI, responseDTO.Message)
				}

				if responseDTO.State == definitions.RegistrationStateDuplicate {
					fmt.Printf("DevEUI already Registered - %s\n", devEUI)
				}

				if responseDTO.State == definitions.RegistrationStateSuccess {
					fmt.Printf("Registered DevEUI %s - short code - %s \n", devEUI, responseDTO.DevEUI.ShortCode())
				}
			}(devEUI)
		}
	}

	// Wait for all requests to finish
	wg.Wait()
}
