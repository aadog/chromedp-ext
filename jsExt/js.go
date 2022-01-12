package jsExt

import "fmt"

func QuerySelector(selector string) string {
	return fmt.Sprintf(`document.querySelector("%s")`, selector)
}
func QuerySelectorVisible(selector string) string {
	return fmt.Sprintf(`Boolean(document.querySelector("%s").offsetWidth||document.querySelector("%s").offsetHeight||document.querySelector("%s").getClientRects().length)`, selector, selector, selector)
}
func QuerySelectorObjEvaluate(selector string, expression string) string {
	e := fmt.Sprintf(`var obj=document.querySelector('%s');if(obj!=null){%s;true}else{false}`, selector, expression)
	return e
}
