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

package models


type PanelMonitorData struct {

    /* 聚合方式 (Optional) */
    Aggregator string `json:"aggregator"`

    /* 监控数据 (Optional) */
    DataPoint []DataPoint `json:"dataPoint"`

    /* 采样方式 (Optional) */
    Downsample string `json:"downsample"`

    /* 采样周期 (Optional) */
    DownsamplePeriod string `json:"downsamplePeriod"`

    /* metric (Optional) */
    Metric string `json:"metric"`

    /* metric名字 (Optional) */
    MetricName string `json:"metricName"`

    /* 实例id，汇总图无 (Optional) */
    ResourceId string `json:"resourceId"`

    /* 实例名称，汇总图无；标签资源该值为实例id (Optional) */
    ResourceName string `json:"resourceName"`

    /* 该资源的维度值 (Optional) */
    Tags interface{} `json:"tags"`

    /* metric单位 (Optional) */
    Unit string `json:"unit"`
}