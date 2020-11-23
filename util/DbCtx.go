package util
//
//import (
//	"time"
//)
//
//type dbContext struct {
//	href
//	lastOpTime time.Time
//}
//
//var ctxList = make(map[string]dbContext)
//
//func PutCtx(id string, dbInfo interface{}) {
//	var dbctx = &dbContext{
//		href:       &dbInfo,
//		lastOpTime: time.Now(),
//	}
//	ctxList[id] = *dbctx
//}
//
//func GetCtx(id string) interface{} {
//	if v, ok := ctxList[id]; ok{
//		return v.href
//	}
//	return nil
//}