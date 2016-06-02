// youtube_prog
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/otium/ytdl"
	"github.com/vaughan0/go-ini"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const pageonload = 3
const key = "AIzaSyBqBiYPy9FFoYAqIRmZFK9MIat9YQR3Pao"
const namefilech = "channels"
const styles = `<style>
.image {
	width: 427px; /*width: 333px;*/
	height: 240px;/*height: 250px;*/
	border-radius: 40px;}
.time {
	background: linear-gradient(#7B68EE, #ababab);
	border-radius: 20px;}
.add {
	text-align: center;
	margin-left: 12%;
	margin-top: -22%;
    font-size: 400%;
	font-style: normal;
	color: rgba(0,0,0,0.4);
	text-shadow: 0 1px 0 #bcb8ae, 0 -1px 0 #97958e;
	font-family: Helvetica, Arial, sans-serif;
	display: inline-block;
	opacity: 0.6;}
.add:hover {
	opacity: 1;}
.aaaa:before,.aaa:after {
    display: inline-block;
    font-style: normal;}
.aaaa:before  {
    content: "+";    }
.aaa:before,.aaa:after {
    display: inline-block;
    font-style: normal;}
.aaa:before  {
    content: "←";}
.aaa2:before,.aaa:after {
    display: inline-block;
    font-style: normal;}
.aaa2:before  {
    content: "dl";}
.avatar4:hover {
	opacity: 1;	}
.href {
	font-size: 300%;
	display: inline-block;}
.href:hover {
	background-color:#7B68EE;}
.avatar {
	width: 150px;
	height: 150px;
	border-radius: 50%;}
.avatar20 {
	z-index: 200;
	opacity: 0.6;}
.avatar200 {
	top: 500px;
	opacity: 0.6;
	width: 150px;
	height: 150px;
	border-radius: 50%;
	font-size: 900%;
	cursor: pointer;
	background-size: cover;
	background-color:#7B68EE;
	background-position:  0px 0px;
	outline: 0;
	border: 1px;
	border-style: dotted;
	position: absolute;}
.avatar2 {
	text-align: center;
	position: fixed;
    right: 1%;
    top: 5px;
	z-index: 200;
	opacity: 0.6;}
.avatar2:hover {
	opacity: 1;}
.avatar20:hover {
	opacity: 1;}
.avatar333 {
	font-size: 150%;
	font-style: normal;
	color: rgba(0,0,0,0.4);
	text-shadow: 0 1px 0 #bcb8ae, 0 -1px 0 #97958e;
	font-family: Helvetica, Arial, sans-serif;
	z-index: 201;}
.avatar33 {
	font-size: 600%;
	font-style: normal;
	color: rgba(0,0,0,0.4);
	text-shadow: 0 1px 0 #bcb8ae, 0 -1px 0 #97958e;
	font-family: Helvetica, Arial, sans-serif;
	z-index: 199;}
.avatar3 {
	vertical-align: middle;
	border-radius: 50%;
	text-align: center;
	position: fixed;
	right: 2%;
	top: 35px;
	padding: 10px;
	z-index: 200;	
	text-transform: uppercase;
	font-style: normal;
	font-weight: bold;
	color: rgba(0,0,0,0.4);
	text-shadow: 0 1px 0 #bcb8ae, 0 -1px 0 #97958e;
	/*font-family: Helvetica, Arial, sans-serif;*/
	font-size: 300%;
	/*top: 0%; 
	right: 0%; 
	bottom: 0; 
	left: 0;*/
	opacity: 0.9;}
.avatar4 {
	vertical-align: middle;
	border-radius: 50%;
	text-align: center;
	position: fixed;
	right: 1%;
	top: 330px;
	padding: 10px;
	z-index: 200;
	width:150px;
	height:150px;
	border-radius:50%;
	background-color:#7B68EE;
	text-transform: uppercase;
	font-style: normal;
	font-weight: bold;
	color: rgba(0,0,0,0.4);
	/*text-shadow: 0 1px 0 #bcb8ae, 0 -1px 0 #97958e;*/
	font-family: Helvetica, Arial, sans-serif;
	font-size: 660%;
	opacity: 0.9;}
.avatar5 {
	text-align: center;
	position: fixed;
    right: 1%;
	top: 165px;
    /*padding: 10px;*/
	z-index: 200;
	opacity: 0.6;}
.avatar44 {
	border-radius: 50%;
	/*z-index: 200;*/
	width:150px;
	height:150px;
	border-radius:50%;
	background-color:#7B68EE;
	text-transform: uppercase;
	font-style: normal;
	font-weight: bold;
	color: rgba(0,0,0,0.4);
	/*text-shadow: 0 1px 0 #bcb8ae, 0 -1px 0 #97958e;*/
	font-family: Helvetica, Arial, sans-serif;
	font-size: 660%;
	opacity: 0.9;}
.el {
	font-family: 'Arial'; 
	font-size: 500%;
	width:100%; 
	display: block;
	text-align: center;}
.spis {
	font-family: 'Arial'; 
	font-size: 180%;
	/*top: -100px;*/
	position: relative;
	font-weight: bold;
	text-align: center;
	margin: 30px auto;
	/*white-space: nowrap;*/
	/*width:100%; */
	/*display: block;*/
	/*text-align: center;*/}  
.chlabel {
	text-align: center;
	width:100%;
	font-size: 250%;}
button {
	font-family: 'Arial';
	width:100%;
	font-size: 300%;
	text-align: center;
	border-radius: 40px;
	background: #ababab;
	background: -moz-linear-gradient(#f2f2f2, #ababab);
	background: -ms-linear-gradient(#f2f2f2, #ababab);
	background: -o-linear-gradient(#f2f2f2, #ababab);
	background: -webkit-gradient(linear, 0 0, 0 100%, from(#f2f2f2), to(#ababab));
	background: -webkit-linear-gradient(#f2f2f2, #ababab);
	background: linear-gradient(#f2f2f2, #ababab);
	box-shadow: 0 0 10px rgba(0,0,0,0.3),
	0 1px 1px rgba(0,0,0,0.25);} 
select {
	font-size: 500%;
	width:100%;
	direction: rtl;
	border-radius: 40px;}
*,*:after,*:before {
	-webkit-box-sizing: border-box;
	-moz-box-sizing: border-box;
	box-sizing: border-box;
	padding: 0;
	margin: 0;}
.switch {
	margin: 30px auto;
	position: relative;}
.switch label {
	width: 100%;
	height: 100%;
	position: relative;
	display: block;}
.switch input {
	top: 0; 
	right: 0; 
	bottom: 0; 
	left: 0;
	opacity: 0;
	z-index: 100;
	position: absolute;
	width: 100%;
	height: 100%;
	cursor: pointer;}
.switch.chboxydl {
	width: 98%;
	text-align: center;
	height: 240px;
	/*width: 400px;
	height: 180px;*/
	/* width: 180px;
	height: 50px;*/}
.switch.chboxydl label {
	display: block;
	/*text-align: left;*/
	width: 100%;
	height: 100%;
	background: #FF3333;
	border-radius: 40px;
	box-shadow:
	inset 0 3px 8px 1px rgba(0,0,0,0.2),
	0 1px 0 rgba(255,255,255,0.5);}
.switch.chboxydl label:after {
	content: "";
	position: absolute;
	z-index: -1;
	top: -8px; right: -8px; bottom: -8px; left: -8px;
	border-radius: inherit;
	background: #ababab;
	background: -moz-linear-gradient(#f2f2f2, #ababab);
	background: -ms-linear-gradient(#f2f2f2, #ababab);
	background: -o-linear-gradient(#f2f2f2, #ababab);
	background: -webkit-gradient(linear, 0 0, 0 100%, from(#f2f2f2), to(#ababab));
	background: -webkit-linear-gradient(#f2f2f2, #ababab);
	background: linear-gradient(#f2f2f2, #ababab);
	box-shadow: 0 0 10px rgba(0,0,0,0.3),
	0 1px 1px rgba(0,0,0,0.25);}
.switch.chboxydl label:before {
	content: "";
	position: absolute;
	z-index: -1;
	top: -18px; right: -18px; bottom: -18px; left: -18px;
	border-radius: inherit;
	background: #eee;
	background: -moz-linear-gradient(#e5e7e6, #eee);
	background: -ms-linear-gradient(#e5e7e6, #eee);
	background: -o-linear-gradient(#e5e7e6, #eee);
	background: -webkit-gradient(linear, 0 0, 0 100%, from(#e5e7e6), to(#eee));
	background: -webkit-linear-gradient(#e5e7e6, #eee);
	background: linear-gradient(#e5e7e6, #eee);
	box-shadow:
		0 1px 0 rgba(255,255,255,0.5);
		-webkit-filter: blur(1px);
		-moz-filter: blur(1px);
		-ms-filter: blur(1px);
		-o-filter: blur(1px);
		filter: blur(1px);}
.switch.chboxydl label i {
	display: block;
	height: 100%;
	width: 50%;
	border-radius: inherit;
	background: silver;
	position: absolute;
	z-index: 2;
	right: 50%;
	top: 0;
	background: #b2ac9e;
	background: -moz-linear-gradient(#f7f2f6, #b2ac9e);
	background: -ms-linear-gradient(#f7f2f6, #b2ac9e);
	background: -o-linear-gradient(#f7f2f6, #b2ac9e);
	background: -webkit-gradient(linear, 0 0, 0 100%, from(#f7f2f6), to(#b2ac9e));
	background: -webkit-linear-gradient(#f7f2f6, #b2ac9e);
	background: linear-gradient(#f7f2f6, #b2ac9e);
	box-shadow:
	    inset 0 1px 0 white,
	    0 0 8px rgba(0,0,0,0.3),
	    0 5px 5px rgba(0,0,0,0.2);}
.switch.chboxydl label i:after {
	/*  content: "null size";
	font-size: 120%;
	font-weight: bold;
	text-align: center;*/
	content: ""
	position: absolute;
	left: 15%;
	top: 25%;
	width: 70%;
	height: 50%;
	background: #d2cbc3;
	background: -moz-linear-gradient(#cbc7bc, #d2cbc3);
	background: -ms-linear-gradient(#cbc7bc, #d2cbc3);
	background: -o-linear-gradient(#cbc7bc, #d2cbc3);
	background: -webkit-gradient(linear, 0 0, 0 100%, from(#cbc7bc), to(#d2cbc3));
	background: -webkit-linear-gradient(#cbc7bc, #d2cbc3);
	background: linear-gradient(#cbc7bc, #d2cbc3);
	border-radius: inherit;}
.switch.chboxydl label i:before {
	/*content: "OFF";*/
	text-transform: uppercase;
	font-style: normal;
	font-weight: bold;
	color: rgba(0,0,0,0.4);
	text-shadow: 0 1px 0 #bcb8ae, 0 -1px 0 #97958e;
	font-family: Helvetica, Arial, sans-serif;
	font-size: 30px;
	position: absolute;
	top: 50%;
	margin-top: -12px;
	right: -40%;}
.switch.chboxydl input:checked ~ label {
	background: #9abb82;}
.switch.chboxydl input:checked ~ label i {
	right: -1%;}
.switch.chboxydl input:checked ~ label i:before {
	right: 130%;
	color: #82a06a;
	text-shadow: 
	  0 1px 0 #afcb9b,
	  0 -1px 0 #6b8659;}
.switch.chboxydl label .spis {
	position: relative;
	padding-left: 50%;
	padding-right: 3%;}
.switch.chboxydl input:checked ~ label .spis {
	position: relative;
	padding-right: 50%;
	padding-left: 3%;}
</style>`

type Query struct {
	Kind     string    `json:"kind"`
	Etag     string    `json:"etag"`
	PageInfo *PageInfo `json:"pageInfo"`
	Items    []Items_  `json:"items"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type ContentDetails struct {
	RelatedPlaylists *RelatedPlaylists `json:"relatedPlaylists"`
}

type RelatedPlaylists struct {
	Likes   string `json:"likes"`
	Uploads string `json:"uploads"`
}

type Items_ struct {
	Kind           string          `json:"kind"`
	Etag           string          `json:"etag"`
	Id             string          `json:"id"`
	ContentDetails *ContentDetails `json:"contentDetails"`
}

func (slice *QueryV) AddItem(item ItemsV) []ItemsV {
	return append(slice.Items, item)
}

type QueryV struct {
	Kind          string    `json:"kind"`
	Etag          string    `json:"etag"`
	NextPageToken string    `json:"nextPageToken"`
	PageInfo      *PageInfo `json:"pageInfo"`
	Items         []ItemsV  `json:"items"`
}

type PageInfoV struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type ItemsV struct {
	Kind    string    `json:"kind"`
	Etag    string    `json:"etag"`
	Id      string    `json:"id"`
	Snippet *SnippetV `json:"snippet"`
}

type SnippetV struct {
	PublishedAt  string          `json:"publishedAt"`
	ChannelId    string          `json:"channelId"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	ChannelTitle string          `json:"channelTitle"`
	PlaylistId   string          `json:"playlistId"`
	Position     int             `json:"position"`
	ResourceId   ResourceIdV     `json:"resourceId"`
	Thumbnails   *Quality_medium `json:"thumbnails"`
}

type ResourceIdV struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}

type Quality_default struct {
	Quality_Item Quality_Item `json:"default"`
}

type Quality_medium struct {
	Quality_Item Quality_Item `json:"medium"`
}

type Quality_high struct {
	Quality_Item Quality_Item `json:"high"`
}

type Quality_standard struct {
	Quality_Item Quality_Item `json:"standard"`
}

type Quality_maxres struct {
	Quality_Item Quality_Item `json:"maxres"`
}

type Quality_Item struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Item_box struct {
	Name string
	Text string
}

type Items_box struct {
	Items []Item_box
}

func (slice Items_box) Len() int {
	return len(slice.Items)
}

func (slice Items_box) Less(i, j int) bool {
	return slice.Items[i].Name < slice.Items[j].Name
}

func (slice Items_box) Swap(i, j int) {
	slice.Items[i], slice.Items[j] = slice.Items[j], slice.Items[i]
}

func (slice *Items_box) AddItem(item Item_box) []Item_box {
	return append(slice.Items, item)
}

func write_channels_file() string {
	d1 := []byte("")
	ioutil.WriteFile("./"+namefilech, d1, 0777)
	return ""
}

func read_channels() []string {
	dat, _ := ioutil.ReadFile("./" + namefilech)
	new := strings.Replace(string(dat), "<ch>", "", -1)
	new = strings.TrimRight(new, "</ch>")
	return strings.Split(new, "</ch>")
}

func add_channel(text string) {
	for _, val := range read_channels() {
		if val == text {
			return
		}
	}
	if strings.TrimSpace(text) != "" {
		f, err := os.OpenFile("./"+namefilech, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err = f.WriteString("<ch>" + text + "</ch>"); err != nil {
			panic(err)
		}
	}
	return

}

func del_channel(text string, texth string) {
	dat, _ := ioutil.ReadFile("./" + namefilech)
	new := ""
	if text == texth {
		new = strings.Replace(string(dat), "<ch>"+text+"</ch>", "", 1)
	} else {
		new = strings.Replace(string(dat), "<ch>"+texth+"</ch>", "<ch>"+text+"</ch>", 1)
	}
	d1 := []byte(new)
	ioutil.WriteFile("./"+namefilech, d1, 0777)
	return

}

func ulist(data string) string {
	res := Query{}
	json.Unmarshal([]byte(data), &res)
	return res.Items[0].ContentDetails.RelatedPlaylists.Uploads
}

func ulist_pic_ch(data string) string {
	res := QueryV{}
	json.Unmarshal([]byte(data), &res)
	if res.PageInfo.TotalResults == 0 {
		return ""
	}
	return res.Items[0].Snippet.Thumbnails.Quality_Item.Url
}

func ulist_videos(data string) (QueryV, string, int, int) {
	res := QueryV{}
	json.Unmarshal([]byte(data), &res)
	return res, res.NextPageToken, res.PageInfo.ResultsPerPage, res.PageInfo.TotalResults
}

func ulistvideo(data string) QueryV {
	res := QueryV{}
	json.Unmarshal([]byte(data), &res)
	return res
}

func info_ch(username string, seckey string) string {
	response, err := http.Get("https://www.googleapis.com/youtube/v3/channels?part=contentDetails&forUsername=" + username + "&key=" + seckey)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return ulist(string(contents))

	}
	return ""
}

func info_video(videokey string, seckey string) QueryV {
	var res QueryV
	response, err := http.Get("https://www.googleapis.com/youtube/v3/videos?part=snippet&id=" + videokey + "&key=" + seckey)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return ulistvideo(string(contents))

	}
	return res
}

func info_ch_(username string, seckey string) string {
	response, err := http.Get("https://www.googleapis.com/youtube/v3/channels?part=snippet&forUsername=" + username + "&key=" + seckey)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		return ulist_pic_ch(string(contents))

	}
	return ""
}

func videos(ulist string, seckey string, nextpage string) QueryV {
	var res QueryV
	textresp := ""
	if nextpage == "" {
		textresp = "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50&playlistId=" + ulist + "&key=" + seckey
	} else {
		textresp = "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50&pageToken=" + nextpage + "&playlistId=" + ulist + "&key=" + seckey
	}

	response, err := http.Get(textresp)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		res, _, _, _ = ulist_videos(string(contents))
		return res
	}
	return res
}

func read_ini(fname string, sec string) ini.Section {
	//dir, _ := os.Getwd()
	//file, _ := ini.LoadFile(dir + "\\" + fname)
	file, _ := ini.LoadFile(fname)
	return file[sec]
}

func load(sett string, null_ string) {
	ini_ := read_ini(sett, "settings")
	channels := strings.Split(ini_["channels"], "{!}")
	channels_name := strings.Split(ini_["channels_name"], "{!}")
	channels_last_day := strings.Split(ini_["channels_last_day"], "{!}")
	channels_b_w := strings.Split(ini_["channels_b_w"], "{!}")
	for i := 0; i < len(channels); i++ {
		if channels[i] != "" {
			response, err := http.Get(channels[i])
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			} else {
				defer response.Body.Close()
				tokenizer := html.NewTokenizer(response.Body)
				var (
					islink bool
					link   string
				)
				links := make(map[string]string)
			loop:
				for {
					tok := tokenizer.Next()
					switch tok {
					case html.ErrorToken:
						break loop
					case html.StartTagToken:
						tag, _ := tokenizer.TagName()
						if string(tag) == "a" {
							islink = true
							_, linkbytes, _ := tokenizer.TagAttr()
							link = string(linkbytes)
						}
					case html.TextToken:
						if islink {
							links[link] = string(tokenizer.Text())
							islink = false
						}
					}
				}
				for url, _ := range links {
					if strings.Contains(url, "/watch?v=") {
						info, _ := ytdl.GetVideoInfo("https://www.youtube.com" + url)
						day_, _ := strconv.Atoi(channels_last_day[i])
						day_str := strconv.Itoa(info.DatePublished.YearDay() + info.DatePublished.Year()*1000)
						day_to0 := time.Now().YearDay() - day_ + time.Now().Year()*1000
						day_to1 := info.DatePublished.YearDay() + info.DatePublished.Year()*1000
						prov_ := time.Now().YearDay() - day_
						for prov_ < 0 {
							day_--
						}
						if info.Author == channels_name[i] && day_to1 >= day_to0 {
							fmt.Print(info.DatePublished.Format("02-01-2006"), "	", info.Title)
							format := info.Formats.Best("720p")[0]
							if channels_b_w[i] == "low" {
								format = info.Formats.Worst(ytdl.FormatResolutionKey)[0]
							}
							if channels_b_w[i] == "high" {
								format = info.Formats.Best("720p")[0]
							}
							os.Mkdir(info.Author, 0777)
							if _, err := os.Stat(info.Author + "/" + day_str + "_" + strings.Replace(info.Title, ":", " ", 99) + "_" + info.DatePublished.Format("02-01-2006") + "_[" + strings.Replace(url, "/watch?v=", "", 1) + "]_." + format.Extension); os.IsNotExist(err) {
								file, _ := os.Create(info.Author + "/" + day_str + "_" + strings.Replace(info.Title, ":", " ", 99) + "_" + info.DatePublished.Format("02-01-2006") + "_[" + strings.Replace(url, "/watch?v=", "", 1) + "]_." + format.Extension)
								//if null_file[i] != "yes" {
								if null_ != "null" {
									info.Download(format, file)
								}
							}
							fmt.Println(" ----- success")
						}
					}
				}
			}
		}
	}
}

func load_video(url string, title_ string, datepub string) {
	info, _ := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + url)
	const longForm = "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(longForm, datepub)
	format := info.Formats.Best("720p")[0]
	os.Mkdir(info.Author+"", 0777)
	filename_ := access_filename(info.Author) + "/" + t.Format("20060102") + "_" + access_filename(title_) + "_" + t.Format("02-01-2006") + "_[" + url + "]_." + format.Extension
	//	filename_ := access_filename(info.Author) + "/" + info.DatePublished.Format("20060102") + "_" + access_filename(info.Title) + "_" + info.DatePublished.Format("02-01-2006") + "_[" + url + "]_." + format.Extension
	if _, err := os.Stat(filename_); os.IsNotExist(err) {
		file, _ := os.Create(filename_)
		info.Download(format, file)
	}
	//fmt.Println(" ----- success")
}

func video_link(url string) (string, string) {
	info, _ := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + url)
	format := info.Formats.Best("720p")[0]
	link, _ := info.GetDownloadURL(format)
	return link.String(), info.Title
}

func files_in_dir() []os.FileInfo {
	files, _ := ioutil.ReadDir("./")
	return files
}

func access_filename(filename string) string {
	str_ := `*|\:"<>?/~@#$%^&`
	for i := 0; i < len(str_); i++ {
		filename = strings.Replace(filename, string(str_[i]), " ", -1)
	}
	return filename
}

func main() {
	m := martini.Classic()
	m.Post("/load_on_device", func(w http.ResponseWriter, req *http.Request) string {
		text := ""
		body, _ := ioutil.ReadAll(req.Body)
		v, _ := url.ParseQuery(string(body))
		for _, value := range v {
			zn1, zn2 := video_link(value[0])
			text += `<a s
			class="aaa" href="` + zn1 + `">` + zn2 + `</a><br>`
		}
		http.Redirect(w, req, "/", 302)
		return `
		<!DOCTYPE HTML>
		<html>
		<head>
		<meta charset="utf-8">
		<title>Youtube Download</title>
		` + styles + `
		</head>
		<body>
		<br><br><br>
		<a href='../'>НАЗАД</a>
		<br><br><br>
		<form action="/addnew" method="post">
		` + text + `
		<button type="submit" >Добавить</button>
		</form>
		</body>
		</html>`
	})

	m.Post("/download", func(w http.ResponseWriter, req *http.Request) string {
		loadonserv := false
		loadlist := false
		text := ""
		text1 := ""
		text2 := ""
		body, _ := ioutil.ReadAll(req.Body)
		v, _ := url.ParseQuery(string(body))
		for key, _ := range v {
			if key == "plus" {
				loadonserv = true
			}
			if key == "dlist" {
				loadlist = true
			}
		}
		for key_, value := range v {
			if loadonserv && key_ != "plus" && key_ != "dlist" {
				inform := info_video(value[0], key)
				go load_video(value[0], inform.Items[0].Snippet.Title, inform.Items[0].Snippet.PublishedAt)
				http.Redirect(w, req, "/", 302)
			}
			if loadlist && key_ != "plus" && key_ != "dlist" {
				zn1, zn2 := video_link(value[0])
				text += `<a class="href" href="` + zn1 + `">` + zn2 + `</a><br>`
				text1 += `var link = document.createElement('a');
							link.setAttribute('href','` + zn1 + `');
							link.setAttribute('download','` + zn2 + `');
							onload=link.click();
							`

				text2 += `var xhr = new XMLHttpRequest();
						xhr.open('GET', '` + zn1 + `', true);
						xhr.responseType = 'blob';
						xhr.onload = function(e) {
						var blob = this.response;
						var link = document.createElement("a");
						link.href = window.URL.createObjectURL(blob);
						link.innerHTML = "` + zn2 + `";
						link.setAttribute("download","` + zn2 + `.mp4");
						document.body.appendChild(link);
						}
						xhr.send();
						`
			}
		}
		return `
				<!DOCTYPE HTML>
				<html>
				<head>
				<meta charset="utf-8">
				<title>Youtube Download</title>
				` + styles + `
				<script>
				function myFunction() {
				` + text1 + `
				}
				</script>
				</head>
				<body>
				<br><br><br>
				<a class="href" href='../'>НАЗАД</a><br>
				<!--<a onclick="myFunction()" class="href" href='#'>Test</a><br>-->
				<a onclick="myFunction()" class="href" href='../'>Скачать Все</a>
				<br><br><br>
				` + text + `
				</body>
				</html>`
	})

	m.Post("/run", func(req *http.Request) string {
		inidl := ""
		body, _ := ioutil.ReadAll(req.Body)
		v, _ := url.ParseQuery(string(body))
		for key, value := range v {

			if key == "fileini" {
				inidl = value[0]
			}
		}
		//go load(inidl, nulldl)
		textt := ""
		var res QueryV
		//res := videos(info_ch(inidl, key), key, "")
		//nextp := res.NextPageToken
		nextp := ""
		count_page := 1
		//		count_page := res.PageInfo.TotalResults / res.PageInfo.ResultsPerPage
		//		last_page := res.PageInfo.TotalResults % res.PageInfo.ResultsPerPage
		//		if last_page > 0 {
		//			count_page++
		//		}
		//		if pageonload > 0 && count_page > 1 {
		//			count_page = pageonload
		//		}
		// т.к. перенёс всё в 1 цикл, то нужно высчитывать либо заранее либо прямо в цикле наподобие while
		// и в цикле эти страницы высчитывать //тут
		i := 0
		for i < count_page {
			restmp := videos(info_ch(inidl, key), key, nextp)
			//---
			count_page = restmp.PageInfo.TotalResults / restmp.PageInfo.ResultsPerPage
			last_page := restmp.PageInfo.TotalResults % restmp.PageInfo.ResultsPerPage
			if last_page > 0 {
				count_page++
			}
			if pageonload > 0 && count_page >= pageonload {
				count_page = pageonload
			}
			//---
			nextp = restmp.NextPageToken
			for i := 0; i < len(restmp.Items); i++ {
				res.Items = append(res.Items, restmp.Items[i])
			}
			i++
		}

		//		for i := 0; i < count_page; i++ {
		//			restmp := videos(info_ch(inidl, key), key, nextp)
		//			nextp = restmp.NextPageToken
		//			for i := 0; i < len(restmp.Items); i++ {
		//				res.Items = append(res.Items, restmp.Items[i])
		//			}
		//		}
		ccc := len(res.Items)
		const longForm = "2006-01-02T15:04:05.000Z"
		items := []Item_box{}
		box := Items_box{items}
		for i := 0; i < ccc; i++ {
			t, _ := time.Parse(longForm, res.Items[i].Snippet.PublishedAt)
			key_ := t.Format("20060102150405")
			checked := "checked"
			filename_ := access_filename(res.Items[i].Snippet.ChannelTitle) + "/" + t.Format("20060102") + "_" + access_filename(res.Items[i].Snippet.Title) + "_" + t.Format("02-01-2006") + "_[" + res.Items[i].Snippet.ResourceId.VideoId + "]_." + "mp4"
			if _, err := os.Stat(filename_); os.IsNotExist(err) {
				checked = ""
			}
			//texttt := `<div class="switch chboxydl"><input type="checkbox" name="` + res.Items[i].Snippet.ResourceId.VideoId + `" value="` + res.Items[i].Snippet.ResourceId.VideoId + `" ` + checked + `><label><div class="spis"><div class="time">` + t.Format("02-01-2006[15:04]") + `</div><br>` + res.Items[i].Snippet.Title + `</div><i>` + `</i></label></div><br>`

			texttt := `<div class="switch chboxydl"><input type="checkbox" name="` + res.Items[i].Snippet.ResourceId.VideoId + `" value="` + res.Items[i].Snippet.ResourceId.VideoId + `" ` + checked + `><label><div class="spis"><div class="time">` + t.Format("02-01-2006[15:04]") + `</div><br>` + res.Items[i].Snippet.Title + `</div><i><img class="image" src="` + res.Items[i].Snippet.Thumbnails.Quality_Item.Url + `">` + `</i></label></div><br>`
			item1 := Item_box{Name: key_, Text: texttt}
			box.Items = box.AddItem(item1)
		}
		//------------
		//sort.Sort(box)
		sort.Sort(sort.Reverse(box))
		//------------
		for _, k := range box.Items {
			textt += k.Text
		}

		return `
		<!DOCTYPE HTML>
		<html>
		<head>
		<meta charset="utf-8">
		<title>Youtube Download</title>
		` + styles + `
		</head>
		<body>
		<br><br>
		<form action="/download" method="post">
		<!--<div class="avatar3">+</div>-->
		<div class="avatar2"><input type="submit" name="plus" class="avatar33" value="+" style="background-size: cover;background-image: url(` + info_ch_(inidl, key) + `);width:150px;height:150px;border-radius:50%;cursor: pointer;background-position:  0px 0px;outline: 0;border: 0px;" /></div>
		<div class="avatar4"><a class="aaa" href='../' style="width:150px;height:150px;border-radius:50%;cursor: pointer;"/></a></div>
		<div class="avatar5"><input type="submit" name="dlist" class="avatar33" value="links" style="top: 410px;background-size: cover;background-image: url(` + info_ch_(inidl, key) + `);width:150px;height:150px;border-radius:50%;cursor: pointer;background-position:  0px 0px;outline: 0;border: 0px;" /></div>
		<br><br>
		` + textt + `
		</form>
		</body>
		</html>`
	})

	m.Post("/add_ch", func(w http.ResponseWriter, req *http.Request) {
		if _, err := os.Stat(namefilech); os.IsNotExist(err) {
			write_channels_file()
		}
		inidl := ""
		body, _ := ioutil.ReadAll(req.Body)
		v, _ := url.ParseQuery(string(body))
		for key, value := range v {
			if key == "channel_name" {
				inidl = value[0]
			}
		}
		if info_ch_(inidl, key) != "" {
			add_channel(inidl)
		}
		http.Redirect(w, req, "/", 302)
	})

	m.Post("/edit_ch", func(w http.ResponseWriter, req *http.Request) {
		if _, err := os.Stat(namefilech); os.IsNotExist(err) {
			write_channels_file()
		}
		inidl := ""
		inidl_hide := ""
		body, _ := ioutil.ReadAll(req.Body)
		v, _ := url.ParseQuery(string(body))
		for key, value := range v {
			if key == "channel_name" {
				inidl = value[0]
			}
			if key == "channel_name_hide" {
				inidl_hide = value[0]
			}
		}
		if info_ch_(inidl, key) != "" {
			del_channel(inidl, inidl_hide)
		}
		http.Redirect(w, req, "/", 302)
	})

	m.Get("/", func() string {
		if _, err := os.Stat(namefilech); os.IsNotExist(err) {
			write_channels_file()
		}
		text := ""
		text2 := ""
		pos := 0
		for _, val := range read_channels() {
			text += "<option>" + val + "</option>"
			pos += 350
			//onmousedown="mf(this)"
			text2 += `<div id="idd" class="avatar20" style="left:` + strconv.Itoa(pos) + `px"><input name="fileini" class="avatar333" type="submit" value="` + val + `" style="float:left;background-size: cover;background-image: url(` + info_ch_(val, key) + `);width:150px;height:150px;border-radius:50%;cursor: pointer;background-position:  0px 0px;outline: 0;border: 1px;border-style: dotted;" /></div>`
		}
		return `
		<!DOCTYPE HTML>
	    <html>
	    <head>
	    <meta charset="utf-8">
	    <title>Youtube Download</title>
		` + styles + `
		<script>
		function mf_deff() {
		var divs = document.getElementsByTagName("input");
		for (var i = 0; i < divs.length; i++) {
		var node = divs[i];
		//alert(divs[i].innerHTML);
		var longpress = false;
		var presstimer = null;
		var longtarget = null;
		var cancel = function(e) {
		    if (presstimer !== null) {
		        clearTimeout(presstimer);
		        presstimer = null;
		    }
		    this.classList.remove("longpress");
		};
		var click = function(e) {
		    if (presstimer !== null) {
		        clearTimeout(presstimer);
		        presstimer = null;
		    }
		    this.classList.remove("longpress");
		    if (longpress) {
		        return false;
		    }
		    //alert("press");
		};
		
		var start = function(e) {
		    console.log(e);
		    if (e.type === "click" && e.button !== 0) {
		        return;
		    }
		    longpress = false;
		    this.classList.add("longpress");
		    presstimer = setTimeout(function() {
				editch(e.target.value);
		        longpress = true;
		    }, 800);
		    return false;
		};
		
		node.addEventListener("mousedown", start);
		node.addEventListener("touchstart", start);
		node.addEventListener("click", click);
		node.addEventListener("mouseout", cancel);
		node.addEventListener("touchend", cancel);
		node.addEventListener("touchleave", cancel);
		node.addEventListener("touchcancel", cancel);
		
		};
		}
		
		
		
function myFunction() {
	
	if (!document.getElementById('test_add_ch')){
		var inp = document.createElement("input");
		inp.setAttribute("type", "text");
		inp.setAttribute("id", "test_add_ch");
		inp.setAttribute("class", "chlabel");
		inp.setAttribute("name", "channel_name");
		document.getElementById("add_ch").appendChild(inp);
		inp.focus();
	} else {
		document.getElementById('test_add_ch').remove();
	}

	if (!document.getElementById('test_add_ch_but')){
		var but = document.createElement("button");
		but.setAttribute("name", "fileini");
		but.setAttribute("id", "test_add_ch_but");
		but.setAttribute("style", "cursor: pointer;");
		but.setAttribute("type", "submit");
		var textnode = document.createTextNode("Добавить");
		but.appendChild(textnode);
		document.getElementById("add_ch_but").appendChild(but);
	} else {
		document.getElementById('test_add_ch_but').remove();
	}
	
}

function editch(text_) {
	
	if (!document.getElementById('test_edit_ch')){
		var inp = document.createElement("input");
		inp.setAttribute("type", "text");
		inp.setAttribute("id", "test_edit_ch");
		inp.setAttribute("class", "chlabel");
		inp.setAttribute("name", "channel_name");
		inp.setAttribute("value", text_);
		
		document.getElementById("edit_ch1").appendChild(inp);
		inp.focus();
		var inp2 = document.createElement("input");
		inp2.setAttribute("type", "text");
		inp2.setAttribute("id", "test_edit_ch2");
		inp2.setAttribute("class", "chlabel");
		inp2.setAttribute("name", "channel_name_hide");
		
		inp2.setAttribute("value", text_);
		document.getElementById("edit_ch2").appendChild(inp2);
		
	} else {
		document.getElementById('test_edit_ch').remove();
		document.getElementById('test_edit_ch2').remove();
	}

	if (!document.getElementById('test_edit_ch_but')){
		var but = document.createElement("button");
		but.setAttribute("name", "fileini");
		but.setAttribute("id", "test_edit_ch_but");
		but.setAttribute("style", "cursor: pointer;");
		but.setAttribute("type", "submit");
		but.setAttribute("value", "del");
		var textnode = document.createTextNode("Удалить/Изменить");
		but.appendChild(textnode);
		document.getElementById("edit_ch_but").appendChild(but);
		
	} else {
		document.getElementById('test_edit_ch_but').remove();

	}
}
</script>
		</head>
	    <body onload="mf_deff()">
		<br><br><br>
		<form action="/run" method="post">
	    <!--<select name="fileini" size="1">
		` + text + `
        </select>-->
		` + text2 + `
		</form>
    	<a id="add_channel_but" onclick="myFunction()" style="float:left;width:150px;height:150px;border-radius:50%;cursor:pointer;font-size: 300%;background-size: cover;background-color:#7B68EE;"/><div class="add">+</div></a>
		<br><br><br>
		<form action="/add_ch" method="post">
		<div id="add_ch"></div>
		<div id="add_ch_but" class="avatar20"></div>
		</form>
		<form action="/edit_ch" method="post">
		<div id="edit_ch1"></div>
		<div id="edit_ch2" hidden="true"></div>
		<div id="edit_ch_but" class="avatar20"></div>
		<br><br><br>
		</form>
		<!--<button style="cursor: pointer;" type="submit" >Открыть последние 50 видео</button>-->
		<br><br><br>
		<!--<div class="chlabel">ПОЛУЧИТЬ ФАЙЛЫ НУЛЕВОГО РАЗМЕРА<div>
	    <div class="switch chboxydl"><input type="checkbox" name="nulldl" value="null" checked><label><i></i></label></div>-->
		</form>
    	</body>
		</html>`
	})
	http.Handle("/", m)
	m.RunOnAddr(":1234")
}
