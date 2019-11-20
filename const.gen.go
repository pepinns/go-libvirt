// Copyright 2018 The go-libvirt Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// WARNING: This file has automatically been generated
// by https://git.io/c-for-go. DO NOT EDIT.

package libvirt

const (
	// ExportVar as defined in libvirt/libvirt-common.h:57
	ExportVar = 0
	// TypedParamFieldLength as defined in libvirt/libvirt-common.h:170
	TypedParamFieldLength = 80
	// SecurityLabelBuflen as defined in libvirt/libvirt-host.h:84
	SecurityLabelBuflen = 4097
	// SecurityModelBuflen as defined in libvirt/libvirt-host.h:112
	SecurityModelBuflen = 257
	// SecurityDoiBuflen as defined in libvirt/libvirt-host.h:119
	SecurityDoiBuflen = 257
	// NodeCPUStatsFieldLength as defined in libvirt/libvirt-host.h:176
	NodeCPUStatsFieldLength = 80
	// NodeCPUStatsKernel as defined in libvirt/libvirt-host.h:193
	NodeCPUStatsKernel = "kernel"
	// NodeCPUStatsUser as defined in libvirt/libvirt-host.h:201
	NodeCPUStatsUser = "user"
	// NodeCPUStatsIdle as defined in libvirt/libvirt-host.h:209
	NodeCPUStatsIdle = "idle"
	// NodeCPUStatsIowait as defined in libvirt/libvirt-host.h:217
	NodeCPUStatsIowait = "iowait"
	// NodeCPUStatsIntr as defined in libvirt/libvirt-host.h:225
	NodeCPUStatsIntr = "intr"
	// NodeCPUStatsUtilization as defined in libvirt/libvirt-host.h:234
	NodeCPUStatsUtilization = "utilization"
	// NodeMemoryStatsFieldLength as defined in libvirt/libvirt-host.h:254
	NodeMemoryStatsFieldLength = 80
	// NodeMemoryStatsTotal as defined in libvirt/libvirt-host.h:271
	NodeMemoryStatsTotal = "total"
	// NodeMemoryStatsFree as defined in libvirt/libvirt-host.h:280
	NodeMemoryStatsFree = "free"
	// NodeMemoryStatsBuffers as defined in libvirt/libvirt-host.h:288
	NodeMemoryStatsBuffers = "buffers"
	// NodeMemoryStatsCached as defined in libvirt/libvirt-host.h:296
	NodeMemoryStatsCached = "cached"
	// NodeMemorySharedPagesToScan as defined in libvirt/libvirt-host.h:317
	NodeMemorySharedPagesToScan = "shm_pages_to_scan"
	// NodeMemorySharedSleepMillisecs as defined in libvirt/libvirt-host.h:325
	NodeMemorySharedSleepMillisecs = "shm_sleep_millisecs"
	// NodeMemorySharedPagesShared as defined in libvirt/libvirt-host.h:333
	NodeMemorySharedPagesShared = "shm_pages_shared"
	// NodeMemorySharedPagesSharing as defined in libvirt/libvirt-host.h:341
	NodeMemorySharedPagesSharing = "shm_pages_sharing"
	// NodeMemorySharedPagesUnshared as defined in libvirt/libvirt-host.h:349
	NodeMemorySharedPagesUnshared = "shm_pages_unshared"
	// NodeMemorySharedPagesVolatile as defined in libvirt/libvirt-host.h:357
	NodeMemorySharedPagesVolatile = "shm_pages_volatile"
	// NodeMemorySharedFullScans as defined in libvirt/libvirt-host.h:365
	NodeMemorySharedFullScans = "shm_full_scans"
	// NodeMemorySharedMergeAcrossNodes as defined in libvirt/libvirt-host.h:377
	NodeMemorySharedMergeAcrossNodes = "shm_merge_across_nodes"
	// NodeSevPdh as defined in libvirt/libvirt-host.h:445
	NodeSevPdh = "pdh"
	// NodeSevCertChain as defined in libvirt/libvirt-host.h:454
	NodeSevCertChain = "cert-chain"
	// NodeSevCbitpos as defined in libvirt/libvirt-host.h:461
	NodeSevCbitpos = "cbitpos"
	// NodeSevReducedPhysBits as defined in libvirt/libvirt-host.h:469
	NodeSevReducedPhysBits = "reduced-phys-bits"
	// UUIDBuflen as defined in libvirt/libvirt-host.h:554
	UUIDBuflen = 16
	// UUIDStringBuflen as defined in libvirt/libvirt-host.h:563
	UUIDStringBuflen = 37
	// ConnectIdentityUserName as defined in libvirt/libvirt-host.h:588
	ConnectIdentityUserName = "user-name"
	// ConnectIdentityUnixUserID as defined in libvirt/libvirt-host.h:595
	ConnectIdentityUnixUserID = "unix-user-id"
	// ConnectIdentityGroupName as defined in libvirt/libvirt-host.h:602
	ConnectIdentityGroupName = "group-name"
	// ConnectIdentityUnixGroupID as defined in libvirt/libvirt-host.h:609
	ConnectIdentityUnixGroupID = "unix-group-id"
	// ConnectIdentityProcessID as defined in libvirt/libvirt-host.h:616
	ConnectIdentityProcessID = "process-id"
	// ConnectIdentityProcessTime as defined in libvirt/libvirt-host.h:627
	ConnectIdentityProcessTime = "process-time"
	// ConnectIdentitySaslUserName as defined in libvirt/libvirt-host.h:634
	ConnectIdentitySaslUserName = "sasl-user-name"
	// ConnectIdentityX509DistinguishedName as defined in libvirt/libvirt-host.h:641
	ConnectIdentityX509DistinguishedName = "x509-distinguished-name"
	// ConnectIdentitySelinuxContext as defined in libvirt/libvirt-host.h:648
	ConnectIdentitySelinuxContext = "selinux-context"
	// DomainSchedulerCPUShares as defined in libvirt/libvirt-domain.h:316
	DomainSchedulerCPUShares = "cpu_shares"
	// DomainSchedulerGlobalPeriod as defined in libvirt/libvirt-domain.h:324
	DomainSchedulerGlobalPeriod = "global_period"
	// DomainSchedulerGlobalQuota as defined in libvirt/libvirt-domain.h:332
	DomainSchedulerGlobalQuota = "global_quota"
	// DomainSchedulerVCPUPeriod as defined in libvirt/libvirt-domain.h:340
	DomainSchedulerVCPUPeriod = "vcpu_period"
	// DomainSchedulerVCPUQuota as defined in libvirt/libvirt-domain.h:348
	DomainSchedulerVCPUQuota = "vcpu_quota"
	// DomainSchedulerEmulatorPeriod as defined in libvirt/libvirt-domain.h:357
	DomainSchedulerEmulatorPeriod = "emulator_period"
	// DomainSchedulerEmulatorQuota as defined in libvirt/libvirt-domain.h:366
	DomainSchedulerEmulatorQuota = "emulator_quota"
	// DomainSchedulerIothreadPeriod as defined in libvirt/libvirt-domain.h:374
	DomainSchedulerIothreadPeriod = "iothread_period"
	// DomainSchedulerIothreadQuota as defined in libvirt/libvirt-domain.h:382
	DomainSchedulerIothreadQuota = "iothread_quota"
	// DomainSchedulerWeight as defined in libvirt/libvirt-domain.h:390
	DomainSchedulerWeight = "weight"
	// DomainSchedulerCap as defined in libvirt/libvirt-domain.h:398
	DomainSchedulerCap = "cap"
	// DomainSchedulerReservation as defined in libvirt/libvirt-domain.h:406
	DomainSchedulerReservation = "reservation"
	// DomainSchedulerLimit as defined in libvirt/libvirt-domain.h:414
	DomainSchedulerLimit = "limit"
	// DomainSchedulerShares as defined in libvirt/libvirt-domain.h:422
	DomainSchedulerShares = "shares"
	// DomainBlockStatsFieldLength as defined in libvirt/libvirt-domain.h:480
	DomainBlockStatsFieldLength = 80
	// DomainBlockStatsReadBytes as defined in libvirt/libvirt-domain.h:488
	DomainBlockStatsReadBytes = "rd_bytes"
	// DomainBlockStatsReadReq as defined in libvirt/libvirt-domain.h:496
	DomainBlockStatsReadReq = "rd_operations"
	// DomainBlockStatsReadTotalTimes as defined in libvirt/libvirt-domain.h:504
	DomainBlockStatsReadTotalTimes = "rd_total_times"
	// DomainBlockStatsWriteBytes as defined in libvirt/libvirt-domain.h:512
	DomainBlockStatsWriteBytes = "wr_bytes"
	// DomainBlockStatsWriteReq as defined in libvirt/libvirt-domain.h:520
	DomainBlockStatsWriteReq = "wr_operations"
	// DomainBlockStatsWriteTotalTimes as defined in libvirt/libvirt-domain.h:528
	DomainBlockStatsWriteTotalTimes = "wr_total_times"
	// DomainBlockStatsFlushReq as defined in libvirt/libvirt-domain.h:536
	DomainBlockStatsFlushReq = "flush_operations"
	// DomainBlockStatsFlushTotalTimes as defined in libvirt/libvirt-domain.h:544
	DomainBlockStatsFlushTotalTimes = "flush_total_times"
	// DomainBlockStatsErrs as defined in libvirt/libvirt-domain.h:551
	DomainBlockStatsErrs = "errs"
	// MigrateParamURI as defined in libvirt/libvirt-domain.h:868
	MigrateParamURI = "migrate_uri"
	// MigrateParamDestName as defined in libvirt/libvirt-domain.h:878
	MigrateParamDestName = "destination_name"
	// MigrateParamDestXML as defined in libvirt/libvirt-domain.h:897
	MigrateParamDestXML = "destination_xml"
	// MigrateParamPersistXML as defined in libvirt/libvirt-domain.h:912
	MigrateParamPersistXML = "persistent_xml"
	// MigrateParamBandwidth as defined in libvirt/libvirt-domain.h:922
	MigrateParamBandwidth = "bandwidth"
	// MigrateParamBandwidthPostcopy as defined in libvirt/libvirt-domain.h:931
	MigrateParamBandwidthPostcopy = "bandwidth.postcopy"
	// MigrateParamGraphicsURI as defined in libvirt/libvirt-domain.h:952
	MigrateParamGraphicsURI = "graphics_uri"
	// MigrateParamListenAddress as defined in libvirt/libvirt-domain.h:963
	MigrateParamListenAddress = "listen_address"
	// MigrateParamMigrateDisks as defined in libvirt/libvirt-domain.h:972
	MigrateParamMigrateDisks = "migrate_disks"
	// MigrateParamDisksPort as defined in libvirt/libvirt-domain.h:982
	MigrateParamDisksPort = "disks_port"
	// MigrateParamCompression as defined in libvirt/libvirt-domain.h:992
	MigrateParamCompression = "compression"
	// MigrateParamCompressionMtLevel as defined in libvirt/libvirt-domain.h:1001
	MigrateParamCompressionMtLevel = "compression.mt.level"
	// MigrateParamCompressionMtThreads as defined in libvirt/libvirt-domain.h:1009
	MigrateParamCompressionMtThreads = "compression.mt.threads"
	// MigrateParamCompressionMtDthreads as defined in libvirt/libvirt-domain.h:1017
	MigrateParamCompressionMtDthreads = "compression.mt.dthreads"
	// MigrateParamCompressionXbzrleCache as defined in libvirt/libvirt-domain.h:1025
	MigrateParamCompressionXbzrleCache = "compression.xbzrle.cache"
	// MigrateParamAutoConvergeInitial as defined in libvirt/libvirt-domain.h:1034
	MigrateParamAutoConvergeInitial = "auto_converge.initial"
	// MigrateParamAutoConvergeIncrement as defined in libvirt/libvirt-domain.h:1044
	MigrateParamAutoConvergeIncrement = "auto_converge.increment"
	// MigrateParamParallelConnections as defined in libvirt/libvirt-domain.h:1052
	MigrateParamParallelConnections = "parallel.connections"
	// DomainCPUStatsCputime as defined in libvirt/libvirt-domain.h:1312
	DomainCPUStatsCputime = "cpu_time"
	// DomainCPUStatsUsertime as defined in libvirt/libvirt-domain.h:1318
	DomainCPUStatsUsertime = "user_time"
	// DomainCPUStatsSystemtime as defined in libvirt/libvirt-domain.h:1324
	DomainCPUStatsSystemtime = "system_time"
	// DomainCPUStatsVcputime as defined in libvirt/libvirt-domain.h:1331
	DomainCPUStatsVcputime = "vcpu_time"
	// DomainBlkioWeight as defined in libvirt/libvirt-domain.h:1360
	DomainBlkioWeight = "weight"
	// DomainBlkioDeviceWeight as defined in libvirt/libvirt-domain.h:1370
	DomainBlkioDeviceWeight = "device_weight"
	// DomainBlkioDeviceReadIops as defined in libvirt/libvirt-domain.h:1381
	DomainBlkioDeviceReadIops = "device_read_iops_sec"
	// DomainBlkioDeviceWriteIops as defined in libvirt/libvirt-domain.h:1392
	DomainBlkioDeviceWriteIops = "device_write_iops_sec"
	// DomainBlkioDeviceReadBps as defined in libvirt/libvirt-domain.h:1403
	DomainBlkioDeviceReadBps = "device_read_bytes_sec"
	// DomainBlkioDeviceWriteBps as defined in libvirt/libvirt-domain.h:1414
	DomainBlkioDeviceWriteBps = "device_write_bytes_sec"
	// DomainMemoryParamUnlimited as defined in libvirt/libvirt-domain.h:1433
	DomainMemoryParamUnlimited = 9007199254740991
	// DomainMemoryHardLimit as defined in libvirt/libvirt-domain.h:1442
	DomainMemoryHardLimit = "hard_limit"
	// DomainMemorySoftLimit as defined in libvirt/libvirt-domain.h:1451
	DomainMemorySoftLimit = "soft_limit"
	// DomainMemoryMinGuarantee as defined in libvirt/libvirt-domain.h:1460
	DomainMemoryMinGuarantee = "min_guarantee"
	// DomainMemorySwapHardLimit as defined in libvirt/libvirt-domain.h:1470
	DomainMemorySwapHardLimit = "swap_hard_limit"
	// DomainNumaNodeset as defined in libvirt/libvirt-domain.h:1515
	DomainNumaNodeset = "numa_nodeset"
	// DomainNumaMode as defined in libvirt/libvirt-domain.h:1523
	DomainNumaMode = "numa_mode"
	// DomainBandwidthInAverage as defined in libvirt/libvirt-domain.h:1639
	DomainBandwidthInAverage = "inbound.average"
	// DomainBandwidthInPeak as defined in libvirt/libvirt-domain.h:1646
	DomainBandwidthInPeak = "inbound.peak"
	// DomainBandwidthInBurst as defined in libvirt/libvirt-domain.h:1653
	DomainBandwidthInBurst = "inbound.burst"
	// DomainBandwidthInFloor as defined in libvirt/libvirt-domain.h:1660
	DomainBandwidthInFloor = "inbound.floor"
	// DomainBandwidthOutAverage as defined in libvirt/libvirt-domain.h:1667
	DomainBandwidthOutAverage = "outbound.average"
	// DomainBandwidthOutPeak as defined in libvirt/libvirt-domain.h:1674
	DomainBandwidthOutPeak = "outbound.peak"
	// DomainBandwidthOutBurst as defined in libvirt/libvirt-domain.h:1681
	DomainBandwidthOutBurst = "outbound.burst"
	// DomainIothreadPollMaxNs as defined in libvirt/libvirt-domain.h:1981
	DomainIothreadPollMaxNs = "poll_max_ns"
	// DomainIothreadPollGrow as defined in libvirt/libvirt-domain.h:1991
	DomainIothreadPollGrow = "poll_grow"
	// DomainIothreadPollShrink as defined in libvirt/libvirt-domain.h:2002
	DomainIothreadPollShrink = "poll_shrink"
	// PerfParamCmt as defined in libvirt/libvirt-domain.h:2193
	PerfParamCmt = "cmt"
	// PerfParamMbmt as defined in libvirt/libvirt-domain.h:2204
	PerfParamMbmt = "mbmt"
	// PerfParamMbml as defined in libvirt/libvirt-domain.h:2214
	PerfParamMbml = "mbml"
	// PerfParamCacheMisses as defined in libvirt/libvirt-domain.h:2224
	PerfParamCacheMisses = "cache_misses"
	// PerfParamCacheReferences as defined in libvirt/libvirt-domain.h:2234
	PerfParamCacheReferences = "cache_references"
	// PerfParamInstructions as defined in libvirt/libvirt-domain.h:2244
	PerfParamInstructions = "instructions"
	// PerfParamCPUCycles as defined in libvirt/libvirt-domain.h:2254
	PerfParamCPUCycles = "cpu_cycles"
	// PerfParamBranchInstructions as defined in libvirt/libvirt-domain.h:2264
	PerfParamBranchInstructions = "branch_instructions"
	// PerfParamBranchMisses as defined in libvirt/libvirt-domain.h:2274
	PerfParamBranchMisses = "branch_misses"
	// PerfParamBusCycles as defined in libvirt/libvirt-domain.h:2284
	PerfParamBusCycles = "bus_cycles"
	// PerfParamStalledCyclesFrontend as defined in libvirt/libvirt-domain.h:2295
	PerfParamStalledCyclesFrontend = "stalled_cycles_frontend"
	// PerfParamStalledCyclesBackend as defined in libvirt/libvirt-domain.h:2306
	PerfParamStalledCyclesBackend = "stalled_cycles_backend"
	// PerfParamRefCPUCycles as defined in libvirt/libvirt-domain.h:2317
	PerfParamRefCPUCycles = "ref_cpu_cycles"
	// PerfParamCPUClock as defined in libvirt/libvirt-domain.h:2328
	PerfParamCPUClock = "cpu_clock"
	// PerfParamTaskClock as defined in libvirt/libvirt-domain.h:2339
	PerfParamTaskClock = "task_clock"
	// PerfParamPageFaults as defined in libvirt/libvirt-domain.h:2349
	PerfParamPageFaults = "page_faults"
	// PerfParamContextSwitches as defined in libvirt/libvirt-domain.h:2359
	PerfParamContextSwitches = "context_switches"
	// PerfParamCPUMigrations as defined in libvirt/libvirt-domain.h:2369
	PerfParamCPUMigrations = "cpu_migrations"
	// PerfParamPageFaultsMin as defined in libvirt/libvirt-domain.h:2379
	PerfParamPageFaultsMin = "page_faults_min"
	// PerfParamPageFaultsMaj as defined in libvirt/libvirt-domain.h:2389
	PerfParamPageFaultsMaj = "page_faults_maj"
	// PerfParamAlignmentFaults as defined in libvirt/libvirt-domain.h:2399
	PerfParamAlignmentFaults = "alignment_faults"
	// PerfParamEmulationFaults as defined in libvirt/libvirt-domain.h:2409
	PerfParamEmulationFaults = "emulation_faults"
	// DomainBlockCopyBandwidth as defined in libvirt/libvirt-domain.h:2574
	DomainBlockCopyBandwidth = "bandwidth"
	// DomainBlockCopyGranularity as defined in libvirt/libvirt-domain.h:2585
	DomainBlockCopyGranularity = "granularity"
	// DomainBlockCopyBufSize as defined in libvirt/libvirt-domain.h:2594
	DomainBlockCopyBufSize = "buf-size"
	// DomainBlockIotuneTotalBytesSec as defined in libvirt/libvirt-domain.h:2635
	DomainBlockIotuneTotalBytesSec = "total_bytes_sec"
	// DomainBlockIotuneReadBytesSec as defined in libvirt/libvirt-domain.h:2643
	DomainBlockIotuneReadBytesSec = "read_bytes_sec"
	// DomainBlockIotuneWriteBytesSec as defined in libvirt/libvirt-domain.h:2651
	DomainBlockIotuneWriteBytesSec = "write_bytes_sec"
	// DomainBlockIotuneTotalIopsSec as defined in libvirt/libvirt-domain.h:2659
	DomainBlockIotuneTotalIopsSec = "total_iops_sec"
	// DomainBlockIotuneReadIopsSec as defined in libvirt/libvirt-domain.h:2667
	DomainBlockIotuneReadIopsSec = "read_iops_sec"
	// DomainBlockIotuneWriteIopsSec as defined in libvirt/libvirt-domain.h:2674
	DomainBlockIotuneWriteIopsSec = "write_iops_sec"
	// DomainBlockIotuneTotalBytesSecMax as defined in libvirt/libvirt-domain.h:2682
	DomainBlockIotuneTotalBytesSecMax = "total_bytes_sec_max"
	// DomainBlockIotuneReadBytesSecMax as defined in libvirt/libvirt-domain.h:2690
	DomainBlockIotuneReadBytesSecMax = "read_bytes_sec_max"
	// DomainBlockIotuneWriteBytesSecMax as defined in libvirt/libvirt-domain.h:2698
	DomainBlockIotuneWriteBytesSecMax = "write_bytes_sec_max"
	// DomainBlockIotuneTotalIopsSecMax as defined in libvirt/libvirt-domain.h:2706
	DomainBlockIotuneTotalIopsSecMax = "total_iops_sec_max"
	// DomainBlockIotuneReadIopsSecMax as defined in libvirt/libvirt-domain.h:2714
	DomainBlockIotuneReadIopsSecMax = "read_iops_sec_max"
	// DomainBlockIotuneWriteIopsSecMax as defined in libvirt/libvirt-domain.h:2721
	DomainBlockIotuneWriteIopsSecMax = "write_iops_sec_max"
	// DomainBlockIotuneTotalBytesSecMaxLength as defined in libvirt/libvirt-domain.h:2729
	DomainBlockIotuneTotalBytesSecMaxLength = "total_bytes_sec_max_length"
	// DomainBlockIotuneReadBytesSecMaxLength as defined in libvirt/libvirt-domain.h:2737
	DomainBlockIotuneReadBytesSecMaxLength = "read_bytes_sec_max_length"
	// DomainBlockIotuneWriteBytesSecMaxLength as defined in libvirt/libvirt-domain.h:2745
	DomainBlockIotuneWriteBytesSecMaxLength = "write_bytes_sec_max_length"
	// DomainBlockIotuneTotalIopsSecMaxLength as defined in libvirt/libvirt-domain.h:2753
	DomainBlockIotuneTotalIopsSecMaxLength = "total_iops_sec_max_length"
	// DomainBlockIotuneReadIopsSecMaxLength as defined in libvirt/libvirt-domain.h:2761
	DomainBlockIotuneReadIopsSecMaxLength = "read_iops_sec_max_length"
	// DomainBlockIotuneWriteIopsSecMaxLength as defined in libvirt/libvirt-domain.h:2769
	DomainBlockIotuneWriteIopsSecMaxLength = "write_iops_sec_max_length"
	// DomainBlockIotuneSizeIopsSec as defined in libvirt/libvirt-domain.h:2776
	DomainBlockIotuneSizeIopsSec = "size_iops_sec"
	// DomainBlockIotuneGroupName as defined in libvirt/libvirt-domain.h:2783
	DomainBlockIotuneGroupName = "group_name"
	// KeycodeSetRfb as defined in libvirt/libvirt-domain.h:2864
	KeycodeSetRfb = 0
	// DomainSendKeyMaxKeys as defined in libvirt/libvirt-domain.h:2871
	DomainSendKeyMaxKeys = 16
	// DomainJobOperationStr as defined in libvirt/libvirt-domain.h:3283
	DomainJobOperationStr = "operation"
	// DomainJobTimeElapsed as defined in libvirt/libvirt-domain.h:3293
	DomainJobTimeElapsed = "time_elapsed"
	// DomainJobTimeElapsedNet as defined in libvirt/libvirt-domain.h:3303
	DomainJobTimeElapsedNet = "time_elapsed_net"
	// DomainJobTimeRemaining as defined in libvirt/libvirt-domain.h:3313
	DomainJobTimeRemaining = "time_remaining"
	// DomainJobDowntime as defined in libvirt/libvirt-domain.h:3323
	DomainJobDowntime = "downtime"
	// DomainJobDowntimeNet as defined in libvirt/libvirt-domain.h:3332
	DomainJobDowntimeNet = "downtime_net"
	// DomainJobSetupTime as defined in libvirt/libvirt-domain.h:3341
	DomainJobSetupTime = "setup_time"
	// DomainJobDataTotal as defined in libvirt/libvirt-domain.h:3356
	DomainJobDataTotal = "data_total"
	// DomainJobDataProcessed as defined in libvirt/libvirt-domain.h:3366
	DomainJobDataProcessed = "data_processed"
	// DomainJobDataRemaining as defined in libvirt/libvirt-domain.h:3376
	DomainJobDataRemaining = "data_remaining"
	// DomainJobMemoryTotal as defined in libvirt/libvirt-domain.h:3386
	DomainJobMemoryTotal = "memory_total"
	// DomainJobMemoryProcessed as defined in libvirt/libvirt-domain.h:3396
	DomainJobMemoryProcessed = "memory_processed"
	// DomainJobMemoryRemaining as defined in libvirt/libvirt-domain.h:3406
	DomainJobMemoryRemaining = "memory_remaining"
	// DomainJobMemoryConstant as defined in libvirt/libvirt-domain.h:3418
	DomainJobMemoryConstant = "memory_constant"
	// DomainJobMemoryNormal as defined in libvirt/libvirt-domain.h:3428
	DomainJobMemoryNormal = "memory_normal"
	// DomainJobMemoryNormalBytes as defined in libvirt/libvirt-domain.h:3438
	DomainJobMemoryNormalBytes = "memory_normal_bytes"
	// DomainJobMemoryBps as defined in libvirt/libvirt-domain.h:3446
	DomainJobMemoryBps = "memory_bps"
	// DomainJobMemoryDirtyRate as defined in libvirt/libvirt-domain.h:3454
	DomainJobMemoryDirtyRate = "memory_dirty_rate"
	// DomainJobMemoryPageSize as defined in libvirt/libvirt-domain.h:3465
	DomainJobMemoryPageSize = "memory_page_size"
	// DomainJobMemoryIteration as defined in libvirt/libvirt-domain.h:3476
	DomainJobMemoryIteration = "memory_iteration"
	// DomainJobMemoryPostcopyReqs as defined in libvirt/libvirt-domain.h:3486
	DomainJobMemoryPostcopyReqs = "memory_postcopy_requests"
	// DomainJobDiskTotal as defined in libvirt/libvirt-domain.h:3496
	DomainJobDiskTotal = "disk_total"
	// DomainJobDiskProcessed as defined in libvirt/libvirt-domain.h:3506
	DomainJobDiskProcessed = "disk_processed"
	// DomainJobDiskRemaining as defined in libvirt/libvirt-domain.h:3516
	DomainJobDiskRemaining = "disk_remaining"
	// DomainJobDiskBps as defined in libvirt/libvirt-domain.h:3524
	DomainJobDiskBps = "disk_bps"
	// DomainJobCompressionCache as defined in libvirt/libvirt-domain.h:3533
	DomainJobCompressionCache = "compression_cache"
	// DomainJobCompressionBytes as defined in libvirt/libvirt-domain.h:3541
	DomainJobCompressionBytes = "compression_bytes"
	// DomainJobCompressionPages as defined in libvirt/libvirt-domain.h:3549
	DomainJobCompressionPages = "compression_pages"
	// DomainJobCompressionCacheMisses as defined in libvirt/libvirt-domain.h:3558
	DomainJobCompressionCacheMisses = "compression_cache_misses"
	// DomainJobCompressionOverflow as defined in libvirt/libvirt-domain.h:3568
	DomainJobCompressionOverflow = "compression_overflow"
	// DomainJobAutoConvergeThrottle as defined in libvirt/libvirt-domain.h:3577
	DomainJobAutoConvergeThrottle = "auto_converge_throttle"
	// DomainTunableCPUVcpupin as defined in libvirt/libvirt-domain.h:4131
	DomainTunableCPUVcpupin = "cputune.vcpupin%u"
	// DomainTunableCPUEmulatorpin as defined in libvirt/libvirt-domain.h:4139
	DomainTunableCPUEmulatorpin = "cputune.emulatorpin"
	// DomainTunableCPUIothreadspin as defined in libvirt/libvirt-domain.h:4148
	DomainTunableCPUIothreadspin = "cputune.iothreadpin%u"
	// DomainTunableCPUCpuShares as defined in libvirt/libvirt-domain.h:4156
	DomainTunableCPUCpuShares = "cputune.cpu_shares"
	// DomainTunableCPUGlobalPeriod as defined in libvirt/libvirt-domain.h:4164
	DomainTunableCPUGlobalPeriod = "cputune.global_period"
	// DomainTunableCPUGlobalQuota as defined in libvirt/libvirt-domain.h:4172
	DomainTunableCPUGlobalQuota = "cputune.global_quota"
	// DomainTunableCPUVCPUPeriod as defined in libvirt/libvirt-domain.h:4180
	DomainTunableCPUVCPUPeriod = "cputune.vcpu_period"
	// DomainTunableCPUVCPUQuota as defined in libvirt/libvirt-domain.h:4188
	DomainTunableCPUVCPUQuota = "cputune.vcpu_quota"
	// DomainTunableCPUEmulatorPeriod as defined in libvirt/libvirt-domain.h:4197
	DomainTunableCPUEmulatorPeriod = "cputune.emulator_period"
	// DomainTunableCPUEmulatorQuota as defined in libvirt/libvirt-domain.h:4206
	DomainTunableCPUEmulatorQuota = "cputune.emulator_quota"
	// DomainTunableCPUIothreadPeriod as defined in libvirt/libvirt-domain.h:4214
	DomainTunableCPUIothreadPeriod = "cputune.iothread_period"
	// DomainTunableCPUIothreadQuota as defined in libvirt/libvirt-domain.h:4222
	DomainTunableCPUIothreadQuota = "cputune.iothread_quota"
	// DomainTunableBlkdevDisk as defined in libvirt/libvirt-domain.h:4230
	DomainTunableBlkdevDisk = "blkdeviotune.disk"
	// DomainTunableBlkdevTotalBytesSec as defined in libvirt/libvirt-domain.h:4238
	DomainTunableBlkdevTotalBytesSec = "blkdeviotune.total_bytes_sec"
	// DomainTunableBlkdevReadBytesSec as defined in libvirt/libvirt-domain.h:4246
	DomainTunableBlkdevReadBytesSec = "blkdeviotune.read_bytes_sec"
	// DomainTunableBlkdevWriteBytesSec as defined in libvirt/libvirt-domain.h:4254
	DomainTunableBlkdevWriteBytesSec = "blkdeviotune.write_bytes_sec"
	// DomainTunableBlkdevTotalIopsSec as defined in libvirt/libvirt-domain.h:4262
	DomainTunableBlkdevTotalIopsSec = "blkdeviotune.total_iops_sec"
	// DomainTunableBlkdevReadIopsSec as defined in libvirt/libvirt-domain.h:4270
	DomainTunableBlkdevReadIopsSec = "blkdeviotune.read_iops_sec"
	// DomainTunableBlkdevWriteIopsSec as defined in libvirt/libvirt-domain.h:4278
	DomainTunableBlkdevWriteIopsSec = "blkdeviotune.write_iops_sec"
	// DomainTunableBlkdevTotalBytesSecMax as defined in libvirt/libvirt-domain.h:4286
	DomainTunableBlkdevTotalBytesSecMax = "blkdeviotune.total_bytes_sec_max"
	// DomainTunableBlkdevReadBytesSecMax as defined in libvirt/libvirt-domain.h:4294
	DomainTunableBlkdevReadBytesSecMax = "blkdeviotune.read_bytes_sec_max"
	// DomainTunableBlkdevWriteBytesSecMax as defined in libvirt/libvirt-domain.h:4302
	DomainTunableBlkdevWriteBytesSecMax = "blkdeviotune.write_bytes_sec_max"
	// DomainTunableBlkdevTotalIopsSecMax as defined in libvirt/libvirt-domain.h:4310
	DomainTunableBlkdevTotalIopsSecMax = "blkdeviotune.total_iops_sec_max"
	// DomainTunableBlkdevReadIopsSecMax as defined in libvirt/libvirt-domain.h:4318
	DomainTunableBlkdevReadIopsSecMax = "blkdeviotune.read_iops_sec_max"
	// DomainTunableBlkdevWriteIopsSecMax as defined in libvirt/libvirt-domain.h:4326
	DomainTunableBlkdevWriteIopsSecMax = "blkdeviotune.write_iops_sec_max"
	// DomainTunableBlkdevSizeIopsSec as defined in libvirt/libvirt-domain.h:4334
	DomainTunableBlkdevSizeIopsSec = "blkdeviotune.size_iops_sec"
	// DomainTunableBlkdevGroupName as defined in libvirt/libvirt-domain.h:4342
	DomainTunableBlkdevGroupName = "blkdeviotune.group_name"
	// DomainTunableBlkdevTotalBytesSecMaxLength as defined in libvirt/libvirt-domain.h:4351
	DomainTunableBlkdevTotalBytesSecMaxLength = "blkdeviotune.total_bytes_sec_max_length"
	// DomainTunableBlkdevReadBytesSecMaxLength as defined in libvirt/libvirt-domain.h:4360
	DomainTunableBlkdevReadBytesSecMaxLength = "blkdeviotune.read_bytes_sec_max_length"
	// DomainTunableBlkdevWriteBytesSecMaxLength as defined in libvirt/libvirt-domain.h:4369
	DomainTunableBlkdevWriteBytesSecMaxLength = "blkdeviotune.write_bytes_sec_max_length"
	// DomainTunableBlkdevTotalIopsSecMaxLength as defined in libvirt/libvirt-domain.h:4378
	DomainTunableBlkdevTotalIopsSecMaxLength = "blkdeviotune.total_iops_sec_max_length"
	// DomainTunableBlkdevReadIopsSecMaxLength as defined in libvirt/libvirt-domain.h:4387
	DomainTunableBlkdevReadIopsSecMaxLength = "blkdeviotune.read_iops_sec_max_length"
	// DomainTunableBlkdevWriteIopsSecMaxLength as defined in libvirt/libvirt-domain.h:4396
	DomainTunableBlkdevWriteIopsSecMaxLength = "blkdeviotune.write_iops_sec_max_length"
	// DomainSchedFieldLength as defined in libvirt/libvirt-domain.h:4684
	DomainSchedFieldLength = 80
	// DomainBlkioFieldLength as defined in libvirt/libvirt-domain.h:4728
	DomainBlkioFieldLength = 80
	// DomainMemoryFieldLength as defined in libvirt/libvirt-domain.h:4772
	DomainMemoryFieldLength = 80
	// DomainLaunchSecuritySevMeasurement as defined in libvirt/libvirt-domain.h:4898
	DomainLaunchSecuritySevMeasurement = "sev-measurement"
	// NetworkPortBandwidthInAverage as defined in libvirt/libvirt-network.h:391
	NetworkPortBandwidthInAverage = "inbound.average"
	// NetworkPortBandwidthInPeak as defined in libvirt/libvirt-network.h:398
	NetworkPortBandwidthInPeak = "inbound.peak"
	// NetworkPortBandwidthInBurst as defined in libvirt/libvirt-network.h:405
	NetworkPortBandwidthInBurst = "inbound.burst"
	// NetworkPortBandwidthInFloor as defined in libvirt/libvirt-network.h:412
	NetworkPortBandwidthInFloor = "inbound.floor"
	// NetworkPortBandwidthOutAverage as defined in libvirt/libvirt-network.h:419
	NetworkPortBandwidthOutAverage = "outbound.average"
	// NetworkPortBandwidthOutPeak as defined in libvirt/libvirt-network.h:426
	NetworkPortBandwidthOutPeak = "outbound.peak"
	// NetworkPortBandwidthOutBurst as defined in libvirt/libvirt-network.h:433
	NetworkPortBandwidthOutBurst = "outbound.burst"
)

// ConnectCloseReason as declared in libvirt/libvirt-common.h:119
type ConnectCloseReason int32

// ConnectCloseReason enumeration from libvirt/libvirt-common.h:119
const (
	ConnectCloseReasonError     ConnectCloseReason = iota
	ConnectCloseReasonEOF       ConnectCloseReason = 1
	ConnectCloseReasonKeepalive ConnectCloseReason = 2
	ConnectCloseReasonClient    ConnectCloseReason = 3
)

// TypedParameterType as declared in libvirt/libvirt-common.h:138
type TypedParameterType int32

// TypedParameterType enumeration from libvirt/libvirt-common.h:138
const (
	TypedParamInt     TypedParameterType = 1
	TypedParamUint    TypedParameterType = 2
	TypedParamLlong   TypedParameterType = 3
	TypedParamUllong  TypedParameterType = 4
	TypedParamDouble  TypedParameterType = 5
	TypedParamBoolean TypedParameterType = 6
	TypedParamString  TypedParameterType = 7
)

// TypedParameterFlags as declared in libvirt/libvirt-common.h:163
type TypedParameterFlags int32

// TypedParameterFlags enumeration from libvirt/libvirt-common.h:163
const (
	TypedParamStringOkay TypedParameterFlags = 4
)

// NodeSuspendTarget as declared in libvirt/libvirt-host.h:61
type NodeSuspendTarget int32

// NodeSuspendTarget enumeration from libvirt/libvirt-host.h:61
const (
	NodeSuspendTargetMem    NodeSuspendTarget = iota
	NodeSuspendTargetDisk   NodeSuspendTarget = 1
	NodeSuspendTargetHybrid NodeSuspendTarget = 2
)

// NodeGetCPUStatsAllCPUs as declared in libvirt/libvirt-host.h:185
type NodeGetCPUStatsAllCPUs int32

// NodeGetCPUStatsAllCPUs enumeration from libvirt/libvirt-host.h:185
const (
	NodeCPUStatsAllCpus NodeGetCPUStatsAllCPUs = -1
)

// NodeGetMemoryStatsAllCells as declared in libvirt/libvirt-host.h:263
type NodeGetMemoryStatsAllCells int32

// NodeGetMemoryStatsAllCells enumeration from libvirt/libvirt-host.h:263
const (
	NodeMemoryStatsAllCells NodeGetMemoryStatsAllCells = -1
)

// ConnectFlags as declared in libvirt/libvirt-host.h:484
type ConnectFlags int32

// ConnectFlags enumeration from libvirt/libvirt-host.h:484
const (
	ConnectRo        ConnectFlags = 1
	ConnectNoAliases ConnectFlags = 2
)

// ConnectCredentialType as declared in libvirt/libvirt-host.h:501
type ConnectCredentialType int32

// ConnectCredentialType enumeration from libvirt/libvirt-host.h:501
const (
	CredUsername     ConnectCredentialType = 1
	CredAuthname     ConnectCredentialType = 2
	CredLanguage     ConnectCredentialType = 3
	CredCnonce       ConnectCredentialType = 4
	CredPassphrase   ConnectCredentialType = 5
	CredEchoprompt   ConnectCredentialType = 6
	CredNoechoprompt ConnectCredentialType = 7
	CredRealm        ConnectCredentialType = 8
	CredExternal     ConnectCredentialType = 9
)

// CPUCompareResult as declared in libvirt/libvirt-host.h:748
type CPUCompareResult int32

// CPUCompareResult enumeration from libvirt/libvirt-host.h:748
const (
	CPUCompareError        CPUCompareResult = -1
	CPUCompareIncompatible CPUCompareResult = 0
	CPUCompareIdentical    CPUCompareResult = 1
	CPUCompareSuperset     CPUCompareResult = 2
)

// ConnectCompareCPUFlags as declared in libvirt/libvirt-host.h:753
type ConnectCompareCPUFlags int32

// ConnectCompareCPUFlags enumeration from libvirt/libvirt-host.h:753
const (
	ConnectCompareCPUFailIncompatible ConnectCompareCPUFlags = 1
)

// ConnectBaselineCPUFlags as declared in libvirt/libvirt-host.h:779
type ConnectBaselineCPUFlags int32

// ConnectBaselineCPUFlags enumeration from libvirt/libvirt-host.h:779
const (
	ConnectBaselineCPUExpandFeatures ConnectBaselineCPUFlags = 1
	ConnectBaselineCPUMigratable     ConnectBaselineCPUFlags = 2
)

// NodeAllocPagesFlags as declared in libvirt/libvirt-host.h:809
type NodeAllocPagesFlags int32

// NodeAllocPagesFlags enumeration from libvirt/libvirt-host.h:809
const (
	NodeAllocPagesAdd NodeAllocPagesFlags = iota
	NodeAllocPagesSet NodeAllocPagesFlags = 1
)

// DomainState as declared in libvirt/libvirt-domain.h:70
type DomainState int32

// DomainState enumeration from libvirt/libvirt-domain.h:70
const (
	DomainNostate     DomainState = iota
	DomainRunning     DomainState = 1
	DomainBlocked     DomainState = 2
	DomainPaused      DomainState = 3
	DomainShutdown    DomainState = 4
	DomainShutoff     DomainState = 5
	DomainCrashed     DomainState = 6
	DomainPmsuspended DomainState = 7
)

// DomainNostateReason as declared in libvirt/libvirt-domain.h:78
type DomainNostateReason int32

// DomainNostateReason enumeration from libvirt/libvirt-domain.h:78
const (
	DomainNostateUnknown DomainNostateReason = iota
)

// DomainRunningReason as declared in libvirt/libvirt-domain.h:97
type DomainRunningReason int32

// DomainRunningReason enumeration from libvirt/libvirt-domain.h:97
const (
	DomainRunningUnknown           DomainRunningReason = iota
	DomainRunningBooted            DomainRunningReason = 1
	DomainRunningMigrated          DomainRunningReason = 2
	DomainRunningRestored          DomainRunningReason = 3
	DomainRunningFromSnapshot      DomainRunningReason = 4
	DomainRunningUnpaused          DomainRunningReason = 5
	DomainRunningMigrationCanceled DomainRunningReason = 6
	DomainRunningSaveCanceled      DomainRunningReason = 7
	DomainRunningWakeup            DomainRunningReason = 8
	DomainRunningCrashed           DomainRunningReason = 9
	DomainRunningPostcopy          DomainRunningReason = 10
)

// DomainBlockedReason as declared in libvirt/libvirt-domain.h:105
type DomainBlockedReason int32

// DomainBlockedReason enumeration from libvirt/libvirt-domain.h:105
const (
	DomainBlockedUnknown DomainBlockedReason = iota
)

// DomainPausedReason as declared in libvirt/libvirt-domain.h:126
type DomainPausedReason int32

// DomainPausedReason enumeration from libvirt/libvirt-domain.h:126
const (
	DomainPausedUnknown        DomainPausedReason = iota
	DomainPausedUser           DomainPausedReason = 1
	DomainPausedMigration      DomainPausedReason = 2
	DomainPausedSave           DomainPausedReason = 3
	DomainPausedDump           DomainPausedReason = 4
	DomainPausedIoerror        DomainPausedReason = 5
	DomainPausedWatchdog       DomainPausedReason = 6
	DomainPausedFromSnapshot   DomainPausedReason = 7
	DomainPausedShuttingDown   DomainPausedReason = 8
	DomainPausedSnapshot       DomainPausedReason = 9
	DomainPausedCrashed        DomainPausedReason = 10
	DomainPausedStartingUp     DomainPausedReason = 11
	DomainPausedPostcopy       DomainPausedReason = 12
	DomainPausedPostcopyFailed DomainPausedReason = 13
)

// DomainShutdownReason as declared in libvirt/libvirt-domain.h:135
type DomainShutdownReason int32

// DomainShutdownReason enumeration from libvirt/libvirt-domain.h:135
const (
	DomainShutdownUnknown DomainShutdownReason = iota
	DomainShutdownUser    DomainShutdownReason = 1
)

// DomainShutoffReason as declared in libvirt/libvirt-domain.h:152
type DomainShutoffReason int32

// DomainShutoffReason enumeration from libvirt/libvirt-domain.h:152
const (
	DomainShutoffUnknown      DomainShutoffReason = iota
	DomainShutoffShutdown     DomainShutoffReason = 1
	DomainShutoffDestroyed    DomainShutoffReason = 2
	DomainShutoffCrashed      DomainShutoffReason = 3
	DomainShutoffMigrated     DomainShutoffReason = 4
	DomainShutoffSaved        DomainShutoffReason = 5
	DomainShutoffFailed       DomainShutoffReason = 6
	DomainShutoffFromSnapshot DomainShutoffReason = 7
	DomainShutoffDaemon       DomainShutoffReason = 8
)

// DomainCrashedReason as declared in libvirt/libvirt-domain.h:161
type DomainCrashedReason int32

// DomainCrashedReason enumeration from libvirt/libvirt-domain.h:161
const (
	DomainCrashedUnknown  DomainCrashedReason = iota
	DomainCrashedPanicked DomainCrashedReason = 1
)

// DomainPMSuspendedReason as declared in libvirt/libvirt-domain.h:169
type DomainPMSuspendedReason int32

// DomainPMSuspendedReason enumeration from libvirt/libvirt-domain.h:169
const (
	DomainPmsuspendedUnknown DomainPMSuspendedReason = iota
)

// DomainPMSuspendedDiskReason as declared in libvirt/libvirt-domain.h:177
type DomainPMSuspendedDiskReason int32

// DomainPMSuspendedDiskReason enumeration from libvirt/libvirt-domain.h:177
const (
	DomainPmsuspendedDiskUnknown DomainPMSuspendedDiskReason = iota
)

// DomainControlState as declared in libvirt/libvirt-domain.h:197
type DomainControlState int32

// DomainControlState enumeration from libvirt/libvirt-domain.h:197
const (
	DomainControlOk       DomainControlState = iota
	DomainControlJob      DomainControlState = 1
	DomainControlOccupied DomainControlState = 2
	DomainControlError    DomainControlState = 3
)

// DomainControlErrorReason as declared in libvirt/libvirt-domain.h:217
type DomainControlErrorReason int32

// DomainControlErrorReason enumeration from libvirt/libvirt-domain.h:217
const (
	DomainControlErrorReasonNone     DomainControlErrorReason = iota
	DomainControlErrorReasonUnknown  DomainControlErrorReason = 1
	DomainControlErrorReasonMonitor  DomainControlErrorReason = 2
	DomainControlErrorReasonInternal DomainControlErrorReason = 3
)

// DomainModificationImpact as declared in libvirt/libvirt-domain.h:265
type DomainModificationImpact int32

// DomainModificationImpact enumeration from libvirt/libvirt-domain.h:265
const (
	DomainAffectCurrent DomainModificationImpact = iota
	DomainAffectLive    DomainModificationImpact = 1
	DomainAffectConfig  DomainModificationImpact = 2
)

// DomainCreateFlags as declared in libvirt/libvirt-domain.h:305
type DomainCreateFlags int32

// DomainCreateFlags enumeration from libvirt/libvirt-domain.h:305
const (
	DomainNone             DomainCreateFlags = iota
	DomainStartPaused      DomainCreateFlags = 1
	DomainStartAutodestroy DomainCreateFlags = 2
	DomainStartBypassCache DomainCreateFlags = 4
	DomainStartForceBoot   DomainCreateFlags = 8
	DomainStartValidate    DomainCreateFlags = 16
)

// DomainMemoryStatTags as declared in libvirt/libvirt-domain.h:660
type DomainMemoryStatTags int32

// DomainMemoryStatTags enumeration from libvirt/libvirt-domain.h:660
const (
	DomainMemoryStatSwapIn         DomainMemoryStatTags = iota
	DomainMemoryStatSwapOut        DomainMemoryStatTags = 1
	DomainMemoryStatMajorFault     DomainMemoryStatTags = 2
	DomainMemoryStatMinorFault     DomainMemoryStatTags = 3
	DomainMemoryStatUnused         DomainMemoryStatTags = 4
	DomainMemoryStatAvailable      DomainMemoryStatTags = 5
	DomainMemoryStatActualBalloon  DomainMemoryStatTags = 6
	DomainMemoryStatRss            DomainMemoryStatTags = 7
	DomainMemoryStatUsable         DomainMemoryStatTags = 8
	DomainMemoryStatLastUpdate     DomainMemoryStatTags = 9
	DomainMemoryStatDiskCaches     DomainMemoryStatTags = 10
	DomainMemoryStatHugetlbPgalloc DomainMemoryStatTags = 11
	DomainMemoryStatHugetlbPgfail  DomainMemoryStatTags = 12
	DomainMemoryStatNr             DomainMemoryStatTags = 13
)

// DomainCoreDumpFlags as declared in libvirt/libvirt-domain.h:679
type DomainCoreDumpFlags int32

// DomainCoreDumpFlags enumeration from libvirt/libvirt-domain.h:679
const (
	DumpCrash       DomainCoreDumpFlags = 1
	DumpLive        DomainCoreDumpFlags = 2
	DumpBypassCache DomainCoreDumpFlags = 4
	DumpReset       DomainCoreDumpFlags = 8
	DumpMemoryOnly  DomainCoreDumpFlags = 16
)

// DomainCoreDumpFormat as declared in libvirt/libvirt-domain.h:702
type DomainCoreDumpFormat int32

// DomainCoreDumpFormat enumeration from libvirt/libvirt-domain.h:702
const (
	DomainCoreDumpFormatRaw         DomainCoreDumpFormat = iota
	DomainCoreDumpFormatKdumpZlib   DomainCoreDumpFormat = 1
	DomainCoreDumpFormatKdumpLzo    DomainCoreDumpFormat = 2
	DomainCoreDumpFormatKdumpSnappy DomainCoreDumpFormat = 3
)

// DomainMigrateFlags as declared in libvirt/libvirt-domain.h:852
type DomainMigrateFlags int32

// DomainMigrateFlags enumeration from libvirt/libvirt-domain.h:852
const (
	MigrateLive             DomainMigrateFlags = 1
	MigratePeer2peer        DomainMigrateFlags = 2
	MigrateTunnelled        DomainMigrateFlags = 4
	MigratePersistDest      DomainMigrateFlags = 8
	MigrateUndefineSource   DomainMigrateFlags = 16
	MigratePaused           DomainMigrateFlags = 32
	MigrateNonSharedDisk    DomainMigrateFlags = 64
	MigrateNonSharedInc     DomainMigrateFlags = 128
	MigrateChangeProtection DomainMigrateFlags = 256
	MigrateUnsafe           DomainMigrateFlags = 512
	MigrateOffline          DomainMigrateFlags = 1024
	MigrateCompressed       DomainMigrateFlags = 2048
	MigrateAbortOnError     DomainMigrateFlags = 4096
	MigrateAutoConverge     DomainMigrateFlags = 8192
	MigrateRdmaPinAll       DomainMigrateFlags = 16384
	MigratePostcopy         DomainMigrateFlags = 32768
	MigrateTLS              DomainMigrateFlags = 65536
	MigrateParallel         DomainMigrateFlags = 131072
)

// DomainMigrateMaxSpeedFlags as declared in libvirt/libvirt-domain.h:1104
type DomainMigrateMaxSpeedFlags int32

// DomainMigrateMaxSpeedFlags enumeration from libvirt/libvirt-domain.h:1104
const (
	DomainMigrateMaxSpeedPostcopy DomainMigrateMaxSpeedFlags = 1
)

// DomainShutdownFlagValues as declared in libvirt/libvirt-domain.h:1170
type DomainShutdownFlagValues int32

// DomainShutdownFlagValues enumeration from libvirt/libvirt-domain.h:1170
const (
	DomainShutdownDefault      DomainShutdownFlagValues = iota
	DomainShutdownAcpiPowerBtn DomainShutdownFlagValues = 1
	DomainShutdownGuestAgent   DomainShutdownFlagValues = 2
	DomainShutdownInitctl      DomainShutdownFlagValues = 4
	DomainShutdownSignal       DomainShutdownFlagValues = 8
	DomainShutdownParavirt     DomainShutdownFlagValues = 16
)

// DomainRebootFlagValues as declared in libvirt/libvirt-domain.h:1183
type DomainRebootFlagValues int32

// DomainRebootFlagValues enumeration from libvirt/libvirt-domain.h:1183
const (
	DomainRebootDefault      DomainRebootFlagValues = iota
	DomainRebootAcpiPowerBtn DomainRebootFlagValues = 1
	DomainRebootGuestAgent   DomainRebootFlagValues = 2
	DomainRebootInitctl      DomainRebootFlagValues = 4
	DomainRebootSignal       DomainRebootFlagValues = 8
	DomainRebootParavirt     DomainRebootFlagValues = 16
)

// DomainDestroyFlagsValues as declared in libvirt/libvirt-domain.h:1201
type DomainDestroyFlagsValues int32

// DomainDestroyFlagsValues enumeration from libvirt/libvirt-domain.h:1201
const (
	DomainDestroyDefault  DomainDestroyFlagsValues = iota
	DomainDestroyGraceful DomainDestroyFlagsValues = 1
)

// DomainSaveRestoreFlags as declared in libvirt/libvirt-domain.h:1233
type DomainSaveRestoreFlags int32

// DomainSaveRestoreFlags enumeration from libvirt/libvirt-domain.h:1233
const (
	DomainSaveBypassCache DomainSaveRestoreFlags = 1
	DomainSaveRunning     DomainSaveRestoreFlags = 2
	DomainSavePaused      DomainSaveRestoreFlags = 4
)

// DomainMemoryModFlags as declared in libvirt/libvirt-domain.h:1489
type DomainMemoryModFlags int32

// DomainMemoryModFlags enumeration from libvirt/libvirt-domain.h:1489
const (
	DomainMemCurrent DomainMemoryModFlags = iota
	DomainMemLive    DomainMemoryModFlags = 1
	DomainMemConfig  DomainMemoryModFlags = 2
	DomainMemMaximum DomainMemoryModFlags = 4
)

// DomainNumatuneMemMode as declared in libvirt/libvirt-domain.h:1507
type DomainNumatuneMemMode int32

// DomainNumatuneMemMode enumeration from libvirt/libvirt-domain.h:1507
const (
	DomainNumatuneMemStrict     DomainNumatuneMemMode = iota
	DomainNumatuneMemPreferred  DomainNumatuneMemMode = 1
	DomainNumatuneMemInterleave DomainNumatuneMemMode = 2
)

// DomainMetadataType as declared in libvirt/libvirt-domain.h:1569
type DomainMetadataType int32

// DomainMetadataType enumeration from libvirt/libvirt-domain.h:1569
const (
	DomainMetadataDescription DomainMetadataType = iota
	DomainMetadataTitle       DomainMetadataType = 1
	DomainMetadataElement     DomainMetadataType = 2
)

// DomainXMLFlags as declared in libvirt/libvirt-domain.h:1599
type DomainXMLFlags int32

// DomainXMLFlags enumeration from libvirt/libvirt-domain.h:1599
const (
	DomainXMLSecure     DomainXMLFlags = 1
	DomainXMLInactive   DomainXMLFlags = 2
	DomainXMLUpdateCPU  DomainXMLFlags = 4
	DomainXMLMigratable DomainXMLFlags = 8
)

// DomainSaveImageXMLFlags as declared in libvirt/libvirt-domain.h:1603
type DomainSaveImageXMLFlags int32

// DomainSaveImageXMLFlags enumeration from libvirt/libvirt-domain.h:1603
const (
	DomainSaveImageXMLSecure DomainSaveImageXMLFlags = 1
)

// DomainBlockResizeFlags as declared in libvirt/libvirt-domain.h:1708
type DomainBlockResizeFlags int32

// DomainBlockResizeFlags enumeration from libvirt/libvirt-domain.h:1708
const (
	DomainBlockResizeBytes DomainBlockResizeFlags = 1
)

// DomainMemoryFlags as declared in libvirt/libvirt-domain.h:1771
type DomainMemoryFlags int32

// DomainMemoryFlags enumeration from libvirt/libvirt-domain.h:1771
const (
	MemoryVirtual  DomainMemoryFlags = 1
	MemoryPhysical DomainMemoryFlags = 2
)

// DomainDefineFlags as declared in libvirt/libvirt-domain.h:1781
type DomainDefineFlags int32

// DomainDefineFlags enumeration from libvirt/libvirt-domain.h:1781
const (
	DomainDefineValidate DomainDefineFlags = 1
)

// DomainUndefineFlagsValues as declared in libvirt/libvirt-domain.h:1808
type DomainUndefineFlagsValues int32

// DomainUndefineFlagsValues enumeration from libvirt/libvirt-domain.h:1808
const (
	DomainUndefineManagedSave         DomainUndefineFlagsValues = 1
	DomainUndefineSnapshotsMetadata   DomainUndefineFlagsValues = 2
	DomainUndefineNvram               DomainUndefineFlagsValues = 4
	DomainUndefineKeepNvram           DomainUndefineFlagsValues = 8
	DomainUndefineCheckpointsMetadata DomainUndefineFlagsValues = 16
)

// ConnectListAllDomainsFlags as declared in libvirt/libvirt-domain.h:1847
type ConnectListAllDomainsFlags int32

// ConnectListAllDomainsFlags enumeration from libvirt/libvirt-domain.h:1847
const (
	ConnectListDomainsActive        ConnectListAllDomainsFlags = 1
	ConnectListDomainsInactive      ConnectListAllDomainsFlags = 2
	ConnectListDomainsPersistent    ConnectListAllDomainsFlags = 4
	ConnectListDomainsTransient     ConnectListAllDomainsFlags = 8
	ConnectListDomainsRunning       ConnectListAllDomainsFlags = 16
	ConnectListDomainsPaused        ConnectListAllDomainsFlags = 32
	ConnectListDomainsShutoff       ConnectListAllDomainsFlags = 64
	ConnectListDomainsOther         ConnectListAllDomainsFlags = 128
	ConnectListDomainsManagedsave   ConnectListAllDomainsFlags = 256
	ConnectListDomainsNoManagedsave ConnectListAllDomainsFlags = 512
	ConnectListDomainsAutostart     ConnectListAllDomainsFlags = 1024
	ConnectListDomainsNoAutostart   ConnectListAllDomainsFlags = 2048
	ConnectListDomainsHasSnapshot   ConnectListAllDomainsFlags = 4096
	ConnectListDomainsNoSnapshot    ConnectListAllDomainsFlags = 8192
	ConnectListDomainsHasCheckpoint ConnectListAllDomainsFlags = 16384
	ConnectListDomainsNoCheckpoint  ConnectListAllDomainsFlags = 32768
)

// VCPUState as declared in libvirt/libvirt-domain.h:1878
type VCPUState int32

// VCPUState enumeration from libvirt/libvirt-domain.h:1878
const (
	VCPUOffline VCPUState = iota
	VCPURunning VCPUState = 1
	VCPUBlocked VCPUState = 2
)

// DomainVCPUFlags as declared in libvirt/libvirt-domain.h:1900
type DomainVCPUFlags int32

// DomainVCPUFlags enumeration from libvirt/libvirt-domain.h:1900
const (
	DomainVCPUCurrent      DomainVCPUFlags = iota
	DomainVCPULive         DomainVCPUFlags = 1
	DomainVCPUConfig       DomainVCPUFlags = 2
	DomainVCPUMaximum      DomainVCPUFlags = 4
	DomainVCPUGuest        DomainVCPUFlags = 8
	DomainVCPUHotpluggable DomainVCPUFlags = 16
)

// DomainDeviceModifyFlags as declared in libvirt/libvirt-domain.h:2117
type DomainDeviceModifyFlags int32

// DomainDeviceModifyFlags enumeration from libvirt/libvirt-domain.h:2117
const (
	DomainDeviceModifyCurrent DomainDeviceModifyFlags = iota
	DomainDeviceModifyLive    DomainDeviceModifyFlags = 1
	DomainDeviceModifyConfig  DomainDeviceModifyFlags = 2
	DomainDeviceModifyForce   DomainDeviceModifyFlags = 4
)

// DomainStatsTypes as declared in libvirt/libvirt-domain.h:2149
type DomainStatsTypes int32

// DomainStatsTypes enumeration from libvirt/libvirt-domain.h:2149
const (
	DomainStatsState     DomainStatsTypes = 1
	DomainStatsCPUTotal  DomainStatsTypes = 2
	DomainStatsBalloon   DomainStatsTypes = 4
	DomainStatsVCPU      DomainStatsTypes = 8
	DomainStatsInterface DomainStatsTypes = 16
	DomainStatsBlock     DomainStatsTypes = 32
	DomainStatsPerf      DomainStatsTypes = 64
	DomainStatsIothread  DomainStatsTypes = 128
)

// ConnectGetAllDomainStatsFlags as declared in libvirt/libvirt-domain.h:2167
type ConnectGetAllDomainStatsFlags int32

// ConnectGetAllDomainStatsFlags enumeration from libvirt/libvirt-domain.h:2167
const (
	ConnectGetAllDomainsStatsActive       ConnectGetAllDomainStatsFlags = 1
	ConnectGetAllDomainsStatsInactive     ConnectGetAllDomainStatsFlags = 2
	ConnectGetAllDomainsStatsPersistent   ConnectGetAllDomainStatsFlags = 4
	ConnectGetAllDomainsStatsTransient    ConnectGetAllDomainStatsFlags = 8
	ConnectGetAllDomainsStatsRunning      ConnectGetAllDomainStatsFlags = 16
	ConnectGetAllDomainsStatsPaused       ConnectGetAllDomainStatsFlags = 32
	ConnectGetAllDomainsStatsShutoff      ConnectGetAllDomainStatsFlags = 64
	ConnectGetAllDomainsStatsOther        ConnectGetAllDomainStatsFlags = 128
	ConnectGetAllDomainsStatsNowait       ConnectGetAllDomainStatsFlags = 536870912
	ConnectGetAllDomainsStatsBacking      ConnectGetAllDomainStatsFlags = 1073741824
	ConnectGetAllDomainsStatsEnforceStats ConnectGetAllDomainStatsFlags = -2147483648
)

// DomainBlockJobType as declared in libvirt/libvirt-domain.h:2452
type DomainBlockJobType int32

// DomainBlockJobType enumeration from libvirt/libvirt-domain.h:2452
const (
	DomainBlockJobTypeUnknown      DomainBlockJobType = iota
	DomainBlockJobTypePull         DomainBlockJobType = 1
	DomainBlockJobTypeCopy         DomainBlockJobType = 2
	DomainBlockJobTypeCommit       DomainBlockJobType = 3
	DomainBlockJobTypeActiveCommit DomainBlockJobType = 4
)

// DomainBlockJobAbortFlags as declared in libvirt/libvirt-domain.h:2464
type DomainBlockJobAbortFlags int32

// DomainBlockJobAbortFlags enumeration from libvirt/libvirt-domain.h:2464
const (
	DomainBlockJobAbortAsync DomainBlockJobAbortFlags = 1
	DomainBlockJobAbortPivot DomainBlockJobAbortFlags = 2
)

// DomainBlockJobInfoFlags as declared in libvirt/libvirt-domain.h:2473
type DomainBlockJobInfoFlags int32

// DomainBlockJobInfoFlags enumeration from libvirt/libvirt-domain.h:2473
const (
	DomainBlockJobInfoBandwidthBytes DomainBlockJobInfoFlags = 1
)

// DomainBlockJobSetSpeedFlags as declared in libvirt/libvirt-domain.h:2502
type DomainBlockJobSetSpeedFlags int32

// DomainBlockJobSetSpeedFlags enumeration from libvirt/libvirt-domain.h:2502
const (
	DomainBlockJobSpeedBandwidthBytes DomainBlockJobSetSpeedFlags = 1
)

// DomainBlockPullFlags as declared in libvirt/libvirt-domain.h:2512
type DomainBlockPullFlags int32

// DomainBlockPullFlags enumeration from libvirt/libvirt-domain.h:2512
const (
	DomainBlockPullBandwidthBytes DomainBlockPullFlags = 64
)

// DomainBlockRebaseFlags as declared in libvirt/libvirt-domain.h:2536
type DomainBlockRebaseFlags int32

// DomainBlockRebaseFlags enumeration from libvirt/libvirt-domain.h:2536
const (
	DomainBlockRebaseShallow        DomainBlockRebaseFlags = 1
	DomainBlockRebaseReuseExt       DomainBlockRebaseFlags = 2
	DomainBlockRebaseCopyRaw        DomainBlockRebaseFlags = 4
	DomainBlockRebaseCopy           DomainBlockRebaseFlags = 8
	DomainBlockRebaseRelative       DomainBlockRebaseFlags = 16
	DomainBlockRebaseCopyDev        DomainBlockRebaseFlags = 32
	DomainBlockRebaseBandwidthBytes DomainBlockRebaseFlags = 64
)

// DomainBlockCopyFlags as declared in libvirt/libvirt-domain.h:2555
type DomainBlockCopyFlags int32

// DomainBlockCopyFlags enumeration from libvirt/libvirt-domain.h:2555
const (
	DomainBlockCopyShallow      DomainBlockCopyFlags = 1
	DomainBlockCopyReuseExt     DomainBlockCopyFlags = 2
	DomainBlockCopyTransientJob DomainBlockCopyFlags = 4
)

// DomainBlockCommitFlags as declared in libvirt/libvirt-domain.h:2620
type DomainBlockCommitFlags int32

// DomainBlockCommitFlags enumeration from libvirt/libvirt-domain.h:2620
const (
	DomainBlockCommitShallow        DomainBlockCommitFlags = 1
	DomainBlockCommitDelete         DomainBlockCommitFlags = 2
	DomainBlockCommitActive         DomainBlockCommitFlags = 4
	DomainBlockCommitRelative       DomainBlockCommitFlags = 8
	DomainBlockCommitBandwidthBytes DomainBlockCommitFlags = 16
)

// DomainDiskErrorCode as declared in libvirt/libvirt-domain.h:2811
type DomainDiskErrorCode int32

// DomainDiskErrorCode enumeration from libvirt/libvirt-domain.h:2811
const (
	DomainDiskErrorNone    DomainDiskErrorCode = iota
	DomainDiskErrorUnspec  DomainDiskErrorCode = 1
	DomainDiskErrorNoSpace DomainDiskErrorCode = 2
)

// KeycodeSet as declared in libvirt/libvirt-domain.h:2857
type KeycodeSet int32

// KeycodeSet enumeration from libvirt/libvirt-domain.h:2857
const (
	KeycodeSetLinux  KeycodeSet = iota
	KeycodeSetXt     KeycodeSet = 1
	KeycodeSetAtset1 KeycodeSet = 2
	KeycodeSetAtset2 KeycodeSet = 3
	KeycodeSetAtset3 KeycodeSet = 4
	KeycodeSetOsx    KeycodeSet = 5
	KeycodeSetXtKbd  KeycodeSet = 6
	KeycodeSetUsb    KeycodeSet = 7
	KeycodeSetWin32  KeycodeSet = 8
	KeycodeSetQnum   KeycodeSet = 9
)

// DomainProcessSignal as declared in libvirt/libvirt-domain.h:2966
type DomainProcessSignal int32

// DomainProcessSignal enumeration from libvirt/libvirt-domain.h:2966
const (
	DomainProcessSignalNop    DomainProcessSignal = iota
	DomainProcessSignalHup    DomainProcessSignal = 1
	DomainProcessSignalInt    DomainProcessSignal = 2
	DomainProcessSignalQuit   DomainProcessSignal = 3
	DomainProcessSignalIll    DomainProcessSignal = 4
	DomainProcessSignalTrap   DomainProcessSignal = 5
	DomainProcessSignalAbrt   DomainProcessSignal = 6
	DomainProcessSignalBus    DomainProcessSignal = 7
	DomainProcessSignalFpe    DomainProcessSignal = 8
	DomainProcessSignalKill   DomainProcessSignal = 9
	DomainProcessSignalUsr1   DomainProcessSignal = 10
	DomainProcessSignalSegv   DomainProcessSignal = 11
	DomainProcessSignalUsr2   DomainProcessSignal = 12
	DomainProcessSignalPipe   DomainProcessSignal = 13
	DomainProcessSignalAlrm   DomainProcessSignal = 14
	DomainProcessSignalTerm   DomainProcessSignal = 15
	DomainProcessSignalStkflt DomainProcessSignal = 16
	DomainProcessSignalChld   DomainProcessSignal = 17
	DomainProcessSignalCont   DomainProcessSignal = 18
	DomainProcessSignalStop   DomainProcessSignal = 19
	DomainProcessSignalTstp   DomainProcessSignal = 20
	DomainProcessSignalTtin   DomainProcessSignal = 21
	DomainProcessSignalTtou   DomainProcessSignal = 22
	DomainProcessSignalUrg    DomainProcessSignal = 23
	DomainProcessSignalXcpu   DomainProcessSignal = 24
	DomainProcessSignalXfsz   DomainProcessSignal = 25
	DomainProcessSignalVtalrm DomainProcessSignal = 26
	DomainProcessSignalProf   DomainProcessSignal = 27
	DomainProcessSignalWinch  DomainProcessSignal = 28
	DomainProcessSignalPoll   DomainProcessSignal = 29
	DomainProcessSignalPwr    DomainProcessSignal = 30
	DomainProcessSignalSys    DomainProcessSignal = 31
	DomainProcessSignalRt0    DomainProcessSignal = 32
	DomainProcessSignalRt1    DomainProcessSignal = 33
	DomainProcessSignalRt2    DomainProcessSignal = 34
	DomainProcessSignalRt3    DomainProcessSignal = 35
	DomainProcessSignalRt4    DomainProcessSignal = 36
	DomainProcessSignalRt5    DomainProcessSignal = 37
	DomainProcessSignalRt6    DomainProcessSignal = 38
	DomainProcessSignalRt7    DomainProcessSignal = 39
	DomainProcessSignalRt8    DomainProcessSignal = 40
	DomainProcessSignalRt9    DomainProcessSignal = 41
	DomainProcessSignalRt10   DomainProcessSignal = 42
	DomainProcessSignalRt11   DomainProcessSignal = 43
	DomainProcessSignalRt12   DomainProcessSignal = 44
	DomainProcessSignalRt13   DomainProcessSignal = 45
	DomainProcessSignalRt14   DomainProcessSignal = 46
	DomainProcessSignalRt15   DomainProcessSignal = 47
	DomainProcessSignalRt16   DomainProcessSignal = 48
	DomainProcessSignalRt17   DomainProcessSignal = 49
	DomainProcessSignalRt18   DomainProcessSignal = 50
	DomainProcessSignalRt19   DomainProcessSignal = 51
	DomainProcessSignalRt20   DomainProcessSignal = 52
	DomainProcessSignalRt21   DomainProcessSignal = 53
	DomainProcessSignalRt22   DomainProcessSignal = 54
	DomainProcessSignalRt23   DomainProcessSignal = 55
	DomainProcessSignalRt24   DomainProcessSignal = 56
	DomainProcessSignalRt25   DomainProcessSignal = 57
	DomainProcessSignalRt26   DomainProcessSignal = 58
	DomainProcessSignalRt27   DomainProcessSignal = 59
	DomainProcessSignalRt28   DomainProcessSignal = 60
	DomainProcessSignalRt29   DomainProcessSignal = 61
	DomainProcessSignalRt30   DomainProcessSignal = 62
	DomainProcessSignalRt31   DomainProcessSignal = 63
	DomainProcessSignalRt32   DomainProcessSignal = 64
)

// DomainEventType as declared in libvirt/libvirt-domain.h:3004
type DomainEventType int32

// DomainEventType enumeration from libvirt/libvirt-domain.h:3004
const (
	DomainEventDefined     DomainEventType = iota
	DomainEventUndefined   DomainEventType = 1
	DomainEventStarted     DomainEventType = 2
	DomainEventSuspended   DomainEventType = 3
	DomainEventResumed     DomainEventType = 4
	DomainEventStopped     DomainEventType = 5
	DomainEventShutdown    DomainEventType = 6
	DomainEventPmsuspended DomainEventType = 7
	DomainEventCrashed     DomainEventType = 8
)

// DomainEventDefinedDetailType as declared in libvirt/libvirt-domain.h:3020
type DomainEventDefinedDetailType int32

// DomainEventDefinedDetailType enumeration from libvirt/libvirt-domain.h:3020
const (
	DomainEventDefinedAdded        DomainEventDefinedDetailType = iota
	DomainEventDefinedUpdated      DomainEventDefinedDetailType = 1
	DomainEventDefinedRenamed      DomainEventDefinedDetailType = 2
	DomainEventDefinedFromSnapshot DomainEventDefinedDetailType = 3
)

// DomainEventUndefinedDetailType as declared in libvirt/libvirt-domain.h:3034
type DomainEventUndefinedDetailType int32

// DomainEventUndefinedDetailType enumeration from libvirt/libvirt-domain.h:3034
const (
	DomainEventUndefinedRemoved DomainEventUndefinedDetailType = iota
	DomainEventUndefinedRenamed DomainEventUndefinedDetailType = 1
)

// DomainEventStartedDetailType as declared in libvirt/libvirt-domain.h:3051
type DomainEventStartedDetailType int32

// DomainEventStartedDetailType enumeration from libvirt/libvirt-domain.h:3051
const (
	DomainEventStartedBooted       DomainEventStartedDetailType = iota
	DomainEventStartedMigrated     DomainEventStartedDetailType = 1
	DomainEventStartedRestored     DomainEventStartedDetailType = 2
	DomainEventStartedFromSnapshot DomainEventStartedDetailType = 3
	DomainEventStartedWakeup       DomainEventStartedDetailType = 4
)

// DomainEventSuspendedDetailType as declared in libvirt/libvirt-domain.h:3072
type DomainEventSuspendedDetailType int32

// DomainEventSuspendedDetailType enumeration from libvirt/libvirt-domain.h:3072
const (
	DomainEventSuspendedPaused         DomainEventSuspendedDetailType = iota
	DomainEventSuspendedMigrated       DomainEventSuspendedDetailType = 1
	DomainEventSuspendedIoerror        DomainEventSuspendedDetailType = 2
	DomainEventSuspendedWatchdog       DomainEventSuspendedDetailType = 3
	DomainEventSuspendedRestored       DomainEventSuspendedDetailType = 4
	DomainEventSuspendedFromSnapshot   DomainEventSuspendedDetailType = 5
	DomainEventSuspendedAPIError       DomainEventSuspendedDetailType = 6
	DomainEventSuspendedPostcopy       DomainEventSuspendedDetailType = 7
	DomainEventSuspendedPostcopyFailed DomainEventSuspendedDetailType = 8
)

// DomainEventResumedDetailType as declared in libvirt/libvirt-domain.h:3089
type DomainEventResumedDetailType int32

// DomainEventResumedDetailType enumeration from libvirt/libvirt-domain.h:3089
const (
	DomainEventResumedUnpaused     DomainEventResumedDetailType = iota
	DomainEventResumedMigrated     DomainEventResumedDetailType = 1
	DomainEventResumedFromSnapshot DomainEventResumedDetailType = 2
	DomainEventResumedPostcopy     DomainEventResumedDetailType = 3
)

// DomainEventStoppedDetailType as declared in libvirt/libvirt-domain.h:3108
type DomainEventStoppedDetailType int32

// DomainEventStoppedDetailType enumeration from libvirt/libvirt-domain.h:3108
const (
	DomainEventStoppedShutdown     DomainEventStoppedDetailType = iota
	DomainEventStoppedDestroyed    DomainEventStoppedDetailType = 1
	DomainEventStoppedCrashed      DomainEventStoppedDetailType = 2
	DomainEventStoppedMigrated     DomainEventStoppedDetailType = 3
	DomainEventStoppedSaved        DomainEventStoppedDetailType = 4
	DomainEventStoppedFailed       DomainEventStoppedDetailType = 5
	DomainEventStoppedFromSnapshot DomainEventStoppedDetailType = 6
)

// DomainEventShutdownDetailType as declared in libvirt/libvirt-domain.h:3131
type DomainEventShutdownDetailType int32

// DomainEventShutdownDetailType enumeration from libvirt/libvirt-domain.h:3131
const (
	DomainEventShutdownFinished DomainEventShutdownDetailType = iota
	DomainEventShutdownGuest    DomainEventShutdownDetailType = 1
	DomainEventShutdownHost     DomainEventShutdownDetailType = 2
)

// DomainEventPMSuspendedDetailType as declared in libvirt/libvirt-domain.h:3145
type DomainEventPMSuspendedDetailType int32

// DomainEventPMSuspendedDetailType enumeration from libvirt/libvirt-domain.h:3145
const (
	DomainEventPmsuspendedMemory DomainEventPMSuspendedDetailType = iota
	DomainEventPmsuspendedDisk   DomainEventPMSuspendedDetailType = 1
)

// DomainEventCrashedDetailType as declared in libvirt/libvirt-domain.h:3158
type DomainEventCrashedDetailType int32

// DomainEventCrashedDetailType enumeration from libvirt/libvirt-domain.h:3158
const (
	DomainEventCrashedPanicked DomainEventCrashedDetailType = iota
)

// DomainJobType as declared in libvirt/libvirt-domain.h:3202
type DomainJobType int32

// DomainJobType enumeration from libvirt/libvirt-domain.h:3202
const (
	DomainJobNone      DomainJobType = iota
	DomainJobBounded   DomainJobType = 1
	DomainJobUnbounded DomainJobType = 2
	DomainJobCompleted DomainJobType = 3
	DomainJobFailed    DomainJobType = 4
	DomainJobCancelled DomainJobType = 5
)

// DomainGetJobStatsFlags as declared in libvirt/libvirt-domain.h:3249
type DomainGetJobStatsFlags int32

// DomainGetJobStatsFlags enumeration from libvirt/libvirt-domain.h:3249
const (
	DomainJobStatsCompleted DomainGetJobStatsFlags = 1
)

// DomainJobOperation as declared in libvirt/libvirt-domain.h:3274
type DomainJobOperation int32

// DomainJobOperation enumeration from libvirt/libvirt-domain.h:3274
const (
	DomainJobOperationStrUnknown        DomainJobOperation = iota
	DomainJobOperationStrStart          DomainJobOperation = 1
	DomainJobOperationStrSave           DomainJobOperation = 2
	DomainJobOperationStrRestore        DomainJobOperation = 3
	DomainJobOperationStrMigrationIn    DomainJobOperation = 4
	DomainJobOperationStrMigrationOut   DomainJobOperation = 5
	DomainJobOperationStrSnapshot       DomainJobOperation = 6
	DomainJobOperationStrSnapshotRevert DomainJobOperation = 7
	DomainJobOperationStrDump           DomainJobOperation = 8
)

// DomainEventWatchdogAction as declared in libvirt/libvirt-domain.h:3628
type DomainEventWatchdogAction int32

// DomainEventWatchdogAction enumeration from libvirt/libvirt-domain.h:3628
const (
	DomainEventWatchdogNone      DomainEventWatchdogAction = iota
	DomainEventWatchdogPause     DomainEventWatchdogAction = 1
	DomainEventWatchdogReset     DomainEventWatchdogAction = 2
	DomainEventWatchdogPoweroff  DomainEventWatchdogAction = 3
	DomainEventWatchdogShutdown  DomainEventWatchdogAction = 4
	DomainEventWatchdogDebug     DomainEventWatchdogAction = 5
	DomainEventWatchdogInjectnmi DomainEventWatchdogAction = 6
)

// DomainEventIOErrorAction as declared in libvirt/libvirt-domain.h:3659
type DomainEventIOErrorAction int32

// DomainEventIOErrorAction enumeration from libvirt/libvirt-domain.h:3659
const (
	DomainEventIoErrorNone   DomainEventIOErrorAction = iota
	DomainEventIoErrorPause  DomainEventIOErrorAction = 1
	DomainEventIoErrorReport DomainEventIOErrorAction = 2
)

// DomainEventGraphicsPhase as declared in libvirt/libvirt-domain.h:3722
type DomainEventGraphicsPhase int32

// DomainEventGraphicsPhase enumeration from libvirt/libvirt-domain.h:3722
const (
	DomainEventGraphicsConnect    DomainEventGraphicsPhase = iota
	DomainEventGraphicsInitialize DomainEventGraphicsPhase = 1
	DomainEventGraphicsDisconnect DomainEventGraphicsPhase = 2
)

// DomainEventGraphicsAddressType as declared in libvirt/libvirt-domain.h:3737
type DomainEventGraphicsAddressType int32

// DomainEventGraphicsAddressType enumeration from libvirt/libvirt-domain.h:3737
const (
	DomainEventGraphicsAddressIpv4 DomainEventGraphicsAddressType = iota
	DomainEventGraphicsAddressIpv6 DomainEventGraphicsAddressType = 1
	DomainEventGraphicsAddressUnix DomainEventGraphicsAddressType = 2
)

// ConnectDomainEventBlockJobStatus as declared in libvirt/libvirt-domain.h:3825
type ConnectDomainEventBlockJobStatus int32

// ConnectDomainEventBlockJobStatus enumeration from libvirt/libvirt-domain.h:3825
const (
	DomainBlockJobCompleted ConnectDomainEventBlockJobStatus = iota
	DomainBlockJobFailed    ConnectDomainEventBlockJobStatus = 1
	DomainBlockJobCanceled  ConnectDomainEventBlockJobStatus = 2
	DomainBlockJobReady     ConnectDomainEventBlockJobStatus = 3
)

// ConnectDomainEventDiskChangeReason as declared in libvirt/libvirt-domain.h:3875
type ConnectDomainEventDiskChangeReason int32

// ConnectDomainEventDiskChangeReason enumeration from libvirt/libvirt-domain.h:3875
const (
	DomainEventDiskChangeMissingOnStart ConnectDomainEventDiskChangeReason = iota
	DomainEventDiskDropMissingOnStart   ConnectDomainEventDiskChangeReason = 1
)

// DomainEventTrayChangeReason as declared in libvirt/libvirt-domain.h:3916
type DomainEventTrayChangeReason int32

// DomainEventTrayChangeReason enumeration from libvirt/libvirt-domain.h:3916
const (
	DomainEventTrayChangeOpen  DomainEventTrayChangeReason = iota
	DomainEventTrayChangeClose DomainEventTrayChangeReason = 1
)

// ConnectDomainEventAgentLifecycleState as declared in libvirt/libvirt-domain.h:4431
type ConnectDomainEventAgentLifecycleState int32

// ConnectDomainEventAgentLifecycleState enumeration from libvirt/libvirt-domain.h:4431
const (
	ConnectDomainEventAgentLifecycleStateConnected    ConnectDomainEventAgentLifecycleState = 1
	ConnectDomainEventAgentLifecycleStateDisconnected ConnectDomainEventAgentLifecycleState = 2
)

// ConnectDomainEventAgentLifecycleReason as declared in libvirt/libvirt-domain.h:4441
type ConnectDomainEventAgentLifecycleReason int32

// ConnectDomainEventAgentLifecycleReason enumeration from libvirt/libvirt-domain.h:4441
const (
	ConnectDomainEventAgentLifecycleReasonUnknown       ConnectDomainEventAgentLifecycleReason = iota
	ConnectDomainEventAgentLifecycleReasonDomainStarted ConnectDomainEventAgentLifecycleReason = 1
	ConnectDomainEventAgentLifecycleReasonChannel       ConnectDomainEventAgentLifecycleReason = 2
)

// DomainEventID as declared in libvirt/libvirt-domain.h:4545
type DomainEventID int32

// DomainEventID enumeration from libvirt/libvirt-domain.h:4545
const (
	DomainEventIDLifecycle           DomainEventID = iota
	DomainEventIDReboot              DomainEventID = 1
	DomainEventIDRtcChange           DomainEventID = 2
	DomainEventIDWatchdog            DomainEventID = 3
	DomainEventIDIoError             DomainEventID = 4
	DomainEventIDGraphics            DomainEventID = 5
	DomainEventIDIoErrorReason       DomainEventID = 6
	DomainEventIDControlError        DomainEventID = 7
	DomainEventIDBlockJob            DomainEventID = 8
	DomainEventIDDiskChange          DomainEventID = 9
	DomainEventIDTrayChange          DomainEventID = 10
	DomainEventIDPmwakeup            DomainEventID = 11
	DomainEventIDPmsuspend           DomainEventID = 12
	DomainEventIDBalloonChange       DomainEventID = 13
	DomainEventIDPmsuspendDisk       DomainEventID = 14
	DomainEventIDDeviceRemoved       DomainEventID = 15
	DomainEventIDBlockJob2           DomainEventID = 16
	DomainEventIDTunable             DomainEventID = 17
	DomainEventIDAgentLifecycle      DomainEventID = 18
	DomainEventIDDeviceAdded         DomainEventID = 19
	DomainEventIDMigrationIteration  DomainEventID = 20
	DomainEventIDJobCompleted        DomainEventID = 21
	DomainEventIDDeviceRemovalFailed DomainEventID = 22
	DomainEventIDMetadataChange      DomainEventID = 23
	DomainEventIDBlockThreshold      DomainEventID = 24
)

// DomainConsoleFlags as declared in libvirt/libvirt-domain.h:4572
type DomainConsoleFlags int32

// DomainConsoleFlags enumeration from libvirt/libvirt-domain.h:4572
const (
	DomainConsoleForce DomainConsoleFlags = 1
	DomainConsoleSafe  DomainConsoleFlags = 2
)

// DomainChannelFlags as declared in libvirt/libvirt-domain.h:4588
type DomainChannelFlags int32

// DomainChannelFlags enumeration from libvirt/libvirt-domain.h:4588
const (
	DomainChannelForce DomainChannelFlags = 1
)

// DomainOpenGraphicsFlags as declared in libvirt/libvirt-domain.h:4597
type DomainOpenGraphicsFlags int32

// DomainOpenGraphicsFlags enumeration from libvirt/libvirt-domain.h:4597
const (
	DomainOpenGraphicsSkipauth DomainOpenGraphicsFlags = 1
)

// DomainSetTimeFlags as declared in libvirt/libvirt-domain.h:4654
type DomainSetTimeFlags int32

// DomainSetTimeFlags enumeration from libvirt/libvirt-domain.h:4654
const (
	DomainTimeSync DomainSetTimeFlags = 1
)

// SchedParameterType as declared in libvirt/libvirt-domain.h:4675
type SchedParameterType int32

// SchedParameterType enumeration from libvirt/libvirt-domain.h:4675
const (
	DomainSchedFieldInt     SchedParameterType = 1
	DomainSchedFieldUint    SchedParameterType = 2
	DomainSchedFieldLlong   SchedParameterType = 3
	DomainSchedFieldUllong  SchedParameterType = 4
	DomainSchedFieldDouble  SchedParameterType = 5
	DomainSchedFieldBoolean SchedParameterType = 6
)

// BlkioParameterType as declared in libvirt/libvirt-domain.h:4719
type BlkioParameterType int32

// BlkioParameterType enumeration from libvirt/libvirt-domain.h:4719
const (
	DomainBlkioParamInt     BlkioParameterType = 1
	DomainBlkioParamUint    BlkioParameterType = 2
	DomainBlkioParamLlong   BlkioParameterType = 3
	DomainBlkioParamUllong  BlkioParameterType = 4
	DomainBlkioParamDouble  BlkioParameterType = 5
	DomainBlkioParamBoolean BlkioParameterType = 6
)

// MemoryParameterType as declared in libvirt/libvirt-domain.h:4763
type MemoryParameterType int32

// MemoryParameterType enumeration from libvirt/libvirt-domain.h:4763
const (
	DomainMemoryParamInt     MemoryParameterType = 1
	DomainMemoryParamUint    MemoryParameterType = 2
	DomainMemoryParamLlong   MemoryParameterType = 3
	DomainMemoryParamUllong  MemoryParameterType = 4
	DomainMemoryParamDouble  MemoryParameterType = 5
	DomainMemoryParamBoolean MemoryParameterType = 6
)

// DomainInterfaceAddressesSource as declared in libvirt/libvirt-domain.h:4801
type DomainInterfaceAddressesSource int32

// DomainInterfaceAddressesSource enumeration from libvirt/libvirt-domain.h:4801
const (
	DomainInterfaceAddressesSrcLease DomainInterfaceAddressesSource = iota
	DomainInterfaceAddressesSrcAgent DomainInterfaceAddressesSource = 1
	DomainInterfaceAddressesSrcArp   DomainInterfaceAddressesSource = 2
)

// DomainSetUserPasswordFlags as declared in libvirt/libvirt-domain.h:4829
type DomainSetUserPasswordFlags int32

// DomainSetUserPasswordFlags enumeration from libvirt/libvirt-domain.h:4829
const (
	DomainPasswordEncrypted DomainSetUserPasswordFlags = 1
)

// DomainLifecycle as declared in libvirt/libvirt-domain.h:4868
type DomainLifecycle int32

// DomainLifecycle enumeration from libvirt/libvirt-domain.h:4868
const (
	DomainLifecyclePoweroff DomainLifecycle = iota
	DomainLifecycleReboot   DomainLifecycle = 1
	DomainLifecycleCrash    DomainLifecycle = 2
)

// DomainLifecycleAction as declared in libvirt/libvirt-domain.h:4881
type DomainLifecycleAction int32

// DomainLifecycleAction enumeration from libvirt/libvirt-domain.h:4881
const (
	DomainLifecycleActionDestroy         DomainLifecycleAction = iota
	DomainLifecycleActionRestart         DomainLifecycleAction = 1
	DomainLifecycleActionRestartRename   DomainLifecycleAction = 2
	DomainLifecycleActionPreserve        DomainLifecycleAction = 3
	DomainLifecycleActionCoredumpDestroy DomainLifecycleAction = 4
	DomainLifecycleActionCoredumpRestart DomainLifecycleAction = 5
)

// DomainGuestInfoTypes as declared in libvirt/libvirt-domain.h:4911
type DomainGuestInfoTypes int32

// DomainGuestInfoTypes enumeration from libvirt/libvirt-domain.h:4911
const (
	DomainGuestInfoUsers      DomainGuestInfoTypes = 1
	DomainGuestInfoOs         DomainGuestInfoTypes = 2
	DomainGuestInfoTimezone   DomainGuestInfoTypes = 4
	DomainGuestInfoHostname   DomainGuestInfoTypes = 8
	DomainGuestInfoFilesystem DomainGuestInfoTypes = 16
)

// DomainSnapshotCreateFlags as declared in libvirt/libvirt-domain-snapshot.h:76
type DomainSnapshotCreateFlags int32

// DomainSnapshotCreateFlags enumeration from libvirt/libvirt-domain-snapshot.h:76
const (
	DomainSnapshotCreateRedefine   DomainSnapshotCreateFlags = 1
	DomainSnapshotCreateCurrent    DomainSnapshotCreateFlags = 2
	DomainSnapshotCreateNoMetadata DomainSnapshotCreateFlags = 4
	DomainSnapshotCreateHalt       DomainSnapshotCreateFlags = 8
	DomainSnapshotCreateDiskOnly   DomainSnapshotCreateFlags = 16
	DomainSnapshotCreateReuseExt   DomainSnapshotCreateFlags = 32
	DomainSnapshotCreateQuiesce    DomainSnapshotCreateFlags = 64
	DomainSnapshotCreateAtomic     DomainSnapshotCreateFlags = 128
	DomainSnapshotCreateLive       DomainSnapshotCreateFlags = 256
	DomainSnapshotCreateValidate   DomainSnapshotCreateFlags = 512
)

// DomainSnapshotXMLFlags as declared in libvirt/libvirt-domain-snapshot.h:85
type DomainSnapshotXMLFlags int32

// DomainSnapshotXMLFlags enumeration from libvirt/libvirt-domain-snapshot.h:85
const (
	DomainSnapshotXMLSecure DomainSnapshotXMLFlags = 1
)

// DomainSnapshotListFlags as declared in libvirt/libvirt-domain-snapshot.h:144
type DomainSnapshotListFlags int32

// DomainSnapshotListFlags enumeration from libvirt/libvirt-domain-snapshot.h:144
const (
	DomainSnapshotListRoots       DomainSnapshotListFlags = 1
	DomainSnapshotListDescendants DomainSnapshotListFlags = 1
	DomainSnapshotListLeaves      DomainSnapshotListFlags = 4
	DomainSnapshotListNoLeaves    DomainSnapshotListFlags = 8
	DomainSnapshotListMetadata    DomainSnapshotListFlags = 2
	DomainSnapshotListNoMetadata  DomainSnapshotListFlags = 16
	DomainSnapshotListInactive    DomainSnapshotListFlags = 32
	DomainSnapshotListActive      DomainSnapshotListFlags = 64
	DomainSnapshotListDiskOnly    DomainSnapshotListFlags = 128
	DomainSnapshotListInternal    DomainSnapshotListFlags = 256
	DomainSnapshotListExternal    DomainSnapshotListFlags = 512
	DomainSnapshotListTopological DomainSnapshotListFlags = 1024
)

// DomainSnapshotRevertFlags as declared in libvirt/libvirt-domain-snapshot.h:201
type DomainSnapshotRevertFlags int32

// DomainSnapshotRevertFlags enumeration from libvirt/libvirt-domain-snapshot.h:201
const (
	DomainSnapshotRevertRunning DomainSnapshotRevertFlags = 1
	DomainSnapshotRevertPaused  DomainSnapshotRevertFlags = 2
	DomainSnapshotRevertForce   DomainSnapshotRevertFlags = 4
)

// DomainSnapshotDeleteFlags as declared in libvirt/libvirt-domain-snapshot.h:215
type DomainSnapshotDeleteFlags int32

// DomainSnapshotDeleteFlags enumeration from libvirt/libvirt-domain-snapshot.h:215
const (
	DomainSnapshotDeleteChildren     DomainSnapshotDeleteFlags = 1
	DomainSnapshotDeleteMetadataOnly DomainSnapshotDeleteFlags = 2
	DomainSnapshotDeleteChildrenOnly DomainSnapshotDeleteFlags = 4
)

// EventHandleType as declared in libvirt/libvirt-event.h:43
type EventHandleType int32

// EventHandleType enumeration from libvirt/libvirt-event.h:43
const (
	EventHandleReadable EventHandleType = 1
	EventHandleWritable EventHandleType = 2
	EventHandleError    EventHandleType = 4
	EventHandleHangup   EventHandleType = 8
)

// ConnectListAllInterfacesFlags as declared in libvirt/libvirt-interface.h:64
type ConnectListAllInterfacesFlags int32

// ConnectListAllInterfacesFlags enumeration from libvirt/libvirt-interface.h:64
const (
	ConnectListInterfacesInactive ConnectListAllInterfacesFlags = 1
	ConnectListInterfacesActive   ConnectListAllInterfacesFlags = 2
)

// InterfaceXMLFlags as declared in libvirt/libvirt-interface.h:80
type InterfaceXMLFlags int32

// InterfaceXMLFlags enumeration from libvirt/libvirt-interface.h:80
const (
	InterfaceXMLInactive InterfaceXMLFlags = 1
)

// NetworkXMLFlags as declared in libvirt/libvirt-network.h:32
type NetworkXMLFlags int32

// NetworkXMLFlags enumeration from libvirt/libvirt-network.h:32
const (
	NetworkXMLInactive NetworkXMLFlags = 1
)

// ConnectListAllNetworksFlags as declared in libvirt/libvirt-network.h:100
type ConnectListAllNetworksFlags int32

// ConnectListAllNetworksFlags enumeration from libvirt/libvirt-network.h:100
const (
	ConnectListNetworksInactive    ConnectListAllNetworksFlags = 1
	ConnectListNetworksActive      ConnectListAllNetworksFlags = 2
	ConnectListNetworksPersistent  ConnectListAllNetworksFlags = 4
	ConnectListNetworksTransient   ConnectListAllNetworksFlags = 8
	ConnectListNetworksAutostart   ConnectListAllNetworksFlags = 16
	ConnectListNetworksNoAutostart ConnectListAllNetworksFlags = 32
)

// NetworkUpdateCommand as declared in libvirt/libvirt-network.h:149
type NetworkUpdateCommand int32

// NetworkUpdateCommand enumeration from libvirt/libvirt-network.h:149
const (
	NetworkUpdateCommandNone     NetworkUpdateCommand = iota
	NetworkUpdateCommandModify   NetworkUpdateCommand = 1
	NetworkUpdateCommandDelete   NetworkUpdateCommand = 2
	NetworkUpdateCommandAddLast  NetworkUpdateCommand = 3
	NetworkUpdateCommandAddFirst NetworkUpdateCommand = 4
)

// NetworkUpdateSection as declared in libvirt/libvirt-network.h:175
type NetworkUpdateSection int32

// NetworkUpdateSection enumeration from libvirt/libvirt-network.h:175
const (
	NetworkSectionNone             NetworkUpdateSection = iota
	NetworkSectionBridge           NetworkUpdateSection = 1
	NetworkSectionDomain           NetworkUpdateSection = 2
	NetworkSectionIP               NetworkUpdateSection = 3
	NetworkSectionIPDhcpHost       NetworkUpdateSection = 4
	NetworkSectionIPDhcpRange      NetworkUpdateSection = 5
	NetworkSectionForward          NetworkUpdateSection = 6
	NetworkSectionForwardInterface NetworkUpdateSection = 7
	NetworkSectionForwardPf        NetworkUpdateSection = 8
	NetworkSectionPortgroup        NetworkUpdateSection = 9
	NetworkSectionDNSHost          NetworkUpdateSection = 10
	NetworkSectionDNSTxt           NetworkUpdateSection = 11
	NetworkSectionDNSSrv           NetworkUpdateSection = 12
)

// NetworkUpdateFlags as declared in libvirt/libvirt-network.h:187
type NetworkUpdateFlags int32

// NetworkUpdateFlags enumeration from libvirt/libvirt-network.h:187
const (
	NetworkUpdateAffectCurrent NetworkUpdateFlags = iota
	NetworkUpdateAffectLive    NetworkUpdateFlags = 1
	NetworkUpdateAffectConfig  NetworkUpdateFlags = 2
)

// NetworkEventLifecycleType as declared in libvirt/libvirt-network.h:245
type NetworkEventLifecycleType int32

// NetworkEventLifecycleType enumeration from libvirt/libvirt-network.h:245
const (
	NetworkEventDefined   NetworkEventLifecycleType = iota
	NetworkEventUndefined NetworkEventLifecycleType = 1
	NetworkEventStarted   NetworkEventLifecycleType = 2
	NetworkEventStopped   NetworkEventLifecycleType = 3
)

// NetworkEventID as declared in libvirt/libvirt-network.h:293
type NetworkEventID int32

// NetworkEventID enumeration from libvirt/libvirt-network.h:293
const (
	NetworkEventIDLifecycle NetworkEventID = iota
)

// IPAddrType as declared in libvirt/libvirt-network.h:302
type IPAddrType int32

// IPAddrType enumeration from libvirt/libvirt-network.h:302
const (
	IPAddrTypeIpv4 IPAddrType = iota
	IPAddrTypeIpv6 IPAddrType = 1
)

// NetworkPortCreateFlags as declared in libvirt/libvirt-network.h:363
type NetworkPortCreateFlags int32

// NetworkPortCreateFlags enumeration from libvirt/libvirt-network.h:363
const (
	NetworkPortCreateReclaim NetworkPortCreateFlags = 1
)

// ConnectListAllNodeDeviceFlags as declared in libvirt/libvirt-nodedev.h:84
type ConnectListAllNodeDeviceFlags int32

// ConnectListAllNodeDeviceFlags enumeration from libvirt/libvirt-nodedev.h:84
const (
	ConnectListNodeDevicesCapSystem       ConnectListAllNodeDeviceFlags = 1
	ConnectListNodeDevicesCapPciDev       ConnectListAllNodeDeviceFlags = 2
	ConnectListNodeDevicesCapUsbDev       ConnectListAllNodeDeviceFlags = 4
	ConnectListNodeDevicesCapUsbInterface ConnectListAllNodeDeviceFlags = 8
	ConnectListNodeDevicesCapNet          ConnectListAllNodeDeviceFlags = 16
	ConnectListNodeDevicesCapScsiHost     ConnectListAllNodeDeviceFlags = 32
	ConnectListNodeDevicesCapScsiTarget   ConnectListAllNodeDeviceFlags = 64
	ConnectListNodeDevicesCapScsi         ConnectListAllNodeDeviceFlags = 128
	ConnectListNodeDevicesCapStorage      ConnectListAllNodeDeviceFlags = 256
	ConnectListNodeDevicesCapFcHost       ConnectListAllNodeDeviceFlags = 512
	ConnectListNodeDevicesCapVports       ConnectListAllNodeDeviceFlags = 1024
	ConnectListNodeDevicesCapScsiGeneric  ConnectListAllNodeDeviceFlags = 2048
	ConnectListNodeDevicesCapDrm          ConnectListAllNodeDeviceFlags = 4096
	ConnectListNodeDevicesCapMdevTypes    ConnectListAllNodeDeviceFlags = 8192
	ConnectListNodeDevicesCapMdev         ConnectListAllNodeDeviceFlags = 16384
	ConnectListNodeDevicesCapCcwDev       ConnectListAllNodeDeviceFlags = 32768
)

// NodeDeviceEventID as declared in libvirt/libvirt-nodedev.h:154
type NodeDeviceEventID int32

// NodeDeviceEventID enumeration from libvirt/libvirt-nodedev.h:154
const (
	NodeDeviceEventIDLifecycle NodeDeviceEventID = iota
	NodeDeviceEventIDUpdate    NodeDeviceEventID = 1
)

// NodeDeviceEventLifecycleType as declared in libvirt/libvirt-nodedev.h:196
type NodeDeviceEventLifecycleType int32

// NodeDeviceEventLifecycleType enumeration from libvirt/libvirt-nodedev.h:196
const (
	NodeDeviceEventCreated NodeDeviceEventLifecycleType = iota
	NodeDeviceEventDeleted NodeDeviceEventLifecycleType = 1
)

// SecretUsageType as declared in libvirt/libvirt-secret.h:56
type SecretUsageType int32

// SecretUsageType enumeration from libvirt/libvirt-secret.h:56
const (
	SecretUsageTypeNone   SecretUsageType = iota
	SecretUsageTypeVolume SecretUsageType = 1
	SecretUsageTypeCeph   SecretUsageType = 2
	SecretUsageTypeIscsi  SecretUsageType = 3
	SecretUsageTypeTLS    SecretUsageType = 4
	SecretUsageTypeVtpm   SecretUsageType = 5
)

// ConnectListAllSecretsFlags as declared in libvirt/libvirt-secret.h:79
type ConnectListAllSecretsFlags int32

// ConnectListAllSecretsFlags enumeration from libvirt/libvirt-secret.h:79
const (
	ConnectListSecretsEphemeral   ConnectListAllSecretsFlags = 1
	ConnectListSecretsNoEphemeral ConnectListAllSecretsFlags = 2
	ConnectListSecretsPrivate     ConnectListAllSecretsFlags = 4
	ConnectListSecretsNoPrivate   ConnectListAllSecretsFlags = 8
)

// SecretEventID as declared in libvirt/libvirt-secret.h:140
type SecretEventID int32

// SecretEventID enumeration from libvirt/libvirt-secret.h:140
const (
	SecretEventIDLifecycle    SecretEventID = iota
	SecretEventIDValueChanged SecretEventID = 1
)

// SecretEventLifecycleType as declared in libvirt/libvirt-secret.h:182
type SecretEventLifecycleType int32

// SecretEventLifecycleType enumeration from libvirt/libvirt-secret.h:182
const (
	SecretEventDefined   SecretEventLifecycleType = iota
	SecretEventUndefined SecretEventLifecycleType = 1
)

// StoragePoolState as declared in libvirt/libvirt-storage.h:57
type StoragePoolState int32

// StoragePoolState enumeration from libvirt/libvirt-storage.h:57
const (
	StoragePoolInactive     StoragePoolState = iota
	StoragePoolBuilding     StoragePoolState = 1
	StoragePoolRunning      StoragePoolState = 2
	StoragePoolDegraded     StoragePoolState = 3
	StoragePoolInaccessible StoragePoolState = 4
)

// StoragePoolBuildFlags as declared in libvirt/libvirt-storage.h:65
type StoragePoolBuildFlags int32

// StoragePoolBuildFlags enumeration from libvirt/libvirt-storage.h:65
const (
	StoragePoolBuildNew         StoragePoolBuildFlags = iota
	StoragePoolBuildRepair      StoragePoolBuildFlags = 1
	StoragePoolBuildResize      StoragePoolBuildFlags = 2
	StoragePoolBuildNoOverwrite StoragePoolBuildFlags = 4
	StoragePoolBuildOverwrite   StoragePoolBuildFlags = 8
)

// StoragePoolDeleteFlags as declared in libvirt/libvirt-storage.h:70
type StoragePoolDeleteFlags int32

// StoragePoolDeleteFlags enumeration from libvirt/libvirt-storage.h:70
const (
	StoragePoolDeleteNormal StoragePoolDeleteFlags = iota
	StoragePoolDeleteZeroed StoragePoolDeleteFlags = 1
)

// StoragePoolCreateFlags as declared in libvirt/libvirt-storage.h:87
type StoragePoolCreateFlags int32

// StoragePoolCreateFlags enumeration from libvirt/libvirt-storage.h:87
const (
	StoragePoolCreateNormal               StoragePoolCreateFlags = iota
	StoragePoolCreateWithBuild            StoragePoolCreateFlags = 1
	StoragePoolCreateWithBuildOverwrite   StoragePoolCreateFlags = 2
	StoragePoolCreateWithBuildNoOverwrite StoragePoolCreateFlags = 4
)

// StorageVolType as declared in libvirt/libvirt-storage.h:129
type StorageVolType int32

// StorageVolType enumeration from libvirt/libvirt-storage.h:129
const (
	StorageVolFile    StorageVolType = iota
	StorageVolBlock   StorageVolType = 1
	StorageVolDir     StorageVolType = 2
	StorageVolNetwork StorageVolType = 3
	StorageVolNetdir  StorageVolType = 4
	StorageVolPloop   StorageVolType = 5
)

// StorageVolDeleteFlags as declared in libvirt/libvirt-storage.h:135
type StorageVolDeleteFlags int32

// StorageVolDeleteFlags enumeration from libvirt/libvirt-storage.h:135
const (
	StorageVolDeleteNormal        StorageVolDeleteFlags = iota
	StorageVolDeleteZeroed        StorageVolDeleteFlags = 1
	StorageVolDeleteWithSnapshots StorageVolDeleteFlags = 2
)

// StorageVolWipeAlgorithm as declared in libvirt/libvirt-storage.h:167
type StorageVolWipeAlgorithm int32

// StorageVolWipeAlgorithm enumeration from libvirt/libvirt-storage.h:167
const (
	StorageVolWipeAlgZero       StorageVolWipeAlgorithm = iota
	StorageVolWipeAlgNnsa       StorageVolWipeAlgorithm = 1
	StorageVolWipeAlgDod        StorageVolWipeAlgorithm = 2
	StorageVolWipeAlgBsi        StorageVolWipeAlgorithm = 3
	StorageVolWipeAlgGutmann    StorageVolWipeAlgorithm = 4
	StorageVolWipeAlgSchneier   StorageVolWipeAlgorithm = 5
	StorageVolWipeAlgPfitzner7  StorageVolWipeAlgorithm = 6
	StorageVolWipeAlgPfitzner33 StorageVolWipeAlgorithm = 7
	StorageVolWipeAlgRandom     StorageVolWipeAlgorithm = 8
	StorageVolWipeAlgTrim       StorageVolWipeAlgorithm = 9
)

// StorageVolInfoFlags as declared in libvirt/libvirt-storage.h:175
type StorageVolInfoFlags int32

// StorageVolInfoFlags enumeration from libvirt/libvirt-storage.h:175
const (
	StorageVolUseAllocation StorageVolInfoFlags = iota
	StorageVolGetPhysical   StorageVolInfoFlags = 1
)

// StorageXMLFlags as declared in libvirt/libvirt-storage.h:189
type StorageXMLFlags int32

// StorageXMLFlags enumeration from libvirt/libvirt-storage.h:189
const (
	StorageXMLInactive StorageXMLFlags = 1
)

// ConnectListAllStoragePoolsFlags as declared in libvirt/libvirt-storage.h:248
type ConnectListAllStoragePoolsFlags int32

// ConnectListAllStoragePoolsFlags enumeration from libvirt/libvirt-storage.h:248
const (
	ConnectListStoragePoolsInactive    ConnectListAllStoragePoolsFlags = 1
	ConnectListStoragePoolsActive      ConnectListAllStoragePoolsFlags = 2
	ConnectListStoragePoolsPersistent  ConnectListAllStoragePoolsFlags = 4
	ConnectListStoragePoolsTransient   ConnectListAllStoragePoolsFlags = 8
	ConnectListStoragePoolsAutostart   ConnectListAllStoragePoolsFlags = 16
	ConnectListStoragePoolsNoAutostart ConnectListAllStoragePoolsFlags = 32
	ConnectListStoragePoolsDir         ConnectListAllStoragePoolsFlags = 64
	ConnectListStoragePoolsFs          ConnectListAllStoragePoolsFlags = 128
	ConnectListStoragePoolsNetfs       ConnectListAllStoragePoolsFlags = 256
	ConnectListStoragePoolsLogical     ConnectListAllStoragePoolsFlags = 512
	ConnectListStoragePoolsDisk        ConnectListAllStoragePoolsFlags = 1024
	ConnectListStoragePoolsIscsi       ConnectListAllStoragePoolsFlags = 2048
	ConnectListStoragePoolsScsi        ConnectListAllStoragePoolsFlags = 4096
	ConnectListStoragePoolsMpath       ConnectListAllStoragePoolsFlags = 8192
	ConnectListStoragePoolsRbd         ConnectListAllStoragePoolsFlags = 16384
	ConnectListStoragePoolsSheepdog    ConnectListAllStoragePoolsFlags = 32768
	ConnectListStoragePoolsGluster     ConnectListAllStoragePoolsFlags = 65536
	ConnectListStoragePoolsZfs         ConnectListAllStoragePoolsFlags = 131072
	ConnectListStoragePoolsVstorage    ConnectListAllStoragePoolsFlags = 262144
	ConnectListStoragePoolsIscsiDirect ConnectListAllStoragePoolsFlags = 524288
)

// StorageVolCreateFlags as declared in libvirt/libvirt-storage.h:346
type StorageVolCreateFlags int32

// StorageVolCreateFlags enumeration from libvirt/libvirt-storage.h:346
const (
	StorageVolCreatePreallocMetadata StorageVolCreateFlags = 1
	StorageVolCreateReflink          StorageVolCreateFlags = 2
)

// StorageVolDownloadFlags as declared in libvirt/libvirt-storage.h:358
type StorageVolDownloadFlags int32

// StorageVolDownloadFlags enumeration from libvirt/libvirt-storage.h:358
const (
	StorageVolDownloadSparseStream StorageVolDownloadFlags = 1
)

// StorageVolUploadFlags as declared in libvirt/libvirt-storage.h:367
type StorageVolUploadFlags int32

// StorageVolUploadFlags enumeration from libvirt/libvirt-storage.h:367
const (
	StorageVolUploadSparseStream StorageVolUploadFlags = 1
)

// StorageVolResizeFlags as declared in libvirt/libvirt-storage.h:398
type StorageVolResizeFlags int32

// StorageVolResizeFlags enumeration from libvirt/libvirt-storage.h:398
const (
	StorageVolResizeAllocate StorageVolResizeFlags = 1
	StorageVolResizeDelta    StorageVolResizeFlags = 2
	StorageVolResizeShrink   StorageVolResizeFlags = 4
)

// StoragePoolEventID as declared in libvirt/libvirt-storage.h:434
type StoragePoolEventID int32

// StoragePoolEventID enumeration from libvirt/libvirt-storage.h:434
const (
	StoragePoolEventIDLifecycle StoragePoolEventID = iota
	StoragePoolEventIDRefresh   StoragePoolEventID = 1
)

// StoragePoolEventLifecycleType as declared in libvirt/libvirt-storage.h:480
type StoragePoolEventLifecycleType int32

// StoragePoolEventLifecycleType enumeration from libvirt/libvirt-storage.h:480
const (
	StoragePoolEventDefined   StoragePoolEventLifecycleType = iota
	StoragePoolEventUndefined StoragePoolEventLifecycleType = 1
	StoragePoolEventStarted   StoragePoolEventLifecycleType = 2
	StoragePoolEventStopped   StoragePoolEventLifecycleType = 3
	StoragePoolEventCreated   StoragePoolEventLifecycleType = 4
	StoragePoolEventDeleted   StoragePoolEventLifecycleType = 5
)

// StreamFlags as declared in libvirt/libvirt-stream.h:33
type StreamFlags int32

// StreamFlags enumeration from libvirt/libvirt-stream.h:33
const (
	StreamNonblock StreamFlags = 1
)

// StreamRecvFlagsValues as declared in libvirt/libvirt-stream.h:49
type StreamRecvFlagsValues int32

// StreamRecvFlagsValues enumeration from libvirt/libvirt-stream.h:49
const (
	StreamRecvStopAtHole StreamRecvFlagsValues = 1
)

// StreamEventType as declared in libvirt/libvirt-stream.h:237
type StreamEventType int32

// StreamEventType enumeration from libvirt/libvirt-stream.h:237
const (
	StreamEventReadable StreamEventType = 1
	StreamEventWritable StreamEventType = 2
	StreamEventError    StreamEventType = 4
	StreamEventHangup   StreamEventType = 8
)

// errorLevel as declared in libvirt/virterror.h:44
type errorLevel int32

// errorLevel enumeration from libvirt/virterror.h:44
const (
	errNone    errorLevel = iota
	errWarning errorLevel = 1
	errError   errorLevel = 2
)

// errorDomain as declared in libvirt/virterror.h:139
type errorDomain int32

// errorDomain enumeration from libvirt/virterror.h:139
const (
	fromNone           errorDomain = iota
	fromXen            errorDomain = 1
	fromXend           errorDomain = 2
	fromXenstore       errorDomain = 3
	fromSexpr          errorDomain = 4
	fromXML            errorDomain = 5
	fromDom            errorDomain = 6
	fromRPC            errorDomain = 7
	fromProxy          errorDomain = 8
	fromConf           errorDomain = 9
	fromQemu           errorDomain = 10
	fromNet            errorDomain = 11
	fromTest           errorDomain = 12
	fromRemote         errorDomain = 13
	fromOpenvz         errorDomain = 14
	fromXenxm          errorDomain = 15
	fromStatsLinux     errorDomain = 16
	fromLxc            errorDomain = 17
	fromStorage        errorDomain = 18
	fromNetwork        errorDomain = 19
	fromDomain         errorDomain = 20
	fromUml            errorDomain = 21
	fromNodedev        errorDomain = 22
	fromXenInotify     errorDomain = 23
	fromSecurity       errorDomain = 24
	fromVbox           errorDomain = 25
	fromInterface      errorDomain = 26
	fromOne            errorDomain = 27
	fromEsx            errorDomain = 28
	fromPhyp           errorDomain = 29
	fromSecret         errorDomain = 30
	fromCPU            errorDomain = 31
	fromXenapi         errorDomain = 32
	fromNwfilter       errorDomain = 33
	fromHook           errorDomain = 34
	fromDomainSnapshot errorDomain = 35
	fromAudit          errorDomain = 36
	fromSysinfo        errorDomain = 37
	fromStreams        errorDomain = 38
	fromVmware         errorDomain = 39
	fromEvent          errorDomain = 40
	fromLibxl          errorDomain = 41
	fromLocking        errorDomain = 42
	fromHyperv         errorDomain = 43
	fromCapabilities   errorDomain = 44
	fromURI            errorDomain = 45
	fromAuth           errorDomain = 46
	fromDbus           errorDomain = 47
	fromParallels      errorDomain = 48
	fromDevice         errorDomain = 49
	fromSSH            errorDomain = 50
	fromLockspace      errorDomain = 51
	fromInitctl        errorDomain = 52
	fromIdentity       errorDomain = 53
	fromCgroup         errorDomain = 54
	fromAccess         errorDomain = 55
	fromSystemd        errorDomain = 56
	fromBhyve          errorDomain = 57
	fromCrypto         errorDomain = 58
	fromFirewall       errorDomain = 59
	fromPolkit         errorDomain = 60
	fromThread         errorDomain = 61
	fromAdmin          errorDomain = 62
	fromLogging        errorDomain = 63
	fromXenxl          errorDomain = 64
	fromPerf           errorDomain = 65
	fromLibssh         errorDomain = 66
)

// errorNumber as declared in libvirt/virterror.h:322
type errorNumber int32

// errorNumber enumeration from libvirt/virterror.h:322
const (
	errOk                    errorNumber = iota
	errInternalError         errorNumber = 1
	errNoMemory              errorNumber = 2
	errNoSupport             errorNumber = 3
	errUnknownHost           errorNumber = 4
	errNoConnect             errorNumber = 5
	errInvalidConn           errorNumber = 6
	errInvalidDomain         errorNumber = 7
	errInvalidArg            errorNumber = 8
	errOperationFailed       errorNumber = 9
	errGetFailed             errorNumber = 10
	errPostFailed            errorNumber = 11
	errHTTPError             errorNumber = 12
	errSexprSerial           errorNumber = 13
	errNoXen                 errorNumber = 14
	errXenCall               errorNumber = 15
	errOsType                errorNumber = 16
	errNoKernel              errorNumber = 17
	errNoRoot                errorNumber = 18
	errNoSource              errorNumber = 19
	errNoTarget              errorNumber = 20
	errNoName                errorNumber = 21
	errNoOs                  errorNumber = 22
	errNoDevice              errorNumber = 23
	errNoXenstore            errorNumber = 24
	errDriverFull            errorNumber = 25
	errCallFailed            errorNumber = 26
	errXMLError              errorNumber = 27
	errDomExist              errorNumber = 28
	errOperationDenied       errorNumber = 29
	errOpenFailed            errorNumber = 30
	errReadFailed            errorNumber = 31
	errParseFailed           errorNumber = 32
	errConfSyntax            errorNumber = 33
	errWriteFailed           errorNumber = 34
	errXMLDetail             errorNumber = 35
	errInvalidNetwork        errorNumber = 36
	errNetworkExist          errorNumber = 37
	errSystemError           errorNumber = 38
	errRPC                   errorNumber = 39
	errGnutlsError           errorNumber = 40
	warNoNetwork             errorNumber = 41
	errNoDomain              errorNumber = 42
	errNoNetwork             errorNumber = 43
	errInvalidMac            errorNumber = 44
	errAuthFailed            errorNumber = 45
	errInvalidStoragePool    errorNumber = 46
	errInvalidStorageVol     errorNumber = 47
	warNoStorage             errorNumber = 48
	errNoStoragePool         errorNumber = 49
	errNoStorageVol          errorNumber = 50
	warNoNode                errorNumber = 51
	errInvalidNodeDevice     errorNumber = 52
	errNoNodeDevice          errorNumber = 53
	errNoSecurityModel       errorNumber = 54
	errOperationInvalid      errorNumber = 55
	warNoInterface           errorNumber = 56
	errNoInterface           errorNumber = 57
	errInvalidInterface      errorNumber = 58
	errMultipleInterfaces    errorNumber = 59
	warNoNwfilter            errorNumber = 60
	errInvalidNwfilter       errorNumber = 61
	errNoNwfilter            errorNumber = 62
	errBuildFirewall         errorNumber = 63
	warNoSecret              errorNumber = 64
	errInvalidSecret         errorNumber = 65
	errNoSecret              errorNumber = 66
	errConfigUnsupported     errorNumber = 67
	errOperationTimeout      errorNumber = 68
	errMigratePersistFailed  errorNumber = 69
	errHookScriptFailed      errorNumber = 70
	errInvalidDomainSnapshot errorNumber = 71
	errNoDomainSnapshot      errorNumber = 72
	errInvalidStream         errorNumber = 73
	errArgumentUnsupported   errorNumber = 74
	errStorageProbeFailed    errorNumber = 75
	errStoragePoolBuilt      errorNumber = 76
	errSnapshotRevertRisky   errorNumber = 77
	errOperationAborted      errorNumber = 78
	errAuthCancelled         errorNumber = 79
	errNoDomainMetadata      errorNumber = 80
	errMigrateUnsafe         errorNumber = 81
	errOverflow              errorNumber = 82
	errBlockCopyActive       errorNumber = 83
	errOperationUnsupported  errorNumber = 84
	errSSH                   errorNumber = 85
	errAgentUnresponsive     errorNumber = 86
	errResourceBusy          errorNumber = 87
	errAccessDenied          errorNumber = 88
	errDbusService           errorNumber = 89
	errStorageVolExist       errorNumber = 90
	errCPUIncompatible       errorNumber = 91
	errXMLInvalidSchema      errorNumber = 92
	errMigrateFinishOk       errorNumber = 93
	errAuthUnavailable       errorNumber = 94
	errNoServer              errorNumber = 95
	errNoClient              errorNumber = 96
	errAgentUnsynced         errorNumber = 97
	errLibssh                errorNumber = 98
)

// DomainCheckpointCreateFlags as declared in libvirt/libvirt-domain-checkpoint.h:60
type DomainCheckpointCreateFlags int32

// DomainCheckpointCreateFlags enumeration from libvirt/libvirt-domain-checkpoint.h:60
const (
	DomainCheckpointCreateRedefine DomainCheckpointCreateFlags = 1
	DomainCheckpointCreateQuiesce  DomainCheckpointCreateFlags = 2
)

// DomainCheckpointXMLFlags as declared in libvirt/libvirt-domain-checkpoint.h:73
type DomainCheckpointXMLFlags int32

// DomainCheckpointXMLFlags enumeration from libvirt/libvirt-domain-checkpoint.h:73
const (
	DomainCheckpointXMLSecure   DomainCheckpointXMLFlags = 1
	DomainCheckpointXMLNoDomain DomainCheckpointXMLFlags = 2
	DomainCheckpointXMLSize     DomainCheckpointXMLFlags = 4
)

// DomainCheckpointListFlags as declared in libvirt/libvirt-domain-checkpoint.h:103
type DomainCheckpointListFlags int32

// DomainCheckpointListFlags enumeration from libvirt/libvirt-domain-checkpoint.h:103
const (
	DomainCheckpointListRoots       DomainCheckpointListFlags = 1
	DomainCheckpointListDescendants DomainCheckpointListFlags = 1
	DomainCheckpointListTopological DomainCheckpointListFlags = 2
	DomainCheckpointListLeaves      DomainCheckpointListFlags = 4
	DomainCheckpointListNoLeaves    DomainCheckpointListFlags = 8
)

// DomainCheckpointDeleteFlags as declared in libvirt/libvirt-domain-checkpoint.h:129
type DomainCheckpointDeleteFlags int32

// DomainCheckpointDeleteFlags enumeration from libvirt/libvirt-domain-checkpoint.h:129
const (
	DomainCheckpointDeleteChildren     DomainCheckpointDeleteFlags = 1
	DomainCheckpointDeleteMetadataOnly DomainCheckpointDeleteFlags = 2
	DomainCheckpointDeleteChildrenOnly DomainCheckpointDeleteFlags = 4
)
