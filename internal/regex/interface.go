package regex

import (
	"fmt"
	"github.com/johnfercher/chaos/internal/model"
	"regexp"
	"strings"
)

var interfaceName = regexp.MustCompile(`type\s.+interface\s+{`)

func GetInterfaces(file string) []*model.Interface {
	fullInterfaces := interfaceName.FindAllString(file, -1)
	var interfaces []*model.Interface

	for _, fullInterface := range fullInterfaces {
		_interface := strings.ReplaceAll(fullInterface, "type ", "")
		_interface = strings.ReplaceAll(_interface, " interface {", "")
		interfaces = append(interfaces, &model.Interface{
			Name: _interface,
		})
	}

	for i, _interface := range interfaces {
		methods := getInterfaceMethods(file, _interface.Name)
		interfaces[i].Methods = methods
	}

	return interfaces
}

func getInterfaceMethods(file string, name string) []model.Method {
	pattern := fmt.Sprintf(`type\s%s\sinterface\s+{`, name)
	begin := regexp.MustCompile(pattern)

	scope := GetMultiLineScope(file, begin, closeBrackets)
	fmt.Println(scope)

	lines := strings.Split(scope, "\n")

	methods := []model.Method{}
	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		line = strings.ReplaceAll(line, "\t", "")
		m := GetMethod(line)
		methods = append(methods, m)
	}

	return methods
}
