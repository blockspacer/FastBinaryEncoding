// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: Fbe
// Version: 1.3.0.0

import Foundation

// Fast Binary Encoding log listener
public protocol LogListener: class {

    // Receive log message handler
    func onReceiveLog(message: String)
    // Send log message handler
    func onSendLog(message: String)
}

public extension LogListener {
    func onReceiveLog(message: String) { }
    func onSendLog(message: String) { }
}
