// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"testing"

	"mattermost-server/store/storetest"
)

func TestSessionStore(t *testing.T) {
	StoreTest(t, storetest.TestSessionStore)
}
