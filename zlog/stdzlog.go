/**
* FileName: stdzlog
* Description:
* Author:   ww
* Date:     2021/12/7 7:07 下午
 */
package zlog

import "os"

// StdZinxLog 创建全局log
var StdZinxLog = NewZinxLog(os.Stderr, "", BitDefault)

/*
* 获取 StdZinxLog 标记位
 */
func Flags() int {
	return StdZinxLog.Flags()
}

/*
* 设置 StdZinxLog 标记位
 */
func ResetFlags(flag int) {
	StdZinxLog.ResetFlags(flag)
}

/*
* 添加flag标记
 */
func AddFlag(flag int) {
	StdZinxLog.AddFlag(flag)
}

/*
* 设置StdZinxLog 日志头前缀
 */
func SetPrefix(prefix string) {
	StdZinxLog.SetPrefix(prefix)
}

/*
* 设置StdZinxLog绑定的日志文件
 */
func SetLogFile(fileDir string, fileName string) {
	StdZinxLog.SetLogFile(fileDir, fileName)
}

/*
* 设置关闭debug
 */
func CloseDebug() {
	StdZinxLog.CloseDebug()
}

/*
* 设置打开debug
 */
func OpenDebug() {
	StdZinxLog.OpenDebug()
}

//Debugf ====> Debug <====
func Debugf(format string, v ...interface{}) {
	StdZinxLog.Debugf(format, v...)
}

//Debug Debug
func Debug(v ...interface{}) {
	StdZinxLog.Debug(v...)
}

//Infof ====> Info <====
func Infof(format string, v ...interface{}) {
	StdZinxLog.Infof(format, v...)
}

//Info -
func Info(v ...interface{}) {
	StdZinxLog.Info(v...)
}

// ====> Warn <====
func Warnf(format string, v ...interface{}) {
	StdZinxLog.Warnf(format, v...)
}

func Warn(v ...interface{}) {
	StdZinxLog.Warn(v...)
}

// ====> Error <====
func Errorf(format string, v ...interface{}) {
	StdZinxLog.Errorf(format, v...)
}

func Error(v ...interface{}) {
	StdZinxLog.Error(v...)
}

// ====> Fatal 需要终止程序 <====
func Fatalf(format string, v ...interface{}) {
	StdZinxLog.Fatalf(format, v...)
}

func Fatal(v ...interface{}) {
	StdZinxLog.Fatal(v...)
}

// ====> Panic  <====
func Panicf(format string, v ...interface{}) {
	StdZinxLog.Panicf(format, v...)
}

func Panic(v ...interface{}) {
	StdZinxLog.Panic(v...)
}

// ====> Stack  <====
func Stack(v ...interface{}) {
	StdZinxLog.Stack(v...)
}

func init() {
	// StdZinxLog对象 对所有输出方法做了一层包裹，所以在打印调用函数的时候，比正常的logger对象多一层调用
	// 一般的zinxLogger对象 calldDepth=2, StdZinxLog的calldDepth=3
	StdZinxLog.calledDepth = 3
}
