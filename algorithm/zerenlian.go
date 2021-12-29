package algorithm

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct {
}

// Parse Parse
func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse Parse
func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

//func main() {
//	parser := &jsonRuleConfigParser{}
//	parser.Parse([]byte{})
//}

// NewIRuleConfigParser new irule config
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
