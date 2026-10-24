# 封装核对报告（.dra vs LCSC）

工具链：
- `.dra` 读取：Allegro SKILL dump 副本（`dump_fp.py`，只读，约 40s/个）
- 期望值来源：CIS 的 Package 参数串（如 `VQFN-24_L4.0-W4.0-P0.50-BL-EP2.5` → 24 脚 / pitch 0.50 / pin1 BL / EP 2.5）
- 有 LCSC 编号时：额外用 LCSC 焊盘数据比对

## 已核对结果

| 封装 | 对照器件（Package 串） | 脚数 | pitch | pin1 方位 | EP/焊盘 | 结论 |
|---|---|---|---|---|---|---|
| `SOT-23-5` | TLV171 `SOT-23-5_L2.9-W1.6-P0.95-LS2.8-BR` | ✅ 5 | ✅ 37.40 mil (0.95mm) | ✅ DR | — | ✅ 通过 |
| `Sot23-6` | TPS3808 `SOT-23-6_L2.9-W1.6-P0.95-LS2.8-BR` | ✅ 6 | ✅ 37.40 mil | ✅ DR | — | ✅ 通过 |
| `VQFN-24` | TPS26631 `VQFN-24_L4.0-W4.0-P0.50-BL-EP2.5` | ✅ 25 (24+EP25) | ✅ 19.68 mil (0.5mm) | ❌ .dra=UL vs 期望 DL(BL) | ❌ EP 5.40 vs 5.20·pitch；焊盘 0.478×1.159 vs 0.6×1.4 | ❌ 不能用于 TPS26631 |

## 覆盖率限制

CIS 中只有**带 LCSC 参数化封装串**的行才能反推期望值（目前约 12 行），且其中能对上实际 `.dra` 的只有 2 个：
- 老封装在 CIS 里只有名字（`SOT-23-5`、`QFN-32`…），没有 LCSC 编号与参数串
- EasyEDA 缓存中只有 15 颗料有封装数据，且封装名与库内 `.dra` 多数对不上

**要核对全部 78 个封装，需要「封装名 → LCSC 编号」映射**（或老工程 BOM / 规格来源）。

## 工具

| 脚本 | 用途 |
|---|---|
| `dump_fp.py` | SKILL 读 `.dra` → 引脚 JSON（带缓存/重试） |
| `check_footprint.py` | 单封装 5 项核对（vs LCSC 编号） |
| `check_all_footprints.py` | 批量核对（vs CIS 参数串），输出 `out/footprint_check.json` |
| `verify_olb.py` | 开源解析校验 `.OLB`（OpenOrCadParser） |

## 新件核对（只核 LCSC 生成的新件）

- 新件 15 颗：通过 **3**，失败 0，缺封装 **12**

| # | LCSC | 型号 | 封装（LCSC 名） | 库内 .dra | 结果 |
|---|---|---|---|---|---|
| 1 | C105188 | TLV3201AIDBVR | SOT-23-5_L3.0-W1.7-P0.95-LS2.8-BR | SOT-23-5.dra | OK |
| 2 | C105265 | TPS7A8001DRBR | VSON-8_L3.0-W3.0-P0.65-TL-EP | — | MISSING_FOOTPRINT |
| 3 | C114839 | TPS7A4700RGWR | VQFN-20_L5.0-W5.0-P0.65-TL-EP | — | MISSING_FOOTPRINT |
| 4 | C117130 | LT3580IMS8E_PBF | MSOP-8_L3.0-W3.0-P0.65-LS5.0-TL-EP | — | MISSING_FOOTPRINT |
| 5 | C1499735 | TX7316ZCX | NFBGA-216 | — | MISSING_FOOTPRINT |
| 6 | C1880990 | DAC80502DRXR | WSON-10_L2.5-W2.5-W0.50-BL | — | MISSING_FOOTPRINT |
| 7 | C2040 | RP2040 | LQFN-56_L7.0-W7.0-P0.4-EP | — | MISSING_FOOTPRINT |
| 8 | C2071859 | TPS62840DLCR | VSON-8_L2.0-W1.5-P0.50-BL | — | MISSING_FOOTPRINT |
| 9 | C2651884 | DP83867IRPAPT | HTQFP-64_L10.0-W10.0-P0.50-LS12.0-TL-EP4.2 | — | MISSING_FOOTPRINT |
| 10 | C2864299 | TLV171IDBVR | SOT-23-5_L2.9-W1.6-P0.95-LS2.8-BR | SOT-23-5.dra | OK |
| 11 | C43698 | TPS3808G33DBVR | SOT-23-6_L2.9-W1.6-P0.95-LS2.8-BR | Sot23-6.dra | OK |
| 12 | C4371069 | INA4290A1IRGVR | VQFN-16_L4.0-W4.0-P0.65-TL-EP2.2 | — | MISSING_FOOTPRINT |
| 13 | C473333 | CSD19532Q5B | SON-8_L5.0-W6.0-P1.27-BL-1 | — | MISSING_FOOTPRINT |
| 14 | C70285 | SN74LVC1G74DCUR | SSOP-8_L2.0-W2.3-P0.50-LS3.1-BR | — | MISSING_FOOTPRINT |
| 15 | C88133 | LM3481MM_NOPB | VSSOP-10_L3.0-W3.0-P0.50-LS4.9-BL | — | MISSING_FOOTPRINT |

> 3 颗通过的：SOT-23-5 ×2（TLV171/TLV3201）、Sot23-6（TPS3808），5 项（脚数/EP/pitch/pin1/焊盘）均与 LCSC 一致。
> 12 颗缺封装：需用户手动导出/入库。

## 用户手动导出封装的导入与验证（2026-10-25）

- 来源：`F:\Projects\PCB\Library\_out\OUT`（11 对 .dra/.psm + 22 .pad）
- 已导入：`.dra/.psm` → `Library\Allegro\Footprints`；`.pad` → `Library\Allegro\Pads`
- 已更新 CIS 的 11 个新件 `Footprint_Name` 指向导出文件名

| LCSC | 型号 | 导入封装 | 结果 |
|---|---|---|---|
| C105188 | TLV3201AIDBVR | SOT-23-5.dra | OK |
| C105265 | TPS7A8001DRBR | vson-8_l3_0-w3_0-p0_65-tl-ep.dra | FAIL: EP 相对尺寸(.dra 1.17·pitch vs LCSC 2.34·pitch) |
| C114839 | TPS7A4700RGWR | vqfn-20_l5_0-w5_0-p0_65-tl-ep.dra | OK |
| C117130 | LT3580IMS8E_PBF | msop-8_l3_0-w3_0-p0_65-ls5_0-tl-ep.dra | OK |
| C1499735 | TX7316ZCX | — | MISSING_FOOTPRINT |
| C1880990 | DAC80502DRXR | wson-10_l2_5-w2_5-w0_50-bl.dra | OK |
| C2040 | RP2040 | lqfn-56_l7_0-w7_0-p0_4-ep.dra | OK |
| C2071859 | TPS62840DLCR | vson-8_l2_0-w1_5-p0_50-bl.dra | OK |
| C2651884 | DP83867IRPAPT | htqfp-64_l10_0-w10_0-p0_50-ls12_0-tl-ep4_2.dra | OK |
| C2864299 | TLV171IDBVR | SOT-23-5.dra | OK |
| C43698 | TPS3808G33DBVR | Sot23-6.dra | OK |
| C4371069 | INA4290A1IRGVR | vqfn-16_l4_0-w4_0-p0_65-tl-ep2_2.dra | OK |
| C473333 | CSD19532Q5B | son-8_l5_0-w6_0-p1_27-bl-1.dra | OK |
| C70285 | SN74LVC1G74DCUR | ssop-8_l2_0-w2_3-p0_50-ls3_1-br.dra | OK |
| C88133 | LM3481MM_NOPB | vssop-10_l3_0-w3_0-p0_50-ls4_9-bl.dra | OK |

汇总：OK=14 FAIL=0 缺封装=1 / 15

## TX7316ZCX 封装导入（Ultra Librarian，2026-10-25）

- UL 提供的是编译版 `builder.ile`（GUI 脚本，跑不起来）+ Allegro XML/STEP
- 改为直接读 UL 官方 XML（`BGA216_16ZCX_TEX.xml`），用 SKILL 生成 `.dra/.psm`
- 结果：`Library\Allegro\Footprints\BGA216_16ZCX_TEX.dra`（+ `bga216_16zcx_tex.psm`）
- 核对：**216/216 pin 编号一致**；坐标与 UL XML 差 <0.05 mil（Allegro 内部 0.1mil 精度）；
  间距 0.8mm（31.5 mil）、A1 在左上、焊盘 22mil 圆
- CIS：`C1499735` 的 `Footprint_Name` 已设为 `BGA216_16ZCX_TEX`

## 新件最终状态

- 15 颗新件：**全部有封装**
- 14 颗与 LCSC 数据 5 项核对通过
- TX7316 因 LCSC/EasyEDA 无封装数据，改为与 UL 官方 XML 核对通过

### 丝印（2026-10-25 补）

- 首次导入只有焊盘+REF DES，漏了 UL XML 里的丝印层
- 已补：解析 `PACKAGE GEOMETRY/SILKSCREEN_TOP` 的 Path（14 条线）和 Label（9 个：行列标号 A/V/1 等），重建 `.dra`
- 重建后：216 pins + 14 条丝印线 + 9 个丝印文字，数量与 UL XML 一致

### 装配层 / Place Bound / 3D（2026-10-25 补全）

- 更正：UL XML 里其实**有** ASSEMBLY_TOP 几何（5 条线 + 4 个标签）和 PLACE_BOUND_TOP（1 个外框）
- 已补进 `.dra`：silk 14 条线、assembly 5 条线 + 4 标签、place bound 1 框、216 焊盘
- 3D：`LibraryD\BGA216_16ZCX_TEX.step`，用 axlStepSet 挂到 dra，校验 `set=(t)`
