// Cadmium - multiprotocol crossplatform messenger.
// Copyright (c) 2017, Stanislav N aka pztrn.
// Copyright (c) 2017, Cadmium developers and contributors.
//
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package eventerinterface

import (
	// local
	"github.com/pztrn/cadmium/eventer/event"
)

type EventerInterface interface {
	AddEventHandler(event_name string, handler *event.Event)
	Initialize()
	LaunchEvent(event_name string, data map[string]string)
	RemoveEventHandler(event_name string, handler_name string)
}