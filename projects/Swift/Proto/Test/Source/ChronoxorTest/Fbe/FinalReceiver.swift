// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

import Foundation
import ChronoxorFbe
import ChronoxorProto

// Fast Binary Encoding ChronoxorTest final receiver
open class FinalReceiver: ChronoxorFbe.ReceiverProtocol {
    // Imported receivers
    let ProtoReceiver: ChronoxorProto.FinalReceiver?

    // Receiver values accessors

    // Receiver models accessors

    public var buffer: Buffer = Buffer()
    public var final: Bool = false

    public init() {
        ProtoReceiver = ChronoxorProto.FinalReceiver(buffer: buffer)
        build(final: true)
    }

    public init(buffer: ChronoxorFbe.Buffer) {
        ProtoReceiver = ChronoxorProto.FinalReceiver(buffer: buffer)
        build(with: buffer, final: true)
    }

    open func onReceive(type: Int, buffer: Data, offset: Int, size: Int) -> Bool {
        guard let listener = self as? FinalReceiverListener else { return false }
        return onReceiveListener(listener: listener, type: type, buffer: buffer, offset: offset, size: size)
    }

    open func onReceiveListener(listener: FinalReceiverListener, type: Int, buffer: Data, offset: Int, size: Int) -> Bool {

        if let ProtoReceiver = ProtoReceiver, ProtoReceiver.onReceiveListener(listener: listener, type: type, buffer: buffer, offset: offset, size: size) {
            return true
        }

        return false
    }
}
