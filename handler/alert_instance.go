// dbmond
// Copyright (C) 2019 gywndi@gmail.com in kakaoBank
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package handler

import (
	"fmt"

	"github.com/dbmonstar/dbmond/common"
	"github.com/dbmonstar/dbmond/model"

	"github.com/gin-gonic/gin"
)

// startAlertInstanceAPI alert instance API
func startAlertInstanceAPI(r *gin.RouterGroup) {

	// update
	r.PUT("/alert/instance/:instance_id", func(c *gin.Context) {
		var err error
		var params model.AlertInstance

		// bind params (form params)
		err = c.Bind(&params)
		if ErrorIf(c, err) {
			return
		}

		// get id
		params.ID = common.ParseInt(c.Param("instance_id"))

		// check ID
		if params.ID == 0 {
			ErrorIf(c, fmt.Errorf(common.MSG["err.invalid_zero_id"]))
			return
		}

		// target
		target := model.AlertInstance{ID: params.ID}

		// check exists
		if !target.Exist() {
			ErrorIf(c, fmt.Errorf(common.MSG["err.instance_not_exists"]))
			return
		}

		// update
		_, err = target.Update(&params)
		if ErrorIf(c, err) {
			return
		}

		Success(c, params.ID)

	})

	// delete
	r.DELETE("/alert/instance/:instance_id", func(c *gin.Context) {
		var err error
		var params model.AlertInstance

		// get id
		params.ID = common.ParseInt(c.Param("instance_id"))

		// check ID
		if params.ID == 0 {
			ErrorIf(c, fmt.Errorf(common.MSG["err.invalid_zero_id"]))
			return
		}

		// target
		target := model.AlertInstance{ID: params.ID}

		// check exists
		if !target.Exist() {
			ErrorIf(c, fmt.Errorf(common.MSG["err.instance_not_exists"]))
			return
		}

		// delete
		_, err = target.Delete()
		if ErrorIf(c, err) {
			return
		}

		Success(c, params.ID)

	})

	// get one
	r.GET("/alert/instance/:instance_id", func(c *gin.Context) {
		var err error
		var params model.AlertInstance

		// get id
		params.ID = common.ParseInt(c.Param("instance_id"))

		// check ID
		if params.ID == 0 {
			ErrorIf(c, fmt.Errorf(common.MSG["err.invalid_zero_id"]))
			return
		}

		// target
		target := model.AlertInstance{ID: params.ID}

		// get first one
		if params, err = target.GetFirst(); err != nil {
			ErrorIf(c, fmt.Errorf(common.MSG["err.instance_not_exists"]))
			return
		}

		Success(c, params.GetInstanceExt())

	})

	// get list
	r.GET("/alert/instances", func(c *gin.Context) {
		var params model.AlertInstance
		Success(c, params.GetInstanceExtList())
	})

	// get rules in instance
	r.GET("/alert/instance/:instance_id/rules", func(c *gin.Context) {
		var params model.AlertInstance

		// get id
		params.ID = common.ParseInt(c.Param("instance_id"))

		// check ID
		if params.ID == 0 {
			ErrorIf(c, fmt.Errorf(common.MSG["err.invalid_zero_id"]))
			return
		}

		// target
		target := model.AlertInstance{ID: params.ID}

		// check exists
		if !target.Exist() {
			ErrorIf(c, fmt.Errorf(common.MSG["err.instance_not_exists"]))
			return
		}

		Success(c, params.GetRules())

	})

	// get skip rules in instance
	r.GET("/alert/instance/:instance_id/skip_rules", func(c *gin.Context) {
		var params model.AlertInstance

		// get id
		params.ID = common.ParseInt(c.Param("instance_id"))

		// check ID
		if params.ID == 0 {
			ErrorIf(c, fmt.Errorf(common.MSG["err.invalid_zero_id"]))
			return
		}

		// target
		target := model.AlertInstance{ID: params.ID}

		// check exists
		if !target.Exist() {
			ErrorIf(c, fmt.Errorf(common.MSG["err.instance_not_exists"]))
			return
		}

		Success(c, params.GetSkipRules())
	})
}
