// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.protoex.fbe

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import com.chronoxor.fbe.*
import com.chronoxor.protoex.*

import com.google.gson.*

class StateExJson : JsonSerializer<StateEx>, JsonDeserializer<StateEx>
{

    @Override
    override fun serialize(src: StateEx, typeOfSrc: Type, context: JsonSerializationContext): JsonElement
    {
        return JsonPrimitive(src.raw)
    }

    @Throws(JsonParseException::class)
    override fun deserialize(json: JsonElement, type: Type, context: JsonDeserializationContext):StateEx
    {
        return StateEx(json.asJsonPrimitive.asByte)
    }
}