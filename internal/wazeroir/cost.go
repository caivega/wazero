package wazeroir

var OperationCost = map[OperationKind]int64{
	OperationKindUnreachable: 1,
	// OperationKindLabel is the kind for OperationLabel.
	OperationKindLabel: 1,
	// OperationKindBr is the kind for OperationBr.
	OperationKindBr: 1,
	// OperationKindBrIf is the kind for OperationBrIf.
	OperationKindBrIf: 1,
	// OperationKindBrTable is the kind for OperationBrTable.
	OperationKindBrTable: 1,
	// OperationKindCall is the kind for OperationCall.
	OperationKindCall: 1,
	// OperationKindCallIndirect is the kind for OperationCallIndirect.
	OperationKindCallIndirect: 1,
	// OperationKindDrop is the kind for OperationDrop.
	OperationKindDrop: 1,
	// OperationKindSelect is the kind for OperationSelect.
	OperationKindSelect: 1,
	// OperationKindPick is the kind for OperationPick.
	OperationKindPick: 1,
	// OperationKindSet is the kind for OperationSet.
	OperationKindSet: 1,
	// OperationKindGlobalGet is the kind for OperationGlobalGet.
	OperationKindGlobalGet: 1,
	// OperationKindGlobalSet is the kind for OperationGlobalSet.
	OperationKindGlobalSet: 1,
	// OperationKindLoad is the kind for OperationLoad.
	OperationKindLoad: 1,
	// OperationKindLoad8 is the kind for OperationLoad8.
	OperationKindLoad8: 1,
	// OperationKindLoad16 is the kind for OperationLoad16.
	OperationKindLoad16: 1,
	// OperationKindLoad32 is the kind for OperationLoad32.
	OperationKindLoad32: 1,
	// OperationKindStore is the kind for OperationStore.
	OperationKindStore: 1,
	// OperationKindStore8 is the kind for OperationStore8.
	OperationKindStore8: 1,
	// OperationKindStore16 is the kind for OperationStore16.
	OperationKindStore16: 1,
	// OperationKindStore32 is the kind for OperationStore32.
	OperationKindStore32: 1,
	// OperationKindMemorySize is the kind for OperationMemorySize.
	OperationKindMemorySize: 1,
	// OperationKindMemoryGrow is the kind for OperationMemoryGrow.
	OperationKindMemoryGrow: 1,
	// OperationKindConstI32 is the kind for OperationConstI32.
	OperationKindConstI32: 1,
	// OperationKindConstI64 is the kind for OperationConstI64.
	OperationKindConstI64: 1,
	// OperationKindConstF32 is the kind for OperationConstF32.
	OperationKindConstF32: 1,
	// OperationKindConstF64 is the kind for OperationConstF64.
	OperationKindConstF64: 1,
	// OperationKindEq is the kind for OperationEq.
	OperationKindEq: 1,
	// OperationKindNe is the kind for OperationNe.
	OperationKindNe: 1,
	// OperationKindEqz is the kind for OperationEqz.
	OperationKindEqz: 1,
	// OperationKindLt is the kind for OperationLt.
	OperationKindLt: 1,
	// OperationKindGt is the kind for OperationGt.
	OperationKindGt: 1,
	// OperationKindLe is the kind for OperationLe.
	OperationKindLe: 1,
	// OperationKindGe is the kind for OperationGe.
	OperationKindGe: 1,
	// OperationKindAdd is the kind for OperationAdd.
	OperationKindAdd: 1,
	// OperationKindSub is the kind for OperationSub.
	OperationKindSub: 1,
	// OperationKindMul is the kind for OperationMul.
	OperationKindMul: 1,
	// OperationKindClz is the kind for OperationClz.
	OperationKindClz: 1,
	// OperationKindCtz is the kind for OperationCtz.
	OperationKindCtz: 1,
	// OperationKindPopcnt is the kind for OperationPopcnt.
	OperationKindPopcnt: 1,
	// OperationKindDiv is the kind for OperationDiv.
	OperationKindDiv: 1,
	// OperationKindRem is the kind for OperationRem.
	OperationKindRem: 1,
	// OperationKindAnd is the kind for OperationAnd.
	OperationKindAnd: 1,
	// OperationKindOr is the kind for OperationOr.
	OperationKindOr: 1,
	// OperationKindXor is the kind for OperationXor.
	OperationKindXor: 1,
	// OperationKindShl is the kind for OperationShl.
	OperationKindShl: 1,
	// OperationKindShr is the kind for OperationShr.
	OperationKindShr: 1,
	// OperationKindRotl is the kind for OperationRotl.
	OperationKindRotl: 1,
	// OperationKindRotr is the kind for OperationRotr.
	OperationKindRotr: 1,
	// OperationKindAbs is the kind for OperationAbs.
	OperationKindAbs: 1,
	// OperationKindNeg is the kind for OperationNeg.
	OperationKindNeg: 1,
	// OperationKindCeil is the kind for OperationCeil.
	OperationKindCeil: 1,
	// OperationKindFloor is the kind for OperationFloor.
	OperationKindFloor: 1,
	// OperationKindTrunc is the kind for OperationTrunc.
	OperationKindTrunc: 1,
	// OperationKindNearest is the kind for OperationNearest.
	OperationKindNearest: 1,
	// OperationKindSqrt is the kind for OperationSqrt.
	OperationKindSqrt: 1,
	// OperationKindMin is the kind for OperationMin.
	OperationKindMin: 1,
	// OperationKindMax is the kind for OperationMax.
	OperationKindMax: 1,
	// OperationKindCopysign is the kind for OperationCopysign.
	OperationKindCopysign: 1,
	// OperationKindI32WrapFromI64 is the kind for OperationI32WrapFromI64.
	OperationKindI32WrapFromI64: 1,
	// OperationKindITruncFromF is the kind for OperationITruncFromF.
	OperationKindITruncFromF: 1,
	// OperationKindFConvertFromI is the kind for OperationFConvertFromI.
	OperationKindFConvertFromI: 1,
	// OperationKindF32DemoteFromF64 is the kind for OperationF32DemoteFromF64.
	OperationKindF32DemoteFromF64: 1,
	// OperationKindF64PromoteFromF32 is the kind for OperationF64PromoteFromF32.
	OperationKindF64PromoteFromF32: 1,
	// OperationKindI32ReinterpretFromF32 is the kind for OperationI32ReinterpretFromF32.
	OperationKindI32ReinterpretFromF32: 1,
	// OperationKindI64ReinterpretFromF64 is the kind for OperationI64ReinterpretFromF64.
	OperationKindI64ReinterpretFromF64: 1,
	// OperationKindF32ReinterpretFromI32 is the kind for OperationF32ReinterpretFromI32.
	OperationKindF32ReinterpretFromI32: 1,
	// OperationKindF64ReinterpretFromI64 is the kind for OperationF64ReinterpretFromI64.
	OperationKindF64ReinterpretFromI64: 1,
	// OperationKindExtend is the kind for OperationExtend.
	OperationKindExtend: 1,
	// OperationKindSignExtend32From8 is the kind for OperationSignExtend32From8.
	OperationKindSignExtend32From8: 1,
	// OperationKindSignExtend32From16 is the kind for OperationSignExtend32From16.
	OperationKindSignExtend32From16: 1,
	// OperationKindSignExtend64From8 is the kind for OperationSignExtend64From8.
	OperationKindSignExtend64From8: 1,
	// OperationKindSignExtend64From16 is the kind for OperationSignExtend64From16.
	OperationKindSignExtend64From16: 1,
	// OperationKindSignExtend64From32 is the kind for OperationSignExtend64From32.
	OperationKindSignExtend64From32: 1,
	// OperationKindMemoryInit is the kind for OperationMemoryInit.
	OperationKindMemoryInit: 1,
	// OperationKindDataDrop is the kind for OperationDataDrop.
	OperationKindDataDrop: 1,
	// OperationKindMemoryCopy is the kind for OperationMemoryCopy.
	OperationKindMemoryCopy: 1,
	// OperationKindMemoryFill is the kind for OperationMemoryFill.
	OperationKindMemoryFill: 1,
	// OperationKindTableInit is the kind for OperationTableInit.
	OperationKindTableInit: 1,
	// OperationKindElemDrop is the kind for OperationElemDrop.
	OperationKindElemDrop: 1,
	// OperationKindTableCopy is the kind for OperationTableCopy.
	OperationKindTableCopy: 1,
	// OperationKindRefFunc is the kind for OperationRefFunc.
	OperationKindRefFunc: 1,
	// OperationKindTableGet is the kind for OperationTableGet.
	OperationKindTableGet: 1,
	// OperationKindTableSet is the kind for OperationTableSet.
	OperationKindTableSet: 1,
	// OperationKindTableSize is the kind for OperationTableSize.
	OperationKindTableSize: 1,
	// OperationKindTableGrow is the kind for OperationTableGrow.
	OperationKindTableGrow: 1,
	// OperationKindTableFill is the kind for OperationTableFill.
	OperationKindTableFill: 1,

	// Vector value related instructions are prefixed by V128.

	// OperationKindV128Const is the kind for OperationV128Const.
	OperationKindV128Const: 1,
	// OperationKindV128Add is the kind for OperationV128Add.
	OperationKindV128Add: 1,
	// OperationKindV128Sub is the kind for OperationV128Sub.
	OperationKindV128Sub: 1,
	// OperationKindV128Load is the kind for OperationV128Load.
	OperationKindV128Load: 1,
	// OperationKindV128LoadLane is the kind for OperationV128LoadLane.
	OperationKindV128LoadLane: 1,
	// OperationKindV128Store is the kind for OperationV128Store.
	OperationKindV128Store: 1,
	// OperationKindV128StoreLane is the kind for OperationV128StoreLane.
	OperationKindV128StoreLane: 1,
	// OperationKindV128ExtractLane is the kind for OperationV128ExtractLane.
	OperationKindV128ExtractLane: 1,
	// OperationKindV128ReplaceLane is the kind for OperationV128ReplaceLane.
	OperationKindV128ReplaceLane: 1,
	// OperationKindV128Splat is the kind for OperationV128Splat.
	OperationKindV128Splat: 1,
	// OperationKindV128Shuffle is the kind for OperationV128Shuffle.
	OperationKindV128Shuffle: 1,
	// OperationKindV128Swizzle is the kind for OperationV128Swizzle.
	OperationKindV128Swizzle: 1,
	// OperationKindV128AnyTrue is the kind for OperationV128AnyTrue.
	OperationKindV128AnyTrue: 1,
	// OperationKindV128AllTrue is the kind for OperationV128AllTrue.
	OperationKindV128AllTrue: 1,
	// OperationKindV128BitMask is the kind for OperationV128BitMask.
	OperationKindV128BitMask: 1,
	// OperationKindV128And is the kind for OperationV128And.
	OperationKindV128And: 1,
	// OperationKindV128Not is the kind for OperationV128Not.
	OperationKindV128Not: 1,
	// OperationKindV128Or is the kind for OperationV128Or.
	OperationKindV128Or: 1,
	// OperationKindV128Xor is the kind for OperationV128Xor.
	OperationKindV128Xor: 1,
	// OperationKindV128Bitselect is the kind for OperationV128Bitselect.
	OperationKindV128Bitselect: 1,
	// OperationKindV128AndNot is the kind for OperationV128AndNot.
	OperationKindV128AndNot: 1,
	// OperationKindV128Shl is the kind for OperationV128Shl.
	OperationKindV128Shl: 1,
	// OperationKindV128Shr is the kind for OperationV128Shr.
	OperationKindV128Shr: 1,
	// OperationKindV128Cmp is the kind for OperationV128Cmp.
	OperationKindV128Cmp: 1,
	// OperationKindV128AddSat is the kind for OperationV128AddSat.
	OperationKindV128AddSat: 1,
	// OperationKindV128SubSat is the kind for OperationV128SubSat.
	OperationKindV128SubSat: 1,
	// OperationKindV128Mul is the kind for OperationV128Mul.
	OperationKindV128Mul: 1,
	// OperationKindV128Div is the kind for OperationV128Div.
	OperationKindV128Div: 1,
	// OperationKindV128Neg is the kind for OperationV128Neg.
	OperationKindV128Neg: 1,
	// OperationKindV128Sqrt is the kind for OperationV128Sqrt.
	OperationKindV128Sqrt: 1,
	// OperationKindV128Abs is the kind for OperationV128Abs.
	OperationKindV128Abs: 1,
	// OperationKindV128Popcnt is the kind for OperationV128Popcnt.
	OperationKindV128Popcnt: 1,
	// OperationKindV128Min is the kind for OperationV128Min.
	OperationKindV128Min: 1,
	// OperationKindV128Max is the kind for OperationV128Max.
	OperationKindV128Max: 1,
	// OperationKindV128AvgrU is the kind for OperationV128AvgrU.
	OperationKindV128AvgrU: 1,
	// OperationKindV128Pmin is the kind for OperationV128Pmin.
	OperationKindV128Pmin: 1,
	// OperationKindV128Pmax is the kind for OperationV128Pmax.
	OperationKindV128Pmax: 1,
	// OperationKindV128Ceil is the kind for OperationV128Ceil.
	OperationKindV128Ceil: 1,
	// OperationKindV128Floor is the kind for OperationV128Floor.
	OperationKindV128Floor: 1,
	// OperationKindV128Trunc is the kind for OperationV128Trunc.
	OperationKindV128Trunc: 1,
	// OperationKindV128Nearest is the kind for OperationV128Nearest.
	OperationKindV128Nearest: 1,
	// OperationKindV128Extend is the kind for OperationV128Extend.
	OperationKindV128Extend: 1,
	// OperationKindV128ExtMul is the kind for OperationV128ExtMul.
	OperationKindV128ExtMul: 1,
	// OperationKindV128Q15mulrSatS is the kind for OperationV128Q15mulrSatS.
	OperationKindV128Q15mulrSatS: 1,
	// OperationKindV128ExtAddPairwise is the kind for OperationV128ExtAddPairwise.
	OperationKindV128ExtAddPairwise: 1,
	// OperationKindV128FloatPromote is the kind for OperationV128FloatPromote.
	OperationKindV128FloatPromote: 1,
	// OperationKindV128FloatDemote is the kind for OperationV128FloatDemote.
	OperationKindV128FloatDemote: 1,
	// OperationKindV128FConvertFromI is the kind for OperationV128FConvertFromI.
	OperationKindV128FConvertFromI: 1,
	// OperationKindV128Dot is the kind for OperationV128Dot.
	OperationKindV128Dot: 1,
	// OperationKindV128Narrow is the kind for OperationV128Narrow.
	OperationKindV128Narrow: 1,
	// OperationKindV128ITruncSatFromF is the kind for OperationV128ITruncSatFromF.
	OperationKindV128ITruncSatFromF: 1,

	// OperationKindBuiltinFunctionCheckExitCode is the kind for OperationBuiltinFunctionCheckExitCode.
	OperationKindBuiltinFunctionCheckExitCode: 1,

	// operationKindEnd is always placed at the bottom of this iota definition to be used in the test.
	operationKindEnd: 1,
}
