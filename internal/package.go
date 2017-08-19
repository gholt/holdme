// gastly can be found at github.com/gholt/gastly

//go:generate gastly orderednumbersdups.go        ../orderedintsdups.go        holdme NumericType=droptype:int Number=Int number=int
//go:generate gastly orderednumbersdups_test.go   ../orderedintsdups_test.go   holdme NumericType=droptype:int Number=Int number=int
//go:generate gastly orderednumbersnodups.go      ../orderedintsnodups.go      holdme NumericType=droptype:int Number=Int number=int
//go:generate gastly orderednumbersnodups_test.go ../orderedintsnodups_test.go holdme NumericType=droptype:int Number=Int number=int

//go:generate gastly orderednumbersdups.go        ../ordereduint32sdups.go        holdme NumericType=droptype:uint32 Number=Uint32 number=uint32
//go:generate gastly orderednumbersdups_test.go   ../ordereduint32sdups_test.go   holdme NumericType=droptype:uint32 Number=Uint32 number=uint32
//go:generate gastly orderednumbersnodups.go      ../ordereduint32snodups.go      holdme NumericType=droptype:uint32 Number=Uint32 number=uint32
//go:generate gastly orderednumbersnodups_test.go ../ordereduint32snodups_test.go holdme NumericType=droptype:uint32 Number=Uint32 number=uint32

//go:generate gastly keysorderedbyvalues.go      ../intkeysorderedbyintvalues.go      holdme NumericKeyType=droptype:int NumericValueType=droptype:int KeysOrderedByValues=IntKeysOrderedByIntValues
//go:generate gastly keysorderedbyvalues_test.go ../intkeysorderedbyintvalues_test.go holdme NumericKeyType=droptype:int NumericValueType=droptype:int KeysOrderedByValues=IntKeysOrderedByIntValues

package internal
