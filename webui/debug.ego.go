// Generated by ego.
// DO NOT EDIT

//line debug.ego:1

package webui

import "fmt"
import "html"
import "io"
import "context"

import (
	"net/http"
	"runtime"

	"github.com/contribsys/faktory/client"
)

func ego_debug(w io.Writer, req *http.Request) {
	stats := ctx(req).Store().Stats()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

//line debug.ego:16
	_, _ = io.WriteString(w, "\n")
//line debug.ego:16
	ego_layout(w, req, func() {
//line debug.ego:17
		_, _ = io.WriteString(w, "\n\n<h3>")
//line debug.ego:18
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Debugging"))))
//line debug.ego:18
		_, _ = io.WriteString(w, "</h3>\n<div class=\"table_container\">\n  <table class=\"error table table-bordered table-striped\">\n    <tbody>\n      <tr>\n        <th>")
//line debug.ego:23
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Locale"))))
//line debug.ego:23
		_, _ = io.WriteString(w, "</th>\n        <td>\n          <select name=\"locales\" id=\"faktory_locale\">\n            ")
//line debug.ego:26
		sortedLocaleNames(req, func(locale string, current bool) {
//line debug.ego:27
			_, _ = io.WriteString(w, "\n              <option name=\"")
//line debug.ego:27
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(locale)))
//line debug.ego:27
			_, _ = io.WriteString(w, "\" ")
//line debug.ego:27
			if current {
//line debug.ego:27
				_, _ = io.WriteString(w, " selected ")
//line debug.ego:27
			}
//line debug.ego:27
			_, _ = io.WriteString(w, " > ")
//line debug.ego:27
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(locale)))
//line debug.ego:27
			_, _ = io.WriteString(w, " </option>\n            ")
//line debug.ego:28
		})
//line debug.ego:29
		_, _ = io.WriteString(w, "\n          </select>\n          <span style=\"font-size: small\">\n            Want to help us improve the translations?\n            <a href=\"https://github.com/contribsys/faktory/tree/master/webui/static/locales\">Submit a PR</a>.\n          </span>\n        </td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:37
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Version"))))
//line debug.ego:37
		_, _ = io.WriteString(w, "</th>\n      <td>Faktory ")
//line debug.ego:38
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(client.Version)))
//line debug.ego:38
		_, _ = io.WriteString(w, "</td>\n    </tr>\n    <tr>\n    <th>")
//line debug.ego:41
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Data Location"))))
//line debug.ego:41
		_, _ = io.WriteString(w, "</th>\n      <td>")
//line debug.ego:42
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(stats["name"])))
//line debug.ego:42
		_, _ = io.WriteString(w, "</td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:45
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Runtime"))))
//line debug.ego:45
		_, _ = io.WriteString(w, "</th>\n      <td>Goroutines: ")
//line debug.ego:46
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(runtime.NumGoroutine())))
//line debug.ego:46
		_, _ = io.WriteString(w, ", CPUs: ")
//line debug.ego:46
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(runtime.NumCPU())))
//line debug.ego:46
		_, _ = io.WriteString(w, "</td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:49
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Memory"))))
//line debug.ego:49
		_, _ = io.WriteString(w, "</th>\n      <td>\n        Alloc (KB): ")
//line debug.ego:51
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.Alloc/1024)))
//line debug.ego:51
		_, _ = io.WriteString(w, "<br/>\n        Live Objects: ")
//line debug.ego:52
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.Mallocs-m.Frees)))
//line debug.ego:53
		_, _ = io.WriteString(w, "\n        ")
//line debug.ego:53
		if amt := rss(); amt != "" {
//line debug.ego:54
			_, _ = io.WriteString(w, "\n        <br/>RSS: ")
//line debug.ego:54
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(amt)))
//line debug.ego:55
			_, _ = io.WriteString(w, "\n        ")
//line debug.ego:55
		}
//line debug.ego:56
		_, _ = io.WriteString(w, "\n      </td>\n    </tr>\n    <tr>\n      <th>")
//line debug.ego:59
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "GC"))))
//line debug.ego:59
		_, _ = io.WriteString(w, "</th>\n      <td>\n        PauseTotal (µs): ")
//line debug.ego:61
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.PauseTotalNs/1000)))
//line debug.ego:61
		_, _ = io.WriteString(w, "<br/>\n        NumGC: ")
//line debug.ego:62
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(m.NumGC)))
//line debug.ego:63
		_, _ = io.WriteString(w, "\n      </td>\n    </tr>\n  </tbody>\n</table>\n</div>\n\n<h3>")
//line debug.ego:69
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Redis Info"))))
//line debug.ego:69
		_, _ = io.WriteString(w, "</h3>\n<pre>\n")
//line debug.ego:71
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(redis_info(req))))
//line debug.ego:72
		_, _ = io.WriteString(w, "\n</pre>\n\n<h3>")
//line debug.ego:74
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Disk Usage"))))
//line debug.ego:74
		_, _ = io.WriteString(w, "</h3>\n<pre>\n<code>&gt; df -h</code>\n")
//line debug.ego:77
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(df_h())))
//line debug.ego:78
		_, _ = io.WriteString(w, "\n</pre>\n\n")
//line debug.ego:80
	})
//line debug.ego:81
	_, _ = io.WriteString(w, "\n")
//line debug.ego:81
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
