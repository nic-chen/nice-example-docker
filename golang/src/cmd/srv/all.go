package srv

import(
	"github.com/nic-chen/nice/micro/registry"
	opentracing "github.com/opentracing/opentracing-go"
)

func RunAll(register registry.Registry, tracer opentracing.Tracer) {
	go RunMemberSrv(register, tracer)
	RunApi(register, tracer)
}
