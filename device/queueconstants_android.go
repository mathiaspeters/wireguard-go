/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2021 WireGuard LLC. All Rights Reserved.
 */

package device

/* Reduce memory consumption for Android */

const (
	QueueStagedSize            = 128
	QueueOutboundSize          = 2048
	QueueInboundSize           = 2048
	QueueHandshakeSize         = 2048
	MaxSegmentSize             = 2200
	PreallocatedBuffersPerPool = 4096
)
