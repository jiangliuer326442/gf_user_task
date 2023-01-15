package apollo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
	"time"
)

var (
	ctx            = gctx.New()
	configMap      = make(map[string]interface{})
	apolloEndpoint = gconv.String(g.Cfg().MustGetWithEnv(ctx, "apollo.endpoint"))
	appid          = gconv.String(g.Cfg().MustGetWithEnv(ctx, "apollo.appid"))
	cluster        = gconv.String(g.Cfg().MustGetWithEnv(ctx, "apollo.cluster"))
	namespaces     = gstr.Split(gconv.String(g.Cfg().MustGetWithEnv(ctx, "apollo.namespace")), ",")
	pullInterval   = gconv.Int64(g.Cfg().MustGetWithEnv(ctx, "apollo.pullInterval"))
)

func init() {
	pullApollo(ctx)
	var myApolloAdapter = &apolloAdapter{}
	g.Cfg().SetAdapter(myApolloAdapter)
	gtimer.SetInterval(ctx, time.Duration(pullInterval), pullApollo)
}

func pullApollo(ctx context.Context) {
	g.Log().Debug(ctx, "pull apollo")
	for _, namespace := range namespaces {
		url := apolloEndpoint + "configfiles/json/" + appid + "/" + cluster + "/" + namespace
		result, err := (&http.Client{}).Get(url)
		if err != nil {
			glog.Error(ctx, err)
			glog.Debug(ctx, url)
			return
		}
		resultBytes, err := io.ReadAll(result.Body)
		if err != nil {
			glog.Error(ctx, err)
			glog.Debug(ctx, url)
			return
		}
		configMapTmp := make(map[string]interface{})
		if err := json.Unmarshal(resultBytes, &configMapTmp); err != nil {
			fmt.Println("apollo pull result: ", url, string(resultBytes), err)
			return
		}
		for k, v := range configMapTmp {
			if gstr.HasPrefix(gconv.String(v), "[") || gstr.HasPrefix(gconv.String(v), "{") {
				configMap[k] = gjson.New(v, true)
			} else {
				configMap[k] = v
			}

		}
	}
}

type apolloAdapter struct {
}

func (apollo *apolloAdapter) Available(ctx context.Context, resource ...string) (ok bool) {
	return configMap != nil
}

func (apollo *apolloAdapter) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	value, ok := configMap[pattern]

	if ok {
		return value, nil
	} else {
		return nil, nil
	}
}

func (apollo *apolloAdapter) Data(ctx context.Context) (data map[string]interface{}, err error) {
	return configMap, nil
}
