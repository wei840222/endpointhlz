package svc

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
)

type HlzSVC struct {
	endpointSVC *EndpointSVC
}

func NewHlzSVC(svc *EndpointSVC) *HlzSVC {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &HlzSVC{
		endpointSVC: svc,
	}
}

func (svc *HlzSVC) StartCronJob() {
	c := cron.New()
	c.AddFunc("@every 5s", svc.CheckAllEndpoint)
	c.Start()
}

func (svc *HlzSVC) CheckAllEndpoint() {
	e, err := svc.endpointSVC.GetAllEndpoint()
	if err != nil {
		log.Print(err.Error())
	}
	for _, v := range e {
		code := -1
		res, err := http.Get(v.URL)
		if err != nil {
			log.Print(err.Error())
		} else {
			code = res.StatusCode
		}
		if err := svc.endpointSVC.UpdateEndpointByID(*v.ID, &Endpoint{
			Status: &code,
		}); err != nil {
			log.Print(err.Error())
		}
	}
}
