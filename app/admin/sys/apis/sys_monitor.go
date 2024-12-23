package apis

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/host"
	"go-admin/core/dto/api"
	"go-admin/core/lang"
	"go-admin/core/utils/fileutils"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/strutils"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	_ "go-admin/core/dto/response"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Monitor struct {
	api.Api
}

// GetMonitor admin-获取服务器信息
func (e Monitor) GetMonitor(c *gin.Context) {
	e.Context = c

	sysInfo, err := host.Info()
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()
	osDic["ip"] = iputils.GetLocaHost()
	osDic["projectDir"] = fileutils.GetCurrentPath()
	osDic["hostName"] = sysInfo.Hostname
	osDic["time"] = time.Now().Format("2006-01-02 15:04:05")

	dis, _ := disk.Usage("/")
	diskTotalGB := int(dis.Total) / GB
	diskFreeGB := int(dis.Free) / GB
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = diskTotalGB
	diskDic["free"] = diskFreeGB

	mem, _ := mem.VirtualMemory()
	memUsedMB := int(mem.Used) / GB
	memTotalMB := int(mem.Total) / GB
	memFreeMB := int(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = memTotalMB
	memDic["used"] = memUsedMB
	memDic["free"] = memFreeMB
	memDic["usage"] = memUsedPercent

	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuInfo"], _ = cpu.Info()
	percent, _ := cpu.Percent(0, false)
	cpuDic["Percent"] = strutils.Round(percent[0], 2)
	cpus, _ := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	cpuDic["cpus"] = cpus
	cpuDic["cpuNum"], _ = cpu.Counts(false)

	//服务器磁盘信息
	disklist := make([]disk.UsageStat, 0)
	//所有分区
	diskInfo, err := disk.Partitions(true)
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskDetail.Total = diskDetail.Total / 1024 / 1024
				diskDetail.Used = diskDetail.Used / 1024 / 1024
				diskDetail.Free = diskDetail.Free / 1024 / 1024
				disklist = append(disklist, *diskDetail)
			}
		}
	}

	result := map[string]interface{}{
		"os":       osDic,
		"mem":      memDic,
		"cpu":      cpuDic,
		"disk":     diskDic,
		"diskList": disklist,
	}
	e.OK(result, lang.MsgByCode(lang.SuccessCode, e.Lang))
}
