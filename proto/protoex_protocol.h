// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

#pragma once

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "fbe_protocol.h"

#include "protoex_models.h"

#include "proto_protocol.h"

namespace FBE {

namespace protoex {

// Fast Binary Encoding protoex protocol version
struct ProtocolVersion
{
    // Protocol major version
    static const int major = 0;
    // Protocol minor version
    static const int minor = 0;
};

// Fast Binary Encoding protoex sender
class Sender : public virtual FBE::Sender
    , public virtual proto::Sender
{
public:
    Sender()
        : OrderMessageModel(this->_buffer)
        , BalanceMessageModel(this->_buffer)
        , AccountMessageModel(this->_buffer)
    {}
    Sender(const Sender&) = default;
    Sender(Sender&&) noexcept = default;
    virtual ~Sender() = default;

    Sender& operator=(const Sender&) = default;
    Sender& operator=(Sender&&) noexcept = default;

    // Imported senders
    proto::Sender& proto_sender() noexcept { return *this; }

    size_t send(const ::protoex::OrderMessage& value);
    size_t send(const ::protoex::BalanceMessage& value);
    size_t send(const ::protoex::AccountMessage& value);

public:
    // Sender models accessors
    FBE::protoex::OrderMessageModel OrderMessageModel;
    FBE::protoex::BalanceMessageModel BalanceMessageModel;
    FBE::protoex::AccountMessageModel AccountMessageModel;
};

// Fast Binary Encoding protoex receiver
class Receiver : public virtual FBE::Receiver
    , public virtual proto::Receiver
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
    virtual void onReceive(const ::protoex::OrderMessage& value) {}
    virtual void onReceive(const ::protoex::BalanceMessage& value) {}
    virtual void onReceive(const ::protoex::AccountMessage& value) {}

    // Receive message handler
    bool onReceive(size_t type, const void* data, size_t size) override;

private:
    // Receiver values accessors
    ::protoex::OrderMessage OrderMessageValue;
    ::protoex::BalanceMessage BalanceMessageValue;
    ::protoex::AccountMessage AccountMessageValue;

    // Receiver models accessors
    FBE::protoex::OrderMessageModel OrderMessageModel;
    FBE::protoex::BalanceMessageModel BalanceMessageModel;
    FBE::protoex::AccountMessageModel AccountMessageModel;
};

// Fast Binary Encoding protoex proxy
class Proxy : public virtual FBE::Receiver
    , public virtual proto::Proxy
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
    virtual void onProxy(FBE::protoex::OrderMessageModel& model, size_t type, const void* data, size_t size) {}
    virtual void onProxy(FBE::protoex::BalanceMessageModel& model, size_t type, const void* data, size_t size) {}
    virtual void onProxy(FBE::protoex::AccountMessageModel& model, size_t type, const void* data, size_t size) {}

    // Receive message handler
    bool onReceive(size_t type, const void* data, size_t size) override;

private:
    // Proxy models accessors
    FBE::protoex::OrderMessageModel OrderMessageModel;
    FBE::protoex::BalanceMessageModel BalanceMessageModel;
    FBE::protoex::AccountMessageModel AccountMessageModel;
};

// Fast Binary Encoding protoex client
class Client : public virtual Sender, protected virtual Receiver
    , public virtual proto::Client
{
public:
    typedef proto::Client protoClient;

    Client() = default;
    Client(const Client&) = default;
    Client(Client&&) = default;
    virtual ~Client() = default;

    Client& operator=(const Client&) = default;
    Client& operator=(Client&&) = default;

    // Imported clients
    proto::Client& proto_client() noexcept { return *this; }

    // Reset client buffers
    void reset() { std::scoped_lock locker(this->_lock); reset_requests(); }

    // Watchdog for timeouts
    void watchdog(uint64_t utc) { std::scoped_lock locker(this->_lock); watchdog_requests(utc); }

protected:
    virtual bool onReceiveResponse(const ::protoex::OrderMessage& response) { return false; }
    virtual bool onReceiveResponse(const ::protoex::BalanceMessage& response) { return false; }
    virtual bool onReceiveResponse(const ::protoex::AccountMessage& response) { return false; }

    virtual bool onReceiveReject(const ::protoex::OrderMessage& reject) { return false; }
    virtual bool onReceiveReject(const ::protoex::BalanceMessage& reject) { return false; }
    virtual bool onReceiveReject(const ::protoex::AccountMessage& reject) { return false; }

    virtual void onReceiveNotify(const ::protoex::OrderMessage& notify) {}
    virtual void onReceiveNotify(const ::protoex::BalanceMessage& notify) {}
    virtual void onReceiveNotify(const ::protoex::AccountMessage& notify) {}

    virtual void onReceive(const ::protoex::OrderMessage& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }
    virtual void onReceive(const ::protoex::BalanceMessage& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }
    virtual void onReceive(const ::protoex::AccountMessage& value) override { if (!onReceiveResponse(value) && !onReceiveReject(value)) onReceiveNotify(value); }

    // Reset client requests
    virtual void reset_requests();

    // Watchdog client requests for timeouts
    virtual void watchdog_requests(uint64_t utc);
};

} // namespace protoex

} // namespace FBE
