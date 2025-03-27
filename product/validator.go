package product

import (
	"context"
	"fmt"
	"github.com/PaesslerAG/gval"
	"github.com/busy-cloud/iot/calc"
	"github.com/spf13/cast"
	"regexp"
	"strings"
	"time"
)

type Compare struct {
	Type  string  `json:"type"` //= != > >= < <=
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (c *Compare) Evaluate(ctx map[string]any) (bool, error) {
	val, ok := ctx[c.Name]
	if !ok {
		return false, fmt.Errorf("compare evalute field %s not found", c.Name)
	}
	v, err := cast.ToFloat64E(val)
	if err != nil {
		return false, err
	}
	switch c.Type {
	case "=", "==":
		return v == c.Value, nil
	case "!=", "~=", "<>":
		return v != c.Value, nil
	case ">":
		return v > c.Value, nil
	case "<":
		return v < c.Value, nil
	case ">=":
		return v >= c.Value, nil
	case "<=":
		return v <= c.Value, nil
	default:
		return false, fmt.Errorf("unsupported compare type: %s", c.Type)
	}
}

type Validator struct {
	Type       string  `json:"type"` //compare对比， expression表达式
	Compare    Compare `json:"compare,omitempty"`
	Expression string  `json:"expression,omitempty"`
	Title      string  `json:"title,omitempty"`
	Message    string  `json:"message,omitempty"`
	Level      int     `json:"level,omitempty"`
	Delay      int64   `json:"delay,omitempty"`
	Reset      int64   `json:"reset,omitempty"`
	ResetTimes int     `json:"reset_times,omitempty"`
	Disabled   bool    `json:"disabled,omitempty"`

	expression gval.Evaluable
}

func (v *Validator) Build() (err error) {
	v.expression, err = calc.Compile(v.Expression)
	return err
}

func (v *Validator) Evaluate(ctx map[string]any) (*Alarm, error) {
	var err error
	var ret bool

	switch v.Type {
	case "compare":
		ret, err = v.expression.EvalBool(context.Background(), ctx)
	case "expression":
		ret, err = v.Compare.Evaluate(ctx)
	default:
		err = fmt.Errorf("unsupported validator type: %s", v.Type)
	}

	if err != nil {
		return nil, err
	}

	//条件为 假，则重置
	if ret {
		ctx["__start"] = 0
		ctx["__times"] = 0
		return nil, nil
	}

	var start int64 = 0
	var times int = 0

	//起始时间
	now := time.Now().Unix()
	s, ok := ctx["__start"]
	if !ok {
		start = now
		ctx["__start"] = now
	} else {
		start = cast.ToInt64(s)
	}

	t, ok := ctx["__times"]
	if ok {
		times = cast.ToInt(t)
	}

	//延迟报警
	if v.Delay > 0 {
		if now < start+v.Delay {
			return nil, nil
		}
	}

	if times > 0 {
		//重复报警
		if v.Reset <= 0 {
			return nil, nil
		}

		//超过最大次数
		if v.ResetTimes > 0 && times >= v.ResetTimes {
			return nil, nil
		}

		//还没到时间
		if now < start+v.Reset {
			return nil, nil
		}

		ctx["__start"] = now
	}
	ctx["__times"] = times + 1

	//产生报警
	alarm := &Alarm{
		Title:   replaceParams(v.Title, ctx),
		Message: replaceParams(v.Message, ctx),
		Level:   v.Level,
	}

	return alarm, nil
}

type Alarm struct {
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
	Level   int    `json:"level,omitempty"`
}

var paramsRegex *regexp.Regexp

func init() {
	paramsRegex = regexp.MustCompile(`\{\w+\}`)
}

func replaceParams(str string, ctx map[string]any) string {
	return paramsRegex.ReplaceAllStringFunc(str, func(s string) string {
		s = strings.TrimPrefix(s, "{")
		s = strings.TrimSuffix(s, "}")
		return cast.ToString(ctx[s])
	})
}
