// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

import ChronoxorFbe

// Fast Binary Encoding StructOptional final model
public class StructOptionalFinalModel: Model {
    private let _model: FinalModelStructOptional

    public override init(buffer: Buffer = Buffer()) {
        _model = FinalModelStructOptional(buffer: buffer, offset: 8)

        super.init(buffer: buffer)
    }

    // Model type
    public var fbeType: Int = fbeTypeConst

    static let fbeTypeConst: Int = FinalModelStructOptional.fbeTypeConst

    // Check if the struct value is valid
    public func verify() -> Bool {
        if ((buffer.offset + _model.fbeOffset) > buffer.size) {
            return false
        }

        let fbeStructSize = Int(readUInt32(offset: _model.fbeOffset - 8))
        let fbeStructType = Int(readUInt32(offset: _model.fbeOffset - 4))
        if ((fbeStructSize <= 0) || (fbeStructType != fbeType)) {
            return false
        }

        return ((8 + _model.verify()) == fbeStructSize)
    }

    // Serialize the struct value
    public func serialize(value: StructOptional) throws -> Int {
        let fbeInitialSize = buffer.size

        let fbeStructType = fbeType
        var fbeStructSize = 8 + _model.fbeAllocationSize(value: value)
        let fbeStructOffset = try buffer.allocate(size: fbeStructSize) - buffer.offset
        if ((buffer.offset + fbeStructOffset + fbeStructSize) > buffer.size) {
            assertionFailure("Model is broken!")
            return 0
        }

        fbeStructSize = try _model.set(value: value) + 8
        try buffer.resize(size: fbeInitialSize + fbeStructSize)

        write(offset: _model.fbeOffset - 8, value: UInt32(fbeStructSize))
        write(offset: _model.fbeOffset - 4, value: UInt32(fbeStructType))

        return fbeStructSize
    }

    // Deserialize the struct value
    public func deserialize() -> StructOptional { var value = StructOptional(); _ = deserialize(value: &value); return value }

    public func deserialize(value: inout StructOptional) -> Int {
        var valueRef = value

        if ((buffer.offset + _model.fbeOffset) > buffer.size) {
            assertionFailure("Model is broken!")
            return 0
        }

        let fbeStructSize = Int32(readUInt32(offset: _model.fbeOffset - 8))
        let fbeStructType = Int32(readUInt32(offset: _model.fbeOffset - 4))
        if (fbeStructSize <= 0) || (fbeStructType != fbeType) {
            assertionFailure("Model is broken!")
            return 8
        }

        var fbeSize = Size()
        valueRef = _model.get(size: &fbeSize, fbeValue: &valueRef)
        return 8 + fbeSize.value
    }

    // Move to the next struct value
    public func next(prev: Int) {
        _model.fbeShift(size: prev)
    }
}
