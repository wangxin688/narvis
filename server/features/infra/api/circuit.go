package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/helpers/bgtask"
	biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/hooks"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Infra.Circuit
// @Summary Create new circuit
// @X-func {"name": "CreateCircuit"}
// @Description Create new circuit
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param circuit body schemas.CircuitCreate true "Circuit"
// @Success 200 {object} ts.IdResponse
// @Router /infra/circuits [post]
func createCircuit(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var circuit schemas.CircuitCreate
	if err = c.ShouldBindJSON(&circuit); err != nil {
		return
	}
	id, err := biz.NewCircuitService().CreateCircuit(&circuit)
	if err != nil {
		return
	}
	bgtask.BackgroundTask(func() {
		hooks.CircuitCreateHooks(id)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra.Circuit
// @Summary Get circuit
// @X-func {"name": "GetCircuit"}
// @Description Get circuit
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted circuitId"
// @Success 200 {object} schemas.Circuit
// @Router /infra/circuits/{id} [get]
func getCircuit(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	circuit, err := biz.NewCircuitService().GetCircuitById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, circuit)
}

// @Tags Infra.Circuit
// @Summary List circuits
// @X-func {"name": "ListCircuit"}
// @Description List circuits
// @Security BearerAuth
// Accept json
// Produce json
// @Param object query schemas.CircuitQuery true "query"
// @Success 200 {object} ts.ListResponse{results=[]schemas.Circuit}
// @Router /infra/circuits [get]
func listCircuit(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var query schemas.CircuitQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, circuits, err := biz.NewCircuitService().ListCircuit(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: circuits})
}

// @Tags Infra.Circuit
// @Summary Update circuit
// @X-func {"name": "UpdateCircuit"}
// @Description Update circuit
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted circuitId"
// @Param circuit body schemas.CircuitUpdate true "Circuit"
// @Success 200 {object} ts.IdResponse
// @Router /infra/circuits/{id} [put]
func updateCircuit(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	var circuit schemas.CircuitUpdate
	if err = c.ShouldBindJSON(&circuit); err != nil {
		return
	}
	diff, err := biz.NewCircuitService().UpdateCircuit(id, &circuit)
	if err != nil {
		return
	}
	bgtask.BackgroundTask(func() {
		hooks.CircuitUpdateHooks(id, diff[id])
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra.Circuit
// @Summary: Delete circuit
// @X-func {"name": "DeleteCircuit"}
// @Description: Delete circuit
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted circuitId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/circuits/{id} [delete]
func deleteCircuit(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	circuit, err := biz.NewCircuitService().DeleteCircuit(id)
	if err != nil {
		return
	}
	bgtask.BackgroundTask(func() {
		hooks.CircuitDeleteHooks(circuit)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}
