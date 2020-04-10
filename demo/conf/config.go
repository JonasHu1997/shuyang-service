package conf

// SysEnvType 系统运行环境
type SysEnvType uint8

const (
	// SysEnvDev 开发环境
	SysEnvDev SysEnvType = iota
	// SysEnvTest 测试环境
	SysEnvTest
	// SysEnvOLT 线上测试环境
	SysEnvOLT
	// SysEnvPro 生产环境
	SysEnvPro
)

// SysEnv 当前系统配置
const SysEnv = SysEnvDev
