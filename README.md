# Library — OrCAD / Allegro 元件库

本目录是 PCB 设计用的统一元件库（OrCAD Capture + Allegro），2026-10-24 重构为下列结构。
旧库完整备份见 `OldLibrary\backup_20261024_1751`，替换前的旧目录见 `OldLibrary\replaced_at_20261024`（均只读）。

## 目录结构

```
Library\
├─ OrCAD\            6 个分类符号库（.OLB，官方 XML 原样搬运）
│    MCU_Memory / Power_Management / Interface /
│    Analog_SignalChain / Connectors / Discrete_Passive
├─ CIS\              Components_Pro.mdb + TZH_ECAD_CIS_ACCESS_V1.DBC
├─ Allegro\
│   ├─ Footprints\   .dra / .psm
│   ├─ Pads\         padstack
│   └─ Vias\
├─ Datasheet\        数据手册 PDF（含 _Candidate）
├─ 3D\               STEP 模型（<器件名>.step）
├─ docs\             使用手册 / 库报告 / CIS 报告 / 映射表
├─ OldLibrary\       历史备份（backup_* / replaced_at_* / oldlib_26-1）
├─ TZH_ECAD_CIS_ACCESS_V1.DBC   CIS 配置文件（DSN 名 CIS）
├─ config.toml / LICENSE / allegro_L4.prm / allegro_L6.prm
└─ README.md / .gitignore
```

## 各目录说明

| 目录 | 内容 | 配置 |
|---|---|---|
| `OrCAD\` | 6 个功能库：MCU_Memory / Power_Management / Interface / Analog_SignalChain / Connectors / Discrete_Passive | Capture 库路径加 `Library\OrCAD` |
| `CIS\` | 清理后的 CIS 数据库（R/C 为 E96×全封装重生）+ DBC | 32 位 ODBC DSN `CIS` → `CIS\Components_Pro.mdb` |
| `Allegro\Footprints\` | 封装 `.dra/.psm` | `psmpath` |
| `Allegro\Pads\` | padstack | `padpath` |
| `Allegro\Vias\` | 过孔定义 | — |
| `Datasheet\` | 数据手册 | — |
| `3D\` | STEP 模型，命名 `<器件名>.step` | `steppath` |
| `docs\` | **[使用手册](docs/使用手册.md)** / library_report / cis_report / mapping | — |
| `OldLibrary\` | 历史备份，**只读** | — |

## 使用入口

- **使用手册（推荐先读）**：[docs/使用手册.md](docs/使用手册.md)
  - Capture / Allegro 配置、ODBC 设置
  - 新元件入库（MCP `ingest_lcsc_part` / `python pro_ingest.py <LCSC>`）
  - AI 定制建符号（单元/分组/组间留空位）
  - CIS 体检/清理/R-C 重生、回滚、故障排查
- 库报告：[docs/library_report.md](docs/library_report.md)
- CIS 报告：[docs/cis_report.md](docs/cis_report.md)

## 维护规则

1. 新增/修改只动本结构；`OldLibrary\` 只读，永不修改。
2. 新符号进 `OrCAD\<对应分类>.OLB`；新元件入库用 `ingest_lcsc_part`。
3. CIS 新料号用 `ingest_lcsc_part` 自动写入；不要手工改 `Part_ID`。
4. R/C 网格由 `cis_regen_passives.ps1` 生成（E96 × 全封装），改规则请改脚本重跑。
5. 3D 模型命名 `<器件名>.step`，与封装 `.dra` 同名时可用 `attach_3d.py` 自动挂。
6. 大目录（`OldLibrary/`、`3D/`）不进 git；Git 由用户自行提交。
7. **封装由用户手动导出/入库**；MCP 不生成封装、不跑 Allegro 脚本。3D 挂载用 `attach_3d_to_dra` 一次完成。

## 附：自动化工具

- `F:\Projects\PCB\Tools\LyncLibrarySynchronizer`：OrCAD CIS 自动填参数工具（Go，AI 辅助生成）。
- MCP 服务 `orcad-lib`（`F:\Projects\Obsidian\Obsidian\_llm\mcp-servers\orcad-lib-gen\orcad_mcp.py`）：18 个工具，覆盖入库、CIS、3D、AI 定制建符号。
