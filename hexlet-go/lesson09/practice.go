package lesson09

// BEGIN (write your solution here)
func DomainForLocale(domain, locale string) string {
	if locale != "" {
		return locale + "." + domain
	} else {
		return "en." + domain
	}
}

///END
