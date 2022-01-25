package movies

import (
	"fengqi/kodi-metadata-tmdb-cli/config"
	"github.com/fsnotify/fsnotify"
)

type Collector struct {
	config  *config.Config
	watcher *fsnotify.Watcher
	channel chan *Movie
	nfoMode int
}

// Movie 电影目录详情，从名字分析
// Fortress.2021.BluRay.1080p.AVC.DTS-HD.MA5.1-MTeam
type Movie struct {
	Dir             string `json:"dir"`
	OriginTitle     string `json:"origin_title"`   // 原始目录名
	VideoFileName   string `json:"file_name"`      // 视频文件名，仅限：IsSingleFile=true
	Title           string `json:"title"`          // 名称 Hawkeye
	MovieId         int    `json:"tv_id"`          // 电影id
	Year            int    `json:"year"`           // 年份：2020、2021
	IsFile          bool   `json:"is_file"`        // 是否是单文件，而不是目录
	Suffix          string `json:"suffix"`         // 单文件时，文件的后缀
	IsBluray        bool   `json:"is_bluray"`      // 蓝光目录
	IsDvd           bool   `json:"is_dvd"`         // DVD目录
	IsSingleFile    bool   `json:"is_single_file"` // 普通的单文件视频
	IdCacheFile     string `json:"id_cache_file"`
	DetailCacheFile string `json:"detail_cache_file"`
}
