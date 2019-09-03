// hello.go
package main

//void SayHello(const char* s);
import "C"
import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"strings"
)
const XXX_MAIL_TEMPLATE = `    <div>
        <h3>123</h3>
        <p>3456</p>
        <h3>789</h2>
        <table style="border-collapse:collapse;border: 1px solid black;">
            <thead style="border-collapse:collapse;border: 1px solid black;">
                <tr style="border-collapse:collapse;border: 1px solid black;text-align: center;">
                    <th style="border-collapse:collapse;border: 1px solid black;">Case Name</th>
                    <th style="border-collapse:collapse;border: 1px solid black;">Owner</th>
                    <th style="border-collapse:collapse;border: 1px solid black;">Creator</th>
                    <th style="border-collapse:collapse;border: 1px solid black;">Status</th>
                </tr>
            </thead>
            <tbody>
            {{with .Job}}{{range $k, $v := .Cases}}
                <tr style="border-collapse:collapse;border: 1px solid black;text-align: center;">
                    <td style="border-collapse:collapse;border: 1px solid black;">{{$v.Name}}         </td>
                    <td style="border-collapse:collapse;border: 1px solid black;">{{$v.IsSuccess}}          </td>
                    <td style="border-collapse:collapse;border: 1px solid black;">{{$v.Agent}}               </td>

                </tr>
             {{end}}
             {{end}}
            </tbody>
        </table>

    </div>`
func main() {
	//C.SayHello(C.CString("Hello, World\n"))
	//temp := PageInfo{Job: &job}
	MAIL_TEMPLATE := XXX_MAIL_TEMPLATE
	m := gomail.NewMessage()
	m.SetHeader("From", "ckc_it@163.com")
	m.SetHeader("To", strings.Split("943104990@qq.com,476688386@qq.com", ",")...)//send email to multipul persons
	m.SetHeader("Subject", "Hello!")
	t, err := template.New("mail summary template").Parse(MAIL_TEMPLATE)
	if err != nil {
		return
	}
	buffer := new(bytes.Buffer)
	t.Execute(buffer, "")
	m.SetBody("text/html", buffer.String())
	d := gomail.Dialer{Host: "smtp.163.com", Port: 465, Username: "ckc_it@163.com", Password: "liangge666",SSL:true}
	if err := d.DialAndSend(m); err != nil {
		return
	}
	return
}
