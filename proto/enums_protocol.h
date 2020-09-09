// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.4.0.0

#pragma once

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "fbe_protocol.h"

#include "enums_models.h"

#if defined(__GNUC__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wunused-variable"
#elif defined(_MSC_VER)
#pragma warning(push)
#pragma warning(disable:4065) // C4065: switch statement contains 'default' but no 'case' labels
#pragma warning(disable:4100) // C4100: 'identifier' : unreferenced formal parameter
#pragma warning(disable:4702) // C4702: unreachable code
#endif

namespace FBE {

namespace enums {

// Fast Binary Encoding enums protocol version
struct ProtocolVersion
{
    // Protocol major version
    static const int major = 0;
    // Protocol minor version
    static const int minor = 0;
};

// Fast Binary Encoding enums sender
class Sender : public virtual FBE::Sender
{
public:
    Sender()
    {}
    Sender(const Sender&) = default;
    Sender(Sender&&) noexcept = default;
    virtual ~Sender() = default;

    Sender& operator=(const Sender&) = default;
    Sender& operator=(Sender&&) noexcept = default;

public:
    // Sender models accessors
};

// Fast Binary Encoding enums receiver
class Receiver : public virtual FBE::Receiver
{
public:
    Receiver() {}
    Receiver(const Receiver&) = default;
    Receiver(Receiver&&) = default;
    virtual ~Receiver() = default;

    Receiver& operator=(const Receiver&) = default;
    Receiver& operator=(Receiver&&) = default;

protected:
    // Receive handlers

    // Receive message handler
    bool onReceive(size_t type, const void* data, size_t size) override;

private:
    // Receiver values accessors

    // Receiver models accessors
};

// Fast Binary Encoding enums proxy
class Proxy : public virtual FBE::Receiver
{
public:
    Proxy() {}
    Proxy(const Proxy&) = default;
    Proxy(Proxy&&) = default;
    virtual ~Proxy() = default;

    Proxy& operator=(const Proxy&) = default;
    Proxy& operator=(Proxy&&) = default;

protected:
    // Proxy handlers

    // Receive message handler
    bool onReceive(size_t type, const void* data, size_t size) override;

private:
    // Proxy models accessors
};

// Fast Binary Encoding enums client
class Client : public virtual Sender, protected virtual Receiver
{
public:
    Client() = default;
    Client(const Client&) = default;
    Client(Client&&) = default;
    virtual ~Client() = default;

    Client& operator=(const Client&) = default;
    Client& operator=(Client&&) = default;

    // Reset client buffers
    void reset() { std::scoped_lock locker(this->_lock); reset_requests(); }

    // Watchdog for timeouts
    void watchdog(uint64_t utc) { std::scoped_lock locker(this->_lock); watchdog_requests(utc); }

protected:
    std::mutex _lock;
    uint64_t _timestamp{0};

    // Reset client requests
    virtual void reset_requests();

    // Watchdog client requests for timeouts
    virtual void watchdog_requests(uint64_t utc);
};

} // namespace enums

} // namespace FBE

#if defined(__GNUC__)
#pragma GCC diagnostic pop
#elif defined(_MSC_VER)
#pragma warning(pop)
#endif
