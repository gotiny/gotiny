package gotiny

import (
	"fmt"
	"strings"
	"regexp"
)

// Route defines each route
// with a utility to match
// and return a map with vars

type Route struct {
	format string
	variablesNames []string
	regexp *regexp.Regexp
}

func (route *Route) Match (URL string) map[string]string {
	matches := route.regexp.MatchString(URL)
	if matches {
		fmt.Println("> ", matches)
		routeExtractedVars := route.regexp.FindAllStringSubmatch(URL,100)
		var routeVars []string;
		if len(routeExtractedVars) == 1 {
			routeVars = routeExtractedVars[0]
			if len(routeVars) == len(route.variablesNames) + 1 {
				routeVars = routeVars[1:]
			}
		}

		fmt.Println("> ", route.variablesNames)
		fmt.Println("> ", routeExtractedVars)
		fmt.Println("> ", routeVars)

		var variableMapping map[string]string;
		if len(route.variablesNames) == len(routeVars) {
			variableMapping = make(map[string]string)
			for i := range route.variablesNames {
				varName := strings.Trim(route.variablesNames[i], "\\/ ><")
				varValue := strings.Trim(routeVars[i], "/\\")
				variableMapping[varName] = varValue
			}
			return variableMapping
		}
	}
	return nil
}

func NewRoute(routeFormat string) *Route {
	route := new(Route)
	route.format = routeFormat

	// Constants
	const varRegex = "(.*?)"

	// Construct regexp
	routeRegex := routeFormat
	variablesRegexp, _ := regexp.Compile("<(.*?)>")
	variablesTags := variablesRegexp.FindAllString(routeFormat,100)
	variablesNames := make([]string,0)
	for i := range variablesTags {
		// Replace <var> with (.*?) in the regexp
		// and store it in the variablesNames slice/array
		varName := variablesTags[i]
		varNameStripped := strings.Trim(varName,"\\/<>")
		variablesNames = append(variablesNames, varNameStripped)

		routeRegex = strings.Replace( routeRegex, varName, varRegex, 1 )
	}
	route.variablesNames = variablesNames
	routeRegex = fmt.Sprint(routeRegex,"/$")
	fmt.Println("REGEX >>> ", routeRegex)
	route.regexp, _ = regexp.Compile(routeRegex)

	return route
}