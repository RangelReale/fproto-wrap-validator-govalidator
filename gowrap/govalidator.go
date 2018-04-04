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

		check_str := "is not"
		if agv.Source == "false" {
			check_str = "is"
		}

		supported := false

		check_func := ""
		check_desc := ""

		if agn == "is_ascii" {
			check_func = "IsAscii"
			check_desc = "ASCII"
		} else if agn == "is_alpha" {
			check_func = "IsAlpha"
			check_desc = "alpha"
		} else if agn == "is_alphanumeric" {
			supported = true
			g.P("if ", govalidator_alias, ".IsAlphanumeric(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s alphanumeric")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_base64" {
			supported = true
			g.P("if ", govalidator_alias, ".IsBase64(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s base64")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_cidr" {
			supported = true
			g.P("if ", govalidator_alias, ".IsCIDR(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s CIDR")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_creditcard" {
			supported = true
			g.P("if ", govalidator_alias, ".IsCreditCard(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s credit card")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_dnsname" {
			supported = true
			g.P("if ", govalidator_alias, ".IsDNSName(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s DNS name")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_datauri" {
			supported = true
			g.P("if ", govalidator_alias, ".IsDataURI(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s data URI")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_dialstring" {
			supported = true
			g.P("if ", govalidator_alias, ".IsDialString(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s dial string")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_email" {
			supported = true
			g.P("if ", govalidator_alias, ".IsEmail(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s email")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_hexadecimal" {
			supported = true
			g.P("if ", govalidator_alias, ".IsHexadecimal(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s hexadecimal")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_hexcolor" {
			supported = true
			g.P("if ", govalidator_alias, ".IsHexcolor(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s hex color")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_host" {
			supported = true
			g.P("if ", govalidator_alias, ".IsHost(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s host")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_ip" {
			supported = true
			g.P("if ", govalidator_alias, ".IsIP(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s ip")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_ipv4" {
			supported = true
			g.P("if ", govalidator_alias, ".IsIPv4(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s ipv4")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_ipv6" {
			supported = true
			g.P("if ", govalidator_alias, ".IsIPv6(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s ipv6")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		} else if agn == "is_isbn10" {
			supported = true
			g.P("if ", govalidator_alias, ".IsISBN10(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s ISBN10")`, errors_alias, check_str)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		}

		if check_func != "" {
			supported = true
			g.P("if ", govalidator_alias, ".", check_func, "(", varSrc, ") != ", agv.Source, " {")
			g.In()
			error_msg := fmt.Sprintf(`%s.New("Value %s %s")`, errors_alias, check_str, check_desc)
			vh.GenerateValidationErrorAdd(g.G(), error_msg, agn, fproto_gowrap_validator.VEID_INVALID_VALUE, "govalidator", agn, "govalidator_check", agv.Source)
			g.Out()
			g.P("}")
		}

		if !supported {
			return fmt.Errorf("Validation %s not supported for type %s", agn, tp.FullOriginalName())
		}
	}

	return nil
}
