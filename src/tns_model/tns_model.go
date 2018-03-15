/*******************************************************************************
 * Copyright 2018 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

package models

import "gopkg.in/mgo.v2/bson"

// Represents a TNS data, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type TNSdata struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Topic       string        `bson:"topic" json:"topic"`
	Endpoint    string        `bson:"endpoint" json:"endpoint"`
	Schema      string        `bson:"schema" json:"schema"`
}
