// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.protoex.fbe

// Fast Binary Encoding com.chronoxor.protoex final receiver
@Suppress("MemberVisibilityCanBePrivate", "PrivatePropertyName", "UNUSED_PARAMETER")
open class FinalReceiver : com.chronoxor.fbe.Receiver, FinalReceiverListener
{
    // Imported receivers
    var protoReceiver: com.chronoxor.proto.fbe.FinalReceiver? = null

    // Receiver values accessors
    private val OrderValue: com.chronoxor.protoex.Order
    private val BalanceValue: com.chronoxor.protoex.Balance
    private val AccountValue: com.chronoxor.protoex.Account

    // Receiver models accessors
    private val OrderModel: OrderFinalModel
    private val BalanceModel: BalanceFinalModel
    private val AccountModel: AccountFinalModel

    constructor() : super(true)
    {
        protoReceiver = com.chronoxor.proto.fbe.FinalReceiver(buffer)
        OrderValue = com.chronoxor.protoex.Order()
        OrderModel = OrderFinalModel()
        BalanceValue = com.chronoxor.protoex.Balance()
        BalanceModel = BalanceFinalModel()
        AccountValue = com.chronoxor.protoex.Account()
        AccountModel = AccountFinalModel()
    }

    constructor(buffer: com.chronoxor.fbe.Buffer) : super(buffer, true)
    {
        protoReceiver = com.chronoxor.proto.fbe.FinalReceiver(buffer)
        OrderValue = com.chronoxor.protoex.Order()
        OrderModel = OrderFinalModel()
        BalanceValue = com.chronoxor.protoex.Balance()
        BalanceModel = BalanceFinalModel()
        AccountValue = com.chronoxor.protoex.Account()
        AccountModel = AccountFinalModel()
    }

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        return onReceiveListener(this, type, buffer, offset, size)
    }

    open fun onReceiveListener(listener: FinalReceiverListener, type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            com.chronoxor.protoex.fbe.OrderFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                OrderModel.attach(buffer, offset)
                assert(OrderModel.verify()) { "com.chronoxor.protoex.Order validation failed!" }
                val deserialized = OrderModel.deserialize(OrderValue)
                assert(deserialized > 0) { "com.chronoxor.protoex.Order deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = OrderValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(OrderValue)
                return true
            }
            com.chronoxor.protoex.fbe.BalanceFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                BalanceModel.attach(buffer, offset)
                assert(BalanceModel.verify()) { "com.chronoxor.protoex.Balance validation failed!" }
                val deserialized = BalanceModel.deserialize(BalanceValue)
                assert(deserialized > 0) { "com.chronoxor.protoex.Balance deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = BalanceValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(BalanceValue)
                return true
            }
            com.chronoxor.protoex.fbe.AccountFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                AccountModel.attach(buffer, offset)
                assert(AccountModel.verify()) { "com.chronoxor.protoex.Account validation failed!" }
                val deserialized = AccountModel.deserialize(AccountValue)
                assert(deserialized > 0) { "com.chronoxor.protoex.Account deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = AccountValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                listener.onReceive(AccountValue)
                return true
            }
        }

        if ((protoReceiver != null) && protoReceiver!!.onReceiveListener(listener, type, buffer, offset, size))
            return true

        return false
    }
}