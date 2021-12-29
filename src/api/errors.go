/*
highload-wallet-api – API wrapper over high-load TON wallet smart contract

Copyright (C) 2021 Alexander Gapak

This file is part of highload-wallet-api.

highload-wallet-api is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

highload-wallet-api is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with highload-wallet-api.  If not, see <https://www.gnu.org/licenses/>.
*/

package api

type err struct {
	Code      int    `json:"code"`
	ErrorText string `json:"error_text"`
	Status    bool   `json:"status"`
}

var httperrs = struct {
	InternalServerError500 err
}{
	InternalServerError500: err{500, "500 Internal Server Error", false},
}

var apierrs = struct {
	ErrorJsonData err
	ErrorJsonRpc  err
}{
	ErrorJsonData: err{1, "Error Json Data", false},
	ErrorJsonRpc:  err{2, "Json RPC return an unexpected error", false},
}
