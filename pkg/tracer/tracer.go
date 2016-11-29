package tracer

import (
	"gitlab.com/tonyhb/keepupdated/pkg/consts"

	"github.com/opentracing/opentracing-go"
	"github.com/sourcegraph/appdash"
)

func NewTracer() opentracing.Tracer {
	if os.Getenv(consts.EnvTracingDisabled) {
		return opentracing.NoopTracer{}
	}
	return newAppdashTracer()
}

func newAppdashTracer() opentracing.Tracer {
	collector := appdash.NewRemoteCollector(os.Getenv(consts.EnvTracerAddr))
	return appdashtracer.NewTracer(collector)
}
