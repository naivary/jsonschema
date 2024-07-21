package definitions

import "github.com/naivary/specraft/schema"

type Manager interface {
    MarkerToApplier(name string, v []any) schema.Applier
}
