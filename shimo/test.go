package main

import (
	"fmt"
	"regexp"
	"strings"
	"net/http"
	"os"
	"io"
)

func main() {
	md := `
	![图片](https://uploader.shimo.im/f/qRIcjSqoFwErgYwB.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)


	>大家好, 我是可爱的排骨 😝
	# 前言
	
	近几年威联通用猛力发展针对中小企业和家庭用户的网络设备和解决方案, 为解决用户 NAS 设备的网络带宽瓶颈而不断推出千兆以上的**802.3bz (NBASE-T)**标准的**2.5/5/10GbE**网络设备. 不只是快速升级的 Wi-Fi 规格需要突破千兆以太网限制, 网络存储吞吐量同样需要, 这就离不开超千兆的路由器交换机配套支持.
	
	领先于其它友商, 威联通陆续上线了支持 NBASE-T 的全线新品.
	
	⚫ 网络存储: TS-X53D 以及 TS-X73A 等内置 2.5GbE 的高规格的 NAS 产品.
	
	⚫ 交换机: QSW-1105-5T 五口 2.5GbE, QSW-2104-2T 双口 10GbE + 五口 2.5GbE 交换机.
	
	⚫ 路由器: QHora-301W 双口 10GbE + AX3600 双频路由器
	
	⚫ 网卡: QNA-UC5G1T USB 5GbE, QNA-T310G1T 雷电3 10GbE, QXG-10G1T PCIe 3.0 x4 10GbE 有线网卡
	
	**注**: 老的万兆网卡并不能支持 NBASE-T (2.5/5GbE) 标准, 只能在 10/1GbE 两个档切换.
	
	去年底看到群里有基友晒 QHora-301W 这款路由器, 排骨软磨硬泡了几个月终于借 (qiang) 了过来.  排骨为啥会对 QNAP 的这款 QHora-301W 路由器感兴趣? 因为它是市面上很少见的**双万兆电口**无线路由器.
	
	QHora-301W 采用高通 IPQ8072A 64 位 4 核 处理器, 配备 1GB 内存和 4GB 闪存,
	
	![图片](https://uploader.shimo.im/f/5QSFK4O7R1e6oxQv.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)
	
	
	支持最多 6 组 VAP (虚拟 AP), 可对应不同的 VLAN, 支持不同的防火墙规则.
	
	![图片](https://uploader.shimo.im/f/MYz2BfRrNlyrzeEJ.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)有暗色模式？這款瀏覽器外掛應該適合你：Dark ...https://www.mdeditor.tw › zh-tw Translate this page 10 Mar 2020 — ... 裝置設定暗色主題，其實這也就是我們所說的深色模式——Dark Mode。 ... 重要問題，而使用黑色背景圖或深色模式也能夠有效延緩燒屏問題的發生。 ... 的效果已經相當不錯，比如少數派的線上編輯器、石墨文件這類線上文件 ...
	
	![图片](https://uploader.shimo.im/f/5vx95rS1j1B2Hsb1.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/tB0m8q7UDxbVIZVP.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/Jcx0fvMwgjRVOyKj.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/K7cNitMB2F7W72KA.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/gDvNFSwaMyludOzx.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/6KA3fDfkFHtS07f7.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/PMcQrZr33hbKG5nE.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/omxHDKpDsDkJKj0U.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/R0KJZc3Dtul70xlW.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/RpdLCU4vHTLT7fTq.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/Oyj1S8REn684BBju.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/4CxnjnvcEod51hrJ.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)
	
	
	![图片](https://uploader.shimo.im/f/zJMCHvgx047EQUCp.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/u64ssZFlJhFjDA5D.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/tfdFkHQQf8MVJJWf.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)
	
	![图片](https://uploader.shimo.im/f/X3RjAGrPH2UIsM2T.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/sV4wfR3912nseagJ.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/a6gzJJi7VPTSWtfy.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/l0xhMP1k6UOe7XRw.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/fnCQLDLzUQZsiFmJ.PNG!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)
	![图片](https://uploader.shimo.im/f/5dRON0fOh1x0dP30.jpg!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/GhS2yDTKquSMP9eN.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)![图片](https://uploader.shimo.im/f/qi7j0yL7GDZUEcho.png!thumbnail?fileGuid=vDCcrg3WWhjwhGX6)
`
	md = strings.Replace(md, ")![", ")\n![", -1)
	re := regexp.MustCompile(`!\[(.*)\]\((https://uploader.shimo.im/.*)\)`)
	matches := re.FindAllStringSubmatch(md, -1)


	for _, match := range matches {
		url := match[2]
		url = url[0:strings.Index(url, "!thumbnail")]
		fileName := url[strings.LastIndex(url, "/") + 1:]
		url = url + "?accessToken=" + "eyJhbGciOiJIUzI1NiIsImtpZCI6ImRlZmF1bHQiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE2MjAyMzMzMTYsImciOiJ2RENjcmczV1doandoR1g2IiwiaWF0IjoxNjIwMjMxNTE2LCJ1IjoyNzIwOTE0NX0.3lp06H32wsMasAvRGDXryaWXW9JL8r4C9E7aVD3fF14"
		fmt.Println("pulling ", fileName)
	  Download(url, fileName)
	}
}

func Download(url string, filePath string) error {
	file, err := os.Create(filePath)

	if err == nil {
		defer file.Close()

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Referer", "https://shimo.im")
		res, err := client.Do(req)
	
		if err == nil {			
			defer res.Body.Close()
			_, err = io.Copy(file, res.Body)
		}
	}
	
	return err
}