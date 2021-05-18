// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package interfaces

import (
	"bitbucket.org/enesyteam/papo-server/model"
)

type IndexerJobInterface interface {
	MakeWorker() model.Worker
}
