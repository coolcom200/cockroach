// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package opgen

import (
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scop"
	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/scpb"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/catid"
)

func init() {
	opRegistry.register((*scpb.TableLocalitySecondaryRegion)(nil),
		toPublic(
			scpb.Status_ABSENT,
			to(scpb.Status_PUBLIC,
				emit(func(this *scpb.TableLocalitySecondaryRegion) *scop.NotImplemented {
					return notImplemented(this)
				}),
			),
		),
		toAbsent(
			scpb.Status_PUBLIC,
			to(scpb.Status_ABSENT,
				// TODO(postamar): remove revertibility constraint when possible
				revertible(false),
				// TODO(postamar): implement table locality update
				emit(func(this *scpb.TableLocalitySecondaryRegion) *scop.RemoveBackReferenceInTypes {
					return &scop.RemoveBackReferenceInTypes{
						TypeIDs:              []catid.DescID{this.RegionEnumTypeID},
						BackReferencedDescID: this.TableID,
					}
				}),
			),
		),
	)
}
