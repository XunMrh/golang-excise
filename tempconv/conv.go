package tempconv

func CToF(c Celsius) Fahranheit { return Fahranheit(c*9/5 + 32) }
func FToC(f Fahranheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvins    { return Kelvins((c - AbsoluteC)) }
func KToC(k Kelvins) Celsius    { return Celsius(k + Kelvins(AbsoluteC)) }
