package describe_test

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_regexpDescription(t *testing.T) {
	initial := "{{Short description|Canadian computer scientist}}\n{{Use mdy dates|date=March 2019}}\n{{Infobox scientist\n|name              = Yoshua Bengio\n|image             = Yoshua Bengio 2019 cropped.jpg\n|image_size        = \n| honorific_suffix = {{post-nominals|country=CAN|OC|FRS|FRSC|size=100}}\n|caption           = Yoshua Bengio in 2019\n|birth_date        = {{birth date and age|1964|3|5}}\n|birth_place  "
	r := regexp.MustCompile(`\{\{Short description\|(.*)\}\}`)
	fmt.Println(r.FindStringSubmatch(initial))
}
