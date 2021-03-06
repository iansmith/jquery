package jquery

import "github.com/gopherjs/gopherjs/js"

const (
	JQ = "jQuery"
	//keys
	BLUR     = "blur"
	CHANGE   = "change"
	CLICK    = "click"
	DBLCLICK = "dblclick"
	FOCUS    = "focus"
	FOCUSIN  = "focusin"
	FOCUSOUT = "focusout"
	HOVER    = "hover"
	KEYDOWN  = "keydown"
	KEYPRESS = "keypress"
	KEYUP    = "keyup"
	//form
	SUBMIT = "submit"
	LOAD   = "load"
	UNLOAD = "unload"
	RESIZE = "resize"
	//mouse
	MOUSEDOWN  = "mousedown"
	MOUSEENTER = "mouseenter"
	MOUSELEAVE = "mouseleave"
	MOUSEMOVE  = "mousemove"
	MOUSEOUT   = "mouseout"
	MOUSEOVER  = "mouseover"
	MOUSEUP    = "mouseup"
	//touch
	TOUCHSTART  = "touchstart"
	TOUCHMOVE   = "touchmove"
	TOUCHEND    = "touchend"
	TOUCHENTER  = "touchenter"
	TOUCHLEAVE  = "touchleave"
	TOUCHCANCEL = "touchcancel"
	//ajax Events
	AJAXSTART    = "ajaxStart"
	BEFORESEND   = "beforeSend"
	AJAXSEND     = "ajaxSend"
	SUCCESS      = "success"
	AJAXSUCESS   = "ajaxSuccess"
	ERROR        = "error"
	AJAXERROR    = "ajaxError"
	COMPLETE     = "complete"
	AJAXCOMPLETE = "ajaxComplete"
	AJAXSTOP     = "ajaxStop"
)

type JQuery struct {
	o        js.Object
	Jquery   string `js:"jquery"`
	Selector string `js:"selector"` //deprecated according jquery docs
	Length   int    `js:"length"`
	Context  string `js:"context"`
}

type Event struct {
	js.Object
	KeyCode        int       `js:"keyCode"`
	Target         js.Object `js:"target"`
	CurrentTarget  js.Object `js:"currentTarget"`
	DelegateTarget js.Object `js:"delegateTarget"`
	RelatedTarget  js.Object `js:"relatedTarget"`
	Data           js.Object `js:"data"`
	Result         js.Object `js:"result"`
	Which          int       `js:"which"`
	Namespace      string    `js:"namespace"`
	MetaKey        bool      `js:"metaKey"`
	PageX          int       `js:"pageX"`
	PageY          int       `js:"pageY"`
	Type           string    `js:"type"`
}

type JQueryCoordinates struct {
	Left int
	Top  int
}

func (event *Event) PreventDefault() {
	event.Call("preventDefault")
}

func (event *Event) IsDefaultPrevented() bool {
	return event.Call("isDefaultPrevented").Bool()
}

func (event *Event) IsImmediatePropogationStopped() bool {
	return event.Call("isImmediatePropogationStopped").Bool()
}

func (event *Event) IsPropagationStopped() bool {
	return event.Call("isPropagationStopped").Bool()
}

func (event *Event) StopImmediatePropagation() {
	event.Call("stopImmediatePropagation")
}

func (event *Event) StopPropagation() {
	event.Call("stopPropagation")
}

//JQuery constructor
func NewJQuery(args ...interface{}) JQuery {
	return JQuery{o: js.Global.Get(JQ).New(args...)}
}

//static function
func Trim(text string) string {
	return js.Global.Get(JQ).Call("trim", text).Str()
}

//static function
func GlobalEval(cmd string) {
	js.Global.Get(JQ).Call("globalEval", cmd)
}

//static function
func Type(sth interface{}) string {
	return js.Global.Get(JQ).Call("type", sth).Str()
}

//static function
func IsPlainObject(sth interface{}) bool {
	return js.Global.Get(JQ).Call("isPlainObject", sth).Bool()
}

//static function
func IsEmptyObject(sth interface{}) bool {
	return js.Global.Get(JQ).Call("isEmptyObject", sth).Bool()
}

//static function
func IsFunction(sth interface{}) bool {
	return js.Global.Get(JQ).Call("isFunction", sth).Bool()
}

//static function
func IsNumeric(sth interface{}) bool {
	return js.Global.Get(JQ).Call("isNumeric", sth).Bool()
}

//static function
func IsXMLDoc(sth interface{}) bool {
	return js.Global.Get(JQ).Call("isXMLDoc", sth).Bool()
}

//static function
func IsWindow(sth interface{}) bool {
	return js.Global.Get(JQ).Call("isWindow", sth).Bool()
}

//static function
func InArray(val interface{}, arr []interface{}) int {
	return js.Global.Get(JQ).Call("inArray", val, arr).Int()
}

//static function
func Contains(container interface{}, contained interface{}) bool {
	return js.Global.Get(JQ).Call("contains", container, contained).Bool()
}

//static function
func ParseHTML(text string) []interface{} {
	return js.Global.Get(JQ).Call("parseHTML", text).Interface().([]interface{})
}

//static function
func ParseXML(text string) interface{} {
	return js.Global.Get(JQ).Call("parseXML", text).Interface()
}

//static function
func ParseJSON(text string) interface{} {
	return js.Global.Get(JQ).Call("parseJSON", text).Interface()
}

//static function
func Grep(arr []interface{}, fn func(interface{}, int) bool) []interface{} {
	return js.Global.Get(JQ).Call("grep", arr, fn).Interface().([]interface{})
}

//static function
func Noop() interface{} {
	return js.Global.Get(JQ).Get("noop").Interface()
}

//static function
func Now() float64 {
	return js.Global.Get(JQ).Call("now").Float()
}

//static function
func Unique(arr js.Object) js.Object {
	return js.Global.Get(JQ).Call("unique", arr)
}

//methods
func (j JQuery) Each(fn func(int, JQuery)) JQuery {
	j.o = j.o.Call("each", func(idx int, elem js.Object) {
		fn(idx, NewJQuery(elem))
	})
	return j
}

func (j JQuery) Underlying() js.Object {
	return j.o
}

func (j JQuery) Get(i ...interface{}) js.Object {
	return j.o.Call("get", i...)
}

func (j JQuery) Append(i ...interface{}) JQuery {
	return j.dom2args("append", i...)
}

func (j JQuery) Empty() JQuery {
	j.o = j.o.Call("empty")
	return j
}

func (j JQuery) Detach(i ...interface{}) JQuery {
	j.o = j.o.Call("detach", i...)
	return j
}

func (j JQuery) Eq(idx int) JQuery {
	j.o = j.o.Call("eq", idx)
	return j
}
func (j JQuery) FadeIn(i ...interface{}) JQuery {
	j.o = j.o.Call("fadeIn", i...)
	return j
}

func (j JQuery) Delay(i ...interface{}) JQuery {
	j.o = j.o.Call("delay", i...)
	return j
}

//to "range" over selection:
func (j JQuery) ToArray() []interface{} {
	return j.o.Call("toArray").Interface().([]interface{})
}

func (j JQuery) Remove(i ...interface{}) JQuery {
	j.o = j.o.Call("remove", i...)
	return j
}

func (j JQuery) Stop(i ...interface{}) JQuery {
	j.o = j.o.Call("stop", i...)
	return j
}

func (j JQuery) AddBack(i ...interface{}) JQuery {
	j.o = j.o.Call("addBack", i...)
	return j
}

func (j JQuery) Css(name string) string {
	return j.o.Call("css", name).Str()
}

func (j JQuery) CssArray(arr ...string) map[string]interface{} {
	return j.o.Call("css", arr).Interface().(map[string]interface{})
}

func (j JQuery) SetCss(i ...interface{}) JQuery {
	j.o = j.o.Call("css", i...)
	return j
}

func (j JQuery) Text() string {
	return j.o.Call("text").Str()
}

func (j JQuery) SetText(i interface{}) JQuery {

	switch i.(type) {
	case func(int, string) string, string:
	default:
		print("SetText Argument should be 'string' or 'func(int, string) string'")
	}

	j.o = j.o.Call("text", i)
	return j
}

func (j JQuery) Val() string {
	return j.o.Call("val").Str()
}

func (j JQuery) SetVal(i interface{}) JQuery {
	j.o.Call("val", i)
	return j
}

//can return string or bool
func (j JQuery) Prop(property string) interface{} {
	return j.o.Call("prop", property).Interface()
}

func (j JQuery) SetProp(i ...interface{}) JQuery {
	j.o = j.o.Call("prop", i...)
	return j
}

func (j JQuery) RemoveProp(property string) JQuery {
	j.o = j.o.Call("removeProp", property)
	return j
}

func (j JQuery) Attr(property string) string {
	attr := j.o.Call("attr", property)
	if attr == js.Undefined {
		return ""
	}
	return attr.Str()
}

func (j JQuery) SetAttr(i ...interface{}) JQuery {
	j.o = j.o.Call("attr", i...)
	return j
}

func (j JQuery) RemoveAttr(property string) JQuery {
	j.o = j.o.Call("removeAttr", property)
	return j
}

func (j JQuery) HasClass(class string) bool {
	return j.o.Call("hasClass", class).Bool()
}

func (j JQuery) AddClass(i interface{}) JQuery {
	switch i.(type) {
	case func(int, string) string, string:
	default:
		print("addClass Argument should be 'string' or 'func(int, string) string'")
	}
	j.o = j.o.Call("addClass", i)
	return j
}

func (j JQuery) RemoveClass(property string) JQuery {
	j.o = j.o.Call("removeClass", property)
	return j
}

func (j JQuery) ToggleClass(i ...interface{}) JQuery {
	j.o = j.o.Call("toggleClass", i...)
	return j
}

func (j JQuery) Focus() JQuery {
	j.o = j.o.Call("focus")
	return j
}

func (j JQuery) Blur() JQuery {
	j.o = j.o.Call("blur")
	return j
}

func (j JQuery) ReplaceAll(i interface{}) JQuery {
	return j.dom1arg("replaceAll", i)
}
func (j JQuery) ReplaceWith(i interface{}) JQuery {
	return j.dom1arg("replaceWith", i)
}

func (j JQuery) After(i ...interface{}) JQuery {
	return j.dom2args("after", i...)
}

func (j JQuery) Before(i ...interface{}) JQuery {
	return j.dom2args("before", i...)
}

func (j JQuery) Prepend(i ...interface{}) JQuery {
	return j.dom2args("prepend", i...)
}

func (j JQuery) PrependTo(i interface{}) JQuery {
	return j.dom1arg("prependTo", i)
}

func (j JQuery) AppendTo(i interface{}) JQuery {
	return j.dom1arg("appendTo", i)
}

func (j JQuery) InsertAfter(i interface{}) JQuery {
	return j.dom1arg("insertAfter", i)
}

func (j JQuery) InsertBefore(i interface{}) JQuery {
	return j.dom1arg("insertBefore", i)
}

func (j JQuery) Show() JQuery {
	j.o = j.o.Call("show")
	return j
}

func (j JQuery) Hide() JQuery {
	j.o.Call("hide")
	return j
}

func (j JQuery) Toggle(showOrHide bool) JQuery {
	j.o = j.o.Call("toggle", showOrHide)
	return j
}

func (j JQuery) Contents() JQuery {
	j.o = j.o.Call("contents")
	return j
}

func (j JQuery) Html() string {
	return j.o.Call("html").Str()
}

func (j JQuery) SetHtml(i interface{}) JQuery {

	switch i.(type) {
	case func(int, string) string, string:
	default:
		print("SetHtml Argument should be 'string' or 'func(int, string) string'")
	}

	j.o = j.o.Call("html", i)
	return j
}

func (j JQuery) Closest(i ...interface{}) JQuery {
	return j.dom2args("closest", i...)
}

func (j JQuery) End() JQuery {
	j.o = j.o.Call("end")
	return j
}

func (j JQuery) Add(i ...interface{}) JQuery {
	return j.dom2args("add", i...)
}

func (j JQuery) Clone(b ...interface{}) JQuery {
	j.o = j.o.Call("clone", b...)
	return j
}

func (j JQuery) Height() int {
	return j.o.Call("height").Int()
}

func (j JQuery) SetHeight(value string) JQuery {
	j.o = j.o.Call("height", value)
	return j
}

func (j JQuery) Width() int {
	return j.o.Call("width").Int()
}

func (j JQuery) SetWidth(i interface{}) JQuery {

	switch i.(type) {
	case func(int, string) string, string:
	default:
		print("SetWidth Argument should be 'string' or 'func(int, string) string'")
	}

	j.o = j.o.Call("width", i)
	return j
}

func (j JQuery) InnerHeight() int {
	return j.o.Call("innerHeight").Int()
}

func (j JQuery) InnerWidth() int {
	return j.o.Call("innerWidth").Int()
}

func (j JQuery) Offset() JQueryCoordinates {
	obj := j.o.Call("offset")
	return JQueryCoordinates{Left: obj.Get("left").Int(), Top: obj.Get("top").Int()}
}

func (j JQuery) SetOffset(jc JQueryCoordinates) JQuery {
	j.o = j.o.Call("offset", jc)
	return j
}

func (j JQuery) OuterHeight(includeMargin ...bool) int {
	if len(includeMargin) == 0 {
		return j.o.Call("outerHeight").Int()
	}
	return j.o.Call("outerHeight", includeMargin[0]).Int()
}
func (j JQuery) OuterWidth(includeMargin ...bool) int {

	if len(includeMargin) == 0 {
		return j.o.Call("outerWidth").Int()
	}
	return j.o.Call("outerWidth", includeMargin[0]).Int()
}

func (j JQuery) Position() JQueryCoordinates {
	obj := j.o.Call("position")
	return JQueryCoordinates{Left: obj.Get("left").Int(), Top: obj.Get("top").Int()}
}

func (j JQuery) ScrollLeft() int {
	return j.o.Call("scrollLeft").Int()
}
func (j JQuery) SetScrollLeft(value int) JQuery {
	j.o = j.o.Call("scrollLeft", value)
	return j
}

func (j JQuery) ScrollTop() int {
	return j.o.Call("scrollTop").Int()
}
func (j JQuery) SetScrollTop(value int) JQuery {
	j.o = j.o.Call("scrollTop", value)
	return j
}

func (j JQuery) ClearQueue(queueName string) JQuery {
	j.o = j.o.Call("clearQueue", queueName)
	return j
}

func (j JQuery) SetData(key string, value interface{}) JQuery {
	j.o = j.o.Call("data", key, value)
	return j
}

func (j JQuery) Data(key string) interface{} {
	result := j.o.Call("data", key)
	if result == js.Undefined {
		return nil
	}
	return result.Interface()
}

func (j JQuery) Dequeue(queueName string) JQuery {
	j.o = j.o.Call("dequeue", queueName)
	return j
}

func (j JQuery) RemoveData(name string) JQuery {
	j.o = j.o.Call("removeData", name)
	return j
}

func (j JQuery) OffsetParent() JQuery {
	j.o = j.o.Call("offsetParent")
	return j
}

func (j JQuery) Parent(i ...interface{}) JQuery {
	j.o = j.o.Call("parent", i...)
	return j
}

func (j JQuery) Parents(i ...interface{}) JQuery {
	j.o = j.o.Call("parents", i...)
	return j
}

func (j JQuery) ParentsUntil(i ...interface{}) JQuery {
	j.o = j.o.Call("parentsUntil", i...)
	return j
}

func (j JQuery) Prev(i ...interface{}) JQuery {
	j.o = j.o.Call("prev", i...)
	return j
}

func (j JQuery) PrevAll(i ...interface{}) JQuery {
	j.o = j.o.Call("prevAll", i...)
	return j
}

func (j JQuery) PrevUntil(i ...interface{}) JQuery {
	j.o = j.o.Call("prevUntil", i...)
	return j
}

func (j JQuery) Siblings(i ...interface{}) JQuery {
	j.o = j.o.Call("siblings", i...)
	return j
}

func (j JQuery) Slice(i ...interface{}) JQuery {
	j.o = j.o.Call("slice", i...)
	return j
}

func (j JQuery) Children(selector interface{}) JQuery {
	j.o = j.o.Call("children", selector)
	return j
}

func (j JQuery) Unwrap() JQuery {
	j.o = j.o.Call("unwrap")
	return j
}

func (j JQuery) Wrap(obj interface{}) JQuery {
	j.o = j.o.Call("wrap", obj)
	return j
}

func (j JQuery) WrapAll(i interface{}) JQuery {
	return j.dom1arg("wrapAll", i)
}

func (j JQuery) WrapInner(i interface{}) JQuery {
	return j.dom1arg("wrapInner", i)
}

func (j JQuery) Next(i ...interface{}) JQuery {
	j.o = j.o.Call("next", i...)
	return j
}

func (j JQuery) NextAll(i ...interface{}) JQuery {
	j.o = j.o.Call("nextAll", i...)
	return j
}

func (j JQuery) NextUntil(i ...interface{}) JQuery {
	j.o = j.o.Call("nextUntil", i...)
	return j
}

func (j JQuery) Not(i ...interface{}) JQuery {
	j.o = j.o.Call("not", i...)
	return j
}

func (j JQuery) Filter(i ...interface{}) JQuery {
	j.o = j.o.Call("filter", i...)
	return j
}

func (j JQuery) Find(i ...interface{}) JQuery {
	j.o = j.o.Call("find", i...)
	return j
}

func (j JQuery) First() JQuery {
	j.o = j.o.Call("first")
	return j
}

func (j JQuery) Has(selector string) JQuery {
	j.o = j.o.Call("has", selector)
	return j
}

func (j JQuery) Is(i ...interface{}) bool {
	return j.o.Call("is", i...).Bool()
}

func (j JQuery) Last() JQuery {
	j.o = j.o.Call("last")
	return j
}

func (j JQuery) Ready(handler func()) JQuery {
	j.o = j.o.Call("ready", handler)
	return j
}

func (j JQuery) Resize(i ...interface{}) JQuery {
	j.o = j.o.Call("resize", i...)
	return j
}

func (j JQuery) Scroll(i ...interface{}) JQuery {
	return j.handleEvent("scroll", i...)
}

func (j JQuery) FadeOut(i ...interface{}) JQuery {
	j.o = j.o.Call("fadeOut", i...)
	return j
}

func (j JQuery) Select(i ...interface{}) JQuery {
	return j.handleEvent("select", i...)
}

func (j JQuery) Submit(i ...interface{}) JQuery {
	return j.handleEvent("submit", i...)
}

func (j JQuery) handleEvent(evt string, i ...interface{}) JQuery {

	switch len(i) {
	case 0:
		j.o = j.o.Call(evt)
	case 1:
		j.o = j.o.Call(evt, func(e js.Object) {
			i[0].(func(Event))(Event{Object: e})
		})
	case 2:
		j.o = j.o.Call(evt, i[0].(map[string]interface{}), func(e js.Object) {
			i[1].(func(Event))(Event{Object: e})
		})
	default:
		print(evt + " event expects 0 to 2 arguments")
	}
	return j
}

func (j JQuery) Trigger(i ...interface{}) JQuery {
	j.o = j.o.Call("trigger", i...)
	return j
}

func (j JQuery) On(p ...interface{}) JQuery {
	return j.events("on", p...)
}

func (j JQuery) One(p ...interface{}) JQuery {
	return j.events("one", p...)
}

func (j JQuery) Off(p ...interface{}) JQuery {
	return j.events("off", p...)
}

func (j JQuery) events(evt string, p ...interface{}) JQuery {

	count := len(p)

	var isEventFunc bool
	switch p[len(p)-1].(type) {
	case func(Event):
		isEventFunc = true
	default:
		isEventFunc = false
	}

	switch count {
	case 0:
		j.o = j.o.Call(evt)
		return j
	case 1:
		j.o = j.o.Call(evt, p[0])
		return j
	case 2:
		if isEventFunc {
			j.o = j.o.Call(evt, p[0], func(e js.Object) {
				p[1].(func(Event))(Event{Object: e})
			})
			return j
		} else {
			j.o = j.o.Call(evt, p[0], p[1])
			return j
		}
	case 3:
		if isEventFunc {

			j.o = j.o.Call(evt, p[0], p[1], func(e js.Object) {
				p[2].(func(Event))(Event{Object: e})
			})
			return j

		} else {
			j.o = j.o.Call(evt, p[0], p[1], p[2])
			return j
		}
	case 4:
		if isEventFunc {

			j.o = j.o.Call(evt, p[0], p[1], p[2], func(e js.Object) {
				p[3].(func(Event))(Event{Object: e})
			})
			return j

		} else {
			j.o = j.o.Call(evt, p[0], p[1], p[2], p[3])
			return j
		}
	default:
		print(evt + " event should no have more than 4 arguments")
		j.o = j.o.Call(evt, p...)
		return j
	}
}

func (j JQuery) dom2args(method string, i ...interface{}) JQuery {

	switch len(i) {
	case 2:
		selector, selOk := i[0].(JQuery)
		context, ctxOk := i[1].(JQuery)
		if !selOk && !ctxOk {
			j.o = j.o.Call(method, i[0], i[1])
			return j
		} else if selOk && !ctxOk {
			j.o = j.o.Call(method, selector.o, i[1])
			return j
		} else if !selOk && ctxOk {
			j.o = j.o.Call(method, i[0], context.o)
			return j
		}
		j.o = j.o.Call(method, selector.o, context.o)
		return j
	case 1:
		selector, selOk := i[0].(JQuery)
		if !selOk {
			j.o = j.o.Call(method, i[0])
			return j
		}
		j.o = j.o.Call(method, selector.o)
		return j
	default:
		print(" only 1 or 2 parameters allowed for method ", method)
		return j
	}
}

func (j JQuery) dom1arg(method string, i interface{}) JQuery {

	selector, selOk := i.(JQuery)
	if !selOk {
		j.o = j.o.Call(method, i)
		return j
	}
	j.o = j.o.Call(method, selector.o)
	return j
}

//ajax
func Param(params map[string]interface{}) {
	js.Global.Get(JQ).Call("param", params)
}

func (j JQuery) Load(i ...interface{}) JQuery {
	j.o = j.o.Call("load", i...)
	return j
}

func (j JQuery) Serialize() string {
	return j.o.Call("serialize").Str()
}

func (j JQuery) SerializeArray() js.Object {
	return j.o.Call("serializeArray")
}

func Ajax(options map[string]interface{}) Deferred {
	return Deferred{js.Global.Get(JQ).Call("ajax", options)}
}

func AjaxPrefilter(i ...interface{}) {
	js.Global.Get(JQ).Call("ajaxPrefilter", i...)
}

func AjaxSetup(options map[string]interface{}) {
	js.Global.Get(JQ).Call("ajaxSetup", options)
}

func AjaxTransport(i ...interface{}) {
	js.Global.Get(JQ).Call("ajaxTransport", i...)
}

func Get(i ...interface{}) Deferred {
	return Deferred{js.Global.Get(JQ).Call("get", i...)}
}

func Post(i ...interface{}) Deferred {
	return Deferred{js.Global.Get(JQ).Call("post", i...)}
}

func GetJSON(i ...interface{}) Deferred {
	return Deferred{js.Global.Get(JQ).Call("getJSON", i...)}
}

func GetScript(i ...interface{}) Deferred {
	return Deferred{js.Global.Get(JQ).Call("getScript", i...)}
}

func (d Deferred) Promise() js.Object {
	return d.Call("promise")
}

type Deferred struct {
	js.Object
}

func (d Deferred) Then(fn ...interface{}) Deferred {
	return Deferred{d.Call("then", fn...)}
}

func (d Deferred) Always(fn ...interface{}) Deferred {
	return Deferred{d.Call("always", fn...)}
}

func (d Deferred) Done(fn ...interface{}) Deferred {
	return Deferred{d.Call("done", fn...)}
}

func (d Deferred) Fail(fn ...interface{}) Deferred {
	return Deferred{d.Call("fail", fn...)}
}

func (d Deferred) Progress(fn interface{}) Deferred {
	return Deferred{d.Call("progress", fn)}

}

func When(d ...interface{}) Deferred {
	return Deferred{js.Global.Get(JQ).Call("when", d...)}
}

func (d Deferred) State() string {
	return d.Call("state").Str()
}

func NewDeferred() Deferred {
	return Deferred{js.Global.Get(JQ).Call("Deferred")}
}

func (d Deferred) Resolve(i ...interface{}) Deferred {
	return Deferred{d.Call("resolve", i...)}
}

func (d Deferred) Reject(i ...interface{}) Deferred {
	return Deferred{d.Call("reject", i...)}

}

func (d Deferred) Notify(i interface{}) Deferred {
	return Deferred{d.Call("notify", i)}

}

//2do: animations api
//2do: test values against "undefined" values
//2do: more docs
