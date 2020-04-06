package panel

// TODO: 该面板后端需要两套，一套为超级管理员使用， 一套为蜜罐用户使用 先完成超级管理员面板
// TODO：超级管理员能够管理蜜罐用户，能够查看各蜜罐收集的信息等

func Start() {
	s := new(Service)
	s.init()
}
