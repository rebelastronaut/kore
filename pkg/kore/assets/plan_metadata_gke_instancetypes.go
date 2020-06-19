/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package assets

// GKEInstanceTypes is the list of possible instance types for GKE clusters. Will be
// replaced by info sourced from cloudinfo shortly.
const GKEInstanceTypes = `{
    "products": [
        {
            "category": "General purpose",
            "type": "c2-standard-16",
            "onDemandPrice": 1.000992,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.20016
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.20016
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.20016
                }
            ],
            "cpusPerVm": 16,
            "memPerVm": 64,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "16",
                "instanceTypeCategory": "General purpose",
                "memory": "64",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "g1-small",
            "onDemandPrice": 0.032400000000000005,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.008400000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.008400000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.008400000000000001
                }
            ],
            "cpusPerVm": 1,
            "memPerVm": 1.69921875,
            "gpusPerVm": 0,
            "ntwPerf": "2 Gbit/s",
            "ntwPerfCategory": "low",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "1",
                "instanceTypeCategory": "General purpose",
                "memory": "1.69921875",
                "networkPerfCategory": "low"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n1-highmem-4",
            "onDemandPrice": 0.304828,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.06094000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.06094000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.06094000000000001
                }
            ],
            "cpusPerVm": 4,
            "memPerVm": 26,
            "gpusPerVm": 0,
            "ntwPerf": "8 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "4",
                "instanceTypeCategory": "Memory optimized",
                "memory": "26",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-4",
            "onDemandPrice": 0.24479,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.04895000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.04895000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.04895000000000001
                }
            ],
            "cpusPerVm": 4,
            "memPerVm": 15,
            "gpusPerVm": 0,
            "ntwPerf": "8 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "4",
                "instanceTypeCategory": "General purpose",
                "memory": "15",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "c2-standard-30",
            "onDemandPrice": 1.87686,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.3753
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.3753
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.3753
                }
            ],
            "cpusPerVm": 30,
            "memPerVm": 120,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "30",
                "instanceTypeCategory": "General purpose",
                "memory": "120",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "c2-standard-60",
            "onDemandPrice": 3.75372,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.7506
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.7506
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.7506
                }
            ],
            "cpusPerVm": 60,
            "memPerVm": 240,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "60",
                "instanceTypeCategory": "General purpose",
                "memory": "240",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n1-highcpu-32",
            "onDemandPrice": 1.4605493339843751,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.29219178710937505
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.29219178710937505
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.29219178710937505
                }
            ],
            "cpusPerVm": 32,
            "memPerVm": 28.7998046875,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "32",
                "instanceTypeCategory": "Compute optimized",
                "memory": "28.7998046875",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-2",
            "onDemandPrice": 0.168788,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.033740000000000006
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.033740000000000006
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.033740000000000006
                }
            ],
            "cpusPerVm": 2,
            "memPerVm": 16,
            "gpusPerVm": 0,
            "ntwPerf": "4 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "2",
                "instanceTypeCategory": "Memory optimized",
                "memory": "16",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-4",
            "onDemandPrice": 0.337576,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.06748000000000001
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.06748000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.06748000000000001
                }
            ],
            "cpusPerVm": 4,
            "memPerVm": 32,
            "gpusPerVm": 0,
            "ntwPerf": "8 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "4",
                "instanceTypeCategory": "Memory optimized",
                "memory": "32",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-48",
            "onDemandPrice": 4.050912,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.80976
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.80976
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.80976
                }
            ],
            "cpusPerVm": 48,
            "memPerVm": 384,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "48",
                "instanceTypeCategory": "Memory optimized",
                "memory": "384",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-16",
            "onDemandPrice": 1.350304,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.26992000000000005
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.26992000000000005
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.26992000000000005
                }
            ],
            "cpusPerVm": 16,
            "memPerVm": 128,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "16",
                "instanceTypeCategory": "Memory optimized",
                "memory": "128",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-48",
            "onDemandPrice": 3.0029760000000003,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.6004800000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.6004800000000001
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.6004800000000001
                }
            ],
            "cpusPerVm": 48,
            "memPerVm": 192,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "48",
                "instanceTypeCategory": "General purpose",
                "memory": "192",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "c2-standard-8",
            "onDemandPrice": 0.500496,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.10008
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.10008
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.10008
                }
            ],
            "cpusPerVm": 8,
            "memPerVm": 32,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "8",
                "instanceTypeCategory": "General purpose",
                "memory": "32",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n1-highcpu-16",
            "onDemandPrice": 0.73027733203125,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.14609642578125
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.14609642578125
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.14609642578125
                }
            ],
            "cpusPerVm": 16,
            "memPerVm": 14.400390625,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "16",
                "instanceTypeCategory": "Compute optimized",
                "memory": "14.400390625",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n1-highcpu-64",
            "onDemandPrice": 2.9210986679687503,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.5843835742187501
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.5843835742187501
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.5843835742187501
                }
            ],
            "cpusPerVm": 64,
            "memPerVm": 57.599609375,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "64",
                "instanceTypeCategory": "Compute optimized",
                "memory": "57.599609375",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-8",
            "onDemandPrice": 0.48958,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.09790000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.09790000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.09790000000000001
                }
            ],
            "cpusPerVm": 8,
            "memPerVm": 30,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "8",
                "instanceTypeCategory": "General purpose",
                "memory": "30",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "m1-ultramem-80",
            "onDemandPrice": 13.748676,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 2.74698
                },
                {
                    "zone": "europe-west2-a",
                    "price": 2.74698
                },
                {
                    "zone": "europe-west2-b",
                    "price": 2.74698
                }
            ],
            "cpusPerVm": 80,
            "memPerVm": 1922,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "80",
                "instanceTypeCategory": "General purpose",
                "memory": "1922",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n1-highmem-8",
            "onDemandPrice": 0.609656,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.12188000000000002
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.12188000000000002
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.12188000000000002
                }
            ],
            "cpusPerVm": 8,
            "memPerVm": 52,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "8",
                "instanceTypeCategory": "Memory optimized",
                "memory": "52",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-16",
            "onDemandPrice": 0.97916,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.19580000000000003
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.19580000000000003
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.19580000000000003
                }
            ],
            "cpusPerVm": 16,
            "memPerVm": 60,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "16",
                "instanceTypeCategory": "General purpose",
                "memory": "60",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-2",
            "onDemandPrice": 0.122395,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.024475000000000004
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.024475000000000004
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.024475000000000004
                }
            ],
            "cpusPerVm": 2,
            "memPerVm": 7.5,
            "gpusPerVm": 0,
            "ntwPerf": "4 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "2",
                "instanceTypeCategory": "General purpose",
                "memory": "7.5",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-32",
            "onDemandPrice": 2.001984,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.40032
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.40032
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.40032
                }
            ],
            "cpusPerVm": 32,
            "memPerVm": 128,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "32",
                "instanceTypeCategory": "General purpose",
                "memory": "128",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n1-highcpu-4",
            "onDemandPrice": 0.18256666796875,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.03652357421875001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.03652357421875001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.03652357421875001
                }
            ],
            "cpusPerVm": 4,
            "memPerVm": 3.599609375,
            "gpusPerVm": 0,
            "ntwPerf": "8 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "4",
                "instanceTypeCategory": "Compute optimized",
                "memory": "3.599609375",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n1-highcpu-8",
            "onDemandPrice": 0.365138666015625,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.073048212890625
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.073048212890625
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.073048212890625
                }
            ],
            "cpusPerVm": 8,
            "memPerVm": 7.2001953125,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "8",
                "instanceTypeCategory": "Compute optimized",
                "memory": "7.2001953125",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-16",
            "onDemandPrice": 0.739008,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.14784000000000003
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.14784000000000003
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.14784000000000003
                }
            ],
            "cpusPerVm": 16,
            "memPerVm": 16,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "16",
                "instanceTypeCategory": "Compute optimized",
                "memory": "16",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-64",
            "onDemandPrice": 2.956032,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.5913600000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.5913600000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.5913600000000001
                }
            ],
            "cpusPerVm": 64,
            "memPerVm": 64,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "64",
                "instanceTypeCategory": "Compute optimized",
                "memory": "64",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-8",
            "onDemandPrice": 0.500496,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.10008
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.10008
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.10008
                }
            ],
            "cpusPerVm": 8,
            "memPerVm": 32,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "8",
                "instanceTypeCategory": "General purpose",
                "memory": "32",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-80",
            "onDemandPrice": 5.0049600000000005,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 1.0008000000000001
                },
                {
                    "zone": "europe-west2-a",
                    "price": 1.0008000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 1.0008000000000001
                }
            ],
            "cpusPerVm": 80,
            "memPerVm": 320,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "80",
                "instanceTypeCategory": "General purpose",
                "memory": "320",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "f1-micro",
            "onDemandPrice": 0.009600000000000001,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.004200000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.004200000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.004200000000000001
                }
            ],
            "cpusPerVm": 1,
            "memPerVm": 0.599609375,
            "gpusPerVm": 0,
            "ntwPerf": "2 Gbit/s",
            "ntwPerfCategory": "low",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "1",
                "instanceTypeCategory": "General purpose",
                "memory": "0.599609375",
                "networkPerfCategory": "low"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-2",
            "onDemandPrice": 0.092376,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.018480000000000003
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.018480000000000003
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.018480000000000003
                }
            ],
            "cpusPerVm": 2,
            "memPerVm": 2,
            "gpusPerVm": 0,
            "ntwPerf": "4 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "2",
                "instanceTypeCategory": "Compute optimized",
                "memory": "2",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-80",
            "onDemandPrice": 3.69504,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.7392000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.7392000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.7392000000000001
                }
            ],
            "cpusPerVm": 80,
            "memPerVm": 80,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "80",
                "instanceTypeCategory": "Compute optimized",
                "memory": "80",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-64",
            "onDemandPrice": 5.401216,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 1.0796800000000002
                },
                {
                    "zone": "europe-west2-c",
                    "price": 1.0796800000000002
                },
                {
                    "zone": "europe-west2-a",
                    "price": 1.0796800000000002
                }
            ],
            "cpusPerVm": 64,
            "memPerVm": 512,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "64",
                "instanceTypeCategory": "Memory optimized",
                "memory": "512",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-2",
            "onDemandPrice": 0.125124,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.02502
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.02502
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.02502
                }
            ],
            "cpusPerVm": 2,
            "memPerVm": 8,
            "gpusPerVm": 0,
            "ntwPerf": "4 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "2",
                "instanceTypeCategory": "General purpose",
                "memory": "8",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-8",
            "onDemandPrice": 0.675152,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.13496000000000002
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.13496000000000002
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.13496000000000002
                }
            ],
            "cpusPerVm": 8,
            "memPerVm": 64,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "8",
                "instanceTypeCategory": "Memory optimized",
                "memory": "64",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-4",
            "onDemandPrice": 0.250248,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.05004
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.05004
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.05004
                }
            ],
            "cpusPerVm": 4,
            "memPerVm": 16,
            "gpusPerVm": 0,
            "ntwPerf": "8 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "4",
                "instanceTypeCategory": "General purpose",
                "memory": "16",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-48",
            "onDemandPrice": 2.2170240000000003,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.44352
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.44352
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.44352
                }
            ],
            "cpusPerVm": 48,
            "memPerVm": 48,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "48",
                "instanceTypeCategory": "Compute optimized",
                "memory": "48",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-8",
            "onDemandPrice": 0.369504,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.07392000000000001
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.07392000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.07392000000000001
                }
            ],
            "cpusPerVm": 8,
            "memPerVm": 8,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "8",
                "instanceTypeCategory": "Compute optimized",
                "memory": "8",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n1-highcpu-96",
            "onDemandPrice": 4.3816533320312505,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.8765764257812501
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.8765764257812501
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.8765764257812501
                }
            ],
            "cpusPerVm": 96,
            "memPerVm": 86.400390625,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "96",
                "instanceTypeCategory": "Compute optimized",
                "memory": "86.400390625",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n1-highmem-64",
            "onDemandPrice": 4.877248,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.9750400000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.9750400000000001
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.9750400000000001
                }
            ],
            "cpusPerVm": 64,
            "memPerVm": 416,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "64",
                "instanceTypeCategory": "Memory optimized",
                "memory": "416",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-4",
            "onDemandPrice": 0.184752,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.03696000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.03696000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.03696000000000001
                }
            ],
            "cpusPerVm": 4,
            "memPerVm": 4,
            "gpusPerVm": 0,
            "ntwPerf": "8 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "4",
                "instanceTypeCategory": "Compute optimized",
                "memory": "4",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "m1-ultramem-160",
            "onDemandPrice": 27.497352,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 5.49396
                },
                {
                    "zone": "europe-west2-b",
                    "price": 5.49396
                },
                {
                    "zone": "europe-west2-c",
                    "price": 5.49396
                }
            ],
            "cpusPerVm": 160,
            "memPerVm": 3844,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "160",
                "instanceTypeCategory": "General purpose",
                "memory": "3844",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n1-highmem-32",
            "onDemandPrice": 2.438624,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.48752000000000006
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.48752000000000006
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.48752000000000006
                }
            ],
            "cpusPerVm": 32,
            "memPerVm": 208,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "32",
                "instanceTypeCategory": "Memory optimized",
                "memory": "208",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-1",
            "onDemandPrice": 0.0611975,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.012237500000000002
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.012237500000000002
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.012237500000000002
                }
            ],
            "cpusPerVm": 1,
            "memPerVm": 3.75,
            "gpusPerVm": 0,
            "ntwPerf": "2 Gbit/s",
            "ntwPerfCategory": "low",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "1",
                "instanceTypeCategory": "General purpose",
                "memory": "3.75",
                "networkPerfCategory": "low"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-80",
            "onDemandPrice": 6.75152,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 1.3496000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 1.3496000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 1.3496000000000001
                }
            ],
            "cpusPerVm": 80,
            "memPerVm": 640,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "80",
                "instanceTypeCategory": "Memory optimized",
                "memory": "640",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "c2-standard-4",
            "onDemandPrice": 0.250248,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.05004
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.05004
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.05004
                }
            ],
            "cpusPerVm": 4,
            "memPerVm": 16,
            "gpusPerVm": 0,
            "ntwPerf": "8 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "4",
                "instanceTypeCategory": "General purpose",
                "memory": "16",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n1-highcpu-2",
            "onDemandPrice": 0.091283333984375,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.018261787109375004
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.018261787109375004
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.018261787109375004
                }
            ],
            "cpusPerVm": 2,
            "memPerVm": 1.7998046875,
            "gpusPerVm": 0,
            "ntwPerf": "4 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "2",
                "instanceTypeCategory": "Compute optimized",
                "memory": "1.7998046875",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-32",
            "onDemandPrice": 1.95832,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.39160000000000006
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.39160000000000006
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.39160000000000006
                }
            ],
            "cpusPerVm": 32,
            "memPerVm": 120,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "32",
                "instanceTypeCategory": "General purpose",
                "memory": "120",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-64",
            "onDemandPrice": 3.91664,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.7832000000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.7832000000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.7832000000000001
                }
            ],
            "cpusPerVm": 64,
            "memPerVm": 240,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "64",
                "instanceTypeCategory": "General purpose",
                "memory": "240",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n1-standard-96",
            "onDemandPrice": 5.87496,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 1.1748
                },
                {
                    "zone": "europe-west2-a",
                    "price": 1.1748
                },
                {
                    "zone": "europe-west2-b",
                    "price": 1.1748
                }
            ],
            "cpusPerVm": 96,
            "memPerVm": 360,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "96",
                "instanceTypeCategory": "General purpose",
                "memory": "360",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "m1-ultramem-40",
            "onDemandPrice": 6.874338,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 1.37349
                },
                {
                    "zone": "europe-west2-c",
                    "price": 1.37349
                },
                {
                    "zone": "europe-west2-a",
                    "price": 1.37349
                }
            ],
            "cpusPerVm": 40,
            "memPerVm": 961,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "40",
                "instanceTypeCategory": "General purpose",
                "memory": "961",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n1-highmem-16",
            "onDemandPrice": 1.219312,
            "spotPrice": [
                {
                    "zone": "europe-west2-c",
                    "price": 0.24376000000000003
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.24376000000000003
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.24376000000000003
                }
            ],
            "cpusPerVm": 16,
            "memPerVm": 104,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "16",
                "instanceTypeCategory": "Memory optimized",
                "memory": "104",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n1-highmem-96",
            "onDemandPrice": 7.315872000000001,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 1.46256
                },
                {
                    "zone": "europe-west2-c",
                    "price": 1.46256
                },
                {
                    "zone": "europe-west2-a",
                    "price": 1.46256
                }
            ],
            "cpusPerVm": 96,
            "memPerVm": 624,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "96",
                "instanceTypeCategory": "Memory optimized",
                "memory": "624",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-16",
            "onDemandPrice": 1.000992,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.20016
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.20016
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.20016
                }
            ],
            "cpusPerVm": 16,
            "memPerVm": 64,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "16",
                "instanceTypeCategory": "General purpose",
                "memory": "64",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "General purpose",
            "type": "n2-standard-64",
            "onDemandPrice": 4.003968,
            "spotPrice": [
                {
                    "zone": "europe-west2-b",
                    "price": 0.80064
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.80064
                },
                {
                    "zone": "europe-west2-a",
                    "price": 0.80064
                }
            ],
            "cpusPerVm": 64,
            "memPerVm": 256,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "64",
                "instanceTypeCategory": "General purpose",
                "memory": "256",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n1-highmem-2",
            "onDemandPrice": 0.152414,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.030470000000000004
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.030470000000000004
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.030470000000000004
                }
            ],
            "cpusPerVm": 2,
            "memPerVm": 13,
            "gpusPerVm": 0,
            "ntwPerf": "4 Gbit/s",
            "ntwPerfCategory": "medium",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "2",
                "instanceTypeCategory": "Memory optimized",
                "memory": "13",
                "networkPerfCategory": "medium"
            },
            "currentGen": false
        },
        {
            "category": "Compute optimized",
            "type": "n2-highcpu-32",
            "onDemandPrice": 1.478016,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.29568000000000005
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.29568000000000005
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.29568000000000005
                }
            ],
            "cpusPerVm": 32,
            "memPerVm": 32,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "32",
                "instanceTypeCategory": "Compute optimized",
                "memory": "32",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        },
        {
            "category": "Memory optimized",
            "type": "n2-highmem-32",
            "onDemandPrice": 2.700608,
            "spotPrice": [
                {
                    "zone": "europe-west2-a",
                    "price": 0.5398400000000001
                },
                {
                    "zone": "europe-west2-b",
                    "price": 0.5398400000000001
                },
                {
                    "zone": "europe-west2-c",
                    "price": 0.5398400000000001
                }
            ],
            "cpusPerVm": 32,
            "memPerVm": 256,
            "gpusPerVm": 0,
            "ntwPerf": "16 Gbit/s",
            "ntwPerfCategory": "extra",
            "zones": [
                "europe-west2-c",
                "europe-west2-b",
                "europe-west2-a"
            ],
            "attributes": {
                "cpu": "32",
                "instanceTypeCategory": "Memory optimized",
                "memory": "256",
                "networkPerfCategory": "extra"
            },
            "currentGen": false
        }
    ],
    "scrapingTime": "1592309997841"
}`
