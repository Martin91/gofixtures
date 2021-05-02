package fixtures

import (
	"time"

	"github.com/bxcodec/faker/v3"
)

var evaluators = map[string]Evaluator{}

// Evaluator is used to generate data for columns
type Evaluator func() interface{}

func RegisterEvaluator(name string, evaluator Evaluator) {
	if _, existed := evaluators[name]; existed {
		panic("conflict evaluators for " + name)
	}
	evaluators[name] = evaluator
}

func init() {
	RegisterEvaluator("Now", func() interface{} { return time.Now().Unix() })

	// The below list evaluators to generate fake data
	RegisterEvaluator("Latitude", func() interface{} { return faker.Latitude() })
	RegisterEvaluator("Longitude", func() interface{} { return faker.Longitude() })
	RegisterEvaluator("UnixTime", func() interface{} { return faker.UnixTime() })
	RegisterEvaluator("Date", func() interface{} { return faker.Date() })
	RegisterEvaluator("TimeString", func() interface{} { return faker.TimeString() })
	RegisterEvaluator("MonthName", func() interface{} { return faker.MonthName() })
	RegisterEvaluator("YearString", func() interface{} { return faker.YearString() })
	RegisterEvaluator("DayOfWeek", func() interface{} { return faker.DayOfWeek() })
	RegisterEvaluator("DayOfMonth", func() interface{} { return faker.DayOfMonth() })
	RegisterEvaluator("Timestamp", func() interface{} { return faker.Timestamp() })
	RegisterEvaluator("Century", func() interface{} { return faker.Century() })
	RegisterEvaluator("Timezone", func() interface{} { return faker.Timezone() })
	RegisterEvaluator("Timeperiod", func() interface{} { return faker.Timeperiod() })

	RegisterEvaluator("Email", func() interface{} { return faker.Email() })
	RegisterEvaluator("MacAddress", func() interface{} { return faker.MacAddress() })
	RegisterEvaluator("DomainName", func() interface{} { return faker.DomainName() })
	RegisterEvaluator("URL", func() interface{} { return faker.URL() })
	RegisterEvaluator("Username", func() interface{} { return faker.Username() })
	RegisterEvaluator("IPv4", func() interface{} { return faker.IPv4() })
	RegisterEvaluator("IPv6", func() interface{} { return faker.IPv6() })
	RegisterEvaluator("Password", func() interface{} { return faker.Password() })

	RegisterEvaluator("Word", func() interface{} { return faker.Word() })
	RegisterEvaluator("Sentence", func() interface{} { return faker.Sentence() })
	RegisterEvaluator("Paragraph", func() interface{} { return faker.Paragraph() })

	RegisterEvaluator("CCType", func() interface{} { return faker.CCType() })
	RegisterEvaluator("CCNumber", func() interface{} { return faker.CCNumber() })
	RegisterEvaluator("Currency", func() interface{} { return faker.Currency() })
	RegisterEvaluator("AmountWithCurrency", func() interface{} { return faker.AmountWithCurrency() })

	RegisterEvaluator("TitleMale", func() interface{} { return faker.TitleMale() })
	RegisterEvaluator("TitleFemale", func() interface{} { return faker.TitleFemale() })
	RegisterEvaluator("FirstName", func() interface{} { return faker.FirstName() })
	RegisterEvaluator("FirstNameMale", func() interface{} { return faker.FirstNameMale() })
	RegisterEvaluator("FirstNameFemale", func() interface{} { return faker.FirstNameFemale() })
	RegisterEvaluator("LastName", func() interface{} { return faker.LastName() })
	RegisterEvaluator("Name", func() interface{} { return faker.Name() })

	RegisterEvaluator("Phonenumber", func() interface{} { return faker.Phonenumber() })
	RegisterEvaluator("TollFreePhoneNumber", func() interface{} { return faker.TollFreePhoneNumber() })
	RegisterEvaluator("E164PhoneNumber", func() interface{} { return faker.E164PhoneNumber() })

	RegisterEvaluator("UUIDHyphenated", func() interface{} { return faker.UUIDHyphenated() })
	RegisterEvaluator("UUIDDigit", func() interface{} { return faker.UUIDDigit() })
}
