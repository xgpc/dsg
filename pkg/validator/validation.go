package validator

import (
	"sync"

	"github.com/xgpc/dsg/util"
)

var (
	transIns *util.Translations
	once     = &sync.Once{}
)

func GetTranslations() *util.Translations {
	if transIns == nil {
		once.Do(func() {
			transIns = util.NewTranslationIns(util.WithRulesOption(&Rules), util.WithRulesMsgOption(&RulesMsg))
		})
	}

	return transIns
}
