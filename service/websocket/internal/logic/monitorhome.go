package logic

import (
	"github.com/valyala/fasthttp"
	"html/template"
)

func (l *Websocket) MonitorHome(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	monitorHomeTemplate.Execute(ctx, "ws://"+string(ctx.Host())+"/")
}

var monitorHomeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
    document.addEventListener("DOMContentLoaded", fetchLoginToken);
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}",token.value);
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
		
        print("SEND: " + input.value);
		ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});

 function fetchLoginToken() {
      // 登录API地址  
      const loginApiUrl = "http://10.132.105.54:9005/api/1.0/Ptm/tpmt/user/login";
      // 登录参数  
      const loginParams = {
        "account": "admin",
        "password": "adminpublic"
      };
      fetch(loginApiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(loginParams)
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json(); // 解析JSON响应  
        })
        .then(res => {
          // 假设接口返回的内容是要填充到input的token值  
          const tokenInput = document.getElementById("token");
          tokenInput.value = res.data.accessToken; // 假设返回的数据中有一个token字段  
        })
        .catch(error => {
          console.error('There has been a problem with your fetch operation:', error);
          alert("无法获取登录令牌：" + error.message);
        });


    };
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.

<p>
<p>监测点<p>
<p>monitor|[{"id":1},{"id":2}]|<p>


<form>
<p>token<p>
<p><input id="token" type="text" value=""><p>


<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
