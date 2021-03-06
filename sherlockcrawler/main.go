package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/micro/go-micro"
	aproto "github.com/ob-algdatii-20ss/SherlockGopher/analyser/proto"
	proto "github.com/ob-algdatii-20ss/SherlockGopher/sherlockcrawler/proto"

	sherlock "github.com/ob-algdatii-20ss/SherlockGopher/sherlockcrawler/sherlockcrawler"
)

const (
	serviceName = "crawler-service"
)

func main() {
	SetupLogging()
	log.Info("Started analyser")

	service := micro.NewService(micro.Name(serviceName))
	service.Init()

	fmt.Printf("[+] Successfully initialized the serivce %s", serviceName)

	crawlerservice := sherlock.NewSherlockCrawlerService()

	deps := sherlock.SherlockDependencies{
		Analyser: func() aproto.AnalyserService {
			return aproto.NewAnalyserService("analyser-service", service.Client())
		},
	}
	fmt.Printf("[+] Injected dependencies in %s", serviceName)

	crawlerservice.InjectDependency(&deps)

	err := proto.RegisterCrawlerHandler(service.Server(), crawlerservice) //ändern

	if err != nil {
		log.Fatal("Crawler->main.go->RegisterCrawlerHandler failed!")
		log.Fatal(err)
	}

	go crawlerservice.ManageTasks()

	if err = service.Run(); err != nil {
		log.Fatal("Crawler->main.go->service.Run() failed!")
		log.Fatal(err)
	} else {
		log.Infof("Service %s started as intended... ", serviceName)
	}
}

func SetupLogging() {
	_ = os.Remove("info.log")
	file, _ := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0644)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		ForceQuote:                true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
	})

	log.SetOutput(file)
}
