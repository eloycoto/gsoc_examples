package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/pods"
)

func main() {
	fmt.Println("Starting")
	ctx, err := bindings.NewConnection(context.Background(), "unix:/run/podman/podman.sock")

	if err != nil {
		log.Fatalf("cannot connect to podman :%v", err)
	}

	fmt.Println(ctx)

	podList, err := pods.List(ctx, nil)

	if err != nil {
		log.Fatalf("cannot get pods:%v", err)
	}
	for _, pod := range podList {
		for _, container := range pod.Containers {
			data, err := containers.Inspect(ctx, container.Id, nil)

			if err != nil {
				log.Fatalf("cannot get container details:%v", err)
			}
			fmt.Printf("Container %+v\n", data)
		}
		fmt.Printf("Pod %+v\n", pod)
	}

}
