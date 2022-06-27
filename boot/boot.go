package boot

import (
	_ "goApplet/packed"

	"github.com/gogf/gf/frame/g"

    "github.com/gogf/swagger"
    //swagger "github.com/gogf/swagger/v2"
)

// 用于应用初始化。
func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}
