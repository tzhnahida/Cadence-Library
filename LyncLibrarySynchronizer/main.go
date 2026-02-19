package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	LoadConfig()
	reader := bufio.NewReader(os.Stdin)
	ai := NewAiClient()

	fmt.Println("🚀 PCB Library 助手已就绪 (输入 'exit' 退出)")
	
	for {
		fmt.Print("\n🔗 URL: ")
		url, _ := reader.ReadString('\n')
		url = strings.TrimSpace(url)
		if url == "exit" { break }
		if url == "" { continue }

		raw, id, err := fetchLCSC(url)
		if err != nil { fmt.Println("❌ 抓取失败:", err); continue }

		parsed, err := ai.Ask(fmt.Sprintf("编号: %s\n内容: %s", id, raw))
		if err != nil { fmt.Println("❌ AI 错误:", err); continue }

		fmt.Println("\n--- AI 建议 ---")
		fmt.Printf("表: [%s]\n", parsed.TableName)
		for k, v := range parsed.Fields { fmt.Printf("%-15s: %v\n", k, v) }

		fmt.Print("\n✍️ Symbol [回车保持]: ")
		s, _ := reader.ReadString('\n')
		if s = strings.TrimSpace(s); s != "" { parsed.Fields["Symbol_Name"] = s }

		fmt.Print("✍️ Footprint [回车保持]: ")
		f, _ := reader.ReadString('\n')
		if f = strings.TrimSpace(f); f != "" { 
			parsed.Fields["Footprint_Name"] = f 
			parsed.Fields["Package"] = f 
		}

		fmt.Print("❓ 写入数据库? (y/n): ")
		conf, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(conf)) == "y" {
			newID, err := saveToAccess(parsed)
			if err != nil {
				fmt.Println("❌ 写入失败:", err)
			} else {
				fmt.Printf("✅ 成功! Part_ID: %d\n", newID)
				fmt.Println("【System Tag】\n", GlobalConfig.SystemTag)
			}
		}
	}
}