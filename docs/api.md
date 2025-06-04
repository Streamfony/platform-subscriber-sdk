# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [subscriber.proto](#subscriber-proto)
    - [HeartbeatRequest](#platform_subscriber-HeartbeatRequest)
    - [HeartbeatResponse](#platform_subscriber-HeartbeatResponse)
    - [MarkDisconnectedRequest](#platform_subscriber-MarkDisconnectedRequest)
    - [MarkDisconnectedResponse](#platform_subscriber-MarkDisconnectedResponse)
    - [PlatformInfo](#platform_subscriber-PlatformInfo)
    - [Target](#platform_subscriber-Target)
    - [Task](#platform_subscriber-Task)
  
    - [TaskAction](#platform_subscriber-TaskAction)
  
    - [PlatformSubscriber](#platform_subscriber-PlatformSubscriber)
  
- [Scalar Value Types](#scalar-value-types)



<a name="subscriber-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## subscriber.proto



<a name="platform_subscriber-HeartbeatRequest"></a>

### HeartbeatRequest
region Requests


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| platform_info | [PlatformInfo](#platform_subscriber-PlatformInfo) |  |  |
| amount_of_connections | [int64](#int64) |  |  |






<a name="platform_subscriber-HeartbeatResponse"></a>

### HeartbeatResponse
region Responses


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_ok | [bool](#bool) |  |  |
| tasks | [Task](#platform_subscriber-Task) | repeated |  |






<a name="platform_subscriber-MarkDisconnectedRequest"></a>

### MarkDisconnectedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| platform_info | [PlatformInfo](#platform_subscriber-PlatformInfo) |  |  |
| target | [Target](#platform_subscriber-Target) |  |  |






<a name="platform_subscriber-MarkDisconnectedResponse"></a>

### MarkDisconnectedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_ok | [bool](#bool) |  |  |






<a name="platform_subscriber-PlatformInfo"></a>

### PlatformInfo
region Parts


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| platform | [string](#string) |  |  |
| type | [string](#string) |  |  |
| instance_id | [string](#string) |  |  |






<a name="platform_subscriber-Target"></a>

### Target



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  |  |
| additional_info | [string](#string) | optional |  |






<a name="platform_subscriber-Task"></a>

### Task



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [TaskAction](#platform_subscriber-TaskAction) |  |  |
| target | [Target](#platform_subscriber-Target) |  |  |





 


<a name="platform_subscriber-TaskAction"></a>

### TaskAction
region Enums

| Name | Number | Description |
| ---- | ------ | ----------- |
| SUBSCRIBE | 0 |  |
| UNSUBSCRIBE | 1 |  |


 

 


<a name="platform_subscriber-PlatformSubscriber"></a>

### PlatformSubscriber


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Heartbeat | [HeartbeatRequest](#platform_subscriber-HeartbeatRequest) | [HeartbeatResponse](#platform_subscriber-HeartbeatResponse) |  |
| MarkDisconnected | [MarkDisconnectedRequest](#platform_subscriber-MarkDisconnectedRequest) | [MarkDisconnectedResponse](#platform_subscriber-MarkDisconnectedResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

