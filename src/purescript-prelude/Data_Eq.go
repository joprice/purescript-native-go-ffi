package purescript_prelude

import (
	"Data_Eq"
	. "purescript"
)

func eqImpl(a Any) Any {
	return func(b Any) Any {
		return a == b
	}
}

func init() {
	exports := Data_Eq.Foreign

	exports["eqBooleanImpl"] = eqImpl

	exports["eqIntImpl"] = eqImpl

	exports["eqNumberImpl"] = eqImpl

	exports["eqStringImpl"] = eqImpl

	exports["eqCharImpl"] = eqImpl

	exports["eqArrayImpl"] = func(f Any) Any {
		return func(xs_ Any) Any {
			return func(ys_ Any) Any {
				xs := xs_.(Array)
				ys := ys_.(Array)
				if len(xs) != len(ys) {
					return false
				}
				for i := 0; i < len(xs); i++ {
					equal, _ := Apply(f, xs[i], ys[i]).(bool)
					if !equal {
						return false
					}
				}
				return true
			}
		}
	}

}