package jaeger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"sing/app/util/jaeger_service"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {

		var parentSpan opentracing.Span

		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = jaeger_service.Tracer.StartSpan(c.Request.URL.Path)
			defer parentSpan.Finish()

			fmt.Println("======@@@@@@@@@@@@@@@@@@=======")
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()

			fmt.Println("======&&&&&&&&&&&&&&&&&&&&&&=======")
		}

		c.Set("Tracer", jaeger_service.Tracer)
		c.Set("ParentSpanContext", parentSpan.Context())

		c.Next()
	}
}
