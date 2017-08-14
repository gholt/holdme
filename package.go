// got can be found at github.com/gholt/got

//go:generate got orderedslicedups.got          zzz_ordereduint32dups.go        T=Uint32 t=uint32
//go:generate got orderedslicedups_test.got     zzz_ordereduint32dups_test.go   T=Uint32 t=uint32
//go:generate got orderedslicenodups.got        zzz_ordereduint32nodups.go      T=Uint32 t=uint32
//go:generate got orderedslicenodups_test.got   zzz_ordereduint32nodups_test.go T=Uint32 t=uint32

//go:generate got orderedslicedups.got          zzz_orderedintdups.go           T=Int t=int
//go:generate got orderedslicedups_test.got     zzz_orderedintdups_test.go      T=Int t=int
//go:generate got orderedslicenodups.got        zzz_orderedintnodups.go         T=Int t=int
//go:generate got orderedslicenodups_test.got   zzz_orderedintnodups_test.go    T=Int t=int

//go:generate got valueorderedkeys.got          zzz_intvalueorderedintkeys.go       p=holdme new=NewIntValueOrderedIntKeys T=IntValueOrderedIntKeys K=OrderedKeys k=int V=KeyToValue v=int
//go:generate got valueorderedkeys_test.got     zzz_intvalueorderedintkeys_test.go  p=holdme new=NewIntValueOrderedIntKeys T=IntValueOrderedIntKeys K=OrderedKeys k=int V=KeyToValue v=int

// Package holdme contains some generic structures for holding values. I plan
// to add to this over time as I need things.
package holdme
