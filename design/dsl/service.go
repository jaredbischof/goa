package dsl

import (
	"github.com/goadesign/goa/eval"
	"github.com/goadesign/goa/rpc/design"
)

// Service defines a group of related endpoints that are exposed from the same process.
//
// Service is as a top level expression.
//
// Example:
//
//    var _ = Service("divider", func() {
//        Description("divider service") // Optional description
//
//        Endpoint("divide", func() {     // Defines a single endpoint
//            Description("The divide endpoint returns the division of A and B")
//            Request(DividePayload)
//            Response(Division)
//            Error("DivisionByZero", ErrDivByZero) // ErrDivByZero is optional type that describes error body.
//               If gRPC error attribute is added to type, if return error matches design error then
//               error attribute is set otherwise error is returned to gRPC server.
//
//            HTTP(func() {
//                Scheme("https")
//                GET("/{Dividend/{Divisor}") // DividePayload must have Dividend and Divisor attributes
//                POST("/{Dividend}")         // Body is DividePayload minus Dividend attribute and headers
//                POST("/")                   // Body is DividePayload minus headers
//                Header("Account")           // Must match one of DividePayload attributes
//                Response(func() {
//                    Status(OK)              // Default
//                    Header("Result")        // Must be an attribute of Division
//                })
//                Error("DivisionByZero", func() {
//                    Status(BadRequest)      // Default
//                    Header("Message")       // Must be an attribute of ErrDivByZero
//                })
//            })
//
//            GRPC(func() {
//                Proto("divider.divide") // rpc definition in proto file
//                Error("DivisionByZero", func() { // Defines which field contains error if not "Error"
//                    Field("DivByZero")
//                })
//            })
//        })
//    })
//
func Service(name string, adsl ...func()) *design.ServiceExpr {
	s := &design.ServiceExpr{Name: name}
	if len(adsl) > 1 {
		eval.ReportError("too many arguments in call to Service")
		return nil
	}
	if len(adsl) == 1 {
		s.DSLFunc = adsl[0]
	}
	if _, ok := eval.Current().(eval.TopExpr); !ok {
		eval.IncompatibleDSL()
		return nil
	}
	design.Root.Services = append(design.Root.Services, s)
	return s
}
