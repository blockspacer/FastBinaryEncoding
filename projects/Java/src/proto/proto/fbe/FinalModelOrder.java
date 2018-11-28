// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.1.0.0

package proto.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.ByteBuffer;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import fbe.*;
import proto.*;

// Fast Binary Encoding Order final model class
public final class FinalModelOrder extends FinalModel
{
    public final FinalModelInt32 uid;
    public final FinalModelString symbol;
    public final FinalModelOrderSide side;
    public final FinalModelOrderType type;
    public final FinalModelDouble price;
    public final FinalModelDouble volume;

    public FinalModelOrder(Buffer buffer, long offset)
    {
        super(buffer, offset);
        uid = new FinalModelInt32(buffer, 0);
        symbol = new FinalModelString(buffer, 0);
        side = new FinalModelOrderSide(buffer, 0);
        type = new FinalModelOrderType(buffer, 0);
        price = new FinalModelDouble(buffer, 0);
        volume = new FinalModelDouble(buffer, 0);
    }

    // Get the allocation size
    public long fbeAllocationSize(Order fbeValue)
    {
        long fbeResult = 0
            + uid.fbeAllocationSize(fbeValue.uid)
            + symbol.fbeAllocationSize(fbeValue.symbol)
            + side.fbeAllocationSize(fbeValue.side)
            + type.fbeAllocationSize(fbeValue.type)
            + price.fbeAllocationSize(fbeValue.price)
            + volume.fbeAllocationSize(fbeValue.volume)
            ;
        return fbeResult;
    }

    // Get the field type
    public static final long fbeTypeConst = 1;
    public long fbeType() { return fbeTypeConst; }

    // Check if the struct value is valid
    @Override
    public long verify()
    {
        _buffer.shift(fbeOffset());
        long fbeResult = verifyFields();
        _buffer.unshift(fbeOffset());
        return fbeResult;
    }

    // Check if the struct fields are valid
    public long verifyFields()
    {
        long fbeCurrentOffset = 0;
        long fbeFieldSize = 0;

        uid.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = uid.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        symbol.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = symbol.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        side.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = side.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        type.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = type.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        price.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = price.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        volume.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = volume.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        return fbeCurrentOffset;
    }

    // Get the struct value
    public Order get(Size fbeSize) { return get(fbeSize, new Order()); }
    public Order get(Size fbeSize, Order fbeValue)
    {
        _buffer.shift(fbeOffset());
        fbeSize.value = getFields(fbeValue);
        _buffer.unshift(fbeOffset());
        return fbeValue;
    }

    // Get the struct fields values
    public long getFields(Order fbeValue)
    {
        long fbeCurrentOffset = 0;
        long fbeCurrentSize = 0;
        var fbeFieldSize = new Size(0);

        uid.fbeOffset(fbeCurrentOffset);
        fbeValue.uid = uid.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        symbol.fbeOffset(fbeCurrentOffset);
        fbeValue.symbol = symbol.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        side.fbeOffset(fbeCurrentOffset);
        fbeValue.side = side.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        type.fbeOffset(fbeCurrentOffset);
        fbeValue.type = type.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        price.fbeOffset(fbeCurrentOffset);
        fbeValue.price = price.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        volume.fbeOffset(fbeCurrentOffset);
        fbeValue.volume = volume.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        return fbeCurrentSize;
    }

    // Set the struct value
    public long set(Order fbeValue)
    {
        _buffer.shift(fbeOffset());
        long fbeSize = setFields(fbeValue);
        _buffer.unshift(fbeOffset());
        return fbeSize;
    }

    // Set the struct fields values
    public long setFields(Order fbeValue)
    {
        long fbeCurrentOffset = 0;
        long fbeCurrentSize = 0;
        var fbeFieldSize = new Size(0);

        uid.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = uid.set(fbeValue.uid);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        symbol.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = symbol.set(fbeValue.symbol);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        side.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = side.set(fbeValue.side);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        type.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = type.set(fbeValue.type);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        price.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = price.set(fbeValue.price);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        volume.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = volume.set(fbeValue.volume);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        return fbeCurrentSize;
    }
}