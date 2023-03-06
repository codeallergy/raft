/*
 * Copyright (c) 2022-2023 Zander Schwid & Co. LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package api

import (
	"github.com/codeallergy/glue"
	"github.com/codeallergy/sprint"
	"github.com/codeallergy/raftpb"
	"github.com/hashicorp/raft"
	"google.golang.org/grpc"
)

type RaftClientPool interface {
	glue.InitializingBean
	glue.DisposableBean

	GetAPIEndpoint(raftAddress string) (string, error)

	GetAPIConn(raftAddress raft.ServerAddress) (*grpc.ClientConn, error)

	Close() error

}

/**
Finite State Machine Response
 */
type FSMResponse struct {
	Status   *raftpb.Status
	Err      error
}

type RaftService interface {
	glue.InitializingBean
	raft.FSM

}

type RaftServer interface {
	sprint.Server
	sprint.Component

	Transport() (raft.Transport, bool)

	Raft() (*raft.Raft, bool)

}
