package pbmerge

import (
	"fmt"

	"github.com/WiggidyW/etco-go/error/configerror"
)

func newError(
	typeId int32,
	typeMapKey string,
	errStr string,
) configerror.ErrInvalid {
	return configerror.ErrInvalid{
		Err: configerror.ErrShopTypeInvalid{
			Err: fmt.Errorf(
				"'%d - %s': %s",
				typeId,
				typeMapKey,
				errStr,
			),
		},
	}
}