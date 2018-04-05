package fproto_gowrap_validator_govalidator

import (
	"fmt"
	"strings"

	"github.com/RangelReale/fdep"
	"github.com/RangelReale/fproto"
	"github.com/RangelReale/fproto-wrap-validator/gowrap"
	"github.com/RangelReale/fproto-wrap/gowrap"
)

type ValidatorPlugin_Govalidator struct {
}

func (tp *ValidatorPlugin_Govalidator) GetValidator(validatorType *fdep.OptionType) fproto_gowrap_validator.Validator {
	// govalidator.field
	if validatorType.Option != nil &&
		validatorType.Option.DepFile.FilePath == "github.com/RangelReale/fproto-wrap-validator-govalidator/govalidator.proto" &&
		validatorType.Option.DepFile.ProtoFile.PackageName == "govalidator" {
		if validatorType.Name == "field" {
			return &Validator_Govalidator{validatorType: validatorType}
		}
	}
	return nil
}

func (tp *ValidatorPlugin_Govalidator) ValidatorPrefixes() []string {
	return []string{"govalidator"}
}

//
// Validator_Govalidator
//
type Validator_Govalidator struct {
	validatorType *fdep.OptionType
}

func (tp *Validator_Govalidator) FPValidator() {

}

func (t *Validator_Govalidator) GenerateValidation(g *fproto_gowrap.GeneratorFile, vh fproto_gowrap_validator.ValidatorHelper, tp *fdep.DepType, option *fproto.OptionElement, varSrc string) error {
	tinfo := g.G().GetTypeInfo(tp)

	if tinfo.Converter().TCID() == fproto_gowrap.TCID_SCALAR {
		return t.generateValidation_scalar(g, vh, tp, tinfo, option, varSrc)
	}

	tv := vh.GetTypeValidator(t.validatorType, tinfo, tp)

	if tv != nil {
		return tv.GenerateValidation(g, vh, tp, option, varSrc)
	}

	return fmt.Errorf("Unknown type for validator: %s", tp.FullOriginalName())
}

func (t *Validator_Govalidator) generateValidation_scalar(g *fproto_gowrap.GeneratorFile, vh fproto_gowrap_validator.ValidatorHelper, tp *fdep.DepType, tinfo fproto_gowrap.TypeInfo, option *fproto.OptionElement, varSrc string) error {
	var opag []string
	for _, agn := range option.AggregatedSorted() {
		opag = append(opag, fmt.Sprintf("%s=%s", agn, option.AggregatedValues[agn].Source))
	}

	g.P("// ", option.Name, " -- ", option.ParenthesizedName, " ** ", option.NPName, " @@ ", option.Value.Source, " %% ", strings.Join(opag, ", "))

	switch *tp.ScalarType {
	case fproto.StringScalar:
		//
		// STRING
		//
		return t.generateValidation_scalar_string(g, vh, tp, tinfo, option, varSrc)
	}

	return fmt.Errorf("Validation not supported for type %s", tp.FullOriginalName())
}

func (t *Validator_Govalidator) generateValidation_scalar_string(g *fproto_gowrap.GeneratorFile, vh fproto_gowrap_validator.ValidatorHelper, tp *fdep.DepType, tinfo fproto_gowrap.TypeInfo, option *fproto.OptionElement, varSrc string) error {
	errors_alias := g.DeclDep("errors", "errors")
	govalidator_alias := g.DeclDep("github.com/asaskevich/govalidator", "gv")

	for _, agn := range option.AggregatedSorted() {
		agv := option.AggregatedValues[agn]

		check_value := "true"
		check_str := "is not"
		if agv.Source == "false" {
			check_value = "false"
			check_str = "is"
		}

		supported := false

		check_func := ""
		check_desc := ""

		if agn == "ascii" {
			check_func = "IsASCII"
			check_desc = "ASCII"
		} else if agn == "alpha" {
			check_func = "IsAlpha"
			check_desc = "alpha"
		} else if agn == "alphanumeric" {
			check_func = "IsAlphanumeric"
			check_desc = "alphanumeric"
		} else if agn == "base64" {
			check_func = "IsBase64"
			check_desc = "base64"
		} else if agn == "cidr" {
			check_func = "IsCIDR"
			check_desc = "CIDR"
		} else if agn == "creditcard" {
			check_func = "IsCreditCard"
			check_desc = "credit card"
		} else if agn == "dnsname" {
			check_func = "IsDNSName"
			check_desc = "DNS name"
		} else if agn == "datauri" {
			check_func = "IsDataURI"
			check_desc = "data URI"
		} else if agn == "dialstring" {
			check_func = "IsDialString"
			check_desc = "dial string"
		} else if agn == "email" {
			check_func = "IsEmail"
			check_desc = "email"
		} else if agn == "hexadecimal" {
			check_func = "IsHexadecimal"
			check_desc = "hexadecimal"
		} else if agn == "hexcolor" {
			check_func = "IsHexcolor"
			check_desc = "hex color"
		} else if agn == "host" {
			check_func = "IsHost"
			check_desc = "host"
		} else if agn == "ip" {
			check_func = "IsIP"
			check_desc = "ip"
		} else if agn == "ipv4" {
			check_func = "IsIPv4"
			check_desc = "ipv4"
		} else if agn == "ipv6" {
			check_func = "IsIPv6"
			check_desc = "ipv6"
		} else if agn == "isbn10" {
			check_func = "IsISBN10"
			check_desc = "ISBN10"
		} else if agn == "isbn13" {
			check_func = "IsISBN13"
			check_desc = "ISBN13"
		} else if agn == "iso3166alpha2" {
			check_func = "IsISO3166Alpha2"
			check_desc = "ISO3166Alpha2"
		} else if agn == "iso3166alpha3" {
			check_func = "IsISO3166Alpha3"
			check_desc = "ISO3166Alpha3"
		} else if agn == "iso693alpha2" {
			check_func = "IsISO693Alpha2"
			check_desc = "ISO693Alpha2"
		} else if agn == "iso693alpha3b" {
			check_func = "IsISO693Alpha3b"
			check_desc = "ISO693Alpha3b"
		} else if agn == "iso4217" {
			check_func = "IsISO4217"
			check_desc = "ISO4217"
		} else if agn == "json" {
			check_func = "IsJSON"
			check_desc = "JSON"
		} else if agn == "latitude" {
			check_func = "IsLatitude"
			check_desc = "latitude"
		} else if agn == "longitude" {
			check_func = "IsLongitude"
			check_desc = "longitude"
		} else if agn == "lowercase" {
			check_func = "IsLowerCase"
			check_desc = "lowerCase"
		} else if agn == "mac" {
			check_func = "IsMAC"
			check_desc = "MAC"
		} else if agn == "mongoid" {
			check_func = "IsMongoID"
			check_desc = "MongoID"
		} else if agn == "multibyte" {
			check_func = "IsMultibyte"
			check_desc = "multibyte"
		} else if agn == "null" {
			check_func = "IsNull"
			check_desc = "null"
		} else if agn == "numeric" {
			check_func = "IsNumeric"
			check_desc = "numeric"
		} else if agn == "port" {
			check_func = "IsPort"
			check_desc = "port"
		} else if agn == "printableascii" {
			check_func = "IsPrintableASCII"
			check_desc = "printable ASCII"
		} else if agn == "rfc3339" {
			check_func = "IsRFC3339"
			check_desc = "RFC3339"
		} else if agn == "rfc3339withoutzone" {
			check_func = "IsRFC3339WithoutZone"
			check_desc = "RFC3339 without zone"
		} else if agn == "rgbcolor" {
			check_func = "IsRGBcolor"
			check_desc = "RGB color"
		} else if agn == "requesturi" {
			check_func = "IsRequestURI"
			check_desc = "request URI"
		} else if agn == "requesturl" {
			check_func = "IsRequestURL"
			check_desc = "request URL"
		} else if agn == "ssn" {
			check_func = "IsSSN"
			check_desc = "SSN"
		} else if agn == "semver" {
			check_func = "IsSemver"
			check_desc = "semver"
		} else if agn == "url" {
			check_func = "IsURL"
			check_desc = "URL"
		} else if agn == "utfdugit" {
			check_func = "IsUTFDigit"
			check_desc = "UTF digit"
		} else if agn == "utfletter" {
			check_func = "IsUTFLetter"
			check_desc = "UTF letter"
		} else if agn == "utfletternumeric" {
			check_func = "IsUTFLetterNumeric"
			check_desc = "UTF letter numeric"
		} else if agn == "utfnumeric" {
			check_func = "IsUTFNumeric"
			check_desc = "UTF numeric"
		} else if agn == "uuid" {
			check_func = "IsUUID"
			check_desc = "UUID"
		} else if agn == "uuidv3" {
			check_func = "IsUUIDv3"
			check_desc = "UUIDv3"
		} else if agn == "uuidv4" {
			check_func = "IsUUIDv4"
			check_desc = "UUIDv4"
		} else if agn == "uuidv5" {
			check_func = "IsUUIDv5"
			check_desc = "UUIDv5"
		} else if agn == "uppercase" {
			check_func = "IsUpperCase"
			check_desc = "upperCase"
		} else if agn == "variablewidth" {
			check_func = "IsVariableWidth"
			check_desc = "variable width"
		}

		if check_func != "" {
			supported = true
			g.P("if ", govalidator_alias, ".", check_func, "(", varSrc, ") != ", check_value, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s %s")`, errors_alias, check_str, check_desc)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", check_value)
			g.Out()
			g.P("}")
		}

		if !supported {
			return fmt.Errorf("Validation %s not supported for type %s", agn, tp.FullOriginalName())
		}
	}

	return nil
}
