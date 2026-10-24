# 老库 3D 模型（STEP）— 嘉立创获取

来源：`F:\Projects\PCB\Library\Symbols\*.OLB`（老 OrCAD 库，只读扫描）
匹配：型号 → JLCPCB 元件库 API → LCSC 编号 → EasyEDA 官方 3D 模型 → STEP (AP214)

## 结果
- 器件（有具体型号且拿到模型）：**40** 个
- 跳过（通用符号无单一型号 / 厂商无 3D）：**20** 类
- STEP 文件：**79** 个（40 个按老库器件名 `<device>.step` + 39 个按 LCSC 编号 `C<code>.step` 副本）

## 命名与用法
- `<device>.step` 的 device = 老库器件名（多单元取公共前缀，如 AFE5818、STM32G431Vx），
  以后你的 Allegro 封装 `.dra` 与器件同名时，可直接：
  `python attach_3d.py --dra-dir <dra目录> --steppath F:\Projects\Obsidian\Obsidian\_llm\mcp-servers\orcad-lib-gen\out\oldlib_3d_final`
- `report.csv`：每个老库器件的匹配结果（LCSC/MPN/封装/是否成功）
- `manifest.json`：同样的结构化清单

## 跳过的通用符号（无单一型号，故无专属 3D）
CAP / RES / LED / FUSE / INDUCTOR / FERRITEBEAD / CMC / CRYSTAL / MOSFET / SWITCH / TVS / M / FPC / XT30 / TF / TH / HDMI-D 等
如需这些也用“代表性型号”的 3D，请指定型号，我可补。
