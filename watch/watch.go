package watch

import (
	"errors"
	"github.com/violin8/pkg/logs"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Watch 监听文件
// 必须是已经存在的文件，监听文件变化，回调文件内容

// Watcher 监听指定文件的变化
func Watcher(path string, fn func(path string) error) error {
	// 判断文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("watch file " + path + " not exist")
	}
	name := filepath.Clean(path)
	logs.Info("watch file %v", name)
	nw, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer nw.Close()
	if err := nw.Add(name); err != nil {
		return err
	}
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case ev := <-nw.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if ev.Op&fsnotify.Create == fsnotify.Create {
						// logs.Info("创建文件 : %v", ev.Name)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						if ev.Name == name {
							logs.Info("watch file change %v", name)
							if fn != nil {
								_ = fn(path)
							}
						}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						// logs.Info("删除文件 : %v", ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						// logs.Info("重命名文件 : %v", ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						// logs.Info("修改权限 : %v", ev.Name)
					}
				}
			case err := <-nw.Errors:
				{
					logs.Error("watch %v err : %v", name, err)
					return
				}
			}
		}
	}()

	//循环
	select {}
}
