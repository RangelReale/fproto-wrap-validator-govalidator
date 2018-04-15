# fproto-wrap-validator-stddata

[fproto-wrap-validator](https://github.com/RangelReale/fproto-wrap-validator) validator for standard data types.

### example

```protobuf
syntax = "proto3";
package gw_sample;
option go_package = "gwsample/core";

import "github.com/RangelReale/fproto-wrap-validator-stddata/datavalidator.proto";

message User {
    // Checks if the field is NOT an email
    string name = 1 [(datavalidator.field) = {email: false}];
    // Checks if the field is an email
    string email = 2 [(datavalidator.field) = {email: true}];
}
```

### field validators

* (datavalidator.field).ascii [bool]
* (datavalidator.field).alpha [bool]
* (datavalidator.field).alphanumeric [bool]
* (datavalidator.field).base64 [bool]
* (datavalidator.field).cidr [bool]
* (datavalidator.field).creditcard [bool]
* (datavalidator.field).dnsname [bool]
* (datavalidator.field).datauri [bool]
* (datavalidator.field).dialstring [bool]
* (datavalidator.field).email [bool]
* (datavalidator.field).hexadecimal [bool]
* (datavalidator.field).hexcolor [bool]
* (datavalidator.field).host [bool]
* (datavalidator.field).ip [bool]
* (datavalidator.field).ipv4 [bool]
* (datavalidator.field).ipv6 [bool]
* (datavalidator.field).isbn10 [bool]
* (datavalidator.field).isbn13 [bool]
* (datavalidator.field).iso3166alpha2 [bool]
* (datavalidator.field).iso3166alpha3 [bool]
* (datavalidator.field).iso693alpha2 [bool]
* (datavalidator.field).iso693alpha3b [bool]
* (datavalidator.field).iso4217 [bool]
* (datavalidator.field).json [bool]
* (datavalidator.field).latitude [bool]
* (datavalidator.field).longitude [bool]
* (datavalidator.field).lowercase [bool]
* (datavalidator.field).mac [bool]
* (datavalidator.field).mongoid [bool]
* (datavalidator.field).multibyte [bool]
* (datavalidator.field).null [bool]
* (datavalidator.field).numeric [bool]
* (datavalidator.field).port [bool]
* (datavalidator.field).printableascii [bool]
* (datavalidator.field).rfc3339 [bool]
* (datavalidator.field).rfc3339withoutzone [bool]
* (datavalidator.field).rgbcolor [bool]
* (datavalidator.field).requesturi [bool]
* (datavalidator.field).requesturl [bool]
* (datavalidator.field).ssn [bool]
* (datavalidator.field).semver [bool]
* (datavalidator.field).url [bool]
* (datavalidator.field).utfdugit [bool]
* (datavalidator.field).utfletter [bool]
* (datavalidator.field).utfletternumeric [bool]
* (datavalidator.field).utfnumeric [bool]
* (datavalidator.field).uuid [bool]
* (datavalidator.field).uuidv3 [bool]
* (datavalidator.field).uuidv4 [bool]
* (datavalidator.field).uuidv5 [bool]
* (datavalidator.field).uppercase [bool]
* (datavalidator.field).variablewidth [bool]
    
### author

Rangel Reale (rangelspam@gmail.com)

