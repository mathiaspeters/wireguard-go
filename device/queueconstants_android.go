/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2021 WireGuard LLC. All Rights Reserved.
 */

package device

/* Reduce memory consumption for Android */

const (
	QueueStagedSize            = 128
	QueueOutboundSize          = 8192
	QueueInboundSize           = 8192
	QueueHandshakeSize         = 8192
	MaxSegmentSize             = 2200
	PreallocatedBuffersPerPool = 4096
)
