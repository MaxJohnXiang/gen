package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// example for init the database:
//
//  DB, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=true")
//  if err != nil {
//  	panic("failed to connect database: " + err.Error())
//  }
//  defer db.Close()

var DB *gorm.DB

func ConfigRouter() http.Handler {
	router := httprouter.New()
	configOwnAppsRouter(router)
	configTBatchSignsRouter(router)
	configTBrandsRouter(router)
	configTExpressesRouter(router)
	configTInventoryEntriesRouter(router)
	configTInventoryProductsRouter(router)
	configTPicksRouter(router)
	configTPickOrderitemsRouter(router)
	configTProductLogsRouter(router)
	configTProductProcessesRouter(router)
	configTShelvesRouter(router)
	configTShelf2Router(router)
	configTShelfProductsRouter(router)
	configTShelveEntrtiesRouter(router)
	configTShelveEntryProductsRouter(router)
	configTShiftsRouter(router)
	configTShiftProductsRouter(router)
	configTStockEntriesRouter(router)
	configTStockEntryProductsRouter(router)
	configTTransfersRouter(router)
	configTTransferProductsRouter(router)
	configTemOrdersRouter(router)
	configTemShelfProductsRouter(router)
	configTemShelfProduct1Router(router)
	configTmpStoreDupsRouter(router)

	return router
}

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}
	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
