# Library Overview (OrCAD / Allegro)

This directory contains component libraries and related resources for PCB design and simulation, intended for use with **OrCAD** and **Allegro** platforms.

## Directory Structure

```
Library
├─ ComponentLibrary             # OrCAD component library management
│  └─ allegro                  # Allegro project resources linked by OrCAD
├─ Datasheet                   # Component datasheets
├─ Footprints                  # PCB footprint libraries
├─ Pads                        # PCB pad libraries
├─ Symbols                     # Schematic symbols library (shared between OrCAD/Allegro)
├─ Vias                        # PCB via definitions
├─ Components.mdb              # OrCAD CIS component database
├─ TZH_ECAD_CIS_ACCESS_V1.DBC  # OrCAD CIS database access file
├─ TZH_ECAD_CIS_ACCESS_V1.DBCBAK # Backup of the CIS database
└─ .gitignore                  # Git ignore file
```

## Folder Descriptions

### 1. ComponentLibrary (OrCAD Management)
- Contains component libraries managed by OrCAD, including schematic symbols, footprints, pads, and vias.  
- Resources are accessible directly through OrCAD CIS or library management tools.

### 2. ComponentLibrary/allegro (OrCAD-linked Allegro Project)
- Contains Allegro project resources referenced by OrCAD to maintain consistency between OrCAD and Allegro designs.  
- Includes:
  - `devices`: Allegro device and footprint libraries  
  - `symbols`: Allegro schematic symbols  

### 3. Datasheet
- Contains datasheets for all components for easy reference of electrical parameters and package details.

### 4. Footprints / Pads / Vias
- PCB layout resources including footprints, pads, and vias to ensure consistency between schematic and PCB design.

### 5. Symbols
- Schematic symbols library usable in both OrCAD Capture and Allegro Schematic.

### 6. CIS Database Files
- `Components.mdb`: OrCAD CIS component database for managing component attributes and footprint mappings.  
- `TZH_ECAD_CIS_ACCESS_V1.DBC`: CIS database access file for OrCAD.  
- `DBCBAK`: Backup of the CIS database to prevent accidental data loss.

### 7. .gitignore
- Git ignore file to exclude unnecessary files from version control.

> **Note:**  
> The `OldLibrary` folder contains historical files and can be ignored.
