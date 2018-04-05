# fproto-wrap-validator-govalidator

[fproto-wrap-validator](https://github.com/RangelReale/fproto-wrap-validator) validator using [govalidator](https://github.com/asaskevich/govalidator).

### example

```protobuf
syntax = "proto3";
package gw_sample;
option go_package = "gwsample/core";

import "github.com/RangelReale/fproto-wrap-validator-govalidator/govalidator.proto";

message User {
    // Checks if the field is NOT an email
    string name = 1 [(govalidator.field) = {email: false}];
    // Checks if the field is an email
    string email = 2 [(govalidator.field) = {email: true}];
}
```

### field validators

* (govalidator.field).ascii [bool]
* (govalidator.field).alpha [bool]
* (govalidator.field).alphanumeric [bool]
* (govalidator.field).base64 [bool]
* (govalidator.field).cidr [bool]
* (govalidator.field).creditcard [bool]
* (govalidator.field).dnsname [bool]
* (govalidator.field).datauri [bool]
* (govalidator.field).dialstring [bool]
* (govalidator.field).email [bool]
* (govalidator.field).hexadecimal [bool]
* (govalidator.field).hexcolor [bool]
* (govalidator.field).host [bool]
* (govalidator.field).ip [bool]
* (govalidator.field).ipv4 [bool]
* (govalidator.field).ipv6 [bool]
* (govalidator.field).isbn10 [bool]
* (govalidator.field).isbn13 [bool]
* (govalidator.field).iso3166alpha2 [bool]
* (govalidator.field).iso3166alpha3 [bool]
* (govalidator.field).iso693alpha2 [bool]
* (govalidator.field).iso693alpha3b [bool]
* (govalidator.field).iso4217 [bool]
* (govalidator.field).json [bool]
* (govalidator.field).latitude [bool]
* (govalidator.field).longitude [bool]
* (govalidator.field).lowercase [bool]
* (govalidator.field).mac [bool]
* (govalidator.field).mongoid [bool]
* (govalidator.field).multibyte [bool]
* (govalidator.field).null [bool]
* (govalidator.field).numeric [bool]
* (govalidator.field).port [bool]
* (govalidator.field).printableascii [bool]
* (govalidator.field).rfc3339 [bool]
* (govalidator.field).rfc3339withoutzone [bool]
* (govalidator.field).rgbcolor [bool]
* (govalidator.field).requesturi [bool]
* (govalidator.field).requesturl [bool]
* (govalidator.field).ssn [bool]
* (govalidator.field).semver [bool]
* (govalidator.field).url [bool]
* (govalidator.field).utfdugit [bool]
* (govalidator.field).utfletter [bool]
* (govalidator.field).utfletternumeric [bool]
* (govalidator.field).utfnumeric [bool]
* (govalidator.field).uuid [bool]
* (govalidator.field).uuidv3 [bool]
* (govalidator.field).uuidv4 [bool]
* (govalidator.field).uuidv5 [bool]
* (govalidator.field).uppercase [bool]
* (govalidator.field).variablewidth [bool]
    
### author

Rangel Reale (rangelspam@gmail.com)

