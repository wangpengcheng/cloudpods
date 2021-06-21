// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"
)

type DescribeSnapshotRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 快照ID  */
    SnapshotId string `json:"snapshotId"`
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSnapshotRequest(
    regionId string,
    snapshotId string,
) *DescribeSnapshotRequest {

	return &DescribeSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/snapshots/{snapshotId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SnapshotId: snapshotId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 */
func NewDescribeSnapshotRequestWithAllParams(
    regionId string,
    snapshotId string,
) *DescribeSnapshotRequest {

    return &DescribeSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SnapshotId: snapshotId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSnapshotRequestWithoutParam() *DescribeSnapshotRequest {

    return &DescribeSnapshotRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeSnapshotRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param snapshotId: 快照ID(Required) */
func (r *DescribeSnapshotRequest) SetSnapshotId(snapshotId string) {
    r.SnapshotId = snapshotId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSnapshotRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSnapshotResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSnapshotResult `json:"result"`
}

type DescribeSnapshotResult struct {
    Snapshot disk.Snapshot `json:"snapshot"`
}