/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/

package util

//DataPresaleDBInfo struct
type DataPresaleDBInfo struct {
	DataPresaleID string `json:"data_presale_id"`
	UserName      string `json:"user_name"`
	AssetID       string `json:"asset_id"`
	DataReqID     string `json:"data_req_id"`
	Consumer      string `json:"consumer"`
}
