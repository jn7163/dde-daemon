/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
        "os/exec"
        "strings"
)

type areaRange struct {
        x1      int32
        x2      int32
        y1      int32
        y2      int32
}

const (
        DISTANCE         = int32(10)
        MOTION_INSIDE    = "inside"
        EDGE_TOPLEFT     = "TopLeft"
        EDGE_BOTTOMLEFT  = "BottomLeft"
        EDGE_TOPRIGHT    = "TopRight"
        EDGE_BOTTOMRIGHT = "BottomRight"
        ACTION_WORKSPACE = "workspace"
)

var (
        areaId  int32

        topLeftArea     areaRange
        bottomLeftArea  areaRange
        topRightArea    areaRange
        bottomRightArea areaRange
)

func registerZoneArea() {
        //mutex.Lock()
        //defer mutex.Unlock()

        println("Register ...")
        rect := dspObj.PrimaryRect.Get()
        x1 := int32(rect[0].(int16))
        y1 := int32(rect[1].(int16))
        x2 := int32(rect[2].(uint16))
        y2 := int32(rect[3].(uint16))

        topLeftArea = areaRange{x1, x1 + DISTANCE, y1, y1 + DISTANCE}
        bottomLeftArea = areaRange{x1, x1 + DISTANCE, y2 - DISTANCE, y2}
        topRightArea = areaRange{x2 - DISTANCE, x2, y1, y1 + DISTANCE}
        bottomRightArea = areaRange{x2 - DISTANCE, x2, y2 - DISTANCE, y2}

        //areaId = areaObj.RegisterArea(area)
        var err error
        areaId, err = areaObj.RegisterArea([][]interface{}{
                getInterfaceArray(topLeftArea),
                getInterfaceArray(bottomLeftArea),
                getInterfaceArray(topRightArea),
                getInterfaceArray(bottomRightArea)})
        if err != nil {
                logObj.Info("Register area failed: ", err)
                return
        }
        println("Register end")
}

func unregisterZoneArea() {
        areaObj.UnregisterArea(areaId)
}

func execEdgeAction(edge string) {
        if action, ok := edgeActionMap[edge]; ok {
                if action == ACTION_WORKSPACE {
                        return
                }
                strs := strings.Split(action, " ")
                l := len(strs)
                if l < 0 {
                        return
                }

                argv := []string{}
                for i := 1; i < l; i++ {
                        argv = append(argv, strs[i])
                }
                go exec.Command(strs[0], argv...).Run()
        }
}

func isInArea(x, y int32, area areaRange) bool {
        if x >= (area.x1) && x < (area.x2) &&
                y >= (area.y1) && y < (area.y2) {
                return true
        }

        return false
}

func getInterfaceArray(area areaRange) []interface{} {
        return []interface{}{area.x1, area.x2, area.y1, area.y2}
}

func newManager() *Manager {
        println("New Manager...")
        m := &Manager{}

        registerZoneArea()
        println("New Manager end")
        m.listenSignal()

        return m
}