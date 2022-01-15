# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/review-service/review-service.proto](#npool/review-service/review-service.proto)
    - [CreateReviewRequest](#review.service.v1.CreateReviewRequest)
    - [CreateReviewResponse](#review.service.v1.CreateReviewResponse)
    - [CreateReviewRuleRequest](#review.service.v1.CreateReviewRuleRequest)
    - [CreateReviewRuleResponse](#review.service.v1.CreateReviewRuleResponse)
    - [GetReviewRuleByDomainObjectTypeRequest](#review.service.v1.GetReviewRuleByDomainObjectTypeRequest)
    - [GetReviewRuleByDomainObjectTypeResponse](#review.service.v1.GetReviewRuleByDomainObjectTypeResponse)
    - [GetReviewRuleRequest](#review.service.v1.GetReviewRuleRequest)
    - [GetReviewRuleResponse](#review.service.v1.GetReviewRuleResponse)
    - [GetReviewRulesByDomainRequest](#review.service.v1.GetReviewRulesByDomainRequest)
    - [GetReviewRulesByDomainResponse](#review.service.v1.GetReviewRulesByDomainResponse)
    - [GetReviewsByAppDomainObjectTypeIDRequest](#review.service.v1.GetReviewsByAppDomainObjectTypeIDRequest)
    - [GetReviewsByAppDomainObjectTypeIDResponse](#review.service.v1.GetReviewsByAppDomainObjectTypeIDResponse)
    - [GetReviewsByAppDomainRequest](#review.service.v1.GetReviewsByAppDomainRequest)
    - [GetReviewsByAppDomainResponse](#review.service.v1.GetReviewsByAppDomainResponse)
    - [GetReviewsByDomainRequest](#review.service.v1.GetReviewsByDomainRequest)
    - [GetReviewsByDomainResponse](#review.service.v1.GetReviewsByDomainResponse)
    - [Review](#review.service.v1.Review)
    - [ReviewRule](#review.service.v1.ReviewRule)
    - [SubmitReviewRequest](#review.service.v1.SubmitReviewRequest)
    - [SubmitReviewResponse](#review.service.v1.SubmitReviewResponse)
    - [SubmitReviewResultRequest](#review.service.v1.SubmitReviewResultRequest)
    - [SubmitReviewResultResponse](#review.service.v1.SubmitReviewResultResponse)
    - [UpdateReviewRequest](#review.service.v1.UpdateReviewRequest)
    - [UpdateReviewResponse](#review.service.v1.UpdateReviewResponse)
    - [UpdateReviewRuleRequest](#review.service.v1.UpdateReviewRuleRequest)
    - [UpdateReviewRuleResponse](#review.service.v1.UpdateReviewRuleResponse)
    - [VersionResponse](#review.service.v1.VersionResponse)
  
    - [ReviewService](#review.service.v1.ReviewService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/review-service/review-service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/review-service/review-service.proto



<a name="review.service.v1.CreateReviewRequest"></a>

### CreateReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.CreateReviewResponse"></a>

### CreateReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.CreateReviewRuleRequest"></a>

### CreateReviewRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReviewRule](#review.service.v1.ReviewRule) |  |  |






<a name="review.service.v1.CreateReviewRuleResponse"></a>

### CreateReviewRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReviewRule](#review.service.v1.ReviewRule) |  |  |






<a name="review.service.v1.GetReviewRuleByDomainObjectTypeRequest"></a>

### GetReviewRuleByDomainObjectTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Domain | [string](#string) |  |  |
| ObjectType | [string](#string) |  |  |






<a name="review.service.v1.GetReviewRuleByDomainObjectTypeResponse"></a>

### GetReviewRuleByDomainObjectTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReviewRule](#review.service.v1.ReviewRule) |  |  |






<a name="review.service.v1.GetReviewRuleRequest"></a>

### GetReviewRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="review.service.v1.GetReviewRuleResponse"></a>

### GetReviewRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReviewRule](#review.service.v1.ReviewRule) |  |  |






<a name="review.service.v1.GetReviewRulesByDomainRequest"></a>

### GetReviewRulesByDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Domain | [string](#string) |  |  |






<a name="review.service.v1.GetReviewRulesByDomainResponse"></a>

### GetReviewRulesByDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ReviewRule](#review.service.v1.ReviewRule) | repeated |  |






<a name="review.service.v1.GetReviewsByAppDomainObjectTypeIDRequest"></a>

### GetReviewsByAppDomainObjectTypeIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Domain | [string](#string) |  |  |
| ObjectType | [string](#string) |  |  |
| ObjectID | [string](#string) |  |  |






<a name="review.service.v1.GetReviewsByAppDomainObjectTypeIDResponse"></a>

### GetReviewsByAppDomainObjectTypeIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Review](#review.service.v1.Review) | repeated |  |






<a name="review.service.v1.GetReviewsByAppDomainRequest"></a>

### GetReviewsByAppDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Domain | [string](#string) |  |  |






<a name="review.service.v1.GetReviewsByAppDomainResponse"></a>

### GetReviewsByAppDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Review](#review.service.v1.Review) | repeated |  |






<a name="review.service.v1.GetReviewsByDomainRequest"></a>

### GetReviewsByDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Domain | [string](#string) |  |  |






<a name="review.service.v1.GetReviewsByDomainResponse"></a>

### GetReviewsByDomainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Review](#review.service.v1.Review) | repeated |  |






<a name="review.service.v1.Review"></a>

### Review



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| ObjectType | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ReviewerID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| ObjectID | [string](#string) |  |  |
| Domain | [string](#string) |  |  |






<a name="review.service.v1.ReviewRule"></a>

### ReviewRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| ObjectType | [string](#string) |  |  |
| Domain | [string](#string) |  |  |
| Rules | [string](#string) |  |  |






<a name="review.service.v1.SubmitReviewRequest"></a>

### SubmitReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.SubmitReviewResponse"></a>

### SubmitReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.SubmitReviewResultRequest"></a>

### SubmitReviewResultRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.SubmitReviewResultResponse"></a>

### SubmitReviewResultResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.UpdateReviewRequest"></a>

### UpdateReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.UpdateReviewResponse"></a>

### UpdateReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Review](#review.service.v1.Review) |  |  |






<a name="review.service.v1.UpdateReviewRuleRequest"></a>

### UpdateReviewRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReviewRule](#review.service.v1.ReviewRule) |  |  |






<a name="review.service.v1.UpdateReviewRuleResponse"></a>

### UpdateReviewRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReviewRule](#review.service.v1.ReviewRule) |  |  |






<a name="review.service.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="review.service.v1.ReviewService"></a>

### ReviewService
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#review.service.v1.VersionResponse) | Method Version |
| CreateReview | [CreateReviewRequest](#review.service.v1.CreateReviewRequest) | [CreateReviewResponse](#review.service.v1.CreateReviewResponse) |  |
| UpdateReview | [UpdateReviewRequest](#review.service.v1.UpdateReviewRequest) | [UpdateReviewResponse](#review.service.v1.UpdateReviewResponse) |  |
| GetReviewsByDomain | [GetReviewsByDomainRequest](#review.service.v1.GetReviewsByDomainRequest) | [GetReviewsByDomainResponse](#review.service.v1.GetReviewsByDomainResponse) |  |
| GetReviewsByAppDomain | [GetReviewsByAppDomainRequest](#review.service.v1.GetReviewsByAppDomainRequest) | [GetReviewsByAppDomainResponse](#review.service.v1.GetReviewsByAppDomainResponse) |  |
| GetReviewsByAppDomainObjectTypeID | [GetReviewsByAppDomainObjectTypeIDRequest](#review.service.v1.GetReviewsByAppDomainObjectTypeIDRequest) | [GetReviewsByAppDomainObjectTypeIDResponse](#review.service.v1.GetReviewsByAppDomainObjectTypeIDResponse) |  |
| SubmitReview | [SubmitReviewRequest](#review.service.v1.SubmitReviewRequest) | [SubmitReviewResponse](#review.service.v1.SubmitReviewResponse) |  |
| SubmitReviewResult | [SubmitReviewResultRequest](#review.service.v1.SubmitReviewResultRequest) | [SubmitReviewResultResponse](#review.service.v1.SubmitReviewResultResponse) |  |
| CreateReviewRule | [CreateReviewRuleRequest](#review.service.v1.CreateReviewRuleRequest) | [CreateReviewRuleResponse](#review.service.v1.CreateReviewRuleResponse) |  |
| UpdateReviewRule | [UpdateReviewRuleRequest](#review.service.v1.UpdateReviewRuleRequest) | [UpdateReviewRuleResponse](#review.service.v1.UpdateReviewRuleResponse) |  |
| GetReviewRule | [GetReviewRuleRequest](#review.service.v1.GetReviewRuleRequest) | [GetReviewRuleResponse](#review.service.v1.GetReviewRuleResponse) |  |
| GetReviewRulesByDomain | [GetReviewRulesByDomainRequest](#review.service.v1.GetReviewRulesByDomainRequest) | [GetReviewRulesByDomainResponse](#review.service.v1.GetReviewRulesByDomainResponse) |  |
| GetReviewRuleByDomainObjectType | [GetReviewRuleByDomainObjectTypeRequest](#review.service.v1.GetReviewRuleByDomainObjectTypeRequest) | [GetReviewRuleByDomainObjectTypeResponse](#review.service.v1.GetReviewRuleByDomainObjectTypeResponse) |  |

 



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

