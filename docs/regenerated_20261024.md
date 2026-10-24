# 2026-10-24 生成的 13 颗元件修复报告

修复项：D1 名称不一致 / D2 名称带立创编号 / D3 datasheet 指错 / D4 分类归错 / D5 引脚全在一侧。

## 命名规则（唯一真源 `easyeda_lib.clean_symbol_name`）
- 型号去尾部 `_Cxxxxx` 编号，非法字符 → `_`
- CIS `Symbol_Name` 与 OLB 包名使用同一个名字

## 重建后的符号（全部按用户确认的版式）

| 型号 | 库 | 版式 | CIS 表 |
|---|---|---|---|
| LM3481MM_NOPB | Power_Management | 左 VIN/DR/FA/UVLO/ISEN；右 VCC/FB/COMP/AGND/PGND | IC_Params |
| LT3580IMS8E_PBF | Power_Management | 左 VIN/VC/SHDN#/RT/SYNC；右 FB/SW/SS/EP | IC_Params |
| TLV171IDBVR | Analog_SignalChain | 左 IN+/IN-；右 V+/OUT/V- | IC_Params |
| TLV3201AIDBVR | Analog_SignalChain | 左 IN+/IN-；右 V+/OUT/V- | IC_Params（从 Misc 移入） |
| TPS62840DLCR | Power_Management | 左 VIN/EN/MODE/VOS；右 SW/VSET/STOP/GND | IC_Params |
| TPS3808G33DBVR | Power_Management | 左 VDD/MR#/CT；右 RESET#/SENSE/GND | IC_Params（从 Misc 移入） |
| TPS7A8001DRBR | Power_Management | 左 IN/IN/EN/NR；右 OUT/OUT/FB/GND/EP | IC_Params |
| SN74LVC1G74DCUR | Interface | 左 D/CLK/CLR#/PRE#；右 VCC/Q/Q#/GND | IC_Params（从 Misc 移入） |
| CSD19532Q5B | Discrete_Passive | 左 G；右 D×5/S×3 | Transistor_Params |
| DAC80502DRXR | Analog_SignalChain | 左 数字接口×5；右 VDD/VREFIO/VOUTA/VOUTB/AGND | IC_Params |
| INA4290A1IRGVR | Analog_SignalChain | 左 4 通道 IN+/IN-；右 VS×2/OUT×4/GND×2/EP | IC_Params（从 Misc 移入） |
| TPS7A4700RGWR | Power_Management | 左 IN×2/EN/NR/NC×4；右 OUT×2/FB/VSET×9/GND/EP | IC_Params |
| DP83867IRPAPT | Interface | **2 单元**：A 电源/管理/时钟 31 脚；B PHY 信号 34 脚 | IC_Params |

## Datasheet

以下 4 颗 LCSC 给的链接指错器件，已回退为 LCSC 搜索页（需人工补 PDF）：
DP83867IRPAPT、TLV171IDBVR、CSD19532Q5B、SN74LVC1G74DCUR

## 校验

- 15/15 生成行（含 RP2040、TX7316）在 OrCAD 库中都能找到对应包名 ✅
- 4 颗 Miscellaneous → IC_Params 迁移完成 ✅
- 封装未动（由用户手动入库）
