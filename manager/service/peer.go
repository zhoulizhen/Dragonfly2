/*
 *     Copyright 2022 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"
	"strings"

	pkgredis "d7y.io/dragonfly/v2/pkg/redis"
)

func (s *service) GetPeers(ctx context.Context) ([]string, error) {
	rawKeys, err := s.rdb.Keys(ctx, pkgredis.MakeKeyInManager(pkgredis.PeersNamespace, "*")).Result()
	if err != nil {
		return nil, err
	}

	var peers []string
	for _, rawKey := range rawKeys {
		keys := strings.Split(rawKey, ":")
		if len(keys) != 3 {
			continue
		}

		peers = append(peers, keys[2])
	}

	return peers, nil
}
