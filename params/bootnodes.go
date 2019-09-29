// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://566350ad17673cb9a051d67cd4d5acbe159868b710a05a5fc5a319af8265c7aab6d98f4286cfa0942e2fe02caafdfb2f21505b959e866cd183e915a199dca8f2@163.172.163.199:30301",
	"enode://566350ad17673cb9a051d67cd4d5acbe159868b710a05a5fc5a319af8265c7aab6d98f4286cfa0942e2fe02caafdfb2f21505b959e866cd183e915a199dca8f2@163.172.163.199:40301",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// test network.
var TestnetBootnodes = []string{
	"enode://566350ad17673cb9a051d67cd4d5acbe159868b710a05a5fc5a319af8265c7aab6d98f4286cfa0942e2fe02caafdfb2f21505b959e866cd183e915a199dca8f2@163.172.163.199:30301",
	"enode://566350ad17673cb9a051d67cd4d5acbe159868b710a05a5fc5a319af8265c7aab6d98f4286cfa0942e2fe02caafdfb2f21505b959e866cd183e915a199dca8f2@163.172.163.199:40301",
}
