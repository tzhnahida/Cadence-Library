# CIS 数据库报告（Components_Pro.mdb）

- 总行数：**13928**（13 张表）
- 原始库：`OldLibraryackup_20261024_1751\CIS\Components.mdb`

## 各表行数

| 表 | 行数 | 清理时更新行 | 删除重复 |
|---|---|---|---|
| Capacitor_Params | 6912 | 1319 | 0 |
| Connector_Params | 15 | 7 | 0 |
| Crystal_Params | 4 | 4 | 0 |
| Diode_Params | 14 | 10 | 0 |
| FerriteBead_Params | 2 | 1 | 0 |
| Fuse_Params | 6 | 4 | 0 |
| IC_Params | 38 | 16 | 0 |
| Inductor_Params | 6 | 1 | 0 |
| LED_Params | 1 | 1 | 0 |
| Miscellaneous_Params | 6 | 4 | 0 |
| Resistor_Params | 6921 | 6646 | 0 |
| TMPPRTS | 0 | - | - |
| Transistor_Params | 3 | 0 | 0 |

## 本次修复

- 全部字段去首尾空格/制表符
- Datasheet 粘连格式（`X#http://X#`）→ 干净文件名/URL
- 分类命名统一：`IC - …` / `Connector - …` / `Diode - …` / `Capacitor - MLCC` / `Fuse - Resettable` / `Crystal - Passive`
- 精确重复行删除（IC 1 行、Resistor 1 行）
- 符号名修正 7 行（如 `AFE5818ZBV→AFE5818`、`LMR38020SDDAR→LMR38020`、`INA226→INA226AIDGSR`）

## Resistor / Capacitor 网格（重新生成）

- **Resistor_Params：6921 行** = E96（96 值）× 8 个数量级（1Ω–10MΩ）+ 0Ω 跳线 × 9 种封装（R01005/R0201/R0402/R0603/R0805/R1206/R1210/R1812/R2512）
- **Capacitor_Params：6912 行** = E96（96 值）× 8 个数量级（1pF–100µF）× 9 种封装（C01005…C2220）
- 电阻：±1%、功率按封装、温漂 ±100ppm/℃；0Ω 跳线每封装 1 行
- 电容：介质按容值（≤1nF → C0G，其余 X7R），耐压按封装档位（01005=16V / 0201=25V / 其余 50V）
- 旧库中 (封装,值) 能对上的 Part_ID 已保留，避免已有原理图引用错位；新行使用递增编号

## 未解决的 Symbol_Name（库里没有对应符号，需补建或改名）

| 表 | Part_ID | Symbol_Name |
|---|---|---|
| IC_Params | 24 | MP4247GQV |
| IC_Params | 26 | BQ25758RRVR |
| IC_Params | 27 | HUSB320_BA000_QN12R |
| IC_Params | 28 | TPS55288RPMR |
| IC_Params | 29 | TPS54302 |
| IC_Params | 30 | LM5050 |
| IC_Params | 31 | MAX706 |
| IC_Params | 32 | SN74LVC1G14 |
| IC_Params | 33 | SN74AHC1G08 |
| Miscellaneous_Params | 4 | HS20HS072RX |
| Transistor_Params | 3 | NPN_3P |

## 封装完整性

- `.dra` 78 个 / `.psm` 79 个；多余 `.psm`：`m3`（M3 安装孔机械符号，缺 .dra）
- 其余全部成对
