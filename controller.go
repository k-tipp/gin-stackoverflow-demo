package main

import (
	"log"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	// repo *OrganizationRepository
}

func NewOrganizationController() *OrganizationController {
	return &OrganizationController{}
}

func (pc *OrganizationController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		org := Organization{}

		log.Println(c.Request.RequestURI)

		// Here the stack overflow happens
		if err := c.ShouldBind(&org); err == nil {
			// xidstr := pc.repo.NewOrganization(&org)
			loc, err := location.Get(c).Parse(c.Request.RequestURI + "/123" /* + xidstr*/)
			if err != nil {
				log.Println(err)
			}
			c.Header("Location", loc.String())
		} else {
			log.Println(err)
		}
	}
}
